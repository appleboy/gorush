package boltdb

import (
	"github.com/appleboy/gorush/gorush"
	"github.com/asdine/storm"
)

// Storage implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config gorush.ConfYaml, stat gorush.StatusApp) *Storage {
	return &Storage{
		stat:   stat,
		config: config,
	}
}

type Storage struct {
	config gorush.ConfYaml
	stat   gorush.StatusApp
}

func (s *Storage) initBoltDB() {
	s.stat.TotalCount = s.getTotalCount()
	s.stat.Ios.PushSuccess = s.getIosSuccess()
	s.stat.Ios.PushError = s.getIosError()
	s.stat.Android.PushSuccess = s.getAndroidSuccess()
	s.stat.Android.PushError = s.getAndroidError()
}

func (s *Storage) resetBoltDB() {
	s.setBoltDB(gorush.TotalCountKey, 0)
	s.setBoltDB(gorush.IosSuccessKey, 0)
	s.setBoltDB(gorush.IosErrorKey, 0)
	s.setBoltDB(gorush.AndroidSuccessKey, 0)
	s.setBoltDB(gorush.AndroidErrorKey, 0)
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

func (s *Storage) addTotalCount(count int64) {
	total := s.getTotalCount() + count
	s.setBoltDB(gorush.TotalCountKey, total)
}

func (s *Storage) addIosSuccess(count int64) {
	total := s.getIosSuccess() + count
	s.setBoltDB(gorush.IosSuccessKey, total)
}

func (s *Storage) addIosError(count int64) {
	total := s.getIosError() + count
	s.setBoltDB(gorush.IosErrorKey, total)
}

func (s *Storage) addAndroidSuccess(count int64) {
	total := s.getAndroidSuccess() + count
	s.setBoltDB(gorush.AndroidSuccessKey, total)
}

func (s *Storage) addAndroidError(count int64) {
	total := s.getAndroidError() + count
	s.setBoltDB(gorush.AndroidErrorKey, total)
}

func (s *Storage) getTotalCount() int64 {
	var count int64
	s.getBoltDB(gorush.TotalCountKey, &count)

	return count
}

func (s *Storage) getIosSuccess() int64 {
	var count int64
	s.getBoltDB(gorush.IosSuccessKey, &count)

	return count
}

func (s *Storage) getIosError() int64 {
	var count int64
	s.getBoltDB(gorush.IosErrorKey, &count)

	return count
}

func (s *Storage) getAndroidSuccess() int64 {
	var count int64
	s.getBoltDB(gorush.AndroidSuccessKey, &count)

	return count
}

func (s *Storage) getAndroidError() int64 {
	var count int64
	s.getBoltDB(gorush.AndroidErrorKey, &count)

	return count
}
