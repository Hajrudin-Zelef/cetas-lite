---
id: collect-260926-mikrotik/mikrotik/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-e-1
title: "MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-05-07"]
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-ecmp.md
source_anchor: ""
source_lines: [1, 146]
sha256: 2d5a7ecd24057a2b2b26e494fe7e244ab0169a0bfcbf9ebf9f305b5894bf019f
---

# MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing

Last updated: **2026-05-07**

This is a production-minded RouterOS v7 dual-WAN guide. It covers:

- two DHCP WAN uplinks;
- recursive gateway checks;
- per-connection-classifier, or PCC, load balancing;
- automatic failover inside each policy routing table;
- recursive main-table defaults for router-originated traffic, with optional ECMP;
- NAT masquerade across both WANs.

It is written as a guide, not just a pastebin. Review every interface name, subnet, probe IP, firewall rule, and route comment before using it on a live router.

Companion lab validation report: Vultr CHR lab setup, methodology, and results

Want the short pasteable version first? Start with TL;DR default-config paste, then come back here for the rationale and variants.


| Aspect | This design | 
| WAN type | Two DHCP uplinks (PPPoE/static covered in §9) | 
| Load balancing | Per-connection (PCC), connection-sticky | 
| Health check | Recursive route + `check-gateway=ping` (link state alone is not enough) | 
| Failover | Per policy table; backup default at `distance=2` | 
| Failback | Automatic when the probe recovers | 
| NAT | Masquerade on both WANs | 
| IPv6 | Out of scope | 
| Tested on | RouterOS 7.21.4 on CHR, Vultr lab — see companion report | 

**This design is for**: a single router with two ISPs where LAN clients should use both uplinks concurrently and survive a single-WAN failure without manual intervention.

**This design is *not***: bandwidth bonding, session-preserving failover, BGP multihoming, or a substitute for a real firewall policy.

**If you take one thing away from this guide**: a recursive default route through a probe IP that itself resolves through the DHCP-learned gateway, paired with `check-gateway=ping`, is what makes failover trigger on *upstream* outages instead of only on link drops. Everything else in this document is plumbing around that idea. See §5 for the rationale and §6.6 for the routes themselves.


These are the issues most likely to bite a reader applying this guide. Each links to the section that explains how the design handles it.

| Gotcha | What goes wrong | Where to read | 
| Cloud CHR or multi-port routers may not have `ether1` as WAN | Pasting `ether1` /`ether2` examples blindly can turn a management NIC into a "WAN" and lock you out | §1 cloud CHR warning, §2 | 
| FastTrack bypasses mangle and routing marks | Traffic marked for `to_ISP1` /`to_ISP2` exits the wrong WAN, or via`main` | §6.1 | 
| DHCP route-update script can leave placeholders disabled on first apply | First script run during the unbound state disables placeholders; bootstrap once after first apply | §6.7 | 
| Public management default route in `main` masks main-table dual-WAN | Router-originated services (DNS, NTP, package downloads, DDNS) follow the public path instead of the recursive defaults | §6.6 cloud-CHR note, §7.7 | 
| PCC is connection-sticky and changes public IP per-flow | Single-stream speed tests show one WAN; existing sessions break on failover; some banking / anti-fraud flows misbehave | §6.8, §11.1 | 
| Probe set should be disjoint from DNS resolvers | If your only probes are also your only resolvers, route-health failure and DNS-service failure become indistinguishable in triage | §6.6, §10 | 
| `check-gateway=ping` reaction window is ~30–40 s | Link-layer-style HA, not BFD; do not expect sub-second failover | §10 | 
| Local destinations must bypass PCC explicitly | Without the `local` address-list bypass, LAN-to-VLAN/VPN traffic gets PCC-marked and may fail | §6.4, §6.8, §8 | 
| `passthrough=no` on`mark-routing` is required | `passthrough=yes` lets a later mangle rule overwrite the routing mark and silently misroute | §6.8 | 
| Static or PPPoE WANs need a different bootstrap | The DHCP script does not apply; gateway must be set manually or to the PPPoE interface | §9.3, §9.4 | 


1. Assumptions
2. Read this before pasting
3. Topology
4. Packet-flow overview
5. Why recursive routing is used here
6. Core config
7. Validation checklist
8. Troubleshooting
9. Optional variants
10. Operational notes
11. Design trade-offs and when not to use this
12. References
13. Changelog


This example assumes:

| Item | Example value | Change this? | 
| RouterOS | v7.x, lab-tested on 7.21.4 | Yes, test on your target version | 
| WAN 1 interface | `ether1` | Usually yes | 
| WAN 2 interface | `ether2` | Usually yes | 
| LAN bridge | `bridge` | Usually yes | 
| LAN subnet | `192.168.88.0/24` | Usually yes | 
| Router LAN IP | `192.168.88.1/24` | Usually yes | 
| WAN type | DHCP on both WANs | Yes, see static/PPPoE notes | 
| IPv4 NAT | masquerade on both WANs | Usually yes | 
| IPv6 | not covered | Add separately | 

The config does **not** include a complete firewall policy. Keep or build a real input/forward firewall around it.

### Cloud CHR management NIC warning

Cloud-hosted RouterOS instances often have an extra public management NIC before the private WAN/LAN NICs. In the Vultr CHR lab used to validate this guide, `ether1` was the public management interface, while the actual dual-WAN test interfaces were `ether2` and `ether3`.

Do not assume `ether1` is WAN1 on a cloud VM. First map interfaces with DHCP/static probes, then substitute the real WAN and LAN interface names throughout this guide. If you must keep public SSH/WinBox management during testing, leave the public default route in the `main` table and apply the dual-WAN policy routing to LAN traffic only.

Also check the cloud provider's private-network MTU. Vultr VPC 2.0 used MTU `1450` in the validation lab, so the WAN-side CHR interfaces were set to `1450`.


## 2. Read this before pasting

Do not paste this into production blind.

Recommended process:

```
/export file=before-dual-wan-change
/system backup save name=before-dual-wan-change
```

Use **Safe Mode** in WinBox/terminal while applying the routing and firewall changes.

Rollback path if something goes wrong: load the backup from a local console with `/system backup load name=before-dual-wan-change`, or restore selected sections by `/import file-name=before-dual-wan-change.rsc`. Keep at least one console-reachable management path that does not depend on the new policy tables.

Important preflight checks:

1. Replace `ether1` ,`ether2` ,`bridge` , and`192.168.88.0/24` with your real values.
2. Confirm whether the router has a separate public management NIC that should stay out of the WAN list.
3. Disable or bypass FastTrack for traffic that will use routing marks.
4. Add all local LAN, VLAN, and VPN subnets to the `local` address list.
5. Pick probe IPs that are stable and reachable through the intended ISP.
6. Test failover by breaking upstream reachability, not only by unplugging a cable.


```
                  ┌─────────────── ISP 1, DHCP ───────────────┐
                  │                                             │
LAN clients ── bridge ── MikroTik RouterOS v7 ── ether1 / WAN1  │
                  │                           └─ ether2 / WAN2 ─┤
                  │                                             │
                  └─────────────── ISP 2, DHCP ───────────────┘
```


For LAN-originated traffic:

```
LAN client packet
    ↓
/prerouting mangle
    ↓
PCC marks the connection as ISP1_conn or ISP2_conn
    ↓
Routing mark is set to to_ISP1 or to_ISP2
    ↓
Route lookup happens in the selected FIB table
    ↓
Default route points at a recursive probe IP
    ↓
Probe IP resolves through the current DHCP gateway
    ↓
Packet exits WAN1 or WAN2
```

