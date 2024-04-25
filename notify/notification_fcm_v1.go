package notify

import (
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/status"
	"google.golang.org/api/option"
	"time"
)

// Send messages and manage messaging subscriptions for your Firebase
// applications
const firebaseMessagingScope = "https://www.googleapis.com/auth/firebase.messaging"

var fcmV1Client *messaging.Client

func InitiFCMV1Client(ctx context.Context, cfg *config.ConfYaml) (*messaging.Client, error) {
	if fcmV1Client != nil {
		return fcmV1Client, nil
	}

	f, err := firebase.NewApp(ctx,
		&firebase.Config{
			ProjectID: cfg.Android.ProjectID,
		},
		option.WithCredentialsFile(cfg.Android.ServiceAccountKey),
		option.WithScopes(firebaseMessagingScope),
	)

	if err != nil {
		return nil, err
	}

	client, err := f.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	fcmV1Client = client
	return client, err
}

func PushToAndroidV1(ctx context.Context, req *PushNotification, cfg *config.ConfYaml) (resp *ResponsePush, err error) {
	logx.LogAccess.Debug("Start push notification for Android V1")

	// check message
	err = CheckMessage(req)
	if err != nil {
		logx.LogError.Error("request error: " + err.Error())
		return nil, err
	}

	resp = &ResponsePush{}

	notification, err := getAndroidNotificationV1(req)
	if err != nil {
		// FCM server error
		logx.LogError.Error("FCM V1 server error: " + err.Error())
		return resp, err
	}

	client, err := InitiFCMV1Client(ctx, cfg)
	if err != nil {
		// FCM server error
		logx.LogError.Error("FCM V1 server error: " + err.Error())
		return resp, err
	}

	res, err := client.SendEachForMulticast(ctx, notification)
	if err != nil {
		// Send Message error
		logx.LogError.Error("FCM server send message error: " + err.Error())

		for _, token := range req.Tokens {
			errLog := logPush(cfg, core.FailedPush, token, req, err)
			resp.Logs = append(resp.Logs, errLog)
		}

		status.StatStorage.AddAndroidError(int64(len(req.Tokens)))
		return resp, err
	}

	status.StatStorage.AddAndroidSuccess(int64(res.SuccessCount))
	status.StatStorage.AddAndroidError(int64(res.FailureCount))

	// result from Send messages to specific devices
	for k, result := range res.Responses {
		to := req.To
		if k < len(req.Tokens) {
			to = req.Tokens[k]
		}

		if result.Error != nil {
			errLog := logPush(cfg, core.FailedPush, to, req, result.Error)
			resp.Logs = append(resp.Logs, errLog)
			continue
		}

		logPush(cfg, core.SucceededPush, to, req, nil)
	}

	return resp, nil
}

func getAndroidNotificationV1(req *PushNotification) (*messaging.MulticastMessage, error) {

	androidNotification := &messaging.AndroidNotification{}
	if req.Notification != nil {
		notificationCount, err := req.Notification.NotificationCount()
		if err != nil {
			logx.LogError.Error("FCM unsupported badge value", err)
			return nil, errors.New("invalid badge format")
		}

		androidNotification = &messaging.AndroidNotification{
			Title:             req.Notification.Title,
			Body:              req.Notification.Body,
			ChannelID:         req.Notification.ChannelID,
			Icon:              req.Notification.Icon,
			ImageURL:          req.Notification.Image,
			Sound:             req.Notification.Sound,
			NotificationCount: notificationCount,
			Tag:               req.Notification.Tag,
			Color:             req.Notification.Color,
			ClickAction:       req.Notification.ClickAction,
			BodyLocKey:        req.Notification.BodyLocKey,
			BodyLocArgs:       req.Notification.BodyLocArgs,
			TitleLocKey:       req.Notification.TitleLocKey,
			TitleLocArgs:      req.Notification.TitleLocArgs,
			//Ticker:                "",
			//Sticky:                false,
			//EventTimestamp:        nil,
			//LocalOnly:             false,
			//Priority:              0,
			//VibrateTimingMillis:   nil,
			//DefaultVibrateTimings: false,
			//DefaultSound:          false,
			//LightSettings:         nil,
			//DefaultLightSettings:  false,
			//Visibility:            0,
			//NotificationCount:     nil,
		}
	}

	if androidNotification.Title == "" {
		androidNotification.Title = req.Title
	}

	if androidNotification.Body == "" {
		androidNotification.Body = req.Message
	}

	if androidNotification.ImageURL == "" {
		androidNotification.ImageURL = req.Image
	}

	if androidNotification.Sound == "" && req.Sound != nil {
		v, ok := req.Sound.(string)
		if !ok {
			logx.LogError.Errorf("FCM unsupported sound value: %#v", req.Sound)
			return nil, errors.New("invalid sound format")
		}
		androidNotification.Sound = v
	}

	data := make(map[string]string, len(req.Data))
	for k, val := range req.Data {
		v, ok := val.(string)
		if !ok {
			logx.LogError.Errorf("FCM unsupported data value for key %s. value: %#v", k, val)
			return nil, errors.New("invalid data format")
		}
		data[k] = v
	}

	android := &messaging.AndroidConfig{
		CollapseKey: req.CollapseKey,
		Priority:    req.Priority,
		TTL:         nil,
		//RestrictedPackageName: "",
		Data:         data,
		Notification: androidNotification,
		FCMOptions:   nil,
	}

	if req.TimeToLive != nil {
		ttl := time.Second * time.Duration(*req.TimeToLive)
		android.TTL = &ttl
	}

	m := &messaging.MulticastMessage{
		Data: data,
		Notification: &messaging.Notification{
			Title:    req.Title,
			Body:     req.Message,
			ImageURL: req.Image,
		},
		Android:    android,
		Webpush:    nil,
		APNS:       nil,
		FCMOptions: nil,
		Tokens:     req.Tokens,
	}

	return m, nil
}
