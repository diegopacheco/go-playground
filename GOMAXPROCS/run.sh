#!/bin/bash

echo "Running Docker container with resource limits..."
docker run --rm --cpus=2 --memory=512m gomaxprocs-demo
docker ps -q --filter "ancestor=gomaxprocs-demo" | xargs -r docker kill

echo "Running Docker container without resource limits..."
docker run --rm gomaxprocs-demo
docker ps -q --filter "ancestor=gomaxprocs-demo" | xargs -r docker kill