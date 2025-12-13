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

// setupFCMNotification sets up the notification fields on the request.
func setupFCMNotification(req *PushNotification) {
	if req.Title == "" && req.Message == "" && req.Image == "" {
		return
	}
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

// setupFCMContentAvailable sets up the content available config for background notifications.
func setupFCMContentAvailable(req *PushNotification) {
	if !req.ContentAvailable {
		return
	}
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

// setAPNSSound sets the sound on the APNS config, initializing nested structs as needed.
func setAPNSSound(req *PushNotification, sound string) {
	switch {
	case req.APNS == nil:
		req.APNS = &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{Sound: sound},
			},
		}
	case req.APNS.Payload == nil:
		req.APNS.Payload = &messaging.APNSPayload{
			Aps: &messaging.Aps{Sound: sound},
		}
	case req.APNS.Payload.Aps == nil:
		req.APNS.Payload.Aps = &messaging.Aps{Sound: sound}
	default:
		req.APNS.Payload.Aps.Sound = sound
	}
}

// setupFCMSound sets up the sound configuration for FCM notifications.
func setupFCMSound(req *PushNotification) {
	if req.Sound == nil {
		return
	}
	sound, ok := req.Sound.(string)
	if !ok {
		return
	}
	setAPNSSound(req, sound)
	if req.Android == nil {
		req.Android = &messaging.AndroidConfig{
			Priority: req.Priority,
			Notification: &messaging.AndroidNotification{
				Sound: sound,
			},
		}
	}
}

// convertDataToStringMap converts the request data to a string map.
func convertDataToStringMap(data map[string]interface{}) map[string]string {
	if len(data) == 0 {
		return nil
	}
	result := make(map[string]string, len(data))
	for k, v := range data {
		switch v.(type) {
		case string:
			result[k] = fmt.Sprintf("%s", v)
		default:
			if b, err := json.Marshal(v); err == nil {
				result[k] = string(b)
			}
		}
	}
	return result
}

// buildFCMMessage creates a new FCM message with common fields.
func buildFCMMessage(req *PushNotification, data map[string]string) *messaging.Message {
	msg := &messaging.Message{
		Notification: req.Notification,
		Android:      req.Android,
		Webpush:      req.Webpush,
		APNS:         req.APNS,
		FCMOptions:   req.FCMOptions,
	}
	if len(data) > 0 {
		msg.Data = data
	}
	return msg
}

// GetAndroidNotification use for define Android notification.
// HTTP Connection Server Reference for Android
// https://firebase.google.com/docs/cloud-messaging/http-server-ref
func GetAndroidNotification(req *PushNotification) []*messaging.Message {
	setupFCMNotification(req)
	setupFCMContentAvailable(req)
	setupFCMSound(req)

	data := convertDataToStringMap(req.Data)
	var messages []*messaging.Message

	if req.IsTopic() {
		msg := buildFCMMessage(req, data)
		msg.Topic = req.Topic
		msg.Condition = req.Condition
		messages = append(messages, msg)
	}

	for _, token := range req.Tokens {
		msg := buildFCMMessage(req, data)
		msg.Token = token
		messages = append(messages, msg)
	}

	return messages
}

// handleTopicResponse processes the topic message response and returns whether to retry.
func handleTopicResponse(
	cfg *config.ConfYaml, req *PushNotification, res *messaging.BatchResponse, resp *ResponsePush,
) bool {
	if !req.IsTopic() {
		return false
	}

	to := req.Topic
	if req.Condition != "" {
		to = req.Condition
	}
	logx.LogAccess.Debug("Send Topic Message: ", to)

	topicResp := res.Responses[0]
	if topicResp.Success {
		logPush(cfg, core.SucceededPush, to, req, nil)
	}

	retryTopic := false
	if topicResp.Error != nil {
		errLog := logPush(cfg, core.FailedPush, to, req, topicResp.Error)
		resp.Logs = append(resp.Logs, errLog)
		retryTopic = true
	}

	// remove the first response
	res.Responses = res.Responses[1:]
	return retryTopic
}

// handleTokenResponses processes individual token responses and returns tokens that need retry.
func handleTokenResponses(
	cfg *config.ConfYaml, req *PushNotification, res *messaging.BatchResponse, resp *ResponsePush,
) []string {
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
	return newTokens
}

// logDevMessages logs messages in development mode.
func logDevMessages(messages []*messaging.Message) {
	for i, msg := range messages {
		m, _ := json.Marshal(msg)
		logx.LogAccess.Infof("message #%d - %s", i, m)
	}
}

// PushToAndroid provide send notification to Android server.
func PushToAndroid(
	ctx context.Context,
	req *PushNotification,
	cfg *config.ConfYaml,
) (resp *ResponsePush, err error) {
	logx.LogAccess.Debug("Start push notification for Android")

	if err = CheckMessage(req); err != nil {
		logx.LogError.Error("request error: " + err.Error())
		return nil, err
	}

	client, err := InitFCMClient(ctx, cfg)
	if err != nil {
		logx.LogError.Error("FCM server error: " + err.Error())
		return nil, err
	}

	maxRetry := cfg.Android.MaxRetry
	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	resp = &ResponsePush{}
	retryCount := 0

Retry:
	messages := GetAndroidNotification(req)

	if req.Development {
		logDevMessages(messages)
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

	logx.LogAccess.Debug(
		fmt.Sprintf(
			"Android Success count: %d, Failure count: %d",
			res.SuccessCount,
			res.FailureCount,
		),
	)
	status.StatStorage.AddAndroidSuccess(int64(res.SuccessCount))
	status.StatStorage.AddAndroidError(int64(res.FailureCount))

	retryTopic := handleTopicResponse(cfg, req, res, resp)
	newTokens := handleTokenResponses(cfg, req, res, resp)

	if len(newTokens) > 0 && retryCount < maxRetry {
		retryCount++
		if req.IsTopic() && !retryTopic {
			req.Topic = ""
			req.Condition = ""
		}
		req.Tokens = newTokens
		goto Retry
	}

	return resp, nil
}

func logPush(
	cfg *config.ConfYaml,
	status, token string,
	req *PushNotification,
	err error,
) logx.LogPushEntry {
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
