# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gorush is a push notification microserver written in Go that supports sending notifications to iOS (APNS), Android (FCM), and Huawei (HMS) devices. It provides both HTTP REST API and gRPC interfaces.

## Architecture

### Request Flow

1. **main.go** parses CLI flags, loads config, initializes platform clients (APNS/FCM/HMS), then starts HTTP (Gin) and gRPC servers via `graceful.Manager`
2. **router/** receives push requests at `POST /api/push`, validates them, and enqueues `PushNotification` messages into the queue
3. **app/worker.go** creates the queue worker (local/NSQ/NATS/Redis) with `notify.Run(cfg)` as the processing function
4. **notify/** dispatches notifications to platform-specific push functions (`PushToIOS`, `PushToAndroid`, `PushToHuawei`)
5. **status/** + **storage/** track success/failure counts across configurable backends

### Key Packages

- **app/**: Application orchestration — CLI send helpers (`sender.go`), config validation/merge (`config.go`), CLI options (`options.go`), queue worker creation (`worker.go`)
- **config/**: YAML config with Viper, env var overrides. Reference config: `config/testdata/config.yml`
- **core/**: Shared types and interfaces — `Platform` enum, queue engine constants (`core/queue.go`), storage interface (`core/storage.go`), health check interface (`core/health.go`)
- **notify/**: Platform-specific push implementations. Each platform has its own file (`notification_apns.go`, `notification_fcm.go`, `notification_hms.go`). `global.go` holds shared client state
- **router/**: Gin HTTP server with REST endpoints and Prometheus metrics
- **rpc/**: gRPC server. Proto definitions in `rpc/proto/`
- **storage/**: Storage backends (memory, Redis, BoltDB, BuntDB, LevelDB, BadgerDB) all implement `core.Storage`
- **logx/**: Logging utilities wrapping zerolog/logrus

## Development Commands

### Build and Run

```bash
make build                  # Build binary to release/gorush
make install                # Install to $GOPATH/bin
make dev                    # Hot reload with air
```

### Testing

```bash
# Full test suite (requires FCM credentials)
export FCM_CREDENTIAL="/path/to/firebase-credentials.json"
export FCM_TEST_TOKEN="your_test_device_token"
make test

# Run a single test
go test -v -tags sqlite -run TestFunctionName ./package/...

# Run tests for a specific package
go test -v -tags sqlite ./notify/...
go test -v -tags sqlite ./config/...
```

The `-tags sqlite` flag is required for all test commands (it's the default build tag).

### Linting and Formatting

```bash
make lint                   # Run golangci-lint (auto-installs if missing)
make fmt                    # Format code with golangci-lint fmt
```

Linter config is in `.golangci.yml` (v2 format). Uses golangci-lint v2.

### Protocol Buffers

```bash
make generate_proto         # Generate both Go and JS proto files
```

## Build Tags

- `sqlite` — default tag, required for standard builds and tests
- `lambda` — for AWS Lambda builds (replaces sqlite)
- Set custom tags via `TAGS` environment variable

## Configuration

Config uses Viper with YAML files. Key sections: `core`, `grpc`, `android`, `ios`, `huawei`, `queue`, `stat`, `log`, `api`. See `config/testdata/config.yml` for all options. Environment variables can override any config value.

## API Endpoints

- `POST /api/push` — send push notifications
- `GET /api/stat/go` — Go runtime stats
- `GET /api/stat/app` — push statistics
- `GET /api/config` — current config
- `GET /metrics` — Prometheus metrics
- `GET /healthz` — health check
