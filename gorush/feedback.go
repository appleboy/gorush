package gorush

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"time"
)

// DispatchFeedback sends a feedback to the configured gateway.
func DispatchFeedback(log LogPushEntry, url string, timeout int64) error {

	if url == "" {
		return errors.New("The url can't be empty")
	}

	payload, err := json.Marshal(log)

	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var client = &http.Client{
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
