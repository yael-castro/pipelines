package repository

import (
	"context"
	"github.com/yael-castro/pipelines/internal/logic"
	"log"
	"strconv"
	"time"
)

const latencyPerOp = 10 * time.Millisecond

func New() logic.Repository {
	return repository{
		logger: log.Default(),
	}
}

type repository struct {
	logger *log.Logger
}

func (repository) GetCosts(ctx context.Context, start, end time.Time) (float64, error) {
	timer := time.NewTimer(latencyPerOp)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-timer.C: // Simulating latency
		const defaultCosts = 1_000
		return defaultCosts, nil
	}
}

func (repository) GetSales(ctx context.Context, start, end time.Time) (logic.Sales, error) {
	timer := time.NewTimer(latencyPerOp)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-timer.C:
	}

	const defaultSales = 100
	sales := make(logic.Sales, 0, defaultSales)

	now := time.Now()

	for range defaultSales {
		const defaultValue = 100

		sales = append(sales, logic.Sale{
			Date:  now,
			Value: defaultValue,
		})
	}

	return sales, nil
}

func (repository) GetClosings(ctx context.Context, _ logic.ClosingStatus) (logic.Closings, error) {
	timer := time.NewTimer(latencyPerOp)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-timer.C:
	}

	const defaultClosings = 10_000
	closings := make(logic.Closings, 0, defaultClosings)

	now := time.Now()
	start := now.Add(-time.Hour)
	end := now.Add(time.Hour)

	for i := range defaultClosings {
		closings = append(closings, logic.Closing{
			ID:    logic.ClosingID(strconv.FormatInt(int64(i), 10)),
			Start: start,
			End:   end,
		})
	}

	return closings, nil
}

func (r repository) SaveProfit(ctx context.Context, id logic.ClosingID, profit float64) error {
	timer := time.NewTimer(latencyPerOp)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
	}

	r.logger.Println("CLOSING ID:", id, "PROFIT:", profit)
	return nil
}
