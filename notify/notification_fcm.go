package notify

import (
	"context"
	"errors"
	"fmt"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/status"

	"firebase.google.com/go/v4/messaging"
	"github.com/appleboy/go-fcm"
)

// InitFCMClient use for initialize FCM Client.
func InitFCMClient(cfg *config.ConfYaml, key string) (*fcm.Client, error) {
	var err error

	if key == "" && cfg.Android.Credential == "" {
		return nil, errors.New("missing android api key")
	}

	if key != "" && key != cfg.Android.Credential {
		return fcm.NewClient(
			context.Background(),
			fcm.WithCredentialsJSON([]byte(key)),
		)
	}

	if FCMClient == nil {
		FCMClient, err = fcm.NewClient(
			context.Background(),
			fcm.WithCredentialsJSON([]byte(cfg.Android.Credential)),
		)
		return FCMClient, err
	}

	return FCMClient, nil
}

// GetAndroidNotification use for define Android notification.
// HTTP Connection Server Reference for Android
// https://firebase.google.com/docs/cloud-messaging/http-server-ref
func GetAndroidNotification(req *PushNotification) *messaging.Message {
	notification := &messaging.Message{
		Token:        req.To,
		Condition:    req.Condition,
		Notification: req.Notification,
		Android:      req.Android,
		Webpush:      req.Webpush,
		APNS:         req.APNS,
		FCMOptions:   req.FCMOptions,
	}

	// Add another field
	if len(req.Data) > 0 {
		notification.Data = make(map[string]string, len(req.Data))
		for k, v := range req.Data {
			notification.Data[k] = fmt.Sprintf("%s", v)
		}
	}

	if req.Title != "" || req.Message != "" || req.Image != "" {
		if notification.Notification == nil {
			notification.Notification = &messaging.Notification{}
		}
		if req.Title != "" {
			notification.Notification.Title = req.Title
		}
		if req.Message != "" {
			notification.Notification.Body = req.Message
		}
		if req.Image != "" {
			notification.Notification.ImageURL = req.Image
		}
	}

	return notification
}

// PushToAndroid provide send notification to Android server.
func PushToAndroid(req *PushNotification, cfg *config.ConfYaml) (resp *ResponsePush, err error) {
	logx.LogAccess.Debug("Start push notification for Android")

	var (
		client     *fcm.Client
		retryCount = 0
		maxRetry   = cfg.Android.MaxRetry
	)

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	// check message
	err = CheckMessage(req)
	if err != nil {
		logx.LogError.Error("request error: " + err.Error())
		return nil, err
	}

	resp = &ResponsePush{}
	client, err = InitFCMClient(cfg, cfg.Android.Credential)

Retry:
	notification := GetAndroidNotification(req)
	if err != nil {
		// FCM server error
		logx.LogError.Error("FCM server error: " + err.Error())
		return resp, err
	}

	res, err := client.Send(context.Background(), notification)
	if err != nil {
		// Send Message error
		logx.LogError.Error("FCM server send message error: " + err.Error())

		if req.IsTopic() {
			errLog := logPush(cfg, core.FailedPush, req.To, req, err)
			resp.Logs = append(resp.Logs, errLog)
			status.StatStorage.AddAndroidError(1)
		} else {
			for _, token := range req.Tokens {
				errLog := logPush(cfg, core.FailedPush, token, req, err)
				resp.Logs = append(resp.Logs, errLog)
			}
			status.StatStorage.AddAndroidError(int64(len(req.Tokens)))
		}
		return resp, err
	}

	if !req.IsTopic() {
		logx.LogAccess.Debug(fmt.Sprintf("Android Success count: %d, Failure count: %d", res.SuccessCount, res.FailureCount))
	}

	status.StatStorage.AddAndroidSuccess(int64(res.SuccessCount))
	status.StatStorage.AddAndroidError(int64(res.FailureCount))

	var newTokens []string
	// result from Send messages to specific devices
	for k, result := range res.Responses {
		to := ""
		if k < len(req.Tokens) {
			to = req.Tokens[k]
		} else {
			to = req.To
		}

		if result.Error != nil {
			errLog := logPush(cfg, core.FailedPush, to, req, result.Error)
			resp.Logs = append(resp.Logs, errLog)
			continue
		}

		logPush(cfg, core.SucceededPush, to, req, nil)
	}

	// result from Send messages to topics
	if req.IsTopic() {
		to := ""
		if req.Topic != "" {
			to = req.Topic
		} else {
			to = req.Condition
		}
		logx.LogAccess.Debug("Send Topic Message: ", to)
		// Success
		if res.SuccessCount == 1 {
			logPush(cfg, core.SucceededPush, to, req, nil)
		} else {
			// failure
			errLog := logPush(cfg, core.FailedPush, to, req, res.Responses[0].Error)
			resp.Logs = append(resp.Logs, errLog)
		}
	}

	if len(newTokens) > 0 && retryCount < maxRetry {
		retryCount++

		// resend fail token
		req.Tokens = newTokens
		goto Retry
	}

	return resp, nil
}

func logPush(cfg *config.ConfYaml, status, token string, req *PushNotification, err error) logx.LogPushEntry {
	return logx.LogPush(&logx.InputLog{
		ID:          req.ID,
		Status:      status,
		Token:       token,
		Message:     req.Message,
		Platform:    req.Platform,
		Error:       err,
		HideToken:   cfg.Log.HideToken,
		HideMessage: cfg.Log.HideMessages,
		Format:      cfg.Log.Format,
	})
}
