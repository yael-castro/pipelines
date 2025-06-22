package main

import (
	"context"
	"github.com/yael-castro/pipelines/internal/command"
	"github.com/yael-castro/pipelines/internal/logic"
	"github.com/yael-castro/pipelines/internal/repository"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	// Limiting CPUs that can be using simultaneously
	const threads = 1
	_ = runtime.GOMAXPROCS(threads)

	// Building main context
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	// Setting default logger
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Building executable command
	cmd := command.New(logic.New(repository.New()))

	// Executing business logic
	slog.InfoContext(ctx, "calculating_profit")

	err := cmd(ctx)

	slog.InfoContext(ctx, "done", "error", err)
}
