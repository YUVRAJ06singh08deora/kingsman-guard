#!/bin/bash

PROJECT_ROOT="$(dirname "$(realpath "$0")")/.."

echo "Compile Binaries? (yes/no)"
read answer

if [ "$answer" == "yes" ]; then
    echo "Windows Or Linux? (win/lin)"
    read computer
    if [ "$computer" == "win" ]; then
    ./build_windows.sh
    fi
    if [ "$computer" == "lin" ]; then
    ./build_linux.sh
    fi
fi

cd "$PROJECT_ROOT/infra"

terraform init
terraform apply

echo "Terraform apply has been executed."
