tests:
	go test ./... -count=1 -v -cover

generate-mocks:
	go generate -v ./...

lint:
	golangci-lint run ./...


generate-docs:
	cd docs && rm docs.md && go run docs.go


