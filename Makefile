.PHONY: lint lint-fix

lint: 
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./... --fix

generate-v1:
	openapi bundle openapi/v1/main.yaml -o openapi/v1/bundle.yaml
	oapi-codegen --config oapi-codegen.yaml openapi/v1/bundle.yaml
	go mod tidy
