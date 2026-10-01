---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-5
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [949, 1283]
sha256: 12a8f02a6e2c8a163d4c12e0d4b615af03fc75d9727ae8b4fadabbfb55405e32
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

    "action": "ALLOW",
    "macAddressFilter": [
      "string"
    ]
  },
  "blackoutScheduleConfiguration": {
    "days": [
      {
        "type": "string",
        "day": "SUN"
      }
    ]
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updatewifibroadcast
PUT /proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}
Update Wifi Broadcast
Update an existing Wifi Broadcast on the specified site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| wifiBroadcastId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "type": "string",
  "name": "string",
  "network": {
    "type": "string"
  },
  "enabled": true,
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
    "action": "ALLOW",
    "macAddressFilter": [
      "string"
    ]
  },
  "blackoutScheduleConfiguration": {
    "days": [
      {
        "type": "string",
        "day": "SUN"
      }
    ]
  }
}
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
    "action": "ALLOW",
    "macAddressFilter": [
      "string"
    ]
  },
  "blackoutScheduleConfiguration": {
    "days": [
      {
        "type": "string",
        "day": "SUN"
      }
    ]
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "name": "string",
  "network": {
    "type": "string"
  },
  "enabled": true,
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
    "action": "ALLOW",
    "macAddressFilter": [
      "string"
    ]
  },
  "blackoutScheduleConfiguration": {
    "days": [
      {
        "type": "string",
        "day": "SUN"
      }
    ]
  }
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "name": "string",
  "network": {
    "type": "string"
  },
  "enabled": true,
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
    "action": "ALLOW",
    "macAddressFilter": [
      "string"
    ]
  },
  "blackoutScheduleConfiguration": {
    "days": [
      {
        "type": "string",
        "day": "SUN"
      }
    ]
  }
}'
Docs: https://developer.ui.com/network/v10.1.84/deletewifibroadcast
DELETE /proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}
Delete Wifi Broadcast
Delete an existing Wifi Broadcast from the specified site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| wifiBroadcastId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| force |  | boolean |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getwifibroadcastpage
GET /proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts
List Wifi Broadcasts
$1a
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
      "enabled": true,
      "metadata": {
        "origin": "string"
      },
      "network": {
        "type": "string"
      },
      "securityConfiguration": {
        "type": "string"
      },
      "broadcastingDeviceFilter": {
        "type": "string"
      }
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/createwifibroadcast
POST /proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts
Create Wifi Broadcast
Create a new Wifi Broadcast on the specified site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Request Body:
{
  "type": "string",
  "name": "string",
  "network": {
    "type": "string"
  },
  "enabled": true,
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
    "action": "ALLOW",
    "macAddressFilter": [
      "string"
    ]
  },
