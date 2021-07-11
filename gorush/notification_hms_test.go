package gorush

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestMissingHuaweiAppSecret(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Huawei.Enabled = true
	PushConf.Huawei.AppSecret = ""

	err := CheckPushConf()

	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App Secret", err.Error())
}

func TestMissingHuaweiAppID(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Huawei.Enabled = true
	PushConf.Huawei.AppID = ""

	err := CheckPushConf()

	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App ID", err.Error())
}

func TestMissingAppSecretForInitHMSClient(t *testing.T) {
	client, err := InitHMSClient("", "APP_SECRET")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App Secret", err.Error())
}

func TestMissingAppIDForInitHMSClient(t *testing.T) {
	client, err := InitHMSClient("APP_ID", "")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App ID", err.Error())
}
