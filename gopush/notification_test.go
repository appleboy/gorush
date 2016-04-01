package gopush

import (
	"github.com/stretchr/testify/assert"
	"github.com/buger/jsonparser"
	"encoding/json"
	"testing"
	"log"
)

func TestIOSNotificationStructure(t *testing.T) {
	var dat map[string]interface{}

	test := "test"
	message := "Welcome notification Server"
	req := RequestPushNotification{
		ApnsID: test,
		Topic: test,
		Priority: "normal",
		Message: message,
		Badge: 1,
		Sound: test,
		ContentAvailable: true,
		Extend: []ExtendJSON{
			{
				Key: "key1",
				Value: "1",
			},
			{
				Key: "key2",
				Value: "2",
			},
		},
		Category: test,
		URLArgs: []string{"a", "b"},
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
	key1 := dat["key1"].(string)
	key2 := dat["key2"].(string)
	aps := dat["aps"].(map[string]interface{})
	urlArgs := aps["url-args"].([]interface{})

	assert.Equal(t, test, notification.ApnsID)
	assert.Equal(t, test, notification.Topic)
	assert.Equal(t, ApnsPriorityLow, notification.Priority)
	assert.Equal(t, message, alert)
	assert.Equal(t, 1, int(badge))
	assert.Equal(t, test, sound)
	assert.Equal(t, 1, int(contentAvailable))
	assert.Equal(t, "1", key1)
	assert.Equal(t, "2", key2)
	assert.Equal(t, test, category)
	assert.Contains(t, urlArgs, "a")
	assert.Contains(t, urlArgs, "b")
}

func TestIOSAlertNotificationStructure(t *testing.T) {
	var dat map[string]interface{}

	test := "test"
	req := RequestPushNotification{
		Alert: Alert{
			Action: test,
			ActionLocKey: test,
			Body: test,
			LaunchImage: test,
			LocArgs: []string{"a", "b"},
			LocKey: test,
			Title: test,
			TitleLocArgs: []string{"a", "b"},
			TitleLocKey: test,
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
	actionLocKey, _ := jsonparser.GetString(data, "aps", "alert","action-loc-key")
	body, _ := jsonparser.GetString(data, "aps", "alert","body")
	launchImage, _ := jsonparser.GetString(data, "aps", "alert","launch-image")
	locKey, _ := jsonparser.GetString(data, "aps", "alert","loc-key")
	title, _ := jsonparser.GetString(data, "aps", "alert","title")
	titleLocKey, _ := jsonparser.GetString(data, "aps", "alert","title-loc-key")
	aps := dat["aps"].(map[string]interface{})
	alert := aps["alert"].(map[string]interface{})
	titleLocArgs := alert["title-loc-args"].([]interface{})

	assert.Equal(t, test, action)
	assert.Equal(t, test, actionLocKey)
	assert.Equal(t, test, body)
	assert.Equal(t, test, launchImage)
	assert.Equal(t, test, locKey)
	assert.Equal(t, test, title)
	assert.Equal(t, test, titleLocKey)
	assert.Contains(t, titleLocArgs, "a")
	assert.Contains(t, titleLocArgs, "b")
}
