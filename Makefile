tests:
	go test ./... -count=1 -v
lint:
	golangci-lint run ./...