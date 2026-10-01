---
id: collect-261001-cisco/cisco/2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28-2
title: "2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28.md
source_anchor: ""
source_lines: [81, 88]
sha256: 80339a067e66c814f35df83360b3d52b739346fae9b016da6135f6840726106d
---

# 2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28

Being a guest WiFi network, the usual idea is to allow guest devices to the internet but not to any device on the local network. With firewall rules there are many ways to skin a cat, but I initially used the guidance in this OPNsense forum post, which prevents access to any private address ranges (that is, those defined in RFC1918 – 192.168.0.0/16, 172.16.0.0/12 & 10.0.0.0/8). This turned out to be over-strict, since it prevented traffic flowing even within the Guest WiFi subnet. I therefore listed my private LAN’s range specifically instead. Create a rules alias named that covers these ranges:
- 192.168.<your normal LAN>.0/16
- 172.16.0.0/12
- 10.0.0.0/8
Then, add the following two rules at the end of the rules list for the guest WiFi interface:
- block destination alias_defined_above
- allow destination all
These two rules in this order mean that any traffic that tries to reach any private address range via the firewall will be blocked, but any other traffic will flow (i.e. internet traffic). Traffic from one device to another within the guest WiFi subnet will flow, since this setup only blocks traffic that arrives at the firewall, which will only happen if the packet cannot find it’s destination within its source subnet.
