package buntdb

import (
	"fmt"
	"github.com/appleboy/gorush/config"
	"github.com/tidwall/buntdb"
	"strconv"
)

// Stat variable for redis
const (
	TotalCountKey     = "gorush-total-count"
	IosSuccessKey     = "gorush-ios-success-count"
	IosErrorKey       = "gorush-ios-error-count"
	AndroidSuccessKey = "gorush-android-success-count"
	AndroidErrorKey   = "gorush-android-error-count"
)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config config.ConfYaml
}

// Init client storage.
func (s *Storage) Init() error {
	return nil
}

// Reset Client storage.
func (s *Storage) Reset() {
	s.setBuntDB(TotalCountKey, 0)
	s.setBuntDB(IosSuccessKey, 0)
	s.setBuntDB(IosErrorKey, 0)
	s.setBuntDB(AndroidSuccessKey, 0)
	s.setBuntDB(AndroidErrorKey, 0)
}

func (s *Storage) setBuntDB(key string, count int64) {
	db, _ := buntdb.Open(s.config.Stat.BuntDB.Path)

	db.Update(func(tx *buntdb.Tx) error {
		tx.Set(key, fmt.Sprintf("%d", count), nil)
		return nil
	})
	defer db.Close()
}

func (s *Storage) getBuntDB(key string, count *int64) {
	db, _ := buntdb.Open(s.config.Stat.BuntDB.Path)

	db.View(func(tx *buntdb.Tx) error {
		val, _ := tx.Get(key)
		*count, _ = strconv.ParseInt(val, 10, 64)
		return nil
	})
	defer db.Close()
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	s.setBuntDB(TotalCountKey, total)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	s.setBuntDB(IosSuccessKey, total)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	s.setBuntDB(IosErrorKey, total)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	s.setBuntDB(AndroidSuccessKey, total)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	s.setBuntDB(AndroidErrorKey, total)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	s.getBuntDB(TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	s.getBuntDB(IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	s.getBuntDB(IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	s.getBuntDB(AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	s.getBuntDB(AndroidErrorKey, &count)

	return count
}
