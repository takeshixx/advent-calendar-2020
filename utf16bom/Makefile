run: build
	bin/server 4444

build:
	go build -o bin/server main.go

client-run: client-build
	bin/client

client-build:
	go build -o bin/client client/client.go

clean:
	rm -rf bin/
