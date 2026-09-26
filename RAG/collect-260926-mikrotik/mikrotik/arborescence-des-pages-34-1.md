---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-34-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-34.md
source_anchor: ""
source_lines: [1, 14]
sha256: 4d2467ae38a549b960dc7ccf4d94759a957c3a45e79a63c1ad32c91ad0b0f1aa
---

# Introduction

Virtual eXtensible Local Area Network (VXLAN) is a tunneling protocol designed to solve the problem of limited VLAN IDs (4096) in IEEE 802.1Q, and it is described by IETF RFC 7348. With VXLAN the size of the identifier is expanded to 24 bits (16777216). It creates a Layer 2 overlay scheme on a Layer 3 network and the protocol runs over UDP. RouterOS VXLAN interface supports IPv4 or IPv6 (since version 7.6), but dual-stack is not supported.

VXLAN creates a 50-byte overhead for IPv4 and a 70-byte overhead for IPv6. When configuring VXLAN, it is recommended to ensure that the size of the encapsulated Ethernet frame does not exceed the MTU of the underlying network, by configuring the MTU accordingly or by limiting the size of the Ethernet frames.

Only devices within the same VXLAN segment can communicate with each other. Each VXLAN segment is identified through a 24-bit segment ID, termed the VXLAN Network Identifier (VNI). Unlike most tunnels, a VXLAN is a 1-to-N network, not just point-to-point. VXLAN endpoints, which terminate VXLAN tunnels are known as VXLAN tunnel endpoints (VTEPs). RouterOS only supports statically configured remote VTEPs. When unicast traffic needs to be sent over VXLAN, a device can learn the IP address of the other endpoint dynamically in a manner similar to a learning bridge, and forward traffic only to the necessary VTEP. For traffic that needs to be flooded (broadcast, unknown-unicast, and multicast) to all VTEPs on the same segment, VXLAN can use multicast or unicast with head-end replication to send one replica for every remote VTEP.

# Configuration options

This section describes the VXLAN interface and VTEP configuration options.

**Sub-menu:** `/interface vxlan`

