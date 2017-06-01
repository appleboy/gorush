package fcm

import (
	"errors"
	"strings"
)

var (
	// ErrInvalidMessage occurs if push notitication message is nil.
	ErrInvalidMessage = errors.New("message is invalid")

	// ErrInvalidTarget occurs if message topic is empty.
	ErrInvalidTarget = errors.New("topic is invalid or registration ids are not set")

	// ErrToManyRegIDs occurs when registration ids more then 1000.
	ErrToManyRegIDs = errors.New("too many registrations ids")

	// ErrInvalidTimeToLive occurs if TimeToLive more then 2419200.
	ErrInvalidTimeToLive = errors.New("messages time-to-live is invalid")
)

// Notification specifies the predefined, user-visible key-value pairs of the
// notification payload.
type Notification struct {
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	Icon         string `json:"icon,omitempty"`
	Sound        string `json:"sound,omitempty"`
	Badge        string `json:"badge,omitempty"`
	Tag          string `json:"tag,omitempty"`
	Color        string `json:"color,omitempty"`
	ClickAction  string `json:"click_action,omitempty"`
	BodyLocKey   string `json:"body_loc_key,omitempty"`
	BodyLocArgs  string `json:"body_loc_args,omitempty"`
	TitleLocKey  string `json:"title_loc_key,omitempty"`
	TitleLocArgs string `json:"title_loc_args,omitempty"`
}

// Message represents list of targets, options, and payload for HTTP JSON
// messages.
type Message struct {
	To                       string                 `json:"to,omitempty"`
	RegistrationIDs          []string               `json:"registration_ids,omitempty"`
	Condition                string                 `json:"condition,omitempty"`
	CollapseKey              string                 `json:"collapse_key,omitempty"`
	Priority                 string                 `json:"priority,omitempty"`
	ContentAvailable         bool                   `json:"content_available,omitempty"`
	DelayWhileIdle           bool                   `json:"delay_while_idle,omitempty"`
	TimeToLive               *uint                  `json:"time_to_live,omitempty"`
	DeliveryReceiptRequested bool                   `json:"delivery_receipt_requested,omitempty"`
	DryRun                   bool                   `json:"dry_run,omitempty"`
	RestrictedPackageName    string                 `json:"restricted_package_name,omitempty"`
	Notification             *Notification          `json:"notification,omitempty"`
	Data                     map[string]interface{} `json:"data,omitempty"`
}

// Validate returns an error if the message is not well-formed.
func (msg *Message) Validate() error {
	if msg == nil {
		return ErrInvalidMessage
	}

	// validate target identifier: `to` or `condition`, or `registration_ids`
	opCnt := strings.Count(msg.Condition, "&&") + strings.Count(msg.Condition, "||")
	if msg.To == "" && (msg.Condition == "" || opCnt > 2) && len(msg.RegistrationIDs) == 0 {
		return ErrInvalidTarget
	}

	if len(msg.RegistrationIDs) > 1000 {
		return ErrToManyRegIDs
	}

	if msg.TimeToLive != nil && *msg.TimeToLive > uint(2419200) {
		return ErrInvalidTimeToLive
	}
	return nil
}
