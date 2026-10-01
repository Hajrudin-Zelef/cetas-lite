---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/clfigueiredo-hermes-infra-skills-blob-head-hermes-skills-forumtelecom-opnsense-o-319d6e57-1
title: "System status / firmware info, depending on version"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "license", "mcp", "research"]
source: docs/RAG/collect-261001-opnsense-pfsense/clfigueiredo-hermes-infra-skills-blob-head-hermes-skills-forumtelecom-opnsense-o-319d6e57.md
source_anchor: ""
source_lines: [1, 198]
sha256: 55e39abb36fd7592f84811d4b60a5bd91bb2a248802e6e1db2926600a6af7d4e
---

# System status / firmware info, depending on version

| name | opnsense-ops | 
|---|---|
| description | Senior OPNsense firewall engineer for ISP/MSP network operations. Use when the user asks to diagnose, audit, configure, or operate OPNsense firewalls via API, SSH/CLI, or web-GUI guidance: firewall rules, aliases, NAT, VLANs, interfaces, DHCP/Kea/dnsmasq, Unbound DNS, WireGuard/OpenVPN/IPsec status, HAProxy, gateways, routes, pf states/logs, config backup, firmware/plugins, and service health. Triggers include OPNsense, pfSense-like firewall, opn*, firewall rule, alias, NAT port forward, outbound NAT, VLAN OPNsense, Unbound, Kea DHCP, WireGuard OPNsense, HAProxy OPNsense, pfctl, configctl, filter reload, gateway status, CARP/HA. | 
| version | 1.0.0 | 
| author | Hermes Agent | 
| license | MIT | 
| platforms |  | 
| metadata |  | 

| freebsd | linux | 

| hermes | 
|---|
|  | 

| tags | related_skills | source_research | 
|---|---|---|
|  |  |  | 

| opnsense | firewall | networking | api | vpn | nat | dns | dhcp | 

| mikrotik-ops | cisco-ops | huawei-ne-ops | 

Senior firewall/network engineer for OPNsense. Speak Brazilian Portuguese with the user; keep OPNsense API paths, FreeBSD commands, `pfctl`, `configctl`, and firewall terminology in original syntax.

This is a Hermes-native operational playbook based on public OPNsense MCP/API projects and practical OPNsense workflows. It is not copied verbatim from another agent skill. It follows the Forum Telecom safety model: **Identify → Snapshot → Apply → Validate → Report**.

Prefer **API-first** for structured read/write operations when API credentials exist. Use SSH/CLI for diagnostics that the API does not expose (`pfctl`, packet capture, logs, configd status). Use web-GUI guidance only for operations with poor/no API coverage.

Public research strongly supports these safety principles:

- read-only default;
- explicit opt-in for writes;
- savepoint/rollback for firewall changes;
- hard blocklist for halt/reboot/firmware upgrade unless user explicitly confirms;
- never expose API key/secret or config secrets in chat output.

Typical connection fields:

| Variable | Purpose | 
|---|---|
| `OPNSENSE_URL` | GUI/API base URL, e.g. `https://10.0.0.1` or`https://fw.example.com:10443` | 
| `OPNSENSE_API_KEY` | API key from System > Access > Users | 
| `OPNSENSE_API_SECRET` | API secret; cannot be recovered after creation | 
| `OPNSENSE_VERIFY_SSL` | `true` for valid certificate,`false` for self-signed/internal | 
| `OPNSENSE_SSH_HOST` | SSH host/IP when CLI diagnostics are needed | 
| `OPNSENSE_SSH_USERNAME` | usually `root` or an admin user | 
| `OPNSENSE_SSH_KEY_PATH` / password | SSH auth material | 

API URLs usually use `/api/<module>/<controller>/<command>`. Some tools expect `OPNSENSE_URL` without `/api`; others expect it with `/api`. Verify the tool convention before calling.

Completion criterion: device version, API reachability, active interfaces, gateways, services, and target object are known.

API probes:

```
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/core/system/status"
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/diagnostics/interface/getInterfaceNames"
```
SSH/CLI probes:

```
opnsense-version
configctl system status
ifconfig
netstat -rn
pfctl -s info
pfctl -s rules
pfctl -s nat
pfctl -s state | head -50
```
For a request involving a rule/alias/NAT, always identify UUIDs before mutating:

```
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/firewall/filter/searchRule"
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/firewall/alias/searchItem"
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/firewall/nat/searchRule"
```
Completion criterion: current config/rule state is saved locally or an OPNsense savepoint exists.

API config backup, stripping sensitive content before reporting:

```
mkdir -p ~/opnsense-backups
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/core/backup/download/this" \
  -o ~/opnsense-backups/opnsense-$(date +%Y%m%d-%H%M%S).xml
```
Firewall savepoint pattern when supported:

```
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  -X POST "$OPNSENSE_URL/api/firewall/filter/savepoint"
```
If using SSH, snapshot `/conf/config.xml` before direct changes:

`cp /conf/config.xml /conf/config.xml.pre-hermes-$(date +%Y%m%d-%H%M%S)`
Do **not** paste full config XML into Telegram; it may contain secrets, certificates, VPN keys, tokens, and user hashes.

Completion criterion: only the requested delta was applied, using the narrowest API/CLI call.

Prefer API CRUD for aliases/rules/NAT. For firewall rules after a change:

```
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  -X POST "$OPNSENSE_URL/api/firewall/filter/apply"
```
Some APIs use `reconfigure` rather than `apply`:

```
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  -X POST "$OPNSENSE_URL/api/unbound/service/reconfigure"
```
For CLI-only diagnostics, prefer read-only commands. Avoid direct XML edits unless the API cannot do it and the user explicitly approves.

Completion criterion: API/CLI state and traffic diagnostics prove success.

```
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/firewall/filter/searchRule"
curl -sk -u "$OPNSENSE_API_KEY:$OPNSENSE_API_SECRET" \
  "$OPNSENSE_URL/api/diagnostics/firewall/log"
```
SSH/CLI validation:

```
pfctl -s rules | grep -i '<description-or-ip>'
pfctl -s nat | grep -i '<port-or-ip>'
pfctl -s state | grep '<ip>'
configctl filter reload
configctl service list
```
Use the report template at the end. Include commands and summarized results, not secrets.

Endpoint availability varies by OPNsense version and plugin. OPNsense 24.7+ has broader MVC API coverage; older versions may need SSH or GUI.

```
# System status / firmware info, depending on version
GET  /api/core/system/status
GET  /api/core/firmware/status
GET  /api/core/firmware/info
# Config backup
GET  /api/core/backup/download/this
# Services/controllers vary by plugin and version
GET  /api/core/service/search
POST /api/core/service/restart/<service>
```
Hard-block unless explicitly confirmed: halt, poweroff, reboot, firmware update/upgrade.

```
# Rules
GET  /api/firewall/filter/searchRule
GET  /api/firewall/filter/getRule/<uuid>
POST /api/firewall/filter/addRule
POST /api/firewall/filter/setRule/<uuid>
POST /api/firewall/filter/delRule/<uuid>
POST /api/firewall/filter/toggleRule/<uuid>
POST /api/firewall/filter/apply
POST /api/firewall/filter/savepoint
POST /api/firewall/filter/cancelRollback/<revision>
# Aliases
GET  /api/firewall/alias/searchItem
GET  /api/firewall/alias/getItem/<uuid>
POST /api/firewall/alias/addItem
POST /api/firewall/alias/setItem/<uuid>
POST /api/firewall/alias/delItem/<uuid>
POST /api/firewall/alias/reconfigure
```
Before deleting an alias, search firewall/NAT references. Do not remove aliases blindly.

```
GET  /api/firewall/nat/searchRule
GET  /api/firewall/nat/getRule/<uuid>
POST /api/firewall/nat/addRule
POST /api/firewall/nat/setRule/<uuid>
POST /api/firewall/nat/delRule/<uuid>
POST /api/firewall/nat/apply
# Outbound NAT settings/rules vary by version
GET  /api/firewall/nat/outbound/searchRule
GET  /api/firewall/nat/settings/get
```
NAT changes are high impact in provider networks. Validate both `pfctl -s nat` and live states/logs.

```
GET /api/diagnostics/interface/getInterfaceNames
GET /api/diagnostics/interface/getInterfaceStatistics
GET /api/diagnostics/interface/getArp
GET /api/routes/routes/searchRoutes
```
Interface assignments and VLAN creation may be API-limited. If API does not support the action, guide through GUI or use a vetted helper script only after backup and confirmation.

Unbound:

