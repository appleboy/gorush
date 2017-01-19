package leveldb

import (
	"fmt"
	"strconv"

	"github.com/appleboy/gorush/config"
	"github.com/syndtr/goleveldb/leveldb"
)

// Stat variable for redis
const (
	TotalCountKey     = "gorush-total-count"
	IosSuccessKey     = "gorush-ios-success-count"
	IosErrorKey       = "gorush-ios-error-count"
	AndroidSuccessKey = "gorush-android-success-count"
	AndroidErrorKey   = "gorush-android-error-count"
)

var dbPath string

func setLevelDB(key string, count int64) {
	db, _ := leveldb.OpenFile(dbPath, nil)
	value := fmt.Sprintf("%d", count)

	_ = db.Put([]byte(key), []byte(value), nil)

	defer db.Close()
}

func getLevelDB(key string, count *int64) {
	db, _ := leveldb.OpenFile(dbPath, nil)

	data, _ := db.Get([]byte(key), nil)
	*count, _ = strconv.ParseInt(string(data), 10, 64)

	defer db.Close()
}

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
	dbPath = s.config.Stat.LevelDB.Path
	return nil
}

// Reset Client storage.
func (s *Storage) Reset() {
	setLevelDB(TotalCountKey, 0)
	setLevelDB(IosSuccessKey, 0)
	setLevelDB(IosErrorKey, 0)
	setLevelDB(AndroidSuccessKey, 0)
	setLevelDB(AndroidErrorKey, 0)
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	setLevelDB(TotalCountKey, total)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	setLevelDB(IosSuccessKey, total)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	setLevelDB(IosErrorKey, total)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	setLevelDB(AndroidSuccessKey, total)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	setLevelDB(AndroidErrorKey, total)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	getLevelDB(TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	getLevelDB(IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	getLevelDB(IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	getLevelDB(AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	getLevelDB(AndroidErrorKey, &count)

	return count
}
