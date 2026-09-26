---
id: collect-260926-mikrotik/mikrotik/port-based-vlan-tagging-trunk-and-access-ports-mikrotik-wifinederland-nl
title: "port-based-vlan-tagging-trunk-and-access-ports-mikrotik-wifinederland-nl"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/port-based-vlan-tagging-trunk-and-access-ports-mikrotik-wifinederland-nl.md
source_anchor: ""
source_lines: [1, 23]
sha256: 2271bdf02bff7d2d825c00b47f96307c6453ed41799f59699a3628b23b059799
---

# port-based-vlan-tagging-trunk-and-access-ports-mikrotik-wifinederland-nl

- Add necessary VLAN interfaces on ethernet interface to make it as a VLAN trunk port.

/interface vlan
add interface=ether2 name=eth2-vlan200 vlan-id=200
add interface=ether2 name=eth2-vlan300 vlan-id=300
add interface=ether2 name=eth2-vlan400 vlan-id=400

- Add bridges for each VLAN

/interface bridge
add name=bridge-vlan200
add name=bridge-vlan300
add name=bridge-vlan400

- Add VLAN interfaces to their corresponding bridges and ethernet interfaces where untagged traffic is necessary

/interface bridge port
add bridge=bridge-vlan200 interface=eth2-vlan200
add bridge=bridge-vlan200 interface=ether6
add bridge=bridge-vlan300 interface=eth2-vlan300
add bridge=bridge-vlan300 interface=ether7
add bridge=bridge-vlan400 interface=eth2-vlan400
add bridge=bridge-vlan400 interface=ether8
