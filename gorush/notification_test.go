package gorush

import (
	"encoding/json"
	"github.com/appleboy/gorush/config"
	"github.com/buger/jsonparser"
	"github.com/google/go-gcm"
	"github.com/sideshow/apns2"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestDisabledAndroidIosConf(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	err := CheckPushConf()

	assert.Error(t, err)
	assert.Equal(t, "Please enable iOS or Android config in yml config", err.Error())
}

func TestMissingIOSCertificate(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = ""

	err := CheckPushConf()

	assert.Error(t, err)
	assert.Equal(t, "Missing iOS certificate path", err.Error())
}

func TestMissingAndroidAPIKey(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = ""

	err := CheckPushConf()

	assert.Error(t, err)
	assert.Equal(t, "Missing Android API Key", err.Error())
}

func TestCorrectConf(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = "xxxxx"

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "xxxxx"

	err := CheckPushConf()

	assert.NoError(t, err)
}

func TestIOSNotificationStructure(t *testing.T) {
	var dat map[string]interface{}
	var unix = time.Now().Unix()

	test := "test"
	message := "Welcome notification Server"
	req := PushNotification{
		ApnsID:           test,
		Topic:            test,
		Expiration:       time.Now().Unix(),
		Priority:         "normal",
		Message:          message,
		Badge:            1,
		Sound:            test,
		ContentAvailable: true,
		Data: D{
			"key1": "test",
			"key2": 2,
		},
		Category: test,
		URLArgs:  []string{"a", "b"},
	}

	notification := GetIOSNotification(req)

	dump, _ := json.Marshal(notification.Payload)
	data := []byte(string(dump))

	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println(err)
		panic(err)
	}

	alert, _ := jsonparser.GetString(data, "aps", "alert")
	badge, _ := jsonparser.GetInt(data, "aps", "badge")
	sound, _ := jsonparser.GetString(data, "aps", "sound")
	contentAvailable, _ := jsonparser.GetInt(data, "aps", "content-available")
	category, _ := jsonparser.GetString(data, "aps", "category")
	key1 := dat["key1"].(interface{})
	key2 := dat["key2"].(interface{})
	aps := dat["aps"].(map[string]interface{})
	urlArgs := aps["url-args"].([]interface{})

	assert.Equal(t, test, notification.ApnsID)
	assert.Equal(t, test, notification.Topic)
	assert.Equal(t, unix, notification.Expiration.Unix())
	assert.Equal(t, ApnsPriorityLow, notification.Priority)
	assert.Equal(t, message, alert)
	assert.Equal(t, 1, int(badge))
	assert.Equal(t, test, sound)
	assert.Equal(t, 1, int(contentAvailable))
	assert.Equal(t, "test", key1)
	assert.Equal(t, 2, int(key2.(float64)))
	assert.Equal(t, test, category)
	assert.Contains(t, urlArgs, "a")
	assert.Contains(t, urlArgs, "b")
}

func TestIOSAlertNotificationStructure(t *testing.T) {
	var dat map[string]interface{}

	test := "test"
	req := PushNotification{
		Message: "Welcome",
		Title:   test,
		Alert: Alert{
			Action:       test,
			ActionLocKey: test,
			Body:         test,
			LaunchImage:  test,
			LocArgs:      []string{"a", "b"},
			LocKey:       test,
			Subtitle:     test,
			TitleLocArgs: []string{"a", "b"},
			TitleLocKey:  test,
		},
	}

	notification := GetIOSNotification(req)

	dump, _ := json.Marshal(notification.Payload)
	data := []byte(string(dump))

	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println(err)
		panic(err)
	}

	action, _ := jsonparser.GetString(data, "aps", "alert", "action")
	actionLocKey, _ := jsonparser.GetString(data, "aps", "alert", "action-loc-key")
	body, _ := jsonparser.GetString(data, "aps", "alert", "body")
	launchImage, _ := jsonparser.GetString(data, "aps", "alert", "launch-image")
	locKey, _ := jsonparser.GetString(data, "aps", "alert", "loc-key")
	title, _ := jsonparser.GetString(data, "aps", "alert", "title")
	subtitle, _ := jsonparser.GetString(data, "aps", "alert", "subtitle")
	titleLocKey, _ := jsonparser.GetString(data, "aps", "alert", "title-loc-key")
	aps := dat["aps"].(map[string]interface{})
	alert := aps["alert"].(map[string]interface{})
	titleLocArgs := alert["title-loc-args"].([]interface{})
	locArgs := alert["loc-args"].([]interface{})

	assert.Equal(t, test, action)
	assert.Equal(t, test, actionLocKey)
	assert.Equal(t, test, body)
	assert.Equal(t, test, launchImage)
	assert.Equal(t, test, locKey)
	assert.Equal(t, test, title)
	assert.Equal(t, test, subtitle)
	assert.Equal(t, test, titleLocKey)
	assert.Contains(t, titleLocArgs, "a")
	assert.Contains(t, titleLocArgs, "b")
	assert.Contains(t, locArgs, "a")
	assert.Contains(t, locArgs, "b")
}

func TestAndroidNotificationStructure(t *testing.T) {

	test := "test"
	timeToLive := uint(100)
	req := PushNotification{
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
		Notification: gcm.Notification{
			Color: test,
			Tag:   test,
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
}

func TestPushToIOS(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	InitAPNSClient()
	InitAppStatus()

	req := PushNotification{
		Tokens:   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
		Platform: 1,
		Message:  "Welcome",
	}

	isError := PushToIOS(req)
	assert.True(t, isError)
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

	success := PushToAndroid(req)
	assert.False(t, success)
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

	success := PushToAndroid(req)
	assert.True(t, success)
}

func TestPushToAndroidRightTokenForJSONLog(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")
	// log for json
	PushConf.Log.Format = "json"

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := PushNotification{
		Tokens:   []string{androidToken, "bbbbb"},
		Platform: PlatFormAndroid,
		Message:  "Welcome",
	}

	success := PushToAndroid(req)
	assert.True(t, success)
}

func TestPushToAndroidRightTokenForStringLog(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	req := PushNotification{
		Tokens:   []string{androidToken, "bbbbb"},
		Platform: PlatFormAndroid,
		Message:  "Welcome",
	}

	success := PushToAndroid(req)
	assert.True(t, success)
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

	success := PushToAndroid(req)
	assert.False(t, success)
}

func TestSenMultipleNotifications(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	InitWorkers(int64(2), 2)

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	InitAPNSClient()

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

	count := queueNotification(req)
	assert.Equal(t, 3, count)
}

func TestDisabledAndroidNotifications(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	InitAPNSClient()

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

	count := queueNotification(req)
	assert.Equal(t, 1, count)
}

func TestDisabledIosNotifications(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = false
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	InitAPNSClient()

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

	count := queueNotification(req)
	assert.Equal(t, 2, count)
}

func TestWrongIosCertificateExt(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "test"
	err := InitAPNSClient()

	assert.Error(t, err)
	assert.Equal(t, "Wrong Certificate key extension.", err.Error())
}

func TestAPNSClientDevHost(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.p12"
	InitAPNSClient()

	assert.Equal(t, apns2.HostDevelopment, ApnsClient.Host)
}

func TestAPNSClientProdHost(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.Production = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	InitAPNSClient()

	assert.Equal(t, apns2.HostProduction, ApnsClient.Host)
}

func TestGCMMessage(t *testing.T) {
	var req PushNotification
	var err error

	// the message must not be empty
	req = PushNotification{
		Message: "",
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// the message must specify at least one registration ID
	req = PushNotification{
		Message: "Test",
		Tokens:  []string{},
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// the token must not be empty
	req = PushNotification{
		Message: "Test",
		Tokens:  []string{""},
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// the message may specify at most 1000 registration IDs
	req = PushNotification{
		Message:  "Test",
		Platform: PlatFormAndroid,
		Tokens:   make([]string, 1001),
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// the message's TimeToLive field must be an integer
	// between 0 and 2419200 (4 weeks)
	timeToLive := uint(2419201)
	req = PushNotification{
		Message:    "Test",
		Platform:   PlatFormAndroid,
		Tokens:     []string{"XXXXXXXXX"},
		TimeToLive: &timeToLive,
	}

	err = CheckMessage(req)
	assert.Error(t, err)

	// Pass
	timeToLive = uint(86400)
	req = PushNotification{
		Message:    "Test",
		Platform:   PlatFormAndroid,
		Tokens:     []string{"XXXXXXXXX"},
		TimeToLive: &timeToLive,
	}

	err = CheckMessage(req)
	assert.NoError(t, err)
}

func TestCheckAndroidMessage(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	timeToLive := uint(2419201)
	req := PushNotification{
		Tokens:     []string{"aaaaaa", "bbbbb"},
		Platform:   PlatFormAndroid,
		Message:    "Welcome",
		TimeToLive: &timeToLive,
	}

	success := PushToAndroid(req)
	assert.False(t, success)
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
