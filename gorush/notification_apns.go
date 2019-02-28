package gorush

import (
	"crypto/ecdsa"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"path/filepath"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"
)

// Sound sets the aps sound on the payload.
type Sound struct {
	Critical int     `json:"critical,omitempty"`
	Name     string  `json:"name,omitempty"`
	Volume   float32 `json:"volume,omitempty"`
}

// InitAPNSClient use for initialize APNs Client.
func InitAPNSClient() error {
	if PushConf.Ios.Enabled {
		var err error
		var authKey *ecdsa.PrivateKey
		var certificateKey tls.Certificate
		var ext string

		if PushConf.Ios.KeyPath != "" {
			ext = filepath.Ext(PushConf.Ios.KeyPath)

			switch ext {
			case ".p12":
				certificateKey, err = certificate.FromP12File(PushConf.Ios.KeyPath, PushConf.Ios.Password)
			case ".pem":
				certificateKey, err = certificate.FromPemFile(PushConf.Ios.KeyPath, PushConf.Ios.Password)
			case ".p8":
				authKey, err = token.AuthKeyFromFile(PushConf.Ios.KeyPath)
			default:
				err = errors.New("wrong certificate key extension")
			}

			if err != nil {
				LogError.Error("Cert Error:", err.Error())

				return err
			}
		} else if PushConf.Ios.KeyBase64 != "" {
			ext = "." + PushConf.Ios.KeyType
			key, err := base64.StdEncoding.DecodeString(PushConf.Ios.KeyBase64)
			if err != nil {
				LogError.Error("base64 decode error:", err.Error())

				return err
			}
			switch ext {
			case ".p12":
				certificateKey, err = certificate.FromP12Bytes(key, PushConf.Ios.Password)
			case ".pem":
				certificateKey, err = certificate.FromPemBytes(key, PushConf.Ios.Password)
			case ".p8":
				authKey, err = token.AuthKeyFromBytes(key)
			default:
				err = errors.New("wrong certificate key type")
			}

			if err != nil {
				LogError.Error("Cert Error:", err.Error())

				return err
			}
		}

		if ext == ".p8" && PushConf.Ios.KeyID != "" && PushConf.Ios.TeamID != "" {
			token := &token.Token{
				AuthKey: authKey,
				// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
				KeyID: PushConf.Ios.KeyID,
				// TeamID from developer account (View Account -> Membership)
				TeamID: PushConf.Ios.TeamID,
			}
			if PushConf.Ios.Production {
				ApnsClient = apns2.NewTokenClient(token).Production()
			} else {
				ApnsClient = apns2.NewTokenClient(token).Development()
			}
		} else {
			if PushConf.Ios.Production {
				ApnsClient = apns2.NewClient(certificateKey).Production()
			} else {
				ApnsClient = apns2.NewClient(certificateKey).Development()
			}
		}
	}

	return nil
}

func iosAlertDictionary(payload *payload.Payload, req PushNotification) *payload.Payload {
	// Alert dictionary

	if len(req.Title) > 0 {
		payload.AlertTitle(req.Title)
	}

	if len(req.Alert.Title) > 0 {
		payload.AlertTitle(req.Alert.Title)
	}

	// Apple Watch & Safari display this string as part of the notification interface.
	if len(req.Alert.Subtitle) > 0 {
		payload.AlertSubtitle(req.Alert.Subtitle)
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

	if len(req.Alert.SummaryArg) > 0 {
		payload.AlertSummaryArg(req.Alert.SummaryArg)
	}

	if req.Alert.SummaryArgCount > 0 {
		payload.AlertSummaryArgCount(req.Alert.SummaryArgCount)
	}

	return payload
}

// GetIOSNotification use for define iOS notification.
// The iOS Notification Payload
// ref: https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html#//apple_ref/doc/uid/TP40008194-CH17-SW1
func GetIOSNotification(req PushNotification) *apns2.Notification {
	notification := &apns2.Notification{
		ApnsID:     req.ApnsID,
		Topic:      req.Topic,
		CollapseID: req.CollapseID,
	}

	if req.Expiration > 0 {
		notification.Expiration = time.Unix(req.Expiration, 0)
	}

	if len(req.Priority) > 0 && req.Priority == "normal" {
		notification.Priority = apns2.PriorityLow
	}

	payload := payload.NewPayload()

	// add alert object if message length > 0
	if len(req.Message) > 0 {
		payload.Alert(req.Message)
	}

	// zero value for clear the badge on the app icon.
	if req.Badge != nil && *req.Badge >= 0 {
		payload.Badge(*req.Badge)
	}

	if req.MutableContent {
		payload.MutableContent()
	}

	switch req.Sound.(type) {
	// from http request binding
	case map[string]interface{}:
		result := &Sound{}
		_ = mapstructure.Decode(req.Sound, &result)
		payload.Sound(result)
	// from http request binding for non critical alerts
	case string:
		payload.Sound(&req.Sound)
	case Sound:
		payload.Sound(&req.Sound)
	}

	if len(req.SoundName) > 0 {
		payload.SoundName(req.SoundName)
	}

	if req.SoundVolume > 0 {
		payload.SoundVolume(req.SoundVolume)
	}

	if req.ContentAvailable {
		payload.ContentAvailable()
	}

	if len(req.URLArgs) > 0 {
		payload.URLArgs(req.URLArgs)
	}

	if len(req.ThreadID) > 0 {
		payload.ThreadID(req.ThreadID)
	}

	for k, v := range req.Data {
		payload.Custom(k, v)
	}

	payload = iosAlertDictionary(payload, req)

	notification.Payload = payload

	return notification
}

func getApnsClient(req PushNotification) (client *apns2.Client) {
	if req.Production {
		client = ApnsClient.Production()
	} else if req.Development {
		client = ApnsClient.Development()
	} else {
		if PushConf.Ios.Production {
			client = ApnsClient.Production()
		} else {
			client = ApnsClient.Development()
		}
	}
	return
}

// PushToIOS provide send notification to APNs server.
func PushToIOS(req PushNotification) bool {
	LogAccess.Debug("Start push notification for iOS")
	if PushConf.Core.Sync {
		defer req.WaitDone()
	}

	var (
		retryCount = 0
		maxRetry   = PushConf.Ios.MaxRetry
	)

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

Retry:
	var (
		isError   = false
		newTokens []string
	)

	notification := GetIOSNotification(req)
	client := getApnsClient(req)

	for _, token := range req.Tokens {
		notification.DeviceToken = token

		// send ios notification
		res, err := client.Push(notification)

		if err != nil {
			// apns server error
			LogPush(FailedPush, token, req, err)
			if PushConf.Core.Sync {
				req.AddLog(getLogPushEntry(FailedPush, token, req, err))
			}
			StatStorage.AddIosError(1)
			newTokens = append(newTokens, token)
			isError = true
			continue
		}

		if res.StatusCode != 200 {
			// error message:
			// ref: https://github.com/sideshow/apns2/blob/master/response.go#L14-L65
			LogPush(FailedPush, token, req, errors.New(res.Reason))
			if PushConf.Core.Sync {
				req.AddLog(getLogPushEntry(FailedPush, token, req, errors.New(res.Reason)))
			}
			StatStorage.AddIosError(1)
			newTokens = append(newTokens, token)
			isError = true
			continue
		}

		if res.Sent() {
			LogPush(SucceededPush, token, req, nil)
			StatStorage.AddIosSuccess(1)
		}
	}

	if isError && retryCount < maxRetry {
		retryCount++

		// resend fail token
		req.Tokens = newTokens
		goto Retry
	}

	return isError
}
