package gorush

const (
	// Version is gorush server version.
	Version = "1.2.1"
)

const (
	// PlatFormIos constant is 1 for iOS
	PlatFormIos = iota + 1
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
	gorushTotalCount     = "gorush-total-count"
	gorushIosSuccess     = "gorush-ios-success-count"
	gorushIosError       = "gorush-ios-error-count"
	gorushAndroidSuccess = "gorush-android-success-count"
	gorushAndroidError   = "gorush-android-error-count"
)
