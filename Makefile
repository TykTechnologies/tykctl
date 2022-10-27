tests:
	go test ./... -count=1 -v
generate-mocks:
	go generate -v ./...