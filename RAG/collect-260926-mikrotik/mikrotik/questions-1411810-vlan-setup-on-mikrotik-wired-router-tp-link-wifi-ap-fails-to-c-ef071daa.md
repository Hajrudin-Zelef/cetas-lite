---
id: collect-260926-mikrotik/mikrotik/questions-1411810-vlan-setup-on-mikrotik-wired-router-tp-link-wifi-ap-fails-to-c-ef071daa
title: "questions-1411810-vlan-setup-on-mikrotik-wired-router-tp-link-wifi-ap-fails-to-c-ef071daa"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1411810-vlan-setup-on-mikrotik-wired-router-tp-link-wifi-ap-fails-to-c-ef071daa.md
source_anchor: ""
source_lines: [1, 14]
sha256: 94e78bbf9e0fad00bbcaaf123ded2ebaba8cab95410f2fde4a683c18a139e311
---

# questions-1411810-vlan-setup-on-mikrotik-wired-router-tp-link-wifi-ap-fails-to-c-ef071daa

I have a MikroTik hEX router and TP-Link EAP225v3 wireless access point. I created VLAN with id 11 on the router, and created an SSID that uses vlan 11 on the AP. When i try to connect to the wifi, it fails to get an IP address, and says "no internet connection". My other SSID (without VLAN) works just fine.
For router setup i followed this manual:
/interface vlan
add name=VLAN11 vlan-id=11 interface=ether1 disabled=no
/ip address
add address=192.168.10.3/24 interface=VLAN11
/ip pool
add name=dhcp_pool11 ranges=192.168.10.5-192.168.10.254
/ip dhcp-server
add address-pool=dhcp_pool11 disabled=no interface=VLAN11 name=dhcp11
/ip dhcp-server network
add address=192.168.10.0/24 dns-server=8.8.8.8 gateway=192.168.10.3
On the EAP225 the config is fairly straight forward:
It seems I'm missing a step or two here. Any help appreciated!
