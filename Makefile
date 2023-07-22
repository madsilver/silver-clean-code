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
	@go test -v -cover -p=1 -covermode=count -coverprofile=coverage.out  ./...

mock: ## Build mocks
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=internal/usecase/account/repository.go -destination=internal/usecase/account/mock/repository.go
	@~/go/bin/mockgen -source=internal/adapter/controller/account/controller.go -destination=internal/adapter/controller/account/mock/usecase.go
	@~/go/bin/mockgen -source=internal/infra/db/db.go -destination=internal/infra/db/mock/db.go
	@~/go/bin/mockgen -source=internal/adapter/context_server.go -destination=internal/adapter/mock/context_server.go
	@~/go/bin/mockgen -source=internal/usecase/transaction/repository.go -destination=internal/usecase/transaction/mock/repository.go