---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-2
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [128, 402]
sha256: 8fa19d5b7193c4b16e5e73c348098a660bf2b897d6b507c9535579e67fbe06f3
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

Retrieve a paginated list of local sites managed by this Network application. Site ID is required for other UniFi Network API calls.
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
      "id": "00000000-0000-0000-0000-000000000000",
      "internalReference": "string",
      "name": "string"
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getadopteddeviceoverviewpage
GET /proxy/network/integration/v1/sites/{siteId}/devices
List Adopted Devices
Retrieve a paginated list of all adopted devices on a site, including basic device information.
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
      "id": "00000000-0000-0000-0000-000000000000",
      "macAddress": "94:2a:6f:26:c6:ca",
      "ipAddress": "192.168.1.55",
      "name": "IW HD",
      "model": "UHDIW",
      "state": "ONLINE",
      "supported": true,
      "firmwareVersion": "6.6.55",
      "firmwareUpdatable": true,
      "features": [
        "switching"
      ],
      "interfaces": [
        "ports"
      ]
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/devices" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/devices" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/adoptdevice
POST /proxy/network/integration/v1/sites/{siteId}/devices
Adopt Devices
Adopt a device to a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Request Body:
{
  "macAddress": "string",
  "ignoreDeviceLimit": true
}
Response:
{
  "id": "00000000-0000-0000-0000-000000000000",
  "macAddress": "94:2a:6f:26:c6:ca",
  "ipAddress": "192.168.1.55",
  "name": "IW HD",
  "model": "UHDIW",
  "supported": true,
  "state": "ONLINE",
  "firmwareVersion": "6.6.55",
  "firmwareUpdatable": true,
  "adoptedAt": "2024-01-01T00:00:00Z",
  "provisionedAt": "2024-01-01T00:00:00Z",
  "configurationId": "7596498d2f367dc2",
  "uplink": {
    "deviceId": "00000000-0000-0000-0000-000000000000"
  },
  "features": {
    "switching": {},
    "accessPoint": {}
  },
  "interfaces": {
    "ports": [
      {
        "idx": 1,
        "state": "UP",
        "connector": "RJ45",
        "maxSpeedMbps": 10000,
        "speedMbps": 1000,
        "poe": {
          "standard": "802.3bt",
          "type": 3,
          "enabled": true,
          "state": "UP"
        }
      }
    ],
    "radios": [
      {
        "wlanStandard": "802.11a",
        "frequencyGHz": 0.0,
        "channelWidthMHz": 40,
        "channel": 36
      }
    ]
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/devices" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "macAddress": "string",
  "ignoreDeviceLimit": true
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/devices" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "macAddress": "string",
  "ignoreDeviceLimit": true
}'
Docs: https://developer.ui.com/network/v10.1.84/executeportaction
POST /proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/interfaces/ports/{portIdx}/actions
Execute Port Action
Perform an action on a specific device port. The request body must include the action name and any applicable input arguments.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| portIdx | ✓ | integer |  | 
| siteId | ✓ | string |  | 
| deviceId | ✓ | string |  | 
Request Body:
{
  "action": "string"
}
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/interfaces/ports/{portIdx}/actions" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "action": "string"
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/interfaces/ports/{portIdx}/actions" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "action": "string"
}'
Docs: https://developer.ui.com/network/v10.1.84/executeadopteddeviceaction
POST /proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/actions
Execute Adopted Device Action
Perform an action on an specific adopted device. The request body must include the action name and any applicable input arguments.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
| deviceId | ✓ | string |  | 
Request Body:
{
  "action": "string"
}
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/actions" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "action": "string"
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}/actions" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "action": "string"
}'
Docs: https://developer.ui.com/network/v10.1.84/getadopteddevicedetails
GET /proxy/network/integration/v1/sites/{siteId}/devices/{deviceId}
Get Adopted Device Details
Retrieve detailed information about a specific adopted device, including firmware versioning, uplink state, details about device features and interfaces (ports, radios) and other key attributes.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
| deviceId | ✓ | string |  | 
Response:
{
  "id": "00000000-0000-0000-0000-000000000000",
  "macAddress": "94:2a:6f:26:c6:ca",
  "ipAddress": "192.168.1.55",
  "name": "IW HD",
  "model": "UHDIW",
  "supported": true,
  "state": "ONLINE",
  "firmwareVersion": "6.6.55",
  "firmwareUpdatable": true,
  "adoptedAt": "2024-01-01T00:00:00Z",
  "provisionedAt": "2024-01-01T00:00:00Z",
  "configurationId": "7596498d2f367dc2",
  "uplink": {
    "deviceId": "00000000-0000-0000-0000-000000000000"
  },
  "features": {
    "switching": {},
    "accessPoint": {}
  },
  "interfaces": {
    "ports": [
      {
        "idx": 1,
        "state": "UP",
        "connector": "RJ45",
        "maxSpeedMbps": 10000,
        "speedMbps": 1000,
        "poe": {
          "standard": "802.3bt",
          "type": 3,
          "enabled": true,
          "state": "UP"
        }
      }
    ],
    "radios": [
      {
        "wlanStandard": "802.11a",
        "frequencyGHz": 0.0,
        "channelWidthMHz": 40,
