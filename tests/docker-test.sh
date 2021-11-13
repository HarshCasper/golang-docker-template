#!/usr/bin/env bash
set -euo pipefail

docker-compose up -d

echo "Running HTTP API smoke tests..."
hurl --test tests/*.hurl

docker-compose stop
