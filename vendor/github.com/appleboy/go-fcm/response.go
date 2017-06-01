package fcm

import (
	"encoding/json"
	"errors"
)

var (
	// ErrMissingRegistration occurs if registration token is not set.
	ErrMissingRegistration = errors.New("missing registration token")

	// ErrInvalidRegistration occurs if registration token is invalid.
	ErrInvalidRegistration = errors.New("invalid registration token")

	// ErrNotRegistered occurs when application was deleted from device and
	// token is not registered in FCM.
	ErrNotRegistered = errors.New("unregistered device")

	// ErrInvalidPackageName occurs if package name in message is invalid.
	ErrInvalidPackageName = errors.New("invalid package name")

	// ErrMismatchSenderID occurs when application has a new registration token.
	ErrMismatchSenderID = errors.New("mismatched sender id")

	// ErrMessageTooBig occurs when message is too big.
	ErrMessageTooBig = errors.New("message is too big")

	// ErrInvalidDataKey occurs if data key is invalid.
	ErrInvalidDataKey = errors.New("invalid data key")

	// ErrInvalidTTL occurs when message has invalid TTL.
	ErrInvalidTTL = errors.New("invalid time to live")

	// ErrUnavailable occurs when FCM service is unavailable. It makes sense
	// to retry after this error.
	ErrUnavailable = connectionError("timeout")

	// ErrInternalServerError is internal FCM error. It makes sense to retry
	// after this error.
	ErrInternalServerError = serverError("internal server error")

	// ErrDeviceMessageRateExceeded occurs when client sent to many requests to
	// the device.
	ErrDeviceMessageRateExceeded = errors.New("device message rate exceeded")

	// ErrTopicsMessageRateExceeded occurs when client sent to many requests to
	// the topics.
	ErrTopicsMessageRateExceeded = errors.New("topics message rate exceeded")
)

var (
	errMap = map[string]error{
		"MissingRegistration":       ErrMissingRegistration,
		"InvalidRegistration":       ErrInvalidRegistration,
		"NotRegistered":             ErrNotRegistered,
		"InvalidPackageName":        ErrInvalidPackageName,
		"MismatchSenderId":          ErrMismatchSenderID,
		"MessageTooBig":             ErrMessageTooBig,
		"InvalidDataKey":            ErrInvalidDataKey,
		"InvalidTtl":                ErrInvalidTTL,
		"Unavailable":               ErrUnavailable,
		"InternalServerError":       ErrInternalServerError,
		"DeviceMessageRateExceeded": ErrDeviceMessageRateExceeded,
		"TopicsMessageRateExceeded": ErrTopicsMessageRateExceeded,
	}
)

// connectionError represents connection errors such as timeout error, etc.
// Implements `net.Error` interface.
type connectionError string

func (err connectionError) Error() string {
	return string(err)
}

func (err connectionError) Temporary() bool {
	return true
}

func (err connectionError) Timeout() bool {
	return true
}

// serverError represents internal server errors.
// Implements `net.Error` interface.
type serverError string

func (err serverError) Error() string {
	return string(err)
}

func (serverError) Temporary() bool {
	return true
}

func (serverError) Timeout() bool {
	return false
}

// Response represents the FCM server's response to the application
// server's sent message.
type Response struct {
	MulticastID  int64    `json:"multicast_id"`
	Success      int      `json:"success"`
	Failure      int      `json:"failure"`
	CanonicalIDs int      `json:"canonical_ids"`
	Results      []Result `json:"results"`
}

// Result represents the status of a processed message.
type Result struct {
	MessageID      string `json:"message_id"`
	RegistrationID string `json:"registration_id"`
	Error          error  `json:"error"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (r *Result) UnmarshalJSON(data []byte) error {
	var result struct {
		MessageID      string `json:"message_id"`
		RegistrationID string `json:"registration_id"`
		Error          string `json:"error"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	r.MessageID = result.MessageID
	r.RegistrationID = result.RegistrationID
	r.Error = errMap[result.Error]

	return nil
}

// Unregistered checks if the device token is unregistered,
// according to response from FCM server. Useful to determine
// if app is uninstalled.
func (r Result) Unregistered() bool {
	switch r.Error {
	case ErrNotRegistered, ErrMismatchSenderID, ErrMissingRegistration, ErrInvalidRegistration:
		return true

	default:
		return false
	}
}
