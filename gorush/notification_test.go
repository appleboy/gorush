package gorush

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestCorrectConf(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = "xxxxx"

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"

	err := CheckPushConf()

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
