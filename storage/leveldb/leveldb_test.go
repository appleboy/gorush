package leveldb

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLevelDBEngine(t *testing.T) {
	var val int64

	levelDB := New("")
	err := levelDB.Init()
	require.NoError(t, err)

	// reset the value of the key to 0
	levelDB.Set(core.HuaweiSuccessKey, 0)
	val = levelDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	levelDB.Add(core.HuaweiSuccessKey, 10)
	val = levelDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	levelDB.Add(core.HuaweiSuccessKey, 10)
	val = levelDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	levelDB.Set(core.HuaweiSuccessKey, 0)
	val = levelDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() {
			levelDB.Add(core.HuaweiSuccessKey, 1)
		})
	}
	wg.Wait()
	val = levelDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, levelDB.Close())
}
