package rpc

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/appleboy/gorush/gorush"
	"github.com/appleboy/gorush/rpc/proto"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server is used to implement gorush grpc server.
type Server struct {
	mu sync.Mutex
	// statusMap stores the serving status of the services this Server monitors.
	statusMap map[string]proto.HealthCheckResponse_ServingStatus
}

// NewServer returns a new Server.
func NewServer() *Server {
	return &Server{
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
	var badge = int(in.Badge)
	notification := gorush.PushNotification{
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
	}

	if badge > 0 {
		notification.Badge = &badge
	}

	if in.Alert != nil {
		notification.Alert = gorush.Alert{
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
		notification.Data = map[string]interface{}{}
		for k, v := range in.Data.Fields {
			notification.Data[k] = v
		}
	}

	go gorush.SendNotification(notification)

	return &proto.NotificationReply{
		Success: true,
		Counts:  int32(len(notification.Tokens)),
	}, nil
}

// RunGRPCServer run gorush grpc server
func RunGRPCServer(ctx context.Context) error {
	if !gorush.PushConf.GRPC.Enabled {
		gorush.LogAccess.Info("gRPC server is disabled.")
		return nil
	}

	s := grpc.NewServer()
	rpcSrv := NewServer()
	proto.RegisterGorushServer(s, rpcSrv)
	proto.RegisterHealthServer(s, rpcSrv)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	gorush.LogAccess.Info("gRPC server is running on " + gorush.PushConf.GRPC.Port + " port.")

	srv := &http.Server{
		Addr:    ":" + gorush.PushConf.GRPC.Port,
		Handler: s,
	}

	var g errgroup.Group
	g.Go(func() error {
		select {
		case <-ctx.Done():
			timeout := time.Duration(gorush.PushConf.Core.ShutdownTimeout) * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			return srv.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return srv.ListenAndServe()
	})
	return g.Wait()
}
