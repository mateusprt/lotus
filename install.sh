#!/bin/bash

set -e

LOTUS_VERSION="${1:-1}"
INSTALL_DIR="$HOME/.lotus/versions/$LOTUS_VERSION"
BIN_DIR="$HOME/.lotus/bin"

# Detect OS (Linux only)
OS="$(uname -s)"
if [[ "$OS" != "Linux" ]]; then
    echo "This installer is only supported on Linux."
    exit 1
fi

URL="https://github.com/mateusprt/lotus/releases/download/$LOTUS_VERSION/lotus"

mkdir -p "$INSTALL_DIR"
mkdir -p "$BIN_DIR"
cd "$INSTALL_DIR"

echo "Downloading Lotus $LOTUS_VERSION..."
curl -L -o lotus "$URL"

chmod +x lotus

# Create symlink for the active version
ln -sf "$INSTALL_DIR/lotus" "$BIN_DIR/lotus"

cd ~

# Add to PATH in the current shell and suggest adding to shell profile
if [[ ":$PATH:" != *":$BIN_DIR:"* ]]; then
    export PATH="$BIN_DIR:$PATH"
    echo -e "\n\033[1;33m==============================\nLotus added to PATH. Please, restart your terminal.\n==============================\033[0m\n"
    echo "Lotus $LOTUS_VERSION successfully installed in $INSTALL_DIR"
    fi
echo -e "\n\033[1;33m==============================\nPlease restart your terminal. After that, run 'lotus --version' to verify the installed version.\n==============================\033[0m\n"