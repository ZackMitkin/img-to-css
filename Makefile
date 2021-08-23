
run-wasm: wasm start

wasm: export GOOS=js
wasm: export GOARCH=wasm

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

wasm:
	go build -o ./web/public/main.wasm ./cmd/wasm/main.go
#cp "$(GOPATH)/misc/wasm/wasm_exec.js" ./static

start:
	go run ./cmd/server/main.go