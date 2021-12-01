REPO_NAME := "queue"
REPO_PATH := "adenix"
REPO_HOST := "github.com"
PKG := "${REPO_HOST}/${REPO_PATH}/${REPO_NAME}"
PKG_LIST := $(shell go list ${PKG}/...)

.PHONY: clean-cover clean-build clean test race msan cover cover-html build

clean-cover: ## Remove previous coverage
	@rm -rf cover

clean-build: ## Remove previous build
	@rm -rf bin

clean: clean-test clean-build

test: ## Run unit tests
	@go test -cover ./...

race: ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

cover: clean-cover ## Generate global code coverage report
	@mkdir -p cover
	@echo 'mode: count' > cover/coverage.out
	@echo ${PKG_LIST} | xargs -n1 -I{} sh -c 'go test -covermode=count -coverprofile=cover/coverage.tmp {} && tail -n +2 cover/coverage.tmp >> cover/coverage.out' && rm cover/coverage.tmp
	@go tool cover -func=cover/coverage.out

cover-html: cover ## Generate global code coverage report in HTML
	@go tool cover -html=cover/coverage.out

build: clean-build ## Build the binary file
	@go build $(flags) -o bin/ ${PKG_LIST}
