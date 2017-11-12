package main

import (
	"log"

	"github.com/appleboy/gorush/rpc/proto"

	"golang.org/x/net/context"
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

	c := proto.NewHealthClient(conn)
	r, err := c.Check(context.Background(), &proto.HealthCheckRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Health: %d\n", r.GetStatus())
}
