CURRENT_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')

BUILD_FLAGS := -ldflags="-X 'main.Version=v0.1.0' -X 'main.BuildDate=${CURRENT_DATE}'"

build:
	go build ${BUILD_FLAGS} -o quicktool cmd/main.go

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build ${BUILD_FLAGS} -o quicktool-darwin-amd64 cmd/main.go

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build ${BUILD_FLAGS} -o quicktool-darwin-arm64 cmd/main.go

install: build
	sudo mv quicktool /usr/local/bin/f

.PHONY: docs
docs:
	go run ./docs/main.go

test:
	go test ./...
