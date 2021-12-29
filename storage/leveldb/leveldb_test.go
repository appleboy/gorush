package leveldb

import (
	"github.com/appleboy/gorush/storage"
	"os"
	"sync"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestLevelDBEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	if _, err := os.Stat(cfg.Stat.LevelDB.Path); os.IsNotExist(err) {
		err = os.RemoveAll(cfg.Stat.LevelDB.Path)
		assert.Nil(t, err)
	}

	levelDB := New(cfg)
	err := levelDB.Init()
	assert.Nil(t, err)

	levelDB.Add(storage.HuaweiSuccessKey, 10)
	val = levelDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	levelDB.Add(storage.HuaweiSuccessKey, 20)
	val = levelDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	levelDB.Set(storage.HuaweiSuccessKey, 0)
	val = levelDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			levelDB.Add(storage.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = levelDB.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, levelDB.Close())
}
