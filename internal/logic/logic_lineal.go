//go:build lineal

package logic

import (
	"context"
)

func (l logic) CalculateProfit(ctx context.Context) (err error) {
	closings, err := l.GetClosings(ctx, Open) // STEP 0 (SOURCE)
	if err != nil {
		return
	}

	// EXECUTION TIME = 10,000 * (10ms + 10ms + 1ms + 10ms)
	for _, closing := range closings {
		closing.Sales, err = l.GetSales(ctx, closing.Start, closing.End) // STEP 1 (LATENCY >= 10ms)
		if err != nil {
			continue
		}

		closing.Costs, err = l.GetCosts(ctx, closing.Start, closing.End) // STEP 2 (LATENCY = 10ms)
		if err != nil {
			continue
		}

		closing.CalculateNetProfit() // STEP 3 (LATENCY >= 0ms)

		_ = l.SaveProfit(ctx, closing.ID, closing.NetProfit) // STEP 4 (LATENCY = 10ms)
	}

	return
}
