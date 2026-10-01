---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-7
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [1604, 1874]
sha256: a8c4653ba3bc477f4a1aff9157c68143188c8cfa1545ea85d05e9fef4e798ee9
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers/{voucherId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers/{voucherId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/deletevoucher
DELETE /proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers/{voucherId}
Delete Voucher
Remove a specific Hotspot voucher.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| voucherId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "vouchersDeleted": 0
}
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers/{voucherId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/hotspot/vouchers/{voucherId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getfirewallzone
GET /proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}
Get Firewall Zone
Get a firewall zone on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| firewallZoneId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "id": "ffcdb32c-6278-4364-8947-df4f77118df8",
  "name": "Hotspot|My custom zone",
  "networkIds": [
    "dfb21062-8ea0-4dca-b1d8-1eb3da00e58b"
  ],
  "metadata": {
    "origin": "string"
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updatefirewallzone
PUT /proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}
Update Firewall Zone
Update a firewall zone on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| firewallZoneId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "name": "Hotspot|My custom zone",
  "networkIds": [
    "dfb21062-8ea0-4dca-b1d8-1eb3da00e58b"
  ]
}
Response:
{
  "id": "ffcdb32c-6278-4364-8947-df4f77118df8",
  "name": "Hotspot|My custom zone",
  "networkIds": [
    "dfb21062-8ea0-4dca-b1d8-1eb3da00e58b"
  ],
  "metadata": {
    "origin": "string"
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "name": "Hotspot|My custom zone",
  "networkIds": [
    "dfb21062-8ea0-4dca-b1d8-1eb3da00e58b"
  ]
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "name": "Hotspot|My custom zone",
  "networkIds": [
    "dfb21062-8ea0-4dca-b1d8-1eb3da00e58b"
  ]
}'
Docs: https://developer.ui.com/network/v10.1.84/deletefirewallzone
DELETE /proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}
Delete Custom Firewall Zone
Delete a custom firewall zone from a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| firewallZoneId | ✓ | string |  | 
| siteId | ✓ | string |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/zones/{firewallZoneId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getfirewallpolicy
GET /proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}
Get Firewall Policy
Retrieve specific firewall policy.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| firewallPolicyId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "id": "00000000-0000-0000-0000-000000000000",
  "enabled": true,
  "name": "My firewall policy",
  "description": "A description for my firewall policy",
  "index": 0,
  "action": {
    "type": "string"
  },
  "source": {
    "zoneId": "00000000-0000-0000-0000-000000000000",
    "trafficFilter": {
      "type": "string"
    }
  },
  "destination": {
    "zoneId": "00000000-0000-0000-0000-000000000000",
    "trafficFilter": {
      "type": "string"
    }
  },
  "ipProtocolScope": {
    "ipVersion": "string"
  },
  "connectionStateFilter": [
    "NEW"
  ],
  "ipsecFilter": "MATCH_ENCRYPTED",
  "loggingEnabled": true,
  "schedule": {
    "mode": "string"
  },
  "metadata": {
    "origin": "string"
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updatefirewallpolicy
PUT /proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}
Update Firewall Policy
Update an existing firewall policy on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| firewallPolicyId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "enabled": true,
  "name": "My firewall policy",
  "description": "A description for my firewall policy",
  "action": {
    "type": "string"
  },
  "source": {
    "zoneId": "00000000-0000-0000-0000-000000000000",
    "trafficFilter": {
      "type": "string"
    }
  },
  "destination": {
    "zoneId": "00000000-0000-0000-0000-000000000000",
    "trafficFilter": {
      "type": "string"
    }
  },
  "ipProtocolScope": {
    "ipVersion": "string"
  },
  "connectionStateFilter": [
    "NEW"
  ],
  "ipsecFilter": "MATCH_ENCRYPTED",
  "loggingEnabled": true,
  "schedule": {
    "mode": "string"
  }
}
Response:
{
  "id": "00000000-0000-0000-0000-000000000000",
  "enabled": true,
  "name": "My firewall policy",
  "description": "A description for my firewall policy",
  "index": 0,
  "action": {
    "type": "string"
  },
  "source": {
    "zoneId": "00000000-0000-0000-0000-000000000000",
    "trafficFilter": {
      "type": "string"
    }
  },
  "destination": {
    "zoneId": "00000000-0000-0000-0000-000000000000",
    "trafficFilter": {
      "type": "string"
    }
  },
  "ipProtocolScope": {
    "ipVersion": "string"
  },
  "connectionStateFilter": [
    "NEW"
  ],
  "ipsecFilter": "MATCH_ENCRYPTED",
  "loggingEnabled": true,
  "schedule": {
    "mode": "string"
  },
  "metadata": {
    "origin": "string"
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
