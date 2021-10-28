package redis

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/storage"

	"github.com/go-redis/redis/v8"
)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config *config.ConfYaml) *Storage {
	return &Storage{
		ctx:    context.Background(),
		config: config,
	}
}

func (s *Storage) getInt64(key string, count *int64) {
	val, _ := s.client.Get(s.ctx, key).Result()
	*count, _ = strconv.ParseInt(val, 10, 64)
}

// Storage is interface structure
type Storage struct {
	ctx    context.Context
	config *config.ConfYaml
	client redis.Cmdable
}

// Init client storage.
func (s *Storage) Init() error {
	if s.config.Stat.Redis.Cluster {
		s.client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(s.config.Stat.Redis.Addr, ","),
			Password: s.config.Stat.Redis.Password,
		})
	} else {
		s.client = redis.NewClient(&redis.Options{
			Addr:     s.config.Stat.Redis.Addr,
			Password: s.config.Stat.Redis.Password,
			DB:       s.config.Stat.Redis.DB,
		})
	}

	if err := s.client.Ping(s.ctx).Err(); err != nil {
		return err
	}

	return nil
}

// Close the storage connection
func (s *Storage) Close() error {
	switch v := s.client.(type) {
	case *redis.Client:
		return v.Close()
	case *redis.ClusterClient:
		return v.Close()
	case nil:
		return nil
	default:
		// this will not happen anyway, unless we mishandle it on `Init`
		panic(fmt.Sprintf("invalid redis client: %v", reflect.TypeOf(v)))
	}
}

// Reset Client storage.
func (s *Storage) Reset() {
	s.client.Set(s.ctx, storage.TotalCountKey, int64(0), 0)
	s.client.Set(s.ctx, storage.IosSuccessKey, int64(0), 0)
	s.client.Set(s.ctx, storage.IosErrorKey, int64(0), 0)
	s.client.Set(s.ctx, storage.AndroidSuccessKey, int64(0), 0)
	s.client.Set(s.ctx, storage.AndroidErrorKey, int64(0), 0)
	s.client.Set(s.ctx, storage.HuaweiSuccessKey, int64(0), 0)
	s.client.Set(s.ctx, storage.HuaweiErrorKey, int64(0), 0)
}

// AddTotalCount record push notification count.
func (s *Storage) AddTotalCount(count int64) {
	s.client.IncrBy(s.ctx, storage.TotalCountKey, count)
}

// AddIosSuccess record counts of success iOS push notification.
func (s *Storage) AddIosSuccess(count int64) {
	s.client.IncrBy(s.ctx, storage.IosSuccessKey, count)
}

// AddIosError record counts of error iOS push notification.
func (s *Storage) AddIosError(count int64) {
	s.client.IncrBy(s.ctx, storage.IosErrorKey, count)
}

// AddAndroidSuccess record counts of success Android push notification.
func (s *Storage) AddAndroidSuccess(count int64) {
	s.client.IncrBy(s.ctx, storage.AndroidSuccessKey, count)
}

// AddAndroidError record counts of error Android push notification.
func (s *Storage) AddAndroidError(count int64) {
	s.client.IncrBy(s.ctx, storage.AndroidErrorKey, count)
}

// AddHuaweiSuccess record counts of success Android push notification.
func (s *Storage) AddHuaweiSuccess(count int64) {
	s.client.IncrBy(s.ctx, storage.HuaweiSuccessKey, count)
}

// AddHuaweiError record counts of error Android push notification.
func (s *Storage) AddHuaweiError(count int64) {
	s.client.IncrBy(s.ctx, storage.HuaweiErrorKey, count)
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
