// +build windows

package gorush

import (
	"net/http"

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
		err = s.ListenAndServeTLS("", "")
	} else if PushConf.Core.SSL && PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
		err = http.ListenAndServeTLS(":"+PushConf.Core.Port, PushConf.Core.CertPath, PushConf.Core.KeyPath, routerEngine())
	} else {
		err = http.ListenAndServe(":"+PushConf.Core.Port, routerEngine())
	}

	return
}
