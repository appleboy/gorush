package gorush

// Storage interface
type Storage interface {
	addTotalCount(int64)
	addIosSuccess(int64)
	addIosError(int64)
	addAndroidSuccess(int64)
	addAndroidError(int64)
	getTotalCount() int64
	getIosSuccess() int64
	getIosError() int64
	getAndroidSuccess() int64
	getAndroidError() int64
}
