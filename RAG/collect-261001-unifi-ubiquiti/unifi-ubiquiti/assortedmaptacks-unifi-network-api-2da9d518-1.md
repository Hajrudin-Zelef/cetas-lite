---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/assortedmaptacks-unifi-network-api-2da9d518-1
title: "Clone the repository"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["mcp"]
source: docs/RAG/collect-261001-unifi-ubiquiti/assortedmaptacks-unifi-network-api-2da9d518.md
source_anchor: ""
source_lines: [1, 124]
sha256: 70a0dbba02661f2a3cfbee00756979b98899c54dc37784af50dcc7d9695cf1f0
---

# Clone the repository

This project originally was taken on in order for me to build an MCP server to connect to my UniFi network devices. Unfortunately, there was not any good documentation to support this, so I started on a quest of identifying all of the API endpoints that a controller or device could have.
UniFi uses multiple different types of APIs and authentication. In my review, I researched the backend APIs, the frontend JavaScripts, and also their new official network API and compiled them into a single place here.
I hope this to be a comprehensive, community-driven documentation project for UniFi Network Controller APIs (v9.4+). This repository contains detailed OpenAPI specifications reverse-engineered from the UniFi Network Controller interface.
This project aims to document all available APIs in the UniFi Network ecosystem through collaborative effort. The API specifications have been compiled using multiple reverse engineering methods including:
- Analysis of UniFi Controller web interface network traffic
- Inspection of JavaScript bundles and UI components
- Official documentation cross-referencing
- Community testing and validation
# Clone the repository
git clone https://github.com/tmcpro/unifi-network-api.git
cd unifi-network-api
# Start a local HTTP server
python -m http.server 8000
# Open your browser to view the documentation
# http://localhost:8000
Simply open the index.html file in your web browser to view the OpenAPI documentation. This is loading a remote JS to render the OpenAPI Spec.
# Install Redoc CLI
npm install -g redoc-cli
# Generate static documentation
redoc-cli serve openapi/openapi.yaml --watch
# Or build static HTML
redoc-cli build openapi/openapi.yaml --output docs/index.html
The UniFi Network Controller supports multiple authentication methods depending on the API version and use case:
- Path Pattern: /integration/v1/*
- Authentication: X-API-Key header
- Setup: Generate API keys in the UniFi Controller under Settings → Integrations
curl -H "X-API-Key: YOUR_API_KEY" \
  "https://{{your-controller}}/proxy/network/integration/v1/sites"
- Path Pattern: /v2/api/* and/api/*
- Authentication: Session cookies + CSRF tokens
- Use Case: Browser-based applications, advanced controller features
- Setup: Login via /api/login endpoint
# Login and save cookies
curl -c cookies.jar -X POST \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"your_password"}' \
  "https://{{your-controller}}/proxy/network/api/login"
# Use cookies for subsequent requests
curl -b cookies.jar \
  "https://{{your-controller}}/proxy/network/v2/api/info"
- Integration v1: Limited subset focused on common operations (15 endpoints)
- Controller v2: Complete feature set with full controller functionality (227+ endpoints)
- v1 Responses: Include pagination metadata (offset ,limit ,count ,totalCount )
- v2 Responses: Use data +meta envelope pattern
All APIs return structured error responses:
{
  "code": "ERROR_CODE",
  "message": "Human readable error description",
  "details": {}
}
API endpoint availability depends on your UniFi hardware and firmware:
- UniFi Dream Machine (UDM/UDM Pro): Likely Full API support
- CloudKey Gen2+: Likely API support
- Self-hosted Controllers: Likely Partial API support
- Legacy Hardware: Limited endpoint availability
- UniFi OS vs Traditional Controller: Some endpoints may differ
Note: Always test endpoints with your specific hardware configuration before production deployment.
We encourage community contributions to improve and validate this documentation:
- Fork this repository
- Test API endpoints with your UniFi setup
- Document your findings (success/failure, hardware compatibility)
- Submit pull requests with improvements
- Report issues for incorrect or missing documentation
- Validate endpoints before submitting changes
- Include hardware/firmware version information
- Provide example requests and responses
- Follow the existing OpenAPI specification format
- Add appropriate tags and descriptions
- Validate all documented endpoints across different hardware
- Expand coverage of undocumented APIs
- Improve example requests and responses
- Document hardware-specific limitations
- Create SDKs and client libraries
The following table lists all 327 documented API endpoints organized by category:
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| About Application | /v1/info | GET | Get Application Info | 
| Sites | /v1/sites | GET | List Local Sites | 
| UniFi Devices | /v1/sites/{siteId}/devices | GET | List Site Devices | 
|  | /v1/sites/{siteId}/devices/{deviceId} | GET | Get Device Details | 
|  | /v1/sites/{siteId}/devices/{deviceId}/statistics | GET | Get Device Statistics | 
|  | /v1/sites/{siteId}/devices/{deviceId}/actions | POST | Execute Device Action | 
|  | /v1/sites/{siteId}/devices/{deviceId}/interfaces/ports/{portIdx}/actions | POST | Execute Port Action | 
| Clients | /v1/sites/{siteId}/clients | GET | List Site Clients | 
|  | /v1/sites/{siteId}/clients/{clientId} | GET | Get Client Details | 
|  | /v1/sites/{siteId}/clients/{clientId}/actions | POST | Execute Client Action | 
| Hotspot Vouchers | /v1/sites/{siteId}/hotspot/vouchers | GET | List Vouchers | 
|  | /v1/sites/{siteId}/hotspot/vouchers | POST | Create Voucher | 
|  | /v1/sites/{siteId}/hotspot/vouchers/{voucherId} | GET | Get Voucher Details | 
|  | /v1/sites/{siteId}/hotspot/vouchers/{voucherId} | DELETE | Delete Voucher | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Settings | /api/login | POST | Authenticate to Controller | 
|  | /api/logout | POST | End Session | 
|  | /v2/api/info | GET | Get Controller Info | 
|  | /v2/api/log-levels/defaults | GET | Get Default Log Levels | 
|  | /v2/api/timezones | GET | List Supported Timezones | 
|  | /v2/api/site/{site}/settings/mgmt | GET | Get Management Settings | 
|  | /v2/api/site/{site}/settings/mgmt | PUT | Update Management Settings | 
|  | /v2/api/site/{site}/settings/snmp | GET | Get SNMP Settings | 
|  | /v2/api/site/{site}/settings/snmp | PUT | Update SNMP Settings | 
|  | /v2/api/site/{site}/settings/radio-ai | GET | Get Radio AI Settings | 
|  | /v2/api/site/{site}/settings/radio-ai | PUT | Update Radio AI Settings | 
|  | /v2/api/site/{site}/settings/syslog | GET | Get Syslog Settings | 
|  | /v2/api/site/{site}/settings/syslog | PUT | Update Syslog Settings | 
|  | /v2/api/site/{site}/settings/super-mgmt | GET | Get Super Management Settings | 
|  | /v2/api/site/{site}/settings/super-mgmt | PUT | Update Super Management Settings | 
|  | /v2/api/site/{site}/settings/radius | GET | Get RADIUS Settings | 
|  | /v2/api/site/{site}/settings/radius | PUT | Update RADIUS Settings | 
|  | /v2/api/site/{site}/settings/mdns | GET | Get mDNS Settings | 
|  | /v2/api/site/{site}/settings/mdns | PUT | Update mDNS Settings | 
|  | /v2/api/site/{site}/settings/country | GET | Get Country Settings | 
|  | /v2/api/site/{site}/settings/country | PUT | Update Country Settings | 
|  | /v2/api/site/{site}/settings/locale | GET | Get Locale Settings | 
|  | /v2/api/site/{site}/settings/locale | PUT | Update Locale Settings | 
|  | /v2/api/site/{site}/settings/network-optimization | GET | Get Network Optimization Settings | 
|  | /v2/api/site/{site}/settings/network-optimization | PUT | Update Network Optimization Settings | 
|  | /v2/api/site/{site}/settings/baresip | GET | Get Baresip Settings | 
|  | /v2/api/site/{site}/settings/baresip | PUT | Update Baresip Settings | 
|  | /v2/api/site/{site}/settings/super-identity | GET | Get Super Identity Settings | 
|  | /v2/api/site/{site}/settings/super-identity | PUT | Update Super Identity Settings | 
|  | /v2/api/site/{site}/settings/super-mail | GET | Get Super Mail Settings | 
|  | /v2/api/site/{site}/settings/super-mail | PUT | Update Super Mail Settings | 
|  | /v2/api/site/{site}/settings/super-smtp | GET | Get Super SMTP Settings | 
