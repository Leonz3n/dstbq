VERSION=$(shell git describe --tags --always)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
 	       --go_out=paths=source_relative:./internal \
	       internal/config/config.proto

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

