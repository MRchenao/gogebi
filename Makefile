.PHONY:build
build:
	@swag init
	@go build gebi

.PHONY:run
run:
	@swag init
	@go run gebi

.PHONY: proto
proto:
	protoc  --proto_path=./proto --go_out=plugins=grpc:./proto --go_opt=paths=source_relative ./proto/*.proto

.PHONY:test
test:
	@go test ./test -v

.PHONY:linux
linux:
	@set GOOS=linux
	@set CGO_ENABLED=0
	@set GOHOSTOS=linux

.PHONY: help
help:
	@echo "make build - go build"
	@echo "make proto - build proto"
	@echo "make run - go run"
	@echo "make test - go test"
	@set GOOS=linux
	@set CGO_ENABLED=0
	@set GOHOSTOS=linux