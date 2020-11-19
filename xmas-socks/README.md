# xmas-socks

xmas-socks is a simple portable parallel secure SOCKS server written in Go.

## Description

This is the most secure SOCKS proxy server ever to protect the XMAS socks:

![XMAS Socks](permitted/sock.png)
![XMAS Socks](permitted/sock.png)

Only source IPs allowed in the access control list can get the SOCKS.

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

Add flag to ``flag.txt``.

```bash
docker-compose up --build
```

The sock server is exposed on TCP port 24.

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

Print usage:
```bash
./socks --help
```

Start server:
```bash
./socks
```

## Solution

    IyBSZXBsYWNlIGxvY2FsaG9zdCAyNCB3aXRoIFNPQ0tTIHNlcnZlciBhZGRyZXNzCmVjaG8gLWUg
    InN0cmljdF9jaGFpblxuW1Byb3h5TGlzdF1cbnNvY2tzNSAJMTI3LjAuMC4xIDI0IiA+IHByb3h5
    Y2hhaW5zLmNvbmYKcHJveHljaGFpbnMgLWYgcHJveHljaGFpbnMuY29uZiBjdXJsIC12diAtLXNv
    Y2tzNSBsb2NhbGhvc3Q6OTAwMCBodHRwOi8vZXhhbXBsZS5jb20KcHJveHljaGFpbnMgLWYgcHJv
    eHljaGFpbnMuY29uZiBjdXJsIC12diAtLXNvY2tzNSBsb2NhbGhvc3Q6OTAwMCBodHRwOi8vZXhh
    bXBsZS5jb20vZmxhZy50eHQ=
