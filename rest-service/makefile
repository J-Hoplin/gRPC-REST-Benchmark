BIN_DIR = bin
.PHONY: service
project := service

all: $(project)
service: $@
$(project):
	go build -o ${BIN_DIR}/server main.go