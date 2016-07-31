package memory

import (
	"sync/atomic"
)

// StatusApp is app status structure
type statApp struct {
	TotalCount int64         `json:"total_count"`
	Ios        IosStatus     `json:"ios"`
	Android    AndroidStatus `json:"android"`
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

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New() *Storage {
	return &Storage{
		stat: &statApp{},
	}
}

type Storage struct {
	stat *statApp
}

func (s *Storage) Init() error {
	return nil
}

func (s *Storage) Reset() {
}

func (s *Storage) AddTotalCount(count int64) {
	atomic.AddInt64(&s.stat.TotalCount, count)
}

func (s *Storage) AddIosSuccess(count int64) {
	atomic.AddInt64(&s.stat.Ios.PushSuccess, count)
}

func (s *Storage) AddIosError(count int64) {
	atomic.AddInt64(&s.stat.Ios.PushError, count)
}

func (s *Storage) AddAndroidSuccess(count int64) {
	atomic.AddInt64(&s.stat.Android.PushSuccess, count)
}

func (s *Storage) AddAndroidError(count int64) {
	atomic.AddInt64(&s.stat.Android.PushError, count)
}

func (s *Storage) GetTotalCount() int64 {
	count := atomic.LoadInt64(&s.stat.TotalCount)

	return count
}

func (s *Storage) GetIosSuccess() int64 {
	count := atomic.LoadInt64(&s.stat.Ios.PushSuccess)

	return count
}

func (s *Storage) GetIosError() int64 {
	count := atomic.LoadInt64(&s.stat.Ios.PushError)

	return count
}

func (s *Storage) GetAndroidSuccess() int64 {
	count := atomic.LoadInt64(&s.stat.Android.PushSuccess)

	return count
}

func (s *Storage) GetAndroidError() int64 {
	count := atomic.LoadInt64(&s.stat.Android.PushError)

	return count
}
