package status

import "github.com/appleboy/gorush/storage"

type StateStorage struct {
	store storage.Storage
}

func NewStateStorage(store storage.Storage) *StateStorage {
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
	s.store.Set(storage.TotalCountKey, 0)
	s.store.Set(storage.IosSuccessKey, 0)
	s.store.Set(storage.IosErrorKey, 0)
	s.store.Set(storage.AndroidSuccessKey, 0)
	s.store.Set(storage.AndroidErrorKey, 0)
	s.store.Set(storage.HuaweiSuccessKey, 0)
	s.store.Set(storage.HuaweiErrorKey, 0)
}

// AddTotalCount record push notification count.
func (s *StateStorage) AddTotalCount(count int64) {
	s.store.Add(storage.TotalCountKey, count)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *StateStorage) AddIosSuccess(count int64) {
	s.store.Add(storage.IosSuccessKey, count)
}

// AddIosError record counts of error iOS push notification.
func (s *StateStorage) AddIosError(count int64) {
	s.store.Add(storage.IosErrorKey, count)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *StateStorage) AddAndroidSuccess(count int64) {
	s.store.Add(storage.AndroidSuccessKey, count)
}

// AddAndroidError record counts of error Android push notification.
func (s *StateStorage) AddAndroidError(count int64) {
	s.store.Add(storage.AndroidErrorKey, count)
}

// AddHuaweiSuccess record counts of success Huawei push notification.
func (s *StateStorage) AddHuaweiSuccess(count int64) {
	s.store.Add(storage.HuaweiSuccessKey, count)
}

// AddHuaweiError record counts of error Huawei push notification.
func (s *StateStorage) AddHuaweiError(count int64) {
	s.store.Add(storage.HuaweiErrorKey, count)
}

// GetTotalCount show counts of all notification.
func (s *StateStorage) GetTotalCount() int64 {
	return s.store.Get(storage.TotalCountKey)
}

// GetIosSuccess show success counts of iOS notification.
func (s *StateStorage) GetIosSuccess() int64 {
	return s.store.Get(storage.IosSuccessKey)
}

// GetIosError show error counts of iOS notification.
func (s *StateStorage) GetIosError() int64 {
	return s.store.Get(storage.IosErrorKey)
}

// GetAndroidSuccess show success counts of Android notification.
func (s *StateStorage) GetAndroidSuccess() int64 {
	return s.store.Get(storage.AndroidSuccessKey)
}

// GetAndroidError show error counts of Android notification.
func (s *StateStorage) GetAndroidError() int64 {
	return s.store.Get(storage.AndroidErrorKey)
}

// GetHuaweiSuccess show success counts of Huawei notification.
func (s *StateStorage) GetHuaweiSuccess() int64 {
	return s.store.Get(storage.HuaweiSuccessKey)
}

// GetHuaweiError show error counts of Huawei notification.
func (s *StateStorage) GetHuaweiError() int64 {
	return s.store.Get(storage.HuaweiErrorKey)
}
