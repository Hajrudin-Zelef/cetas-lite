BINARY  := cetas-lite
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)
CROSS   := linux-amd64 linux-arm64 windows-amd64 windows-arm64 darwin-amd64 darwin-arm64

.PHONY: test build cross desktop run clean fmt fmt-check vet smoke ci

DESKTOP_LDFLAGS := -s -w -H windowsgui -X main.version=$(VERSION)

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

# Application bureau Windows : binaire GUI (sans console), double-clic = ouvre Cetas.
desktop:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "$(DESKTOP_LDFLAGS)" -o bin/$(BINARY)-desktop-windows-amd64.exe ./cmd/cetas-lite
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -trimpath -ldflags "$(DESKTOP_LDFLAGS)" -o bin/$(BINARY)-desktop-windows-arm64.exe ./cmd/cetas-lite

run: build
	./bin/$(BINARY) serve

clean:
	rm -rf bin
