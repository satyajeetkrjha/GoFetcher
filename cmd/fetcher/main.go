package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"worker-go-pool/internal/app"
)

func main() {
	// CLI flags
	filePath := flag.String("file", "", "path to file containing URLs (one per line)")
	concurrency := flag.Int("concurrency", 5, "number of concurrent workers")
	timeout := flag.Duration("timeout", 5*time.Second, "per-request timeout")

	flag.Parse()

	if *filePath == "" {
		log.Fatal("missing required --file flag")
	}

	// Root context with Ctrl+C / SIGTERM handling
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Run the app
	err := app.Run(ctx, *filePath, *concurrency, *timeout, os.Stdout)
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
}
