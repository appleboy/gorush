package gorush

import (
	"crypto/tls"

	"github.com/axiomzen/gorush/config"
	"github.com/axiomzen/gorush/storage"

	"github.com/appleboy/go-fcm"
	apns "github.com/sideshow/apns2"
	"github.com/sirupsen/logrus"
)

var (
	// PushConf is gorush config
	PushConf config.ConfYaml
	// QueueNotification is chan type
	QueueNotification chan PushNotification
	// CertificatePemIos is ios certificate file
	CertificatePemIos tls.Certificate
	// ApnsClient is apns client
	ApnsClient *apns.Client
	// FCMClient is apns client
	FCMClient *fcm.Client
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
	// StatStorage implements the storage interface
	StatStorage storage.Storage
)
