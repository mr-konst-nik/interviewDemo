.PHONY: run
run:
	@go run cmd/demo/demo.go $(args)

.PHONY: deps
deps:
	go mod tidy

.PHONY: lint
lint:
	golint ./...
	golangci-lint run ./... 		