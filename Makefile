PKG_NAME := "queue"
PKG_HOST := "go.adenix.dev"
PKG := "${PKG_HOST}/${PKG_NAME}"
PKG_LIST := $(shell go list ${PKG}/...)

.PHONY: clean
clean: ## Remove previous coverage
	@rm -rf cover

.PHONY: test
test: ## Run unit tests
	go test -cover ./...

.PHONY: race
race: ## Run data race detector
	go test -race -short ${PKG_LIST}

.PHONY: masan
msan: ## Run memory sanitizer
	go test -msan -short ${PKG_LIST}

.PHONY: cover
cover: clean ## Generate global code coverage report
	@mkdir -p cover
	@echo 'mode: count' > cover/coverage.out
	@echo ${PKG_LIST} | xargs -n1 -I{} sh -c 'go test -covermode=count -coverprofile=cover/coverage.tmp {} && tail -n +2 cover/coverage.tmp >> cover/coverage.out' && rm cover/coverage.tmp
	go tool cover -func=cover/coverage.out

.PHONY: cover-html
cover-html: cover ## Generate global code coverage report in HTML
	go tool cover -html=cover/coverage.out

