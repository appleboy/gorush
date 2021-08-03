package notify

import (
	"os"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"

	"github.com/appleboy/go-fcm"
	"github.com/stretchr/testify/assert"
)

func TestMissingAndroidAPIKey(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.APIKey = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "Missing Android API Key", err.Error())
}

func TestMissingKeyForInitFCMClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	cfg.Android.APIKey = ""
	client, err := InitFCMClient(cfg, "")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "Missing Android API Key", err.Error())
}

func TestPushToAndroidWrongToken(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.APIKey = os.Getenv("ANDROID_API_KEY")

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
	cfg.Android.APIKey = os.Getenv("ANDROID_API_KEY")
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
	cfg.Android.APIKey = os.Getenv("ANDROID_API_KEY")

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

func TestOverwriteAndroidAPIKey(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Core.Sync = true
	cfg.Android.Enabled = true
	cfg.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := &PushNotification{
		Tokens:   []string{androidToken, "bbbbb"},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
		// overwrite android api key
		APIKey: "1234",
	}

	// FCM server error: 401 error: 401 Unauthorized (Wrong API Key)
	resp, err := PushToAndroid(req, cfg)

	assert.Error(t, err)
	assert.Len(t, resp.Logs, 2)
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

	// the message's TimeToLive field must be an integer
	// between 0 and 2419200 (4 weeks)
	timeToLive := uint(2419201)
	req = &PushNotification{
		Message:    "Test",
		Platform:   core.PlatFormAndroid,
		Tokens:     []string{"XXXXXXXXX"},
		TimeToLive: &timeToLive,
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// Pass
	timeToLive = uint(86400)
	req = &PushNotification{
		Message:    "Test",
		Platform:   core.PlatFormAndroid,
		Tokens:     []string{"XXXXXXXXX"},
		TimeToLive: &timeToLive,
	}

	err = CheckMessage(req)
	assert.NoError(t, err)
}

func TestCheckAndroidMessage(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	timeToLive := uint(2419201)
	req := &PushNotification{
		Tokens:     []string{"aaaaaa", "bbbbb"},
		Platform:   core.PlatFormAndroid,
		Message:    "Welcome",
		TimeToLive: &timeToLive,
	}

	// the message's TimeToLive field must be an integer between 0 and 2419200 (4 weeks)
	resp, err := PushToAndroid(req, cfg)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestAndroidNotificationStructure(t *testing.T) {
	test := "test"
	timeToLive := uint(100)
	req := &PushNotification{
		Tokens:                []string{"a", "b"},
		Message:               "Welcome",
		To:                    test,
		Priority:              "high",
		CollapseKey:           "1",
		ContentAvailable:      true,
		DelayWhileIdle:        true,
		TimeToLive:            &timeToLive,
		RestrictedPackageName: test,
		DryRun:                true,
		Title:                 test,
		Sound:                 test,
		Data: D{
			"a": "1",
			"b": 2,
		},
		Notification: &fcm.Notification{
			Color: test,
			Tag:   test,
			Body:  "",
		},
	}

	notification := GetAndroidNotification(req)

	assert.Equal(t, test, notification.To)
	assert.Equal(t, "high", notification.Priority)
	assert.Equal(t, "1", notification.CollapseKey)
	assert.True(t, notification.ContentAvailable)
	assert.True(t, notification.DelayWhileIdle)
	assert.Equal(t, uint(100), *notification.TimeToLive)
	assert.Equal(t, test, notification.RestrictedPackageName)
	assert.True(t, notification.DryRun)
	assert.Equal(t, test, notification.Notification.Title)
	assert.Equal(t, test, notification.Notification.Sound)
	assert.Equal(t, test, notification.Notification.Color)
	assert.Equal(t, test, notification.Notification.Tag)
	assert.Equal(t, "Welcome", notification.Notification.Body)
	assert.Equal(t, "1", notification.Data["a"])
	assert.Equal(t, 2, notification.Data["b"])

	// test empty body
	req = &PushNotification{
		Tokens: []string{"a", "b"},
		To:     test,
		Notification: &fcm.Notification{
			Body: "",
		},
	}
	notification = GetAndroidNotification(req)

	assert.Equal(t, test, notification.To)
	assert.Equal(t, "", notification.Notification.Body)
}
