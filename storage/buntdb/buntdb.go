package buntdb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/axiomzen/gorush/config"
	"github.com/axiomzen/gorush/storage"

	"github.com/tidwall/buntdb"
)

// New func implements the storage interface for gorush (https://github.com/axiomzen/gorush)
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
	s.setBuntDB(storage.TotalCountKey, 0)
	s.setBuntDB(storage.IosSuccessKey, 0)
	s.setBuntDB(storage.IosErrorKey, 0)
	s.setBuntDB(storage.AndroidSuccessKey, 0)
	s.setBuntDB(storage.AndroidErrorKey, 0)
}

func (s *Storage) setBuntDB(key string, count int64) {
	db, _ := buntdb.Open(s.config.Stat.BuntDB.Path)

	err := db.Update(func(tx *buntdb.Tx) error {
		if _, _, err := tx.Set(key, fmt.Sprintf("%d", count), nil); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Println("BuntDB update error:", err.Error())
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("BuntDB error:", err.Error())
		}
	}()
}

func (s *Storage) getBuntDB(key string, count *int64) {
	db, _ := buntdb.Open(s.config.Stat.BuntDB.Path)

	err := db.View(func(tx *buntdb.Tx) error {
		val, _ := tx.Get(key)
		*count, _ = strconv.ParseInt(val, 10, 64)
		return nil
	})

	if err != nil {
		log.Println("BuntDB get error:", err.Error())
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("BuntDB error:", err.Error())
		}
	}()
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	s.setBuntDB(storage.TotalCountKey, total)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	s.setBuntDB(storage.IosSuccessKey, total)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	s.setBuntDB(storage.IosErrorKey, total)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	s.setBuntDB(storage.AndroidSuccessKey, total)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	s.setBuntDB(storage.AndroidErrorKey, total)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	s.getBuntDB(storage.TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	s.getBuntDB(storage.IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	s.getBuntDB(storage.IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	s.getBuntDB(storage.AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	s.getBuntDB(storage.AndroidErrorKey, &count)

	return count
}
