---
id: collect-261001-ia-llm/ia-llm/alanlora-core-mpls-ip-5e46e156
title: "alanlora-core-mpls-ip-5e46e156"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-ia-llm/alanlora-core-mpls-ip-5e46e156.md
source_anchor: ""
source_lines: [1, 22]
sha256: 8b50262c1a72c5a71363d100b6fb057c7afe37d380ece8ad3bc3529a65f05da3
---

# alanlora-core-mpls-ip-5e46e156

This project simulates a service provider MPLS backbone built using Huawei eNSP with Huawei VRP (AR series routers).
The topology represents an ISP-style network delivering L3VPN and L2VPN services to multiple customers over an MPLS core, including traffic engineering and customer segmentation mechanisms.
- MPLS core running IS-IS as the IGP
- MP-BGP vpnv4 for L3VPN route distribution
- VRF implementation for customer segmentation and traffic isolation
- MPLS Traffic Engineering (MPLS-TE)
- Bandwidth reservation on TE tunnels
- Explicit path tunnel configuration
- Dynamic L2VPN services for two customers
- VLAN extension between remote sites using 802.1Q subinterfaces
- Customer gateways located at the remote edge of the network
- Access mode ports for end devices
- 802.1Q trunk links for multi-VLAN transport
- Ether-Trunk with two Gigabit interfaces for redundancy and load balancing
- OSPF configured in selected segments and integrated with the MPLS core
- Loopback interfaces used for router ID and routing protocol sessions
- MPLS backbone design
- Multi-tenant segmentation using VRFs
- L3VPN and L2VPN service provisioning
- Traffic engineering and explicit path control
- Interaction between IS-IS, OSPF, and MP-BGP
- ISP-style troubleshooting and optimization
