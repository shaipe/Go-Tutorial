GOOS=linux GOARCH=amd64 go build trs.go

zip -q -r trs.zip trs

scp trs.zip root@192.168.4.136:/srv -p366eclinux.
