package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetrics(t *testing.T) {
	m := NewMetrics()
	assert.Equal(t, 0, m.GetQueueUsage())

	m = NewMetrics(func() int { return 1 })
	assert.Equal(t, 1, m.GetQueueUsage())
}
