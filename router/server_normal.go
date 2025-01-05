//go:build !lambda
// +build !lambda

package router

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"

	"github.com/golang-queue/queue"
	"golang.org/x/sync/errgroup"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer(ctx context.Context, cfg *config.ConfYaml, q *queue.Queue, s ...*http.Server) (err error) {
	var server *http.Server

	if !cfg.Core.Enabled {
		logx.LogAccess.Info("httpd server is disabled.")
		return nil
	}

	if len(s) == 0 {
		//nolint:gosec
		server = &http.Server{
			Addr:    cfg.Core.Address + ":" + cfg.Core.Port,
			Handler: routerEngine(cfg, q),
		}
	} else {
		server = s[0]
	}

	logx.LogAccess.Info("HTTPD server is running on " + cfg.Core.Port + " port.")
	if cfg.Core.AutoTLS.Enabled {
		return startServer(ctx, autoTLSServer(cfg, q), cfg)
	} else if cfg.Core.SSL {
		config := &tls.Config{
			MinVersion: tls.VersionTLS12,
		}

		if config.NextProtos == nil {
			config.NextProtos = []string{"http/1.1"}
		}

		config.Certificates = make([]tls.Certificate, 1)
		//nolint:gocritic
		if cfg.Core.CertPath != "" && cfg.Core.KeyPath != "" {
			config.Certificates[0], err = tls.LoadX509KeyPair(cfg.Core.CertPath, cfg.Core.KeyPath)
			if err != nil {
				logx.LogError.Error("Failed to load https cert file: ", err)
				return err
			}
		} else if cfg.Core.CertBase64 != "" && cfg.Core.KeyBase64 != "" {
			cert, err := base64.StdEncoding.DecodeString(cfg.Core.CertBase64)
			if err != nil {
				logx.LogError.Error("base64 decode error:", err.Error())
				return err
			}
			key, err := base64.StdEncoding.DecodeString(cfg.Core.KeyBase64)
			if err != nil {
				logx.LogError.Error("base64 decode error:", err.Error())
				return err
			}
			if config.Certificates[0], err = tls.X509KeyPair(cert, key); err != nil {
				logx.LogError.Error("tls key pair error:", err.Error())
				return err
			}
		} else {
			return errors.New("missing https cert config")
		}

		server.TLSConfig = config
	}

	return startServer(ctx, server, cfg)
}

func listenAndServe(ctx context.Context, s *http.Server, cfg *config.ConfYaml) error {
	var g errgroup.Group
	g.Go(func() error {
		<-ctx.Done()
		timeout := time.Duration(cfg.Core.ShutdownTimeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return s.Shutdown(ctx)
	})
	g.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return g.Wait()
}

func listenAndServeTLS(ctx context.Context, s *http.Server, cfg *config.ConfYaml) error {
	var g errgroup.Group
	g.Go(func() error {
		<-ctx.Done()
		timeout := time.Duration(cfg.Core.ShutdownTimeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return s.Shutdown(ctx)
	})
	g.Go(func() error {
		if err := s.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return g.Wait()
}

func startServer(ctx context.Context, s *http.Server, cfg *config.ConfYaml) error {
	if s.TLSConfig == nil {
		return listenAndServe(ctx, s, cfg)
	}

	return listenAndServeTLS(ctx, s, cfg)
}
