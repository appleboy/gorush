package rpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/notify"
	"github.com/appleboy/gorush/rpc/proto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server is used to implement gorush grpc server.
type Server struct {
	cfg *config.ConfYaml
	mu  sync.Mutex
	// statusMap stores the serving status of the services this Server monitors.
	statusMap map[string]proto.HealthCheckResponse_ServingStatus
}

// NewServer returns a new Server.
func NewServer(cfg *config.ConfYaml) *Server {
	return &Server{
		cfg:       cfg,
		statusMap: make(map[string]proto.HealthCheckResponse_ServingStatus),
	}
}

// Check implements `service Health`.
func (s *Server) Check(ctx context.Context, in *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if in.Service == "" {
		// check the server overall health status.
		return &proto.HealthCheckResponse{
			Status: proto.HealthCheckResponse_SERVING,
		}, nil
	}
	if status, ok := s.statusMap[in.Service]; ok {
		return &proto.HealthCheckResponse{
			Status: status,
		}, nil
	}
	return nil, status.Error(codes.NotFound, "unknown service")
}

// Send implements helloworld.GreeterServer
func (s *Server) Send(ctx context.Context, in *proto.NotificationRequest) (*proto.NotificationReply, error) {
	badge := int(in.Badge)
	notification := notify.PushNotification{
		ID:               in.ID,
		Platform:         int(in.Platform),
		Tokens:           in.Tokens,
		Message:          in.Message,
		Title:            in.Title,
		Topic:            in.Topic,
		APIKey:           in.Key,
		Category:         in.Category,
		Sound:            in.Sound,
		ContentAvailable: in.ContentAvailable,
		ThreadID:         in.ThreadID,
		MutableContent:   in.MutableContent,
		Image:            in.Image,
		Priority:         strings.ToLower(in.GetPriority().String()),
		PushType:         in.PushType,
		Development:      in.Development,
	}

	if badge > 0 {
		notification.Badge = &badge
	}

	if in.Topic != "" && in.Platform == core.PlatFormAndroid {
		notification.To = in.Topic
	}

	if in.Alert != nil {
		notification.Alert = notify.Alert{
			Title:        in.Alert.Title,
			Body:         in.Alert.Body,
			Subtitle:     in.Alert.Subtitle,
			Action:       in.Alert.Action,
			ActionLocKey: in.Alert.Action,
			LaunchImage:  in.Alert.LaunchImage,
			LocArgs:      in.Alert.LocArgs,
			LocKey:       in.Alert.LocKey,
			TitleLocArgs: in.Alert.TitleLocArgs,
			TitleLocKey:  in.Alert.TitleLocKey,
		}
	}

	if in.Data != nil {
		notification.Data = in.Data.AsMap()
	}

	go func() {
		_, err := notify.SendNotification(ctx, &notification, s.cfg)
		if err != nil {
			logx.LogError.Error(err)
		}
	}()

	return &proto.NotificationReply{
		Success: true,
		Counts:  int32(len(notification.Tokens)),
	}, nil
}

// RunGRPCServer run gorush grpc server
func RunGRPCServer(ctx context.Context, cfg *config.ConfYaml) error {
	if !cfg.GRPC.Enabled {
		logx.LogAccess.Info("gRPC server is disabled.")
		return nil
	}

	recoveryOpt := grpc_recovery.WithRecoveryHandlerContext(
		func(ctx context.Context, p interface{}) error {
			fmt.Printf("[PANIC] %s\n%s", p, string(debug.Stack()))
			return status.Error(codes.Internal, "system has been broken")
		},
	)

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(recoveryOpt),
	}

	var s *grpc.Server

	if cfg.Core.SSL && cfg.Core.CertPath != "" && cfg.Core.KeyPath != "" {
		tlsCert, err := tls.LoadX509KeyPair(cfg.Core.CertPath, cfg.Core.KeyPath)
		if err != nil {
			logx.LogError.Error("failed to load tls cert file: ", err)
			return err
		}

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			ClientAuth:   tls.NoClientCert,
		}

		s = grpc.NewServer(
			grpc.Creds(credentials.NewTLS(tlsConfig)),
			grpc.StatsHandler(&ocgrpc.ServerHandler{}),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		)
	} else {
		s = grpc.NewServer(
			grpc.StatsHandler(&ocgrpc.ServerHandler{}),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		)
	}

	rpcSrv := NewServer(cfg)
	proto.RegisterGorushServer(s, rpcSrv)
	proto.RegisterHealthServer(s, rpcSrv)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		logx.LogError.Fatalln(err)
		return err
	}
	logx.LogAccess.Info("gRPC server is running on " + cfg.GRPC.Port + " port.")
	go func() {
		<-ctx.Done()
		s.GracefulStop() // graceful shutdown
		logx.LogAccess.Info("shutdown the gRPC server")
	}()
	if err = s.Serve(lis); err != nil {
		logx.LogError.Fatalln(err)
	}
	return err
}
