package gorush

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"encoding/json"

	"github.com/msalihkarakasli/go-hms-push/push/model"
	"github.com/msalihkarakasli/go-hms-push/push/core"
	"github.com/msalihkarakasli/go-hms-push/push/config"
)

var (
	pushError   error
	pushClient  *core.HMSClient
	once        sync.Once
)

func GetPushClient(conf *config.Config) (*core.HMSClient, error){
	once.Do(func() {
		client, err := core.NewHttpClient(conf)
		if err != nil {
			fmt.Printf("Failed to new common client! Error is %s\n", err.Error())
			panic(err)
		}
		pushClient = client
		pushError = err
	})

	return pushClient, pushError
}

// InitHMSClient use for initialize HMS Client.
func InitHMSClient(apiKey string, appId string) (*core.HMSClient, error) {

	if apiKey == "" {
		return nil, errors.New("Missing Huawei API Key")
	}

	if appId == "" {
		return nil, errors.New("Missing Huawei App Id")
	}

	var conf = &config.Config{
		AppId:     appId,
		AppSecret: apiKey,
		AuthUrl: "https://login.cloud.huawei.com/oauth2/v2/token",
		PushUrl: "https://api.push.hicloud.com",
	}

	if apiKey != PushConf.Huawei.APIKey || appId != PushConf.Huawei.APPId {
		return GetPushClient(conf)
	}

	if HMSClient == nil {
		return GetPushClient(conf)
	}

	return HMSClient, nil
}

// GetHuaweiNotification use for define HMS notification.
// HTTP Connection Server Reference for HMS
// https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi
func GetHuaweiNotification(req PushNotification) (*model.MessageRequest, error) {
	
	msgRequest := model.NewNotificationMsgRequest()
	msgRequest.Message.Token = req.Tokens
	msgRequest.Message.Android = model.GetDefaultAndroid()
	msgRequest.Message.Android.Notification = model.GetDefaultAndroidNotification()

	b, err := json.Marshal(msgRequest)
	if err != nil {
		fmt.Printf("Failed to marshal the default message! Error is %s\n", err.Error())
		return nil, err
	}

	fmt.Printf("Default message is %s\n", string(b))
	return msgRequest, nil

	//TODO: Fix This part

	// notification := &fcm.Message{
	// 	To:                    req.To,
	// 	Condition:             req.Condition,
	// 	CollapseKey:           req.CollapseKey,
	// 	ContentAvailable:      req.ContentAvailable,
	// 	MutableContent:        req.MutableContent,
	// 	DelayWhileIdle:        req.DelayWhileIdle,
	// 	TimeToLive:            req.TimeToLive,
	// 	RestrictedPackageName: req.RestrictedPackageName,
	// 	DryRun:                req.DryRun,
	// }

	// if len(req.Tokens) > 0 {
	// 	notification.RegistrationIDs = req.Tokens
	// }

	// if len(req.Priority) > 0 && req.Priority == "high" {
	// 	notification.Priority = "high"
	// }

	// // Add another field
	// if len(req.Data) > 0 {
	// 	notification.Data = make(map[string]interface{})
	// 	for k, v := range req.Data {
	// 		notification.Data[k] = v
	// 	}
	// }

	// n := &fcm.Notification{}
	// isNotificationSet := false
	// if req.Notification != nil {
	// 	isNotificationSet = true
	// 	n = req.Notification
	// }

	// if len(req.Message) > 0 {
	// 	isNotificationSet = true
	// 	n.Body = req.Message
	// }

	// if len(req.Title) > 0 {
	// 	isNotificationSet = true
	// 	n.Title = req.Title
	// }

	// if len(req.Image) > 0 {
	// 	isNotificationSet = true
	// 	n.Image = req.Image
	// }

	// if v, ok := req.Sound.(string); ok && len(v) > 0 {
	// 	isNotificationSet = true
	// 	n.Sound = v
	// }

	// if isNotificationSet {
	// 	notification.Notification = n
	// }

	// // handle iOS apns in fcm

	// if len(req.Apns) > 0 {
	// 	notification.Apns = req.Apns
	// }

	// return notification
}

// PushToHuawei provide send notification to Android server.
func PushToHuawei(req PushNotification) bool {
	LogAccess.Debug("Start push notification for Huawei")

	var (
		client     *core.HMSClient
		retryCount = 0
		maxRetry   = PushConf.Huawei.MaxRetry
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

	notification, err := GetHuaweiNotification(req)

	client, err = InitHMSClient(PushConf.Huawei.APIKey, PushConf.Huawei.APPId)

	if err != nil {
		// HMS server error
		LogError.Error("HMS server error: " + err.Error())
		return false
	}

	res, err := client.SendMessage(context.Background(), notification)
	if err != nil {
		// Send Message error
		LogError.Error("HMS server send message error: " + err.Error())
		return false
	}

	fmt.Println(res);
	// if !req.IsTopic() {
	// 	LogAccess.Debug(fmt.Sprintf("Android Success count: %d, Failure count: %d", res.Success, res.Failure))
	// }

	
	//StatStorage.AddAndroidSuccess(int64(res.Success))
	//StatStorage.AddAndroidError(int64(res.Failure))

	//var newTokens []string
	// result from Send messages to specific devices
	// for k, result := range res.Results {
	// 	to := ""
	// 	if k < len(req.Tokens) {
	// 		to = req.Tokens[k]
	// 	} else {
	// 		to = req.To
	// 	}

	// 	if result.Error != nil {
	// 		// We should retry only "retryable" statuses. More info about response:
	// 		// https://firebase.google.com/docs/cloud-messaging/http-server-ref#downstream-http-messages-plain-text
	// 		if !result.Unregistered() {
	// 			newTokens = append(newTokens, to)
	// 		}
	// 		isError = true

	// 		LogPush(FailedPush, to, req, result.Error)
	// 		if PushConf.Core.Sync {
	// 			req.AddLog(getLogPushEntry(FailedPush, to, req, result.Error))
	// 		} else if PushConf.Core.FeedbackURL != "" {
	// 			go func(logger *logrus.Logger, log LogPushEntry, url string, timeout int64) {
	// 				err := DispatchFeedback(log, url, timeout)
	// 				if err != nil {
	// 					logger.Error(err)
	// 				}
	// 			}(LogError, getLogPushEntry(FailedPush, to, req, result.Error), PushConf.Core.FeedbackURL, PushConf.Core.FeedbackTimeout)
	// 		}
	// 		continue
	// 	}

	// 	LogPush(SucceededPush, to, req, nil)
	// }

	// result from Send messages to topics
	// if req.IsTopic() {
	// 	to := ""
	// 	if req.To != "" {
	// 		to = req.To
	// 	} else {
	// 		to = req.Condition
	// 	}
	// 	LogAccess.Debug("Send Topic Message: ", to)
	// 	// Success
	// 	if res.MessageID != 0 {
	// 		LogPush(SucceededPush, to, req, nil)
	// 	} else {
	// 		isError = true
	// 		// failure
	// 		LogPush(FailedPush, to, req, res.Error)
	// 		if PushConf.Core.Sync {
	// 			req.AddLog(getLogPushEntry(FailedPush, to, req, res.Error))
	// 		}
	// 	}
	// }

	// // Device Group HTTP Response
	// if len(res.FailedRegistrationIDs) > 0 {
	// 	isError = true
	// 	newTokens = append(newTokens, res.FailedRegistrationIDs...)

	// 	LogPush(FailedPush, notification.To, req, errors.New("device group: partial success or all fails"))
	// 	if PushConf.Core.Sync {
	// 		req.AddLog(getLogPushEntry(FailedPush, notification.To, req, errors.New("device group: partial success or all fails")))
	// 	}
	// }

	if isError && retryCount < maxRetry {
	 	retryCount++

	 	// resend fail token
	 	//req.Tokens = newTokens
	 	goto Retry
	}

	return false;
	//return isError
}
