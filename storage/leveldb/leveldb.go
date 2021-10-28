package leveldb

import (
	"fmt"
	"strconv"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/storage"

	"github.com/syndtr/goleveldb/leveldb"
)

func (s *Storage) setLevelDB(key string, count int64) {
	value := fmt.Sprintf("%d", count)
	_ = s.db.Put([]byte(key), []byte(value), nil)
}

func (s *Storage) getLevelDB(key string, count *int64) {
	data, _ := s.db.Get([]byte(key), nil)
	*count, _ = strconv.ParseInt(string(data), 10, 64)
}

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config *config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config *config.ConfYaml
	db     *leveldb.DB
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.db, err = leveldb.OpenFile(s.config.Stat.LevelDB.Path, nil)
	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}

	return s.db.Close()
}

// Reset Client storage.
func (s *Storage) Reset() {
	s.setLevelDB(storage.TotalCountKey, 0)
	s.setLevelDB(storage.IosSuccessKey, 0)
	s.setLevelDB(storage.IosErrorKey, 0)
	s.setLevelDB(storage.AndroidSuccessKey, 0)
	s.setLevelDB(storage.AndroidErrorKey, 0)
	s.setLevelDB(storage.HuaweiSuccessKey, 0)
	s.setLevelDB(storage.HuaweiErrorKey, 0)
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	s.setLevelDB(storage.TotalCountKey, total)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	s.setLevelDB(storage.IosSuccessKey, total)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	s.setLevelDB(storage.IosErrorKey, total)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	s.setLevelDB(storage.AndroidSuccessKey, total)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	s.setLevelDB(storage.AndroidErrorKey, total)
}

// AddHuaweiSuccess record counts of success Huawei push notification.
func (s *Storage) AddHuaweiSuccess(count int64) {
	total := s.GetHuaweiSuccess() + count
	s.setLevelDB(storage.HuaweiSuccessKey, total)
}

// AddHuaweiError record counts of error Huawei push notification.
func (s *Storage) AddHuaweiError(count int64) {
	total := s.GetHuaweiError() + count
	s.setLevelDB(storage.HuaweiErrorKey, total)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	s.getLevelDB(storage.TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	s.getLevelDB(storage.IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	s.getLevelDB(storage.IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	s.getLevelDB(storage.AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	s.getLevelDB(storage.AndroidErrorKey, &count)

	return count
}

// GetHuaweiSuccess show success counts of Huawei notification.
func (s *Storage) GetHuaweiSuccess() int64 {
	var count int64
	s.getLevelDB(storage.HuaweiSuccessKey, &count)

	return count
}

// GetHuaweiError show error counts of Huawei notification.
func (s *Storage) GetHuaweiError() int64 {
	var count int64
	s.getLevelDB(storage.HuaweiErrorKey, &count)

	return count
}
