include .env
export

build:
	go build -o civi-backend-challenge -a -v

run:
	go run main.go

deps:
	go mod vendor

run-tests:
	go clean -testcache && go test ./... -cover