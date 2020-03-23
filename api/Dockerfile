# docker 向けに linux build する container
FROM golang:1.14 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
COPY . /app
WORKDIR /app
RUN make

# 実際の server
FROM alpine

RUN apk add --no-cache mysql-client mysql-dev

COPY --from=builder /app /app
