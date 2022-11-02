tests:
	go test ./... -count=1 -v -cover

generate-mocks:
	go generate -v ./...

lint:
	golangci-lint run ./...

