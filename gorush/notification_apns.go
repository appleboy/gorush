package gorush

import (
	"errors"
	"path/filepath"
	"time"
	"os"
	apns "github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"github.com/eencloud/goeen/dhash"
	"log"
	"fmt"
)

type Watch struct {
	*dhash.WatchValue
	values []string
}

func RemoveToken(token string, esn string) {

	for i := 0; i<5; i++ { // Try up to 5 times
		ret := NewWatch(esn)
		if ret == nil {
			LogError.Error("Could not get watch for %s", esn)
		}

		err := ret.Get()

		if err == dhash.WatchEmpty {
			LogError.Error("Attempted to remove token from null list");
		} else if err != nil {
			LogError.Error("Dhash Error: ", err.Error())
			continue;
		} else {
			b := ret.values[:0]
			for i, x := range ret.values {
				if x == token {
					ret.values = append(ret.values[:i], ret.values[i+1:]...)
					break
				}
			}

			ret.values = b;

			saveErr := ret.Save()

			if saveErr != nil { // data most likely changed, retry the process
				LogError.Error("Save Error: ", saveErr.Error());
				continue
			} else {
				log.Printf("Saved string array: %v", ret.values)
				break;
			}
		}
	}
}

func NewWatch(esn string) *Watch {
	w := &Watch{ }
	str := fmt.Sprintf("com.eencloud.push_tokens.%s.ios", esn)
	dh, err := dhash.Resolve(str)
	log.Printf("dhash key: %s", str)
	if err != nil {
		LogError.Error("Failed to resolve dhash token", err.Error());
		return nil
	}

	w.WatchValue = dhash.NewWatchValue(dh, &w.values, str)
	return w
}

// InitAPNSClient use for initialize APNs Client.
func InitAPNSClient() error {
	var addr= os.Getenv("EEN_DHASH_ADDRESS");
	log.Printf("Initializing dhash service at %s", addr)

	dhash.Initialize(addr)



	if PushConf.Ios.Enabled {
		var err error
		ext := filepath.Ext(PushConf.Ios.KeyPath)

		switch ext {
		case ".p12":
			CertificatePemIos, err = certificate.FromP12File(PushConf.Ios.KeyPath, PushConf.Ios.Password)
		case ".pem":
			CertificatePemIos, err = certificate.FromPemFile(PushConf.Ios.KeyPath, PushConf.Ios.Password)
		default:
			err = errors.New("wrong certificate key extension")
		}

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

	return payload
}

// GetIOSNotification use for define iOS notification.
// The iOS Notification Payload
// ref: https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html#//apple_ref/doc/uid/TP40008194-CH17-SW1
func GetIOSNotification(req PushNotification) *apns.Notification {
	notification := &apns.Notification{
		ApnsID: req.ApnsID,
		Topic:  req.Topic,
	}

	if req.Expiration > 0 {
		notification.Expiration = time.Unix(req.Expiration, 0)
	}

	if len(req.Priority) > 0 && req.Priority == "normal" {
		notification.Priority = apns.PriorityLow
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

	payload.MutableContent()

	if len(req.Sound) > 0 {
		payload.Sound(req.Sound)
	}

	if req.ContentAvailable {
		payload.ContentAvailable()
	}

	if len(req.URLArgs) > 0 {
		payload.URLArgs(req.URLArgs)
	}

	for k, v := range req.Data {
		payload.Custom(k, v)
	}

	payload = iosAlertDictionary(payload, req)

	notification.Payload = payload

	return notification
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

	for i :=0; i < len(req.Tokens); i++ {
		token := req.Tokens[i]
		userId := req.UserIds[i]
		notification.DeviceToken = token

		// send ios notification
		res, err := ApnsClient.Push(notification)

		if err != nil {
			// apns server error
			LogPush(FailedPush, token, userId, req, err)
			if PushConf.Core.Sync {
				req.AddLog(getLogPushEntry(FailedPush, token, userId, req, err))
			}
			StatStorage.AddIosError(1)
			newTokens = append(newTokens, token)
			isError = true
			continue
		}

		if res.StatusCode != 200 {
			// error message:
			// ref: https://github.com/sideshow/apns2/blob/master/response.go#L14-L65
			LogPush(FailedPush, token, userId, req, errors.New(res.Reason))
			if PushConf.Core.Sync {
				req.AddLog(getLogPushEntry(FailedPush, token, userId, req, errors.New(res.Reason)))
			}
			StatStorage.AddIosError(1)
			newTokens = append(newTokens, token)
			isError = true

			reasons := []string{apns.ReasonBadDeviceToken, apns.ReasonDeviceTokenNotForTopic, apns.ReasonUnregistered}

			for _, a := range(reasons) {
				if (a == res.Reason) {
					go RemoveToken(token, userId)
					break
				}
			}
			continue
		}

		if res.Sent() {
			LogPush(SucceededPush, token, userId, req, nil)
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
