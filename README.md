# gRPC Sample
Sample of using gRPC.

## Requirements
- Docker 25.0.3
- Docker Compose v2.24.6

## Usage
1. Set up docker containers.
  - Server will start automatically.
```bash
$ docker compose up -d
```

2. Run request from client.
```bash
$ docker compose exec client go run main.go
```
