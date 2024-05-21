#!/bin/bash

# Set the output directory for the binaries
OUTPUT_DIR="bin"

# Build for Linux (amd64)
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_DIR/file_server_linux_amd64 main.go

# Build for Linux (arm64)
echo "Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -o $OUTPUT_DIR/file_server_linux_arm64 main.go

# Build for macOS (amd64)
echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o $OUTPUT_DIR/file_server_darwin_amd64 main.go

# Build for macOS (arm64)
echo "Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o $OUTPUT_DIR/file_server_darwin_arm64 main.go

# Build for Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_DIR/file_server_windows_amd64.exe main.go

# Build for Android
echo "Building for Android..."
GOOS=android GOARCH=arm64 go build -o $OUTPUT_DIR/file_server_android_arm64 main.go