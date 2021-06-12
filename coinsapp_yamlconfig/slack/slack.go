package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

func SendMessage(webhook string, message string) error {
	// Convert the struct to bytes
	textBody, err := json.Marshal(SlackRequestBody{Text: message})

	if err != nil {
		return err
	}

	// Define the HTTP request
	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(textBody))

	if err != nil {
		return err
	}

	// Add the request headers
	req.Header.Add("content-Type", "application/json")

	// Create the HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Set a buffer for the response body
	buf := new(bytes.Buffer)

	// Read the response body
	buf.ReadFrom(resp.Body)

	if buf.String() != "ok" {
		return errors.New("Non-Ok response returned from slack")
	}

	return nil

}
