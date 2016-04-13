package gopush

import (
	"crypto/tls"
	"github.com/Sirupsen/logrus"
	apns "github.com/sideshow/apns2"
)

var (
	// PushConf is gorush config
	PushConf ConfYaml
	// CertificatePemIos is ios certificate file
	CertificatePemIos tls.Certificate
	// ApnsClient is apns client
	ApnsClient *apns.Client
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
)
