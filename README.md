# Go in WebAssembly Example 

## Compile for WebAssembly

```
GOOS=js GOARCH=wasm go build -o web/Build_wasm_js cmd/wasm/main.go
```

## Copy wasm js dependency
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
```

## Serve
```
go run cmd/web/main.go -dir web/
```

## Links

* https://github.com/golang/go/wiki/WebAssembly
