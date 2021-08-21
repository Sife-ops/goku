#!/bin/sh

. ./.env
HOST_PORT=${HOST_PORT:-80} # override with PORT=<port>

dev(){
    DOCKER_PORT=${HOST_PORT} go run main.go
}

prod(){
    docker build -t yw/goku:dev . || exit 1
    docker rmi -f $(docker images -f "dangling=true" -q)
    docker run \
            -e DOCKER_PORT=${DOCKER_PORT} \
            -e POSTGRES_DBNAME=${POSTGRES_DBNAME} \
            -e POSTGRES_HOST=${POSTGRES_HOST} \
            -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
            -e POSTGRES_PORT=${POSTGRES_PORT} \
            -e POSTGRES_USER=${POSTGRES_USER} \
            -e ACCESS_TOKEN_SECRET=${ACCESS_TOKEN_SECRET} \
            -e REFRESH_TOKEN_SECRET=${REFRESH_TOKEN_SECRET} \
            -e REDIS_DSN=${REDIS_DSN} \
            --rm \
            -p ${HOST_PORT}:${DOCKER_PORT} \
            --name goku \
            yw/goku:dev
}

case $1 in
    dev)
        dev
        ;;
    prod)
        prod
        ;;
    *)
        exit 1
        ;;
esac
