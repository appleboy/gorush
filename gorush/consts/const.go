package consts

import (
	"github.com/appleboy/gorush/gorush/structs"
	"github.com/sideshow/apns2"
)

const (
	// PlatFormIos constant is 1 for iOS
	PlatFormIos structs.Platform = iota + 1
	// PlatFormAndroid constant is 2 for Android
	PlatFormAndroid
)

const (
	// SucceededPush is log block
	SucceededPush = "succeeded-push"
	// FailedPush is log block
	FailedPush = "failed-push"
)

// Stat variable for redis
const (
	TotalCountKey     = "gorush-total-count"
	IosSuccessKey     = "gorush-ios-success-count"
	IosErrorKey       = "gorush-ios-error-count"
	AndroidSuccessKey = "gorush-android-success-count"
	AndroidErrorKey   = "gorush-android-error-count"
)

const (
	// ApnsPriorityLow will tell APNs to send the push message at a time that takes
	// into account power considerations for the device. Notifications with this
	// priority might be grouped and delivered in bursts. They are throttled, and
	// in some cases are not delivered.
	ApnsPriorityLow = apns2.PriorityLow

	// ApnsPriorityHigh will tell APNs to send the push message immediately.
	// Notifications with this priority must trigger an alert, sound, or badge on
	// the target device. It is an error to use this priority for a push
	// notification that contains only the content-available key.
	ApnsPriorityHigh = apns2.PriorityHigh
)

const (
	// PriorityNormal sets push notification with normal priority
	PriorityNormal structs.Priority = "normal"
	// PriorityHigh sets push notification with high priority
	PriorityHigh structs.Priority = "high"
)

const (
	// PushTypeAlert is used for notifications that trigger a user interaction —
	// for example, an alert, badge, or sound. If you set this push type, the
	// apns-topic header field must use your app’s bundle ID as the topic. The
	// alert push type is required on watchOS 6 and later. It is recommended on
	// macOS, iOS, tvOS, and iPadOS.
	PushTypeAlert = apns2.PushTypeAlert

	// PushTypeBackground is used for notifications that deliver content in the
	// background, and don’t trigger any user interactions. If you set this push
	// type, the apns-topic header field must use your app’s bundle ID as the
	// topic. The background push type is required on watchOS 6 and later. It is
	// recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeBackground = apns2.PushTypeBackground

	// PushTypeVOIP is used for notifications that provide information about an
	// incoming Voice-over-IP (VoIP) call. If you set this push type, the
	// apns-topic header field must use your app’s bundle ID with .voip appended
	// to the end. If you’re using certificate-based authentication, you must
	// also register the certificate for VoIP services. The voip push type is
	// not available on watchOS. It is recommended on macOS, iOS, tvOS, and
	// iPadOS.
	PushTypeVOIP = apns2.PushTypeVOIP

	// PushTypeComplication is used for notifications that contain update
	// information for a watchOS app’s complications. If you set this push type,
	// the apns-topic header field must use your app’s bundle ID with
	// .complication appended to the end. If you’re using certificate-based
	// authentication, you must also register the certificate for WatchKit
	// services. The complication push type is recommended for watchOS and iOS.
	// It is not available on macOS, tvOS, and iPadOS.
	PushTypeComplication = apns2.PushTypeComplication

	// PushTypeFileProvider is used to signal changes to a File Provider
	// extension. If you set this push type, the apns-topic header field must
	// use your app’s bundle ID with .pushkit.fileprovider appended to the end.
	// The fileprovider push type is not available on watchOS. It is recommended
	// on macOS, iOS, tvOS, and iPadOS.
	PushTypeFileProvider = apns2.PushTypeFileProvider

	// PushTypeMDM is used for notifications that tell managed devices to
	// contact the MDM server. If you set this push type, you must use the topic
	// from the UID attribute in the subject of your MDM push certificate.
	PushTypeMDM = apns2.PushTypeMDM
)
