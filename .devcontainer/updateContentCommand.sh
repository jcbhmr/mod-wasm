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

echo "# Install 'cargo binstall'"
curl -L --proto '=https' --tlsv1.2 -sSf https://raw.githubusercontent.com/cargo-bins/cargo-binstall/main/install-from-binstall-release.sh | bash
source ~/.bashrc

echo "# Install 'wkg' and 'wasm-tools'"
cargo binstall --no-confirm wkg wasm-tools
