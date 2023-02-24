test:
	go test ./...

build:
	go build cmd/cliapp/aznum2words-cli.go

.DEFAULT_GOAL := build