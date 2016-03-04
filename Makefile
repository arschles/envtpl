SHORT_NAME := envtpl
REPO_PATH := github.com/arschles/${SHORT_NAME}
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.9.0
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}

DOCKER_REPO ?= quay.io/
DOCKER_ORG ?= arschles
DOCKER_IMAGE ?= envtpl
DOCKER_TAG ?= devel
DOCKER_IMAGE := ${DOCKER_REPO}${DOCKER_ORG}/${DOCKER_IMAGE}:${DOCKER_TAG}

bootstrap:
	${DEV_ENV_CMD} glide install

glideup:
	${DEV_ENV_CMD} glide up

test:
	@echo "no tests yet"

build:
	${DEV_ENV_PREFIX} -e CGO_ENABLED=0 ${DEV_ENV_IMAGE} go build -a -installsuffix cgo -ldflags '-s' -o envtpl

build-cross:
	${DEV_ENV_PREFIX} -e CGO_ENABLED=0 ${DEV_ENV_IMAGE} gox

docker-build:
	docker build --rm -t ${DOCKER_IMAGE} .

docker-push:
	docker push ${DOCKER_IMAGE}
