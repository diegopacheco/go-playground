#!/usr/bin/env bash
set -euo pipefail

go run .
echo
echo "Running tests..."
go test ./...
