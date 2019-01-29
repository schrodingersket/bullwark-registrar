VET_REPORT = vet.report
TEST_REPORT = tests.xml
LIB_PATHS = src/github.com src/golang.org
BIN_DIR = bin
PKG_PATHS = pkg
RM_PATHS = ${TEST_REPORT} ${VET_REPORT} ${BIN_DIR} ${PKG_PATHS}
GO = $(shell which go)
DOCKER = $(shell which docker)
REPOSITORY_NAME = bullwark-registrar
BUILD_DIR = gitlab.com/${REPOSITORY_NAME}
BIN_NAME = bullwark-registrar

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

ifeq ($(OS),Windows_NT)
    GOARCH = 386
    GOOS = windows
else
    GOOS = linux
    ifeq ($(UNAME_S),Darwin)
        GOARCH = arm
	else
        GOARCH = amd64
	endif
endif

# Build the project
all: test build

goget-test:
	GOPATH=$(CURDIR) GOOS=${GOOS} ${GO} get github.com/stretchr/testify
	GOPATH=$(CURDIR) GOOS=${GOOS} ${GO} get github.com/tebeka/go2xunit

goget:
	GOPATH=$(CURDIR) GOOS=${GOOS} ${GO} get github.com/dghubble/sling
	GOPATH=$(CURDIR) GOOS=${GOOS} ${GO} get github.com/gorilla/mux
	GOPATH=$(CURDIR) GOOS=${GOOS} ${GO} get github.com/gorilla/handlers
	GOPATH=$(CURDIR) GOOS=${GOOS} ${GO} get github.com/ghodss/yaml


build: goget
	GOPATH=$(CURDIR) GOOS=${GOOS} GOARCH=${GOARCH} ${GO} build -o ${BIN_DIR}/${BIN_NAME} ${LDFLAGS} ${BUILD_DIR} ; \

build.tar.gz: build
	tar -C ${BIN_DIR} -cvzf ${BIN_DIR}/${BIN_NAME}.tar.gz ${BIN_NAME}

build.zip: build
	zip ${BIN_DIR}/${BIN_NAME}.tar.gz ${BIN_DIR} -r

test: goget goget-test
	GOPATH=$(CURDIR) GOOS=${GOOS} ${GO} test -v ${BUILD_DIR}/... | bin/go2xunit -output ${TEST_REPORT} ; \

clean:
	-rm -rf ${RM_PATHS}

.PHONY: linux.tar.gz

# Each of the above commands has a corresponding Docker implementation, so that no dependencies need be installed.
docker:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make

docker-clean:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make clean

docker-test:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make test

docker-build:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make build

docker-build.tar.gz:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make build.tar.gz

docker-build.zip:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make build.zip

docker-goget:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make goget

docker-goget-test:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make goget-test

docker-all:
	docker run --rm -v $(CURDIR):/usr/src/${REPOSITORY_NAME} -w /usr/src/${REPOSITORY_NAME} golang:1.8 make all
