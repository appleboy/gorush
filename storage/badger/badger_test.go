package badger

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
)

func TestBadgerEngine(t *testing.T) {
	var val int64

	badger := New("")
	err := badger.Init()
	assert.Nil(t, err)

	// reset the value of the key to 0
	badger.Set(core.HuaweiSuccessKey, 0)
	val = badger.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	badger.Add(core.HuaweiSuccessKey, 10)
	val = badger.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	badger.Add(core.HuaweiSuccessKey, 10)
	val = badger.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	badger.Set(core.HuaweiSuccessKey, 0)
	val = badger.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			badger.Add(core.HuaweiSuccessKey, 1)
		}()
	}
	wg.Wait()
	val = badger.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(100), val)

	assert.NoError(t, badger.Close())
}
