---
id: collect-260926-mikrotik/mikrotik/questions-1864149-dhcp-not-assigning-ip-addresses-for-wifi-on-vlan-40-mikrotik-a6ad66b4
title: "questions-1864149-dhcp-not-assigning-ip-addresses-for-wifi-on-vlan-40-mikrotik-a6ad66b4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1864149-dhcp-not-assigning-ip-addresses-for-wifi-on-vlan-40-mikrotik-a6ad66b4.md
source_anchor: ""
source_lines: [1, 25]
sha256: 411d64f905b303eb0d87164561f09514671854a805cff58ff6af7b0a7eb2d5da
---

# questions-1864149-dhcp-not-assigning-ip-addresses-for-wifi-on-vlan-40-mikrotik-a6ad66b4

I’m trying to configure a VLAN (VLAN 40) on my MikroTik router that includes both Ethernet ports (ether3, ether4) and a WiFi network. However, I’m encountering an issue where devices connecting to the WiFi network do not receive an IP address via DHCP.
VLAN 40 with ether3 and ether4 - Works Fine, IP addresses are correctly assigned to devices connected to ether3 and ether4.
/interface bridge add name=bridge_vlan40
/interface/bridge/port/set bridge=bridge_vlan40 [find interface=ether3]
/interface/bridge/port/set bridge=bridge_vlan40 [find interface=ether4]
/interface vlan add name=vlan40 vlan-id=40 interface=bridge_vlan40
/ip address add address=192.168.40.1/24 interface=bridge_vlan40
/ip pool add name=vlan40_pool ranges=192.168.40.10-192.168.40.254
/ip dhcp-server add name=vlan40_dhcp address-pool=vlan40_pool interface=bridge_vlan40
/ip dhcp-server network add address=192.168.40.0/24 dns-server=8.8.8.8,1.1.1.1 gateway=192.168.40.1
/ip firewall nat add chain=srcnat action=masquerade out-interface=ether1 src-address=192.168.40.0/24
I created a WiFi network and added it to VLAN 40 as follows:
/interface/wifi add name=wifi_vlan40 mtu=1500 l2mtu=1560 master-interface=wifi1 security.authentication-types=wpa2-psk,wpa3-psk security.passphrase="XXXX" datapath.vlan-id=40 configuration.ssid="XXXX" disabled=no
/interface/bridge/port/add bridge=bridge_vlan40 interface=wifi_vlan40
When connecting to the WiFi network (wifi_vlan40), devices do not receive an IP address from the VLAN 40 DHCP server (vlan40_dhcp).
For reference, I have a working configuration for a different VLAN (VLAN 10) that is only for a WiFi network and assigns IP addresses correctly:
/interface/wifi add name=wifi_vlan10 mtu=1500 l2mtu=1560 master-interface=wifi1 security.authentication-types=wpa2-psk,wpa3-psk security.passphrase="XXX" datapath.vlan-id=10 configuration.ssid="XXX" configuration.mode=ap disabled=no
/interface vlan add name=vlan10 vlan-id=10 interface=wifi_vlan10
/ip address add address=192.168.10.1/24 interface=vlan10
/ip pool add name=vlan10 ranges=192.168.10.10-192.168.10.254
/ip dhcp-server add address-pool=vlan10 interface=vlan10 name=vlan10
/ip dhcp-server network add address=192.168.10.0/24 dns-server=8.8.8.8,1.1.1.1 gateway=192.168.10.1
/ip firewall nat add chain=srcnat action=masquerade out-interface=ether1 src-address=192.168.10.0/24
I suspect I might be missing a configuration step, but I can’t pinpoint the issue. Could anyone help me understand why DHCP is not working for WiFi on VLAN 40?
Thanks in advance!
