---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/uchkunr-unifi-best-practices-65075d48-3
title: "uchkunr-unifi-best-practices-65075d48"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/uchkunr-unifi-best-practices-65075d48.md
source_anchor: ""
source_lines: [111, 203]
sha256: f3f5b284951b4063344bc762c39b96b773e25bf8105dc561585390128d3fd513
---

# uchkunr-unifi-best-practices-65075d48

| Create Network | POST | /rest/networkconf | Add network | 
| Update Network | PUT | /rest/networkconf/{id} | Modify | 
| Delete Network | DELETE | /rest/networkconf/{id} | Remove | 
| Operation | Method | Endpoint | Description | 
| List WLANs | GET | /rest/wlanconf | All SSIDs | 
| WLAN Details | GET | /rest/wlanconf/{id} | Specific WLAN | 
| Create WLAN | POST | /rest/wlanconf | New SSID | 
| Update WLAN | PUT | /rest/wlanconf/{id} | Modify SSID | 
| Delete WLAN | DELETE | /rest/wlanconf/{id} | Remove SSID | 
| Operation | Method | Endpoint | Description | 
| List Rules | GET | /rest/firewallrule | User-defined rules | 
| Update Rule | PUT | /rest/firewallrule/{id} | Enable / edit | 
| List Groups | GET | /rest/firewallgroup | Address/port groups | 
| Update Group | PUT | /rest/firewallgroup/{id} | Edit group members | 
| IPS/IDS Settings | GET/PUT | /rest/setting/ips | Threat management | 
| Operation | Method | Endpoint | Description | 
| List Forwards | GET | /rest/portforward | All forward rules | 
| Create Forward | POST | /rest/portforward | New rule | 
| Toggle Forward | PUT | /rest/portforward/{rule-id} | Enable / disable | 
| Delete Forward | DELETE | /rest/portforward/{rule-id} | Remove | 
| Port Profiles | GET/PUT | /rest/portconf | Switch port profiles | 
| Operation | Method | Endpoint | Description | 
| Static Routes | GET/POST/PUT/DELETE | /rest/routing | User-defined routes | 
| DynamicDNS Config | GET/POST/PUT | /rest/dynamicdns | DDNS entries | 
| Operation | Method | Endpoint | Description | 
| List Sites | GET | /api/self/sites | Accessible sites | 
| Sites + Health | GET | /api/stat/sites | Health / alert summary | 
| Site Health | GET | /stat/health | Per-subsystem health | 
| Create Site | POST | /api/s/default/cmd/sitemgr | add-site | 
| Delete Site | POST | /api/s/default/cmd/sitemgr | delete-site | 
| Operation | Method | Endpoint | Description | 
| List Vouchers | GET | /stat/voucher | All vouchers | 
| Create Voucher | POST | /cmd/hotspot | create-voucher | 
| Delete Voucher | POST | /cmd/hotspot | delete-voucher | 
| Guest Operators | GET | /rest/hotspotop | Hotspot operators | 
| Payments | GET | /stat/payment | Guest payments | 
| Operation | Method | Endpoint | Description | 
| List Events | GET | /stat/event | Site events (capped at 3000) | 
| List Alarms | GET | /stat/alarm | Active alarms | 
| Archive Alarm | POST | /cmd/evtmgr | archive-alarm | 
| Operation | Method | Endpoint | Description | 
| List Backups | POST | /cmd/backup | list-backups | 
| Create Backup | POST | /cmd/backup | backup | 
| Delete Backup | POST | /cmd/backup | delete-backup | 
| Download Backup | GET | /dl/backup/{filename} | Raw download | 
| Operation | Method | Endpoint | Description | 
| DPI Stats | GET | /stat/dpi | Deep Packet Inspection | 
| Client DPI | GET | /stat/stadpi | Per-client DPI | 
| Daily Report | GET | /stat/report/daily.site | Daily metrics | 
| Hourly Report | GET | /stat/report/hourly.site | Hourly metrics | 
| 5-min Report | GET | /stat/report/5minutes.site | 5-minute granularity | 
| Gateway Stats | GET | /stat/gateway | Gateway performance | 
| Rogue APs | GET | /stat/rogueap | Detected rogue access points | 
| Operation | Method | Endpoint | Description | 
| List Admins | GET | /api/stat/admin | All admins & roles | 
| Invite Admin | POST | /cmd/sitemgr | invite-admin | 
| Revoke Admin | POST | /cmd/sitemgr | revoke-admin | 
| Parameter | Description | Example | 
| _limit | Max items to return | /stat/sta?_limit=50 | 
| _start | Offset | /stat/sta?_start=50 | 
| _sort | Sort field (prefix - for descending) | /stat/sta?_sort=-last_seen | 
| mac | Filter by MAC | /stat/sta?mac=001122334455 | 
| ip | Filter by IP | /stat/sta?ip=192.168.1.100 | 
| within | Historical window in hours | /stat/sta?within=24 | 
| attrs | Comma-separated attributes to include | /stat/sta?attrs=mac,ip,hostname | 
const response = await api.get("/api/s/default/stat/sta", {
  params: {
    _limit: 10,
    _sort: "-last_seen",
    attrs: "mac,hostname,ip,signal,tx_bytes,rx_bytes",
  },
});
Rate Limiting & Error Handling
- Official APIs enforce per-minute quotas (10k/min for Site Manager GA). Overflow returns 429 with an RFC-compliantRetry-After header.
- Classic API has no documented quota but will return 401 on stale sessions - re-login onmeta.rc = "error" withapi.err.LoginRequired .
- Common error codes (meta.msg ):
  - api.err.Invalid - bad request body
  - api.err.LoginRequired - session expired
  - api.err.NoPermission - admin role too low
  - api.err.Ubic2faTokenRequired - 2FA challenge
try {
  await client.post(...);
} catch (err) {
  if (err.response?.status === 429) {
    const wait = Number(err.response.headers["retry-after"] ?? 1);
    await new Promise(r => setTimeout(r, wait * 1000));
  }
}
WebSocket Events (Real-time)
The controller exposes a real-time feed:
- Standalone: wss://{controller}:8443/wss/s/{site}/events
- UniFi OS: wss://{console}/proxy/network/wss/s/{site}/events
Send the session cookie (or X-API-KEY where supported) on the upgrade request. Useful for live client join/leave, device state, alarm push, and voucher consumption without polling /stat/event.
