package gopush

import (
	"github.com/google/go-gcm"
	apns "github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"log"
)

type ExtendJSON struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

const (
	// PriorityLow will tell APNs to send the push message at a time that takes
	// into account power considerations for the device. Notifications with this
	// priority might be grouped and delivered in bursts. They are throttled, and
	// in some cases are not delivered.
	ApnsPriorityLow = 5

	// PriorityHigh will tell APNs to send the push message immediately.
	// Notifications with this priority must trigger an alert, sound, or badge on
	// the target device. It is an error to use this priority for a push
	// notification that contains only the content-available key.
	ApnsPriorityHigh = 10
)

type Alert struct {
	Action       string   `json:"action,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	Body         string   `json:"body,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	Title        string   `json:"title,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
}

type RequestPushNotification struct {
	// Common
	Tokens           []string `json:"tokens" binding:"required"`
	Platform         int      `json:"platform" binding:"required"`
	Message          string   `json:"message" binding:"required"`
	Priority         string   `json:"priority,omitempty"`
	ContentAvailable bool     `json:"content_available,omitempty"`

	// Android
	To                    string           `json:"to,omitempty"`
	CollapseKey           string           `json:"collapse_key,omitempty"`
	DelayWhileIdle        bool             `json:"delay_while_idle,omitempty"`
	TimeToLive            uint             `json:"time_to_live,omitempty"`
	RestrictedPackageName string           `json:"restricted_package_name,omitempty"`
	DryRun                bool             `json:"dry_run,omitempty"`
	Data                  gcm.Data         `json:"data,omitempty"`
	Notification          gcm.Notification `json:"notification,omitempty"`

	// iOS
	ApnsID   string       `json:"apns_id,omitempty"`
	Topic    string       `json:"topic,omitempty"`
	Badge    int          `json:"badge,omitempty"`
	Sound    string       `json:"sound,omitempty"`
	Category string       `json:"category,omitempty"`
	URLArgs  []string     `json:"url-args,omitempty"`
	Extend   []ExtendJSON `json:"extend,omitempty"`
	Alert    Alert        `json:"alert,omitempty"`

	// meta
	IDs []uint64 `json:"seq_id,omitempty"`
}

func InitAPNSClient() {
	if PushConf.Ios.Enabled {
		var err error

		CertificatePemIos, err = certificate.FromPemFile(PushConf.Ios.PemKeyPath, "")

		if err != nil {
			log.Println("Cert Error:", err)

			return
		}

		if PushConf.Ios.Production {
			ApnsClient = apns.NewClient(CertificatePemIos).Production()
		} else {
			ApnsClient = apns.NewClient(CertificatePemIos).Development()
		}
	}
}

func pushNotification(notification RequestPushNotification) bool {
	var (
		success bool
	)

	switch notification.Platform {
	case PlatFormIos:
		if !PushConf.Ios.Enabled {
			return false
		}
		success = pushNotificationIos(notification)
	case PlatFormAndroid:
		if !PushConf.Android.Enabled {
			return false
		}
		success = pushNotificationAndroid(notification)
	}

	return success
}

func GetIOSNotification(req RequestPushNotification) *apns.Notification {
	notification := &apns.Notification{}

	if len(req.ApnsID) > 0 {
			notification.ApnsID = req.ApnsID
	}

	if len(req.Topic) > 0 {
		notification.Topic = req.Topic
	}

	if len(req.Priority) > 0 && req.Priority == "normal" {
		notification.Priority = apns.PriorityLow
	}

	payload := payload.NewPayload().Alert(req.Message)

	if req.Badge > 0 {
		payload.Badge(req.Badge)
	}

	if len(req.Sound) > 0 {
		payload.Sound(req.Sound)
	}

	if req.ContentAvailable {
		payload.ContentAvailable()
	}

	if len(req.Extend) > 0 {
		for _, extend := range req.Extend {
			payload.Custom(extend.Key, extend.Value)
		}
	}

	// Alert dictionary

	if len(req.Alert.Title) > 0 {
		payload.AlertTitle(req.Alert.Title)
	}

	if len(req.Alert.TitleLocKey) > 0 {
		payload.AlertTitleLocKey(req.Alert.TitleLocKey)
	}

	// if len(req.Alert.LocArgs) > 0 {
	// 	payload.AlertLocArgs(req.Alert.LocArgs)
	// }

	if len(req.Alert.TitleLocArgs) > 0 {
		payload.AlertTitleLocArgs(req.Alert.TitleLocArgs)
	}

	if len(req.Alert.Body) > 0 {
		payload.AlertBody(req.Alert.Body)
	}

	if len(req.Alert.LaunchImage) > 0 {
		payload.AlertLaunchImage(req.Alert.LaunchImage)
	}

	if len(req.Alert.LocKey) > 0 {
		payload.AlertLocKey(req.Alert.LocKey)
	}

	if len(req.Alert.Action) > 0 {
		payload.AlertAction(req.Alert.Action)
	}

	if len(req.Alert.ActionLocKey) > 0 {
		payload.AlertActionLocKey(req.Alert.ActionLocKey)
	}

	// General

	if len(req.Category) > 0 {
		payload.Category(req.Category)
	}

	if len(req.URLArgs) > 0 {
		payload.URLArgs(req.URLArgs)
	}

	notification.Payload = payload

	return notification
}

func pushNotificationIos(req RequestPushNotification) bool {

	notification := GetIOSNotification(req)

	// The Remote Notification Payload
	// https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/TheNotificationPayload.html
	for _, token := range req.Tokens {
		notification.DeviceToken = token

		// send ios notification
		res, err := ApnsClient.Push(notification)

		if err != nil {
			log.Println("There was an error: ", err)

			return false
		}

		if res.Sent() {
			log.Println("APNs ID:", res.ApnsID)

			return true
		}
	}

	return true
}

func pushNotificationAndroid(req RequestPushNotification) bool {

	// HTTP Connection Server Reference for Android
	// https://developers.google.com/cloud-messaging/http-server-ref
	notification := gcm.HttpMessage{}

	notification.RegistrationIds = req.Tokens

	if len(req.To) > 0 {
		notification.To = req.To
	}

	if len(req.Priority) > 0 && req.Priority == "high" {
		notification.Priority = "high"
	}

	if len(req.CollapseKey) > 0 {
		notification.CollapseKey = req.CollapseKey
	}

	if req.ContentAvailable {
		notification.ContentAvailable = true
	}

	if req.DelayWhileIdle {
		notification.DelayWhileIdle = true
	}

	if req.TimeToLive > 0 {
		notification.TimeToLive = req.TimeToLive
	}

	if len(req.RestrictedPackageName) > 0 {
		notification.RestrictedPackageName = req.RestrictedPackageName
	}

	if req.DryRun {
		notification.DryRun = true
	}

	if len(req.Data) > 0 {
		notification.Data = req.Data
	}

	notification.Notification = &req.Notification

	// Set request message if body is empty
	if len(notification.Notification.Body) == 0 {
		notification.Notification.Body = req.Message
	}

	res, err := gcm.SendHttp(PushConf.Android.ApiKey, notification)

	if err != nil {
		log.Println(err)

		return false
	}

	if res.Error != "" {
		log.Println("GCM Error Message: " + res.Error)
	}

	if res.Success > 0 {
		log.Printf("Success count: %d, Failure count: %d", res.Success, res.Failure)

		return true
	}

	return true
}
