.PHONY: run
run:
	@go run cmd/demo/demo.go $(args)
imports:
	@go mod tidy	