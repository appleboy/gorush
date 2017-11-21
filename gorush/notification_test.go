package gorush

import (
	"os"
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

func TestSenMultipleNotifications(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	InitWorkers(int64(2), 2)

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	err := InitAPNSClient()
	assert.Nil(t, err)

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	PushConf.Web.Enabled = true
	PushConf.Web.APIKey = os.Getenv("ANDROID_API_KEY")
	err2 := InitWebClient()
	assert.Nil(t, err2)

	req := RequestPush{
		Notifications: []PushNotification{
			//ios
			{
				Tokens:   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
				Platform: PlatformIos,
				Message:  "Welcome",
			},
			// android
			{
				Tokens:   []string{androidToken, "bbbbb"},
				Platform: PlatformAndroid,
				Message:  "Welcome",
			},
			// web
			{
				Subscriptions: []Subscription{{"https://updates.push.services.mozilla.com/wpush/v1/gAAAAABZdwpXbtIhiT_gXZZ_lUrs0AqbMROAwW8-LpTVx_LNYTU-xrcvIoZ7LXNNeTSSO525EYuKCeGueKtSqi626yCOAaFWYAzRu9hOIwvVmJFfN3BIlMjR9PJU28s7JsNVKywp4_wb", "BIqGnYTkDOAyPNTInUdE7AeYnA2LrHNu6jKpfYwcfl3Z8EVyRtftqpHgJTku3YjQBGwqJyzxsYwc9tHNB5jUks8", "j7NMANtEQ5zoFUGtiCdqRQ"}},
				Platform:      PlatformWeb,
				Message:       "Welcome",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 4, count)
	assert.Equal(t, 0, len(logs))
}

func TestDisabledAndroidNotifications(t *testing.T) {
	PushConf, _ = config.LoadConf("")

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
				Platform: PlatformIos,
				Message:  "Welcome",
			},
			// android
			{
				Tokens:   []string{androidToken, "bbbbb"},
				Platform: PlatformAndroid,
				Message:  "Welcome",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 1, count)
	assert.Equal(t, 0, len(logs))
}

func TestSyncModeForNotifications(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	err := InitAPNSClient()
	assert.Nil(t, err)

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	// enable sync mode
	PushConf.Core.Sync = true

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	PushConf.Web.Enabled = true
	PushConf.Web.APIKey = os.Getenv("ANDROID_API_KEY")
	err2 := InitWebClient()
	assert.Nil(t, err2)

	req := RequestPush{
		Notifications: []PushNotification{
			//ios
			{
				Tokens:   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
				Platform: PlatformIos,
				Message:  "Welcome",
			},
			// android
			{
				Tokens:   []string{androidToken, "bbbbb"},
				Platform: PlatformAndroid,
				Message:  "Welcome",
			},
			// web
			{
				Subscriptions: []Subscription{
					{
						"https://updates.push.services.mozilla.com/wpush/v1/gAAAAABZdwpXbtIhiT_gXZZ_lUrs0AqbMROAwW8-LpTVx_LNYTU-xrcvIoZ7LXNNeTSSO525EYuKCeGueKtSqi626yCOAaFWYAzRu9hOIwvVmJFfN3BIlMjR9PJU28s7JsNVKywp4_wb",
						"BIqGnYTkDOAyPNTInUdE7AeYnA2LrHNu6jKpfYwcfl3Z8EVyRtftqpHgJTku3YjQBGwqJyzxsYwc9tHNB5jUks8",
						"j7NMANtEQ5zoFUGtiCdqRQ",
					},
					{
						"https://updates.push.services.mozilla.com/wpush/v1/g",
						"BIqGnYTkDOAyPNTInUdE7AeYnA2LrHNu6jKpfYwcfl3Z8EVyRtftqpHgJTku3YjQBGwqJyzxsYwc9tHNB5jUks8",
						"j7NMANtEQ5zoFUGtiCdqRQ",
					},
					{
						"aaaaa",
						"bbbbb",
						"ccccc",
					},
				},
				Platform: PlatformWeb,
				Message:  "Welcome",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 6, count)
	assert.Equal(t, 4, len(logs))
}

func TestSyncModeForTopicNotification(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")
	PushConf.Log.HideToken = false

	// enable sync mode
	PushConf.Core.Sync = true

	req := RequestPush{
		Notifications: []PushNotification{
			// android
			{
				// error:InvalidParameters
				// Check that the provided parameters have the right name and type.
				To:       "/topics/foo-bar@@@##",
				Platform: PlatformAndroid,
				Message:  "This is a Firebase Cloud Messaging Topic Message!",
			},
			// android
			{
				// success
				To:       "/topics/foo-bar",
				Platform: PlatformAndroid,
				Message:  "This is a Firebase Cloud Messaging Topic Message!",
			},
			// android
			{
				// success
				Condition: "'dogs' in topics || 'cats' in topics",
				Platform:  PlatformAndroid,
				Message:   "This is a Firebase Cloud Messaging Topic Message!",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 2, count)
	assert.Equal(t, 1, len(logs))
}

func TestSyncModeForDeviceGroupNotification(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")
	PushConf.Log.HideToken = false

	// enable sync mode
	PushConf.Core.Sync = true

	req := RequestPush{
		Notifications: []PushNotification{
			// android
			{
				To:       "aUniqueKey",
				Platform: PlatformAndroid,
				Message:  "This is a Firebase Cloud Messaging Device Group Message!",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, len(logs))
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
