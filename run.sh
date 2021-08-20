#!/bin/sh
docker build -t api-asdf . || exit 1
docker run \
        --rm \
        -p 8080:80 \
        --name api-asdf \
        api-asdf
