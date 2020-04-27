// +build !lambda

package gorush

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer(ctx context.Context) (err error) {
	if !PushConf.Core.Enabled {
		LogAccess.Info("httpd server is disabled.")
		return nil
	}

	server := &http.Server{
		Addr:    PushConf.Core.Address + ":" + PushConf.Core.Port,
		Handler: routerEngine(),
	}

	LogAccess.Info("HTTPD server is running on " + PushConf.Core.Port + " port.")
	if PushConf.Core.AutoTLS.Enabled {
		return startServer(ctx, autoTLSServer())
	} else if PushConf.Core.SSL {
		config := &tls.Config{
			MinVersion: tls.VersionTLS10,
		}

		if config.NextProtos == nil {
			config.NextProtos = []string{"http/1.1"}
		}

		config.Certificates = make([]tls.Certificate, 1)
		if PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
			config.Certificates[0], err = tls.LoadX509KeyPair(PushConf.Core.CertPath, PushConf.Core.KeyPath)
			if err != nil {
				LogError.Error("Failed to load https cert file: ", err)
				return err
			}
		} else if PushConf.Core.CertBase64 != "" && PushConf.Core.KeyBase64 != "" {
			cert, err := base64.StdEncoding.DecodeString(PushConf.Core.CertBase64)
			if err != nil {
				LogError.Error("base64 decode error:", err.Error())
				return err
			}
			key, err := base64.StdEncoding.DecodeString(PushConf.Core.KeyBase64)
			if err != nil {
				LogError.Error("base64 decode error:", err.Error())
				return err
			}
			if config.Certificates[0], err = tls.X509KeyPair(cert, key); err != nil {
				LogError.Error("tls key pair error:", err.Error())
				return err
			}
		} else {
			return errors.New("missing https cert config")
		}

		server.TLSConfig = config
	}

	return startServer(ctx, server)
}

func listenAndServe(ctx context.Context, s *http.Server) error {
	var g errgroup.Group
	g.Go(func() error {
		select {
		case <-ctx.Done():
			timeout := time.Duration(PushConf.Core.ShutdownTimeout) * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			return s.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return g.Wait()
}

func listenAndServeTLS(ctx context.Context, s *http.Server) error {
	var g errgroup.Group
	g.Go(func() error {
		select {
		case <-ctx.Done():
			timeout := time.Duration(PushConf.Core.ShutdownTimeout) * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			return s.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		if err := s.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	return g.Wait()
}

func startServer(ctx context.Context, s *http.Server) error {
	if s.TLSConfig == nil {
		return listenAndServe(ctx, s)
	}

	return listenAndServeTLS(ctx, s)
}
