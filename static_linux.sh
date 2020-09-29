#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags "-s" -a
