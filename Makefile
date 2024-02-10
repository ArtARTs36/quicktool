build:
	go build -o quicktool cmd/main.go

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o quicktool-darwin-amd64 cmd/main.go

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o quicktool-darwin-arm64 cmd/main.go

install: build
	sudo mv quicktool /usr/local/bin/f

.PHONY: docs
docs:
	go run ./docs/main.go

test:
	go test ./...
