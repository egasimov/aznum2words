test:
	go test ./...

build:
	go build cmd/aznum2words-cli/aznum2words-cli.go

.DEFAULT_GOAL := build