package app

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"worker-go-pool/internal/fetch"
	"worker-go-pool/internal/workerpool"
)

func Run(ctx context.Context, filepath string, concurrency int, timeout time.Duration, out io.Writer) error {

	if concurrency <= 0 {
		concurrency = 1
	}

	f, err := os.Open(filepath)

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()

	jobsCh := make(chan fetch.Job, concurrency)

	go func() {
		defer close(jobsCh)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			select {
			case <-ctx.Done():
				return
			case jobsCh <- fetch.Job{URL: line}:

			}
		}

	}()

	workFn := func(ctx context.Context, job fetch.Job) fetch.Result {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return fetch.FetchOnce(ctx, job)
	}

	resultsCh := workerpool.Run(ctx, jobsCh, concurrency, workFn)

	enc := json.NewEncoder(out)

	for res := range resultsCh {
		if err := enc.Encode(res); err != nil {
			return fmt.Errorf("write jsonl: %w", err)
		}
	}

	return nil

}
