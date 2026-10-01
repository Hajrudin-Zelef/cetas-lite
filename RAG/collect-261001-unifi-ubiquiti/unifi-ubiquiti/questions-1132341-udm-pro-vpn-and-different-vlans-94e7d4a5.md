---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1132341-udm-pro-vpn-and-different-vlans-94e7d4a5
title: "questions-1132341-udm-pro-vpn-and-different-vlans-94e7d4a5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1132341-udm-pro-vpn-and-different-vlans-94e7d4a5.md
source_anchor: ""
source_lines: [1, 8]
sha256: 060019e4b280b841cd4fbd075c09c80c3b5f6e265601eadfb28ae0ea0b1a854d
---

# questions-1132341-udm-pro-vpn-and-different-vlans-94e7d4a5

I'm running a UDM Pro with 3 USW24-POE which are connected via fibre (OS2) which is running fine meanwhile.
On the UDM there are a few VLANS which are all managed by the UDM Pro.
The Main Lan is 192.168.10.0/24 which is my "tech" Lan, then there is a "facility" lan 192.168.5.0/24 with Vlan-ID of 2 and a "gastronomy" lan 172.168.16.0/24 with Vlan-ID 60.
I'm trying to get a VPN connection directly over the Windows VPN. The VPN Server is configurated correctly and authenticates via RADIUS with a subnet of 192.168.70.0/24.
The main problem is, that if I am connected with VPN I can reach every client which is in the 192.168.10.0/24 range but no others. I'd like to reach the clients in the 192.168.5.0/24 and the only thing I can ping is the DHCP Server (192.168.5.1).
The problem exists only when using L2TP, when using wireguard everything works fine (the bad thing is that my client wants to use the built-in Windows VPN).
I also tried to configure the radius server to "Assigned Vlan Support" Wired and Wireless enabled with the Vlan-ID 2 set on the user which doesn't do the trick. And I also tried to route the IP Range from 192.168.70.1 - 192.168.70.255 to the 192.168.5.0/24 network...
Could someone please tell me which configuration part I'm missing?
