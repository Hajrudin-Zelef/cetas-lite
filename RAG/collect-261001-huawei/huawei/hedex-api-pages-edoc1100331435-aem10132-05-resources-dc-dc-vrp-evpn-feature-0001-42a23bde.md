---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-evpn-feature-0001-42a23bde
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-evpn-feature-0001-42a23bde"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-evpn-feature-0001-42a23bde.md
source_anchor: ""
source_lines: [1, 5]
sha256: ed7c777ad868c0bf8edb4f54a7e71c43bd1f2ac4a6338210386eb56e238628b7
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-evpn-feature-0001-42a23bde

An Ethernet virtual private network (EVPN) is a VPN used for Layer 2 interworking. EVPN is similar to Border Gateway Protocol (BGP)/Multiprotocol Label Switching (MPLS) IP VPN. Using extended reachability information, EVPN implements MAC address learning and advertisement between Layer 2 networks at different sites on the control plane rather than on the data plane.
Figure 1 shows the basic EVPN model.
Originally, the Virtual Extensible LAN (VXLAN) solution did not provide the control plane. VXLAN used traffic flooding on the data plane to implement VXLAN tunnel endpoint (VTEP) discovery and host information learning, including IP and MAC addresses, VXLAN Network Identifiers (VNIs), and gateway VTEP IP addresses. This way resulted in high traffic volumes on data center networks. To resolve this problem, VXLAN uses EVPN as the control plane. EVPN allows VTEPs to exchange EVPN routes so that VTEPs can be automatically discovered and host information can be mutually advertised. Therefore, EVPN prevents unnecessary data traffic flooding.
EVPN is similar to BGP/MPLS IP VPN and communicates EVPN routes over a public network, improving security of customers' private data.
Additionally, manual configuration of the original VXLAN solution is time-consuming on a large-scale network. Using EVPN helps reduce the manual configuration workload.
