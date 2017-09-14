package gorush

import (
	"os"
	"testing"

	"github.com/axiomzen/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestCorrectConf(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = "xxxxx"

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"

	err := CheckPushConf()

	assert.NoError(t, err)
}

func TestSenMultipleNotifications(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	InitWorkers(int64(2), 2)

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	err := InitAPNSClient()
	assert.Nil(t, err)

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := RequestPush{
		Notifications: []PushNotification{
			//ios
			{
				Tokens:   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
				Platform: PlatFormIos,
				Message:  "Welcome",
			},
			// android
			{
				Tokens:   []string{androidToken, "bbbbb"},
				Platform: PlatFormAndroid,
				Message:  "Welcome",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 3, count)
	assert.Equal(t, 0, len(logs))
}

func TestDisabledAndroidNotifications(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	err := InitAPNSClient()
	assert.Nil(t, err)

	PushConf.Android.Enabled = false
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := RequestPush{
		Notifications: []PushNotification{
			//ios
			{
				Tokens:   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
				Platform: PlatFormIos,
				Message:  "Welcome",
			},
			// android
			{
				Tokens:   []string{androidToken, "bbbbb"},
				Platform: PlatFormAndroid,
				Message:  "Welcome",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 1, count)
	assert.Equal(t, 0, len(logs))
}

func TestSyncModeForNotifications(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	err := InitAPNSClient()
	assert.Nil(t, err)

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	// enable sync mode
	PushConf.Core.Sync = true

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := RequestPush{
		Notifications: []PushNotification{
			//ios
			{
				Tokens:   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
				Platform: PlatFormIos,
				Message:  "Welcome",
			},
			// android
			{
				Tokens:   []string{androidToken, "bbbbb"},
				Platform: PlatFormAndroid,
				Message:  "Welcome",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 3, count)
	assert.Equal(t, 2, len(logs))
}

func TestSetProxyURL(t *testing.T) {

	err := SetProxy("87.236.233.92:8080")
	assert.Error(t, err)
	assert.Equal(t, "parse 87.236.233.92:8080: invalid URI for request", err.Error())

	err = SetProxy("a.html")
	assert.Error(t, err)

	err = SetProxy("http://87.236.233.92:8080")
	assert.NoError(t, err)
}
