check:
    go fmt ./... && \
    go vet ./... && \
    golangci-lint run ./...
