package gorush

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryEnqueue(t *testing.T) {
	chn := make(chan PushNotification, 2)
	assert.True(t, tryEnqueue(PushNotification{}, chn))
	assert.Equal(t, 1, len(chn))
	assert.True(t, tryEnqueue(PushNotification{}, chn))
	assert.Equal(t, 2, len(chn))
	assert.False(t, tryEnqueue(PushNotification{}, chn))
	assert.Equal(t, 2, len(chn))
}
