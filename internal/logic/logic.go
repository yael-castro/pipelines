package logic

import (
	"context"
	"github.com/yael-castro/pipelines/internal/repository"
)

func New(repo repository.Repository) Logic {
	return logic{
		Repository: repo,
	}
}

type Logic interface {
	CalculateProfit(context.Context) error
}

type logic struct {
	repository.Repository
}
