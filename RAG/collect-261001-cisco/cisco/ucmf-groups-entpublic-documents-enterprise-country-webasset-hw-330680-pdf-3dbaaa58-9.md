---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-9
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [1016, 1183]
sha256: fb67988bfa934721b9ffd4144095a8525f5dd3fb7bfb0ffe15492c71b071eda1
---

# A line starting with the # sign is comments.

An MPLS-based VPN has the following characteristics:
l PEs manage VPN users, set up LSPs between PEs, and allocate routes to sites on a VPN.
l The route allocation between PEs is implemented by LDP or MP-BGP.
l The MPLS-based VPN supports IP address multiplexing between sites as well as the
interconnection of different VPNs.
1.3.2 MPLS-based TE
On traditional IP networks, routers select the shortest path as the route regardless of other factors
such as bandwidth. Traffic on a path is not switched to other paths even if the path is congested.
As more applications are deployed on the Internet, this shortest path first rule causes severe
problems on networks.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
17

Traffic engineering (TE) adjusts parameters including traffic management, routing, and resource
restraint parameters in real time to dynamically monitor the network traffic and the load of the
network components, which prevents network congestion caused by unbalanced traffic
distribution.
Figure 1-14 MPLS-based TE
R1
R3
R7 R6
R5 R4
R2
80 Mbit/s 
bandwidth
30 Mbit/s 
bandwidth
30 Mbit/s 
bandwidth
80 Mbit/s 
bandwidth80 Mbit/s 
bandwidth
30 Mbit/s traffic
50 Mbit/s traffic
 
As shown in Figure 1-14, two paths are set up between R1 and R7: R1 -> R2 -> R3 -> R6 ->
R7 and R1 -> R2 -> R4 -> R5 -> R6 -> R7. Bandwidth of the first path is 30 Mbit/s, and bandwidth
of the second path is 80 Mbit/s. TE allocates traffic properly based on bandwidth, preventing
link congestion. For example, 30 Mbit/s and 50 Mbit/s services are running between R1 and R7.
TE distributes the 30 Mbit/s traffic to the 30 Mbit/s path and the 50 Mbit/s traffic to the 80 Mbit/
s path.
The following characteristics of MPLS make TE implementation possible:
l Explicit paths can be specified for LSPs.
l Label forwarding is easier to manage and maintain than IP forwarding.
l MPLS TE occupies fewer resources than other TE implementations.
The MPLS TE technology integrates the MPLS technology with TE. MPLS TE can reserve
resources by setting up LSPs along a specified path to prevent network congestion and balance
network traffic. MPLS TE has the following advantages:
l MPLS TE can reserve resources to ensure the quality of services during the establishment
of LSPs.
l The behaviors of an LSP can be easily controlled based on the attributes of the LSP such
as priority and bandwidth.
l LSP establishment consumes a few resources and does not affect other network services.
l MPLS allows traffic aggregation and disaggregation, which is more flexible than IP
forwarding.
l Backup path and fast reroute (FRR) protect the network communication upon a failure of
a link or a node.
These advantages make MPLS TE the optimal TE solution. MPLS TE allows service providers
(SPs) to fully leverage existing network resources to provide diverse services, optimize network
resources, and efficiently manage the network.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
18

1.3.3 MPLS-based 6PE
IPv6 Provider Edge (6PE) is a technology for transition from IPv4 to IPv6. The 6PE technology
allows ISPs to provide access services for scattered IPv6 networks over the existing IPv4
backbone network. In this way, CEs on IPv6 islands can communicate with each other through
existing IPv4 PEs.
On an MPLS 6PE network shown in Figure 1-15:
l 6PE routers exchange IPv6 routing information with CEs using IPv6 routing protocols.
l 6PE routers exchange IPv6 routing information with each other using MP-BGP and allocate
MPLS labels to IPv6 prefixes.
l 6PE routers exchange IPv4 routing information with Ps using IPv4 routing protocols and
establish LSPs between 6PE routers and Ps using MPLS.
Figure 1-15 Process of packet forwarding using MPLS 6PE
P
IPv6 
Network 
Customer 
site
IPv6 
Network 
Customer
 site
6PE
IPv4/MPLS network
6PECE
MP-BGP
IPv6 
routing 
protocol CE
IPv6 
routing 
protocol
IPv4 routing
 protocol IPv4 routing
 protocol
IPv6 L2 L1 IPv6 IPv6 L2 IPv6 
 
Figure 1-15 shows the process of IPv6 packet forwarding on an MPLS 6PE network. IPv6
packets must carry outer and inner labels when being forwarded on the IPv4 backbone network.
The inner label maps the IPv6 prefix, while the outer label maps the LSP between 6PEs.
The MPLS 6PE technology allows ISPs to connect existing IPv4/MPLS networks to IPv6
networks by simply upgrading PEs. To ISPs, the MPLS 6PE technology is an efficient solution
for transition to IPv6.
1.3.4 PBR to an LSP
Policy-based routing (PBR) selects a route according to a user-defined policy for security and
load balancing. The router supports the PBR to an LSP. On an MPLS network, IP packets that
meet the filtering policy can be forwarded through a specified LSP.
In Figure 1-16, RouterA, RouterB, RouterC, RouterD, and RouterE are on the existing network.
RouterF and RouterG are added to provide new services. Traffic is forwarded as follows:
l Traffic for existing services is forwarded through the existing network.
l Traffic for new services is forwarded by RouterF and RouterG.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
19

Figure 1-16 Application of the PBR to an LSP
RouterA RouterC
RouterB
RouterF
RouterD
RouterG
RouterE
 
To forward some traffic of new services through the existing network, configure the PBR to an
LSP on RouterA. In this manner, traffic meeting the specified policy can be forwarded through
the LSP on the existing network.
You can also use the PBR to the LSP with LDP FRR to divert some traffic to the backup LSP
for load balancing when the backup LSP is idle relatively.
1.4 References
The following table lists the references.
Document No. Description
RFC3031 Multiprotocol Label Switching Architecture
RFC3036 LDP Specification
RFC3032 MPLS Label Stack Encoding
RFC3443 Time To Live (TTL) Processing in Multi-Protocol Label Switching
(MPLS) Networks
RFC3034 Use of Label Switching on Frame Relay Networks Specification
RFC2702 Requirements for Traffic Engineering Over MPLS
RFC3209 RSVP-TE: Extensions to RSVP for LSP Tunnels
RFC4364 BGP/MPLS IP Virtual Private Networks (VPNs)
RFC2598 An Expedited Forwarding PHB
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
20

2 MPLS LDP
About This Chapter
2.1 Introduction to MPLS LDP
2.2 Principles
2.3 References
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
21

