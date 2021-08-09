run:
	@echo off
	nodemon -e go --exec go run . --signal SIGTERM
build:
	@echo off
	del /f popomepost.exe
	go generate
	go build -ldflags "-H windowsgui" -o popomepost.exe
dev:
	cd vue && yarn serve