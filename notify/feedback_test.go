package notify

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"

	"github.com/stretchr/testify/assert"
)

func TestEmptyFeedbackURL(t *testing.T) {
	cfg, _ := config.LoadConf()
	logEntry := logx.LogPushEntry{
		ID:       "",
		Type:     "",
		Platform: "",
		Token:    "",
		Message:  "",
		Error:    "",
	}

	err := DispatchFeedback(
		context.Background(),
		logEntry,
		cfg.Core.FeedbackURL,
		cfg.Core.FeedbackTimeout,
		cfg.Core.FeedbackHeader,
	)
	assert.NotNil(t, err)
}

func TestHTTPErrorInFeedbackCall(t *testing.T) {
	cfg, _ := config.LoadConf()
	cfg.Core.FeedbackURL = "http://test.example.com/api/"
	logEntry := logx.LogPushEntry{
		ID:       "",
		Type:     "",
		Platform: "",
		Token:    "",
		Message:  "",
		Error:    "",
	}

	err := DispatchFeedback(
		context.Background(),
		logEntry,
		cfg.Core.FeedbackURL,
		cfg.Core.FeedbackTimeout,
		cfg.Core.FeedbackHeader,
	)
	assert.NotNil(t, err)
}

func TestSuccessfulFeedbackCall(t *testing.T) {
	// Mock http server
	httpMock := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/dispatch" {
				// check http header
				if r.Header.Get("x-gorush-token") != "1234" {
					panic("x-gorush-token header is not set")
				}

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

	cfg, _ := config.LoadConf()
	cfg.Core.FeedbackURL = httpMock.URL + "/dispatch"
	cfg.Core.FeedbackHeader = []string{
		"x-gorush-token: 1234",
	}
	logEntry := logx.LogPushEntry{
		ID:       "",
		Type:     "",
		Platform: "",
		Token:    "",
		Message:  "",
		Error:    "",
	}

	err := DispatchFeedback(
		context.Background(),
		logEntry,
		cfg.Core.FeedbackURL,
		cfg.Core.FeedbackTimeout,
		cfg.Core.FeedbackHeader,
	)
	assert.Nil(t, err)
}
