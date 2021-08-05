package notify

import (
	"bytes"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/appleboy/gorush/logx"
)

// DispatchFeedback sends a feedback to the configured gateway.
func DispatchFeedback(log logx.LogPushEntry, url string, timeout int64) error {
	if url == "" {
		return errors.New("url can't be empty")
	}

	payload, err := json.Marshal(log)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
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
