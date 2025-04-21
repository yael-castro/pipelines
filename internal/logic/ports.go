package logic

import (
	"context"
	"time"
)

type Logic interface {
	CalculateProfit(context.Context) error
}

type Repository interface {
	GetClosings(context.Context, ClosingStatus) (Closings, error)
	GetCosts(context.Context, time.Time, time.Time) (float64, error)
	GetSales(context.Context, time.Time, time.Time) (Sales, error)
	SaveProfit(context.Context, ClosingID, float64) error
}
