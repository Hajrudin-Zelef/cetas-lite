---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8-2
title: "noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8.md
source_anchor: ""
source_lines: [63, 122]
sha256: 4fb6a793aa42eeb8724d03c77142b3e1afb3de2ea1a46ef347faa821a1fbb8d5
---

# noahjenkins-unifi-cli-blob-head-docs-superpowers-plans-2026-07-29-api-key-only-a-27cde3d8

Step 1: Write failing storage tests around API-key records, not sessions Copy the existing injected-keyring test fixture into internal/authstore/store_test.go . Write tests for save/load normalization, controller isolation, native-store preference, explicit fallback-only saving, fallback directory0700 and file0600 , atomic replacement, and deletion. Add a test with a legacy session JSON payload in the old keyring account and oldsessions fallback path;Load must returnfound == false and never return the JSON as an API key. Add tests thatSave removes the legacy fallback after a successful new-key save andDelete removes both current and legacy local entries.key, found, err := store.Load("https://controller.example:443") if err != nil || found || key != "" { t.Fatalf("legacy record was usable: key=%q found=%t err=%v", key, found, err) }
- 
Step 2: Run storage tests to verify they fail before implementation Run: go test ./internal/authstoreExpected: FAIL because internal/authstore and the API-keyStore interface do not exist.
- 
Step 3: Implement the API-key store and narrow legacy cleanup Move the current keyring abstraction, URL normalization, account hash, state-home selection, atomic write, and permission handling into internal/authstore . Persist a JSON record with controller and API-key fields soLoad can validate that the controller matches and distinguish the record from the legacy session JSON. Use the existingunifi-cli keyring service and controller account hash: a successful new-key save replaces a legacy keyring session record in place. Use a new fallback subdirectory such askeys ; retain knowledge of the oldsessions subdirectory only to delete it.Implement these safety rules exactly: // Load reads only a valid API-key record from the new storage location. // A legacy or malformed record is not an API key and returns found == false. // Save uses the keyring first; only ErrKeyringUnavailable plus // allowFileFallback writes the protected fallback. // Delete removes the keyring account, new fallback, and legacy fallback; // missing entries are successful. Error text from storage must identify an operation but never include encoded records, fallback contents, or key values.
- 
Step 4: Run storage tests and the package-wide test suite Run: gofmt -w internal/authstore/store.go internal/authstore/store_test.go && go test ./internal/authstore && go test ./...Expected: PASS. The existing session package remains temporarily untouched until the client has migrated in Task 3, so the repository must still compile at this task boundary.
- 
Step 5: Commit the storage replacement git add internal/authstore/store.go internal/authstore/store_test.go
git commit -m "feat: persist controller API keys securely"
Files:
- Modify: internal/client/client.go
- Modify: internal/client/auth.go
- Modify: internal/client/client_test.go
- Modify: internal/cli/helpers.go
- Modify: internal/config/config.go
- Delete: internal/session/store.go
- Delete: internal/session/store_test.go
Interfaces:
- 
Consumes: authstore.Store ,apperr.NotAuthenticated , and non-secretconfig.Config .
- 
Produces: func New(cfg config.Config) (*Client, error) func NewWithStore(cfg config.Config, store authstore.Store) (*Client, error) func NewWithAPIKey(cfg config.Config, apiKey, method string) (*Client, error) func (c *Client) AuthMethod() string func (c *Client) Validate(ctx context.Context) error
- 
Contract: AuthMethod() returns exactly"environment_api_key" or"saved_api_key" for normal clients. The temporary validation client uses"interactive_api_key" and never persists or deletes state.
- 
Step 1: Replace session tests with failing API-key source tests Remove password-login, cookie-rotation, CSRF, and saved-session tests. Add injected in-memory authstore.Store tests that assert:
  - UNIFI_API_KEY is sent asX-API-KEY and wins over a saved key.
  - A saved key is loaded once and sent as X-API-KEY by a fresh client.
  - No environment key and no saved key produces apperr.NotAuthenticated with hintrun 'unifi login' before an HTTP request.
  - A 401 fromsaved_api_key deletes only that controller's store record and returns the login hint.
  - A 401 fromenvironment_api_key leaves the store untouched.
  - NewWithAPIKey validates an entered key without accessing the store.
 err := c.Do(ctx, http.MethodGet, client.PathSelfSites, nil, nil) if !apperr.Is(err, apperr.NotAuthenticated) || requests != 0 { t.Fatalf("missing-key behavior: err=%v requests=%d", err, requests) }
- 
Step 2: Run client tests to verify the new cases fail Run: go test ./internal/clientExpected: FAIL because the API-key-only constructors, private credential source, and invalidation behavior do not exist yet.
- 
Step 3: Implement key resolution, request headers, and saved-key invalidation Delete Username ,Password , andAPIKey fromconfig.Config , then delete cookie-jar, CSRF token, response-cookie, password-login, read-only-session, and session-write code frominternal/client . Store the active API key privately onClient ; never put it back onconfig.Config . InNewWithStore , use non-emptyUNIFI_API_KEY first, otherwise load fromauthstore.Store .NewWithAPIKey must construct the same transport with the supplied key andinteractive_api_key method without loading the store.ensureAuth must return:apperr.WithHint( apperr.New(apperr.NotAuthenticated, "not authenticated"), "run 'unifi login' to save an API key", ) Set X-API-KEY from the private client field. After anapperr.AuthFailed , callstore.Delete(c.baseURL) only whenAuthMethod() == "saved_api_key" ; preserve a deletion failure as a cause without rendering it. UpdateloadRuntime to always construct the standard API-key client and remove the--no-session-write branch.
- 
Step 4: Run the client tests and preserve the CLI boundary for Task 4 Run: gofmt -w internal/client/client.go internal/client/auth.go internal/client/client_test.go internal/cli/helpers.go internal/config/config.go && go test ./internal/clientExpected: PASS. Leave command registration and CLI command-test updates to Task 4.
- 
Step 5: Commit the API-key-only client behavior git add internal/client/client.go internal/client/auth.go internal/client/client_test.go internal/cli/helpers.go internal/config/config.go internal/session/store.go internal/session/store_test.go
git commit -m "refactor: make client API-key-only"
Files:
- Create: internal/cli/login.go
- Create: internal/cli/login_test.go
- Create: internal/cli/prompt.go
- Modify: internal/cli/auth.go
- Modify: internal/cli/auth_test.go
- Modify: internal/cli/root.go
- Modify: go.mod
- Modify: go.sum
Interfaces:
- 
Consumes: client.NewWithAPIKey ,client.NewWithStore ,authstore.NewStore , andgolang.org/x/term .
- 
Produces: func newLoginCmd() *cobra.Command func newLogoutCmd() *cobra.Command func newAuthStatusCmd() *cobra.Command func promptAPIKey(in *os.File, out io.Writer) (string, error)
- 
Contract: only unifi login accepts interactive key input;unifi logout performs no controller request;unifi auth status validates the resolved source withGET client.PathSelfSites .
- 
