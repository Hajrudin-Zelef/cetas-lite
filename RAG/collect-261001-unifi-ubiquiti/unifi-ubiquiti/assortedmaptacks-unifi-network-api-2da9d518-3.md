---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/assortedmaptacks-unifi-network-api-2da9d518-3
title: "Clone the repository"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/assortedmaptacks-unifi-network-api-2da9d518.md
source_anchor: ""
source_lines: [236, 355]
sha256: 9171e8a346bc6303487ba6eeb27b75201afd1587c1b3845ffe56af24ec83ca2c
---

# Clone the repository

| QoS | /v2/api/site/{site}/qos-rules | GET | List QoS Rules | 
|  | /v2/api/site/{site}/qos-rules | POST | Create QoS Rule | 
|  | /v2/api/site/{site}/qos-rules/{id} | GET | Get QoS Rule Details | 
|  | /v2/api/site/{site}/qos-rules/{id} | PUT | Update QoS Rule | 
|  | /v2/api/site/{site}/qos-rules/{id} | DELETE | Delete QoS Rule | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Firewall | /v2/api/site/{site}/firewall-rules | GET | List Firewall Rules | 
|  | /v2/api/site/{site}/firewall-rules | POST | Create Firewall Rule | 
|  | /v2/api/site/{site}/firewall-rules/{id} | GET | Get Firewall Rule Details | 
|  | /v2/api/site/{site}/firewall-rules/{id} | PUT | Update Firewall Rule | 
|  | /v2/api/site/{site}/firewall-rules/{id} | DELETE | Delete Firewall Rule | 
|  | /v2/api/site/{site}/firewall-groups | GET | List Firewall Groups | 
|  | /v2/api/site/{site}/firewall-groups | POST | Create Firewall Group | 
|  | /v2/api/site/{site}/firewall-groups/{id} | GET | Get Firewall Group Details | 
|  | /v2/api/site/{site}/firewall-groups/{id} | PUT | Update Firewall Group | 
|  | /v2/api/site/{site}/firewall-groups/{id} | DELETE | Delete Firewall Group | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Port Profiles | /v2/api/site/{site}/port-profiles | GET | List Port Profiles | 
|  | /v2/api/site/{site}/port-profiles | POST | Create Port Profile | 
|  | /v2/api/site/{site}/port-profiles/{id} | GET | Get Port Profile Details | 
|  | /v2/api/site/{site}/port-profiles/{id} | PUT | Update Port Profile | 
|  | /v2/api/site/{site}/port-profiles/{id} | DELETE | Delete Port Profile | 
|  | /v2/api/site/{site}/port-profiles/defaults | GET | Get Port Profile Defaults | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Port Management | /v2/api/site/{site}/devices/{mac}/ports/{portIndex} | GET | Get Port Configuration | 
|  | /v2/api/site/{site}/devices/{mac}/ports/{portIndex} | PUT | Update Port Configuration | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Routing | /v2/api/site/{site}/static-routes | GET | List Static Routes | 
|  | /v2/api/site/{site}/static-routes | POST | Create Static Route | 
|  | /v2/api/site/{site}/static-routes/{id} | GET | Get Static Route Details | 
|  | /v2/api/site/{site}/static-routes/{id} | PUT | Update Static Route | 
|  | /v2/api/site/{site}/static-routes/{id} | DELETE | Delete Static Route | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| WAN | /v2/api/site/{site}/port-forward-advanced | GET | List Advanced Port Forwarding Rules | 
|  | /v2/api/site/{site}/port-forward-advanced | POST | Create Advanced Port Forwarding Rule | 
|  | /v2/api/site/{site}/port-forward-advanced/{id} | GET | Get Advanced Port Forwarding Rule Details | 
|  | /v2/api/site/{site}/port-forward-advanced/{id} | PUT | Update Advanced Port Forwarding Rule | 
|  | /v2/api/site/{site}/port-forward-advanced/{id} | DELETE | Delete Advanced Port Forwarding Rule | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Floorplan | /v2/api/site/{site}/maps | GET | List Site Maps | 
|  | /v2/api/site/{site}/maps | POST | Create Site Map | 
|  | /v2/api/site/{site}/maps/{id} | GET | Get Site Map Details | 
|  | /v2/api/site/{site}/maps/{id} | PUT | Update Site Map | 
|  | /v2/api/site/{site}/maps/{id} | DELETE | Delete Site Map | 
|  | /v2/api/site/{site}/walls | GET | List Walls | 
|  | /v2/api/site/{site}/walls | POST | Create Wall | 
|  | /v2/api/site/{site}/walls/{id} | GET | Get Wall Details | 
|  | /v2/api/site/{site}/walls/{id} | PUT | Update Wall | 
|  | /v2/api/site/{site}/walls/{id} | DELETE | Delete Wall | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Legacy Config | /api/s/{site}/stat/sites | GET | List Sites Statistics | 
|  | /api/s/{site}/stat/device | GET | List Device Statistics | 
|  | /api/s/{site}/stat/client | GET | List Client Statistics | 
|  | /api/s/{site}/stat/sessions | GET | Get Session Statistics | 
|  | /api/s/{site}/list/user | GET | List Users | 
|  | /api/s/{site}/list/usergroup | GET | List User Groups | 
|  | /api/s/{site}/list/wlanconf | GET | List WLAN Configurations | 
|  | /api/s/{site}/list/networkconf | GET | List Network Configurations | 
|  | /api/s/{site}/list/firewallrule | GET | List Firewall Rules | 
|  | /api/s/{site}/list/firewallgroup | GET | List Firewall Groups | 
|  | /api/s/{site}/list/portconf | GET | List Port Configurations | 
|  | /api/s/{site}/list/radiusprofile | GET | List RADIUS Profiles | 
|  | /api/s/{site}/list/account | GET | List Accounts | 
|  | /api/s/{site}/list/wlangroup | GET | List WLAN Groups | 
|  | /api/s/{site}/list/rogueap | GET | List Rogue Access Points | 
|  | /api/s/{site}/list/known_rogueap | GET | List Known Rogue Access Points | 
|  | /api/s/{site}/list/event | GET | List Events | 
|  | /api/s/{site}/list/alarm | GET | List Alarms | 
|  | /api/s/{site}/list/hotspotop | GET | List Hotspot Operators | 
|  | /api/s/{site}/list/hotspotpackage | GET | List Hotspot Packages | 
|  | /api/s/{site}/list/voucher | GET | List Vouchers | 
|  | /api/s/{site}/list/self | GET | Get Self Information | 
|  | /api/s/{site}/list/settings | GET | List Settings | 
|  | /api/s/{site}/get/setting/{section} | GET | Get Setting Section | 
|  | /api/s/{site}/list/health | GET | List Health Status | 
|  | /api/s/{site}/list/dashboard | GET | List Dashboard Data | 
|  | /api/s/{site}/list/portforward | GET | List Port Forwarding Rules | 
|  | /api/s/{site}/list/dynamicdns | GET | List Dynamic DNS Configurations | 
|  | /api/s/{site}/list/portprofile | GET | List Port Profiles | 
|  | /api/s/{site}/list/dpi_stats | GET | List DPI Statistics | 
|  | /api/s/{site}/list/dpi_stats_fields | GET | List DPI Statistics Fields | 
|  | /api/s/{site}/list/dpiapp | GET | List DPI Applications | 
|  | /api/s/{site}/list/dpigroup | GET | List DPI Groups | 
|  | /api/s/{site}/list/current_channels | GET | List Current Channels | 
|  | /api/s/{site}/list/country_codes | GET | List Country Codes | 
|  | /api/s/{site}/list/admins | GET | List Administrators | 
|  | /api/s/{site}/stat/admin | GET | Get Admin Statistics | 
|  | /api/s/{site}/cmd/evtmgr | POST | Event Manager Command | 
|  | /api/s/{site}/cmd/stamgr | POST | Station Manager Command | 
|  | /api/s/{site}/set/setting/{section} | POST | Update Setting Section | 
|  | /api/s/{site}/rest/user | POST | Create User | 
|  | /api/s/{site}/rest/usergroup | POST | Create User Group | 
|  | /api/s/{site}/rest/wlanconf | POST | Create WLAN Configuration | 
|  | /api/s/{site}/rest/networkconf | POST | Create Network Configuration | 
|  | /api/s/{site}/rest/firewallrule | POST | Create Firewall Rule | 
|  | /api/s/{site}/rest/firewallgroup | POST | Create Firewall Group | 
|  | /api/s/{site}/rest/portconf | POST | Create Port Configuration | 
|  | /api/s/{site}/rest/radiusprofile | POST | Create RADIUS Profile | 
|  | /api/s/{site}/rest/account | POST | Create Account | 
|  | /api/s/{site}/rest/wlangroup | POST | Create WLAN Group | 
|  | /api/s/{site}/rest/hotspotop | POST | Create Hotspot Operator | 
|  | /api/s/{site}/rest/hotspotpackage | POST | Create Hotspot Package | 
|  | /api/s/{site}/rest/voucher | POST | Create Voucher | 
|  | /api/s/{site}/rest/portforward | POST | Create Port Forwarding Rule | 
|  | /api/s/{site}/rest/dynamicdns | POST | Create Dynamic DNS Configuration | 
|  | /api/s/{site}/rest/portprofile | POST | Create Port Profile | 
|  | /api/s/{site}/rest/dpiapp | POST | Create DPI Application | 
|  | /api/s/{site}/rest/dpigroup | POST | Create DPI Group | 
|  | /api/s/{site}/upd/user/{id} | PUT | Update User | 
|  | /api/s/{site}/upd/usergroup/{id} | PUT | Update User Group | 
|  | /api/s/{site}/upd/wlanconf/{id} | PUT | Update WLAN Configuration | 
|  | /api/s/{site}/upd/networkconf/{id} | PUT | Update Network Configuration | 
|  | /api/s/{site}/upd/firewallrule/{id} | PUT | Update Firewall Rule | 
