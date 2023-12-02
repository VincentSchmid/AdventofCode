day := "02"

init:
    cd ./Day{{day}} && \
    go mod init main

check:
    cd ./Day{{day}} && \
    go fmt ./... && \
    go vet ./... && \
    golangci-lint run ./...

run:
    cd ./Day{{day}} && \
    go run main.go

playground:
    cd ./playground && \
    go run main.go
