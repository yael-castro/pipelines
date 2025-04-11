package repository

import (
	"context"
	"strconv"
	"time"
)

const latencyPerOp = 10 * time.Millisecond

func New() Repository {
	return repository{}
}

type Repository interface {
	GetClosings(context.Context, ClosingStatus) (Closings, error)
	GetCosts(context.Context, time.Time, time.Time) (float64, error)
	GetSales(context.Context, time.Time, time.Time) (Sales, error)
	SaveProfit(context.Context, ClosingID, float64) error
}

type repository struct{}

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

func (repository) GetSales(ctx context.Context, start, end time.Time) (Sales, error) {
	timer := time.NewTimer(latencyPerOp)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-timer.C:
	}

	const defaultSales = 100
	sales := make(Sales, 0, defaultSales)

	now := time.Now()

	for range defaultSales {
		const defaultValue = 100

		sales = append(sales, Sale{
			Date:  now,
			Value: defaultValue,
		})
	}

	return sales, nil
}

func (repository) GetClosings(ctx context.Context, _ ClosingStatus) (Closings, error) {
	timer := time.NewTimer(latencyPerOp)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-timer.C:
	}

	const defaultClosings = 10_000
	closings := make(Closings, 0, defaultClosings)

	now := time.Now()
	start := now.Add(-time.Hour)
	end := now.Add(time.Hour)

	for i := range defaultClosings {
		closings = append(closings, Closing{
			ID:    ClosingID(strconv.FormatInt(int64(i), 10)),
			Start: start,
			End:   end,
		})
	}

	return closings, nil
}

func (repository) SaveProfit(ctx context.Context, id ClosingID, profit float64) error {
	timer := time.NewTimer(latencyPerOp)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
	}

	return nil
}
