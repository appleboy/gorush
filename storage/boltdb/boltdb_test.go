package boltdb

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/storage"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestBoltDBEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	boltDB := New(cfg)
	err := boltDB.Init()
	assert.Nil(t, err)

	boltDB.Add(storage.HuaweiSuccessKey, 10)
	val = boltDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	boltDB.Add(storage.HuaweiSuccessKey, 10)
	val = boltDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	boltDB.Set(storage.HuaweiSuccessKey, 0)
	val = boltDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			boltDB.Add(storage.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = boltDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, boltDB.Close())
}
