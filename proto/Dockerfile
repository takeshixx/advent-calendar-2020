FROM golang:1.15.5 as builder
WORKDIR /src
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o socks .
RUN chmod -R go+rX .

FROM scratch
WORKDIR /app
USER 1001
COPY --from=builder /src/advent-proto .
EXPOSE 9000
CMD ["./socks"]