---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-1
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [1, 127]
sha256: d704aac8852cb498e93479e94f11df789013f12d31042b87bac370468d88a71a
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

Source: https://developer.ui.com/network/v10.1.84
Auth: Add X-API-Key: <your-api-key> header to every request.
Generate the key in UniFi OS → Settings → Control Plane → Integrations.
Base URL (local): https://<controller-ip>/proxy/network/integration/v1/...
Base URL (cloud): https://api.ui.com/proxy/network/integration/v1/...
Docs: https://developer.ui.com/network/v10.1.84/gettingstarted
ℹ️ No endpoint data found (may be a docs/info page)
Docs: https://developer.ui.com/network/v10.1.84/filtering
ℹ️ No endpoint data found (may be a docs/info page)
Docs: https://developer.ui.com/network/v10.1.84/error-handling
ℹ️ No endpoint data found (may be a docs/info page)
Docs: https://developer.ui.com/network/v10.1.84/connectorpost
POST /proxy/network/integration/v1/connector/consoles/{id}/*path
Connector - POST
Forward POST requests to UniFi applications using RESTful HTTP methods. Request Flow: The request is proxied through api.ui.com cloud endpoint to the remote console at http://127.0.0.1/proxy/[path] with POST method. Requirements: - Console firmware version must >= 5.0.3- For non-organization API keys: Limited to API key owner's consoles only (cannot access other admins' consoles)- For organization API keys: Can access any console within the organizationAPI Documentation: - Network API: https://developer.ui.com/network- Protect API: https://developer.ui.com/protectResponse: On success, the upstream API response is passed through directly. On error, a standardized error schema is returned.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| id | ✓ | string | Host ID to proxy the request to | 
| path | ✓ | string | API path to proxy | 
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/connectorget
GET /proxy/network/integration/v1/connector/consoles/{id}/*path
Connector - GET
$1a
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| id | ✓ | string | Host ID to proxy the request to | 
| path | ✓ | string | API path to proxy | 
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/connectorput
PUT /proxy/network/integration/v1/connector/consoles/{id}/*path
Connector - PUT
Forward PUT requests to UniFi applications using RESTful HTTP methods. Request Flow: The request is proxied through api.ui.com cloud endpoint to the remote console at http://127.0.0.1/proxy/[path] with PUT method. Requirements: - Console firmware version must >= 5.0.3- For non-organization API keys: Limited to API key owner's consoles only (cannot access other admins' consoles)- For organization API keys: Can access any console within the organizationAPI Documentation: - Network API: https://developer.ui.com/network- Protect API: https://developer.ui.com/protectResponse: On success, the upstream API response is passed through directly. On error, a standardized error schema is returned.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| id | ✓ | string | Host ID to proxy the request to | 
| path | ✓ | string | API path to proxy | 
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/connectordelete
DELETE /proxy/network/integration/v1/connector/consoles/{id}/*path
Connector - DELETE
Forward DELETE requests to UniFi applications using RESTful HTTP methods. Request Flow: The request is proxied through api.ui.com cloud endpoint to the remote console at http://127.0.0.1/proxy/[path] with DELETE method. Requirements: - Console firmware version must >= 5.0.3- For non-organization API keys: Limited to API key owner's consoles only (cannot access other admins' consoles)- For organization API keys: Can access any console within the organizationAPI Documentation: - Network API: https://developer.ui.com/network- Protect API: https://developer.ui.com/protectResponse: On success, the upstream API response is passed through directly. On error, a standardized error schema is returned.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| id | ✓ | string | Host ID to proxy the request to | 
| path | ✓ | string | API path to proxy | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/connectorpatch
PATCH /proxy/network/integration/v1/connector/consoles/{id}/*path
Connector - PATCH
Forward PATCH requests to UniFi applications using RESTful HTTP methods. Request Flow: The request is proxied through api.ui.com cloud endpoint to the remote console at http://127.0.0.1/proxy/[path] with PATCH method. Requirements: - Console firmware version must >= 5.0.3- For non-organization API keys: Limited to API key owner's consoles only (cannot access other admins' consoles)- For organization API keys: Can access any console within the organizationAPI Documentation: - Network API: https://developer.ui.com/network- Protect API: https://developer.ui.com/protectResponse: On success, the upstream API response is passed through directly. On error, a standardized error schema is returned.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| id | ✓ | string | Host ID to proxy the request to | 
| path | ✓ | string | API path to proxy | 
curl (local — direct to controller):
curl -sS -L \
  -X PATCH "https://192.168.1.1/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PATCH "https://api.ui.com/proxy/network/integration/v1/connector/consoles/{id}/*path" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getinfo
GET /proxy/network/integration/v1/info
Get Application Info
Retrieve general information about the UniFi Network application.
Response:
{
  "applicationVersion": "9.1.0"
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/info" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/info" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getsiteoverviewpage
GET /proxy/network/integration/v1/sites
List Local Sites
