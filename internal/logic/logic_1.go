//go:build lineal

package logic

import (
	"context"
	"github.com/yael-castro/pipelines/internal/repository"
)

func (l logic) CalculateProfit(ctx context.Context) (err error) {
	closings, err := l.GetClosings(ctx, repository.Open) // STEP 0 (SOURCE)
	if err != nil {
		return
	}

	for _, closing := range closings { // 10,000 * (10ms + 10ms + 1ms + 10ms)
		closing.Sales, err = l.GetSales(ctx, closing.Start, closing.End) // STEP 1 (LATENCY > 10ms)
		if err != nil {
			continue
		}

		closing.Costs, err = l.GetCosts(ctx, closing.Start, closing.End) // STEP 2 (LATENCY = 10ms)
		if err != nil {
			continue
		}

		closing.CalculateProfit() // STEP 3 (LATENCY >= 0ms)

		_ = l.SaveProfit(ctx, closing.ID, closing.Profit) // STEP 4 (LATENCY = 10ms)
	}

	return
}
