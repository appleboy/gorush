package gopush

import (
	"github.com/stretchr/testify/assert"
	"github.com/Sirupsen/logrus"
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	log := logrus.New()

	err := SetLogLevel(log, "debug")
	assert.Nil(t, err)

	err = SetLogLevel(log, "invalid")
	assert.Equal(t, "not a valid logrus Level: \"invalid\"", err.Error())
}

func TestSetLogOut(t *testing.T) {
	log := logrus.New()

	err := SetLogOut(log, "stdout")
	assert.Nil(t, err)

	err = SetLogOut(log, "stderr")
	assert.Nil(t, err)

	err = SetLogOut(log, "access.log")
	assert.Nil(t, err)

	// missing create logs folder.
	err = SetLogOut(log, "logs/access.log")
	assert.NotNil(t, err)
}

func TestInitDefaultLog(t *testing.T) {
	PushConf = BuildDefaultPushConf()

	// no errors on default config
	assert.Nil(t, InitLog())

	PushConf.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog())
}

func TestAccessLevel(t *testing.T) {
	PushConf = BuildDefaultPushConf()

	PushConf.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog())
}

func TestErrorLevel(t *testing.T) {
	PushConf = BuildDefaultPushConf()

	PushConf.Log.ErrorLevel = "invalid"

	assert.NotNil(t, InitLog())
}

func TestAccessLogPath(t *testing.T) {
	PushConf = BuildDefaultPushConf()

	PushConf.Log.AccessLog = "logs/access.log"

	assert.NotNil(t, InitLog())
}

func TestErrorLogPath(t *testing.T) {
	PushConf = BuildDefaultPushConf()

	PushConf.Log.ErrorLog = "logs/error.log"

	assert.NotNil(t, InitLog())
}
