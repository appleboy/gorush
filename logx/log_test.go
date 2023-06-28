package logx

import (
	"errors"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var invalidLevel = "invalid"

func TestSetLogLevel(t *testing.T) {
	log := logrus.New()

	err := SetLogLevel(log, "debug")
	assert.Nil(t, err)

	err = SetLogLevel(log, invalidLevel)
	assert.Equal(t, "not a valid logrus Level: \"invalid\"", err.Error())
}

func TestSetLogOut(t *testing.T) {
	log := logrus.New()

	err := SetLogOut(log, "stdout")
	assert.Nil(t, err)

	err = SetLogOut(log, "stderr")
	assert.Nil(t, err)

	err = SetLogOut(log, "log/access.log")
	assert.Nil(t, err)

	// missing create logs folder.
	err = SetLogOut(log, "logs/access.log")
	assert.NotNil(t, err)
}

func TestInitDefaultLog(t *testing.T) {
	cfg, _ := config.LoadConf()

	// no errors on default config
	assert.Nil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))

	cfg.Log.AccessLevel = invalidLevel

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))

	isTerm = true

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestAccessLevel(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.AccessLevel = invalidLevel

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestErrorLevel(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.ErrorLevel = invalidLevel

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestAccessLogPath(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.AccessLog = "logs/access.log"

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestErrorLogPath(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Log.ErrorLog = "logs/error.log"

	assert.NotNil(t, InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	))
}

func TestPlatFormType(t *testing.T) {
	assert.Equal(t, "ios", typeForPlatForm(core.PlatFormIos))
	assert.Equal(t, "android", typeForPlatForm(core.PlatFormAndroid))
	assert.Equal(t, "huawei", typeForPlatForm(core.PlatFormHuawei))
	assert.Equal(t, "", typeForPlatForm(10000))
}

func TestPlatFormColor(t *testing.T) {
	assert.Equal(t, blue, colorForPlatForm(core.PlatFormIos))
	assert.Equal(t, yellow, colorForPlatForm(core.PlatFormAndroid))
	assert.Equal(t, green, colorForPlatForm(core.PlatFormHuawei))
	assert.Equal(t, reset, colorForPlatForm(1000000))
}

func TestHideToken(t *testing.T) {
	assert.Equal(t, "", hideToken("", 2))
	assert.Equal(t, "**345678**", hideToken("1234567890", 2))
	assert.Equal(t, "*****", hideToken("12345", 10))
}

func TestLogPushEntry(t *testing.T) {
	in := InputLog{}

	in.Platform = 1
	assert.Equal(t, "ios", GetLogPushEntry(&in).Platform)

	in.Error = errors.New("error")
	assert.Equal(t, "error", GetLogPushEntry(&in).Error)

	in.Token = "1234567890"
	in.HideToken = true
	assert.Equal(t, "**********", GetLogPushEntry(&in).Token)

	in.Message = "hellothisisamessage"
	in.HideMessage = true
	assert.Equal(t, "(message redacted)", GetLogPushEntry(&in).Message)
}

func TestLogPush(t *testing.T) {
	in := InputLog{}
	isTerm = true

	in.Format = "json"
	in.Status = "succeeded-push"
	assert.Equal(t, "succeeded-push", LogPush(&in).Type)

	in.Format = ""
	in.Message = "success"
	assert.Equal(t, "success", LogPush(&in).Message)

	in.Status = "failed-push"
	in.Message = "failed"
	assert.Equal(t, "failed", LogPush(&in).Message)
}
