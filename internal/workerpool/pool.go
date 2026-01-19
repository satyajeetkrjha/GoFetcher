package workerpool

import (
	"context"
	"sync"
	"worker-go-pool/fetch"
)

type WorkFn func(ctx context.Context, job fetch.Job) fetch.Result

func Run(ctx context.Context, jobs <-chan fetch.Job, concurrency int, workFn WorkFn) <-chan fetch.Result {
	out := make(chan fetch.Result, concurrency)

	if concurrency <= 0 {
		concurrency = 1
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			worker(ctx, jobs, out, workFn)
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}

func worker(ctx context.Context, jobs <-chan fetch.Job, out chan<- fetch.Result, workFn WorkFn) {
	for {
		var job fetch.Job
		var ok bool

		select {
		case <-ctx.Done():
			return
		case job, ok = <-jobs:
			if !ok {
				return
			}
		}

		res := workFn(ctx, job)

		select {
		case <-ctx.Done():
			return
		case out <- res:
		}
	}
}
