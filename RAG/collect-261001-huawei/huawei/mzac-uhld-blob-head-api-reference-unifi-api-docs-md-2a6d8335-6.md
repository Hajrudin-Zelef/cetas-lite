---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-6
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [1284, 1603]
sha256: 83f9565c1bbedf5b2f4e0a97a3dd8fa0a84c6ef7273b6e84d6a61f5079cdd2d2
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

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
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts" \
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
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/wifi/broadcasts" \
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
Docs: https://developer.ui.com/network/v10.1.84/getvouchers
GET /proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers
List Vouchers
Retrieve a paginated list of Hotspot vouchers.
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
      "createdAt": "2024-01-01T00:00:00Z",
      "name": "hotel-guest",
      "code": 4861409510,
      "authorizedGuestLimit": 1,
      "authorizedGuestCount": 0,
      "activatedAt": "2024-01-01T00:00:00Z",
      "expiresAt": "2024-01-01T00:00:00Z",
      "expired": true,
      "timeLimitMinutes": 1440,
      "dataUsageLimitMBytes": 1024,
      "rxRateLimitKbps": 1000,
      "txRateLimitKbps": 1000
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/createvouchers
POST /proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers
Generate Vouchers
Create one or more Hotspot vouchers.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Request Body:
{
  "count": 1,
  "name": "string",
  "authorizedGuestLimit": 1,
  "timeLimitMinutes": 0,
  "dataUsageLimitMBytes": 0,
  "rxRateLimitKbps": 0,
  "txRateLimitKbps": 0
}
Response:
{
  "vouchers": [
    {
      "id": "00000000-0000-0000-0000-000000000000",
      "createdAt": "2024-01-01T00:00:00Z",
      "name": "hotel-guest",
      "code": 4861409510,
      "authorizedGuestLimit": 1,
      "authorizedGuestCount": 0,
      "activatedAt": "2024-01-01T00:00:00Z",
      "expiresAt": "2024-01-01T00:00:00Z",
      "expired": true,
      "timeLimitMinutes": 1440,
      "dataUsageLimitMBytes": 1024,
      "rxRateLimitKbps": 1000,
      "txRateLimitKbps": 1000
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "count": 1,
  "name": "string",
  "authorizedGuestLimit": 1,
  "timeLimitMinutes": 0,
  "dataUsageLimitMBytes": 0,
  "rxRateLimitKbps": 0,
  "txRateLimitKbps": 0
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "count": 1,
  "name": "string",
  "authorizedGuestLimit": 1,
  "timeLimitMinutes": 0,
  "dataUsageLimitMBytes": 0,
  "rxRateLimitKbps": 0,
  "txRateLimitKbps": 0
}'
Docs: https://developer.ui.com/network/v10.1.84/deletevouchers
DELETE /proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers
Delete Vouchers
Remove Hotspot vouchers based on the specified filter criteria.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| filter | ✓ | string |  | 
Response:
{
  "vouchersDeleted": 0
}
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getvoucher
GET /proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers/{voucherId}
Get Voucher Details
Retrieve details of a specific Hotspot voucher.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| voucherId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "id": "00000000-0000-0000-0000-000000000000",
  "createdAt": "2024-01-01T00:00:00Z",
  "name": "hotel-guest",
  "code": 4861409510,
  "authorizedGuestLimit": 1,
  "authorizedGuestCount": 0,
  "activatedAt": "2024-01-01T00:00:00Z",
  "expiresAt": "2024-01-01T00:00:00Z",
  "expired": true,
  "timeLimitMinutes": 1440,
  "dataUsageLimitMBytes": 1024,
  "rxRateLimitKbps": 1000,
  "txRateLimitKbps": 1000
}
curl (local — direct to controller):
