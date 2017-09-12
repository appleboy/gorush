package main

import (
	"log"

	pb "github.com/axiomzen/gorush/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGorushClient(conn)

	r, err := c.Send(context.Background(), &pb.NotificationRequest{
		Platform: 2,
		Tokens:   []string{"1234567890"},
		Message:  "test message",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Success: %t\n", r.Success)
	log.Printf("Count: %d\n", r.Counts)
}
