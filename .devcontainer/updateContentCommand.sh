#!/usr/bin/env bash
set -Eeuo pipefail

echo "# Install latest version of TinyGo"
temp_dir_path=$(mktemp --directory)
pushd "$temp_dir_path"
tinygo_version=$(curl --location https://ungh.cc/repos/tinygo-org/tinygo/releases/latest | jq --raw-output '.release.tag[1:]')
wget "https://github.com/tinygo-org/tinygo/releases/download/v${tinygo_version}/tinygo_${tinygo_version}_amd64.deb"
sudo dpkg --install "tinygo_${tinygo_version}_amd64.deb"
popd
rm -rf "$temp_dir_path"

echo "# Install Wasmtime"
curl https://wasmtime.dev/install.sh -sSf | bash
