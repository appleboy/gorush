package gorush

import (
	"github.com/appleboy/gorush/config"

	"github.com/appleboy/go-fcm"
	"github.com/msalihkarakasli/go-hms-push/push/core"
	"github.com/sideshow/apns2"
)

var (
	// PushConf is gorush config
	PushConf config.ConfYaml
	// QueueNotification is chan type
	QueueNotification chan PushNotification
	// ApnsClient is apns client
	ApnsClient *apns2.Client
	// FCMClient is apns client
	FCMClient *fcm.Client
	// HMSClient is Huawei push client
	HMSClient *core.HMSClient
	// MaxConcurrentIOSPushes pool to limit the number of concurrent iOS pushes
	MaxConcurrentIOSPushes chan struct{}
)
