#!/bin/sh

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/alog alog.go
echo "build complete start zip..."
zip -q -r alog.zip build/alog build/alog.conf build/alog.service