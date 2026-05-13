#!/usr/bin/env bash
set -Eeuo pipefail

wkg wit fetch
rm -rf ./internal/wit-bindgen-go/
go tool wit-bindgen-go generate --world exports --out ./internal/wit-bindgen-go/ ./wit/
