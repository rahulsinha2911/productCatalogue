#!/bin/bash
# Script to generate Connect RPC code from proto files

set -e

echo "Generating Connect RPC code from proto files..."

# Check if buf is installed
if ! command -v buf &> /dev/null; then
    echo "ERROR: buf is not installed. Installing..."
    go install github.com/bufbuild/buf/cmd/buf@latest
fi

# Generate code
buf generate proto/user/user.proto

echo "Code generation complete!"

