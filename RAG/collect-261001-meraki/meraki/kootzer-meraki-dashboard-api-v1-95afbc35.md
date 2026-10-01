---
id: collect-261001-meraki/meraki/kootzer-meraki-dashboard-api-v1-95afbc35
title: "kootzer-meraki-dashboard-api-v1-95afbc35"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "license", "parameters"]
source: docs/RAG/collect-261001-meraki/kootzer-meraki-dashboard-api-v1-95afbc35.md
source_anchor: ""
source_lines: [1, 86]
sha256: 62461e274f5efa1384924e1492b03be69830006dfd1e9cdc1b5dcff70a16c3ff
---

# kootzer-meraki-dashboard-api-v1-95afbc35

A Node-RED node for interacting with the Cisco Meraki Dashboard API v1.
This node provides access to 992 API endpoints covering the full Meraki Dashboard API v1.73.0, including devices, networks, organizations, appliances, switches, wireless, cameras, sensors, and more. All endpoints utilize OpenAPI Specification version 3, pulled directly from Meraki's spec3.json file from GitHub.
Note: This endpoint count reflects Meraki's stable, generally-available (GA) API spec only. Cisco's changelog may cite a higher total that also includes Early API Access (beta) operations, which are excluded here since they're subject to breaking changes without notice.
- All Meraki Dashboard API v1 endpoints available from a single node
- Cascading category/sub-category/filter selectors for easy endpoint discovery
- Parameters configurable in the node editor or passed dynamically via msg properties
- API host and key overridable per-message for multi-org workflows
- Body examples shown inline for POST/PUT operations
- Data-driven architecture — endpoints defined in a single JSON file for easy maintenance
Install from npm in your Node-RED user directory (typically ~/.node-red):
npm install @kootzer/meraki-dashboard-api-v1
Or, to install from a local copy of this folder instead of npm:
Copy the 'meraki-dashboard-api-v1' folder into your Node-RED modules directory.
Default location: C:\Users\username\.node-red\node_modules\@kootzer\meraki-dashboard-api-v1
Then restart Node-RED.
- Drag the Meraki Dashboard API v1 node onto your flow
- Double-click to configure
- Add a new Service configuration:
  - API URL: https://api.meraki.com/api/v1 (default)
  - Header Name: X-Cisco-Meraki-API-Key
  - API Key: Your Meraki Dashboard API key
- API URL: 
- Select a Category, Sub-Category, and Method
- Fill in the required parameters
- Log in to the Meraki Dashboard
- Navigate to Organization → API & Webhooks
- Under API keys and access, generate your API key
See the Meraki API documentation for more details.
| Property | Type | Description | 
|---|---|---|
| msg.payload | object | For POST/PUT methods, the request body as a JSON object | 
| msg.service | object | Optional. Override API host and/or key: { host: "https://...", apiKey: "..." } | 
| msg.operationId | string | Optional. Dynamically select which API method to call | 
| msg.<paramName> | string | Any parameter can be passed as a message property (e.g., msg.serial ,msg.networkId ) | 
| Property | Type | Description | 
|---|---|---|
| msg.payload | object | The API response body | 
| msg.statusCode | number | HTTP status code | 
| msg.headers | object | Response headers | 
| msg.responseUrl | string | Final URL after redirects | 
A ready-to-import example flow ("Meraki API Examples") is included with this package and appears under Import → Examples → @kootzer/meraki-dashboard-api-v1 in the Node-RED editor. It contains several grouped, read-only (GET) example use cases:
| Group | Demonstrates | 
|---|---|
| Return API Key Owner | The simplest possible call: getAdministeredIdentitiesMe with no parameters. Good first test that your API key and Service config are working. | 
| Return API Key Owner's Organizations | Listing every organization the API key has access to ( getOrganizations ). | 
| Find Network ID by Name | Looking up a network's ID from its name ( getOrganizationNetworks + aswitch node), useful if you only know a network by name. | 
| Return All Organization Networks | Listing every network in an organization. | 
| Org-Wide Client Search | Searching for clients across an organization. | 
| Return All Uplinks Within an Organization | Listing uplink status for every device in an organization. | 
| Return All Organization Network Health Alerts | Listing active assurance alerts across an organization's networks. | 
| Return Organization Device Inventory | Listing an organization's full device inventory. | 
Every Meraki node in the example ships with no Service config selected — Node-RED will flag them as needing configuration until you select or create your own Service config with a valid API key.
This node uses a data-driven architecture. All API endpoints are defined in endpoints.json and the code is fully generic. No per-endpoint logic exists in the source files.
To update to the latest Meraki API version:
cd node_modules/@kootzer/meraki-dashboard-api-v1
node update_from_openapi.js
node build.js
This downloads the latest Meraki OpenAPI spec from GitHub, regenerates endpoints.json, and rebuilds node.html. Restart Node-RED to pick up the changes.
| File | Purpose | 
|---|---|
| endpoints.json | All API endpoint definitions (operationId, method, path, params, tags, summary) | 
| lib.js | Generic API client with a single callEndpoint() method | 
| node.js | Node-RED node logic — generic parameter resolution and API invocation | 
| node.html | Node-RED editor UI — dynamically generated from endpoint data | 
| node.html.template | Template for node.html with placeholder for endpoint data | 
| build.js | Build script that injects endpoints.json into the HTML template | 
| update_from_openapi.js | Downloads the latest Meraki OpenAPI spec and regenerates endpoints.json | 
If you need to add an endpoint manually (e.g., a beta or private API), add an entry to endpoints.json:
{
    "operationId": "getMyCustomEndpoint",
    "method": "GET",
    "path": "/organizations/{organizationId}/custom/endpoint",
    "params": [
        { "name": "organizationId", "in": "path", "type": "string", "required": true },
        { "name": "perPage", "in": "query", "type": "string", "required": false }
    ],
    "tags": ["organizations", "custom"],
    "summary": "Description of what this endpoint does"
}
Then run node build.js and restart Node-RED.
This node has zero runtime dependencies — the HTTP client is built on Node's native fetch. Requires Node.js 18 or later (the same minimum Node-RED itself requires as of v3.0).
This project is licensed under the Apache License 2.0.
See the LICENSE file for details.
Maintained by Austin Kutzer.
This node began as a fork of Cory Guynn's original node-red-contrib-meraki-dashboard-api-v1. Since then it's been rewritten with a generic, data-driven architecture: a single callEndpoint() method driven by endpoints.json in place of the original's ~500 individually generated methods, cutting the runtime code from roughly 68,000 lines to under 500 while adding full OpenAPI v3 coverage and nearly double the endpoints. Credit to Cory for the original concept and groundwork.
The Meraki Dashboard API is developed and maintained by Cisco Meraki. This node is a community project and is not officially supported by Cisco.
