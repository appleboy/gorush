package notify

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"

	qcore "github.com/golang-queue/queue/core"
	jsoniter "github.com/json-iterator/go"
	"github.com/msalihkarakasli/go-hms-push/push/model"
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
	APIKey                string           `json:"api_key,omitempty"`
	To                    string           `json:"to,omitempty"`
	CollapseKey           string           `json:"collapse_key,omitempty"`
	TimeToLive            *uint            `json:"time_to_live,omitempty"`
	RestrictedPackageName string           `json:"restricted_package_name,omitempty"`
	DryRun                bool             `json:"dry_run,omitempty"`
	Condition             string           `json:"condition,omitempty"`
	Notification          *FCMNotification `json:"notification,omitempty"`

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

	// ref: https://github.com/sideshow/apns2/blob/54928d6193dfe300b6b88dad72b7e2ae138d4f0a/payload/builder.go#L7-L24
	InterruptionLevel string `json:"interruption_level,omitempty"`
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
	if p.Platform == core.PlatFormAndroid {
		return p.To != "" && strings.HasPrefix(p.To, "/topics/") || p.Condition != ""
	}

	if p.Platform == core.PlatFormHuawei {
		return p.Topic != "" || p.Condition != ""
	}

	return false
}

// FCMNotification specifies the predefined, user-visible key-value pairs of the
// notification payload.
// Copied as is from go-fcm (old FCM API) to keep backward compatibility in external contracts
// Only TitleLockArgs and BodyLocArgs fixed to be an arrays according docs
type FCMNotification struct {
	Title        string   `json:"title,omitempty"`
	Body         string   `json:"body,omitempty"`
	ChannelID    string   `json:"android_channel_id,omitempty"`
	Icon         string   `json:"icon,omitempty"`
	Image        string   `json:"image,omitempty"`
	Sound        string   `json:"sound,omitempty"`
	Badge        string   `json:"badge,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	Color        string   `json:"color,omitempty"`
	ClickAction  string   `json:"click_action,omitempty"`
	BodyLocKey   string   `json:"body_loc_key,omitempty"`
	BodyLocArgs  []string `json:"body_loc_args,omitempty"`
	TitleLocKey  string   `json:"title_loc_key,omitempty"`
	TitleLocArgs []string `json:"title_loc_args,omitempty"`
}

func (f FCMNotification) NotificationCount() (*int, error) {
	if f.Badge == "" {
		return nil, nil
	}

	v, err := strconv.Atoi(f.Badge)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

// CheckMessage for check request message
func CheckMessage(req *PushNotification) error {
	var msg string

	if req.Platform == core.PlatFormAndroid && req.IsTopic() {
		msg = "android topics not supported yet"
		logx.LogAccess.Debug(msg)
		return errors.New(msg)
	}

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

	if req.Platform == core.PlatFormAndroid && len(req.Tokens) > 500 {
		msg = "the message may specify at most 500 registration IDs"
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
		if cfg.Android.ServiceAccountKey == "" {
			return errors.New("missing service account key")
		}

		if cfg.Android.ProjectID == "" {
			return errors.New("missing project id")
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
		resp, err = PushToIOS(v, cfg)
	case core.PlatFormAndroid:
		resp, err = PushToAndroidV1(ctx, v, cfg)
	case core.PlatFormHuawei:
		resp, err = PushToHuawei(v, cfg)
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
