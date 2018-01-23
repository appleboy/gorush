// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package lambda

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

// Start takes a handler, and talks to and internal Lambda endpoint to pass Invoke requests to the handler.  If a
// handler does not match one of the supported types, the lambda package will respond to new invokes served by in
// internal endpoint with an appropriate error message.  Start blocks, and does not return after being called.
//
// Rules:
// * handler must be a function
// * handler may take between 0 and two arguments.
//   * If there are two arguments, the first argument must implement "context.Context".
// * handler may return between 0 and two arguments.
//   * If there are two return values, the second argument must implement "error".
//   * If there is one return value it must implement "error".
//
// func ()
// func () error
// func (TIn) error
// func () (TOut, error)
// func (TIn) (TOut, error)
// func (context.Context) error
// func (context.Context, TIn) error
// func (context.Context) (TOut, error)
// func (context.Context, TIn) (TOut, error)
//
// Where '''TIn''' and '''TOut''' are types compatible with the ''encoding/json'' standard library.
// See https://golang.org/pkg/encoding/json/#Unmarshal for how deserialization behaves
func Start(handler interface{}) {
	port := os.Getenv("_LAMBDA_SERVER_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	wrappedHandler := newHandler(handler)
	function := new(Function)
	function.handler = wrappedHandler
	err = rpc.Register(function)
	if err != nil {
		log.Fatal("failed to register handler function")
	}
	rpc.Accept(lis)
	log.Fatal("accept should not have returned")
}
