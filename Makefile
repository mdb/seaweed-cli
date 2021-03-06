NAME=seaweed
HOMEPAGE=https://github.com/mdb/seaweed-cli
VERSION=0.1.2
TAG=v$(VERSION)
ARCH=$(shell uname -m)
PREFIX=/usr/local
VETARGS?=-all

all: lint vet test

test: unit acceptance

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/$(NAME) $(PREFIX)/bin/$(NAME)

uninstall:
	rm -vf $(PREFIX)/bin/$(NAME)

unit: dependencies
	go test

acceptance: build
	bats test

build: dependencies
	go build -ldflags "-X main.version=$(VERSION)" -o bin/$(NAME)

build_releases: dependencies
	mkdir -p build/Linux  && GOOS=linux  go build -ldflags "-X main.version=$(VERSION)" -o build/Linux/$(NAME)
	mkdir -p build/Darwin && GOOS=darwin go build -ldflags "-X main.version=$(VERSION)" -o build/Darwin/$(NAME)
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(ARCH).tgz -C build/Linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(ARCH).tgz -C build/Darwin $(NAME)

dependencies:
	go get -t
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go get github.com/golang/lint/golint

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

# NOTE: TravisCI will auto-deploy a GitHub release when a tag is pushed
tag:
	git tag $(TAG)
	git push origin $(TAG)

lint: dependencies
	golint -set_exit_status

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "go tool vet $(VETARGS)"
	@go tool vet $(VETARGS) . ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: acceptance build build_releases dependencies install test uninstall unit
