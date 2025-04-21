package logic_test

import (
	"context"
	"github.com/yael-castro/pipelines/internal/logic"
	"github.com/yael-castro/pipelines/internal/repository"
	"io"
	"log"
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// Limiting CPUs that can be using simultaneously
	const threads = 1
	_ = runtime.GOMAXPROCS(threads)

	// Disabling logs
	log.SetOutput(io.Discard)

	os.Exit(m.Run())
}

func BenchmarkLogic_CalculateProfit(b *testing.B) {
	// Context to avoid go routine leaks
	ctx, cancel := context.WithCancel(context.Background())

	b.Cleanup(func() {
		cancel()
	})

	l := logic.New(repository.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := l.CalculateProfit(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}
