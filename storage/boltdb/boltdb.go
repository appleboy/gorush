package boltdb

import (
	"log"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/storage"

	"github.com/asdine/storm/v3"
)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config *config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config *config.ConfYaml
	db     *storm.DB
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.db, err = storm.Open(s.config.Stat.BoltDB.Path)
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
	s.setBoltDB(storage.TotalCountKey, 0)
	s.setBoltDB(storage.IosSuccessKey, 0)
	s.setBoltDB(storage.IosErrorKey, 0)
	s.setBoltDB(storage.AndroidSuccessKey, 0)
	s.setBoltDB(storage.AndroidErrorKey, 0)
	s.setBoltDB(storage.HuaweiSuccessKey, 0)
	s.setBoltDB(storage.HuaweiErrorKey, 0)
}

func (s *Storage) setBoltDB(key string, count int64) {
	err := s.db.Set(s.config.Stat.BoltDB.Bucket, key, count)
	if err != nil {
		log.Println("BoltDB set error:", err.Error())
	}
}

func (s *Storage) getBoltDB(key string, count *int64) {
	err := s.db.Get(s.config.Stat.BoltDB.Bucket, key, count)
	if err != nil {
		log.Println("BoltDB get error:", err.Error())
	}
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	s.setBoltDB(storage.TotalCountKey, total)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	s.setBoltDB(storage.IosSuccessKey, total)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	s.setBoltDB(storage.IosErrorKey, total)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	s.setBoltDB(storage.AndroidSuccessKey, total)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	s.setBoltDB(storage.AndroidErrorKey, total)
}

// AddHuaweiSuccess record counts of success Huawei push notification.
func (s *Storage) AddHuaweiSuccess(count int64) {
	total := s.GetHuaweiSuccess() + count
	s.setBoltDB(storage.HuaweiSuccessKey, total)
}

// AddHuaweiError record counts of error Huawei push notification.
func (s *Storage) AddHuaweiError(count int64) {
	total := s.GetHuaweiError() + count
	s.setBoltDB(storage.HuaweiErrorKey, total)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	s.getBoltDB(storage.TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	s.getBoltDB(storage.IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	s.getBoltDB(storage.IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	s.getBoltDB(storage.AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	s.getBoltDB(storage.AndroidErrorKey, &count)

	return count
}

// GetHuaweiSuccess show success counts of Huawei notification.
func (s *Storage) GetHuaweiSuccess() int64 {
	var count int64
	s.getBoltDB(storage.HuaweiSuccessKey, &count)

	return count
}

// GetHuaweiError show error counts of Huawei notification.
func (s *Storage) GetHuaweiError() int64 {
	var count int64
	s.getBoltDB(storage.HuaweiErrorKey, &count)

	return count
}
