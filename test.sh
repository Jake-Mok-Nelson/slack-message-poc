#! /bin/bash

go clean -testcache
SLACK_API_TOKEN=${SLACK_API_TOKEN} go test ./...
