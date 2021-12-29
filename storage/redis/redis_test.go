package redis

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/storage"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestRedisServerError(t *testing.T) {
	cfg, _ := config.LoadConf()
	cfg.Stat.Redis.Addr = "redis:6370"

	redis := New(cfg)
	err := redis.Init()

	assert.Error(t, err)
}

func TestRedisEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()
	cfg.Stat.Redis.Addr = "redis:6379"

	redis := New(cfg)
	err := redis.Init()
	assert.Nil(t, err)

	redis.Add(storage.HuaweiSuccessKey, 10)
	val = redis.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	redis.Add(storage.HuaweiSuccessKey, 10)
	val = redis.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	redis.Set(storage.HuaweiSuccessKey, 0)
	val = redis.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			redis.Add(storage.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = redis.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, redis.Close())
}
