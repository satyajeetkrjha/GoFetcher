package fetch

import (
	"context"
	"io"
	"net/http"
	"time"
)

func FetchOnce(ctx context.Context, job Job) Result {
	// Simulate fetching the URL (this is a placeholder for actual fetch logic)
	// In a real implementation, you would perform an HTTP request here.

	startTime := time.Now()

	result := Result{
		URL: job.URL,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, job.URL, nil)

	if err != nil {
		result.Error = err.Error()
		result.Status = 0
		result.DurationMS = time.Since(startTime).Milliseconds()
		result.Bytes = 0
		return result
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		result.Error = err.Error()
		result.Status = 0
		result.DurationMS = time.Since(startTime).Milliseconds()
		result.Bytes = 0
		return result
	}

	defer resp.Body.Close()

	result.Status = resp.StatusCode

	n, readErr := io.Copy(io.Discard, resp.Body)
	result.Bytes = int64(n)

	if readErr != nil {
		result.Error = readErr.Error()
	} else {
		result.Error = ""
	}

	result.DurationMS = time.Since(startTime).Milliseconds()
	return result
}
