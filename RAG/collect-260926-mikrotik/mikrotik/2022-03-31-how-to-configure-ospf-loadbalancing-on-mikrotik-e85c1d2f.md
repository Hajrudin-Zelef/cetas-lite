---
id: collect-260926-mikrotik/mikrotik/2022-03-31-how-to-configure-ospf-loadbalancing-on-mikrotik-e85c1d2f
title: "How To Configure OSPF LoadBalancing on Mikrotik"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/2022-03-31-how-to-configure-ospf-loadbalancing-on-mikrotik-e85c1d2f.md
source_anchor: ""
source_lines: [1, 81]
sha256: 898177324ea435387f2b2e3433dab1c2223dec97b952fcd67e4005c6f9e41bb0
---

# How To Configure OSPF LoadBalancing on Mikrotik

*Source : https://www.taufiknurhuda.web.id/2022/03/31/how-to-configure-ospf-loadbalancing-on-mikrotik/*

Use case: more than one network path running OSPF and you want load balancing. OSPF selects the best path by cost: in the topology, R1→R5 via R2–R4–R5 costs 20, via R2–R3–R4 costs 30, so the best path is R2–R4. Goal: load-balance so R1→R5 uses both R2–R4 and R2–R3–R4.

## Basic configuration per router

**R1**

```
/system identity set name=R1
/ip address add address=192.168.10.2/24 interface=ether1 network=192.168.10.0
/ip route add distance=1 gateway=192.168.10.1
```

**R5**

```
/system identity set name=R5
/ip address add address=192.168.20.2/24 interface=ether1 network=192.168.20.0
/ip route add distance=1 gateway=192.168.20.1
```

**R2**

```
/ip address
add address=192.168.10.1/24 interface=ether1 network=192.168.10.0
add address=23.23.23.1/24 interface=ether2 network=23.23.23.0
add address=24.24.24.1/24 interface=ether3 network=24.24.24.0

/routing ospf instance set [ find default=yes ] router-id=2.2.2.2

/routing ospf network
add area=backbone network=192.168.10.0/24
add area=backbone network=23.23.23.0/24
add area=backbone network=24.24.24.0/24
```

**R3**

```
/system identity set name=R3
/ip address
add address=23.23.23.2/24 interface=ether1 network=23.23.23.0
add address=34.34.34.1/24 interface=ether2 network=34.34.34.0

/routing ospf instance set [ find default=yes ] router-id=3.3.3.3

/routing ospf network
add area=backbone network=23.23.23.0/24
add area=backbone network=34.34.34.0/24
```

**R4**

```
/system identity set name=R4
/ip address
add address=34.34.34.2/24 interface=ether1 network=34.34.34.0
add address=24.24.24.2/24 interface=ether2 network=24.24.24.0
add address=192.168.20.1/24 interface=ether3 network=192.168.20.0

/routing ospf instance set [ find default=yes ] router-id=4.4.4.4

/routing ospf network
add area=backbone network=192.168.20.0/24
add area=backbone network=34.34.34.0/24
add area=backbone network=24.24.24.0/24
```

## Enabling load balancing

After basic config, R2's route to `192.168.20.0/24` uses gateway `24.24.24.2` (R2–R4 only). Increase the cost on R2's interface toward R4 (ether3) so both paths have equal cost:

```
[admin@R2] > routing ospf interface add interface=ether3 cost=20
```

Re-check routes on R2: the route to `192.168.20.0/24` now shows **2 gateways** — load balancing is working.
