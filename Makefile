Version := $(shell date "+%Y%m%d%H%M")
GitCommit := $(shell git rev-parse HEAD)
DIR := $(shell pwd)
LDFLAGS := -s -w -X main.Version=$(Version) -X main.GitCommit=$(GitCommit)

.PHONY: run
run: build
	./bin/wizard-enhance --debug

.PHONY: build
build: build-orm
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/wizard-enhance cmd/*.go

.PHONY: build-orm
build-orm:
	orm internal/service/model/*.yaml
	gofmt -s -w internal/service/model/*.go

.PHONY: dist
dist: build-orm
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o bin/wizard-enhance-linux cmd/*.go