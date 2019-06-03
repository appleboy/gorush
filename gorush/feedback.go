package gorush

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// DispatchFeedback sends a feedback to the configured gateway.
func DispatchFeedback(log LogPushEntry) (*http.Response, error) {

	if PushConf.Core.FeedbackURL == "" {
		return nil, nil
	}

	payload, err := json.Marshal(log)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", PushConf.Core.FeedbackURL, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	HTTPClient := &http.Client{}
	httpRes, err := HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	return httpRes, nil
}
