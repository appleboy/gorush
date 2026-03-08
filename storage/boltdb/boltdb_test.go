package boltdb

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoltDBEngine(t *testing.T) {
	var val int64

	boltDB := New("", "gorush")
	err := boltDB.Init()
	require.NoError(t, err)

	// reset the value of the key to 0
	boltDB.Set(core.HuaweiSuccessKey, 0)
	val = boltDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

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
	for range 10 {
		wg.Go(func() {
			boltDB.Add(core.HuaweiSuccessKey, 1)
		})
	}
	wg.Wait()
	val = boltDB.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, boltDB.Close())
}
