---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8-1
title: "noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agentic", "decode", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8.md
source_anchor: ""
source_lines: [1, 62]
sha256: 4effbaa51717c5a742929cb05b86101c9f442dd9b07f83a1ddc108e3466e0d0a
---

# noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8

Implemented historical plan. Use the root README.md for current commands and configuration.
For agentic workers: REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (- [ ]) syntax for tracking.
Goal: Replace password and cookie-session authentication with persistent, controller-scoped API-key authentication that supports interactive login and non-interactive automation.
Architecture: internal/authstore replaces the session package and owns serialized API-key records in the native credential store plus an explicit protected-file fallback. internal/client resolves either UNIFI_API_KEY or the saved key and sends only X-API-KEY; it removes a failed saved key on 401. The CLI provides top-level login and logout commands and a hidden TTY prompt, while configuration remains non-secret.
Tech Stack: Go 1.26.5, Cobra, github.com/zalando/go-keyring, golang.org/x/term, gopkg.in/yaml.v3, Go net/http.
- Authentication is API-key-only; no password request, cookie restoration, CSRF handling, or legacy session-auth path remains usable.
- Persist a key only after a successful read-only validation request.
- Resolve credentials in this exact order: UNIFI_API_KEY , saved controller-scoped key, thennot_authenticated with aunifi login hint.
- Never print an API key, include one in normal config, accept one in a flag or positional argument, or expose it in an error, JSON envelope, or log.
- Save state in macOS Keychain, Windows Credential Manager, or Linux Secret Service. Permit a protected-file fallback only when unifi login --file-fallback is explicitly supplied.
- Scope saved state by normalized scheme, host, and port. An environment override never writes or deletes saved state.
- Existing username ,password , and YAMLapi_key settings must fail with migration guidance that does not repeat their values.
- Preserve unrelated working-tree files and existing generated artifacts; stage only files changed for the task being committed.
- internal/config/config.go — non-secret connection config plus explicit detection of deprecated credential settings.
- internal/config/config_test.go — configuration defaults, overrides, and non-leaking migration errors.
- internal/authstore/store.go — controller-scoped API-key record, keyring integration, fallback writes, and legacy-session cleanup.
- internal/authstore/store_test.go — storage behavior using an injected in-memory keyring and temporary state directories.
- internal/session/store.go andinternal/session/store_test.go — delete after their cleanup-compatible replacement exists; no live session storage remains.
- internal/client/client.go — API-key-only request construction, credential source tracking, validation, and saved-key invalidation.
- internal/client/client_test.go — HTTP-level tests for header use, source precedence, missing auth, and401 behavior.
- internal/apperr/apperr.go andinternal/apperr/apperr_test.go — add the stablenot_authenticated error code.
- internal/cli/auth.go — retain onlyauth status behavior and shared safe auth metadata.
- internal/cli/login.go — top-level interactivelogin and locallogout commands.
- internal/cli/prompt.go — hidden TTY API-key prompt with an injectable test seam.
- internal/cli/root.go ,internal/cli/helpers.go ,internal/cli/context.go ,internal/cli/configcmd.go — command registration and removal of session-era flags and secret fields.
- internal/cli/auth_test.go ,internal/cli/cli_test.go ,internal/cli/login_test.go — command behavior, prompt, redaction, and help coverage.
- go.mod ,go.sum — add the direct terminal-input dependency only.
- README.md ,configs/config.example.yaml — public API-key-only setup and persistent-state documentation.
Files:
- Modify: internal/config/config.go
- Modify: internal/config/config_test.go
- Modify: internal/apperr/apperr.go
- Modify: internal/apperr/apperr_test.go
Interfaces:
- 
Produces: config.Load that never accepts or populates password/API-key configuration.
- 
Produces: apperr.NotAuthenticated with the string value"not_authenticated" .
- 
Consumed by later tasks: config.Load(path string) (config.Config, error) rejects legacy config keys andUNIFI_USERNAME /UNIFI_PASSWORD without exposing values.
- 
Step 1: Write failing configuration and error-code tests Add a table-driven TestLoadRejectsLegacyCredentials covering YAMLusername ,password , andapi_key , plus non-emptyUNIFI_USERNAME andUNIFI_PASSWORD . For every case, assertconfig.Load fails, the error contains"no longer supported" and"unifi login" , and it does not contain the sentinel secret. Add a control case showingUNIFI_API_KEY does not changeconfig.Config or makeconfig.Load fail. Add anapperr test thatapperr.New(apperr.NotAuthenticated, "not authenticated") has the expected code.if strings.Contains(err.Error(), "legacy-secret") { t.Fatalf("migration error leaked secret: %v", err) } if !apperr.Is(apperr.New(apperr.NotAuthenticated, "not authenticated"), apperr.NotAuthenticated) { t.Fatal("missing NotAuthenticated code") }
- 
Step 2: Run the new tests to verify they fail Run: go test ./internal/config ./internal/apperrExpected: FAIL because legacy credentials are still accepted and NotAuthenticated does not exist.
- 
Step 3: Implement safe migration checks without breaking the intermediate client build Remove all credential environment overrides. Before unmarshalling into Config , decode the YAML top-level mapping and reject only the keysusername ,password , andapi_key with a fixed error such as:func legacyCredentialError(name string) error { return fmt.Errorf("config %q is no longer supported; remove it and run 'unifi login'", name) } Reject non-empty UNIFI_USERNAME andUNIFI_PASSWORD with equivalent fixed messages. AddNotAuthenticated Code = "not_authenticated" to the application error constants. Do not inspect or interpolate any credential value. Keep the three deprecatedConfig fields temporarily, but leave them unpopulated; Task 3 removes them in the same change that removes all consumers, preserving a buildable commit boundary.
- 
Step 4: Run focused tests and format changed Go files Run: gofmt -w internal/config/config.go internal/config/config_test.go internal/apperr/apperr.go internal/apperr/apperr_test.go && go test ./internal/config ./internal/apperrExpected: PASS.
- 
Step 5: Commit the independently testable configuration boundary git add internal/config/config.go internal/config/config_test.go internal/apperr/apperr.go internal/apperr/apperr_test.go
git commit -m "refactor: remove credential config fields"
Files:
- Create: internal/authstore/store.go
- Create: internal/authstore/store_test.go
Interfaces:
- 
Consumes: config.Config.BaseURL() values and the existinggo-keyring dependency.
- 
Produces: type Store interface { Load(controller string) (apiKey string, found bool, err error) Save(controller, apiKey string, allowFileFallback bool) error Delete(controller string) error } func NewStore(options Options) *KeyringStore func NormalizeController(controller string) (string, error)
- 
Consumed by later tasks: client.NewWithStore ,unifi login , andunifi logout .
- 
