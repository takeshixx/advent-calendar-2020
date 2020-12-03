FROM golang:1.15.5 as builder
WORKDIR /src
RUN echo "9lslhc67y8WqKAQDbUlAvlNaFqNqB3Z6" > flag.txt
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
RUN chmod -R go+rX .

FROM scratch
WORKDIR /app
USER 1001
COPY --from=builder /src/proto .
COPY --from=builder /src/flag.txt .
EXPOSE 8443
CMD ["./proto"]