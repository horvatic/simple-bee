.PHONY: build
.NOTPARALLEL:
build: clean test
	go mod tidy
	GOHOSTARCH=amd64 GOOS=linux GOARCH=arm64 go build -o bin/simplebee cmd/simplebee.go

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