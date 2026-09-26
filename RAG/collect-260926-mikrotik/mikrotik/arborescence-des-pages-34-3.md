---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-34-3
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-34.md
source_anchor: ""
source_lines: [74, 106]
sha256: 37a2a3d6d12bf16b1b734db916e7a290c1d988144821494fe5e1219831ded8b6
---

# Introduction

Starting from RouterOS version 7.18, initial support for hardware-offloaded VXLAN was introduced. This makes offloaded VXLAN data plane possible, supporting encapsulation and decapsulation, and allowing for static one-to-one VLAN-to-VXLAN mappings within a vlan-filtering bridge. Refer to the L3HW Device Support documentation for a list of compatible devices.

At this point, some known features are not yet implemented.

**Underlay (routing encapsulated VXLAN packets):**

1. VTEPs are not supported over ECMP,

2. VTEPs are not supported over bond, bridge, VLAN interfaces (only stand-alone routed Ethernet interfaces are supported),

3. VTEPs are not supported over multicast,

4. VTEPs cannot operate within VRFs,

5. VTEPs are not supported with IPv6.

**Overlay (forwarding between Ethernet and VXLAN):**

1. VLAN tagging over VXLAN is not supported,

2. Routing between different VXLAN VNIs is not supported,

3. VTEPs are isolated, and there is no mechanism to control "horizon" between them.

4. Bridged VXLAN interfaces do not support IGMP snooping. When snooping is enabled, MDB entries on VXLAN are not offloaded, and multicast traffic gets restricted between Ethernet and VXLAN.

5. Bridged VXLAN interfaces are not supported by MLAG.

## Basic configuration example

In this example, static routing is used to reach remote VTEPs, but dynamic routing protocols like OSPF or BGP could also be used. The upstream interface has a higher MTU to support large packets and VXLAN encapsulation. Below is a network topology overview:

**sfp-sfpplus1** - upstream (underlay) interface**sfp-sfpplus3** - bridged port for untagged VLAN 10**sfp-sfpplus4** - bridged port for untagged VLAN 20**vxlan-10010** - overlay port for untagged VLAN 10**vxlan-10020** - overlay port for untagged VLAN 20
