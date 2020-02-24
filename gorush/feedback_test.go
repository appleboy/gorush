package gorush

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/stretchr/testify/assert"
)

func TestEmptyFeedbackURL(t *testing.T) {
	// PushConf, _ = config.LoadConf("")
	logEntry := LogPushEntry{
		ID:       "",
		Type:     "",
		Platform: "",
		Token:    "",
		Message:  "",
		Error:    "",
	}

	err := DispatchFeedback(logEntry, PushConf.Core.FeedbackURL, PushConf.Core.FeedbackTimeout)
	assert.NotNil(t, err)
}

func TestHTTPErrorInFeedbackCall(t *testing.T) {
	config, _ := config.LoadConf("")
	config.Core.FeedbackURL = "http://test.example.com/api/"
	logEntry := LogPushEntry{
		ID:       "",
		Type:     "",
		Platform: "",
		Token:    "",
		Message:  "",
		Error:    "",
	}

	err := DispatchFeedback(logEntry, config.Core.FeedbackURL, config.Core.FeedbackTimeout)
	assert.NotNil(t, err)
}

func TestSuccessfulFeedbackCall(t *testing.T) {

	// Mock http server
	httpMock := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/dispatch" {
				w.Header().Add("Content-Type", "application/json")
				_, err := w.Write([]byte(`{}`))

				if err != nil {
					log.Println(err)
					panic(err)
				}
			}
		}),
	)
	defer httpMock.Close()

	config, _ := config.LoadConf("")
	config.Core.FeedbackURL = httpMock.URL
	logEntry := LogPushEntry{
		ID:       "",
		Type:     "",
		Platform: "",
		Token:    "",
		Message:  "",
		Error:    "",
	}

	err := DispatchFeedback(logEntry, config.Core.FeedbackURL, config.Core.FeedbackTimeout)
	assert.Nil(t, err)
}
