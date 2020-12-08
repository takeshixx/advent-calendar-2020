run: build
	http_server -no-auth -root www -port 9090

build: clean
	GOOS=js GOARCH=wasm go build -o www/xmas.wasm main.go
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js www/wasm_exec.js

clean:
	rm www/xmas.wasm www/wasm_exec.js || true

build-crypter:
	go build -o bin/crypter ./crypter/crypter.go