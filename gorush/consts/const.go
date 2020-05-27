package consts

import "github.com/appleboy/gorush/gorush/structs"

const (
	// PlatFormIos constant is 1 for iOS
	PlatFormIos = structs.Platform(iota + 1)
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
	ApnsPriorityLow = 5

	// ApnsPriorityHigh will tell APNs to send the push message immediately.
	// Notifications with this priority must trigger an alert, sound, or badge on
	// the target device. It is an error to use this priority for a push
	// notification that contains only the content-available key.
	ApnsPriorityHigh = 10
)

const (
	// PriorityNormal sets push notification with normal priority
	PriorityNormal = structs.Priority("normal")
	// PriorityHigh sets push notification with high priority
	PriorityHigh = structs.Priority("high")
)

const (
	// PushTypeAlert for notifications that trigger a user interaction—for example, an alert, badge, or sound.
	PushTypeAlert = structs.APNSPushType("alert")
	// PushTypeBackground for notifications that deliver content in the background, and don’t trigger any user interactions
	PushTypeBackground = structs.APNSPushType("background")
	// PushTypeVOIP for notifications that provide information about an incoming Voice-over-IP (VoIP) call
	PushTypeVOIP = structs.APNSPushType("voip")
	// PushTypeComplication for notifications that contain update information for a watchOS app’s complications
	PushTypeComplication = structs.APNSPushType("complication")
	// PushTypeFileProvider to signal changes to a File Provider extension
	PushTypeFileProvider = structs.APNSPushType("fileprovider")
	// PushTypeMDM for notifications that tell managed devices to contact the MDM server
	PushTypeMDM = structs.APNSPushType("mdm")
)
