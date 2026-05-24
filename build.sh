#!/bin/bash

set -e

APP_NAME="lotus"

# Linux amd64
GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

# Windows amd64
GOOS=windows GOARCH=amd64 go build -o ${APP_NAME}.exe

echo "Binaries successfully built:"
ls -lh ${APP_NAME}*