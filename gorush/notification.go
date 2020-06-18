package gorush

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/appleboy/gorush/gorush/consts"
	"github.com/appleboy/gorush/gorush/structs"
)

// RequestPush support multiple notification request.
type RequestPush struct {
	Notifications []PushNotification `json:"notifications" binding:"required"`
}

// PushNotification is single notification request
type PushNotification struct {
	wg  *sync.WaitGroup
	log *[]LogPushEntry

	structs.PushNotification
}

// WaitDone decrements the WaitGroup counter.
func (p *PushNotification) WaitDone() {
	if p.wg != nil {
		p.wg.Done()
	}
}

// AddWaitCount increments the WaitGroup counter.
func (p *PushNotification) AddWaitCount() {
	if p.wg != nil {
		p.wg.Add(1)
	}
}

// AddLog record fail log of notification
func (p *PushNotification) AddLog(log LogPushEntry) {
	if p.log != nil {
		*p.log = append(*p.log, log)
	}
}

// IsTopic check if message format is topic for FCM
// ref: https://firebase.google.com/docs/cloud-messaging/send-message#topic-http-post-request
func (p *PushNotification) IsTopic() bool {
	return (p.Platform == consts.PlatFormAndroid && p.To != "" && strings.HasPrefix(p.To, "/topics/")) ||
		p.Condition != ""
}

// CheckMessage for check request message
func CheckMessage(req PushNotification) error {
	var msg string

	// ignore send topic mesaage from FCM
	if !req.IsTopic() && len(req.Tokens) == 0 && len(req.To) == 0 {
		msg = "the message must specify at least one registration ID"
		LogAccess.Debug(msg)
		return errors.New(msg)
	}

	if len(req.Tokens) == 1 && len(req.Tokens[0]) == 0 {
		msg = "the token must not be empty"
		LogAccess.Debug(msg)
		return errors.New(msg)
	}

	if req.Platform == consts.PlatFormAndroid && len(req.Tokens) > 1000 {
		msg = "the message may specify at most 1000 registration IDs"
		LogAccess.Debug(msg)
		return errors.New(msg)
	}

	// ref: https://firebase.google.com/docs/cloud-messaging/http-server-ref
	if req.Platform == consts.PlatFormAndroid && req.TimeToLive != nil && (*req.TimeToLive < uint(0) || uint(2419200) < *req.TimeToLive) {
		msg = "the message's TimeToLive field must be an integer " +
			"between 0 and 2419200 (4 weeks)"
		LogAccess.Debug(msg)
		return errors.New(msg)
	}

	return nil
}

// SetProxy only working for FCM server.
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
	if !PushConf.Ios.Enabled && !PushConf.Android.Enabled {
		return errors.New("Please enable iOS or Android config in yml config")
	}

	if PushConf.Ios.Enabled {
		if PushConf.Ios.KeyPath == "" && PushConf.Ios.KeyBase64 == "" {
			return errors.New("Missing iOS certificate key")
		}

		// check certificate file exist
		if PushConf.Ios.KeyPath != "" {
			if _, err := os.Stat(PushConf.Ios.KeyPath); os.IsNotExist(err) {
				return errors.New("certificate file does not exist")
			}
		}
	}

	if PushConf.Android.Enabled {
		if PushConf.Android.APIKey == "" {
			return errors.New("Missing Android API Key")
		}
	}

	return nil
}
