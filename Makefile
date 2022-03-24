.PHONY: parser test

all: parser test

parser:
	@pigeon yue.peg > yue.go

test:
	@go test -v
