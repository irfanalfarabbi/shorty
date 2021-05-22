.PHONY: all dep test run compile build push

IMAGE := irfanalfarabbi/shorty
UNIT_TEST_PATHS := . ./handler ./service

all: compile build push

dep: #tobeadded

test:
	@go test -race ${UNIT_TEST_PATHS} -v -coverprofile=coverage.out
	@go tool cover -func=coverage.out

testlocal:
	@go test -failfast ${UNIT_TEST_PATHS}

run:
	go run shorty.go

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build shorty.go

build:
	docker build -t $(IMAGE):latest -f ./Dockerfile .

push:
	docker push $(IMAGE):latest
