#
# Makefile for juju/names
#

ifndef GOPATH
$(warning You need to set up a GOPATH.  See the README file.)
endif

PROJECT := gopkg.in/juju/names.v3
PROJECT_DIR := $(shell go list -e -f '{{.Dir}}' $(PROJECT))
PROJECT_PACKAGES := $(shell go list $(PROJECT)/... | grep -v /vendor/ | grep -v /acceptancetests/)
TEST_TIMEOUT := 600s

default: build

ifeq ($(NAMES_SKIP_DEP),true)
dep:
	@echo "skipping dep"
else
$(GOPATH)/bin/dep:
	go get -u github.com/golang/dep/cmd/dep

# populate vendor/ from Gopkg.lock without updating it first (lock file is the single source of truth for machine).
dep: $(GOPATH)/bin/dep
	$(GOPATH)/bin/dep ensure -vendor-only $(verbose)
endif

build: dep go-build

# update Gopkg.lock (if needed), but do not update `vendor/`.
rebuild-dependencies:
	dep ensure -v -no-vendor $(dep-update)

# Reformat source files.
format:
	gofmt -w -l .

go-build:
	@go build $(PROJECT_PACKAGES)

test: build
	go test $(CHECK_ARGS) -test.timeout=$(TEST_TIMEOUT) $(PROJECT_PACKAGES) -check.v
