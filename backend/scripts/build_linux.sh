#!/bin/bash

PROJECT_ROOT="$(dirname "$(realpath "$0")")/.."

cd "$PROJECT_ROOT/cmd"

for DIR in */; do
    cd "$PROJECT_ROOT/cmd/$DIR"
    GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
    cd "$PROJECT_ROOT"
    mkdir -p dist
    zip -j "dist/${DIR%/}.zip" "cmd/${DIR}bootstrap"
    rm "cmd/${DIR}bootstrap"
done

echo "All binaries compiled and zipped successfully."
