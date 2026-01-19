```
go build -o fetcher ./cmd/fetcher
./fetcher --file urls.txt --concurrency 3 --timeout 3s
```