VERSION=$(shell git describe --tags --always)
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

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

.PHONY: proto
# proto
proto:
	protoc -I=$(ROOT_DIR)/proto \
				 --go_out=$(ROOT_DIR)/proto \
				 --go_opt=module=github.com/Leonz3n/kulery/proto \
				 $(ROOT_DIR)/proto/kulery.proto $(ROOT_DIR)/proto/config.proto

