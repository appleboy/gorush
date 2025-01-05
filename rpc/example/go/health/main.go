package main

import (
	"context"
	"log"
	"time"

	"github.com/appleboy/gorush/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	address = "localhost:9000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := rpc.NewGrpcHealthClient(conn)

	for {
		ok, err := client.Check(context.Background())
		if !ok || err != nil {
			log.Printf("can't connect grpc server: %v, code: %v\n", err, status.Code(err))
		} else {
			log.Println("connect the grpc server successfully")
		}

		<-time.After(time.Second)
	}
}
