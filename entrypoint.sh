#!/bin/sh -l

set -e

go get -v -d ./...
go build -v -o github-action-commit-workflow-changes

./github-action-commit-workflow-changes
