.PHONY:
.SILENT:

build:
	go build -o ./.bin/cmd cmd/main.go

run: build
	./.bin/cmd

format:
	${call colored, formatting is running...}
	go vet ./...
	go fmt ./...