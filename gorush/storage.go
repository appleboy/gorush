package gorush

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
