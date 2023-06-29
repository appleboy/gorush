package notify

import (
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/stretchr/testify/assert"
)

func TestMissingHuaweiAppSecret(t *testing.T) {
	cfg, _ := config.LoadConf()
	tenantId := "tenant_id1"
	tenant := cfg.Tenants[tenantId]

	tenant.Huawei.Enabled = true
	tenant.Huawei.APIKey = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "missing huawei app secret", err.Error())
}

func TestMissingHuaweiAppID(t *testing.T) {
	cfg, _ := config.LoadConf()
	tenantId := "tenant_id1"
	tenant := cfg.Tenants[tenantId]

	tenant.Huawei.Enabled = true
	tenant.Huawei.APPId = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "missing huawei app id", err.Error())
}

func TestMissingAppSecretForInitHMSClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	client, err := InitHMSClient(cfg, "", "APP_SECRET", "")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "missing huawei app secret", err.Error())
}

func TestMissingAppIDForInitHMSClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	client, err := InitHMSClient(cfg, "tenant_id1", "APP_ID", "")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "missing huawei app id", err.Error())
}
