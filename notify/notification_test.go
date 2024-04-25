package notify

import (
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/stretchr/testify/assert"
)

func TestCorrectConf(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.ServiceAccountKey = "xxxxx"
	cfg.Android.ServiceAccountKey = "xxxxx"

	cfg.Ios.Enabled = true
	cfg.Ios.KeyPath = testKeyPath

	err := CheckPushConf(cfg)

	assert.NoError(t, err)
}

func TestSetProxyURL(t *testing.T) {
	err := SetProxy("87.236.233.92:8080")
	assert.Error(t, err)
	assert.Equal(t, "parse \"87.236.233.92:8080\": invalid URI for request", err.Error())

	err = SetProxy("a.html")
	assert.Error(t, err)

	err = SetProxy("http://87.236.233.92:8080")
	assert.NoError(t, err)
}
