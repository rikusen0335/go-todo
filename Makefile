# note: call scripts from /scripts
export PATH := $(PWD)/bin:$(PATH)

TARGET_ENV := dev
TASK_PARALLEL := 1

install:
	$(eval BIN:=$(abspath bin))
	mkdir -p ./bin
	go mod download golang.org/x/tools

build:
	go build -o ./bin/go-todo-cli ./cmd/go-todo