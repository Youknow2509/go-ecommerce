#!/bin/bash
# Script to generate Go code from user.proto using protoc and protoc-gen-go

PROTO_DIR="$(dirname "$0")/../proto"
OUT_DIR="${PROTO_DIR}"
PROTO_FILE="${PROTO_DIR}/user.proto"

# Ensure protoc and protoc-gen-go are installed
if ! command -v protoc &> /dev/null; then
  echo "protoc could not be found. Please install Protocol Buffers compiler."
  exit 1
fi
if ! command -v protoc-gen-go &> /dev/null; then
  echo "protoc-gen-go could not be found. Please install it with: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
  exit 1
fi

# Generate Go code from proto
protoc --go_out=paths=source_relative:${OUT_DIR} --go-grpc_out=paths=source_relative:${OUT_DIR} ${PROTO_FILE}

if [ $? -eq 0 ]; then
  echo "Proto generation successful."
else
  echo "Proto generation failed."
  exit 1
fi
