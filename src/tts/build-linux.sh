GOOS=linux GOARCH=amd64 go build tts.go

zip -q -r tts.zip tts README.md tts.cnf tts.service