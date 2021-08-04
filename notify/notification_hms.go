package notify

import (
	"context"
	"errors"
	"sync"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/status"

	c "github.com/msalihkarakasli/go-hms-push/push/config"
	client "github.com/msalihkarakasli/go-hms-push/push/core"
	"github.com/msalihkarakasli/go-hms-push/push/model"
)

var (
	pushError  error
	pushClient *client.HMSClient
	once       sync.Once
)

// GetPushClient use for create HMS Push
func GetPushClient(conf *c.Config) (*client.HMSClient, error) {
	once.Do(func() {
		client, err := client.NewHttpClient(conf)
		if err != nil {
			panic(err)
		}
		pushClient = client
		pushError = err
	})

	return pushClient, pushError
}

// InitHMSClient use for initialize HMS Client.
func InitHMSClient(cfg *config.ConfYaml, appSecret, appID string) (*client.HMSClient, error) {
	if appSecret == "" {
		return nil, errors.New("Missing Huawei App Secret")
	}

	if appID == "" {
		return nil, errors.New("Missing Huawei App ID")
	}

	conf := &c.Config{
		AppId:     appID,
		AppSecret: appSecret,
		AuthUrl:   "https://oauth-login.cloud.huawei.com/oauth2/v3/token",
		PushUrl:   "https://push-api.cloud.huawei.com",
	}

	if appSecret != cfg.Huawei.AppSecret || appID != cfg.Huawei.AppID {
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
func GetHuaweiNotification(req *PushNotification) (*model.MessageRequest, error) {
	msgRequest := model.NewNotificationMsgRequest()

	msgRequest.Message.Android = model.GetDefaultAndroid()

	if len(req.Tokens) > 0 {
		msgRequest.Message.Token = req.Tokens
	}

	if len(req.Topic) > 0 {
		msgRequest.Message.Topic = req.Topic
	}

	if len(req.To) > 0 {
		msgRequest.Message.Topic = req.To
	}

	if len(req.Condition) > 0 {
		msgRequest.Message.Condition = req.Condition
	}

	if req.Priority == "high" {
		msgRequest.Message.Android.Urgency = "HIGH"
	}

	// if req.HuaweiCollapseKey != nil {
	msgRequest.Message.Android.CollapseKey = req.HuaweiCollapseKey
	//}

	if len(req.Category) > 0 {
		msgRequest.Message.Android.Category = req.Category
	}

	if len(req.HuaweiTTL) > 0 {
		msgRequest.Message.Android.TTL = req.HuaweiTTL
	}

	if len(req.BiTag) > 0 {
		msgRequest.Message.Android.BiTag = req.BiTag
	}

	msgRequest.Message.Android.FastAppTarget = req.FastAppTarget

	// Add data fields
	if len(req.HuaweiData) > 0 {
		msgRequest.Message.Data = req.HuaweiData
	} else {
		// Notification Message
		msgRequest.Message.Android.Notification = model.GetDefaultAndroidNotification()

		n := msgRequest.Message.Android.Notification
		isNotificationSet := false

		if req.HuaweiNotification != nil {
			isNotificationSet = true
			n = req.HuaweiNotification

			if n.ClickAction == nil {
				n.ClickAction = model.GetDefaultClickAction()
			}
		}

		if len(req.Message) > 0 {
			isNotificationSet = true
			n.Body = req.Message
		}

		if len(req.Title) > 0 {
			isNotificationSet = true
			n.Title = req.Title
		}

		if len(req.Image) > 0 {
			isNotificationSet = true
			n.Image = req.Image
		}

		if v, ok := req.Sound.(string); ok && len(v) > 0 {
			isNotificationSet = true
			n.Sound = v
		} else {
			n.DefaultSound = true
		}

		if isNotificationSet {
			msgRequest.Message.Android.Notification = n
		}
	}

	b, err := json.Marshal(msgRequest)
	if err != nil {
		logx.LogError.Error("Failed to marshal the default message! Error is " + err.Error())
		return nil, err
	}

	logx.LogAccess.Debugf("Default message is %s", string(b))
	return msgRequest, nil
}

// PushToHuawei provide send notification to Android server.
func PushToHuawei(req *PushNotification, cfg *config.ConfYaml) (resp *ResponsePush, err error) {
	logx.LogAccess.Debug("Start push notification for Huawei")

	var (
		client     *client.HMSClient
		retryCount = 0
		maxRetry   = cfg.Huawei.MaxRetry
	)

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	// check message
	err = CheckMessage(req)
	if err != nil {
		logx.LogError.Error("request error: " + err.Error())
		return
	}

	client, err = InitHMSClient(cfg, cfg.Huawei.AppSecret, cfg.Huawei.AppID)

	if err != nil {
		// HMS server error
		logx.LogError.Error("HMS server error: " + err.Error())
		return
	}

	resp = &ResponsePush{}

Retry:
	isError := false

	notification, _ := GetHuaweiNotification(req)

	res, err := client.SendMessage(context.Background(), notification)
	if err != nil {
		// Send Message error
		errLog := logPush(cfg, core.FailedPush, req.To, req, err)
		resp.Logs = append(resp.Logs, errLog)
		logx.LogError.Error("HMS server send message error: " + err.Error())
		return
	}

	// Huawei Push Send API does not support exact results for each token
	if res.Code == "80000000" {
		status.StatStorage.AddHuaweiSuccess(int64(1))
		logx.LogAccess.Debug("Huwaei Send Notification is completed successfully!")
	} else {
		isError = true
		status.StatStorage.AddHuaweiError(int64(1))
		logx.LogAccess.Debug("Huawei Send Notification is failed! Code: " + res.Code)
	}

	if isError && retryCount < maxRetry {
		retryCount++

		// resend all tokens
		goto Retry
	}

	return resp, nil
}
