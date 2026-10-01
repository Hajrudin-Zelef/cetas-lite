---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/lexfrei-go-unifi-blob-head-dig-findings-md-8aca4575
title: "lexfrei-go-unifi-blob-head-dig-findings-md-8aca4575"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-unifi-ubiquiti/lexfrei-go-unifi-blob-head-dig-findings-md-8aca4575.md
source_anchor: ""
source_lines: [1, 109]
sha256: 1bf21a024be3ff0252331ddb56d5d543bfecd6e98175b28141dc5c2328c4e8fd
---

# lexfrei-go-unifi-blob-head-dig-findings-md-8aca4575

- Model: UDR7 (UniFi Dream Router 7)
- UniFi Network Version: 9.5.21
- UniFi OS Version: 4.x
- IP: 172.16.0.1
Location: /usr/lib/unifi/webapps/ROOT/api-docs/integration.json
Base URL: /integration
Version: 9.5.21
This is the official public API documented by Ubiquiti at developer.ui.com.
- GET /v1/sites - List local sites
- GET /v1/sites/{siteId}/devices - List devices
- GET /v1/sites/{siteId}/devices/{deviceId} - Get device details
- GET /v1/sites/{siteId}/devices/{deviceId}/statistics/latest - Device stats
- GET /v1/sites/{siteId}/clients - List connected clients
- GET /v1/sites/{siteId}/clients/{clientId} - Get client details
- POST /v1/sites/{siteId}/clients/{clientId}/actions - Client actions
- POST /v1/sites/{siteId}/devices/{deviceId}/actions - Device actions
- POST /v1/sites/{siteId}/devices/{deviceId}/interfaces/ports/{portIdx}/actions - Port actions
- GET /v1/sites/{siteId}/hotspot/vouchers - List vouchers
- POST /v1/sites/{siteId}/hotspot/vouchers - Generate vouchers
- DELETE /v1/sites/{siteId}/hotspot/vouchers - Delete vouchers
- GET /v1/info - Get application info
- Uses X-API-KEY header
| Port | Service | Description | 
|---|---|---|
| 8080 | unifi | Device inform port | 
| 8081 | unifi | Internal API (login required) | 
| 8443 | unifi | Web UI HTTPS | 
| 8843 | unifi | Guest portal HTTPS | 
| 8880 | unifi | Guest portal HTTP | 
| 6789 | unifi | Mobile speedtest | 
| 9080 | ulp-go-app | User management API | 
| 9443 | ulp-go-app | User management HTTPS | 
| 11010 | uos-agent | Internal UOS Agent | 
| 11011 | uos-agent | UOS Agent API | 
| 11051 | node22 | gRPC | 
| 11081 | node22 | IPC | 
| 1080 | udapi-bridge | UDAPI bridge | 
| 27117 | mongod | MongoDB (WiredTiger) | 
| 5432 | postgres | PostgreSQL | 
curl http://127.0.0.1:8081/status
Returns:
{
  "meta": {
    "rc": "ok",
    "uuid": "...",
    "server_version": "9.5.21",
    "server_running": true,
    "db_running": true,
    "db_connected": true,
    "ucore_installation": true,
    "udm_connected": true
  }
}
Ports: 11010 (internal), 11011 (API)
Binary: /usr/bin/uos-agent (Rust-based)
Found API paths:
- /api/v1/backups/export
- /api/v1/backups/import
- /api/v2/alarms/{source} (protect, network, access, connect)
- /ucore/api-docs (OpenAPI documentation)
Port: 9080 (HTTP), 9443 (HTTPS)
Binary: ulp-go-app
Configuration at /usr/lib/ulp-go/:
- scopes/ - Permission scopes for API keys
- capabilities/ - Controller capabilities
- config.props - Main configuration
Port: 1080 (udapi-bridge)
Config: /usr/share/ubios-udapi-server/
Device-specific configs:
- udr7-a67a.default - Default configuration
- config-board/udr7-a67a.json - Board capabilities
Located at /usr/share/unifi-core/http/ and /data/unifi-core/config/http/
Key routes:
- /proxy/network/ →http://network_api_backend/ (port 8081)
- /api/ →http://uos_api_backend/api/
- /api/v2/ →http://uos_agent_api_backend/api/v2/
- network-integration-api-9.5.21.json - Full OpenAPI spec for Network API
- unifi-core-config.yaml - UniFi Core service configuration
- unifi-os-console-models.json - All console model definitions
- uidb-devices.json - Device database (all UniFi device models)
- scopes/ - Permission scopes for API keys
- capabilities/ - Controller capabilities per version
- nginx-http/ - Nginx HTTP configurations
- udr7-board-config.json - UDR7 board capabilities
- udr7-default-config.json - UDR7 default configuration
- listening-services.txt - All listening TCP services
- running-processes.txt - Running UniFi processes
From scopes/api-key/access.yaml:
Access API scopes:
- view:open_api.user /edit:open_api.user
- view:open_api.visitor /edit:open_api.visitor
- view:open_api.policy /edit:open_api.policy
- view:open_api.credential /edit:open_api.credential
- view:open_api.space /edit:open_api.space
- view:open_api.device /edit:open_api.device
- view:open_api.system_log /edit:open_api.system_log
- view:open_api.webhook /edit:open_api.webhook
- view:open_api.api_server /edit:open_api.api_server
- UniFi Core (Node.js) orchestrates all services
- ulp-go (Go) handles user management and authentication
- uos-agent (Rust) handles alarms and backups
- unifi (Java/GraalVM native) is the main Network controller
- MongoDB stores Network configuration
- PostgreSQL stores user/session data
- Update api/network/openapi.yaml with new Integration API schema
- Add hotspot voucher management endpoints
- Add device actions (restart, etc.)
- Add client actions (authorize/unauthorize guest)
- Consider adding internal status endpoint
