.PHONY: build clean run test podman docker

BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
BIN := goapp
VERSION := 1.0.0
TARGET_OS := linux
TARGET_ARCH := amd64

# TARGET LIST \
TARGET_OS			TARGET_ARCH \
anroid					arm \
linux					amd64 \
linux					arm \
linux					arm64 \
linux					386 \
windows					amd64 \
windows					386 \

build: clean
	env CGO_ENABLED=0 GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -o ./bin/${BIN} -ldflags="-X 'main.buildVersion=${VERSION}' -X 'main.buildDate=${BUILD_DATE}'"

run:
	go run main.go

clean:
	go clean
	rm -rf ./bin/*

podman: clean build
	podman build -t docker.io/piinalpin/learn-goapp:latest --build-arg APP_NAME=goapp  . -f build/Dockerfile

docker: clean build
	docker build -t piinalpin/learn-goapp:latest --build-arg APP_NAME=goapp  . -f build/Dockerfile