package redis

import (
	"github.com/appleboy/gorush/gorush"
	"gopkg.in/redis.v3"
	"log"
	"strconv"
)

var RedisClient *redis.Client

// Storage implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config gorush.ConfYaml, stat gorush.StatusApp) *Storage {
	return &Storage{
		stat:   stat,
		config: config,
	}
}

func getRedisInt64Result(key string, count *int64) {
	val, _ := RedisClient.Get(key).Result()
	*count, _ = strconv.ParseInt(val, 10, 64)
}

type Storage struct {
	config gorush.ConfYaml
	stat   gorush.StatusApp
}

func (s *Storage) initRedis() error {
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

	s.stat.TotalCount = s.getTotalCount()
	s.stat.Ios.PushSuccess = s.getIosSuccess()
	s.stat.Ios.PushError = s.getIosError()
	s.stat.Android.PushSuccess = s.getAndroidSuccess()
	s.stat.Android.PushError = s.getAndroidError()

	return nil
}

func (s *Storage) addTotalCount(count int64) {
	RedisClient.Set(gorush.TotalCountKey, strconv.Itoa(int(count)), 0)
}

func (s *Storage) addIosSuccess(count int64) {
	RedisClient.Set(gorush.IosSuccessKey, strconv.Itoa(int(count)), 0)
}

func (s *Storage) addIosError(count int64) {
	RedisClient.Set(gorush.IosErrorKey, strconv.Itoa(int(count)), 0)
}

func (s *Storage) addAndroidSuccess(count int64) {
	RedisClient.Set(gorush.AndroidSuccessKey, strconv.Itoa(int(count)), 0)
}

func (s *Storage) addAndroidError(count int64) {
	RedisClient.Set(gorush.AndroidErrorKey, strconv.Itoa(int(count)), 0)
}

func (s *Storage) getTotalCount() int64 {
	var count int64
	getRedisInt64Result(gorush.TotalCountKey, &count)

	return count
}

func (s *Storage) getIosSuccess() int64 {
	var count int64
	getRedisInt64Result(gorush.IosSuccessKey, &count)

	return count
}

func (s *Storage) getIosError() int64 {
	var count int64
	getRedisInt64Result(gorush.IosErrorKey, &count)

	return count
}

func (s *Storage) getAndroidSuccess() int64 {
	var count int64
	getRedisInt64Result(gorush.AndroidSuccessKey, &count)

	return count
}

func (s *Storage) getAndroidError() int64 {
	var count int64
	getRedisInt64Result(gorush.AndroidErrorKey, &count)

	return count
}
