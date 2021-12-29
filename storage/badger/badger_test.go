package badger

import (
	"github.com/appleboy/gorush/storage"
	"sync"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestBadgerEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	badger := New(cfg)
	err := badger.Init()
	assert.Nil(t, err)

	badger.Add(storage.HuaweiSuccessKey, 10)
	val = badger.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	badger.Add(storage.HuaweiSuccessKey, 20)
	val = badger.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	badger.Set(storage.HuaweiSuccessKey, 0)
	val = badger.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			badger.Add(storage.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = badger.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, badger.Close())
}
