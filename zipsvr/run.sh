#!/usr/bin/env bash
docker run -d \
-p 443:443 \
--name zipsvr \
-v /Users/Adam/go/src/github.com/abourn/info344-in-class/zipsvr/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
abourn/zipserver