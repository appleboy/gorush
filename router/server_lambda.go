// +build lambda

package router

import (
	"context"
	"net/http"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"

	"github.com/apex/gateway"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer(ctx context.Context, cfg config.ConfYaml, s ...*http.Server) (err error) {
	if !cfg.Core.Enabled {
		logx.LogAccess.Debug("httpd server is disabled.")
		return nil
	}

	logx.LogAccess.Info("HTTPD server is running on " + cfg.Core.Port + " port.")

	return gateway.ListenAndServe(cfg.Core.Address+":"+cfg.Core.Port, routerEngine(cfg))
}
