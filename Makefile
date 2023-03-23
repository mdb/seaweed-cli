NAME=seaweed
HOMEPAGE=https://github.com/mdb/seaweed-cli
VERSION=0.1.2
TAG=v$(VERSION)
ARCH=$(shell uname -m)
PREFIX=/usr/local

.DEFAULT_GOAL := test

test: vet test-fmt unit acceptance
.PHONY: test

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/$(NAME) $(PREFIX)/bin/$(NAME)
.PHONY: install

uninstall:
	rm -vf $(PREFIX)/bin/$(NAME)
.PHONY: uninstall

unit:
	go test
.PHONY: test

acceptance: build
	bats test
.PHONY: acceptance

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/$(NAME)
.PHONY: build

build_releases:
	mkdir -p build/Linux  && GOOS=linux  go build -ldflags "-X main.version=$(VERSION)" -o build/Linux/$(NAME)
	mkdir -p build/Darwin && GOOS=darwin go build -ldflags "-X main.version=$(VERSION)" -o build/Darwin/$(NAME)
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(ARCH).tgz -C build/Linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(ARCH).tgz -C build/Darwin $(NAME)
.PHONY: build_releases

vet:
	go vet $(SOURCE)
.PHONY: vet

test-fmt:
	test -z $(shell go fmt $(SOURCE))
.PHONY: test-fmt

release: build_releases
	go get github.com/aktau/github-release
	github-release release \
		--user mdb \
		--repo seaweed-cli \
		--tag $(TAG) \
		--name "$(TAG)" \
		--description "seaweed-cli version $(VERSION)"
	ls release/*.tgz | xargs -I FILE github-release upload \
		--user mdb \
		--repo seaweed-cli \
		--tag $(TAG) \
		--name FILE \
		--file FILE
.PHONY: release

# NOTE: TravisCI will auto-deploy a GitHub release when a tag is pushed
tag:
	git tag $(TAG)
	git push origin $(TAG)
.PHONY: tag
