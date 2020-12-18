FROM golang:1.15.5 as builder
WORKDIR /src
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
RUN chmod -R go+rX .

FROM scratch
WORKDIR /app
USER 1001
COPY --from=builder /src/static ./static
COPY --from=builder /src/templates ./templates
COPY --from=builder /src/http3only .
EXPOSE 8080
CMD ["./http3only", "--flag", "ltim2wewL7p5SgS1ig2Rbc0zBwXQWS8P"]