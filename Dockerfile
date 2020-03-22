FROM alpine

RUN apk add --no-cache mysql-client mysql-dev

COPY . /app
