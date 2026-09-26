---
id: collect-260926-mikrotik/mikrotik/github-andre4freelance-mikrotik-vrrp-mikrotik-vrrp-virtual-router-redundancy-protocol-to-a
title: "github-andre4freelance-mikrotik-vrrp-mikrotik-vrrp-virtual-router-redundancy-protocol-to-achieve-dhc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/github-andre4freelance-mikrotik-vrrp-mikrotik-vrrp-virtual-router-redundancy-protocol-to-achieve-dhc.md
source_anchor: ""
source_lines: [1, 18]
sha256: 2e1a0e7e045f4bbd1bbb9da1091b3920fd7625a43006ede0dfa230b78d33a6f7
---

# github-andre4freelance-mikrotik-vrrp-mikrotik-vrrp-virtual-router-redundancy-protocol-to-achieve-dhc

This repository provides configuration examples for **MikroTik VRRP (Virtual Router Redundancy Protocol)** to achieve DHCP Server high availability.

The scenario simulates an office LAN network where **RO-MAIN** acts as the primary router with VRRP priority 254, and **RO-BACKUP** serves as the standby router with VRRP priority 10. Both routers share a virtual IP address (192.168.1.1) that clients use as their default gateway and DHCP source.

This solution is highly suitable for office environments that require **high network availability** and **minimal downtime**.

When both links are active, the **RO-MAIN** router handles all DHCP and gateway traffic for clients on VLAN 10.

When the distribution link on the primary router (RO-MAIN) is disconnected, the **RO-BACKUP** router takes over. Client devices experience only a **brief RTO (Request Timeout)**, demonstrating optimal network continuity.

- **Router** : MikroTik RouterOS 7.14.3
- **Switch** : Cisco NX-OS

Origin:

https://github.com/andre4freelance/MikroTik-vrrp

Facebook post: https://www.facebook.com/share/p/15SR1AjWQVx/
