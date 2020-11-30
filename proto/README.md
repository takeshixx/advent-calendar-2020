# proto

A custom TLS server written in go.

## Description

**SPOILER:** The proto challenge offers the token if you negotiate the correct [TLS ALPN protocol](https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation).

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

Add flag to ``flag.txt``.

```bash
docker-compose up --build
```

The TLS server is exposed on TCP port 443. You can change it in the ``docker-compose.yml`` file.

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
./proto --help
```

Start server:
```bash
./proto
```

## Solution

```base64
IyBEaXNjb3ZlciBhdmFpbGFibGUgcHJvdG9jb2xzOgpmb3IgcCBpbiAkKGNhdCB+L2dpdC9TZWNM
aXN0cy9EaXNjb3ZlcnkvV2ViLUNvbnRlbnQvcmFmdC1zbWFsbC13b3Jkcy50eHQpOyBkbyAgZWNo
byAidGVzdCIgfCBvcGVuc3NsIHNfY2xpZW50IC1hbHBuICIkcCIgbG9jYWxob3N0Ojg0NDMgMj4v
ZGV2L251bGwgIHwgZ3JlcCAiQUxQTiBwcm90b2NvbCI7IGRvbmUKIyBHZXQgdG9rZW46Cm9wZW5z
c2wgc19jbGllbnQgLWFscG4geG1hcyBsb2NhbGhvc3Q6ODQ0Mw==
```
