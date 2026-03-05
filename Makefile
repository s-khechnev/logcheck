.PHONY: fmt test gen_golangci_lint build

fmt:
	goimports -w .
	go fmt ./...

test:
	go test -v ./...

gen_golangci_lint:
	golangci-lint custom -v
	@echo "Use ./custom-gcl"

build:
	go build -o build/logcheck cmd/logcheck/logcheck.go

deps:
	@echo "Installing main dependencies..."
	go mod download
	@echo "Installing testdata dependencies..."
	@find internal/logcheck -path "*/testdata/*/go.mod" -exec dirname {} \; | while read dir; do \
		echo "Processing $$dir"; \
		cd $$dir && go mod download && go mod tidy && cd -; \
	done
