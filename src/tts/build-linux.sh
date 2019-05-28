#!/bin/sh

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build tts.go
echo "build complete start zip..."
zip -q -r tts.zip tts README.md tts.cnf tts.service