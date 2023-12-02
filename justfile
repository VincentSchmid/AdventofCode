set shell := ["zsh", "-uc"]

day := "01"

init:
    cd ./src/Day{{day}} && \
    go mod init main

check:
    cd ./src/Day{{day}} && \
    go fmt ./... && \
    go vet ./... && \
    golangci-lint run ./...

run:
    cd ./src/Day{{day}} && \
    go run main.go

playground:
    cd ./src/playground && \
    go run main.go

# just --set day "03" run
