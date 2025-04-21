package command

import (
	"context"
	"github.com/yael-castro/pipelines/internal/logic"
)

func New(l logic.Logic) func(ctx context.Context, args ...string) error {
	return func(ctx context.Context, args ...string) error {
		return l.CalculateProfit(ctx)
	}
}
