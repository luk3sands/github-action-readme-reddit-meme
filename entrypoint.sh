#!/bin/sh -l

set -e

go get -v -t -d ./...
go build -v

./<your-program-name>
