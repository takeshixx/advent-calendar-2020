FROM golang:1.15.5 as builder
WORKDIR /src
RUN echo "JAHr7BBNh0fp3IlbC2DItfWOV0CYIj3J" > flag.txt

COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o socks .
RUN chmod -R go+rX .
RUN apt update && apt install curl
RUN curl --output /proxy.sh https://raw.githubusercontent.com/takeshixx/tools/master/net/proxy.sh

WORKDIR /app
USER 1001
COPY --from=builder /src/socks .
COPY --from=builder /src/anonymous ./anonymous
COPY --from=builder /src/permitted ./permitted
COPY --from=builder /src/flag.txt /app/permitted/flag.txt
COPY --from=builder /proxy.sh /proxy.sh
EXPOSE 9000
CMD ["./socks"]