# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

export CGO_ENABLED=1
export GO111MODULE=on

VERSION                     ?= $(shell cat ./VERSION)

DOCKER_REPOSITORY           ?= onosproject/
ONOS_REFERENCE_ADAPTER_VERSION ?= latest

## Docker labels. Only set ref and commit date if committed
DOCKER_LABEL_VCS_URL        ?= $(shell git remote get-url $(shell git remote))
DOCKER_LABEL_VCS_REF        = $(shell git rev-parse HEAD)
DOCKER_LABEL_BUILD_DATE     ?= $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")
DOCKER_LABEL_COMMIT_DATE    = $(shell git show -s --format=%cd --date=iso-strict HEAD)

DOCKER_EXTRA_ARGS           ?=
DOCKER_BUILD_ARGS ?= \
        ${DOCKER_EXTRA_ARGS} \
        --build-arg org_label_schema_version="${VERSION}" \
        --build-arg org_label_schema_vcs_url="${DOCKER_LABEL_VCS_URL}" \
        --build-arg org_label_schema_vcs_ref="${DOCKER_LABEL_VCS_REF}" \
        --build-arg org_label_schema_build_date="${DOCKER_LABEL_BUILD_DATE}" \
        --build-arg org_opencord_vcs_commit_date="${DOCKER_LABEL_COMMIT_DATE}" \
        --build-arg org_opencord_vcs_dirty="${DOCKER_LABEL_VCS_DIRTY}"

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then mkdir -p build && cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

.PHONY: images
images: # @HELP build simulators image
images: reference-adapter-docker

reference-adapter-docker: # @HELP build the referencee adapter docker container
reference-adapter-docker:
	docker build . -f Dockerfile \
	$(DOCKER_BUILD_ARGS) \
	-t ${DOCKER_REPOSITORY}reference-adapter:${ONOS_REFERENCE_ADAPTER_VERSION}

.PHONY: build
build: # @HELP build the Go binaries (default)
build:
	go build -o build/_output/reference-adapter ./cmd/reference-adapter

mod-update: # @HELP Download the dependencies to the vendor folder
	go mod tidy
	go mod vendor
mod-lint: mod-update # @HELP ensure that the required dependencies are in place
	# dependencies are vendored, but not committed, go.sum is the only thing we need to check
	bash -c "diff -u <(echo -n) <(git diff go.sum)"


test: # @HELP run the unit tests and source code validation  producing a golang style report
test: mod-lint build linters license
	go test -race github.com/onosproject/config-adapter/pkg/...
	go test -race github.com/onosproject/config-adapter/cmd/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: mod-lint build linters license
	TEST_PACKAGES=github.com/onosproject/config-adapter/pkg/... ./build/build-tools/build/jenkins/make-unit

publish: # @HELP publish version on github and dockerhub
	./build/build-tools/publish-version ${VERSION}

jenkins-publish: jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/build-tools/release-merge-commit

all: test

clean:: # @HELP remove all the build artifacts
	go clean -testcache github.com/onosproject/config-adapter/...
