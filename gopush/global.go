package gopush

import (
	"crypto/tls"
	apns "github.com/sideshow/apns2"
)

var (
	PushConf          ConfYaml
	CertificatePemIos tls.Certificate
	ApnsClient        *apns.Client
)
