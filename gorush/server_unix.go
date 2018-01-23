// +build !windows,!lambda

package gorush

import (
	"crypto/tls"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() (err error) {
	if !PushConf.Core.Enabled {
		LogAccess.Debug("httpd server is disabled.")
		return nil
	}

	LogAccess.Debug("HTTPD server is running on " + PushConf.Core.Port + " port.")
	if PushConf.Core.AutoTLS.Enabled {
		err = gracehttp.Serve(autoTLSServer())
	} else if PushConf.Core.SSL && PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
		config := &tls.Config{
			MinVersion: tls.VersionTLS10,
		}

		if config.NextProtos == nil {
			config.NextProtos = []string{"http/1.1"}
		}

		config.Certificates = make([]tls.Certificate, 1)
		config.Certificates[0], err = tls.LoadX509KeyPair(PushConf.Core.CertPath, PushConf.Core.KeyPath)
		if err != nil {
			LogError.Error("Failed to load https cert file: ", err)
			return err
		}

		err = gracehttp.Serve(&http.Server{
			Addr:      PushConf.Core.Address + ":" + PushConf.Core.Port,
			Handler:   routerEngine(),
			TLSConfig: config,
		})
	} else {
		err = gracehttp.Serve(&http.Server{
			Addr:    PushConf.Core.Address + ":" + PushConf.Core.Port,
			Handler: routerEngine(),
		})
	}

	return
}
