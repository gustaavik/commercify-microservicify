# Commercify Microservicify

This project is a microservices-based e-commerce backend written in Go, using gRPC for service communication and Consul for service discovery. It is structured for scalability and maintainability, with separate services for products, orders, and a gateway.

## Features

- Product and Order microservices
- gRPC APIs for inter-service communication
- Consul for service discovery and health checks
- Docker and Docker Compose support
- Common utilities and models for shared logic

## Project Structure

```
api/            # gRPC API definitions (protobuf)
bin/            # Compiled binaries
cmd/            # Service entrypoints (main.go for each service)
internal/       # Service implementations (orders, products, gateway)
pkg/            # Shared packages (common, db, models, trpc)
proto/          # Protobuf source files
Dockerfile      # Docker build file
docker-compose.yml # Multi-service orchestration
Makefile        # Build and management commands
go.mod, go.sum  # Go module files
```

## Getting Started

### Prerequisites

- Go 1.20+
- Docker & Docker Compose

### Quickstart (Development)

1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd commercify-microservicify
   ```
2. **Generate gRPC code:**
   ```sh
   make proto
   ```
3. **Copy environment file:**
   ```sh
   cp .env.example .env.local
   ```
4. **Start all services (recommended in a standalone terminal):**
   ```sh
   make run-grid
   ```
   This will launch Consul and all services in a tmux session for local development.
5. **Access Consul UI:**
   Visit [http://localhost:8500](http://localhost:8500)

### Build Binaries (Manual)

```sh
make build
```

### Running Services Locally

Each service can be run individually from its `cmd/<service>/main.go` entrypoint.

```sh
cd cmd/product && go run main.go
cd cmd/order && go run main.go
cd cmd/gateway && go run main.go
```

## API

- gRPC endpoints are defined in `proto/` and generated into `api/`.
- Use a gRPC client or tools like [grpcurl](https://github.com/fullstorydev/grpcurl) to interact with the APIs.

## Service Discovery

- Services register themselves with Consul on startup.
- The gateway and other services discover each other via Consul DNS or HTTP API.

## License

MIT
