# HELLO XMAS/3.0

HTTP/3.0 only server on UDP.

## Description

tbd.

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

Add flag to the [``docker-compose.yml``](./docker-compose.yml).

```bash
docker-compose up --build
```

The web server is exposed on UDP port 26. You can change it in the ``docker-compose.yml`` file.

## Build without Docker

First install [go](https://golang.org/). Next you have to install the go dependencies:

```bash
go get
```

The project is using [Go Modules](https://github.com/golang/go/wiki/Modules).

```bash
# Linux AMD64 portable
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

## Usage without Docker

Start server:
```bash
./http3only
```

## Solution

Use curl with HTTP/3.0 support:

    docker run -it --net host  --rm ymuski/curl-http3 curl -v https://localhost:6121 --http3
