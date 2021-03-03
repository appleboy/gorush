package gorush

const (
	// PlatFormIos constant is 1 for iOS
	PlatFormIos = iota + 1
	// PlatFormAndroid constant is 2 for Android
	PlatFormAndroid
	// PlatFormHuawei constant is 3 for Huawei
	PlatFormHuawei
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
	HuaweiSuccessKey  = "gorush-huawei-success-count"
	HuaweiErrorKey    = "gorush-huawei-error-count"
)
