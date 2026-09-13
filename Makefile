BINARY  := cetas-lite
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: test build cross run clean fmt vet smoke

test:
	go test ./... -race

vet:
	go vet ./...

smoke: build
	python3 tests/smoke/ui_smoke.py

fmt:
	gofmt -w .

build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./cmd/cetas-lite

cross:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY)-linux-amd64 ./cmd/cetas-lite
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY)-windows-amd64.exe ./cmd/cetas-lite
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY)-darwin-arm64 ./cmd/cetas-lite
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY)-darwin-amd64 ./cmd/cetas-lite

run: build
	./bin/$(BINARY) serve

clean:
	rm -rf bin
