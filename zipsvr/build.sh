#!/usr/bin/env bash
set -e 
echo "building linux executable"
GOOS=linux go build
docker build -t abourn/zipserver .
docker push abourn/zipserver
go clean 