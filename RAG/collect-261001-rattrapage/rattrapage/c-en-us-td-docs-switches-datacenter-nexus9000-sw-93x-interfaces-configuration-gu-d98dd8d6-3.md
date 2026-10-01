---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6-3
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6.md
source_anchor: ""
source_lines: [63, 79]
sha256: bf9ae0e7625ee23aecc41d3d5c2f32d026a9ded303d5e69b66151e1fe6f2a7a7
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6

The operational state of this interface is governed by the state of the various ports in its corresponding VLAN. An SVI interface on a VLAN comes up when at least one port in that VLAN is in the Spanning Tree Protocol (STP) forwarding state. Similarly, this interface goes down when the last STP forwarding port goes down or goes to another STP state.
High Availability
High availability features ensure continuous network operation and redundancy for Layer 2 interfaces on Cisco Nexus 9000 Series switches.
See the Cisco Nexus 9000 Series NX-OS High Availability and Redundancy Guide for complete information about high availability features.
Counter Values
See the following information on the configuration, packet size, incremented counter values, and traffic.
| Configuration | Packet Size | Incremented Counters | Traffic | 
|---|---|---|---|
| L2 port – without any MTU configuration | 6400 and 10000 | Jumbo, giant, and input error | Dropped | 
| L2 port – with jumbo MTU 9216 in network-qos configuration | 6400 | Jumbo | Forwarded | 
| L2 port – with jumbo MTU 9216 in network-qos configuration | 10000 | Jumbo, giant, and input error | Dropped | 
| Layer 3 port with default Layer 3 MTU and jumbo MTU 9216 in network-qos configuration | 6400 | Jumbo | Packets are punted to the CPU (subjected to CoPP configs), get fragmented, and then they are forwarded by the software. | 
| Layer 3 port with default Layer 3 MTU and jumbo MTU 9216 in network-qos configuration | 10000 | Jumbo, giant, and input error | Dropped | 
| Layer 3 port with jumbo Layer 3 MTU and jumbo MTU 9216 in network-qos configuration | 6400 | Jumbo | Forwarded without any fragmentation. | 
| Layer 3 port with jumbo Layer 3 MTU and jumbo MTU 9216 in network-qos configuration | 10000 | Jumbo, giant, and input error | Dropped | 
| Layer 3 port with jumbo Layer 3 MTU and default L2 MTU configuration | 6400 and 10000 | Jumbo, giant, and input error | Dropped | 
| Note |  |
