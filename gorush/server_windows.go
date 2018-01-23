// +build windows,!lambda

package gorush

import (
	"net/http"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() (err error) {
	if !PushConf.Core.Enabled {
		LogAccess.Debug("httpd server is disabled.")
		return nil
	}

	if PushConf.Core.AutoTLS.Enabled {
		s := autoTLSServer()
		err = s.ListenAndServeTLS("", "")
	} else if PushConf.Core.SSL && PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
		err = http.ListenAndServeTLS(PushConf.Core.Address+":"+PushConf.Core.Port, PushConf.Core.CertPath, PushConf.Core.KeyPath, routerEngine())
	} else {
		err = http.ListenAndServe(PushConf.Core.Address+":"+PushConf.Core.Port, routerEngine())
	}

	return
}
