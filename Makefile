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
	@go test ./...

mock: ## Build mocks
	@mockgen -source=internal/usecase/account/repository.go -destination=internal/usecase/account/mock/repository.go