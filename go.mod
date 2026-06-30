module go.jcbhmr.com.internal/mod-wasm

go 1.26.0

toolchain go1.26.2

tool (
	go.bytecodealliance.org/cmd/wit-bindgen-go
	go.jcbhmr.com.internal/mod-wasm/internal/cmd/build
	go.jcbhmr.com/wasmpkgtools/cmd/wkg
)

require (
	go.bytecodealliance.org/cm v0.3.0
	golang.org/x/mod v0.36.0
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/klauspost/compress v1.18.6 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/regclient/regclient v0.11.5 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	github.com/tetratelabs/wazero v1.9.0 // indirect
	github.com/ulikunitz/xz v0.5.15 // indirect
	github.com/urfave/cli/v3 v3.3.3 // indirect
	go.bytecodealliance.org v0.7.0 // indirect
	go.jcbhmr.com/crossexec v1.1.1 // indirect
	go.jcbhmr.com/wasmpkgtools v0.15.1-0.20260524070841-580d2f50203b // indirect
	golang.org/x/sys v0.45.0 // indirect
)
