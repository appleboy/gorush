package gorush

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestMissingKeyForInitFCMClient(t *testing.T) {
	config.BuildDefaultPushConf()

	client, err := InitFCMClient("")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "Missing Android API Key", err.Error())
}
