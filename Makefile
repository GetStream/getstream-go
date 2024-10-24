.PHONY: install-gofumpt
install-gofumpt:
	@go install mvdan.cc/gofumpt@latest

.PHONY: install-golangci
install-golangci:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: build
build:
	@go build ./...

.PHONY: lint
lint:
	./lint.sh

.PHONY: fmt
fmt: install-gofumpt
	@gofmt -s -w .
	@gofumpt -l -w .

.PHONY: test
test:
	@go test -v ./pkg/...

.PHONY: generate
generate:
	./generate.sh
