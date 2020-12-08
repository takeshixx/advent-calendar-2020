FROM golang:alpine
WORKDIR /srv
RUN apk update && apk add make git
COPY greeter_server /srv/greeter_server
COPY xmasgreetings /srv/xmasgreetings
COPY go.mod go.sum Makefile /srv/
RUN go mod download
RUN make build
ENV XMAS_SECRET "WkBDHJryg3wqWsHRA6JJZe5G2rZEAdKH"
CMD /srv/bin/server