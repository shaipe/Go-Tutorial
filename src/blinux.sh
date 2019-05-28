#!/bin/sh

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o web myweb.go

# CGO_ENABLED=0
# GOOS=linux
# GOARCH=amd64

# go build -o web myweb.go

zip -q -r web.zip web