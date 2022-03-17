include .env
export

build:
	go build -o civi-backend-challenge -a -v

run:
	go run main.go

deps:
	go mod vendor

swagger-docs:
	go get -u github.com/swaggo/swag/cmd/swag@v1.6.7 && swag init --parseDependency --parseInternal -g main.go

run-tests:
	go clean -testcache && go test ./... -cover