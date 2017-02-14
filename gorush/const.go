package gorush

const (
	// PlatformIos constant is 1 for iOS
	PlatformIos = iota + 1
	// PlatformAndroid constant is 2 for Android
	PlatformAndroid
	// PlatformWeb constant is 3 for Web
	PlatformWeb
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
