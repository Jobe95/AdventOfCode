#!/bin/bash
# Install script for AOC CLI

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INSTALL_DIR="$HOME/.local/bin"
BINARY="$INSTALL_DIR/aoc"

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Build directly to install directory
cd "$SCRIPT_DIR"
go build -o "$BINARY" ./cmd/aoc

if [ ! -f "$BINARY" ]; then
    echo "✗ Build failed"
    exit 1
fi

chmod +x "$BINARY"

echo "✓ AOC CLI installed to $BINARY"
echo ""

# Check if directory is in PATH
if [[ ":$PATH:" == *":$INSTALL_DIR:"* ]]; then
    echo "✓ $INSTALL_DIR is in your PATH"
    echo "You can now run: aoc setup"
else
    echo "⚠️  $INSTALL_DIR is not in your PATH"
    echo ""
    echo "Add this line to your ~/.zshrc:"
    echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
    echo ""
    echo "Then run: source ~/.zshrc"
fi
