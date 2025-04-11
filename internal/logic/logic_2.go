//go:build pipeline

package logic

import (
	"context"
	"github.com/yael-castro/pipelines/internal/repository"
)

func (l logic) CalculateProfit(ctx context.Context) (err error) {
	closings, err := l.GetClosings(ctx, repository.Open)
	if err != nil {
		return
	}

	closingCh := l.closingChannel(closings)      // STEP 0 (SOURCE)
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
	closingCh = l.salesPipeline(ctx, closingCh)  // STEP 1
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
		defer close(salesCh)

		for closing := range closingsCh {
			sales, err := l.GetSales(ctx, closing.Start, closing.End)
			if err != nil {
				continue
			}

			closing.Sales = sales
			salesCh <- closing
		}
	}()

	return salesCh
}

func (l logic) costsPipeline(ctx context.Context, closingsCh <-chan repository.Closing) <-chan repository.Closing {
	costsCh := make(chan repository.Closing)

	go func() {
		defer close(costsCh)

		for closing := range closingsCh {
			costs, err := l.GetCosts(ctx, closing.Start, closing.End)
			if err != nil {
				return
			}

			closing.Costs = costs
			costsCh <- closing
		}
	}()

	return costsCh
}

func (logic) profitPipeline(_ context.Context, closingsCh <-chan repository.Closing) <-chan repository.Closing {
	profitCh := make(chan repository.Closing)

	go func() {
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
			doneCh <- struct{}{}
			close(doneCh)
		}()

		for closing := range closingsCh {
			_ = l.SaveProfit(ctx, closing.ID, closing.Profit)
		}
	}()

	return doneCh
}
