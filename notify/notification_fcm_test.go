package notify

import (
	"context"
	"os"
	"reflect"
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
	client, err := InitFCMClient(context.Background(), cfg)

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "missing fcm credential data", err.Error())
}

func TestPushToAndroidWrongToken(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = os.Getenv("FCM_CREDENTIAL")

	req := &PushNotification{
		Tokens:   []string{"aaaaaa", "bbbbb"},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
	}

	// Android Success count: 0, Failure count: 2
	resp, err := PushToAndroid(context.Background(), req, cfg)
	assert.Nil(t, err)
	assert.Len(t, resp.Logs, 2)
}

func TestPushToAndroidRightTokenForJSONLog(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = os.Getenv("FCM_CREDENTIAL")
	// log for json
	cfg.Log.Format = "json"

	androidToken := os.Getenv("FCM_TEST_TOKEN")

	req := &PushNotification{
		Tokens:   []string{androidToken},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
	}

	resp, err := PushToAndroid(context.Background(), req, cfg)
	assert.Nil(t, err)
	assert.Len(t, resp.Logs, 0)
}

func TestPushToAndroidRightTokenForStringLog(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = os.Getenv("FCM_CREDENTIAL")

	androidToken := os.Getenv("FCM_TEST_TOKEN")

	req := &PushNotification{
		Tokens:   []string{androidToken},
		Platform: core.PlatFormAndroid,
		Message:  "Welcome",
	}

	resp, err := PushToAndroid(context.Background(), req, cfg)
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

	// ignore check token length if send topic message
	req = &PushNotification{
		Message:  "Test",
		Platform: core.PlatFormAndroid,
		Topic:    "/topics/foo-bar",
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
		Tokens:   make([]string, 501),
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

func TestAndroidNotificationStructure(t *testing.T) {
	test := "test"
	req := &PushNotification{
		Tokens:         []string{"a", "b"},
		Message:        "Welcome",
		To:             test,
		Priority:       HIGH,
		MutableContent: true,
		Title:          test,
		Sound:          test,
		Data: D{
			"a": "1",
			"b": 2,
			"json": map[string]interface{}{
				"c": "3",
				"d": 4,
			},
		},
		Notification: &messaging.Notification{
			Title: test,
			Body:  "",
		},
	}

	messages := GetAndroidNotification(req)

	assert.Equal(t, test, messages[0].Notification.Title)
	assert.Equal(t, "Welcome", messages[0].Notification.Body)
	assert.Equal(t, "1", messages[0].Data["a"])
	assert.Equal(t, "2", messages[0].Data["b"])
	assert.Equal(t, "{\"c\":\"3\",\"d\":4}", messages[0].Data["json"])
	assert.NotNil(t, messages[0].APNS)
	assert.Equal(t, req.Sound, messages[0].APNS.Payload.Aps.Sound)
	assert.Equal(t, req.MutableContent, messages[0].APNS.Payload.Aps.MutableContent)

	// test empty body
	req = &PushNotification{
		Tokens: []string{"a", "b"},
		To:     test,
		Notification: &messaging.Notification{
			Body: "",
		},
	}
	messages = GetAndroidNotification(req)

	assert.Equal(t, "", messages[0].Notification.Body)
}

func TestAndroidBackgroundNotificationStructure(t *testing.T) {
	data := map[string]any{
		"a": "1",
		"b": 2,
		"json": map[string]interface{}{
			"c": "3",
			"d": 4,
		},
	}
	req := &PushNotification{
		Tokens:           []string{"a", "b"},
		Priority:         HIGH,
		ContentAvailable: true,
		Data:             data,
	}

	messages := GetAndroidNotification(req)

	assert.Equal(t, "1", messages[0].Data["a"])
	assert.Equal(t, "2", messages[0].Data["b"])
	assert.Equal(t, "{\"c\":\"3\",\"d\":4}", messages[0].Data["json"])
	assert.NotNil(t, messages[0].APNS)
	assert.Equal(t, req.ContentAvailable, messages[0].APNS.Payload.Aps.ContentAvailable)
	assert.True(t, reflect.DeepEqual(data, messages[0].APNS.Payload.Aps.CustomData))
}

// Tests for refactored helper functions

func TestSetupFCMNotification(t *testing.T) {
	tests := []struct {
		name        string
		req         *PushNotification
		wantTitle   string
		wantBody    string
		wantImage   string
		wantMutable bool
	}{
		{
			name: "title message and image",
			req: &PushNotification{
				Tokens:  []string{"token"},
				Title:   "Test Title",
				Message: "Test Message",
				Image:   "https://example.com/image.png",
			},
			wantTitle: "Test Title",
			wantBody:  "Test Message",
			wantImage: "https://example.com/image.png",
		},
		{
			name: "mutable content",
			req: &PushNotification{
				Tokens:         []string{"token"},
				Title:          "Title",
				MutableContent: true,
			},
			wantTitle:   "Title",
			wantMutable: true,
		},
		{
			name: "empty notification fields",
			req: &PushNotification{
				Tokens: []string{"token"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messages := GetAndroidNotification(tt.req)
			if len(messages) > 0 {
				msg := messages[0]
				if tt.wantTitle != "" {
					assert.Equal(t, tt.wantTitle, msg.Notification.Title)
				}
				if tt.wantBody != "" {
					assert.Equal(t, tt.wantBody, msg.Notification.Body)
				}
				if tt.wantImage != "" {
					assert.Equal(t, tt.wantImage, msg.Notification.ImageURL)
				}
				if tt.wantMutable {
					assert.NotNil(t, msg.APNS)
					assert.True(t, msg.APNS.Payload.Aps.MutableContent)
				}
			}
		})
	}
}

func TestSetupFCMContentAvailable(t *testing.T) {
	data := D{"key": "value"}
	req := &PushNotification{
		Tokens:           []string{"token"},
		ContentAvailable: true,
		Data:             data,
	}

	messages := GetAndroidNotification(req)
	assert.Len(t, messages, 1)
	msg := messages[0]

	assert.NotNil(t, msg.APNS)
	assert.Equal(t, "5", msg.APNS.Headers["apns-priority"])
	assert.True(t, msg.APNS.Payload.Aps.ContentAvailable)
	assert.Equal(t, data, D(msg.APNS.Payload.Aps.CustomData))
}

func TestSetAPNSSound(t *testing.T) {
	tests := []struct {
		name      string
		req       *PushNotification
		wantSound string
	}{
		{
			name: "sound with no existing APNS",
			req: &PushNotification{
				Tokens: []string{"token"},
				Sound:  "default",
			},
			wantSound: "default",
		},
		{
			name: "sound with existing APNS from mutable content",
			req: &PushNotification{
				Tokens:         []string{"token"},
				Title:          "Title",
				MutableContent: true,
				Sound:          "custom.aiff",
			},
			wantSound: "custom.aiff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messages := GetAndroidNotification(tt.req)
			if len(messages) > 0 {
				msg := messages[0]
				assert.NotNil(t, msg.APNS)
				assert.Equal(t, tt.wantSound, msg.APNS.Payload.Aps.Sound)
			}
		})
	}
}

func TestSetupFCMSound(t *testing.T) {
	req := &PushNotification{
		Tokens:   []string{"token"},
		Sound:    "test.aiff",
		Priority: HIGH,
	}

	messages := GetAndroidNotification(req)
	assert.Len(t, messages, 1)
	msg := messages[0]

	// Check APNS sound is set
	assert.NotNil(t, msg.APNS)
	assert.Equal(t, "test.aiff", msg.APNS.Payload.Aps.Sound)

	// Check Android config is set with sound
	assert.NotNil(t, msg.Android)
	assert.Equal(t, "test.aiff", msg.Android.Notification.Sound)
	assert.Equal(t, HIGH, msg.Android.Priority)
}

func TestConvertDataToStringMap(t *testing.T) {
	tests := []struct {
		name string
		data D
		want map[string]string
	}{
		{
			name: "empty data",
			data: D{},
			want: nil,
		},
		{
			name: "string values",
			data: D{"key1": "value1", "key2": "value2"},
			want: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name: "mixed values",
			data: D{"str": "text", "num": 42, "bool": true},
			want: map[string]string{"str": "text", "num": "42", "bool": "true"},
		},
		{
			name: "nested object",
			data: D{"nested": map[string]interface{}{"a": 1}},
			want: map[string]string{"nested": `{"a":1}`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertDataToStringMap(tt.data)
			if tt.want == nil {
				assert.Nil(t, got)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestBuildFCMMessage(t *testing.T) {
	req := &PushNotification{
		Tokens: []string{"token"},
		Title:  "Title",
		Notification: &messaging.Notification{
			Title: "Notification Title",
			Body:  "Notification Body",
		},
		Android: &messaging.AndroidConfig{
			Priority: HIGH,
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{"TTL": "3600"},
		},
		FCMOptions: &messaging.FCMOptions{
			AnalyticsLabel: "test-label",
		},
	}

	data := map[string]string{"key": "value"}
	msg := buildFCMMessage(req, data)

	assert.Equal(t, req.Notification, msg.Notification)
	assert.Equal(t, req.Android, msg.Android)
	assert.Equal(t, req.Webpush, msg.Webpush)
	assert.Equal(t, req.FCMOptions, msg.FCMOptions)
	assert.Equal(t, data, msg.Data)
}

func TestGetAndroidNotificationWithTopic(t *testing.T) {
	req := &PushNotification{
		Topic:     "test-topic",
		Condition: "'dogs' in topics",
		Message:   "Topic message",
	}

	messages := GetAndroidNotification(req)
	assert.Len(t, messages, 1)
	assert.Equal(t, "test-topic", messages[0].Topic)
	assert.Equal(t, "'dogs' in topics", messages[0].Condition)
}

func TestGetAndroidNotificationWithTokens(t *testing.T) {
	req := &PushNotification{
		Tokens:  []string{"token1", "token2", "token3"},
		Message: "Multi token message",
	}

	messages := GetAndroidNotification(req)
	assert.Len(t, messages, 3)
	assert.Equal(t, "token1", messages[0].Token)
	assert.Equal(t, "token2", messages[1].Token)
	assert.Equal(t, "token3", messages[2].Token)
}

func TestGetAndroidNotificationWithTopicAndTokens(t *testing.T) {
	req := &PushNotification{
		Topic:   "test-topic",
		Tokens:  []string{"token1", "token2"},
		Message: "Combined message",
	}

	messages := GetAndroidNotification(req)
	// 1 for topic + 2 for tokens = 3 messages
	assert.Len(t, messages, 3)
	assert.Equal(t, "test-topic", messages[0].Topic)
	assert.Equal(t, "token1", messages[1].Token)
	assert.Equal(t, "token2", messages[2].Token)
}
