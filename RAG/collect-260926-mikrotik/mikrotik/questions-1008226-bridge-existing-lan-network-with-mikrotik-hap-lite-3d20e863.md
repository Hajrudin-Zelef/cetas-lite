---
id: collect-260926-mikrotik/mikrotik/questions-1008226-bridge-existing-lan-network-with-mikrotik-hap-lite-3d20e863
title: "questions-1008226-bridge-existing-lan-network-with-mikrotik-hap-lite-3d20e863"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1008226-bridge-existing-lan-network-with-mikrotik-hap-lite-3d20e863.md
source_anchor: ""
source_lines: [1, 15]
sha256: fe42047a26359b3bdc7cefafe35c44c3a6298c97df4d1687128569af27715fd1
---

# questions-1008226-bridge-existing-lan-network-with-mikrotik-hap-lite-3d20e863

I have existing network on 192.168.0.1 subnet, now I bought Mikrotik hap Lite which I want to use to extend existing network, and be able to connect clients on Mikrotik both wired (printer) and wirelles ( smartphone, notebook ) but all devices must see each other. First modem/router will be used as DHCP server and its address is 192.168.0.1. How should I configure Mikrotik to achieve mentioned behavior?
1 Answer 1
Just setup hap Lite to use a subnet of your 192.168.0.0/255.255.255.0(?) primary network and connect hap Lite WAN to your primary router LAN.
If you prefer DHCP usage:
- Setup hap Lite WAN port to get IP address from DHCP.
- Bind a hap Lite WAN IP address with its MAC address on yourprimary router DHCP configuration.
- Choose hap Lite subnet from outside of DHCP range.
- Add route to your primary router to push network traffic tohap Lite LAN subnet troughhap Lite WAN IP.
If you prefer static addressing:
- Setup hap Lite WAN port to unique IP address in yourprimary router LAN.
- Choose hap Lite subnet as unique IP range in yourprimary router LAN.
- Add route to your primary router to push network traffic tohap Lite LAN subnet troughhap Lite WAN IP.
You can mix:
- Static hap Lite WAN port IP address.
- DHCP on hap Lite LAN.
