export GO111MODULE=on
GO_SRC=$(shell find . -path ./.build -prune -false -o -name \*.go)

lint: ./.golangcilint.yaml
	./bin/golangci-lint --version || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.46.2
	./bin/golangci-lint --config ./.golangcilint.yaml run ./...

test: $(GO_SRC)
	go test -v -race -cover -coverpkg ./... -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: all
all: test lint

.PHONY: mod
mod:
	go get -u
	go mod tidy

.PHONY: clean
clean:
	-rm coverage.txt