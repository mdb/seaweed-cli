NAME=seaweed-cli
HOMEPAGE=https://github.com/mdb/seaweed-cli
VERSION=`cat VERSION`
TAG=v$(VERSION)
ARCH=$(shell uname -m)
PREFIX=/usr/local

test: unit acceptance

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/seaweed-cli $(PREFIX)/bin/seaweed-cli

uninstall:
	rm -vf $(PREFIX)/bin/seaweed-cli

unit: dependencies
	go test

acceptance: build
	bats test

build: dependencies
	go build -o bin/seaweed-cli

build_releases: dependencies
	go get github.com/progrium/gh-release
	mkdir -p build/Linux  && GOOS=linux  go build -ldflags "-X main.Version $(VERSION)" -o build/Linux/$(NAME)
	mkdir -p build/Darwin && GOOS=darwin go build -ldflags "-X main.Version $(VERSION)" -o build/Darwin/$(NAME)

dependencies:
	go get -t

release: build_releases
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(ARCH).tgz -C build/Linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(ARCH).tgz -C build/Darwin $(NAME)

release: build_releases
	gh-release create mdb/$(NAME) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: acceptance build build_releases dependencies install test uninstall unit
