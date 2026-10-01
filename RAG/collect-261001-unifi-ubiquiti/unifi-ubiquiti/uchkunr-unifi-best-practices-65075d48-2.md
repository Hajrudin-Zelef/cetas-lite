---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/uchkunr-unifi-best-practices-65075d48-2
title: "uchkunr-unifi-best-practices-65075d48"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/uchkunr-unifi-best-practices-65075d48.md
source_anchor: ""
source_lines: [3, 110]
sha256: d5020894a8a502dbadb434e15d3268567cf95e12a30a6197d6b85c5e94103814
---

# uchkunr-unifi-best-practices-65075d48

The UniFi API allows programmatic control over your entire UniFi network - devices, clients, WLANs, firewalls, vouchers, and more. Ubiquiti provides officially supported Cloud APIs, a secure Cloud Connector Proxy, and local REST endpoints, alongside the legacy internal controller API.
Depending on your architecture, select the appropriate Integration API:
| Flavor | Scope | Auth | Base URL / Path | Status | 
| Site Manager API | Cloud - monitor/manage deployments at scale | X-API-KEY | https://api.ui.com/v1/ | Official, GA | 
| Connector Proxy | Cloud-to-Local - secure remote access to local consoles | X-API-KEY | https://api.ui.com/v1/connector/consoles/{consoleId}/proxy/ | Official, firmware >= 5.0.3 | 
| Network Integration API | Local - per-application control on a local console | X-API-KEY | https://{console}/proxy/network/integration/ | Official, Network App >= 9.3 | 
| Classic Controller API | Local - full internal legacy interface | Cookie session | https://{controller}:8443/api/ or/proxy/network/api/ | Unofficial but stable | 
Prefer the Official Network Integration API (either directly or routed via the Connector Proxy) for new applications. Use the Classic API only for features not yet exposed officially (e.g. firewall rules, port forwards, routing).
Cloud Connector Proxy (Proxied Control)
Firmware version 5.0.3+ introduces the UniFi Connector Proxy. This allows your cloud applications to execute local API actions remotely by proxying the request through Ubiquiti's cloud endpoint to the physical console without exposing ports or establishing VPNs.
Proxy Routing Example (Axios)
Simply replace the base URL and supply the consoleId:
const client = axios.create({
  baseURL:
    "https://api.ui.com/v1/connector/consoles/YOUR_CONSOLE_ID/proxy/network/integration/v1",
  headers: {
    "X-API-KEY": process.env.UNIFI_API_KEY,
    Accept: "application/json",
  },
});
// Retrieves local network sites securely via the Cloud Proxy
const { data } = await client.get("/sites");
- Site Structure: Most endpoints require a site identifier (default is the primary site).
- Response Format: JSON with standard structure { "meta": { "rc": "ok" }, "data": [...] } . On error:{ "meta": { "rc": "error", "msg": "api.err.Invalid" } } .
- MAC Addresses: Use lowercase, typically without separators (e.g., 00112233aabb ).
- IDs: Most objects use a MongoDB-style _id field.
Official API Key (X-API-KEY)
Available on UniFi OS-based consoles (UDM, UDR, UCG, UX, UDW, UCG-Ultra, UniFi OS Server). The key inherits the creator admin's permissions.
Generating a key:
- Cloud: unifi.ui.com → your profile → API Keys
- Local: UniFi Network → Settings → Control Plane → Integrations
Stateless - no login/logout flow required. Ideal for scripts, CI, and long-running services.
Classic Login (Cookie Session)
| Operation | Method | Endpoint | Description | 
| Login (legacy) | POST | /api/login | Standalone controller | 
| Login (UniFi OS) | POST | /api/auth/login | UDM / UCG / UniFi OS Server | 
| Verify Session | GET | /api/self | Check session validity | 
| Logout | POST | /api/logout //api/auth/logout | End session | 
{
  "username": "admin",
  "password": "your_password",
  "remember": true
}
UniFi OS vs Standalone Controller (Classic)
| Aspect | Standalone (software controller) | UniFi OS (UDM/UCG/UDR) | 
| Port | 8443 | 443 (or11443 for UniFi OS Server) | 
| Login | /api/login | /api/auth/login | 
| API prefix | /api/... | /proxy/network/api/... | 
| CSRF | Not required | Required on state-changing calls | 
On UniFi OS, the login response includes a X-CSRF-Token header (or TOKEN cookie). Include this header on every POST/PUT/DELETE:
api.defaults.headers.common["X-CSRF-Token"] = loginResp.headers["x-csrf-token"];
If the admin has 2FA enabled, the first login returns HTTP 499 with meta.msg = "api.err.Ubic2faTokenRequired". Resubmit with:
{ "username": "admin", "password": "pw", "token": "123456" }
Official Site Manager API (Cloud)
Base URL: https://api.ui.com/v1/ · Auth: X-API-KEY · Rate limit: 10,000 req/min (GA), 429 with Retry-After header on overflow.
| Operation | Method | Endpoint | Description | 
| List Hosts | GET | /hosts | All UniFi OS hosts on the account | 
| Host Details | GET | /hosts/{id} | Host metadata, status, uptime | 
| List Sites | GET | /sites | All sites across hosts | 
| List Devices | GET | /devices | Devices across every site | 
| ISP Metrics | GET | /isp-metrics | ISP performance samples | 
| Query ISP Metrics | POST | /isp-metrics/query | Filtered ISP metric query | 
| List SD-WAN | GET | /sdwan-configs | SD-WAN configurations | 
| SD-WAN Details | GET | /sdwan-configs/{id} | Specific SD-WAN config | 
| SD-WAN Status | GET | /sdwan-configs/{id}/status | Deployment status | 
Official Network Integration API (Local & Proxied)
API paths are relative to /v1 locally or proxied through /connector/consoles/{id}/proxy/network/integration/v1.
| Operation | Method | Endpoint | Description | 
| App Info | GET | /info | Controller version ( applicationVersion ) & capabilities | 
| List Sites | GET | /sites | Paginated list of local sites | 
| List Devices | GET | /sites/{siteId}/devices | Devices associated with the site | 
| Device Details | GET | /sites/{siteId}/devices/{deviceId} | Single device info | 
| Device Stats | GET | /sites/{siteId}/devices/{deviceId}/statistics/latest | Real-time statistics | 
| Adopt Device | POST | /sites/{siteId}/devices | Programmatic device adoption | 
| Device Action | POST | /sites/{siteId}/devices/{deviceId}/actions | Trigger actions (e.g. RESTART ) | 
| Port Action | POST | /sites/{siteId}/devices/{deviceId}/interfaces/ports/{idx}/actions | Power-cycle switch ports | 
| List Clients | GET | /sites/{siteId}/clients | Connected clients | 
| Client Details | GET | /sites/{siteId}/clients/{clientId} | Specific client | 
| Client Action | POST | /sites/{siteId}/clients/{clientId}/actions | AUTHORIZE_GUEST_ACCESS , etc. | 
| WiFi Broadcast Info | GET | /sites/{siteId}/wifi/broadcasts/{id} | SSID, MLO, & client filtering policies | 
| List Vouchers | GET | /sites/{siteId}/hotspot/vouchers | Hotspot vouchers | 
| Generate Vouchers | POST | /sites/{siteId}/hotspot/vouchers | Batch create vouchers | 
| Delete Voucher | DELETE | /sites/{siteId}/hotspot/vouchers/{id} | Remove voucher | 
Classic API Endpoint Reference
All endpoints below are prefixed with /api/s/{site}/ on a standalone controller, or /proxy/network/api/s/{site}/ on UniFi OS.
| Operation | Method | Endpoint | Description | 
| List All Devices | GET | /stat/device | Full device payload | 
| Basic Device List | GET | /stat/device-basic | Lightweight list | 
| Device Details | GET | /stat/device/{mac} | Specific device | 
| Device Commands | POST | /cmd/devmgr | e.g. adopt, restart, upgrade, power-cycle | 
| Locate (LED) | POST | /cmd/devmgr | set-locate /unset-locate | 
{ "cmd": "restart", "mac": "00:11:22:33:44:55" }
| Operation | Method | Endpoint | Description | 
| Active Clients | GET | /stat/sta | Connected clients | 
| All Clients | GET | /stat/alluser | Historical client logs | 
| Client Details | GET | /stat/user/{mac} | Specific client | 
| Known Clients | GET | /rest/user | Configured clients | 
| Block / Unblock / Kick | POST | /cmd/stamgr | block-sta /unblock-sta /kick-sta | 
| Authorize Guest | POST | /cmd/stamgr | authorize-guest | 
| Unauthorize Guest | POST | /cmd/stamgr | unauthorize-guest | 
| Forget Client | POST | /cmd/stamgr | forget-sta | 
{ "cmd": "authorize-guest", "mac": "aa:bb:cc:dd:ee:ff", "minutes": 60 }
| Operation | Method | Endpoint | Description | 
| Update Client | PUT | /rest/user/{client_id} | Set fixed IP, group, access | 
| Traffic Rules | GET/POST/PUT/DELETE | /rest/trafficrule | L7/bandwidth rules | 
| Traffic Routes | GET/POST/PUT/DELETE | /rest/trafficroute | Policy-based routing | 
| Operation | Method | Endpoint | Description | 
| List Networks | GET | /rest/networkconf | All LAN/VLAN networks | 
