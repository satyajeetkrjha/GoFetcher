package fetch

type Result struct {
	Status     int    `json:"status"`
	URL        string `json:"url"`
	DurationMS int64  `json:"duration_ms"`
	Error      string `json:"error"`
	Bytes      int64  `json:"bytes"`
}

type Job struct {
	URL string
}
