# Learn REST API Golang using Gin

## How to Run?

```bash
go run main.go
```

## Work with containerize

### Build binary and images

```bash
CGO_ENABLED=0 go build -o bin/learn-rest-api

docker build -t learn-rest-api:latest . -f build/Dockerfile
```

### Run container

```bash
docker run -d --name learn_go -p 8080:8080 --env-file=dev.env --network my-network learn-rest-api
```