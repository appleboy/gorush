package rpc

import (
	"net"

	"github.com/axiomzen/gorush/gorush"
	pb "github.com/axiomzen/gorush/rpc/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement gorush grpc server.
type server struct{}

// Send implements helloworld.GreeterServer
func (s *server) Send(ctx context.Context, in *pb.NotificationRequest) (*pb.NotificationReply, error) {
	notification := gorush.PushNotification{
		Platform: int(in.Platform),
		Tokens:   in.Tokens,
		Message:  in.Message,
		Title:    in.Title,
		Topic:    in.Topic,
		APIKey:   in.Key,
	}

	go gorush.SendNotification(notification)

	return &pb.NotificationReply{
		Success: false,
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
		gorush.LogError.Error("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	pb.RegisterGorushServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	gorush.LogAccess.Debug("gRPC server is running on " + gorush.PushConf.GRPC.Port + " port.")
	if err := s.Serve(lis); err != nil {
		gorush.LogError.Error("failed to serve: %v", err)
		return err
	}

	return nil
}
