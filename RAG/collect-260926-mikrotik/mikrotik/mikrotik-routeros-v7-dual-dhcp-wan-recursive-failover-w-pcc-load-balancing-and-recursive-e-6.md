---
id: collect-260926-mikrotik/mikrotik/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-e-6
title: "MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-05-07"]
keywords: ["latency"]
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-ecmp.md
source_anchor: ""
source_lines: [673, 714]
sha256: 4eba25217637cabcff8f22f703aa6ae07f53049b16d47b10c17bb8b455b0e478
---

# MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing

This guide only covers routing, mangle, and NAT. It does not replace a firewall policy. Keep normal protections such as:

- input chain drop for unsolicited WAN traffic;
- established/related accept rules;
- invalid drop rules where appropriate;
- explicit management access restrictions;
- dst-nat rules only for services you intend to expose.

Be cautious with generic invalid drops if your network has asymmetric routing, tunnels, or advanced policy routing.


## 11. Design trade-offs and when not to use this

### 11.1 What you give up by choosing PCC

PCC is connection-level, not packet-level. That means:

- A single TCP flow uses one WAN. Single-stream speed tests measure one link, not the sum of both.
- Sessions that bind to a public IP — banking, anti-fraud / Cloudflare Turnstile, some video conferencing, IMAP IDLE on certain providers — can misbehave when consecutive connections from the same client leave through different ISPs. If most users hit a small set of such services, route those destinations through one WAN via address-list + routing mark, alongside or instead of PCC.
- Failover is not session-preserving. Live SSH, VoIP, and streaming sessions on the failed WAN drop. New connections recover automatically.

### 11.2 When a simpler design is enough

If you only need **failover** (no load balancing), drop the PCC mangle rules and keep the recursive default routes in `main` at distances 1 and 2. One WAN is active at a time, and `check-gateway=ping` triggers failover the same way.

If you only need **per-client steering** (e.g. send the guest VLAN out WAN2), skip PCC and use `src-address-list` → `mark-routing` directly. Static assignment is easier to debug than connection-classifier hashing.

### 11.3 When PCC is the wrong tool

- Two ISPs with very different latency or jitter where you want application-aware steering. Use a real SD-WAN appliance.
- Public-facing services that must answer on a stable IP. PCC + masquerade gives a different public IP per WAN; you want BGP multihoming or an upstream front-door (DNS failover, anycast, cloud LB).
- True bandwidth aggregation. Look at MLPPP, L2TP+BCP, or commercial bonding services. None are RouterOS-native at the level you might expect.

### 11.4 Why recursive routes instead of `netwatch` scripts

Some guides drive failover with `/tool netwatch` scripts that enable and disable default routes when a probe goes down. That works, but it is an out-of-band controller racing with the routing table. Recursive routing keeps the health check inside the FIB: a route is active if and only if its probe is reachable through the intended next-hop. There is one source of truth, no script timing to tune, and `check-gateway=ping` is documented behavior rather than a custom script you have to maintain across upgrades.



- **2026-05-07** — Initial public version. Includes the TL;DR default-config paste, route probes deliberately disjoint from DNS resolvers, and validation against RouterOS 7.21.4 in a Vultr CHR lab. PCC split, per-policy-table failover and failback, recursive`main` -table ECMP for router-originated traffic, and the TL;DR paste applied to a default-like RouterBOARD baseline are all confirmed. See companion lab report.

@bwraith I removed the ICMP accept rule and updated the output to check for destination. It seems to work.
