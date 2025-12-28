# App Backend

## Development

### Run with auto-reload (recommended for development)

```bash
air
```

### Run normally

```bash
go run cmd/server/main.go
```

### Build

```bash
go build -o bin/server cmd/server/main.go
```

## Configuration

Set the following environment variables or create a `.env` file:

```env
APP_NAME=messenger
ENV=development
HTTP_PORT=8080
GRPC_PORT=50051
```

## Features

- 🔄 Auto HTTP & gRPC service registration
- 📝 Colorful async logging for HTTP & gRPC
- 🏥 Health check endpoints (HTTP & gRPC)
