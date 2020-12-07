FROM golang:alpine
WORKDIR /srv
RUN apk update && apk add make git
COPY go.mod go.sum main.go Makefile /srv/
RUN go mod download
RUN make build
ENV XMAS_SECRET "ddCADpw1yVUy0Xk8Lv0BjnhuD7Wuv8oT"
CMD /srv/bin/server 4444