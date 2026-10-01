---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-11
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [2766, 3013]
sha256: 98327aaab60476d8795d5c8016acb33444d128b9909ff5bbeb9c945c9f420051
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

  "enabled": true,
  "metadata": {
    "origin": "string"
  },
  "domain": "string"
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updatednspolicy
PUT /proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}
Update DNS Policy
Update an existing DNS policy on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| dnsPolicyId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "type": "string",
  "enabled": true
}
Response:
{
  "type": "string",
  "id": "00000000-0000-0000-0000-000000000000",
  "enabled": true,
  "metadata": {
    "origin": "string"
  },
  "domain": "string"
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true
}'
Docs: https://developer.ui.com/network/v10.1.84/deletednspolicy
DELETE /proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}
Delete DNS Policy
Delete an existing DNS policy on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| dnsPolicyId | ✓ | string |  | 
| siteId | ✓ | string |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/dns/policies/{dnsPolicyId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getdnspolicypage
GET /proxy/network/integration/v1/sites/{siteId}/dns/policies
List DNS Policies
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
      "metadata": {
        "origin": "string"
      },
      "domain": "string"
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/dns/policies" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/dns/policies" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/creatednspolicy
POST /proxy/network/integration/v1/sites/{siteId}/dns/policies
Create DNS Policy
Create a new DNS policy on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Request Body:
{
  "type": "string",
  "enabled": true
}
Response:
{
  "type": "string",
  "id": "00000000-0000-0000-0000-000000000000",
  "enabled": true,
  "metadata": {
    "origin": "string"
  },
  "domain": "string"
}
curl (local — direct to controller):
curl -sS -L \
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/dns/policies" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/dns/policies" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "enabled": true
}'
Docs: https://developer.ui.com/network/v10.1.84/gettrafficmatchinglist
GET /proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}
Get Traffic Matching List
Get an exist traffic matching list on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| trafficMatchingListId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Response:
{
  "type": "string",
  "id": "ffcdb32c-6278-4364-8947-df4f77118df8",
  "name": "Allowed port list|Protected IP list"
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/updatetrafficmatchinglist
PUT /proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}
Update Traffic Matching List
Update an exist traffic matching list on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| trafficMatchingListId | ✓ | string |  | 
| siteId | ✓ | string |  | 
Request Body:
{
  "type": "string",
  "name": "Allowed port list|Protected IP list"
}
Response:
{
  "type": "string",
  "id": "ffcdb32c-6278-4364-8947-df4f77118df8",
  "name": "Allowed port list|Protected IP list"
}
curl (local — direct to controller):
curl -sS -L \
  -X PUT "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "name": "Allowed port list|Protected IP list"
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X PUT "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "name": "Allowed port list|Protected IP list"
}'
Docs: https://developer.ui.com/network/v10.1.84/deletetrafficmatchinglist
DELETE /proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}
Delete Traffic Matching List
Delete an exist traffic matching list on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| trafficMatchingListId | ✓ | string |  | 
| siteId | ✓ | string |  | 
curl (local — direct to controller):
curl -sS -L \
  -X DELETE "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X DELETE "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists/{trafficMatchingListId}" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
