package memory

import (
	"github.com/appleboy/gorush/storage"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryEngine(t *testing.T) {
	var val int64

	memory := New()
	err := memory.Init()
	assert.Nil(t, err)

	memory.Add(storage.HuaweiSuccessKey, 10)
	val = memory.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)
	memory.Add(storage.HuaweiSuccessKey, 20)
	val = memory.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(20), val)

	memory.Set(storage.HuaweiSuccessKey, 0)
	val = memory.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(0), val)

	// test concurrency issues
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			memory.Add(storage.HuaweiSuccessKey, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	val = memory.Get(storage.HuaweiSuccessKey)
	assert.Equal(t, int64(10), val)

	assert.NoError(t, memory.Close())
}
