.PHONY: build
.NOTPARALLEL:
build: clean
	go mod tidy
	GOHOSTARCH=amd64 GOOS=linux GOARCH=arm64 go build -o bin/simplebee cmd/simplebee.go

.PHONY: clean
.NOTPARALLEL:
clean:
	rm -rf bin