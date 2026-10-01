---
id: collect-261001-ia-llm/ia-llm/sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b-1
title: "sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["accelerator", "copyright", "license", "licenses", "packaging"]
source: docs/RAG/collect-261001-ia-llm/sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b.md
source_anchor: ""
source_lines: [1, 103]
sha256: e58d917e1850ae7d0eb6ce71856864c7b20ebb39f5fe2bb941b0d506a4ddedd6
---

# sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b

go-unifi2mqtt is a pure-Go daemon that bridges a local UniFi Network installation (UDM / UDM-Pro / UDR / UCG / UX / UniFi OS Server, or a standalone software controller) to an MQTT broker, with optional Home Assistant MQTT auto-discovery.
It polls the console's official Network Integration API
(/proxy/network/integration/v1, API-key authenticated) for sites,
devices, per-device statistics and clients, and — where that API has
gaps — an optional classic controller API client
(/proxy/network/api/s/<site>/, cookie session) for site health, client
blocking and WLAN toggles. Values are published as MQTT topics plus HA
discovery payloads; inbound HA commands (restart device, power-cycle a
PoE port, block a client, …) are written back to the console.
Read CONCEPT.md before touching anything — it is
the implementation concept: API strategy, package layout, MQTT topic
tree, HA entity model, client filtering, polling design and the phased
roadmap. 1.0.0 is released (phases 0–8): the daemon polls the
console, publishes devices, ports, radios, the WLAN catalogue, filtered
clients with presence detection and — with the optional classic layer —
site health, per-client SSID/signal and PoE wattage, announces all of it
to Home Assistant, accepts write-back commands, and serves a diagnostic
web UI. Only the optional WebSocket accelerator (phase 9) remains.
- Language: Go 1.26+ (see go.mod / CIGO_VERSION ).
- Module path: github.com/SukramJ/go-unifi2mqtt .
- License: MIT. Every Go source file starts with:
Note this differs from the sibling project// SPDX-License-Identifier: MIT // Copyright (C) 2026 SukramJ go-mtec2mqtt , which is
LGPL-3.0-or-later because it derives from Christian Rödel'sMTECmqtt . Do not copy LGPL headers from there into this repo.
- Deployment: one static binary (CGO_ENABLED=0 ), theunifi2mqtt daemon. Also ships as a Docker image (distroless runtime) and a Home
Assistant add-on (addon/ ).
- Dependencies are deliberately minimal: golang.org/x/sync (errgroup),gopkg.in/yaml.v3 , andgithub.com/SukramJ/go-mqtt (the shared MQTT client, MIT — MQTT 5.0 by default, 3.1.1 selectable
viaTCPConfig.ProtocolVersion ). The UniFi API clients are
hand-rolled onnet/http — no third-party UniFi SDK.
- Config: YAML (config-template.yaml is the annotated reference),
overridable per-key viaUNIFI_<KEY> env vars. Loaded from--config , then$XDG_CONFIG_HOME/unifi2mqtt/config.yaml (or$APPDATA on Windows), then~/.config/unifi2mqtt/config.yaml .
cmd/unifi2mqtt/          daemon entry point (main.go)
internal/config/         YAML loader + UNIFI_* env overlay + validation
internal/unifi/          shared HTTP transport: TLS, auth, retry, rate limiting
internal/unifi/integration/  official Network Integration API v1 client
internal/unifi/classic/  classic controller API client (cookie session + CSRF)
internal/model/          API-neutral domain types (Site, Device, Client, Health)
internal/coordinator/    orchestration: poll loops, command queue, reconcile
internal/hass/           Home Assistant discovery payload builder
internal/state/          thread-safe live-value cache shared with the web UI
internal/web/            optional diagnostic web UI / HA add-on Ingress panel
internal/version/        build-info package (Version/Commit/BuildDate, ldflags)
addon/                   Home Assistant add-on packaging (Dockerfile, config.yaml, DOCS.md)
script/                  run.sh (add-on entrypoint), capture-fixtures.sh, extract-release-notes.sh
config-template.yaml     annotated reference config
CONCEPT.md               implementation concept — the design source of truth
.github/workflows/       ci.yml, docker-build-push.yml, addon-image.yml, release-on-tag.yml, codeql.yml, dependabot-auto-merge.yml
Every package above exists as of phase 7.
The MQTT transport is not part of this tree: it comes from the external
github.com/SukramJ/go-mqtt module as a regular go.mod dependency,
not an internal/ package.
All defined in the Makefile (make help lists them):
make build          # build the daemon into bin/
make run            # build then run against ./config.yaml
make test           # go test -race -count=1 -timeout=60s ./... (CGO_ENABLED=1 for race detector)
make test-cover     # tests + coverage report (coverage.out)
make vet            # go vet ./...
make fmt            # gofumpt -w . && goimports -w -local <module> .
make fmt-check      # fail if gofumpt would rewrite anything (CI gate)
make lint           # golangci-lint run ./...
make vuln           # govulncheck ./...
make licenses       # go-licenses check, forbids GPL/AGPL/LGPL(reciprocal)/MPL deps
make check          # vet + fmt-check + lint + test — the pre-commit/pre-push gate
make docker         # build a tagged container image
make release        # cross-compile linux/amd64, linux/arm64, darwin/arm64 archives into dist/
make setup          # install gofumpt/goimports/golangci-lint/govulncheck/go-licenses + git hooks
make tidy           # go mod tidy
make clean          # remove bin/, dist/, coverage.out
Run a single package's tests directly with go test ./internal/unifi/...
etc. — no special test runner beyond go test.
- License header on every .go file:SPDX-License-Identifier: MIT  - Copyright (C) 2026 SukramJ .
- No CGo in the default build (CGO_ENABLED=0 ); CGo is only
re-enabled transiently in CI/Makefile to get the race detector duringmake test .
- golangci-lint v2 config (.golangci.yaml ) enables:bodyclose ,contextcheck ,copyloopvar ,errcheck ,errorlint ,exhaustive ,gocritic ,gosec ,govet ,intrange ,makezero ,nilerr ,noctx ,prealloc ,reassign ,revive ,sloglint ,staticcheck ,thelper ,tparallel ,unconvert ,unparam ,unused ,usestdlibvars ,wastedassign .
- Formatting: gofumpt (stricter gofmt) +goimports -local github.com/SukramJ/go-unifi2mqtt for import grouping.
- Structured logging: log/slog (enforced bysloglint ).
- Package layout mirrors the data flow: config → unifi → model →
coordinator → mqtt/hass/web, each as its own internal/ package with
colocated_test.go files.
- Never log secrets. The API key, the classic password and the MQTT
password must not reach a log line, an error message, the web UI or an
MQTT payload. Redact them in any String() /diagnostic output.
- HTTP correctness is tested against golden fixtures: UniFi API
responses live as JSON under internal/unifi/integration/testdata/ and are replayed viahttptest.Server , so the decoders are anchored
without a live console.script/capture-fixtures.sh replaces them
with redacted output from a real console; see the testdata README for
the provenance rules.
- Dependency licensing is enforced by tooling, not just policy —
make licenses (go-licenses check --disallowed_types=forbidden, restricted,reciprocal ) blocks GPL/AGPL/LGPL/MPL-style dependencies
from entering the tree.
- Git hooks: make setup (ormake hooks ) pointscore.hooksPath at.githooks/ , which blocks direct commits onmain /master .
- Commit style: Conventional Commits with a scope, e.g.
feat(unifi): add integration API device client ,fix(addon): ... ,chore(deps): ... .
- Release bookkeeping — four files move together. A version bump
touches internal/version/version.go ,addon/config.yaml and theBUILD_VERSION default inaddon/Dockerfile , and
everychangelog.md entry must ALWAYS be mirrored intoaddon/CHANGELOG.md (the file Home Assistant renders in the add-on
UI's Changelog tab) — keep the two changelog files identical.
- CI (.github/workflows/ci.yml ) runs three jobs:lint (go vet +
gofumpt check),test (matrix across ubuntu/macos/windows with the
race detector),build (compiles the binary and checks the--version banner). Separate workflows publish the Docker image, the
HA add-on image, and run CodeQL + Dependabot auto-merge.
- List endpoints are not detail endpoints. GET /networks omits the
subnets andGET /devices omitsuplink /interfaces . Both fan out a
per-object detail call. Collapsing that back would compile, pass a
