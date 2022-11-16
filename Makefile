tests:
	go test ./... -count=1 -v -cover

generate-mocks:
	go generate -v ./...

lint:
	golangci-lint run ./...

generate-docs:
	cd docs && [ ! -e docs.md ] || rm docs.md && go run docs.go


