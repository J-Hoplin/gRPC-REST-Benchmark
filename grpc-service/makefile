BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client

SHELL := bash
SHELL_VERSION = $(shell echo $$BASH_VERSION)
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
UNAME := $(shell uname -s)
VERSION_AND_ARCH = $(shell uname -rm)
ifeq ($(UNAME),Darwin)
	OS = macos ${VERSION_AND_ARCH}
else ifeq ($(UNAME),Linux)
	OS = linux ${VERSION_AND_ARCH}
else
	$(error OS not supported by this Makefile)
endif


.DEFAULT_GOAL := help
.PHONY: service
project := service

all: $(project)

service: $@

$(project):
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/server main.go


about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"
