package gorush

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/jaraxasoftware/gorush/web"
	"github.com/google/go-gcm"
	apns "github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
)

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
	Subtitle     string   `json:"subtitle,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
}

// RequestPush support multiple notification request.
type RequestPush struct {
	Notifications []PushNotification `json:"notifications" binding:"required"`
	Sync          bool               `json:"sync,omitempty"`
}

type Subscription struct {
	Endpoint  string    `json:"endpoint" binding:"required"`
	Key       string    `json:"key" binding:"required"`
	Auth      string    `json:"auth" binding:"required"`
}

// PushNotification is single notification request
type PushNotification struct {
	// Common
	Platform         int                     `json:"platform" binding:"required"`
	Message          string                  `json:"message,omitempty"`
	Title            string                  `json:"title,omitempty"`
	Priority         string                  `json:"priority,omitempty"`
	ContentAvailable bool                    `json:"content_available,omitempty"`
	Sound            string                  `json:"sound,omitempty"`
	Data             map[string]interface{}  `json:"data,omitempty"`
	Retry            int                     `json:"retry,omitempty"`
	// Android + iOS
	Tokens           []string `json:"tokens,omitempty`
	// Android + web
	APIKey           string   `json:"api_key,omitempty"`
	TimeToLive       *uint    `json:"time_to_live,omitempty"`

	// Android
	To                    string           `json:"to,omitempty"`
	CollapseKey           string           `json:"collapse_key,omitempty"`
	DelayWhileIdle        bool             `json:"delay_while_idle,omitempty"`
	RestrictedPackageName string           `json:"restricted_package_name,omitempty"`
	DryRun                bool             `json:"dry_run,omitempty"`
	Notification          gcm.Notification `json:"notification,omitempty"`

	// iOS
	Expiration     int64    `json:"expiration,omitempty"`
	ApnsID         string   `json:"apns_id,omitempty"`
	Topic          string   `json:"topic,omitempty"`
	Badge          *int     `json:"badge,omitempty"`
	Category       string   `json:"category,omitempty"`
	URLArgs        []string `json:"url-args,omitempty"`
	Alert          Alert    `json:"alert,omitempty"`
	MutableContent bool     `json:"mutable-content,omitempty"`
	Voip           bool     `json:"voip,omitempty"`

	// Web
	Subscriptions  []Subscription `json:"subscriptions,omitempty"`
}

// CheckMessage for check request message
func CheckMessage(req PushNotification) error {
	var msg string

    if req.Platform == PlatformIos || req.Platform == PlatformAndroid {
		if len(req.Tokens) == 0 {
			msg = "the message must specify at least one registration ID"
			LogAccess.Debug(msg)
			return errors.New(msg)
		}

		if req.Platform == PlatformIos && len(req.Tokens[0]) == 0 {
			msg = "the token must not be empty"
			LogAccess.Debug(msg)
			return errors.New(msg)
		}

		if req.Platform == PlatformAndroid && len(req.Tokens) > 1000 {
			msg = "the message may specify at most 1000 registration IDs"
			LogAccess.Debug(msg)
			return errors.New(msg)
		}

		// ref: https://developers.google.com/cloud-messaging/http-server-ref
		if req.Platform == PlatformAndroid && req.TimeToLive != nil && (*req.TimeToLive < uint(0) || uint(2419200) < *req.TimeToLive) {
			msg = "the message's TimeToLive field must be an integer " +
				"between 0 and 2419200 (4 weeks)"
			LogAccess.Debug(msg)
			return errors.New(msg)
		}
	} else if req.Platform == PlatformWeb {
		if len(req.Subscriptions) == 0 {
			msg = "the message must specify at least one subscription"
			LogAccess.Debug(msg)
			return errors.New(msg)
		}
	}

	return nil
}

// SetProxy only working for GCM server.
func SetProxy(proxy string) error {

	proxyURL, err := url.ParseRequestURI(proxy)

	if err != nil {
		return err
	}

	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	LogAccess.Debug("Set http proxy as " + proxy)

	return nil
}

// CheckPushConf provide check your yml config.
func CheckPushConf() error {
	if !PushConf.Ios.VoipEnabled && !PushConf.Ios.Enabled && !PushConf.Android.Enabled && !PushConf.Web.Enabled {
		return errors.New("Please enable iOS, VoIP iOS, Android or Web config in yml config")
	}

	if PushConf.Ios.Enabled {
		if PushConf.Ios.KeyPath == "" {
			return errors.New("Missing iOS certificate path")
		}
	}

	if PushConf.Ios.VoipEnabled {
		if PushConf.Ios.VoipKeyPath == "" {
			return errors.New("Missing VoIP iOS certificate path")
		}
	}

	if PushConf.Android.Enabled {
		if PushConf.Android.APIKey == "" {
			return errors.New("Missing Android API Key")
		}
	}

	if PushConf.Web.Enabled {
		if PushConf.Web.APIKey == "" {
			return errors.New("Missing GCM API Key for Chrome")
		}
	}

	return nil
}

// InitAPNSClient use for initialize APNs Client.
func InitAPNSClient() error {
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

	if PushConf.Ios.VoipEnabled {
		var err error
		ext := filepath.Ext(PushConf.Ios.VoipKeyPath)

		switch ext {
		case ".p12":
			VoipCertificatePemIos, err = certificate.FromP12File(PushConf.Ios.VoipKeyPath, PushConf.Ios.VoipPassword)
		case ".pem":
			VoipCertificatePemIos, err = certificate.FromPemFile(PushConf.Ios.VoipKeyPath, PushConf.Ios.VoipPassword)
		default:
			err = errors.New("wrong VoIP certificate key extension")
		}

		if err != nil {
			LogError.Error("Cert Error:", err.Error())

			return err
		}

		if PushConf.Ios.VoipProduction {
			VoipApnsClient = apns.NewClient(VoipCertificatePemIos).Production()
		} else {
			VoipApnsClient = apns.NewClient(VoipCertificatePemIos).Development()
		}
	}

	return nil
}

// InitWebClient use for initialize APNs Client.
func InitWebClient() error {
	if PushConf.Web.Enabled {
		//var err error
		WebClient = web.NewClient(PushConf.Web.APIKey)
	}

	return nil
}

// InitWorkers for initialize all workers.
func InitWorkers(workerNum int64, queueNum int64) {
	LogAccess.Debug("worker number is ", workerNum, ", queue number is ", queueNum)
	QueueNotification = make(chan PushNotification, queueNum)
	for i := int64(0); i < workerNum; i++ {
		go startWorker()
	}
}

func startWorker() {
	for {
		notification := <-QueueNotification
		switch notification.Platform {
		case PlatformIos:
			PushToIOS(notification)
		case PlatformAndroid:
			PushToAndroid(notification)
		case PlatformWeb:
			PushToWeb(notification)
		}
	}
}

// queueNotification add notification to queue list.
func queueNotification(req RequestPush) int {
	var count int
	for _, notification := range req.Notifications {
		switch notification.Platform {
		case PlatformIos:
			if !PushConf.Ios.Enabled && !notification.Voip {
				continue
			}
			if !PushConf.Ios.VoipEnabled && notification.Voip {
				continue
			}
		case PlatformAndroid:
			if !PushConf.Android.Enabled {
				continue
			}
		}
		QueueNotification <- notification

		count += len(notification.Tokens)
	}

	StatStorage.AddTotalCount(int64(count))

	return count
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

// GetIOSNotification use for define iOS notificaiton.
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

	if req.MutableContent {
		payload.MutableContent()
	}

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
	var isError bool
	_, isError = PushToIOSWithErrorResult(req)
	return isError
}

// PushToIOSWithErrorResult provide send notification to APNs server and return response array for failed requests.
func PushToIOSWithErrorResult(req PushNotification)  (*map[string]*apns.Response,bool) {
	LogAccess.Debug("Start push notification for iOS")

	var retryCount = 0
	var maxRetry = PushConf.Ios.MaxRetry

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	// check message
	err := CheckMessage(req)

	if err != nil {
		errorString := "request error: " + err.Error()
		LogError.Error(errorString)
		var returnResultList map[string]*apns.Response
		returnResultList = make(map[string]*apns.Response)
		for _,token := range req.Tokens {
			time := apns.Time{time.Now()}
			response := apns.Response{StatusCode: 500, Reason: errorString, Timestamp: time}
			returnResultList[token] = &response
		}
		return &returnResultList, true
	}

Retry:
	var isError = false
	var newTokens []string
	var returnResultList map[string]*apns.Response
	returnResultList = make(map[string]*apns.Response)

	notification := GetIOSNotification(req)

	for _, token := range req.Tokens {
		notification.DeviceToken = token

		// send ios notification
		var res *apns.Response
		var err error

		if req.Voip {
			res, err = VoipApnsClient.Push(notification)
		} else {
			res, err = ApnsClient.Push(notification)
		}

		if err != nil {
			// apns server error
			LogPush(FailedPush, token, req, err)
			StatStorage.AddIosError(1)
			newTokens = append(newTokens, token)
			returnResultList[token] = res
			isError = true
			continue
		}

		if res.StatusCode != 200 {
			// error message:
			// ref: https://github.com/sideshow/apns2/blob/master/response.go#L14-L65
			LogPush(FailedPush, token, req, errors.New(res.Reason))
			StatStorage.AddIosError(1)
			newTokens = append(newTokens, token)
			returnResultList[token] = res
			isError = true
			continue
		}

		if res.Sent() {
			LogPush(SucceededPush, token, req, nil)
			StatStorage.AddIosSuccess(1)
		}
	}

	if isError == true && retryCount < maxRetry {
		retryCount++

		// resend fail token
		req.Tokens = newTokens
		goto Retry
	}

	return &returnResultList,isError
}

// GetAndroidNotification use for define Android notificaiton.
// HTTP Connection Server Reference for Android
// https://developers.google.com/cloud-messaging/http-server-ref
func GetAndroidNotification(req PushNotification) gcm.HttpMessage {
	notification := gcm.HttpMessage{
		To:                    req.To,
		CollapseKey:           req.CollapseKey,
		ContentAvailable:      req.ContentAvailable,
		DelayWhileIdle:        req.DelayWhileIdle,
		TimeToLive:            req.TimeToLive,
		RestrictedPackageName: req.RestrictedPackageName,
		DryRun:                req.DryRun,
	}

	notification.RegistrationIds = req.Tokens

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
	var isError bool
	_, isError = PushToAndroidWithErrorResult(req)
	return isError
}

// PushToAndroidWithErrorResult provide send notification to Android server.
func PushToAndroidWithErrorResult(req PushNotification) (*map[string]string,bool) {
	LogAccess.Debug("Start push notification for Android")

	var APIKey string
	var retryCount = 0
	var maxRetry = PushConf.Android.MaxRetry

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	// check message
	err := CheckMessage(req)

	if err != nil {
		errorString := "request error: " + err.Error()
		LogError.Error(errorString)
		var returnResultList map[string]string
		returnResultList = make(map[string]string)
		for _,token := range req.Tokens {
			returnResultList[token] = errorString
		}
		return &returnResultList, true
	}

Retry:
	var isError = false
	notification := GetAndroidNotification(req)

	if APIKey = PushConf.Android.APIKey; req.APIKey != "" {
		APIKey = req.APIKey
	}

	res, err := gcm.SendHttp(APIKey, notification)

	if err != nil {
		// GCM server error
		errorString := "GCM server error: " + err.Error()
		LogError.Error(errorString)
		var returnResultList map[string]string
		returnResultList = make(map[string]string)
		for _,token := range req.Tokens {
			returnResultList[token] = errorString
		}
		return &returnResultList, true
	}

	LogAccess.Debug(fmt.Sprintf("Android Success count: %d, Failure count: %d", res.Success, res.Failure))
	StatStorage.AddAndroidSuccess(int64(res.Success))
	StatStorage.AddAndroidError(int64(res.Failure))

	var newTokens []string
	var returnResultList map[string]string
	returnResultList = make(map[string]string)
	for k, result := range res.Results {
		if result.Error != "" {
			isError = true
			newTokens = append(newTokens, req.Tokens[k])
			returnResultList[req.Tokens[k]] = result.Error
			LogPush(FailedPush, req.Tokens[k], req, errors.New(result.Error))
			continue
		}
		delete(returnResultList, req.Tokens[k])

		LogPush(SucceededPush, req.Tokens[k], req, nil)
	}

	if isError == true && retryCount < maxRetry {
		retryCount++

		// resend fail token
		req.Tokens = newTokens
		goto Retry
	}

	return &returnResultList, isError
}

func GetWebNotification(req PushNotification, subscription *Subscription) *web.Notification {
	notification := &web.Notification{
		Payload: &req.Data,
		Subscription: &web.Subscription{
			Endpoint: subscription.Endpoint, 
			Key: subscription.Key, 
			Auth: subscription.Auth,
		},
		TimeToLive: req.TimeToLive,
	}
	return notification
}

// PushToAndroid provide send notification to Android server.
func PushToWeb(req PushNotification) bool {
	var isError bool
	_, isError = PushToWebWithErrorResult(req)
	return isError
}

// PushToAndroidWithErrorResult provide send notification to Android server.
func PushToWebWithErrorResult(req PushNotification) (*map[string]*web.Response,bool) {
	LogAccess.Debug("Start push notification for Web")

	var retryCount = 0
	var maxRetry = PushConf.Web.MaxRetry

	if req.Retry > 0 && req.Retry < maxRetry {
		maxRetry = req.Retry
	}

	// check message
	err := CheckMessage(req)

	if err != nil {
		errorString := "request error: " + err.Error()
		LogError.Error(errorString)
		var returnResultList map[string]*web.Response
		returnResultList = make(map[string]*web.Response)
		for _,subscription := range req.Subscriptions {
			response := web.Response{StatusCode: 500, Body: errorString}
			returnResultList[subscription.Endpoint] = &response
		}
		return &returnResultList, true
	}

Retry:
	var isError = false
	var returnResultList map[string]*web.Response
	successCount := 0
	failureCount := 0
	returnResultList = make(map[string]*web.Response)

	for _, subscription := range req.Subscriptions {
		notification := GetWebNotification(req, &subscription)
		response , err := WebClient.Push(notification)
		if err != nil {
			failureCount++
			LogPush(FailedPush, subscription.Endpoint, req, err)
			fmt.Println(err)
			returnResultList[subscription.Endpoint] = response
		} else {
			successCount++
			LogPush(SucceededPush, subscription.Endpoint, req, nil)
		}
	}

	LogAccess.Debug(fmt.Sprintf("Web Success count: %d, Failure count: %d", successCount, failureCount))
	StatStorage.AddWebSuccess(int64(successCount))
	StatStorage.AddWebError(int64(failureCount))

	if isError == true && retryCount < maxRetry {
		retryCount++

		// resend fail token
		//FIXME
		//req.Tokens = newTokens
		goto Retry
	}	
	return	&returnResultList, isError
}