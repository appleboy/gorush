package gorush

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// DispatchFeedback sends a feedback to the configured gateway.
func DispatchFeedback(log LogPushEntry, url string) error {

	if url == "" {
		return errors.New("The url can't be empty")
	}

	payload, err := json.Marshal(log)

	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	HTTPClient := &http.Client{}
	resp, err := HTTPClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}

	return nil
}
