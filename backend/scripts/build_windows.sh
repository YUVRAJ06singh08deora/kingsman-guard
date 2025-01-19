#!/bin/bash

PROJECT_ROOT="$(dirname "$(realpath "$0")")/.."

cd "$PROJECT_ROOT/cmd"

for DIR in */; do
    cd "$PROJECT_ROOT/cmd/$DIR"
    GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
    cd "$PROJECT_ROOT"
    
    # Create dist directory if it doesn't exist
    mkdir -p dist
    
    # Remove the existing ZIP file if it already exists
    if [ -f "dist/${DIR%/}.zip" ]; then
        echo "Removing existing ZIP file: dist/${DIR%/}.zip"
        rm "dist/${DIR%/}.zip"
    fi
    
    # Create the new ZIP file
    powershell -Command "Compress-Archive -Path 'cmd/${DIR}bootstrap' -DestinationPath 'dist/${DIR%/}.zip'"

    # Remove the bootstrap binary after zipping it
    rm "cmd/${DIR}bootstrap"
done

echo "All binaries compiled and zipped successfully."
