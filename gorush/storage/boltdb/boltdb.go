package boltdb

import (
	"github.com/appleboy/gorush/gorush/config"
	"github.com/asdine/storm"
)

// Stat variable for redis
const (
	TotalCountKey     = "gorush-total-count"
	IosSuccessKey     = "gorush-ios-success-count"
	IosErrorKey       = "gorush-ios-error-count"
	AndroidSuccessKey = "gorush-android-success-count"
	AndroidErrorKey   = "gorush-android-error-count"
)

// Storage implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

type Storage struct {
	config config.ConfYaml
}

func (s *Storage) Init() error {
	return nil
}

func (s *Storage) Reset() {
	s.setBoltDB(TotalCountKey, 0)
	s.setBoltDB(IosSuccessKey, 0)
	s.setBoltDB(IosErrorKey, 0)
	s.setBoltDB(AndroidSuccessKey, 0)
	s.setBoltDB(AndroidErrorKey, 0)
}

func (s *Storage) setBoltDB(key string, count int64) {
	db, _ := storm.Open(s.config.Stat.BoltDB.Path)
	db.Set(s.config.Stat.BoltDB.Bucket, key, count)
	defer db.Close()
}

func (s *Storage) getBoltDB(key string, count *int64) {
	db, _ := storm.Open(s.config.Stat.BoltDB.Path)
	db.Get(s.config.Stat.BoltDB.Bucket, key, count)
	defer db.Close()
}

func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	s.setBoltDB(TotalCountKey, total)
}

func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	s.setBoltDB(IosSuccessKey, total)
}

func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	s.setBoltDB(IosErrorKey, total)
}

func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	s.setBoltDB(AndroidSuccessKey, total)
}

func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	s.setBoltDB(AndroidErrorKey, total)
}

func (s *Storage) GetTotalCount() int64 {
	var count int64
	s.getBoltDB(TotalCountKey, &count)

	return count
}

func (s *Storage) GetIosSuccess() int64 {
	var count int64
	s.getBoltDB(IosSuccessKey, &count)

	return count
}

func (s *Storage) GetIosError() int64 {
	var count int64
	s.getBoltDB(IosErrorKey, &count)

	return count
}

func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	s.getBoltDB(AndroidSuccessKey, &count)

	return count
}

func (s *Storage) GetAndroidError() int64 {
	var count int64
	s.getBoltDB(AndroidErrorKey, &count)

	return count
}
