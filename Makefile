BINARY  := cetas-lite
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)
CROSS   := linux-amd64 linux-arm64 windows-amd64 windows-arm64 darwin-amd64 darwin-arm64

.PHONY: test build cross run clean fmt fmt-check vet smoke ci

test:
	go test ./... -race

vet:
	go vet ./...

fmt:
	gofmt -w .

fmt-check:
	@out="$$(gofmt -l .)"; if [ -n "$$out" ]; then echo "gofmt requis:"; echo "$$out"; exit 1; fi

ci: fmt-check vet test

smoke: build
	python3 tests/smoke/ui_smoke.py

build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./cmd/cetas-lite

cross:
	mkdir -p bin
	@set -e; for t in $(CROSS); do \
	  os=$${t%-*}; arch=$${t#*-}; ext=; \
	  if [ "$$os" = windows ]; then ext=.exe; fi; \
	  echo "build $$t"; \
	  CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY)-$$t$$ext ./cmd/cetas-lite; \
	done

run: build
	./bin/$(BINARY) serve

clean:
	rm -rf bin
