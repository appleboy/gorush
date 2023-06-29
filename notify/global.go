package notify

import (
	"github.com/appleboy/go-fcm"
	"github.com/msalihkarakasli/go-hms-push/push/core"
	"github.com/sideshow/apns2"
)

var (
	// ApnsClients is apns client
	ApnsClients map[string]*apns2.Client
	// FCMClients is apns client
	FCMClients map[string]*fcm.Client
	// HMSClients is Huawei push client
	HMSClients map[string]*core.HMSClient
	// MaxConcurrentIOSPushes pool to limit the number of concurrent iOS pushes
	MaxConcurrentIOSPushes map[string]chan struct{}
)

const (
	HIGH   = "high"
	NORMAL = "nornal"
)
