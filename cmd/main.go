package main

import (
	"context"
	"github.com/yael-castro/pipelines/internal/command"
	"github.com/yael-castro/pipelines/internal/logic"
	"github.com/yael-castro/pipelines/internal/repository"
	"log"
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

	// Building executable command
	cmd := command.New(logic.New(repository.New()))

	// Executing business logic
	log.Println("Calculating profit... ")

	_ = cmd(ctx)

	log.Println("Done!")
}
