package gorush

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/appleboy/go-fcm"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/msalihkarakasli/go-hms-push/push/model"
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
	Action          string   `json:"action,omitempty"`
	ActionLocKey    string   `json:"action-loc-key,omitempty"`
	Body            string   `json:"body,omitempty"`
	LaunchImage     string   `json:"launch-image,omitempty"`
	LocArgs         []string `json:"loc-args,omitempty"`
	LocKey          string   `json:"loc-key,omitempty"`
	Title           string   `json:"title,omitempty"`
	Subtitle        string   `json:"subtitle,omitempty"`
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`
	TitleLocKey     string   `json:"title-loc-key,omitempty"`
	SummaryArg      string   `json:"summary-arg,omitempty"`
	SummaryArgCount int      `json:"summary-arg-count,omitempty"`
}

// RequestPush support multiple notification request.
type RequestPush struct {
	Notifications []PushNotification `json:"notifications" binding:"required"`
}

// PushNotification is single notification request
type PushNotification struct {
	Wg  *sync.WaitGroup
	Log *[]logx.LogPushEntry

	// Common
	ID               string      `json:"notif_id,omitempty"`
	Tokens           []string    `json:"tokens" binding:"required"`
	Platform         int         `json:"platform" binding:"required"`
	Message          string      `json:"message,omitempty"`
	Title            string      `json:"title,omitempty"`
	Image            string      `json:"image,omitempty"`
	Priority         string      `json:"priority,omitempty"`
	ContentAvailable bool        `json:"content_available,omitempty"`
	MutableContent   bool        `json:"mutable_content,omitempty"`
	Sound            interface{} `json:"sound,omitempty"`
	Data             D           `json:"data,omitempty"`
	Retry            int         `json:"retry,omitempty"`

	// Android
	APIKey                string            `json:"api_key,omitempty"`
	To                    string            `json:"to,omitempty"`
	CollapseKey           string            `json:"collapse_key,omitempty"`
	DelayWhileIdle        bool              `json:"delay_while_idle,omitempty"`
	TimeToLive            *uint             `json:"time_to_live,omitempty"`
	RestrictedPackageName string            `json:"restricted_package_name,omitempty"`
	DryRun                bool              `json:"dry_run,omitempty"`
	Condition             string            `json:"condition,omitempty"`
	Notification          *fcm.Notification `json:"notification,omitempty"`

	// Huawei
	AppID              string                     `json:"app_id,omitempty"`
	AppSecret          string                     `json:"app_secret,omitempty"`
	HuaweiNotification *model.AndroidNotification `json:"huawei_notification,omitempty"`
	HuaweiData         string                     `json:"huawei_data,omitempty"`
	HuaweiCollapseKey  int                        `json:"huawei_collapse_key,omitempty"`
	HuaweiTTL          string                     `json:"huawei_ttl,omitempty"`
	BiTag              string                     `json:"bi_tag,omitempty"`
	FastAppTarget      int                        `json:"fast_app_target,omitempty"`

	// iOS
	Expiration  *int64   `json:"expiration,omitempty"`
	ApnsID      string   `json:"apns_id,omitempty"`
	CollapseID  string   `json:"collapse_id,omitempty"`
	Topic       string   `json:"topic,omitempty"`
	PushType    string   `json:"push_type,omitempty"`
	Badge       *int     `json:"badge,omitempty"`
	Category    string   `json:"category,omitempty"`
	ThreadID    string   `json:"thread-id,omitempty"`
	URLArgs     []string `json:"url-args,omitempty"`
	Alert       Alert    `json:"alert,omitempty"`
	Production  bool     `json:"production,omitempty"`
	Development bool     `json:"development,omitempty"`
	SoundName   string   `json:"name,omitempty"`
	SoundVolume float32  `json:"volume,omitempty"`
	Apns        D        `json:"apns,omitempty"`
}

// WaitDone decrements the WaitGroup counter.
func (p *PushNotification) WaitDone() {
	if p.Wg != nil {
		p.Wg.Done()
	}
}

// AddWaitCount increments the WaitGroup counter.
func (p *PushNotification) AddWaitCount() {
	if p.Wg != nil {
		p.Wg.Add(1)
	}
}

// AddLog record fail log of notification
func (p *PushNotification) AddLog(log logx.LogPushEntry) {
	if p.Log != nil {
		*p.Log = append(*p.Log, log)
	}
}

// IsTopic check if message format is topic for FCM
// ref: https://firebase.google.com/docs/cloud-messaging/send-message#topic-http-post-request
func (p *PushNotification) IsTopic() bool {
	if p.Platform == core.PlatFormAndroid {
		return p.To != "" && strings.HasPrefix(p.To, "/topics/") || p.Condition != ""
	}

	if p.Platform == core.PlatFormHuawei {
		return p.Topic != "" || p.Condition != ""
	}

	return false
}

// CheckMessage for check request message
func CheckMessage(req PushNotification) error {
	var msg string

	// ignore send topic mesaage from FCM
	if !req.IsTopic() && len(req.Tokens) == 0 && req.To == "" {
		msg = "the message must specify at least one registration ID"
		logx.LogAccess.Debug(msg)
		return errors.New(msg)
	}

	if len(req.Tokens) == core.PlatFormIos && req.Tokens[0] == "" {
		msg = "the token must not be empty"
		logx.LogAccess.Debug(msg)
		return errors.New(msg)
	}

	if req.Platform == core.PlatFormAndroid && len(req.Tokens) > 1000 {
		msg = "the message may specify at most 1000 registration IDs"
		logx.LogAccess.Debug(msg)
		return errors.New(msg)
	}

	if req.Platform == core.PlatFormHuawei && len(req.Tokens) > 500 {
		msg = "the message may specify at most 500 registration IDs for Huawei"
		logx.LogAccess.Debug(msg)
		return errors.New(msg)
	}

	// ref: https://firebase.google.com/docs/cloud-messaging/http-server-ref
	if req.Platform == core.PlatFormAndroid && req.TimeToLive != nil && *req.TimeToLive > uint(2419200) {
		msg = "the message's TimeToLive field must be an integer " +
			"between 0 and 2419200 (4 weeks)"
		logx.LogAccess.Debug(msg)
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
	logx.LogAccess.Debug("Set http proxy as " + proxy)

	return nil
}

// CheckPushConf provide check your yml config.
func CheckPushConf(cfg config.ConfYaml) error {
	if !cfg.Ios.Enabled && !cfg.Android.Enabled && !cfg.Huawei.Enabled {
		return errors.New("Please enable iOS, Android or Huawei config in yml config")
	}

	if cfg.Ios.Enabled {
		if cfg.Ios.KeyPath == "" && cfg.Ios.KeyBase64 == "" {
			return errors.New("Missing iOS certificate key")
		}

		// check certificate file exist
		if cfg.Ios.KeyPath != "" {
			if _, err := os.Stat(cfg.Ios.KeyPath); os.IsNotExist(err) {
				return errors.New("certificate file does not exist")
			}
		}
	}

	if cfg.Android.Enabled {
		if cfg.Android.APIKey == "" {
			return errors.New("Missing Android API Key")
		}
	}

	if cfg.Huawei.Enabled {
		if cfg.Huawei.AppSecret == "" {
			return errors.New("Missing Huawei App Secret")
		}

		if cfg.Huawei.AppID == "" {
			return errors.New("Missing Huawei App ID")
		}
	}

	return nil
}

// SendNotification send notification
func SendNotification(cfg config.ConfYaml, req PushNotification) {
	defer func() {
		req.WaitDone()
	}()

	switch req.Platform {
	case core.PlatFormIos:
		PushToIOS(cfg, req)
	case core.PlatFormAndroid:
		PushToAndroid(cfg, req)
	case core.PlatFormHuawei:
		PushToHuawei(cfg, req)
	}
}
