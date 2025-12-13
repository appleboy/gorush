package notify

import (
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/appleboy/go-hms-push/push/model"
	"github.com/stretchr/testify/assert"
)

func TestMissingHuaweiAppSecret(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = false
	cfg.Huawei.Enabled = true
	cfg.Huawei.AppSecret = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "missing huawei app secret", err.Error())
}

func TestMissingHuaweiAppID(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = false
	cfg.Huawei.Enabled = true
	cfg.Huawei.AppID = ""

	err := CheckPushConf(cfg)

	assert.Error(t, err)
	assert.Equal(t, "missing huawei app id", err.Error())
}

func TestMissingAppSecretForInitHMSClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	client, err := InitHMSClient(cfg, "", "APP_SECRET")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "missing huawei app secret", err.Error())
}

func TestMissingAppIDForInitHMSClient(t *testing.T) {
	cfg, _ := config.LoadConf()
	client, err := InitHMSClient(cfg, "APP_ID", "")

	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, "missing huawei app id", err.Error())
}

// Tests for refactored helper functions

func TestSetHuaweiMessageTarget(t *testing.T) {
	tests := []struct {
		name          string
		req           *PushNotification
		wantTokens    []string
		wantTopic     string
		wantCondition string
	}{
		{
			name: "tokens only",
			req: &PushNotification{
				Tokens: []string{"token1", "token2"},
			},
			wantTokens: []string{"token1", "token2"},
		},
		{
			name: "topic only",
			req: &PushNotification{
				Topic: "test-topic",
			},
			wantTopic: "test-topic",
		},
		{
			name: "condition only",
			req: &PushNotification{
				Condition: "'dogs' in topics",
			},
			wantCondition: "'dogs' in topics",
		},
		{
			name: "all fields",
			req: &PushNotification{
				Tokens:    []string{"token1"},
				Topic:     "topic",
				Condition: "condition",
			},
			wantTokens:    []string{"token1"},
			wantTopic:     "topic",
			wantCondition: "condition",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msgRequest, err := GetHuaweiNotification(tt.req)
			assert.NoError(t, err)

			if len(tt.wantTokens) > 0 {
				assert.Equal(t, tt.wantTokens, msgRequest.Message.Token)
			}
			if tt.wantTopic != "" {
				assert.Equal(t, tt.wantTopic, msgRequest.Message.Topic)
			}
			if tt.wantCondition != "" {
				assert.Equal(t, tt.wantCondition, msgRequest.Message.Condition)
			}
		})
	}
}

func TestSetHuaweiAndroidConfig(t *testing.T) {
	collapseKey := 1
	req := &PushNotification{
		Tokens:            []string{"token"},
		Priority:          HIGH,
		HuaweiCollapseKey: collapseKey,
		Category:          "test-category",
		HuaweiTTL:         "86400s",
		BiTag:             "bi-tag-123",
		FastAppTarget:     1,
	}

	msgRequest, err := GetHuaweiNotification(req)
	assert.NoError(t, err)

	android := msgRequest.Message.Android
	assert.Equal(t, "HIGH", android.Urgency)
	assert.Equal(t, collapseKey, android.CollapseKey)
	assert.Equal(t, "test-category", android.Category)
	assert.Equal(t, "86400s", android.TTL)
	assert.Equal(t, "bi-tag-123", android.BiTag)
	assert.Equal(t, 1, android.FastAppTarget)
}

func TestSetHuaweiNotificationContent(t *testing.T) {
	tests := []struct {
		name      string
		req       *PushNotification
		wantTitle string
		wantBody  string
		wantImage string
		wantSound string
	}{
		{
			name: "title and message",
			req: &PushNotification{
				Tokens:  []string{"token"},
				Title:   "Test Title",
				Message: "Test Message",
			},
			wantTitle: "Test Title",
			wantBody:  "Test Message",
		},
		{
			name: "image",
			req: &PushNotification{
				Tokens: []string{"token"},
				Image:  "https://example.com/image.png",
			},
			wantImage: "https://example.com/image.png",
		},
		{
			name: "custom sound",
			req: &PushNotification{
				Tokens: []string{"token"},
				Sound:  "custom.mp3",
			},
			wantSound: "custom.mp3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msgRequest, err := GetHuaweiNotification(tt.req)
			assert.NoError(t, err)

			notification := msgRequest.Message.Android.Notification
			if tt.wantTitle != "" {
				assert.Equal(t, tt.wantTitle, notification.Title)
			}
			if tt.wantBody != "" {
				assert.Equal(t, tt.wantBody, notification.Body)
			}
			if tt.wantImage != "" {
				assert.Equal(t, tt.wantImage, notification.Image)
			}
			if tt.wantSound != "" {
				assert.Equal(t, tt.wantSound, notification.Sound)
			}
		})
	}
}

func TestHuaweiNotificationDefaultSound(t *testing.T) {
	req := &PushNotification{
		Tokens:  []string{"token"},
		Title:   "Title",
		Message: "Message",
		// No sound specified
	}

	msgRequest, err := GetHuaweiNotification(req)
	assert.NoError(t, err)

	notification := msgRequest.Message.Android.Notification
	assert.True(t, notification.DefaultSound)
}

func TestHuaweiNotificationWithData(t *testing.T) {
	req := &PushNotification{
		Tokens:     []string{"token"},
		HuaweiData: `{"key":"value"}`,
	}

	msgRequest, err := GetHuaweiNotification(req)
	assert.NoError(t, err)

	assert.Equal(t, `{"key":"value"}`, msgRequest.Message.Data)
}

func TestHuaweiNotificationWithCustomNotification(t *testing.T) {
	req := &PushNotification{
		Tokens: []string{"token"},
		HuaweiNotification: &model.AndroidNotification{
			Title: "Custom Title",
			Body:  "Custom Body",
			Image: "custom-image.png",
		},
	}

	msgRequest, err := GetHuaweiNotification(req)
	assert.NoError(t, err)

	notification := msgRequest.Message.Android.Notification
	assert.Equal(t, "Custom Title", notification.Title)
	assert.Equal(t, "Custom Body", notification.Body)
	assert.Equal(t, "custom-image.png", notification.Image)
	// ClickAction should be set to default
	assert.NotNil(t, notification.ClickAction)
}
