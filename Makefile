GREEN=\033[36m
DEFAULT=\033[0m

.PHONY: test vendor

help: ## Display help screen
	@echo "Usage:"
	@echo "\tmake [COMMAND]"
	@echo "Commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%s\t$(DEFAULT)%s\n", $$1, $$2}'

run: ## Run application
	@go run cmd/api/main.go

tidy: ## Downloads go dependencies
	@go mod tidy

vendor: ## Copy of all packages needed
	@go mod vendor

test: ## Run the tests of the project
	@go test -v -cover -p=1 -covermode=count -coverprofile=cov.out  ./...

mock: ## Build mocks
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=internal/usecase/account/iface.go -destination=internal/usecase/account/mock/iface.go
	@~/go/bin/mockgen -source=internal/adapter/iface.go -destination=internal/adapter/mock/iface.go
	@~/go/bin/mockgen -source=internal/adapter/account/iface.go -destination=internal/adapter/account/mock/iface.go