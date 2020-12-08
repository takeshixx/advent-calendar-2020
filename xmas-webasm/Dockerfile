FROM golang:alpine
WORKDIR /srv
RUN mkdir /srv/logs/
RUN mkdir /run/nginx/
COPY www /srv/www
COPY nginx.conf /etc/nginx/nginx.conf
RUN apk update && apk add make git nginx
COPY go.mod go.sum main.go Makefile /srv/
RUN go mod download
RUN make build
ENV XMAS_SECRET "MOCtzhkRF27k9MunPy1RwMN873hTBTmK"
EXPOSE 80
CMD nginx