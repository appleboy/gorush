package rpc

import (
	"net"
	"sync"

	"github.com/appleboy/gorush/gorush"
	"github.com/appleboy/gorush/rpc/proto"

	"golang.org/x/net/context"
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
func RunGRPCServer() error {
	if !gorush.PushConf.GRPC.Enabled {
		gorush.LogAccess.Debug("gRPC server is disabled.")
		return nil
	}

	lis, err := net.Listen("tcp", ":"+gorush.PushConf.GRPC.Port)
	if err != nil {
		gorush.LogError.Errorf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	srv := NewServer()
	proto.RegisterGorushServer(s, srv)
	proto.RegisterHealthServer(s, srv)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	gorush.LogAccess.Debug("gRPC server is running on " + gorush.PushConf.GRPC.Port + " port.")
	if err := s.Serve(lis); err != nil {
		gorush.LogError.Errorf("failed to serve: %v", err)
		return err
	}

	return nil
}
