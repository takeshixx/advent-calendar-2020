run: build
	bin/server

build:
	go build -o bin/server ./greeter_server/main.go

build-client:
	go build -o bin/client ./greeter_client/main.go

run-client: build-client
	bin/client

clean:
	rm -rf bin/