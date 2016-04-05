package gopush

import (
	"crypto/tls"
	"github.com/Sirupsen/logrus"
	apns "github.com/sideshow/apns2"
)

var (
	PushConf          ConfYaml
	CertificatePemIos tls.Certificate
	ApnsClient        *apns.Client
	LogAccess         *logrus.Logger
	LogError          *logrus.Logger
)
