export PKGS=$(shell go list ./...)

test:
	@go test -v -cover -race $(PKGS)

build:
	@go build -v -o ./cmd/billing-engine ./cmd/main.go

run:
	make build
	@./cmd/billing-engine

race: 
	@go run -race ./cmd/main.go