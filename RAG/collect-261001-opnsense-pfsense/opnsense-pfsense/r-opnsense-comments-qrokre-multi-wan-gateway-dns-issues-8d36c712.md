---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-qrokre-multi-wan-gateway-dns-issues-8d36c712
title: "r-opnsense-comments-qrokre-multi-wan-gateway-dns-issues-8d36c712"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-qrokre-multi-wan-gateway-dns-issues-8d36c712.md
source_anchor: ""
source_lines: [1, 31]
sha256: b5d834841ac7aa7fab6a92a4c06101fc133719a69fb8353873e7365ab590aea7
---

# r-opnsense-comments-qrokre-multi-wan-gateway-dns-issues-8d36c712

I'm struggling to get opnsense working when there are two gateways:

Base Setup: APU4

WAN -> OPT1

LAN -> OPT2 IP/Subnet 172.31.255.1/24 DHCP Range 172.31.255.11-15

WIFI -> Part of the LAN subnet - Atheros based card, mostly seems to work but does throw some errors occasionally. IP/Subnet 172.31.255.5/24 DHCP Range 172.31.255.16-20

WWAN -> Sierra Wireless AX7700 with Project Fi sim. Definitely connects and works.

Unbound is setup and active using Quad9 for upstream.

DHCP is setup to handout the firewall address for DNS (internal 172.31.255.1)

If I just use a WAN plug without setting up the WWAN everything will work fine.

If I setup the WWAN and swap the WAN Interface to the WWAN one everything will work fine.

If I setup the WWAN and the WAN as gateways at the same time DNS is gone. From both LAN and WIFI. Including when I setup a group and setup WWAN as a failover.

I've had similar issues when setting up a dual WAN gateway with failover using just the OPT1&2 ports.

I've used pf for years, but OpnSense has the hardware support for the WWAN.

I'm always struggling with DNS on OPNSense, What am I missing?

no comments yet

Be the first to share what you think!
