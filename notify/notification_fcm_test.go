package notify

import (
	"os"
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
)

func TestMissingAndroidCredential(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "missing fcm credential data", err.Error())
}

func TestMissingKeyForInitFCMClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	cfg.Android.Credential = ""
	cfg.Android.KeyPath = ""
	client, err := InitFCMClient(cfg)

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "missing fcm credential data", err.Error())
}

func TestPushToAndroidWrongToken(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = os.Getenv("ANDROID_API_KEY")

	req := &PushNotification{
		Tokens:   []string{"aaaaaa", "bbbbb"},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
	}

	// Android Success count: 0, Failure count: 2
	resp, err := PushToAndroid(req, cfg)
	assert.Nil(t, err)
	assert.Len(t, resp.Logs, 2)
}

func TestPushToAndroidRightTokenForJSONLog(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = os.Getenv("ANDROID_API_KEY")
	// log for json
	cfg.Log.Format = "json"

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := &PushNotification{
		Tokens:   []string{androidToken},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
	}

	resp, err := PushToAndroid(req, cfg)
	assert.Nil(t, err)
	assert.Len(t, resp.Logs, 0)
}

func TestPushToAndroidRightTokenForStringLog(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := &PushNotification{
		Tokens:   []string{androidToken},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
	}

	resp, err := PushToAndroid(req, cfg)
	assert.Nil(t, err)
	assert.Len(t, resp.Logs, 0)
}

func TestFCMMessage(t *testing.T) {
	var err error

	// the message must specify at least one registration ID
	req := &PushNotification{
		Message: "Test",
		Tokens:  []string{},
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// the token must not be empty
	req = &PushNotification{
		Message: "Test",
		Tokens:  []string{""},
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// ignore check token length if send topic message
	req = &PushNotification{
		Message:  "Test",
		Platform: core.PlatFormAndroid,
		To:       "/topics/foo-bar",
	}

	err = CheckMessage(req)
	assert.NoError(t, err)

	// "condition": "'dogs' in topics || 'cats' in topics",
	req = &PushNotification{
		Message:   "Test",
		Platform:  core.PlatFormAndroid,
		Condition: "'dogs' in topics || 'cats' in topics",
	}

	err = CheckMessage(req)
	assert.NoError(t, err)

	// the message may specify at most 1000 registration IDs
	req = &PushNotification{
		Message:  "Test",
		Platform: core.PlatFormAndroid,
		Tokens:   make([]string, 1001),
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	req = &PushNotification{
		Message:  "Test",
		Platform: core.PlatFormAndroid,
		Tokens:   []string{"XXXXXXXXX"},
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// Pass
	req = &PushNotification{
		Message:  "Test",
		Platform: core.PlatFormAndroid,
		Tokens:   []string{"XXXXXXXXX"},
	}

	err = CheckMessage(req)
	assert.NoError(t, err)
}

func TestCheckAndroidMessage(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = os.Getenv("ANDROID_API_KEY")

	req := &PushNotification{
		Tokens:   []string{"aaaaaa", "bbbbb"},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
	}

	// the message's TimeToLive field must be an integer between 0 and 2419200 (4 weeks)
	resp, err := PushToAndroid(req, cfg)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestAndroidNotificationStructure(t *testing.T) {
	test := "test"
	req := &PushNotification{
		Tokens:           []string{"a", "b"},
		Message:          "Welcome",
		To:               test,
		Priority:         HIGH,
		ContentAvailable: true,
		Title:            test,
		Sound:            test,
		Data: D{
			"a": "1",
			"b": 2,
		},
		Notification: &messaging.Notification{
			Title: test,
			Body:  "",
		},
	}

	notification := GetAndroidNotification(req)

	assert.Equal(t, test, notification.Notification.Title)
	assert.Equal(t, "Welcome", notification.Notification.Body)
	assert.Equal(t, "1", notification.Data["a"])
	assert.Equal(t, 2, notification.Data["b"])

	// test empty body
	req = &PushNotification{
		Tokens: []string{"a", "b"},
		To:     test,
		Notification: &messaging.Notification{
			Body: "",
		},
	}
	notification = GetAndroidNotification(req)

	assert.Equal(t, "", notification.Notification.Body)
}
