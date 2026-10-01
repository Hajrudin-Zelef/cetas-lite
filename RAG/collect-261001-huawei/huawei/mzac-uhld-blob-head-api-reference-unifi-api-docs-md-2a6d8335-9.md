---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-9
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [2159, 2475]
sha256: fc3404847734e4c868c739d79069bf1901302cf59ce6d776cf0e93b7eeaffd8b
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

      "metadata": {
        "origin": "string"
      }
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/zones" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/zones" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/createfirewallzone
POST /proxy/network/integration/v1/sites/{siteId}/firewall/zones
Create Custom Firewall Zone
Create a new custom firewall zone on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
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
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/zones" \
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
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/zones" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "name": "Hotspot|My custom zone",
  "networkIds": [
    "dfb21062-8ea0-4dca-b1d8-1eb3da00e58b"
  ]
}'
Docs: https://developer.ui.com/network/v10.1.84/getfirewallpolicies
GET /proxy/network/integration/v1/sites/{siteId}/firewall/policies
List Firewall Policies
Retrieve a list of all firewall policies on a site.
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
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/createfirewallpolicy
POST /proxy/network/integration/v1/sites/{siteId}/firewall/policies
Create Firewall Policy
Create a new firewall policy on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
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
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/firewall/policies" \
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
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/firewall/policies" \
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
Docs: https://developer.ui.com/network/v10.1.84/getaclrule
GET /proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}
Get ACL Rule
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| aclRuleId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "type": "string",
  "id": "00000000-0000-0000-0000-000000000000",
  "enabled": true,
  "name": "string",
  "description": "string",
  "action": "ALLOW|BLOCK",
  "enforcingDeviceFilter": {
    "type": "string"
  },
  "index": 0,
  "metadata": {
    "origin": "string"
  }
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
