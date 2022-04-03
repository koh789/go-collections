.PHONY: help
help: ## help 表示 `make help` でタスクの一覧を確認できます
	@echo "------- タスク一覧 ------"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36mmake %-20s\033[0m %s\n", $$1, $$2}'

.PHONY: install-golangci-lint
install-golangci-lint: ## golangci-lint をインストールします。既に存在する場合はインストールしません
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.45.0

.PHONY: test-all
test-all: ## 静的解析＆テスト
	make lint
	make test

.PHONY: lint
lint: ## 静的解析 golangci-lint
	make install-golangci-lint
	golangci-lint cache clean
	golangci-lint run

.PHONY: test
test: ## test.
	@echo "+ go test..."
	go clean ./... && go test -v ./...
	@echo "+ go test clear."
