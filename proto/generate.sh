#!/bin/bash
set -e

# Remove vendor if it exists
[ -d "vendor" ] && rm -rf vendor

# Set correct GOOGLEAPIS_PATH
GOOGLEAPIS_PATH=${GOOGLEAPIS_PATH:-"$HOME/googleapis"}

# Generate code
protoc -I . \
  -I "$GOOGLEAPIS_PATH" \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  --descriptor_set_out=../envoy/desc.desc \
  --include_imports \
  --include_source_info \
  api_gaurd.proto
