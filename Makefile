.DEFAULT_GOAL:=help

help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Format
format:  ## Format code with `go fmt`
	go fmt $$(go list ./...)

vet:  ## Vet with `go fmt`
	go vet $$(go list ./...)

##@ Build
build: ## Build application
	go build -o dist/

##@ Test
test: ## Run tests in all packages and generate coverage report
	go test $$(go list ./...) -coverprofile tmp/coverage.out

test-verbose: ## Run tests in all packages and generate coverage report
	go test $$(go list ./...) -coverprofile tmp/coverage.out -v

cov: ## Print coverage report
	go tool cover -func tmp/coverage.out

cov-html: ## View coverage report as HTML
	go tool cover -html=tmp/coverage.out

.PHONY: format vet make test test-verbose coverage coverage-html
