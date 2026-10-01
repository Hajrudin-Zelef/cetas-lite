---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-10
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [2476, 2765]
sha256: c83ed977599c290983eac21d604dbf60607dfc4ee1105fefbfa2ae5e723fd6f8
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updateaclrule
PUT /proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}
Update ACL Rule
Update an existing user defined ACL rule on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| aclRuleId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "type": "string",
  "enabled": true,
  "name": "string",
  "description": "string",
  "action": "ALLOW|BLOCK",
  "enforcingDeviceFilter": {
    "type": "string"
  },
  "index": 0
}
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
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true,
  "name": "string",
  "description": "string",
  "action": "ALLOW|BLOCK",
  "enforcingDeviceFilter": {
    "type": "string"
  },
  "index": 0
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true,
  "name": "string",
  "description": "string",
  "action": "ALLOW|BLOCK",
  "enforcingDeviceFilter": {
    "type": "string"
  },
  "index": 0
}'
Docs: https://developer.ui.com/network/v10.1.84/deleteaclrule
DELETE /proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}
Delete ACL Rule
Delete an existing user defined ACL rule on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| aclRuleId | ✓ | string |  | 
| siteId | ✓ | string |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/acl-rules/{aclRuleId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getaclruleordering
GET /proxy/network/integration/v1/sites/{siteId}/acl-rules/ordering
Get User-Defined ACL Rule Ordering
Retrieve user-defined ACL rule ordering on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Response:
{
  "orderedAclRuleIds": [
    "00000000-0000-0000-0000-000000000000"
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/acl-rules/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/acl-rules/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updateaclruleordering
PUT /proxy/network/integration/v1/sites/{siteId}/acl-rules/ordering
Reorder User-Defined ACL Rules
Reorder user-defined ACL rules on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Request Body:
{
  "orderedAclRuleIds": [
    "00000000-0000-0000-0000-000000000000"
  ]
}
Response:
{
  "orderedAclRuleIds": [
    "00000000-0000-0000-0000-000000000000"
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/acl-rules/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "orderedAclRuleIds": [
    "00000000-0000-0000-0000-000000000000"
  ]
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/acl-rules/ordering" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "orderedAclRuleIds": [
    "00000000-0000-0000-0000-000000000000"
  ]
}'
Docs: https://developer.ui.com/network/v10.1.84/getaclrulepage
GET /proxy/network/integration/v1/sites/{siteId}/acl-rules
List ACL Rules
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
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/acl-rules" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/acl-rules" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/createaclrule
POST /proxy/network/integration/v1/sites/{siteId}/acl-rules
Create ACL Rule
Create a new user defined ACL rule on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Request Body:
{
  "type": "string",
  "enabled": true,
  "name": "string",
  "description": "string",
  "action": "ALLOW|BLOCK",
  "enforcingDeviceFilter": {
    "type": "string"
  },
  "index": 0
}
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
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/acl-rules" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true,
  "name": "string",
  "description": "string",
  "action": "ALLOW|BLOCK",
  "enforcingDeviceFilter": {
    "type": "string"
  },
  "index": 0
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/acl-rules" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true,
  "name": "string",
  "description": "string",
  "action": "ALLOW|BLOCK",
  "enforcingDeviceFilter": {
    "type": "string"
  },
  "index": 0
}'
Docs: https://developer.ui.com/network/v10.1.84/getdnspolicy
GET /proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}
Get DNS Policy
Retrieve specific DNS policy.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| dnsPolicyId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "type": "string",
  "id": "00000000-0000-0000-0000-000000000000",
