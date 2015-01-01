.PHONY: all test build

NAME=gorsync

all: test

test:
	@go test

build:
	go build -o ${NAME}
