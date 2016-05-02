package redis

import (
	"github.com/appleboy/gorush/gorush/config"
	"gopkg.in/redis.v3"
	"log"
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

var RedisClient *redis.Client

// Storage implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

func getRedisInt64Result(key string, count *int64) {
	val, _ := RedisClient.Get(key).Result()
	*count, _ = strconv.ParseInt(val, 10, 64)
}

type Storage struct {
	config config.ConfYaml
}

func (s *Storage) Init() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     s.config.Stat.Redis.Addr,
		Password: s.config.Stat.Redis.Password,
		DB:       s.config.Stat.Redis.DB,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		// redis server error
		log.Println("Can't connect redis server: " + err.Error())

		return err
	}

	return nil
}

func (s *Storage) Reset() {
	RedisClient.Set(TotalCountKey, strconv.Itoa(0), 0)
	RedisClient.Set(IosSuccessKey, strconv.Itoa(0), 0)
	RedisClient.Set(IosErrorKey, strconv.Itoa(0), 0)
	RedisClient.Set(AndroidSuccessKey, strconv.Itoa(0), 0)
	RedisClient.Set(AndroidErrorKey, strconv.Itoa(0), 0)
}

func (s *Storage) AddTotalCount(count int64) {
	total := s.GetTotalCount() + count
	RedisClient.Set(TotalCountKey, strconv.Itoa(int(total)), 0)
}

func (s *Storage) AddIosSuccess(count int64) {
	total := s.GetIosSuccess() + count
	RedisClient.Set(IosSuccessKey, strconv.Itoa(int(total)), 0)
}

func (s *Storage) AddIosError(count int64) {
	total := s.GetIosError() + count
	RedisClient.Set(IosErrorKey, strconv.Itoa(int(total)), 0)
}

func (s *Storage) AddAndroidSuccess(count int64) {
	total := s.GetAndroidSuccess() + count
	RedisClient.Set(AndroidSuccessKey, strconv.Itoa(int(total)), 0)
}

func (s *Storage) AddAndroidError(count int64) {
	total := s.GetAndroidError() + count
	RedisClient.Set(AndroidErrorKey, strconv.Itoa(int(total)), 0)
}

func (s *Storage) GetTotalCount() int64 {
	var count int64
	getRedisInt64Result(TotalCountKey, &count)

	return count
}

func (s *Storage) GetIosSuccess() int64 {
	var count int64
	getRedisInt64Result(IosSuccessKey, &count)

	return count
}

func (s *Storage) GetIosError() int64 {
	var count int64
	getRedisInt64Result(IosErrorKey, &count)

	return count
}

func (s *Storage) GetAndroidSuccess() int64 {
	var count int64
	getRedisInt64Result(AndroidSuccessKey, &count)

	return count
}

func (s *Storage) GetAndroidError() int64 {
	var count int64
	getRedisInt64Result(AndroidErrorKey, &count)

	return count
}
