package notify

import (
	"context"
	"errors"
	"sync"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/status"

	c "github.com/appleboy/go-hms-push/push/config"
	client "github.com/appleboy/go-hms-push/push/core"
	"github.com/appleboy/go-hms-push/push/model"
)

var (
	pushError  error
	pushClient *client.HMSClient
	once       sync.Once
)

// GetPushClient use for create HMS Push.
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
		return nil, errors.New("missing huawei app secret")
	}

	if appID == "" {
		return nil, errors.New("missing huawei app id")
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

// setHuaweiMessageTarget sets the target (tokens, topic, condition) on the message.
func setHuaweiMessageTarget(msg *model.Message, req *PushNotification) {
	if len(req.Tokens) > 0 {
		msg.Token = req.Tokens
	}
	if len(req.Topic) > 0 {
		msg.Topic = req.Topic
	}
	if len(req.Condition) > 0 {
		msg.Condition = req.Condition
	}
}

// setHuaweiAndroidConfig sets Android-specific configuration on the message.
func setHuaweiAndroidConfig(android *model.AndroidConfig, req *PushNotification) {
	if req.Priority == HIGH {
		android.Urgency = "HIGH"
	}
	android.CollapseKey = req.HuaweiCollapseKey
	if len(req.Category) > 0 {
		android.Category = req.Category
	}
	if len(req.HuaweiTTL) > 0 {
		android.TTL = req.HuaweiTTL
	}
	if len(req.BiTag) > 0 {
		android.BiTag = req.BiTag
	}
	android.FastAppTarget = req.FastAppTarget
}

// setHuaweiNotificationContent sets the notification content fields.
func setHuaweiNotificationContent(android *model.AndroidConfig, req *PushNotification) {
	ensureNotification := func() {
		if android.Notification == nil {
			android.Notification = model.GetDefaultAndroidNotification()
		}
	}

	if len(req.Message) > 0 {
		ensureNotification()
		android.Notification.Body = req.Message
	}
	if len(req.Title) > 0 {
		ensureNotification()
		android.Notification.Title = req.Title
	}
	if len(req.Image) > 0 {
		ensureNotification()
		android.Notification.Image = req.Image
	}
	if v, ok := req.Sound.(string); ok && len(v) > 0 {
		ensureNotification()
		android.Notification.Sound = v
	} else if android.Notification != nil {
		android.Notification.DefaultSound = true
	}
}

// GetHuaweiNotification use for define HMS notification.
// HTTP Connection Server Reference for HMS
// https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi
func GetHuaweiNotification(req *PushNotification) (*model.MessageRequest, error) {
	msgRequest := model.NewNotificationMsgRequest()
	msgRequest.Message.Android = model.GetDefaultAndroid()

	setHuaweiMessageTarget(msgRequest.Message, req)
	setHuaweiAndroidConfig(msgRequest.Message.Android, req)

	if len(req.HuaweiData) > 0 {
		msgRequest.Message.Data = req.HuaweiData
	}

	if req.HuaweiNotification != nil {
		msgRequest.Message.Android.Notification = req.HuaweiNotification
		if msgRequest.Message.Android.Notification.ClickAction == nil {
			msgRequest.Message.Android.Notification.ClickAction = model.GetDefaultClickAction()
		}
	}

	setHuaweiNotificationContent(msgRequest.Message.Android, req)

	b, err := json.Marshal(msgRequest)
	if err != nil {
		logx.LogError.Error("Failed to marshal the default message! Error is " + err.Error())
		return nil, err
	}

	logx.LogAccess.Debugf("Default message is %s", string(b))
	return msgRequest, nil
}

// PushToHuawei provide send notification to Android server.
func PushToHuawei(
	ctx context.Context,
	req *PushNotification,
	cfg *config.ConfYaml,
) (resp *ResponsePush, err error) {
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
		return nil, err
	}

	client, err = InitHMSClient(cfg, cfg.Huawei.AppSecret, cfg.Huawei.AppID)
	if err != nil {
		// HMS server error
		logx.LogError.Error("HMS server error: " + err.Error())
		return nil, err
	}

	resp = &ResponsePush{}

Retry:
	isError := false

	notification, _ := GetHuaweiNotification(req)

	res, err := client.SendMessage(ctx, notification)
	if err != nil {
		// Send Message error
		errLog := logPush(cfg, core.FailedPush, req.Topic, req, err)
		resp.Logs = append(resp.Logs, errLog)
		logx.LogError.Error("HMS server send message error: " + err.Error())
		return resp, err
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
