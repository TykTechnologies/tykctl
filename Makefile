tests:
	go test ./... -count=1 -v -cover
lint:
	golangci-lint run ./...