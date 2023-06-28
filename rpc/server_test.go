package rpc

// const gRPCAddr = "localhost:9000"

// func initTest() *config.ConfYaml {
// 	cfg, _ := config.LoadConf()
// 	cfg.Core.Mode = "test"
// 	return cfg
// }

// func TestGracefulShutDownGRPCServer(t *testing.T) {
// 	cfg := initTest()
// 	cfg.GRPC.Enabled = true
// 	cfg.GRPC.Port = "9000"
// 	cfg.Log.Format = "json"

// 	// Run gRPC server
// 	ctx, gRPCContextCancel := context.WithCancel(context.Background())
// 	go func() {
// 		if err := RunGRPCServer(ctx, cfg); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// gRPC client conn
// 	conn, err := grpc.Dial(
// 		gRPCAddr,
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
// 	) // wait for server ready
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Stop gRPC server
// 	go gRPCContextCancel()

// 	// wait for client connection would be closed
// 	for conn.GetState() != connectivity.TransientFailure {
// 	}
// 	conn.Close()
// }
