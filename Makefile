format:
	go fmt $$(go list ./...)

vet:
	go vet $$(go list ./...)

test:
	go test $$(go list ./...) -coverprofile tmp/coverage.out

coverage:
	go tool cover -func tmp/coverage.out

coverage-html:
	go tool cover -html=tmp/coverage.out

.PHONY: format vet test coverage coverage-html
