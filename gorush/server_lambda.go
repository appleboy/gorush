// +build lambda

package gorush

import (
	"github.com/apex/gateway"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() error {
	if !PushConf.Core.Enabled {
		LogAccess.Debug("httpd server is disabled.")
		return nil
	}

	LogAccess.Debug("HTTPD server is running on " + PushConf.Core.Port + " port.")

	return gateway.ListenAndServe(PushConf.Core.Address+":"+PushConf.Core.Port, routerEngine())
}
