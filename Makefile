HOMEPAGE=https://github.com/mdb/seaweed-cli
PREFIX=/usr/local

COVERAGE_FILE = coverage.out

VERSION=0.0.1
TAG=v$(VERSION)

test: acceptance

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/seaweed-cli $(PREFIX)/bin/seaweed-cli

uninstall:
	rm -vf $(PREFIX)/bin/seaweed-cli

acceptance: build
	bats test

build: dependencies
	go build -o bin/seaweed-cli

dependencies:
	go get -t

.PHONY: acceptance build dependencies install test uninstall
