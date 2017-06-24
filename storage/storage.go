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
	GetTotalCount() int64
	GetIosSuccess() int64
	GetIosError() int64
	GetAndroidSuccess() int64
	GetAndroidError() int64
}
