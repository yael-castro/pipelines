//go:build buffered

package logic

import (
	"context"
	"github.com/yael-castro/pipelines/internal/repository"
	"log"
	"sync"
)

const bufferSize = 500

func (l logic) CalculateProfit(ctx context.Context) (err error) {
	closings, err := l.GetClosings(ctx, repository.Open)
	if err != nil {
		return
	}

	closingCh := l.closingChannel(closings)      // STEP 0 (SOURCE)
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.costsPipeline(ctx, closingCh)  // STEP 2
	closingCh = l.profitPipeline(ctx, closingCh) // STEP 3
	doneCh := l.savePipeline(ctx, closingCh)     // STEP 4 (LAST)

	// Waiting to finish all processes
	<-doneCh

	return
}

func (logic) closingChannel(closings []repository.Closing) <-chan repository.Closing {
	closingsCh := make(chan repository.Closing)

	go func() {
		defer close(closingsCh)

		for i := range closings {
			closingsCh <- closings[i]
		}
	}()

	return closingsCh
}

func (l logic) salesPipeline(ctx context.Context, closingsCh <-chan repository.Closing) <-chan repository.Closing {
	salesCh := make(chan repository.Closing)

	go func() {
		defer log.Println("SALES PIPELINE IS DONE!")

		defer close(salesCh)

		var wg sync.WaitGroup
		semaphoreCh := make(chan struct{}, bufferSize)

		for closing := range closingsCh {
			semaphoreCh <- struct{}{}
			wg.Add(1)
			go func() {
				defer func() {
					<-semaphoreCh
					wg.Done()
				}()

				sales, err := l.GetSales(ctx, closing.Start, closing.End)
				if err != nil {
					return
				}

				closing.Sales = sales
				salesCh <- closing
			}()
		}

		wg.Wait()
	}()

	return salesCh
}

func (l logic) costsPipeline(ctx context.Context, closingsCh <-chan repository.Closing) <-chan repository.Closing {
	costsCh := make(chan repository.Closing)

	go func() {
		defer log.Println("COST PIPELINE IS DONE!")
		defer close(costsCh)

		var wg sync.WaitGroup
		semaphoreCh := make(chan struct{}, bufferSize)

		for closing := range closingsCh {
			semaphoreCh <- struct{}{}

			wg.Add(1)
			go func() {
				defer func() {
					<-semaphoreCh
					wg.Done()
				}()

				costs, err := l.GetCosts(ctx, closing.Start, closing.End)
				if err != nil {
					return
				}

				closing.Costs = costs
				costsCh <- closing
			}()
		}

		wg.Wait()
	}()

	return costsCh
}

func (logic) profitPipeline(_ context.Context, closingsCh <-chan repository.Closing) <-chan repository.Closing {
	profitCh := make(chan repository.Closing, bufferSize)

	go func() {
		defer log.Println("PROFIT PIPELINE IS DONE!")
		defer close(profitCh)

		for closing := range closingsCh {
			closing.CalculateProfit()
			profitCh <- closing
		}
	}()

	return profitCh
}

func (l logic) savePipeline(ctx context.Context, closingsCh <-chan repository.Closing) <-chan struct{} {
	doneCh := make(chan struct{})

	go func() {
		defer func() {
			log.Println("SAVE PIPELINE IS DONE!")
			doneCh <- struct{}{}
			close(doneCh)
		}()

		var wg sync.WaitGroup
		semaphoreCh := make(chan struct{}, bufferSize)

		for closing := range closingsCh {
			semaphoreCh <- struct{}{}

			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
					<-semaphoreCh
				}()
				_ = l.SaveProfit(ctx, closing.ID, closing.Profit)
			}()
		}

		wg.Wait()
	}()

	return doneCh
}
