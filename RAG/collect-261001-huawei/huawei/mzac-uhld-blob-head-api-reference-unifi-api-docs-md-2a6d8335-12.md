---
id: collect-261001-huawei/huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335-12
title: "mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335.md
source_anchor: ""
source_lines: [3014, 3284]
sha256: 384c63ff5b2799c3d4b3bb3ac19153646e7751c9cb8a36b92e6bd2f5442f6819
---

# mzac-uhld-blob-head-api-reference-unifi-api-docs-md-2a6d8335

Docs: https://developer.ui.com/network/v10.1.84/gettrafficmatchinglists
GET /proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists
List Traffic Matching Lists
Retrieve all traffic matching lists on a site.
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
      "id": "ffcdb32c-6278-4364-8947-df4f77118df8",
      "name": "Allowed port list|Protected IP list"
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/createtrafficmatchinglist
POST /proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists
Create Traffic Matching List
Create a new traffic matching list on a site.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
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
  -X POST "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "name": "Allowed port list|Protected IP list"
}'
curl (cloud — via api.ui.com):
curl -sS -L \
  -X POST "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/traffic-matching-lists" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>" \
  -H "Content-Type: application/json" \
  -d '{
  "type": "string",
  "name": "Allowed port list|Protected IP list"
}'
Docs: https://developer.ui.com/network/v10.1.84/getwansoverviewpage
GET /proxy/network/integration/v1/sites/{siteId}/wans
List WAN Interfaces
Returns available WAN interface definitions for a given site, including identifiers and names. Useful for network and NAT configuration.
Path Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| siteId | ✓ | string |  | 
Query Parameters:
| Name | Required | Type | Description | 
|---|---|---|---|
| offset |  | integer |  | 
| limit |  | integer |  | 
Response:
{
  "offset": 0,
  "limit": 25,
  "count": 10,
  "totalCount": 1000,
  "data": [
    {
      "id": "00000000-0000-0000-0000-000000000000",
      "name": "Internet 1"
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/wans" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/wans" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getsitetositevpntunnelpage
GET /proxy/network/integration/v1/sites/{siteId}/vpn/site-to-site-tunnels
List Site-To-Site VPN Tunnels
Retrieve a paginated list of all site-to-site VPN tunnels on a site.
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
      "metadata": {
        "origin": "string"
      }
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/vpn/site-to-site-tunnels" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/vpn/site-to-site-tunnels" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getvpnserverpage
GET /proxy/network/integration/v1/sites/{siteId}/vpn/servers
List VPN Servers
Retrieve a paginated list of all VPN servers on a site.
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
      }
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/vpn/servers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/vpn/servers" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getradiusprofileoverviewpage
GET /proxy/network/integration/v1/sites/{siteId}/radius/profiles
List Radius Profiles
Returns available RADIUS authentication profiles, including configuration origin metadata.
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
      "name": "string",
      "metadata": {
        "origin": "string"
      }
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
  -X GET "https://192.168.1.1/proxy/network/integration/v1/sites/{siteId}/radius/profiles" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
curl (cloud — via api.ui.com):
curl -sS -L \
  -X GET "https://api.ui.com/proxy/network/integration/v1/sites/{siteId}/radius/profiles" \
  -H "Accept: application/json" \
  -H "X-API-Key: <your-api-key>"
Docs: https://developer.ui.com/network/v10.1.84/getdevicetagpage
GET /proxy/network/integration/v1/sites/{siteId}/device-tags
List Device Tags
Returns all device tags defined within a site, which can be used for WiFi Broadcast assignments.
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
      "name": "string",
      "deviceIds": [
        "00000000-0000-0000-0000-000000000000"
      ],
      "metadata": {
        "origin": "string"
      }
    }
  ]
}
curl (local — direct to controller):
curl -sS -L \
