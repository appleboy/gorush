package redis

import (
	"strconv"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/storage"

	"github.com/go-redis/redis/v7"
)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) getInt64(key string, count *int64) {
	val, _ := s.client.Get(key).Result()
	*count, _ = strconv.ParseInt(val, 10, 64)
}

// Storage is interface structure
type Storage struct {
	config config.ConfYaml
	client *redis.Client
}

// Init client storage.
func (s *Storage) Init() error {
	s.client = redis.NewClient(&redis.Options{
		Addr:     s.config.Stat.Redis.Addr,
		Password: s.config.Stat.Redis.Password,
		DB:       s.config.Stat.Redis.DB,
	})
	_, err := s.client.Ping().Result()

	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.client == nil {
		return nil
	}

	return s.client.Close()
}

// Reset Client storage.
func (s *Storage) Reset() {
	s.client.Set(storage.TotalCountKey, int64(0), 0)
	s.client.Set(storage.IosSuccessKey, int64(0), 0)
	s.client.Set(storage.IosErrorKey, int64(0), 0)
	s.client.Set(storage.AndroidSuccessKey, int64(0), 0)
	s.client.Set(storage.AndroidErrorKey, int64(0), 0)
	s.client.Set(storage.HuaweiSuccessKey, int64(0), 0)
	s.client.Set(storage.HuaweiErrorKey, int64(0), 0)
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	s.client.IncrBy(storage.TotalCountKey, count)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	s.client.IncrBy(storage.IosSuccessKey, count)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	s.client.IncrBy(storage.IosErrorKey, count)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	s.client.IncrBy(storage.AndroidSuccessKey, count)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	s.client.IncrBy(storage.AndroidErrorKey, count)
}

// AddHuaweiSuccess record counts of success Android push notification.
func (s *Storage) AddHuaweiSuccess(count int64) {
	s.client.IncrBy(storage.HuaweiSuccessKey, count)
}

// AddHuaweiError record counts of error Android push notification.
func (s *Storage) AddHuaweiError(count int64) {
	s.client.IncrBy(storage.HuaweiErrorKey, count)
}

// GetTotalCount show counts of all notification.
func (s *Storage) GetTotalCount() int64 {
	var count int64
	s.getInt64(storage.TotalCountKey, &count)

	return count
}

// GetIosSuccess show success counts of iOS notification.
func (s *Storage) GetIosSuccess() int64 {
	var count int64
	s.getInt64(storage.IosSuccessKey, &count)

	return count
}

// GetIosError show error counts of iOS notification.
func (s *Storage) GetIosError() int64 {
	var count int64
	s.getInt64(storage.IosErrorKey, &count)

	return count
}

// GetAndroidSuccess show success counts of Android notification.
func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	s.getInt64(storage.AndroidSuccessKey, &count)

	return count
}

// GetAndroidError show error counts of Android notification.
func (s *Storage) GetAndroidError() int64 {
	var count int64
	s.getInt64(storage.AndroidErrorKey, &count)

	return count
}

// GetHuaweiSuccess show success counts of Huawei notification.
func (s *Storage) GetHuaweiSuccess() int64 {
	var count int64
	s.getInt64(storage.HuaweiSuccessKey, &count)

	return count
}

// GetHuaweiError show error counts of Huawei notification.
func (s *Storage) GetHuaweiError() int64 {
	var count int64
	s.getInt64(storage.HuaweiErrorKey, &count)

	return count
}
