// +build !windows

package gorush

import (
	"crypto/tls"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"golang.org/x/crypto/acme/autocert"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() (err error) {
	if PushConf.Core.AutoTLS.Enabled {
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(PushConf.Core.AutoTLS.Host),
			Cache:      autocert.DirCache(PushConf.Core.AutoTLS.Folder),
		}

		s := &http.Server{
			Addr:      ":https",
			TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
			Handler:   routerEngine(),
		}
		err = gracehttp.Serve(s)
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
			Addr:      ":" + PushConf.Core.Port,
			Handler:   routerEngine(),
			TLSConfig: config,
		})
	} else {
		err = gracehttp.Serve(&http.Server{
			Addr:    ":" + PushConf.Core.Port,
			Handler: routerEngine(),
		})
	}

	return
}
