package redis

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/appleboy/gorush/core"

	"github.com/redis/go-redis/v9"
)

var _ core.Storage = (*Storage)(nil)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(
	addr string,
	password string,
	db int,
	isCluster bool,
) *Storage {
	return &Storage{
		ctx:       context.Background(),
		addr:      addr,
		password:  password,
		db:        db,
		isCluster: isCluster,
	}
}

// Storage is interface structure
type Storage struct {
	ctx       context.Context
	client    redis.Cmdable
	addr      string
	password  string
	db        int
	isCluster bool
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
	if s.isCluster {
		s.client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(s.addr, ","),
			Password: s.password,
		})
	} else {
		s.client = redis.NewClient(&redis.Options{
			Addr:     s.addr,
			Password: s.password,
			DB:       s.db,
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
