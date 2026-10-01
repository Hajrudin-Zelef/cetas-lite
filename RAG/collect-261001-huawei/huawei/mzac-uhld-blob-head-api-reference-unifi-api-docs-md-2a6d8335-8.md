---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-8
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [1875, 2158]
sha256: 845471f54b8293b3f942e98a145b3e477f6abc3bbe6812b6c1a8fe582b35957f
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

  -H "Content-Type: application/json" \
  -d '{
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
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
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
}'
Docs: https://developer.ui.com/network/v10.1.84/deletefirewallpolicy
DELETE /proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}
Delete Firewall Policy
Delete an existing firewall policy on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| firewallPolicyId | ✓ | string |  | 
| siteId | ✓ | string |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/patchfirewallpolicy
PATCH /proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}
Patch Firewall Policy
Patch an existing firewall policy on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| firewallPolicyId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "loggingEnabled": true
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
  -X PATCH "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "loggingEnabled": true
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PATCH "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies/{firewallPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "loggingEnabled": true
}'
Docs: https://developer.ui.com/network/v10.1.84/getfirewallpolicyordering
GET /proxy/network/integration/v1/sites/{siteId}/firewall/policies/ordering
Get User-Defined Firewall Policy Ordering
Retrieve user-defined firewall policy ordering for a specific source/destination zone pair.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| sourceFirewallZoneId | ✓ | string |  | 
| destinationFirewallZoneId | ✓ | string |  | 
Response:
{
  "orderedFirewallPolicyIds": {
    "beforeSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ],
    "afterSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ]
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updatefirewallpolicyordering
PUT /proxy/network/integration/v1/sites/{siteId}/firewall/policies/ordering
Reorder User-Defined Firewall Policies
Reorder user-defined firewall policies for a specific source/destination zone pair.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| sourceFirewallZoneId | ✓ | string |  | 
| destinationFirewallZoneId | ✓ | string |  | 
Request Body:
{
  "orderedFirewallPolicyIds": {
    "beforeSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ],
    "afterSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ]
  }
}
Response:
{
  "orderedFirewallPolicyIds": {
    "beforeSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ],
    "afterSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ]
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "orderedFirewallPolicyIds": {
    "beforeSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ],
    "afterSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ]
  }
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "orderedFirewallPolicyIds": {
    "beforeSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ],
    "afterSystemDefined": [
      "00000000-0000-0000-0000-000000000000"
    ]
  }
}'
Docs: https://developer.ui.com/network/v10.1.84/getfirewallzones
GET /proxy/network/integration/v1/sites/{siteId}/firewall/zones
List Firewall Zones
Retrieve a list of all firewall zones on a site.
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
      "id": "ffcdb32c-6278-4364-8947-df4f77118df8",
      "name": "Hotspot|My custom zone",
      "networkIds": [
        "dfb21062-8ea0-4dca-b1d8-1eb3da00e58b"
      ],
