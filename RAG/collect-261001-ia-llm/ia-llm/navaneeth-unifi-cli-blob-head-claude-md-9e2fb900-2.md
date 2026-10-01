---
id: collect-261001-ia-llm/ia-llm/navaneeth-unifi-cli-blob-head-claude-md-9e2fb900-2
title: "List clients (table format)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["claude"]
source: docs/RAG/collect-261001-ia-llm/navaneeth-unifi-cli-blob-head-claude-md-9e2fb900.md
source_anchor: ""
source_lines: [194, 276]
sha256: d9385c981df6689f7c46008152a5ae91b70e402aeeb9d59b3839de7bf1e997b7
---

# List clients (table format)

- unifi devices restart <mac> - Restart a device
- Add formatter to internal/output/
- Update --format flag validation
- Add format to switch statement in command
Example formats to add:
- CSV
- YAML
- Colored table (using fatih/color)
- Wide table (more columns)
Direct Dependencies:
github.com/spf13/cobra v1.10.2
github.com/spf13/viper v1.21.0
github.com/olekukonko/tablewriter v1.1.2
Test Dependencies:
- Standard library: testing ,net/http/httptest
Transitive Dependencies (auto-managed):
- github.com/spf13/pflag
- github.com/spf13/afero
- github.com/spf13/cast
- github.com/fsnotify/fsnotify
- And more (see go.sum)
- Make changes to source files
- Run tests: make test
- Build: make build
- Test manually: ./bin/unifi clients list
- Install: make install (copies to $GOPATH/bin)
Repository: Local repository (no remote configured initially)
Branch: master
Ignore patterns: See .gitignore
- bin/ directory
- .unifi-cli.yaml config file
- Standard Go ignores (*.test, *.out)
- IDE files (.vscode, .idea)
- 
TLS Verification: Disabled by default for self-signed certs (common in Unifi setups) 
  - Can be enabled by setting --insecure=false
- Can be enabled by setting 
- 
API Key Storage: 
  - Never commit API keys to git
  - Config file is in .gitignore
  - Prefer environment variables for CI/CD
- 
Credential Exposure: 
  - API key visible in process list if passed via command line
  - Prefer env vars or config file
- HTTP timeout: 30 seconds
- No caching implemented
- Each command makes fresh API call
- Table rendering is fast even with 100+ clients
- Ensure UNIFI_HOST andUNIFI_API_KEY are set
- Or configure in ~/.unifi-cli.yaml
- Or pass via --host flag
- Expected with self-signed certs
- Already handled: --insecure defaults to true
- To enforce verification: --insecure=false
- Invalid API key
- Check key in Unifi controller: Settings > Control Plane > Integrations
- This was fixed in the implementation
- If you see it, a field type needs changing from int64 to float64
- All exported functions have comments
- Error messages include context
- Tests cover edge cases and error paths
- Code follows Go conventions
- No external logging framework (uses standard fmt)
- Always check API response types - The tx_bytes-r float issue
- Environment variables can interfere with tests - Need cleanup
- Naming conflicts are common - Use descriptive names (APIClient vs Client)
- Viper automatically reads environment variables - Set prefix with SetEnvPrefix
- Table output formatting is tricky - tablewriter API changed between versions
- Go Version: 1.21+ (uses go.mod)
- Tested on: Linux 6.17.12-300.fc43.x86_64
- Unifi Network Application: Works with versions supporting the official API (9.1.105+)
For Unifi API documentation:
- Official Unifi API Docs
- Controller-specific docs: Navigate to Settings > Control Plane > Integrations in your controller
- Created full CLI structure with Cobra and Viper
- Implemented clients list command
- Added table and JSON output formats
- Created comprehensive test suite (>90% coverage)
- Fixed float64 type issue for tx_bytes-r and rx_bytes-r
- Added Makefile for build automation
- Documentation: README.md, spec.md, CLAUDE.md
