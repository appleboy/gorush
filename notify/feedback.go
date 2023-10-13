package notify

import (
	"bytes"
	"context"
	"errors"
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
func DispatchFeedback(ctx context.Context, log logx.LogPushEntry, url string, timeout int64, header []string) error {
	if url == "" {
		return errors.New("url can't be empty")
	}

	payload, err := json.Marshal(log)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	headers := extractHeaders(header)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	feedbackClient.Timeout = time.Duration(timeout) * time.Second
	resp, err := feedbackClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}

	return nil
}
