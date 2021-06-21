package cloudflare

import (
	"net/http"
	"time"
)

func getResponseData(url string, timeoutLimit time.Duration) (http.Header, int) {
	client := http.Client{
		Timeout: timeoutLimit * time.Second,
	}

	r, err := client.Get(url)

	if err != nil {
		panic(err)
	}

	return r.Header, r.StatusCode
}
