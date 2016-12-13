// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package gorush

import (
	"github.com/fvbock/endless"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() error {
	var err error

	if PushConf.Core.SSL && PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
		err = endless.ListenAndServeTLS(":"+PushConf.Core.Port, PushConf.Core.CertPath, PushConf.Core.KeyPath, routerEngine())
	} else {
		err = endless.ListenAndServe(":"+PushConf.Core.Port, routerEngine())
	}

	return err
}
