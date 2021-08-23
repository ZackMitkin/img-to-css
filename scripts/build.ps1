set GOOS js
set GOARCH wasm

cd ./cmd/wasm

go build -o ../../static/main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./cmd/wasm