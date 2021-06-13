package monitor

import (
	"net/http"
	"time"
)

func GetStatusCode(url string, timeoutLimit time.Duration) (int, error) {

	client := http.Client{
		Timeout: timeoutLimit * time.Second,
	}

	r, err := client.Get(url)

	if err != nil {
		return 0, err
	}

	return r.StatusCode, nil
}
