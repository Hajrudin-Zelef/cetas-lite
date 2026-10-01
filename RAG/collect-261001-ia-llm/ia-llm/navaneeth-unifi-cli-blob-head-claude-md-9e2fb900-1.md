---
id: collect-261001-ia-llm/ia-llm/navaneeth-unifi-cli-blob-head-claude-md-9e2fb900-1
title: "List clients (table format)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: ["2026-01-11"]
keywords: ["claude"]
source: docs/RAG/collect-261001-ia-llm/navaneeth-unifi-cli-blob-head-claude-md-9e2fb900.md
source_anchor: ""
source_lines: [1, 193]
sha256: 118348cc8a786fc0602bd993ed3cb142c1f48769fe57a8c0da458f29b29ff1c4
---

# List clients (table format)

This document contains comprehensive implementation details for future reference and development.
Project Name: Unifi CLI Language: Go (Golang) Purpose: Command-line interface for managing Unifi Network devices using the official Unifi Network API Created: 2026-01-11
- CLI Framework: Cobra v1.10.2 - Command structure and parsing
- Configuration: Viper v1.21.0 - Config file, env vars, and flags management
- Table Output: tablewriter v1.1.2 - ASCII table formatting
unifi-cli/
├── cmd/
│   ├── root.go           # Root command with global flags and config initialization
│   └── clients.go        # Clients subcommand (list connected clients)
├── internal/
│   ├── api/
│   │   ├── client.go     # HTTP API client implementation
│   │   ├── client_test.go # API client tests (94.8% coverage)
│   │   ├── types.go      # API response types and helper methods
│   │   └── types_test.go # Type tests
│   ├── config/
│   │   ├── config.go     # Viper-based configuration management
│   │   └── config_test.go # Config tests (93.1% coverage)
│   └── output/
│       ├── table.go      # Table output formatter
│       ├── table_test.go # Table output tests
│       ├── json.go       # JSON output formatter
│       └── json_test.go  # JSON output tests (91.7% coverage)
├── main.go               # Entry point - calls cmd.Execute()
├── Makefile              # Build automation
├── go.mod                # Go module definition
├── go.sum                # Dependency checksums
├── spec.md               # Original specification
├── README.md             # User documentation
├── .gitignore            # Git ignore rules
└── CLAUDE.md             # This file
File: internal/config/config.go
Configuration Priority (highest to lowest):
- Command-line flags
- Environment variables (with UNIFI_ prefix)
- Config file (~/.unifi-cli.yaml )
- Default values
Environment Variables:
- UNIFI_HOST - Controller host URL
- UNIFI_API_KEY - API authentication key
- UNIFI_SITE - Site ID (default: "default")
Default Values:
- site : "default"
- insecure : true (TLS verification skipped by default for self-signed certs)
Important Functions:
- Init(cfgFile string) - Initialize configuration from file or default location
- Get() - Returns singleton config instance
- Validate() - Validates required fields (host, api_key)
- GetConfigPath() - Returns current or default config file path
File: internal/api/client.go
Key Design Decisions:
- Uses APIClient struct (notClient to avoid naming conflict withClient type in types.go)
- TLS verification skipped by default (InsecureSkipVerify: true )
- 30-second HTTP timeout
- Trailing slash automatically removed from host URL
API Endpoints:
- List Clients: {host}/proxy/network/api/s/{site}/stat/sta
- List Sites: {host}/proxy/network/api/self/sites
Authentication:
- Header: X-API-KEY: {api_key}
- Content-Type: application/json
Important Note: The API returns floating-point numbers for tx_bytes-r and rx_bytes-r fields, so these are typed as float64 in the Client struct.
File: internal/api/types.go
Main Types:
type Client struct {
    MAC       string
    Name      string
    Hostname  string
    IP        string
    IsWired   bool
    Essid     string  // SSID for wireless
    Signal    int     // dBm for wireless
    Uptime    int64   // seconds
    RxBytes   int64
    TxBytes   int64
    TxBytesR  float64 // IMPORTANT: float64 not int64!
    RxBytesR  float64 // IMPORTANT: float64 not int64!
    // ... many more fields
}
Helper Methods:
- GetDisplayName() - Returns Name → Hostname → MAC (fallback chain)
- GetConnectionType() - Returns "Wired" or "Wireless"
- GetSSID() - Returns SSID for wireless clients, empty for wired
- GetSignal() - Returns formatted signal strength (e.g., "-65 dBm")
- GetUptime() - Returns human-readable uptime (e.g., "5d 3h 15m")
Utility Functions:
- FormatBytes(bytes int64) - Converts bytes to human-readable format (KB, MB, GB, TB)
Table Format (internal/output/table.go):
- Uses tablewriter library
- Columns: MAC, Name, IP, Type, SSID, Signal, Uptime, RX/TX
- SSID and Signal only shown for wireless clients
- RX/TX shows formatted bytes (e.g., "35.8 MB / 25.5 MB")
JSON Format (internal/output/json.go):
- Pretty-printed with 2-space indentation
- Returns raw API response data
Root Command (cmd/root.go):
unifi [command] [flags]
Global Flags:
- --config, -c - Config file path
- --host - Unifi controller host
- --site - Site ID (default: "default")
- --insecure, -k - Skip TLS verification (default: true)
Clients Command (cmd/clients.go):
unifi clients list [flags]
Flags:
- --format, -f - Output format: "table" (default) or "json"
Makefile Targets:
make build      # Build to ./bin/unifi
make compile    # Alias for build
make install    # Run 'go install'
make clean      # Remove ./bin directory
make test       # Run tests
make lint       # Run golangci-lint (requires installation)
make help       # Show available targets
Test Coverage:
- internal/api: 94.8%
- internal/config: 93.1%
- internal/output: 91.7%
Test Files:
- client_test.go - Uses httptest for mocking API responses
- config_test.go - Tests config loading, validation, environment variable handling
- types_test.go - Tests all helper methods and formatters
- json_test.go - Captures stdout and validates JSON output
- table_test.go - Validates table formatting and content
Running Tests:
go test ./...              # Run all tests
go test ./... -v           # Verbose output
go test ./... -cover       # With coverage
Problem: API returns floating-point numbers but code expected int64
Error: json: cannot unmarshal number 43.08318972270092 into Go struct field Client.data.tx_bytes-r of type int64
Solution: Changed type from int64 to float64 in internal/api/types.go
Problem: Config tests failed because UNIFI_API_KEY from environment overrode test config
Solution: Added env var cleanup in TestInitWithValidConfigFile
Problem: Both API client and response data use "Client" name
Solution: Renamed API client struct to APIClient, kept data type as Client
Official Unifi API Documentation:
- Access via: Unifi Network > Settings > Control Plane > Integrations
- Generate API keys in the same location
- Documentation is version-specific to your controller
API Response Structure:
{
  "meta": {
    "rc": "ok"  // "ok" for success, "error" for failure
  },
  "data": [
    {
      "_id": "...",
      "mac": "aa:bb:cc:dd:ee:ff",
      "ip": "192.168.1.100",
      "is_wired": true,
      // ... many more fields
    }
  ]
}# Set environment variables
export UNIFI_HOST="https://unifi.example.com"
export UNIFI_API_KEY="your-api-key"
# List clients (table format)
./bin/unifi clients list
# List clients (JSON format)
./bin/unifi clients list --format json
./bin/unifi clients list -f json# Create config file
cat > ~/.unifi-cli.yaml <<EOF
host: https://unifi.example.com
api_key: your-api-key
site: default
insecure: true
EOF
# Use it
./bin/unifi clients list./bin/unifi --host https://unifi.example.com clients list
./bin/unifi --site custom-site clients list
- Create command file in cmd/ (e.g.,devices.go )
- Define command structure:
var devicesCmd = &cobra.Command{
    Use:   "devices",
    Short: "Manage Unifi devices",
}
var devicesListCmd = &cobra.Command{
    Use:   "list",
    Short: "List devices",
    RunE:  runDevicesList,
}
func init() {
    rootCmd.AddCommand(devicesCmd)
    devicesCmd.AddCommand(devicesListCmd)
}
- Add API method to internal/api/client.go
- Add response types to internal/api/types.go
- Add output formatters to internal/output/
- unifi sites list - List all sites
- unifi devices list - List network devices (APs, switches, etc.)
- unifi networks list - List networks/SSIDs
- unifi clients block <mac> - Block a client
- unifi clients unblock <mac> - Unblock a client
