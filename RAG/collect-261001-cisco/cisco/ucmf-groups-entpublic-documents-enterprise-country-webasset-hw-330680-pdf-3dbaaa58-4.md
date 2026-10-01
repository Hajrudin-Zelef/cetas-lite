---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-4
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["asic", "copyright", "distribution"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [230, 387]
sha256: 83a23218aa90f077a2801fd0e66cb3526a9a06ef55bcbabbd7a711a8746cf198
---

# A line starting with the # sign is comments.

1.1 Introduction to MPLS
Definition
Multiprotocol Label Switching (MPLS) is technology used on IP backbone networks. MPLS
uses connection-oriented label switching on connectionless IP networks. By combining Layer
3 routing technologies and Layer 2 switching technologies, MPLS leverages flexibility of IP
routing and simplicity of Layer 2 switching.
MPLS is based on the Internet Protocol version 4 (IPv4). The core MPLS technology can be
extended to multiple network protocols, such as the Internet Protocol version 6 (IPv6), Internet
Packet Exchange (IPX), and Connectionless Network Protocol (CLNP). Multiprotocol in MPLS
means that multiple network protocols are supported.
In fact, the MPLS technology is a tunneling technology but not a service or an application. It
supports multiple protocols and services. Moreover, it ensures security of data transmission.
Purpose
The IP-based Internet in the middle 1990s stimulated data growth. However, IP technology is
inefficient in forwarding packets because software must search for routes using the longest match
algorithm. As a result, the forwarding capability of IP technology becomes a bottleneck of the
network development.
Asynchronous transfer mode (ATM) technology has been created from the evolution of network
technologies. It uses labels (particularly, cells) of fixed length and maintains a label table that
is much smaller than a routing table. Compared to IP technology, ATM technology is much
more efficient in forwarding packets. ATM technology, however, is a complex protocol with
high deployment costs, which hinders its popularity and growth.
Traditional IP technology, however, is simple with less deployment costs. People are eager to
use technology that combines advantages of IP and ATM technologies. The MPLS technology
is used.
Initially, MPLS was created to increase forwarding rates. Different from the manner in which
packets are routed and forwarded using IP technology, MPLS analyzes a packet header only on
the edge of the network rather than at each hop. In this manner, the packet processing time is
shortened.
Application-specific integrated circuit (ASIC) technology has now been developed and the
routing rate is no longer a bottleneck to the network development. As a result, MPLS no longer
has the high-speed forwarding advantages. MPLS supports multi-layer labels, and its forwarding
plane is connection-oriented. MPLS is widely used in virtual private network (VPN), traffic
engineering (TE), and quality of service (QoS).
1.2 Principles
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
2

1.2.1 Basic MPLS Architecture
MPLS Network Structure
On a typical MPLS network shown in Figure 1-1, all routers function as label switching routers
(LSRs) that exchange labels and forward packets. These LSRs construct an MPLS domain.LSRs
that reside at the edge of the MPLS domain and connect to other networks are called label edge
routers (LERs). LSRs within an MPLS domain are core LSRs.
Figure 1-1 MPLS network structure
LER
MPLS Domain 
LER
LER
Core LSRCore LSRLER LER
IP Network
IP Network
LSP
IP 
HeaderData1028 IP 
Header Data1030 IP 
Header Data1032
IP 
Header DataIP 
Header Data
IP Network
 
On IP networks, packets are forwarded based on IP addresses; in MPLS domains, packets are
forwarded based on labels.
When receiving IP packets from the connected IP network, an LER tags labels on the packets
and then forwards the labeled packets to a core LSR. When receiving labeled packets from the
core LSR, the LER removes the labels and forwards the packets to the IP network. LSRs only
forward packets based on labels.
LSPs are determined using different protocols and are established before packet forwarding. IP
packets are transmitted through the specified label switched paths (LSPs) on an MPLS network.
As shown in Figure 1-2, an LSP is a unidirectional path whose direction is the same as the data
flow. The nodes on an LSP include the ingress, transit, and egress nodes. The number of transit
nodes on an LSP varies (none, one, or multiple), but only one ingress node and one egress node
exist on the LSP.
To an LSR, all LSRs that send MPLS packets to the LSR are the upstream LSRs, and all next-
hop LSRs that receive MPLS packets from the LSR are the downstream LSRs. As shown in
Figure 1-2, for the data flow that are destined for 192.168.1.0/24, the ingress node is the upstream
to the transit node, and the transit node is the downstream to the ingress node. Similarly, the
transit node is the upstream to the egress node, and the egress node is the downstream to the
transit node.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
3

Figure 1-2 Upstream and downstream LSRs
Ingress Transit Egress
192.168.1.0/24
Downstream 
data flow
LSP
Downstream 
data flow
 
MPLS Architecture
The MPLS architecture consists of a control plane and a forwarding plane.
Figure 1-3 shows the MPLS architecture.
Figure 1-3 MPLS architecture
IP Routing Protocol
Routing Information 
Base (RIB)
Label Distributiion 
Protocol
Label Forwarding 
Information Base(LFIB)
Forwarding  Information 
Base (FIB)
Control Plane
Forwarding Plane
Receiving 
IP packets
Receiving 
labeled packets
Sending IP 
packets
Sending 
labeled packets
 
l The connectionless control plane generates and maintains routing information and labels.
On the control plane, the IP Routing Protocol module transmits routing information and
generates a routing information base (RIB); the Label Distribution Protocol module
switches labels and establishes LSPs.
l The forwarding plane, also called data plane, is connection-oriented and forwards common
IP packets and labeled MPLS packets.
The forwarding plane consists of the modules IP forwarding information base (FIB) and
label forwarding information base (LFIB). When receiving common IP packets, the
forwarding plane forwards the packets based on the IP FIB or LFIB as required. When
receiving labeled packets, the forwarding plane forwards the packets based on the LFIB.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
4

If the destination locates on an IP network, the data plane removes the labels and forwards
the packets based on the IP FIB.
1.2.2 MPLS Label
Forwarding Equivalence Class
Forwarding equivalence class (FEC) is a class-based forwarding technology that classifies the
packets with the same forwarding mode based on the destination address or mask. Packets with
the same FEC are forwarded in the same way on an MPLS network.
FEC can be defined based on the destination IP address and mask. For example, during IP
forwarding, packets with the same destination belong to a FEC according to the longest match
algorithm.
Label
A label is a short identifier that is 4 bytes long and has only local significance. It uniquely
identifies a FEC to which a packet belongs. In some cases, such as load balancing, a FEC can
be mapped to multiple incoming labels. Each label, however, represents only one FEC on a
device.
Figure 1-4 shows the encapsulation structure of the label.
Figure 1-4 Structure of an MPLS label
Label Exp S TTL
0 3119 22 23
 
