package gorush

import (
	"errors"
	"fmt"
	"github.com/google/go-gcm"
	apns "github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"time"
)

// D provide string array
type D map[string]interface{}

const (
	// ApnsPriorityLow will tell APNs to send the push message at a time that takes
	// into account power considerations for the device. Notifications with this
	// priority might be grouped and delivered in bursts. They are throttled, and
	// in some cases are not delivered.
	ApnsPriorityLow = 5

	// ApnsPriorityHigh will tell APNs to send the push message immediately.
	// Notifications with this priority must trigger an alert, sound, or badge on
	// the target device. It is an error to use this priority for a push
	// notification that contains only the content-available key.
	ApnsPriorityHigh = 10
)

// Alert is APNs payload
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

// RequestPush support multiple notification request.
type RequestPush struct {
	Notifications []PushNotification `json:"notifications" binding:"required"`
}

// PushNotification is single notification request
type PushNotification struct {
	// Common
	Tokens           []string `json:"tokens" binding:"required"`
	Platform         int      `json:"platform" binding:"required"`
	Message          string   `json:"message" binding:"required"`
	Title            string   `json:"title,omitempty"`
	Priority         string   `json:"priority,omitempty"`
	ContentAvailable bool     `json:"content_available,omitempty"`
	Sound            string   `json:"sound,omitempty"`
	Data             D        `json:"data,omitempty"`

	// Android
	APIKey                string           `json:"api_key,omitempty"`
	To                    string           `json:"to,omitempty"`
	CollapseKey           string           `json:"collapse_key,omitempty"`
	DelayWhileIdle        bool             `json:"delay_while_idle,omitempty"`
	TimeToLive            uint             `json:"time_to_live,omitempty"`
	RestrictedPackageName string           `json:"restricted_package_name,omitempty"`
	DryRun                bool             `json:"dry_run,omitempty"`
	Notification          gcm.Notification `json:"notification,omitempty"`

	// iOS
	Expiration int64    `json:"expiration,omitempty"`
	ApnsID     string   `json:"apns_id,omitempty"`
	Topic      string   `json:"topic,omitempty"`
	Badge      int      `json:"badge,omitempty"`
	Category   string   `json:"category,omitempty"`
	URLArgs    []string `json:"url-args,omitempty"`
	Alert      Alert    `json:"alert,omitempty"`
}

// CheckPushConf provide check your yml config.
func CheckPushConf() error {
	if !PushConf.Ios.Enabled && !PushConf.Android.Enabled {
		return errors.New("Please enable iOS or Android config in yml config")
	}

	if PushConf.Ios.Enabled {
		if PushConf.Ios.PemKeyPath == "" {
			return errors.New("Missing iOS certificate path")
		}
	}

	if PushConf.Android.Enabled {
		if PushConf.Android.APIKey == "" {
			return errors.New("Missing Android API Key")
		}
	}

	return nil
}

// InitAPNSClient use for initialize APNs Client.
func InitAPNSClient() error {
	if PushConf.Ios.Enabled {
		var err error

		CertificatePemIos, err = certificate.FromPemFile(PushConf.Ios.PemKeyPath, "")

		if err != nil {
			LogError.Error("Cert Error:", err.Error())

			return err
		}

		if PushConf.Ios.Production {
			ApnsClient = apns.NewClient(CertificatePemIos).Production()
		} else {
			ApnsClient = apns.NewClient(CertificatePemIos).Development()
		}
	}

	return nil
}

// SendNotification provide send all push request.
func SendNotification(req RequestPush) int {
	var count int
	for _, notification := range req.Notifications {
		switch notification.Platform {
		case PlatFormIos:
			if !PushConf.Ios.Enabled {
				continue
			}

			count++
			go PushToIOS(notification)
		case PlatFormAndroid:
			if !PushConf.Android.Enabled {
				continue
			}

			count++
			go PushToAndroid(notification)
		}
	}

	return count
}

// GetIOSNotification use for define iOS notificaiton.
// The iOS Notification Payload
// ref: https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/TheNotificationPayload.html
func GetIOSNotification(req PushNotification) *apns.Notification {
	notification := &apns.Notification{}

	if len(req.ApnsID) > 0 {
		notification.ApnsID = req.ApnsID
	}

	if len(req.Topic) > 0 {
		notification.Topic = req.Topic
	}

	if req.Expiration > 0 {
		notification.Expiration = time.Unix(req.Expiration, 0)
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

	if len(req.Data) > 0 {
		for k, v := range req.Data {
			payload.Custom(k, v)
		}
	}

	// Alert dictionary

	if len(req.Title) > 0 {
		payload.AlertTitle(req.Title)
	}

	if len(req.Alert.TitleLocKey) > 0 {
		payload.AlertTitleLocKey(req.Alert.TitleLocKey)
	}

	if len(req.Alert.LocArgs) > 0 {
		payload.AlertLocArgs(req.Alert.LocArgs)
	}

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

// PushToIOS provide send notification to APNs server.
func PushToIOS(req PushNotification) bool {

	notification := GetIOSNotification(req)

	for _, token := range req.Tokens {
		notification.DeviceToken = token

		// send ios notification
		res, err := ApnsClient.Push(notification)

		if err != nil {
			// apns server error
			LogPush(FailedPush, token, req, err)

			return false
		}

		if res.StatusCode != 200 {
			// error message:
			// ref: https://github.com/sideshow/apns2/blob/master/response.go#L14-L65
			LogPush(FailedPush, token, req, errors.New(res.Reason))

			return false
		}

		if res.Sent() {
			LogPush(SucceededPush, token, req, nil)
		}
	}

	return true
}

// GetAndroidNotification use for define Android notificaiton.
// HTTP Connection Server Reference for Android
// https://developers.google.com/cloud-messaging/http-server-ref
func GetAndroidNotification(req PushNotification) gcm.HttpMessage {
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

	// Add another field
	if len(req.Data) > 0 {
		notification.Data = make(map[string]interface{})
		for k, v := range req.Data {
			notification.Data[k] = v
		}
	}

	notification.Notification = &req.Notification

	// Set request message if body is empty
	if len(notification.Notification.Body) == 0 {
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
	var APIKey string

	notification := GetAndroidNotification(req)

	if APIKey = PushConf.Android.APIKey; req.APIKey != "" {
		APIKey = req.APIKey
	}

	res, err := gcm.SendHttp(APIKey, notification)

	if err != nil {
		// GCM server error
		LogError.Error("GCM server error: " + err.Error())

		return false
	}

	LogAccess.Debug(fmt.Sprintf("Android Success count: %d, Failure count: %d", res.Success, res.Failure))

	for k, result := range res.Results {
		if result.Error != "" {
			LogPush(FailedPush, req.Tokens[k], req, errors.New(result.Error))
			continue
		}

		LogPush(SucceededPush, req.Tokens[k], req, nil)
	}

	return true
}
