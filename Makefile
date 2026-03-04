.PHONY: fmt test gen_golangci_lint build

fmt:
	goimports -w .
	go fmt ./...

test:
	go test -v ./...

gen_golangci_lint:
	golangci-lint custom -v
	@echo Use ./custom-gcl

build:
	go build -o build/logcheck cmd/logcheck/logcheck.go
