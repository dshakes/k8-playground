

VERSION:=$(shell git describe --tags --always --dirty)
ifeq ($(VERSION),)
VERSION:=noversion
endif

OS := linux
OS_MAC := darwin
ARCH := amd64
REPO_NAME := skaffold-argocd-api
API_SERVICE_NAME := go-rest-microservice


# The pipeline docker image working directory
export PROJECT_PATH :=/go/src/$(GITHUB_HOST)/$(GITHUB_OWNER)/$(GITHUB_REPO)

SRC := $(shell find * -type d -name vendor -prune -o -type f -name "*.go" -print)
DOCKER = docker run --rm -v $(PWD):$(PROJECT_PATH) -w $(PROJECT_PATH)

# Commands and Other Helper Variables
export ENV_COMPILE := CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH)
export ENV_COMPILE_LOCAL := CGO_ENABLED=0 GOOS=$(OS_MAC) GOARCH=$(ARCH)
export CMD_COMPILE := go build -a -installsuffix cgo -ldflags '-s' -o
export CMD_LINT := gometalinter --config=./metalinter-config.json ./...
export CMD_FAST_LINT := gometalinter --config=./metalinter-config.json --fast --deadline=30s ./...
export CMD_TEST := go test ./...

deletetag:
	git tag -d $(tagName)

newtag:
	git tag -a $(tagName) -m $(msg)

goreleaser:
	goreleaser --rm-dist

pipeline-new: test newtag goreleaser

lint:
	$(CMD_LINT)

test:
	$(CMD_TEST)

build:
	$(ENV_COMPILE) $(CMD_COMPILE) target/$(API_SERVICE_NAME)-$(OS)-$(ARCH)

pipeline: build

build-local:
	$(ENV_COMPILE_LOCAL) $(CMD_COMPILE) target/local/$(API_SERVICE_NAME)-$(OS_MAC)-$(ARCH)


local: test build-local


export ENCRYPTION_BASE :=mybase 
export ENCRYPTION_SALT :=mysalt
export ENCRYPTION_BITS :=12345

run: 
	@docker run -p 8000:8000 -t docker.io/iotcloudstack/golang-rest-api:latest

run-local:
	ulimit -n 1024; appenv="development" go run -ldflags "-X ${PKG}/pkg/version.VERSION=${VERSION}" main.go

version:
	@echo ${VERSION}




down:
	docker-compose down

run-vss-api:
	docker-compose -f docker-compose.yml run --rm api
	docker-compose -f docker-compose.yml down
