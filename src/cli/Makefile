.PHONY: build
.NOTPARALLEL:
build: clean test
	go mod tidy
	go build -o bin/beekeeper cmd/beekeeper.go

.PHONY: test
.NOTPARALLEL:
test:
	go test ./...

.PHONY: clean
.NOTPARALLEL:
clean:
	rm -rf bin

.PHONY: fmt
.NOTPARALLEL:
fmt:
	go fmt ./...