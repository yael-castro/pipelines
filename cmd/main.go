package main

import (
	"context"
	"github.com/yael-castro/pipelines/internal/logic"
	"github.com/yael-castro/pipelines/internal/repository"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	// Limiting CPUs that can be executing simultaneously
	const threads = 1
	_ = runtime.GOMAXPROCS(threads)

	// Building main context
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	// Building executable command
	l := logic.New(repository.New())

	// Executing business logic
	println("Calculating profit... ")

	_ = l.CalculateProfit(ctx)

	println("Done!")
}
