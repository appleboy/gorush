package notify

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestMissingHuaweiAppSecret(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Huawei.Enabled = true
	cfg.Huawei.AppSecret = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App Secret", err.Error())
}

func TestMissingHuaweiAppID(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Huawei.Enabled = true
	cfg.Huawei.AppID = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App ID", err.Error())
}

func TestMissingAppSecretForInitHMSClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	client, err := InitHMSClient(cfg, "", "APP_SECRET")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App Secret", err.Error())
}

func TestMissingAppIDForInitHMSClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	client, err := InitHMSClient(cfg, "APP_ID", "")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "Missing Huawei App ID", err.Error())
}
