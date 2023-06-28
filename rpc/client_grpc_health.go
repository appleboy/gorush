package rpc

import (
	"context"

	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/rpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// generate protobuffs
//   protoc --go_out=plugins=grpc,import_path=proto:. *.proto

type healthClient struct {
	client proto.HealthClient
	conn   *grpc.ClientConn
}

// NewGrpcHealthClient returns a new grpc Client.
func NewGrpcHealthClient(conn *grpc.ClientConn) core.Health {
	client := new(healthClient)
	client.client = proto.NewHealthClient(conn)
	client.conn = conn
	return client
}

func (c *healthClient) Close() error {
	return c.conn.Close()
}

func (c *healthClient) Check(ctx context.Context) (bool, error) {
	var res *proto.HealthCheckResponse
	var err error
	req := new(proto.HealthCheckRequest)

	res, err = c.client.Check(ctx, req)
	if err == nil {
		if res.GetStatus() == proto.HealthCheckResponse_SERVING {
			return true, nil
		}
		return false, nil
	}
	//nolint:exhaustive
	switch status.Code(err) {
	case
		codes.Aborted,
		codes.DataLoss,
		codes.DeadlineExceeded,
		codes.Internal,
		codes.Unavailable:
		// non-fatal errors
	default:
		return false, err
	}

	return false, err
}
