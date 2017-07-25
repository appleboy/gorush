package gorush

import (
	"log"
	"os"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func init() {
	PushConf = config.BuildDefaultPushConf()
	if err := InitLog(); err != nil {
		log.Fatal(err)
	}

	if err := InitAppStatus(); err != nil {
		log.Fatal(err)
	}
}

func TestMissingKeyForInitFCMClient(t *testing.T) {
	client, err := InitFCMClient("")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "Missing Android API Key", err.Error())
}

func TestPushToAndroidWrongAPIKey(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY") + "a"

	req := PushNotification{
		Tokens:   []string{"aaaaaa", "bbbbb"},
		Platform: PlatFormAndroid,
		Message:  "Welcome",
	}

	// FCM server send message error: 401 error: 401 Unauthorized
	err := PushToAndroid(req)
	assert.False(t, err)
}

func TestPushToAndroidWrongToken(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	req := PushNotification{
		Tokens:   []string{"aaaaaa", "bbbbb"},
		Platform: PlatFormAndroid,
		Message:  "Welcome",
	}

	// Android Success count: 0, Failure count: 2
	isError := PushToAndroid(req)
	assert.True(t, isError)
}

func TestPushToAndroidRightTokenForJSONLog(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")
	// log for json
	PushConf.Log.Format = "json"

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := PushNotification{
		Tokens:   []string{androidToken},
		Platform: PlatFormAndroid,
		Message:  "Welcome",
	}

	isError := PushToAndroid(req)
	assert.False(t, isError)
}

func TestPushToAndroidRightTokenForStringLog(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := PushNotification{
		Tokens:   []string{androidToken},
		Platform: PlatFormAndroid,
		Message:  "Welcome",
	}

	isError := PushToAndroid(req)
	assert.False(t, isError)
}

func TestOverwriteAndroidAPIKey(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := PushNotification{
		Tokens:   []string{androidToken, "bbbbb"},
		Platform: PlatFormAndroid,
		Message:  "Welcome",
		// overwrite android api key
		APIKey: "1234",
	}

	// FCM server error: 401 error: 401 Unauthorized (Wrong API Key)
	err := PushToAndroid(req)
	assert.False(t, err)
}
