# XMAS Greetings

A simple gRPC service where the client has to call a specific function. The `.proto` file will be provided.

## Description

```html

```

## Solution

Compile protobuf:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    xmasgreetings/xmasgreetings.proto
```

`greeter_client/main.go` is an example implementation to solve this challenge.

## Bulding & Running

Run the server:

```bash
make
```

Run the client with default settings:

```bash
make run-client
```

Run the client with the actual solution:

```bash
bin/client xmas.rip:12 xmas
```