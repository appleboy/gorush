package core

const (
	// TotalCountKey is key name for total count of storage
	TotalCountKey = "gorush-total-count"

	// IosSuccessKey is key name or ios success count of storage
	/* #nosec */
	IosSuccessKey = "gorush-ios-success-count"

	// IosErrorKey is key name or ios success error of storage
	IosErrorKey = "gorush-ios-error-count"

	// AndroidSuccessKey is key name for android success count of storage
	AndroidSuccessKey = "gorush-android-success-count"

	// AndroidErrorKey is key name for android error count of storage
	AndroidErrorKey = "gorush-android-error-count"

	// HuaweiSuccessKey is key name for huawei success count of storage
	HuaweiSuccessKey = "gorush-huawei-success-count"

	// HuaweiErrorKey is key name for huawei error count of storage
	HuaweiErrorKey = "gorush-huawei-error-count"
)

// Storage interface
type Storage interface {
	Init() error
	Add(key string, count int64)
	Set(key string, count int64)
	Get(key string) int64
	Close() error
}
