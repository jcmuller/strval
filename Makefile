.PHONY: all
all: test

.PHONY: test
test: install-gotestsum
	@mkdir -p tmp/output/results tmp/output/coverage
	@bin/gotestsum \
		--junitfile tmp/output/results/unit-tests.xml \
		--post-run-command "make coverage" \
		-- \
		-coverprofile=tmp/output/coverage/coverage.out \
		./...

.PHONY: coverage
coverage:
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
	@bin/golangci-lint run

.PHONY: install-golang-ci-lint
install-golang-ci-lint:
	@GOBIN=$(PWD)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: install-gotestsum
install-gotestsum:
	@GOBIN=$(PWD)/bin go install gotest.tools/gotestsum@latest
