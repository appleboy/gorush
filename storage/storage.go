package storage

const (
	// TotalCountKey is key name for total count of storage
	TotalCountKey = "gorush-total-count"

	// IosSuccessKey is key name or ios success count of storage
	IosSuccessKey = "gorush-ios-success-count"

	// IosErrorKey is key name or ios success error of storage
	IosErrorKey = "gorush-ios-error-count"

	// AndroidSuccessKey is key name for android success count of storage
	AndroidSuccessKey = "gorush-android-success-count"

	// AndroidErrorKey is key name for android error count of storage
	AndroidErrorKey = "gorush-android-error-count"

	// WebSuccessKey is key name for web success count of storage
	WebSuccessKey = "gorush-web-success-count"

	// WebErrorKey is key name for web error count of storage
	WebErrorKey = "gorush-web-error-count"
)

// Storage interface
type Storage interface {
	Init() error
	Reset()
	AddTotalCount(int64)
	AddIosSuccess(int64)
	AddIosError(int64)
	AddAndroidSuccess(int64)
	AddAndroidError(int64)
	AddWebSuccess(int64)
	AddWebError(int64)
	GetTotalCount() int64
	GetIosSuccess() int64
	GetIosError() int64
	GetAndroidSuccess() int64
	GetAndroidError() int64
	GetWebSuccess() int64
	GetWebError() int64
}
