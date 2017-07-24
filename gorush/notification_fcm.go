package gorush

import (
	"fmt"

	"github.com/appleboy/go-fcm"
)

// GetAndroidNotification use for define Android notification.
// HTTP Connection Server Reference for Android
// https://firebase.google.com/docs/cloud-messaging/http-server-ref
func GetAndroidNotification(req PushNotification) *fcm.Message {
	notification := &fcm.Message{
		To:                    req.To,
		CollapseKey:           req.CollapseKey,
		ContentAvailable:      req.ContentAvailable,
		DelayWhileIdle:        req.DelayWhileIdle,
		TimeToLive:            req.TimeToLive,
		RestrictedPackageName: req.RestrictedPackageName,
		DryRun:                req.DryRun,
	}

	notification.RegistrationIDs = req.Tokens

	if len(req.Priority) > 0 && req.Priority == "high" {
		notification.Priority = "high"
	}

	// Add another field
	if len(req.Data) > 0 {
		notification.Data = make(map[string]interface{})
		for k, v := range req.Data {
			notification.Data[k] = v
		}
	}

	notification.Notification = &req.Notification

	// Set request message if body is empty
	if len(req.Message) > 0 {
		notification.Notification.Body = req.Message
	}

	if len(req.Title) > 0 {
		notification.Notification.Title = req.Title
	}

	if len(req.Sound) > 0 {
		notification.Notification.Sound = req.Sound
	}

	return notification
}

// PushToAndroid provide send notification to Android server.
func PushToAndroid(req PushNotification) bool {
	LogAccess.Debug("Start push notification for Android")
	if PushConf.Core.Sync {
		defer req.WaitDone()
	}

	var (
		APIKey     string
		retryCount = 0
		maxRetry   = PushConf.Android.MaxRetry
	)

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	// check message
	err := CheckMessage(req)

	if err != nil {
		LogError.Error("request error: " + err.Error())
		return false
	}

Retry:
	var isError = false
	notification := GetAndroidNotification(req)

	if APIKey = PushConf.Android.APIKey; req.APIKey != "" {
		APIKey = req.APIKey
	}

	client, err := fcm.NewClient(APIKey)
	if err != nil {
		// FCM server error
		LogError.Error("FCM server error: " + err.Error())
		return false
	}

	res, err := client.Send(notification)
	if err != nil {
		// FCM server error
		LogError.Error("FCM server error: " + err.Error())
		return false
	}

	LogAccess.Debug(fmt.Sprintf("Android Success count: %d, Failure count: %d", res.Success, res.Failure))
	StatStorage.AddAndroidSuccess(int64(res.Success))
	StatStorage.AddAndroidError(int64(res.Failure))

	var newTokens []string
	for k, result := range res.Results {
		if result.Error != nil {
			isError = true
			newTokens = append(newTokens, req.Tokens[k])
			LogPush(FailedPush, req.Tokens[k], req, result.Error)
			if PushConf.Core.Sync {
				req.AddLog(getLogPushEntry(FailedPush, req.Tokens[k], req, result.Error))
			}
			continue
		}

		LogPush(SucceededPush, req.Tokens[k], req, nil)
	}

	if isError == true && retryCount < maxRetry {
		retryCount++

		// resend fail token
		req.Tokens = newTokens
		goto Retry
	}

	return isError
}
