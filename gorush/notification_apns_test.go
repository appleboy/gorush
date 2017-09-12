package gorush

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"github.com/axiomzen/gorush/config"
	"github.com/buger/jsonparser"
	"github.com/sideshow/apns2"
	"github.com/stretchr/testify/assert"
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

	PushConf.Ios.KeyPath = "test.pem"
	err = CheckPushConf()

	assert.Error(t, err)
	assert.Equal(t, "certificate file does not exist", err.Error())

}
func TestIOSNotificationStructure(t *testing.T) {
	var dat map[string]interface{}
	var unix = time.Now().Unix()

	test := "test"
	expectBadge := 0
	message := "Welcome notification Server"
	req := PushNotification{
		ApnsID:           test,
		Topic:            test,
		Expiration:       time.Now().Unix(),
		Priority:         "normal",
		Message:          message,
		Badge:            &expectBadge,
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
	assert.Equal(t, expectBadge, int(badge))
	assert.Equal(t, expectBadge, *req.Badge)
	assert.Equal(t, test, sound)
	assert.Equal(t, 1, int(contentAvailable))
	assert.Equal(t, "test", key1)
	assert.Equal(t, 2, int(key2.(float64)))
	assert.Equal(t, test, category)
	assert.Contains(t, urlArgs, "a")
	assert.Contains(t, urlArgs, "b")
}

// Silent Notification which payload’s aps dictionary must not contain the alert, sound, or badge keys.
// ref: https://goo.gl/m9xyqG
func TestSendZeroValueForBadgeKey(t *testing.T) {
	var dat map[string]interface{}

	test := "test"
	message := "Welcome notification Server"
	req := PushNotification{
		ApnsID:           test,
		Topic:            test,
		Priority:         "normal",
		Message:          message,
		Sound:            test,
		ContentAvailable: true,
		MutableContent:   true,
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
	mutableContent, _ := jsonparser.GetInt(data, "aps", "mutable-content")

	if req.Badge != nil {
		t.Errorf("req.Badge must be nil")
	}

	assert.Equal(t, test, notification.ApnsID)
	assert.Equal(t, test, notification.Topic)
	assert.Equal(t, ApnsPriorityLow, notification.Priority)
	assert.Equal(t, message, alert)
	assert.Equal(t, 0, int(badge))
	assert.Equal(t, test, sound)
	assert.Equal(t, 1, int(contentAvailable))
	assert.Equal(t, 1, int(mutableContent))

	// Add Bage
	expectBadge := 10
	req.Badge = &expectBadge

	notification = GetIOSNotification(req)

	dump, _ = json.Marshal(notification.Payload)
	data = []byte(string(dump))

	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println(err)
		panic(err)
	}

	if req.Badge == nil {
		t.Errorf("req.Badge must be equal %d", *req.Badge)
	}

	badge, _ = jsonparser.GetInt(data, "aps", "badge")
	assert.Equal(t, expectBadge, *req.Badge)
	assert.Equal(t, expectBadge, int(badge))
}

// Silent Notification:
// The payload’s aps dictionary must include the content-available key with a value of 1.
// The payload’s aps dictionary must not contain the alert, sound, or badge keys.
// ref: https://goo.gl/m9xyqG
func TestCheckSilentNotification(t *testing.T) {
	var dat map[string]interface{}

	test := "test"
	req := PushNotification{
		ApnsID:           test,
		Topic:            test,
		Priority:         "normal",
		ContentAvailable: true,
	}

	notification := GetIOSNotification(req)

	dump, _ := json.Marshal(notification.Payload)
	data := []byte(string(dump))

	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println(err)
		panic(err)
	}

	assert.Nil(t, dat["aps"].(map[string]interface{})["alert"])
	assert.Nil(t, dat["aps"].(map[string]interface{})["sound"])
	assert.Nil(t, dat["aps"].(map[string]interface{})["badge"])
}

// URL: https://goo.gl/5xFo3C
// Example 2
// {
//     "aps" : {
//         "alert" : {
//             "title" : "Game Request",
//             "body" : "Bob wants to play poker",
//             "action-loc-key" : "PLAY"
//         },
//         "badge" : 5
//     },
//     "acme1" : "bar",
//     "acme2" : [ "bang",  "whiz" ]
// }
func TestAlertStringExample2ForIos(t *testing.T) {
	var dat map[string]interface{}

	test := "test"
	title := "Game Request"
	body := "Bob wants to play poker"
	actionLocKey := "PLAY"
	req := PushNotification{
		ApnsID:   test,
		Topic:    test,
		Priority: "normal",
		Alert: Alert{
			Title:        title,
			Body:         body,
			ActionLocKey: actionLocKey,
		},
	}

	notification := GetIOSNotification(req)

	dump, _ := json.Marshal(notification.Payload)
	data := []byte(string(dump))

	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println(err)
		panic(err)
	}

	assert.Equal(t, title, dat["aps"].(map[string]interface{})["alert"].(map[string]interface{})["title"])
	assert.Equal(t, body, dat["aps"].(map[string]interface{})["alert"].(map[string]interface{})["body"])
	assert.Equal(t, actionLocKey, dat["aps"].(map[string]interface{})["alert"].(map[string]interface{})["action-loc-key"])
}

// URL: https://goo.gl/5xFo3C
// Example 3
// {
//     "aps" : {
//         "alert" : "You got your emails.",
//         "badge" : 9,
//         "sound" : "bingbong.aiff"
//     },
//     "acme1" : "bar",
//     "acme2" : 42
// }
func TestAlertStringExample3ForIos(t *testing.T) {
	var dat map[string]interface{}

	test := "test"
	badge := 9
	sound := "bingbong.aiff"
	req := PushNotification{
		ApnsID:           test,
		Topic:            test,
		Priority:         "normal",
		ContentAvailable: true,
		Message:          test,
		Badge:            &badge,
		Sound:            sound,
	}

	notification := GetIOSNotification(req)

	dump, _ := json.Marshal(notification.Payload)
	data := []byte(string(dump))

	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println(err)
		panic(err)
	}

	assert.Equal(t, sound, dat["aps"].(map[string]interface{})["sound"])
	assert.Equal(t, float64(badge), dat["aps"].(map[string]interface{})["badge"].(float64))
	assert.Equal(t, test, dat["aps"].(map[string]interface{})["alert"])
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

func TestDisabledIosNotifications(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = false
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
				Tokens:   []string{androidToken, androidToken + "_"},
				Platform: PlatFormAndroid,
				Message:  "Welcome",
			},
		},
	}

	count, logs := queueNotification(req)
	assert.Equal(t, 2, count)
	assert.Equal(t, 0, len(logs))
}

func TestWrongIosCertificateExt(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "test"
	err := InitAPNSClient()

	assert.Error(t, err)
	assert.Equal(t, "wrong certificate key extension", err.Error())
}

func TestAPNSClientDevHost(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.p12"
	err := InitAPNSClient()
	assert.Nil(t, err)
	assert.Equal(t, apns2.HostDevelopment, ApnsClient.Host)
}

func TestAPNSClientProdHost(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.Production = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	err := InitAPNSClient()
	assert.Nil(t, err)
	assert.Equal(t, apns2.HostProduction, ApnsClient.Host)
}

func TestPushToIOS(t *testing.T) {
	PushConf = config.BuildDefaultPushConf()

	PushConf.Ios.Enabled = true
	PushConf.Ios.KeyPath = "../certificate/certificate-valid.pem"
	err := InitAPNSClient()
	assert.Nil(t, err)
	err = InitAppStatus()
	assert.Nil(t, err)

	req := PushNotification{
		Tokens:   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
		Platform: 1,
		Message:  "Welcome",
	}

	// send fail
	isError := PushToIOS(req)
	assert.True(t, isError)
}
