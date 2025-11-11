LOCAL_BIN := $(CURDIR)/bin
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint
GO_TEST=$(LOCAL_BIN)/gotest
GO_TEST_ARGS=-race -v -tags=integration_test ./...

.PHONY: lint
lint:
	echo 'Running linter on files...'
	$(GOLANGCI_BIN) run \
	--config=.golangci.yaml \
	--sort-results \
	--max-issues-per-linter=0 \
	--max-same-issues=0

.PHONY: test
test:
	echo 'Running tests...'
	${GO_TEST} ${GO_TEST_ARGS}


bin-deps: .bin-deps

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5 && \
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/gotest@v0.0.6 && \
	GOBIN=$(LOCAL_BIN) go install go.uber.org/mock/mockgen@latest


build:
	go mod tidy
	go build -o ./bin/gomodinfo ./cmd/gomodinfo/