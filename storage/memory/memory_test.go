package memory

import (
	"sync"
	"testing"

	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryEngine(t *testing.T) {
	var val int64

	memory := New()
	err := memory.Init()
	require.NoError(t, err)

	// reset the value of the key to 0
	memory.Set(core.HuaweiSuccessKey, 0)
	val = memory.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	memory.Add(core.HuaweiSuccessKey, 10)
	val = memory.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	memory.Add(core.HuaweiSuccessKey, 10)
	val = memory.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	memory.Set(core.HuaweiSuccessKey, 0)
	val = memory.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() {
			memory.Add(core.HuaweiSuccessKey, 1)
		})
	}
	wg.Wait()
	val = memory.Get(core.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, memory.Close())
}
