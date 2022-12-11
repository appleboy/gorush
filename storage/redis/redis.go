package redis

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/appleboy/gorush/config"

	"github.com/go-redis/redis/v9"
)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config *config.ConfYaml) *Storage {
	return &Storage{
		ctx:    context.Background(),
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	ctx    context.Context
	config *config.ConfYaml
	client redis.Cmdable
}

func (s *Storage) Add(key string, count int64) {
	s.client.IncrBy(s.ctx, key, count)
}

func (s *Storage) Set(key string, count int64) {
	s.client.Set(s.ctx, key, count, 0)
}

func (s *Storage) Get(key string) int64 {
	val, _ := s.client.Get(s.ctx, key).Result()
	count, _ := strconv.ParseInt(val, 10, 64)
	return count
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
