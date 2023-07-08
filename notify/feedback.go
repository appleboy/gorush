package notify

import (
	"bytes"
	"context"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/appleboy/gorush/logx"
)

// extractHeaders converts a slice of strings to a map of strings.
func extractHeaders(headers []string) map[string]string {
	result := make(map[string]string)
	for _, header := range headers {
		parts := strings.Split(header, ":")
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}

// DispatchFeedback sends a feedback to the configured gateway.
func DispatchFeedback(log logx.LogPushEntry, url string, timeout int64, header []string) error {
	if url == "" {
		return errors.New("url can't be empty")
	}

	payload, err := json.Marshal(log)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	headers := extractHeaders(header)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: transport,
	}

	resp, err := client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}

	return nil
}
