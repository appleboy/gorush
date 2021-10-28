package memory

import (
	"sync/atomic"
)

// statApp is app status structure
type statApp struct {
	TotalCount int64         `json:"total_count"`
	Ios        IosStatus     `json:"ios"`
	Android    AndroidStatus `json:"android"`
	Huawei     HuaweiStatus  `json:"huawei"`
}

// AndroidStatus is android structure
type AndroidStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// IosStatus is iOS structure
type IosStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// HuaweiStatus is android structure
type HuaweiStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New() *Storage {
	return &Storage{
		stat: &statApp{},
	}
}

// Storage is interface structure
type Storage struct {
	stat *statApp
}

// Init client storage.
func (s *Storage) Init() error {
	return nil
}

// Close the storage connection
func (s *Storage) Close() error {
	return nil
}

// Reset Client storage.
func (s *Storage) Reset() {
	atomic.StoreInt64(&s.stat.TotalCount, 0)
	atomic.StoreInt64(&s.stat.Ios.PushSuccess, 0)
	atomic.StoreInt64(&s.stat.Ios.PushError, 0)
	atomic.StoreInt64(&s.stat.Android.PushSuccess, 0)
	atomic.StoreInt64(&s.stat.Android.PushError, 0)
	atomic.StoreInt64(&s.stat.Huawei.PushSuccess, 0)
	atomic.StoreInt64(&s.stat.Huawei.PushError, 0)
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	atomic.AddInt64(&s.stat.TotalCount, count)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	atomic.AddInt64(&s.stat.Ios.PushSuccess, count)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	atomic.AddInt64(&s.stat.Ios.PushError, count)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	atomic.AddInt64(&s.stat.Android.PushSuccess, count)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	atomic.AddInt64(&s.stat.Android.PushError, count)
}

// AddHuaweiSuccess record counts of success Huawei push notification.
func (s *Storage) AddHuaweiSuccess(count int64) {
	atomic.AddInt64(&s.stat.Huawei.PushSuccess, count)
}

// AddHuaweiError record counts of error Huawei push notification.
func (s *Storage) AddHuaweiError(count int64) {
	atomic.AddInt64(&s.stat.Huawei.PushError, count)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	count := atomic.LoadInt64(&s.stat.TotalCount)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	count := atomic.LoadInt64(&s.stat.Ios.PushSuccess)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	count := atomic.LoadInt64(&s.stat.Ios.PushError)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	count := atomic.LoadInt64(&s.stat.Android.PushSuccess)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	count := atomic.LoadInt64(&s.stat.Android.PushError)

	return count
}

// GetHuaweiSuccess show success counts of Huawei notification.
func (s *Storage) GetHuaweiSuccess() int64 {
	count := atomic.LoadInt64(&s.stat.Huawei.PushSuccess)

	return count
}

// GetHuaweiError show error counts of Huawei notification.
func (s *Storage) GetHuaweiError() int64 {
	count := atomic.LoadInt64(&s.stat.Huawei.PushError)

	return count
}
