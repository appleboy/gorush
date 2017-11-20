package apns2

import (
	"encoding/json"
	"time"
)

const (
	// PriorityLow will tell APNs to send the push message at a time that takes
	// into account power considerations for the device. Notifications with this
	// priority might be grouped and delivered in bursts. They are throttled, and
	// in some cases are not delivered.
	PriorityLow = 5

	// PriorityHigh will tell APNs to send the push message immediately.
	// Notifications with this priority must trigger an alert, sound, or badge on
	// the target device. It is an error to use this priority for a push
	// notification that contains only the content-available key.
	PriorityHigh = 10
)

// Notification represents the the data and metadata for a APNs Remote Notification.
type Notification struct {

	// An optional canonical UUID that identifies the notification. The canonical
	// form is 32 lowercase hexadecimal digits, displayed in five groups separated
	// by hyphens in the form 8-4-4-4-12. An example UUID is as follows:
	//
	// 	123e4567-e89b-12d3-a456-42665544000
	//
	// If you don't set this, a new UUID is created by APNs and returned in the
	// response.
	ApnsID string

	// A string which allows multiple notifications with the same collapse identifier
	// to be displayed to the user as a single notification. The value should not
	// exceed 64 bytes.
	CollapseID string

	// A string containing hexadecimal bytes of the device token for the target device.
	DeviceToken string

	// The topic of the remote notification, which is typically the bundle ID for
	// your app. The certificate you create in the Apple Developer Member Center
	// must include the capability for this topic. If your certificate includes
	// multiple topics, you must specify a value for this header. If you omit this
	// header and your APNs certificate does not specify multiple topics, the APNs
	// server uses the certificate’s Subject as the default topic.
	Topic string

	// An optional time at which the notification is no longer valid and can be
	// discarded by APNs. If this value is in the past, APNs treats the
	// notification as if it expires immediately and does not store the
	// notification or attempt to redeliver it. If this value is left as the
	// default (ie, Expiration.IsZero()) an expiration header will not added to the
	// http request.
	Expiration time.Time

	// The priority of the notification. Specify ether apns.PriorityHigh (10) or
	// apns.PriorityLow (5) If you don't set this, the APNs server will set the
	// priority to 10.
	Priority int

	// A byte array containing the JSON-encoded payload of this push notification.
	// Refer to "The Remote Notification Payload" section in the Apple Local and
	// Remote Notification Programming Guide for more info.
	Payload interface{}
}

// MarshalJSON converts the notification payload to JSON.
func (n *Notification) MarshalJSON() ([]byte, error) {
	switch n.Payload.(type) {
	case string:
		return []byte(n.Payload.(string)), nil
	case []byte:
		return n.Payload.([]byte), nil
	default:
		return json.Marshal(n.Payload)
	}
}
