---
id: collect-260926-mikrotik/mikrotik/app-104q2a7-controladoria-fabriciano-mg-gov-br-ospf-configuration-on-mikrotik-a-b2341954
title: "OSPF Configuration on MikroTik: A Complete Guide"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/app-104q2a7-controladoria-fabriciano-mg-gov-br-ospf-configuration-on-mikrotik-a--b2341954.md
source_anchor: ""
source_lines: [1, 115]
sha256: 8c6401e7bf643490541c8074a10f3d9f0050edacc12c56e9744e7d01ef10a8ae
---

# OSPF Configuration on MikroTik: A Complete Guide

*Source : https://controladoria.fabriciano.mg.gov.br/app/104q2a7/controladoria.fabriciano.mg.gov.br/ospf-configuration-on-mikrotik-a-complete-guide-1764803366*

## Understanding OSPF Fundamentals

**OSPF (Open Shortest Path First)** is a link-state routing protocol: routers build a complete map of network topology, flooding link-state info to all routers in the area so each holds an identical **Link-State Database (LSDB)**, then run the **Dijkstra/SPF algorithm** to compute shortest paths. Fast convergence on changes.

**Areas** subdivide the network for scalability — routers only need detail of their own area plus inter-area links. Typically use **Area 0 (Backbone)**; routers connecting areas are **ABRs**. Single-area (Area 0) is fine for simple networks.

## Prerequisites

- WinBox or SSH access to the MikroTik router; comfort with basic CLI if using terminal.
- Basic IP addressing/subnetting knowledge.
- Clear network topology; OSPF interfaces need IP reachability (same subnet or static route).
- Firewall must not block OSPF: IP protocol **89**, multicast `224.0.0.5` (AllSPFRouters) and `224.0.0.6` (AllDRouters).
- Reasonably up-to-date RouterOS.

## Step-by-Step OSPF Configuration in MikroTik

### Enabling OSPF

WinBox: `Routing` -> `OSPF` -> `+` to add an OSPF instance (usually one).

CLI:

```
/routing ospf instance
add name=default
```

### Configuring Interfaces

WinBox: `Routing` -> `OSPF` -> `Interfaces` tab -> `+`: select interface (e.g., `ether2`), Area (`backbone`/area 0), optional Cost override (lower = preferred), **Passive** if no OSPF neighbor expected on that interface (still advertises the network, sends no OSPF packets).

CLI (ether2 active, ether3 passive LAN):

```
/routing ospf interface
add interface=ether2 area=backbone
add interface=ether3 area=backbone passive=yes
```

### Defining Networks to Advertise (Optional but Recommended)

WinBox: `Routing` -> `OSPF` -> `Networks` tab -> `+`: Network (e.g., `192.168.1.0/24`), Area.

CLI:

```
/routing ospf network
add network=192.168.1.0/24 area=backbone
```

Note: often unnecessary if OSPF interfaces are correctly configured — the router advertises their connected networks automatically.

### Checking OSPF Neighbors

WinBox: `Routing` -> `OSPF` -> `Neighbors` tab — state should be `Full`.

CLI:

```
/routing ospf neighbor print
```

## Advanced OSPF Settings and Considerations

### Cost Tuning

Default cost derives from interface bandwidth; override manually to prefer a link regardless of speed (`Routing` -> `OSPF` -> `Interfaces` -> `Cost`). Useful for load balancing/prioritization.

### Priority and DR/BDR Election

On multi-access segments, OSPF elects a **Designated Router (DR)** and **Backup DR** to flood LSAs, reducing adjacencies. Election uses interface **Priority** (higher wins; default 1; `0` = never DR/BDR). Set at `Routing` -> `OSPF` -> `Interfaces` -> `Priority`.

### Authentication

Plain-text or MD5 authentication (`Routing` -> `OSPF` -> `Areas`/`Interfaces` -> `Authentication`). Same type and key on all routers in the area — prevents rogue routers injecting false routes.

### Route Summarization

Consolidate advertisements into summary routes on ABRs/ASBRs to shrink LSDB and routing tables (`Routing` -> `OSPF` -> `Areas` -> `Redistribute`, or `Networks` with larger masks).

### Timers

`Hello` and `Dead` intervals must match between neighbors (`Routing` -> `OSPF` -> `Interfaces` -> `Timers`). Mismatches prevent adjacencies.

## Troubleshooting Common OSPF Issues

### Neighbor Adjacencies Not Forming

1. **IP connectivity:** can routers ping each other's OSPF interface IPs?
2. **Subnet mismatch:** OSPF interfaces must be in the same subnet.
3. **OSPF process enabled** on both routers?
4. **Interface added to OSPF** on both, correct area?
5. **Area mismatch** on the link (e.g., Area 0 vs Area 1)?
6. **Hello/Dead timers** identical on both sides?
7. **Firewall** blocking protocol 89 / 224.0.0.5 / 224.0.0.6?
8. **MTU mismatch** (or use `mtu-ignore` cautiously).

### Incorrect Routes Being Advertised/Learned

1. **Passive interfaces:** a passive interface's network isn't advertised unless explicitly added under `/routing ospf network`.
2. **Networks statement:** correct network/mask under `/routing ospf network`?
3. **Redistribution** from static/other protocols configured correctly (`Routing` -> `OSPF` -> `Redistribute`)?
4. **Route filtering/prefix lists** blocking routes?

### High CPU Usage

Causes: too many neighbors/LSAs, frequent flapping, misconfigurations causing loops or LSA storms. Check `/routing ospf lsa print` and neighbor status; ensure hardware fits network size.

## Conclusion

Enable the OSPF instance, configure OSPF interfaces (areas, passive settings), verify neighbor adjacencies (`Full`), then fine-tune with cost, DR/BDR priority, authentication, and summarization. OSPF brings automation, resilience, and efficiency to the network.
