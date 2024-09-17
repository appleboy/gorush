package rpc

import (
	"math"
	"testing"
)

func TestSafeIntToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int32
		wantErr bool
	}{
		{"Valid int32", 123, 123, false},
		{"Max int32", math.MaxInt32, math.MaxInt32, false},
		{"Min int32", math.MinInt32, math.MinInt32, false},
		{"Overflow int32", math.MaxInt32 + 1, 0, true},
		{"Underflow int32", math.MinInt32 - 1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safeIntToInt32(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("safeIntToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("safeIntToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
