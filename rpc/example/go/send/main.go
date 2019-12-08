package main

import (
	"context"
	"log"

	"github.com/appleboy/gorush/rpc/proto"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewGorushClient(conn)

	r, err := c.Send(context.Background(), &proto.NotificationRequest{
		Platform: 2,
		Tokens:   []string{"1234567890"},
		Message:  "test message",
		Badge:    1,
		Category: "test",
		Sound:    "test",
		Alert: &proto.Alert{
			Title:    "Test Title",
			Body:     "Test Alert Body",
			Subtitle: "Test Alert Sub Title",
			LocKey:   "Test loc key",
			LocArgs:  []string{"test", "test"},
		},
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"key1": {
					Kind: &structpb.Value_StringValue{StringValue: "welcome"},
				},
				"key2": {
					Kind: &structpb.Value_NumberValue{NumberValue: 2},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Success: %t\n", r.Success)
	log.Printf("Count: %d\n", r.Counts)
}
