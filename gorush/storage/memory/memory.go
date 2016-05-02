package memory

import (
	"github.com/appleboy/gorush/gorush"
	"sync/atomic"
)

// Storage implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(stat gorush.StatusApp) *Storage {
	return &Storage{
		stat: stat,
	}
}

type Storage struct {
	stat gorush.StatusApp
}

func (s *Storage) addTotalCount(count int64) {
	atomic.AddInt64(&s.stat.TotalCount, count)
}

func (s *Storage) addIosSuccess(count int64) {
	atomic.AddInt64(&s.stat.Ios.PushSuccess, count)
}

func (s *Storage) addIosError(count int64) {
	atomic.AddInt64(&s.stat.Ios.PushError, count)
}

func (s *Storage) addAndroidSuccess(count int64) {
	atomic.AddInt64(&s.stat.Android.PushSuccess, count)
}

func (s *Storage) addAndroidError(count int64) {
	atomic.AddInt64(&s.stat.Android.PushError, count)
}

func (s *Storage) getTotalCount() int64 {
	count := atomic.LoadInt64(&s.stat.TotalCount)

	return count
}

func (s *Storage) getIosSuccess() int64 {
	count := atomic.LoadInt64(&s.stat.Ios.PushSuccess)

	return count
}

func (s *Storage) getIosError() int64 {
	count := atomic.LoadInt64(&s.stat.Ios.PushError)

	return count
}

func (s *Storage) getAndroidSuccess() int64 {
	count := atomic.LoadInt64(&s.stat.Android.PushSuccess)

	return count
}

func (s *Storage) getAndroidError() int64 {
	count := atomic.LoadInt64(&s.stat.Android.PushError)

	return count
}
