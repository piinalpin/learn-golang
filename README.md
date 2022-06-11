# Learn RestFul API Golang using Gin

## Overview

![Golang](https://www.seekpng.com/png/detail/399-3990193_building-a-go-web-app-from-scratch-to.png)

This project is documentation of my learning Golang with the best practice based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout/), by Kyle Quest.

## Features
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/index.html)
- Custom Soft Deleted
   - `CreatedAt`
   - `CreatedBy`
   - `UpdatedAt`
   - `DeletedAt`
- Base Api Response
   ```json
      {
         "timestamp": "",
         "response_key": "",
         "message": "",
         "data": {}
      }
   ```
- Custom Exception and Exception Handler
- [Golang Enumerations](https://levelup.gitconnected.com/implementing-enums-in-golang-9537c433d6e2)
- [Golang Generic](https://go.dev/doc/tutorial/generics)
- [Package Validator](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)
- Containerized Support
- Read Environment Variables

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

### Run container using Docker

```bash
docker run -d --name learn_go -p 8080:8080 --env-file=dev.env --network my-network learn-rest-api
```

### Run container using Podman (Arch Linux)

Install `dnsmasq` and `cni-plugins` using package manager

```bash
sudo pacman -S dnsmasq cni-plugins
```
Looking for `cni` binary installed.

```bash
ls -l /opt/cni/bin
ls -l /usr/lib/cni
```

Using git, clone `https://github.com/containers/dnsname.git` see [README_PODMAN.md](https://github.com/containers/dnsname/blob/main/README_PODMAN.md) and change line at `Makefile`

```bash
...
LIBEXECDIR ?= ${PREFIX}/cni/bin
...
```

Install with command `sudo make install PREFIX=/opt`

Change line again at `Makefile`

```bash
...
LIBEXECDIR ?= ${PREFIX}/lib/cni
...
```

Install with command `sudo make install PREFIX=/usr`

Enable `cni` in `/etc/containers/containers.conf`.

```bash
...
cni_plugin_dirs = [
    "/usr/local/libexec/cni",
    "/usr/libexec/cni",
    "/usr/local/lib/cni",
    "/usr/lib/cni",
    "/opt/cni/bin",
]
...
```
Create new network `podman network create my-network`

Update configuration `my-network` at `$HOME/.config/cni/net.d/my-network.conflist`

```bash
{
  "cniVersion": "0.4.0",
  "name": "my-network",
  "plugins": [
     ...
     {
        "type": "dnsname",
        "domainName": "my-network",
        "capabilities": {
           "aliases": true
        }
     }
     ...
  ]
}
```

Run container application

```bash
podman run -d --name learn_go -p 8080:8080 --env-file=dev.env --network my-network learn-rest-api
```

## Reference
- [Dasar Pemrograman Golang](https://dasarpemrogramangolang.novalagung.com/)
- 