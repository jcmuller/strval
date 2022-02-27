.PHONY: all
all: test

.PHONY: test
test:
	@go test ./...

.PHONY: test
fmt:
	@go fmt ./...

.PHONY: test
vet: fmt
	@go vet ./...

.PHONY: test
lint: vet install-golang-ci-lint
	@golangci-lint run

.PHONY: install-golang-ci-lint
install-golang-ci-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
