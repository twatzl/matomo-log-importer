#!/bin/sh

go build -o appname main.go
GOOS=windows GOARCH=amd64 go build -o appname.exe main.go
