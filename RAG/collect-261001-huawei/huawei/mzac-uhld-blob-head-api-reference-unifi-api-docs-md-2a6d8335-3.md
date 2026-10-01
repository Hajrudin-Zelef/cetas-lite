---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-3
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [403, 647]
sha256: 0eccd46ee0f5717a8b2a654ae127b9a9eeef11704e293ffe947b0c111b8333c4
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

        "channel": 36
      }
    ]
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/removedevice
DELETE /proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}
Remove (Unadopt) Device
Removes (unadopts) an adopted device from the site. If the device is online, it will be reset to factory defaults.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
| deviceId | ✓ | string |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getadopteddevicelateststatistics
GET /proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/statistics/latest
Get Latest Adopted Device Statistics
Retrieve the latest real-time statistics of a specific adopted device, such as uptime, data transmission rates, CPU and memory utilization.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
| deviceId | ✓ | string |  | 
Response:
{
  "uptimeSec": 0,
  "lastHeartbeatAt": "2024-01-01T00:00:00Z",
  "nextHeartbeatAt": "2024-01-01T00:00:00Z",
  "loadAverage1Min": 0.0,
  "loadAverage5Min": 0.0,
  "loadAverage15Min": 0.0,
  "cpuUtilizationPct": 0.0,
  "memoryUtilizationPct": 0.0,
  "uplink": {
    "txRateBps": 0,
    "rxRateBps": 0
  },
  "interfaces": {
    "radios": [
      {
        "frequencyGHz": 0.0,
        "txRetriesPct": 0.0
      }
    ]
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/statistics/latest" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/statistics/latest" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getpendingdevicepage
GET /proxy/network/integration/v1/pending-devices
List Devices Pending Adoption
Retrieve a paginated list of devices pending adoption, including basic device information.
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| offset |  | integer |  | 
| limit |  | integer |  | 
| filter |  | string |  | 
Response:
{
  "offset": 0,
  "limit": 25,
  "count": 10,
  "totalCount": 1000,
  "data": [
    {
      "macAddress": "94:2a:6f:26:c6:ca",
      "ipAddress": "192.168.1.55",
      "model": "UHDIW",
      "state": "ONLINE",
      "supported": true,
      "firmwareVersion": "6.6.55",
      "firmwareUpdatable": true,
      "features": [
        "switching"
      ],
      "adoptionTargetSiteIds": [
        "00000000-0000-0000-0000-000000000000"
      ]
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/pending-devices" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/pending-devices" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/executeconnectedclientaction
POST /proxy/network/integration/v1/sites/{siteId}/clients/{clientId}/actions
Execute Client Action
Perform an action on a specific connected client. The request body must include the action name and any applicable input arguments.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| clientId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "action": "string"
}
Response:
{
  "action": "string"
}
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/clients/{clientId}/actions" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "action": "string"
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/clients/{clientId}/actions" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "action": "string"
}'
Docs: https://developer.ui.com/network/v10.1.84/getconnectedclientoverviewpage
GET /proxy/network/integration/v1/sites/{siteId}/clients
List Connected Clients
Retrieve a paginated list of all connected clients on a site, including physical devices (computers, smartphones) and active VPN connections.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| offset |  | integer |  | 
| limit |  | integer |  | 
| filter |  | string |  | 
Response:
{
  "offset": 0,
  "limit": 25,
  "count": 10,
  "totalCount": 1000,
  "data": [
    {
      "type": "string",
      "id": "00000000-0000-0000-0000-000000000000",
      "name": "string",
      "connectedAt": "2024-01-01T00:00:00Z",
      "ipAddress": "string",
      "access": {
        "type": "DEFAULT"
      }
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/clients" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/clients" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getconnectedclientdetails
GET /proxy/network/integration/v1/sites/{siteId}/clients/{clientId}
Get Connected Client Details
Retrieve detailed information about a specific connected client, including name, IP address, MAC address, connection type and access information.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| clientId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "type": "string",
  "id": "00000000-0000-0000-0000-000000000000",
  "name": "string",
  "connectedAt": "2024-01-01T00:00:00Z",
  "ipAddress": "string"
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/clients/{clientId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/clients/{clientId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getnetworkdetails
GET /proxy/network/integration/v1/sites/{siteId}/networks/{networkId}
Get Network Details
Retrieve detailed information about a specific network.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| networkId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "management": "string",
  "id": "00000000-0000-0000-0000-000000000000",
  "name": "Default Network",
  "enabled": true,
  "vlanId": 0,
  "metadata": {
    "origin": "string"
  },
  "dhcpGuarding": {
