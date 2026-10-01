---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-4
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [648, 948]
sha256: a895d8ea32d6ea51b4d37d31941de3a04e53692f76b23197fbde4cba2d1d01bf
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  },
  "default": true
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updatenetwork
PUT /proxy/network/integration/v1/sites/{siteId}/networks/{networkId}
Update Network
Update an existing network on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| networkId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "management": "string",
  "name": "Default Network",
  "enabled": true,
  "vlanId": 0,
  "dhcpGuarding": {
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  }
}
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
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  },
  "default": true
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "management": "string",
  "name": "Default Network",
  "enabled": true,
  "vlanId": 0,
  "dhcpGuarding": {
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  }
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "management": "string",
  "name": "Default Network",
  "enabled": true,
  "vlanId": 0,
  "dhcpGuarding": {
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  }
}'
Docs: https://developer.ui.com/network/v10.1.84/deletenetwork
DELETE /proxy/network/integration/v1/sites/{siteId}/networks/{networkId}
Delete Network
Delete an existing network on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| networkId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| force |  | boolean |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getnetworksoverviewpage
GET /proxy/network/integration/v1/sites/{siteId}/networks
List Networks
Retrieve a paginated list of all Networks on a site.
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
      "management": "string",
      "id": "00000000-0000-0000-0000-000000000000",
      "name": "Default Network",
      "enabled": true,
      "vlanId": 0,
      "metadata": {
        "origin": "string"
      },
      "default": true
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/networks" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/networks" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/createnetwork
POST /proxy/network/integration/v1/sites/{siteId}/networks
Create Network
Create a new network on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Request Body:
{
  "management": "string",
  "name": "Default Network",
  "enabled": true,
  "vlanId": 0,
  "dhcpGuarding": {
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  }
}
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
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  },
  "default": true
}
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/networks" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "management": "string",
  "name": "Default Network",
  "enabled": true,
  "vlanId": 0,
  "dhcpGuarding": {
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  }
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/networks" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "management": "string",
  "name": "Default Network",
  "enabled": true,
  "vlanId": 0,
  "dhcpGuarding": {
    "trustedDhcpServerIpAddresses": [
      "string"
    ]
  }
}'
Docs: https://developer.ui.com/network/v10.1.84/getnetworkreferences
GET /proxy/network/integration/v1/sites/{siteId}/networks/{networkId}/references
Get Network References
Retrieve references to a specific network.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| networkId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "referenceResources": [
    {
      "resourceType": "CLIENT",
      "referenceCount": 0,
      "references": [
        {
          "referenceId": "00000000-0000-0000-0000-000000000000"
        }
      ]
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}/references" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/networks/{networkId}/references" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getwifibroadcastdetails
GET /proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}
Get Wifi Broadcast Details
Retrieve detailed information about a specific Wifi.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| wifiBroadcastId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "type": "string",
  "id": "00000000-0000-0000-0000-000000000000",
  "name": "string",
  "metadata": {
    "origin": "string"
  },
  "enabled": true,
  "network": {
    "type": "string"
  },
  "securityConfiguration": {
    "type": "string"
  },
  "broadcastingDeviceFilter": {
    "type": "string"
  },
  "mdnsProxyConfiguration": {
    "mode": "string"
  },
  "multicastFilteringPolicy": {
    "action": "string"
  },
  "multicastToUnicastConversionEnabled": true,
  "clientIsolationEnabled": true,
  "hideName": true,
  "uapsdEnabled": true,
  "basicDataRateKbpsByFrequencyGHz": {
    "5": 6000,
    "2.4": 2000
  },
  "clientFilteringPolicy": {
