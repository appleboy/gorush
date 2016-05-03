package gorush

import (
	"crypto/tls"
	"github.com/Sirupsen/logrus"
	"github.com/appleboy/gorush/gorush/config"
	apns "github.com/sideshow/apns2"
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
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
	// StatStorage implements the storage interface
	StatStorage Storage
)
