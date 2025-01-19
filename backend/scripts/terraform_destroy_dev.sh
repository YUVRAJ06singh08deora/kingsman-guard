#!/bin/bash

PROJECT_ROOT="$(dirname "$(realpath "$0")")/.."

cd "$PROJECT_ROOT/infra"

terraform destroy

echo "Terraform destroy has been executed."
