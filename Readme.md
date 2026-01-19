```
go build -o fetcher ./cmd/fetcher
./fetcher --file urls.txt --concurrency 3 --timeout 3s
```

```
A Go-based concurrent URL fetcher that reads URLs from a file, processes them using a bounded worker pool, and outputs per-URL results as streaming JSON lines, with proper timeout and cancellation handling.
```
