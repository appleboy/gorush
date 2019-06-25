package gorush

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// DispatchFeedback sends a feedback to the configured gateway.
func DispatchFeedback(log LogPushEntry) {

	if PushConf.Core.FeedbackURL == "" {
		return
	}

	payload, err := json.Marshal(log)

	if err != nil {
		LogError.Error(err)
		return
	}

	req, _ := http.NewRequest("POST", PushConf.Core.FeedbackURL, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	HTTPClient := &http.Client{}
	resp, err := HTTPClient.Do(req)

	if err != nil {
		LogError.Error(err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}
}
