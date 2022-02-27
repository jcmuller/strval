.PHONY: all
all: test

.PHONY: test
test: install-gotestsum
	@mkdir -p tmp/output/results tmp/output/coverage
	@gotestsum \
		--junitfile tmp/output/results/unit-tests.xml \
		-- \
		-coverprofile=tmp/output/coverage/coverage.out \
		./...

.PHONY: coverage
coverage: test
	@go tool cover \
		-html=tmp/output/coverage/coverage.out \
		-o tmp/output/coverage/coverage.html

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

.PHONY: install-gotestsum
install-gotestsum:
	@go install gotest.tools/gotestsum@latest
