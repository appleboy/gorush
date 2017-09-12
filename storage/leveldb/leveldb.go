package leveldb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/axiomzen/gorush/config"
	"github.com/axiomzen/gorush/storage"

	"github.com/syndtr/goleveldb/leveldb"
)

var dbPath string

func setLevelDB(key string, count int64) {
	db, _ := leveldb.OpenFile(dbPath, nil)
	value := fmt.Sprintf("%d", count)

	_ = db.Put([]byte(key), []byte(value), nil)

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("LevelDB error:", err.Error())
		}
	}()
}

func getLevelDB(key string, count *int64) {
	db, _ := leveldb.OpenFile(dbPath, nil)

	data, _ := db.Get([]byte(key), nil)
	*count, _ = strconv.ParseInt(string(data), 10, 64)

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("LevelDB error:", err.Error())
		}
	}()
}

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
	dbPath = s.config.Stat.LevelDB.Path
	return nil
}

// Reset Client storage.
func (s *Storage) Reset() {
	setLevelDB(storage.TotalCountKey, 0)
	setLevelDB(storage.IosSuccessKey, 0)
	setLevelDB(storage.IosErrorKey, 0)
	setLevelDB(storage.AndroidSuccessKey, 0)
	setLevelDB(storage.AndroidErrorKey, 0)
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	setLevelDB(storage.TotalCountKey, total)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	setLevelDB(storage.IosSuccessKey, total)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	setLevelDB(storage.IosErrorKey, total)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	setLevelDB(storage.AndroidSuccessKey, total)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	setLevelDB(storage.AndroidErrorKey, total)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	getLevelDB(storage.TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	getLevelDB(storage.IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	getLevelDB(storage.IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	getLevelDB(storage.AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	getLevelDB(storage.AndroidErrorKey, &count)

	return count
}
