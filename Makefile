.PHONY: install-gofumpt
install-gofumpt:
	@go install mvdan.cc/gofumpt@latest

.PHONY: install-golangci
install-golangci:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

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
	@go test -v ./...

.PHONY: setup-hooks
setup-hooks:
	git config core.hooksPath .githooks
	@echo "Git hooks configured to use .githooks/"

.PHONY: generate
generate:
	./generate.sh
