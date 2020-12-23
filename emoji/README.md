# Emoji

A TCP server that speaks a "emoji" protocol: a emoji "puzzle/quiz".

## Description

TODO

## Build with Docker

```bash
docker-compose build
```

## Run with Docker

Add flag to the [``docker-compose.yml``](./docker-compose.yml).

```bash
docker-compose up --build
```

The web server is exposed on TCP port 26. You can change it in the ``docker-compose.yml`` file.

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
./emoji
```

## Solution

The quiz has three steps:
- in the first step you need to send the emoji same as the received emoji
- the second step is a scissors stone paper game with emojis
- in the third step you need to guess to correct XMAS emoji sequence. Possible emojis are: 🎅🎄☃️🎁❄️⛄🎍🛷🦌🤶