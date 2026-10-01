---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8-3
title: "noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8.md
source_anchor: ""
source_lines: [123, 173]
sha256: 65dd7ede6d7d9104c1da34cb31dae3d633c7cae3b627088449e9d0ab1352dd80
---

# noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8

Step 1: Write failing command and prompt tests Add command tests using injectable newAuthStore ,newClientWithAPIKey , andpromptAPIKey seams. Assert that successfulunifi login prompts once, validates withGET /self/sites , saves only after validation, cleans legacy entries, and emitsauth_method: saved_api_key without the sentinel key. Assert validation failure leaves a pre-existing saved key unchanged. Assertunifi login --file-fallback passestrue only to that save call. Assert non-TTY input returnsvalidation_failed withUNIFI_API_KEY guidance. Assertunifi logout deletes local current and legacy entries without starting a server request. Update status tests forsaved_api_key andenvironment_api_key only.if strings.Contains(output.String(), "api-key-not-for-output") { t.Fatalf("login output leaked API key: %q", output.String()) } if store.saveCalls != 0 { t.Fatalf("failed validation saved key %d times", store.saveCalls) }
- 
Step 2: Run CLI auth tests to verify they fail Run: go test ./internal/cliExpected: FAIL because only nested password-session commands and no hidden prompt exist.
- 
Step 3: Implement the top-level command flow and hidden prompt Add golang.org/x/term as a direct dependency.promptAPIKey must requireterm.IsTerminal(int(in.Fd())) , printAPI key: to the command output, callterm.ReadPassword , print one newline, trim surrounding whitespace, and reject empty input without including it in the error. Keep the function behind a package variable or small injected dependency so tests never need a real TTY.Register newLoginCmd() andnewLogoutCmd() at the root.unifi login loads non-secret config, prompts, builds aninteractive_api_key validation client, callsValidate , then callsstore.Save(cfg.BaseURL(), key, allowFileFallback) . Only after both operations succeed, emit safe metadata withsaved_api_key .unifi logout callsstore.Delete(cfg.BaseURL()) and emitslogged_out ; it must not construct an HTTP client. Keepauth only as the read-onlyauth status command and removeauth login andauth logout .
- 
Step 4: Run focused command tests and inspect help output Run: gofmt -w internal/cli/login.go internal/cli/login_test.go internal/cli/prompt.go internal/cli/auth.go internal/cli/auth_test.go internal/cli/root.go && go test ./internal/cli && go run ./cmd/unifi --helpExpected: tests PASS and root help lists login andlogout ; no password-login or API-key argument flag appears. Task 5 removes the remaining obsolete--no-session-write flag.
- 
Step 5: Commit the user-facing auth workflow git add internal/cli/login.go internal/cli/login_test.go internal/cli/prompt.go internal/cli/auth.go internal/cli/auth_test.go internal/cli/root.go go.mod go.sum
git commit -m "feat: add persistent API-key login"
Files:
- Modify: internal/cli/root.go
- Modify: internal/cli/context.go
- Modify: internal/cli/configcmd.go
- Modify: internal/cli/cli_test.go
- Modify: README.md
- Modify: configs/config.example.yaml
Interfaces:
- 
Consumes: the command names and auth-method strings from Task 4.
- 
Produces: non-secret config show data containinghost ,port ,insecure ,site ,safe_mode , andtimeout only.
- 
Step 1: Write failing help and config-output tests Replace tests that expect nested auth login/logout, secret redaction placeholders, or --no-session-write . Add tests that root help includeslogin andlogout ,auth --help exposes onlystatus ,login --help exposes--file-fallback but not an API-key flag, andconfig show has nousername ,password , orapi_key keys. Use a config with legacy fields and verify the failure text has no sentinel secret.for _, forbidden := range []string{"username", "password", "api_key", "--no-session-write"} { if strings.Contains(out, forbidden) { t.Fatalf("obsolete auth surface %q in output:\n%s", forbidden, out) } }
- 
Step 2: Run CLI surface tests to verify they fail Run: go test ./internal/cliExpected: FAIL because the root still registers session-era flags and config output still includes credential fields.
- 
Step 3: Remove obsolete flags and update user-facing documentation Remove flagNoSessionWrite fromroot.go and itsloadRuntime behavior. Remove all credential fields fromredactedConfig andprintData ordering. Update the README and example config to show non-secret controller configuration followed by:unifi login
unifi auth status --json
unifi logoutDocument the hidden prompt, restart-persistent native stores, explicit unifi login --file-fallback , the warning that fallback is opt-in protected local state, controller scoping,401 re-login behavior, andUNIFI_API_KEY as a process-only CI/script override. Remove every username/password/session instruction and update the documented Go requirement to Go 1.26.5.
- 
Step 4: Run documentation-adjacent tests and inspect forbidden live surfaces Run: gofmt -w internal/cli/root.go internal/cli/context.go internal/cli/configcmd.go internal/cli/cli_test.go && go test ./internal/cli && rg -n "UNIFI_USERNAME|UNIFI_PASSWORD|auth login|auth logout|no-session-write|username: admin" README.md configs/config.example.yaml internal/cli -g '!**/*_test.go'Expected: tests PASS; the rg command prints no obsolete public auth instruction. Historical design and plan documents are intentionally excluded from this check.
- 
Step 5: Commit the surface and documentation cleanup git add internal/cli/root.go internal/cli/context.go internal/cli/configcmd.go internal/cli/cli_test.go README.md configs/config.example.yaml
git commit -m "docs: document API-key-only login"
Files:
- Verify: all files changed by Tasks 1–5
- Modify only if verification exposes a concrete defect in those files.
Interfaces:
- 
Consumes: the complete API-key-only CLI.
- 
Produces: evidence that the repository builds, tests, and does not retain an active password/session authentication path.
- 
Step 1: Run the full automated verification suite Run: go test ./... && go vet ./... && go build -o /tmp/unifi-api-key-plan ./cmd/unifiExpected: all commands exit successfully and the built binary is created at /tmp/unifi-api-key-plan .
- 
Step 2: Verify command registration and safe error behavior with the built binary Run: tmp_config="$(mktemp /tmp/unifi-api-key-config.XXXXXX)" printf 'host: controller.example\n' > "$tmp_config" /tmp/unifi-api-key-plan --help /tmp/unifi-api-key-plan login </dev/null /tmp/unifi-api-key-plan --config "$tmp_config" config show rm "$tmp_config" Expected: help includes top-level login andlogout ; non-TTY login exits with a redactedUNIFI_API_KEY guidance error; config show succeeds when the temporary config has a host value. Create the temporary config with a user-only mode and remove it after the check.
- 
Step 3: Review the migration diff for secrets and obsolete behavior Run: git diff --check fcf5b40..HEAD git diff fcf5b40..HEAD -- internal/config internal/authstore internal/client internal/cli README.md configs/config.example.yaml git status --short Expected: no whitespace errors; no key literal or password-login endpoint remains in active code; only task-scoped files are staged or committed. Treat test sentinel strings as test data, not credentials.
- 
Step 4: Commit only a concrete verification fix, if one was required If the preceding steps reveal and fix a defect, run the affected focused test plus go test ./... , then commit only the changed task files:git add <verified-task-files> git commit -m "fix: complete API-key auth migration" If no defect was found, do not create an empty commit.
