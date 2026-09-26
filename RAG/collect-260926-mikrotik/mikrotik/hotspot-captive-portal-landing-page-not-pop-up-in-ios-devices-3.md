---
id: collect-260926-mikrotik/mikrotik/hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices-3
title: "hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices.md
source_anchor: ""
source_lines: [122, 153]
sha256: eea8dc2f8be909ea46264ec84c01b5ff107ea9d17925cb309fc7b55c2452fbb8
---

# hotspot-captive-portal-landing-page-not-pop-up-in-ios-devices

set [ find default=yes ] limit-bytes-total=65000000
/ip hotspot walled-garden
add dst-host=**ELIDED** server=hotspot1
/ip route
add disabled=no distance=1 dst-address=0.0.0.0/0 gateway=102.216.174.1 pref-src="" routing-table=main scope=30 suppress-hw-offload=no target-scope=10 vrf-interface=vlan126_B2B_Management
add disabled=no distance=1 dst-address=10.100.0.0/16 gateway=10.100.102.1 pref-src="" routing-table=main scope=30 suppress-hw-offload=no target-scope=10 vrf-interface=vlan126_B2B_Management
/ip service
set telnet disabled=yes
set ftp disabled=yes
set www disabled=yes
set ssh address=162.213.64.0/24,162.213.65.0/24,10.100.102.1/32,162.213.66.0/24,162.213.67.0/24,185.132.90.0/24,185.41.96.0/24,185.41.97.0/24,185.41.98.0/24,185.41.99.0/24,45.15.136.0/24,145.224.66.0/24,102.216.172.1/32,41.223.183.25/32 port=32222
set www-ssl address=192.168.150.0/23,192.168.160.0/24 disabled=no
set api disabled=yes
set winbox address="162.213.64.0/24,162.213.65.0/24,162.213.66.0/24,162.213.67.0/24,185.132.90.0/24,185.41.96.0/24,185.41.97.0/24,185.41.98.0/24,185.41.99.0/24,45.15.136.0/24,145.224.66.0/24,102.216.172.1/32,41.223.183.25/32,10.100.10.1/32,192.168.200.0/24,10.100\
    .10.15/32,10.100.102.1/32" port=18291
set api-ssl disabled=yes
/radius
add address=10.100.2.200 service=hotspot
/radius incoming
set accept=yes
/system clock
set time-zone-name=**ELIDED**
/system identity
set name=B2B-CP-NAME-FW
/system note
set show-at-login=no
/system routerboard settings
set enter-setup-on=delete-key
/tool mac-server
set allowed-interface-list=LAN
/tool mac-server mac-winbox
```
