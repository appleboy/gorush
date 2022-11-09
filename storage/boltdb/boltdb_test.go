package boltdb

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
)

func TestBoltDBEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	boltDB := New(cfg)
	err := boltDB.Init()
	assert.Nil(t, err)

	boltDB.Add(core.HuaweiSuccessKey, 10)
	val = boltDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	boltDB.Add(core.HuaweiSuccessKey, 10)
	val = boltDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	boltDB.Set(core.HuaweiSuccessKey, 0)
	val = boltDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			boltDB.Add(core.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = boltDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, boltDB.Close())
}
