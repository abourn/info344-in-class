#! /usr/bin/env bash
set -e 
GOOS=linux go build
docker build -t abourn/testserver .
docker push abourn/testserver