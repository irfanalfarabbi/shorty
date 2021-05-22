.PHONY: all dep test run

UNIT_TEST_PATHS := . ./api

dep: #tobeadded

test:
	@go test -race ${UNIT_TEST_PATHS} -v -coverprofile=coverage.out
	@go tool cover -func=coverage.out

testlocal:
	@go test -failfast ${UNIT_TEST_PATHS}

run:
	go run shorty.go
