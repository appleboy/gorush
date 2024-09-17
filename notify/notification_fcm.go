package notify

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/status"

	"firebase.google.com/go/v4/messaging"
	"github.com/appleboy/go-fcm"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// InitFCMClient use for initialize FCM Client.
func InitFCMClient(ctx context.Context, cfg *config.ConfYaml) (*fcm.Client, error) {
	var opts []fcm.Option

	credential := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if cfg.Android.Credential == "" &&
		cfg.Android.KeyPath == "" &&
		credential == "" {
		return nil, errors.New("missing fcm credential data")
	}

	if cfg.Android.KeyPath != "" && fileExists(cfg.Android.KeyPath) {
		opts = append(opts, fcm.WithCredentialsFile(cfg.Android.KeyPath))
	}

	if cfg.Android.Credential != "" {
		opts = append(opts, fcm.WithCredentialsJSON([]byte(cfg.Android.Credential)))
	}

	if FCMClient != nil {
		return FCMClient, nil
	}

	var err error
	FCMClient, err = fcm.NewClient(
		ctx,
		opts...,
	)

	return FCMClient, err
}

// GetAndroidNotification use for define Android notification.
// HTTP Connection Server Reference for Android
// https://firebase.google.com/docs/cloud-messaging/http-server-ref
func GetAndroidNotification(req *PushNotification) []*messaging.Message {
	var messages []*messaging.Message

	if req.Title != "" || req.Message != "" || req.Image != "" {
		if req.Notification == nil {
			req.Notification = &messaging.Notification{}
		}
		if req.Title != "" {
			req.Notification.Title = req.Title
		}
		if req.Message != "" {
			req.Notification.Body = req.Message
		}
		if req.Image != "" {
			req.Notification.ImageURL = req.Image
		}
		if req.MutableContent {
			req.APNS = &messaging.APNSConfig{
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						MutableContent: req.MutableContent,
					},
				},
			}
		}
	}

	// content-available is for background notifications and a badge, alert
	// and sound keys should not be present.
	// See: https://developer.apple.com/documentation/usernotifications/generating-a-remote-notification
	if req.ContentAvailable {
		req.APNS = &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "5",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					ContentAvailable: req.ContentAvailable,
					CustomData:       req.Data,
				},
			},
		}
	}

	// Check if the notification has a sound
	if req.Sound != nil {
		sound, ok := req.Sound.(string)
		if ok {
			switch {
			case req.APNS == nil:
				req.APNS = &messaging.APNSConfig{
					Payload: &messaging.APNSPayload{
						Aps: &messaging.Aps{
							Sound: sound,
						},
					},
				}
			case req.APNS.Payload == nil:
				req.APNS.Payload = &messaging.APNSPayload{
					Aps: &messaging.Aps{
						Sound: sound,
					},
				}

			case req.APNS.Payload.Aps == nil:
				req.APNS.Payload.Aps = &messaging.Aps{
					Sound: sound,
				}
			default:
				req.APNS.Payload.Aps.Sound = sound
			}

			if req.Android == nil {
				req.Android = &messaging.AndroidConfig{
					Notification: &messaging.AndroidNotification{
						Sound: sound,
					},
				}
			}
		}
	}

	// Check if the notification is a topic
	if req.IsTopic() {
		message := &messaging.Message{
			Notification: req.Notification,
			Android:      req.Android,
			Webpush:      req.Webpush,
			APNS:         req.APNS,
			FCMOptions:   req.FCMOptions,
			Topic:        req.Topic,
			Condition:    req.Condition,
		}

		messages = append(messages, message)
	}

	var data map[string]string
	if len(req.Data) > 0 {
		data = make(map[string]string, len(req.Data))
		for k, v := range req.Data {
			switch v.(type) {
			case string:
				data[k] = fmt.Sprintf("%s", v)
			default:
				if v, err := json.Marshal(v); err == nil {
					data[k] = string(v)
				}
			}
		}
	}

	// Loop through the tokens and create a message for each one
	for _, token := range req.Tokens {
		message := &messaging.Message{
			Token:        token,
			Notification: req.Notification,
			Android:      req.Android,
			Webpush:      req.Webpush,
			APNS:         req.APNS,
			FCMOptions:   req.FCMOptions,
		}

		// Add another field
		if len(req.Data) > 0 {
			message.Data = data
		}

		messages = append(messages, message)
	}

	return messages
}

// PushToAndroid provide send notification to Android server.
func PushToAndroid(ctx context.Context, req *PushNotification, cfg *config.ConfYaml) (resp *ResponsePush, err error) {
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
	client, err = InitFCMClient(ctx, cfg)

Retry:
	messages := GetAndroidNotification(req)
	if err != nil {
		// FCM server error
		logx.LogError.Error("FCM server error: " + err.Error())
		return resp, err
	}

	if req.Development {
		for i, msg := range messages {
			m, _ := json.Marshal(msg)
			logx.LogAccess.Infof("message #%d - %s", i, m)
		}
	}

	res, err := client.Send(ctx, messages...)
	if err != nil {
		newErr := fmt.Errorf("fcm service send message error: %v", err)
		logx.LogError.Error(newErr)
		errLog := logPush(cfg, core.FailedPush, "", req, newErr)
		resp.Logs = append(resp.Logs, errLog)
		status.StatStorage.AddAndroidError(1)

		return resp, newErr
	}

	logx.LogAccess.Debug(fmt.Sprintf("Android Success count: %d, Failure count: %d", res.SuccessCount, res.FailureCount))
	status.StatStorage.AddAndroidSuccess(int64(res.SuccessCount))
	status.StatStorage.AddAndroidError(int64(res.FailureCount))

	// result from Send messages to topics
	retryTopic := false
	if req.IsTopic() {
		to := ""
		if req.Topic != "" {
			to = req.Topic
		}
		if req.Condition != "" {
			to = req.Condition
		}
		logx.LogAccess.Debug("Send Topic Message: ", to)

		newResp := res.Responses[0]
		if newResp.Success {
			logPush(cfg, core.SucceededPush, to, req, nil)
		}

		if newResp.Error != nil {
			// failure
			errLog := logPush(cfg, core.FailedPush, to, req, newResp.Error)
			resp.Logs = append(resp.Logs, errLog)
			retryTopic = true
		}

		// remove the first response
		res.Responses = res.Responses[1:]
	}

	var newTokens []string
	for k, result := range res.Responses {
		if result.Error != nil {
			errLog := logPush(cfg, core.FailedPush, req.Tokens[k], req, result.Error)
			resp.Logs = append(resp.Logs, errLog)
			newTokens = append(newTokens, req.Tokens[k])
			continue
		}
		logPush(cfg, core.SucceededPush, req.Tokens[k], req, nil)
	}

	if len(newTokens) > 0 && retryCount < maxRetry {
		retryCount++

		if req.IsTopic() && !retryTopic {
			req.Topic = ""
			req.Condition = ""
		}

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
