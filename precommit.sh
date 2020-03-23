#!/bin/sh
gofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
[ -z "$gofiles" ] && exit 0
unformatted=$(gofmt -l $gofiles)
[ -z "$unformatted" ] || echo "needs formatting: $unformatted"

# tests
go test -v -race $(go list ./... | grep -v /vendor/)
RESULT=$?
[ $RESULT -ne 0 ] && exit 1
exit 0
