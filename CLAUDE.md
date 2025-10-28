# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gorush is a push notification microserver written in Go that supports sending notifications to iOS (APNS), Android (FCM), and Huawei (HMS) devices. It provides both HTTP REST API and gRPC interfaces for sending push notifications.

## Architecture

### Core Components

- **main.go**: Entry point that handles CLI flags, configuration loading, and server initialization
- **config/**: Configuration management with YAML support and environment variable overrides
- **core/**: Core abstractions for queue, storage, and health check interfaces
- **notify/**: Push notification implementations for different platforms (APNS, FCM, HMS)
- **router/**: HTTP server setup using Gin framework with REST endpoints
- **rpc/**: gRPC server implementation with protocol buffer definitions
- **status/**: Application statistics and metrics collection
- **storage/**: Multiple storage backends (memory, Redis, BoltDB, BuntDB, LevelDB, BadgerDB)
- **logx/**: Logging utilities and interfaces

### Key Features

- Multi-platform support (iOS APNS, Android FCM, Huawei HMS)
- Multiple queue engines (local channels, NSQ, NATS, Redis streams)
- Configurable storage backends for statistics
- Prometheus metrics support
- gRPC and REST API interfaces
- Docker and Kubernetes deployment support
- AWS Lambda support

## Development Commands

### Building

```bash
# Build for current platform
make build

# Build for specific platforms
make build_linux_amd64
make build_darwin_arm64
make build_linux_lambda

# Install locally
make install
```

### Testing

```bash
# Run tests (requires FCM_CREDENTIAL and FCM_TEST_TOKEN environment variables)
make test

# Check required environment variables first
make init
```

### Development Tools

```bash
# Run with hot reload
make dev

# Clean build artifacts
make clean

# Check available commands
make help
```

### Linting and Code Quality

```bash
# The project uses golangci-lint with configuration in .golangci.yml
# Run linting (if golangci-lint is installed):
golangci-lint run

# Supported linters include: bodyclose, errcheck, gosec, govet, staticcheck, etc.
# Formatters: gofmt, gofumpt, goimports
```

### Protocol Buffers

```bash
# Install protoc dependencies
make proto_install

# Generate Go protobuf files
make generate_proto_go

# Generate JavaScript protobuf files
make generate_proto_js

# Generate all protobuf files
make generate_proto
```

## Configuration

The application uses YAML configuration files. See `config/testdata/config.yml` for the complete configuration example.

Key configuration sections:

- **core**: Server settings, workers, queue configuration
- **grpc**: gRPC server settings
- **android**: Firebase Cloud Messaging settings
- **ios**: Apple Push Notification settings
- **huawei**: Huawei Mobile Services settings
- **queue**: Queue engine configuration (local/nsq/nats/redis)
- **stat**: Statistics storage backend configuration
- **log**: Logging configuration

## Running the Server

### Basic Usage

```bash
# Use default config
./gorush

# Use custom config file
./gorush -c config.yml

# CLI notification examples
./gorush -android -m "Hello World" --fcm-key /path/to/key.json -t "device_token"
./gorush -ios -m "Hello World" -i /path/to/cert.pem -t "device_token"
```

### Docker

```bash
docker run --name gorush -p 80:8088 appleboy/gorush
```

## API Endpoints

### REST API

- `GET /api/stat/go` - Go runtime statistics
- `GET /api/stat/app` - Application push statistics
- `GET /api/config` - Current configuration
- `POST /api/push` - Send push notifications
- `GET /metrics` - Prometheus metrics
- `GET /healthz` - Health check

### gRPC

Enable gRPC server in config and use port 9000 by default. Protocol definitions are in `rpc/proto/`.

## Testing Requirements

Tests require FCM credentials and test tokens:

```bash
export FCM_CREDENTIAL="/path/to/firebase-credentials.json"
export FCM_TEST_TOKEN="your_test_device_token"
```

## Build Tags

- Default: `sqlite` tag enabled
- `lambda` tag: For AWS Lambda builds
- Custom tags can be set with `TAGS` environment variable
