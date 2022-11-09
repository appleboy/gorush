package status

import (
	"github.com/appleboy/gorush/core"
)

type StateStorage struct {
	store core.Storage
}

func NewStateStorage(store core.Storage) *StateStorage {
	return &StateStorage{
		store: store,
	}
}

func (s *StateStorage) Init() error {
	return s.store.Init()
}

func (s *StateStorage) Close() error {
	return s.store.Close()
}

// Reset Client storage.
func (s *StateStorage) Reset() {
	s.store.Set(core.TotalCountKey, 0)
	s.store.Set(core.IosSuccessKey, 0)
	s.store.Set(core.IosErrorKey, 0)
	s.store.Set(core.AndroidSuccessKey, 0)
	s.store.Set(core.AndroidErrorKey, 0)
	s.store.Set(core.HuaweiSuccessKey, 0)
	s.store.Set(core.HuaweiErrorKey, 0)
}

// AddTotalCount record push notification count.
func (s *StateStorage) AddTotalCount(count int64) {
	s.store.Add(core.TotalCountKey, count)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *StateStorage) AddIosSuccess(count int64) {
	s.store.Add(core.IosSuccessKey, count)
}

// AddIosError record counts of error iOS push notification.
func (s *StateStorage) AddIosError(count int64) {
	s.store.Add(core.IosErrorKey, count)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *StateStorage) AddAndroidSuccess(count int64) {
	s.store.Add(core.AndroidSuccessKey, count)
}

// AddAndroidError record counts of error Android push notification.
func (s *StateStorage) AddAndroidError(count int64) {
	s.store.Add(core.AndroidErrorKey, count)
}

// AddHuaweiSuccess record counts of success Huawei push notification.
func (s *StateStorage) AddHuaweiSuccess(count int64) {
	s.store.Add(core.HuaweiSuccessKey, count)
}

// AddHuaweiError record counts of error Huawei push notification.
func (s *StateStorage) AddHuaweiError(count int64) {
	s.store.Add(core.HuaweiErrorKey, count)
}

// GetTotalCount show counts of all notification.
func (s *StateStorage) GetTotalCount() int64 {
	return s.store.Get(core.TotalCountKey)
}

// GetIosSuccess show success counts of iOS notification.
func (s *StateStorage) GetIosSuccess() int64 {
	return s.store.Get(core.IosSuccessKey)
}

// GetIosError show error counts of iOS notification.
func (s *StateStorage) GetIosError() int64 {
	return s.store.Get(core.IosErrorKey)
}

// GetAndroidSuccess show success counts of Android notification.
func (s *StateStorage) GetAndroidSuccess() int64 {
	return s.store.Get(core.AndroidSuccessKey)
}

// GetAndroidError show error counts of Android notification.
func (s *StateStorage) GetAndroidError() int64 {
	return s.store.Get(core.AndroidErrorKey)
}

// GetHuaweiSuccess show success counts of Huawei notification.
func (s *StateStorage) GetHuaweiSuccess() int64 {
	return s.store.Get(core.HuaweiSuccessKey)
}

// GetHuaweiError show error counts of Huawei notification.
func (s *StateStorage) GetHuaweiError() int64 {
	return s.store.Get(core.HuaweiErrorKey)
}
