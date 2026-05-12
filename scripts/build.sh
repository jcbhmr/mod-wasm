#!/usr/bin/env bash
set -Eeuo pipefail

temp_dir=$(mktemp --directory)
trap cleanup SIGINT SIGTERM ERR EXIT
cleanup() {
    trap - SIGINT SIGTERM ERR EXIT
    rm -rf "$temp_dir"
}

wkg wit fetch
wkg wit build --output "$temp_dir/package.wasm"
tinygo build -target wasip2 -wit-package "$temp_dir/package.wasm" .
