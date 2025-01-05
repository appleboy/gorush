package buntdb

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
)

func TestBuntDBEngine(t *testing.T) {
	var val int64

	buntDB := New("")
	err := buntDB.Init()
	assert.Nil(t, err)

	// reset the value of the key to 0
	buntDB.Set(core.HuaweiSuccessKey, 0)
	val = buntDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	buntDB.Add(core.HuaweiSuccessKey, 10)
	val = buntDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	buntDB.Add(core.HuaweiSuccessKey, 10)
	val = buntDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	buntDB.Set(core.HuaweiSuccessKey, 0)
	val = buntDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			buntDB.Add(core.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = buntDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, buntDB.Close())
}
