---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1149565-openwrt-in-ubnt-unifi-lr-unstable-ping-31c2115d
title: "questions-1149565-openwrt-in-ubnt-unifi-lr-unstable-ping-31c2115d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1149565-openwrt-in-ubnt-unifi-lr-unstable-ping-31c2115d.md
source_anchor: ""
source_lines: [1, 15]
sha256: a77b00c9995dfb7526089643e92cd3005f44b61ed3b23f0eb9f3743fbb201829
---

# questions-1149565-openwrt-in-ubnt-unifi-lr-unstable-ping-31c2115d

I installed openwrt in UBNT Unifi LR. Here's the config:
config interface 'VLAN10'
option type 'bridge'
option ifname 'eth0.10'
config interface 'VLAN8'
option ifname 'eth0.8'
option proto 'static'
option ipaddr '192.168.101.24'
option netmask '255.255.255.0'
option gateway '192.168.101.1'
option dns '192.168.101.1'
wireless is bridged to VLAN10, VLAN8 for management. I already disable the firewall and the dnsmasq from startup.
I can connect to the AP and I don't experience any lagging in browsing or download when using the AP, but if I ping from my router (which has vlan8 interface), the ping is not stable, it said: "Destination Host Unreachable". When I can ping, and login to the AP, I checked that the AP is not busy, low bandwidth and low CPU usage.
I didn't experience this problem with other hardware (openwrt/ddwrt based), so I think this is not network issue.
What is happening? How can I fix this?
