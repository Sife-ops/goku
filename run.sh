#!/bin/sh
. ./.env
export PORT=${PORT:-80}
docker build -t api-asdf . || exit 1
docker run \
        -e APP_PORT=${APP_PORT} \
        -e POSTGRES_DBNAME=${POSTGRES_DBNAME} \
        -e POSTGRES_HOST=${POSTGRES_HOST} \
        -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
        -e POSTGRES_PORT=${POSTGRES_PORT} \
        -e POSTGRES_USER=${POSTGRES_USER} \
        --rm \
        -p ${PORT}:${APP_PORT} \
        --name api-asdf \
        api-asdf
