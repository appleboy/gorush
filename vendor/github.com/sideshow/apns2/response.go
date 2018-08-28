package apns2

import (
	"net/http"
	"strconv"
	"time"
)

// StatusSent is a 200 response.
const StatusSent = http.StatusOK

// The possible Reason error codes returned from APNs.
// From table 8-6 in the Apple Local and Remote Notification Programming Guide.
const (
	// 400 The collapse identifier exceeds the maximum allowed size
	ReasonBadCollapseID = "BadCollapseId"

	// 400 The specified device token was bad. Verify that the request contains a
	// valid token and that the token matches the environment.
	ReasonBadDeviceToken = "BadDeviceToken"

	// 400 The apns-expiration value is bad.
	ReasonBadExpirationDate = "BadExpirationDate"

	// 400 The apns-id value is bad.
	ReasonBadMessageID = "BadMessageId"

	// 400 The apns-priority value is bad.
	ReasonBadPriority = "BadPriority"

	// 400 The apns-topic was invalid.
	ReasonBadTopic = "BadTopic"

	// 400 The device token does not match the specified topic.
	ReasonDeviceTokenNotForTopic = "DeviceTokenNotForTopic"

	// 400 One or more headers were repeated.
	ReasonDuplicateHeaders = "DuplicateHeaders"

	// 400 Idle time out.
	ReasonIdleTimeout = "IdleTimeout"

	// 400 The device token is not specified in the request :path. Verify that the
	// :path header contains the device token.
	ReasonMissingDeviceToken = "MissingDeviceToken"

	// 400 The apns-topic header of the request was not specified and was
	// required. The apns-topic header is mandatory when the client is connected
	// using a certificate that supports multiple topics.
	ReasonMissingTopic = "MissingTopic"

	// 400 The message payload was empty.
	ReasonPayloadEmpty = "PayloadEmpty"

	// 400 Pushing to this topic is not allowed.
	ReasonTopicDisallowed = "TopicDisallowed"

	// 403 The certificate was bad.
	ReasonBadCertificate = "BadCertificate"

	// 403 The client certificate was for the wrong environment.
	ReasonBadCertificateEnvironment = "BadCertificateEnvironment"

	// 403 The provider token is stale and a new token should be generated.
	ReasonExpiredProviderToken = "ExpiredProviderToken"

	// 403 The specified action is not allowed.
	ReasonForbidden = "Forbidden"

	// 403 The provider token is not valid or the token signature could not be
	// verified.
	ReasonInvalidProviderToken = "InvalidProviderToken"

	// 403 No provider certificate was used to connect to APNs and Authorization
	// header was missing or no provider token was specified.
	ReasonMissingProviderToken = "MissingProviderToken"

	// 404 The request contained a bad :path value.
	ReasonBadPath = "BadPath"

	// 405 The specified :method was not POST.
	ReasonMethodNotAllowed = "MethodNotAllowed"

	// 410 The device token is inactive for the specified topic.
	ReasonUnregistered = "Unregistered"

	// 413 The message payload was too large. See Creating the Remote Notification
	// Payload in the Apple Local and Remote Notification Programming Guide for
	// details on maximum payload size.
	ReasonPayloadTooLarge = "PayloadTooLarge"

	// 429 The provider token is being updated too often.
	ReasonTooManyProviderTokenUpdates = "TooManyProviderTokenUpdates"

	// 429 Too many requests were made consecutively to the same device token.
	ReasonTooManyRequests = "TooManyRequests"

	// 500 An internal server error occurred.
	ReasonInternalServerError = "InternalServerError"

	// 503 The service is unavailable.
	ReasonServiceUnavailable = "ServiceUnavailable"

	// 503 The server is shutting down.
	ReasonShutdown = "Shutdown"
)

// Response represents a result from the APNs gateway indicating whether a
// notification was accepted or rejected and (if applicable) the metadata
// surrounding the rejection.
type Response struct {

	// The HTTP status code returned by APNs.
	// A 200 value indicates that the notification was successfully sent.
	// For a list of other possible status codes, see table 6-4 in the Apple Local
	// and Remote Notification Programming Guide.
	StatusCode int

	// The APNs error string indicating the reason for the notification failure (if
	// any). The error code is specified as a string. For a list of possible
	// values, see the Reason constants above.
	// If the notification was accepted, this value will be "".
	Reason string

	// The APNs ApnsID value from the Notification. If you didn't set an ApnsID on the
	// Notification, this will be a new unique UUID which has been created by APNs.
	ApnsID string

	// If the value of StatusCode is 410, this is the last time at which APNs
	// confirmed that the device token was no longer valid for the topic.
	Timestamp Time
}

// Sent returns whether or not the notification was successfully sent.
// This is the same as checking if the StatusCode == 200.
func (c *Response) Sent() bool {
	return c.StatusCode == StatusSent
}

// Time represents a device uninstall time
type Time struct {
	time.Time
}

// UnmarshalJSON converts an epoch date into a Time struct.
func (t *Time) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(ts/1000, 0)
	return nil
}
