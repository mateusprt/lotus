#!/bin/bash

set -e

APP_NAME="lotus"

# Linux amd64
GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}
zip ${APP_NAME}-linux.zip ${APP_NAME}
rm ${APP_NAME}

# Windows amd64
GOOS=windows GOARCH=amd64 go build -o ${APP_NAME}.exe
zip ${APP_NAME}-windows.zip ${APP_NAME}.exe
rm ${APP_NAME}.exe

# macOS amd64
GOOS=darwin GOARCH=amd64 go build -o ${APP_NAME}
zip ${APP_NAME}-mac.zip ${APP_NAME}
rm ${APP_NAME}

echo "Binários gerados e compactados:"
ls -lh *.zip