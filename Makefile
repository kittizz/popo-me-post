run:
	go run .
build:
	@echo off
	del /f popomepost.exe
	go generate
	go build -ldflags "-H windowsgui" -o popomepost.exe
