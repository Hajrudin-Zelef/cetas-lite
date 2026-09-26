---
id: collect-260926-mikrotik/mikrotik/sharing-routes-between-mikrotik-and-vyos-using-ospf
title: "sharing-routes-between-mikrotik-and-vyos-using-ospf"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/sharing-routes-between-mikrotik-and-vyos-using-ospf.md
source_anchor: ""
source_lines: [1, 52]
sha256: 3d375c06da94a28f414a93411e9febd30c87d6132ce10cf2aa42aa33364e10a9
---

# sharing-routes-between-mikrotik-and-vyos-using-ospf

Planning out my new network topology; I wanted to use multiple routers — and have them share routes between them. So I started looking into OSPF.

I made a simple lab in EVE-NG, and set up route sharing between MikroTik’s RouterOS and VyOS.

## Table of contents

The MikroTik router is connected to two virtual PC’s on `10.10.10.0/24`. VyOS is also connected to two VPC’s; one on `192.168.10.0/24`, and the other on `192.168.20.0/24`.

The two routers are connected together on `10.12.13.0/24` — MikroTik being `10.12.13.1`, and VyOS; `10.12.13.2`.

Add OSPF to networks `10.10.10.0/24` and `10.12.13.0/24`. Make the router-id `10.12.13.1`:

```
routing ospf network add network=10.10.10.0/24 area=backbone
routing ospf network add network=10.12.13.0/24 area=backbone
routing ospf instance set default router-id=10.12.13.1
```
Enter configuration mode, and add OSPF to networks `10.12.13.0/24`, `192.168.10.0/24`, and `192.168.20.0/24`. Make router-id `10.12.13.2`, commit and save:

```
config
set protocols ospf area 0.0.0.0 network 10.12.13.0/24
set protocols ospf area 0.0.0.0 network 192.168.10.0/24
set protocols ospf area 0.0.0.0 network 192.168.20.0/24
set protocols ospf parameters router-id 10.12.13.2
commit && save
```
More information about VyOS OSPF in the documentation.

After the two routers shared their routes — this is what the OSPF routing tables looked like:

```
[admin@MikroTik] > routing ospf route print
 # DST-ADDRESS        STATE          COST     GATEWAY         INTERFACE
 0 10.10.10.0/24      intra-area     10       0.0.0.0         bridge1
 1 10.12.13.0/24      intra-area     10       0.0.0.0         ether2
 2 192.168.10.0/24    intra-area     11       10.12.13.2      ether2
 3 192.168.20.0/24    intra-area     11       10.12.13.2      ether2
```
```
vyos@vyos:~$ show ip ospf route
============ OSPF network routing table ============
N    10.10.10.0/24         [11] area: 0.0.0.0
                           via 10.12.13.1, eth0
N    10.12.13.0/24         [1] area: 0.0.0.0
                           directly attached to eth0
N    192.168.10.0/24       [1] area: 0.0.0.0
                           directly attached to eth1
N    192.168.20.0/24       [1] area: 0.0.0.0
                           directly attached to eth2
```
I’m not sure I’ll end up using multiple routers, time will tell. But it was a good learning experience to play with OSPF in EVE-NG — and now I have it documented 🙂
