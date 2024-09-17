package notify

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"

	"firebase.google.com/go/v4/messaging"
	"github.com/appleboy/go-hms-push/push/model"
	qcore "github.com/golang-queue/queue/core"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

// ResponsePush response of notification request.
type ResponsePush struct {
	Logs []logx.LogPushEntry `json:"logs"`
}

// PushNotification is single notification request
type PushNotification struct {
	// Common
	ID               string      `json:"notif_id,omitempty"`
	To               string      `json:"to,omitempty"`
	Topic            string      `json:"topic,omitempty"` // FCM and iOS only
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
	Notification *messaging.Notification  `json:"notification,omitempty"`
	Android      *messaging.AndroidConfig `json:"android,omitempty"`
	Webpush      *messaging.WebpushConfig `json:"webpush,omitempty"`
	APNS         *messaging.APNSConfig    `json:"apns,omitempty"`
	FCMOptions   *messaging.FCMOptions    `json:"fcm_options,omitempty"`
	Condition    string                   `json:"condition,omitempty"`

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

	// ref: https://github.com/sideshow/apns2/blob/54928d6193dfe300b6b88dad72b7e2ae138d4f0a/payload/builder.go#L7-L24
	InterruptionLevel string `json:"interruption_level,omitempty"`

	// live-activity support
	// ref: https://apple.co/3MLe2DB
	ContentState  D      `json:"content-state,omitempty"`
	StaleDate     int64  `json:"stale-date,omitempty"`
	DismissalDate int64  `json:"dismissal-date"`
	Event         string `json:"event,omitempty"`
	Timestamp     int64  `json:"timestamp,omitempty"`
}

// Bytes for queue message
func (p *PushNotification) Bytes() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return b
}

// IsTopic check if message format is topic for FCM
// ref: https://firebase.google.com/docs/cloud-messaging/send-message#topic-http-post-request
func (p *PushNotification) IsTopic() bool {
	if p.Platform == core.PlatFormHuawei || p.Platform == core.PlatFormAndroid {
		return p.Topic != "" || p.Condition != ""
	}

	return false
}

// CheckMessage for check request message
func CheckMessage(req *PushNotification) error {
	var msg string

	if req.To != "" {
		req.Tokens = append(req.Tokens, req.To)
	}

	// if the message is a topic, the tokens field is not required
	if !req.IsTopic() && len(req.Tokens) == 0 {
		return errors.New("tokens must not be nil or empty")
	}

	switch req.Platform {
	case core.PlatFormIos:
		if len(req.Tokens) == 1 && req.Tokens[0] == "" {
			msg = "the token must not be empty"
			logx.LogAccess.Debug(msg)
			return errors.New(msg)
		}
	case
		core.PlatFormAndroid,
		core.PlatFormHuawei:
		if len(req.Tokens) > 500 {
			msg = "tokens must not contain more than 500 elements"
			logx.LogAccess.Debug(msg)
			return errors.New(msg)
		}
	default:
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
func CheckPushConf(cfg *config.ConfYaml) error {
	if !cfg.Ios.Enabled && !cfg.Android.Enabled && !cfg.Huawei.Enabled {
		return errors.New("please enable iOS, Android or Huawei config in yml config")
	}

	if cfg.Ios.Enabled {
		if cfg.Ios.KeyPath == "" && cfg.Ios.KeyBase64 == "" {
			return errors.New("missing iOS certificate key")
		}

		// check certificate file exist
		if cfg.Ios.KeyPath != "" {
			if _, err := os.Stat(cfg.Ios.KeyPath); os.IsNotExist(err) {
				return errors.New("certificate file does not exist")
			}
		}
	}

	if cfg.Android.Enabled {
		credential := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if cfg.Android.Credential == "" &&
			cfg.Android.KeyPath == "" &&
			credential == "" {
			return errors.New("missing fcm credential data")
		}
	}

	if cfg.Huawei.Enabled {
		if cfg.Huawei.AppSecret == "" {
			return errors.New("missing huawei app secret")
		}

		if cfg.Huawei.AppID == "" {
			return errors.New("missing huawei app id")
		}
	}

	return nil
}

// SendNotification provide send notification.
func SendNotification(
	ctx context.Context,
	req qcore.QueuedMessage,
	cfg *config.ConfYaml,
) (resp *ResponsePush, err error) {
	v, ok := req.(*PushNotification)
	if !ok {
		if err = json.Unmarshal(req.Bytes(), &v); err != nil {
			return nil, err
		}
	}

	switch v.Platform {
	case core.PlatFormIos:
		resp, err = PushToIOS(ctx, v, cfg)
	case core.PlatFormAndroid:
		resp, err = PushToAndroid(ctx, v, cfg)
	case core.PlatFormHuawei:
		resp, err = PushToHuawei(ctx, v, cfg)
	}

	if cfg.Core.FeedbackURL != "" {
		for _, l := range resp.Logs {
			err := DispatchFeedback(ctx, l, cfg.Core.FeedbackURL, cfg.Core.FeedbackTimeout, cfg.Core.FeedbackHeader)
			if err != nil {
				logx.LogError.Error(err)
			}
		}
	}

	return resp, err
}

// Run send notification
var Run = func(cfg *config.ConfYaml) func(ctx context.Context, msg qcore.QueuedMessage) error {
	return func(ctx context.Context, msg qcore.QueuedMessage) error {
		_, err := SendNotification(ctx, msg, cfg)
		return err
	}
}
