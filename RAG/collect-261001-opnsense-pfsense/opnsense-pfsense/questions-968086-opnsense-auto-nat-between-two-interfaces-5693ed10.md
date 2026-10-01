---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-968086-opnsense-auto-nat-between-two-interfaces-5693ed10
title: "questions-968086-opnsense-auto-nat-between-two-interfaces-5693ed10"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-968086-opnsense-auto-nat-between-two-interfaces-5693ed10.md
source_anchor: ""
source_lines: [1, 31]
sha256: 3d3360449888c3906fef0c17cf1aadf2a149b1fec08738d2329399b3a79f5e13
---

# questions-968086-opnsense-auto-nat-between-two-interfaces-5693ed10

Having trouble with ONPSense configuration.

**Explanation**

I've just configured a OPNSense Fw, which has 3 interfaces: WAN, LAN, ZRT

LAN: 192.168.101.1/23 (DHCP server from 192.168.101.50 to 192.168.101.150) Addresses on 192.168.100.0/24 mustn't go outside Addresses on 192.168.101.0/24 are abels to go outside

WAN: 192.168.1.x/24 (ip obtained via DHCP)

ZRT: logical int, zerotier VPN, 192.168.103.176/24

**Problem**

Forgetting about wan int, my problem is between the ZRT and LAN ints.

192.168.100.0/24 network is made up by PLCs, which haven't gw. Said that, i'm expecting, for any ICMP call from 192.168.103.0/24 host to any 192.168.100.0/24 host, a timeout.

But it works.

Things are working fine because LAN int is masquerating ZRT traffic, sending ICMP as 192.168.101.1 to lan network instead the original 192.168.103.x ip.

Then, why is doing auto s-nat?

Searching in NAT rules by i can't find anything on this, it seems OPNSense is only auto-generating WAN nat rules, which are deactivables, but nothing about my s-nat.

Am I missing anything?

Thanks for yoyr time.

**Hele.**
