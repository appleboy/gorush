package structs

import (
	"github.com/appleboy/go-fcm"
)

// D provide string array
type D map[string]interface{}

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

// PushNotification is single notification request
type PushNotification struct {
	// Common
	ID               string      `json:"notif_id,omitempty"`
	Tokens           []string    `json:"tokens" binding:"required"`
	Platform         Platform    `json:"platform" binding:"required"`
	Message          string      `json:"message,omitempty"`
	Title            string      `json:"title,omitempty"`
	Image            string      `json:"image,omitempty"`
	Priority         Priority    `json:"priority,omitempty"`
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

	// iOS
	Expiration  *int64       `json:"expiration,omitempty"`
	ApnsID      string       `json:"apns_id,omitempty"`
	CollapseID  string       `json:"collapse_id,omitempty"`
	Topic       string       `json:"topic,omitempty"`
	PushType    APNSPushType `json:"push_type,omitempty"`
	Badge       *int         `json:"badge,omitempty"`
	Category    string       `json:"category,omitempty"`
	ThreadID    string       `json:"thread-id,omitempty"`
	URLArgs     []string     `json:"url-args,omitempty"`
	Alert       Alert        `json:"alert,omitempty"`
	Production  bool         `json:"production,omitempty"`
	Development bool         `json:"development,omitempty"`
	SoundName   string       `json:"name,omitempty"`
	SoundVolume float32      `json:"volume,omitempty"`
	Apns        D            `json:"apns,omitempty"`
}

type (
	// Platform push notification receiver's platform (iOS/Android)
	Platform int
	// Priority push notification's priority
	Priority string

	// APNSPushType is an iOS13 required field describing the type of push notification
	APNSPushType string
)
