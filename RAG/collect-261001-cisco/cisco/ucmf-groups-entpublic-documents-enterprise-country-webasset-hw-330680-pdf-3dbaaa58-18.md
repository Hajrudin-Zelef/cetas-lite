---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-18
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution", "voice"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [2246, 2414]
sha256: c46ceefde692fef6ed632567d5f7378f6826631b6752053427b7f11dbca8d6e3
---

# A line starting with the # sign is comments.

When LDP over mGRE is used to construct an L3VPN network, traffic between branches is
forwarded through Hubs.
There is a delay in transmitting traffic of voice services in this case because voice services require
point-to-point transmission. DSVPN can be used to dynamically establish a tunnel so that
branches can directly communicate. In addition to the preceding advantages, DSVPN has the
following problems:
l During establishment of a tunnel between Spokes, traffic between branches is forwarded
through the Hub. After the tunnel is set up, traffic is switched to the tunnel. In this process,
packet mis-sequencing may occur on the receiver.
l A tunnel between branches needs to be dynamically maintained. If the Spoke performance
is low, maintaining a large amount of tunnel information will cause the Spoke to deteriorate
and affect services.
In addition, MPLS requires that labels in MPLS packets be swapped on LSPs between ingress
and egress nodes. After LDP over mGRE is used, all Spokes are equivalent to PEs on an MPLS
network. If Spokes directly communicate with each other (this scenario is not used), each two
Spokes need to establish an LDP and distribute labels. Consequently, label resources of Spokes
are insufficient. You can use the Hub to forward traffic between branches (the Hub, similar to
the P on an MPLS network, maintains the LDP neighbor relationship and swaps packet labels).
A Spoke only needs to establish one LSP with the Hub so that the Spoke can communicate with
other Spokes. Because MPLS label swapping is used, only traffic on the Hub is increased. Hub
performance is less affected.
Although LDP over mGRE does not use DSVPN to establish tunnels between Spokes, it provides
the Spoke-Hub-Spoke path to better transmit traffic between Spokes. LDP over mGRE is often
used when the MPLS network needs to be extended.
2.3 References
The following table lists the references.
Document No. Description
RFC3036 LDP Specification
RFC3215 LDP State Machine
RFC5443 LDP IGP Synchronization
RFC3478 Graceful Restart Mechanism for Label Distribution
Protocol
RFC1321 The MD5 Message-Digest Algorithm
RFC3037 LDP Applicability
RFC3899 Maximum Transmission Unit Signalling Extensions
for the Label Distribution Protocol
RFC3270 Multi-Protocol Label Switching (MPLS) Support of
Differentiated Services
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
45

3 MPLS TE
About This Chapter
3.1 Introduction to MPLS TE
3.2 Principles
3.3 Applications
3.4 References
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
46

3.1 Introduction to MPLS TE
Definition
Multiprotocol Label Switching (MPLS) traffic engineering (TE) establishes label switched paths
(LSPs) satisfying specific constraints and transparently transmits traffic over the LSPs based on
labels.
Purpose
A node on a conventional IP network selects the shortest path as an optimal route, regardless of
other factors, such as bandwidth. The shortest path may be congested with traffic, but other
available paths are idle.
Figure 3-1 Conventional routing
R6
R7
R1
R2
R3
R4
R5
80M
40M
Path 1
Path 2
 
Each link on the network shown in Figure 3-1 has a bandwidth of 100 Mbit/s and the same
metric value. R1 sends R4 traffic at 80 Mbit/s, and R7 sends R4 traffic at 40 Mbit/s. Interior
Gateway Protocol (IGP) calculates the shortest path for traffic, such as Path1 and Path2 in Figure
3-1. Traffic on this network is forwarded along the path R2→R3→R4. As a result, the path R2
→R3→R4 may be congested because of overload, while the path R2→R5→R6→R4 is idle.
Traffic engineering techniques can load balance traffic to the idle paths to prevent traffic
congestion caused by uneven resource allocation.
Conventional TE solutions are as follows:
l AP TE: AP TE controls network traffic by adjusting the metric of a path. This method
eliminates congestion only on some links. Adjusting a metric is difficult on a complex
network because a link change affects multiple routes.
l ATM TE: Because existing Interior Gateway Protocols (IGPs) are topology-driven and
consider only network connectivity, they cannot present some dynamic factors such as
bandwidth and traffic characteristics. The IP over asynchronous transfer mode (ATM)
overlay model can solve this problem. ATM TE directs some traffic to virtual connections
(VCs) using the overlay model so that traffic can be properly scheduled and allocated and
QoS guarantee is ensured. However, ATM TE has high costs and low extensibility.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
47

A scalable and simple solution is required to implement TE on a large-scale network. MPLS,
an overlay model, allows a virtual topology to be established over a physical topology and maps
traffic to the virtual topology. MPLS can be integrated with TE. MPLS TE is introduced.
MPLS TE can be used on the network shown in Figure 3-1 to address congestion. MPLS TE
sets up two LSPs: Path1 with bandwidth of 80 Mbit/s and Path2 with bandwidth of 40 Mbit/s.
MPLS TE directs traffic to the two LSPs, preventing congestion.
Figure 3-2 MPLS TE
R6
R7
R1
R2
R3
R4
R5
Path 1
Path 2
 
Benefits
MPLS TE effectively schedules, allocates, and uses existing network resources to provide
sufficient bandwidth and support for quality of service (QoS). MPLS TE helps enterprises
minimize expenditures without requiring hardware upgrades. TE is implemented based on
MPLS techniques and is easy to deploy and maintain on live networks. MPLS TE supports a
range of reliability techniques, which helps backbone networks achieve carrier- and device-class
reliability.
3.2 Principles
3.2.1 Basic Concepts
This section describes the following basic concepts:
l LSP Tunnel
l MPLS TE Tunnel
l Link Attributes
l Tunnel Attributes
LSP Tunnel
After a label is added to a packet on the ingress node of an LSP, the packet is forwarded based
on the label. Traffic forwarding is transparent to intermediate nodes, so an LSP can be considered
as an LSP tunnel.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
48

MPLS TE Tunnel
Multiple LSPs are bound together to form an MPLS TE tunnel. The MPLS TE tunnel involves
the following entities:
l Tunnel interface: a P2P virtual interface that encapsulates packets. Similar to a loopback
interface, a tunnel interface is a logical interface.
l Tunnel ID: a decimal number that identifies an MPLS TE tunnel and facilitates tunnel
planning and management.
l LSP ID: a decimal number that identifies an LSP, which facilitates LSP planning and
management.
A primary LSP with an LSP ID 2 is established along the path LSRA -> LSRB -> LSRC ->
LSRD -> LSRE on the network shown in Figure 3-3. A backup LSP with an LSP ID 1024 is
established along the path LSRA -> LSRF -> LSRG -> LSRH -> LSRE. The two LSPs are in a
tunnel named Tunnel 0/0/0 with a tunnel ID 100.
Figure 3-3 MPLS TE Tunnel and LSPs
Primary LSP
LSRA
LSRB LSRC LSRD
LSRE
LSRF LSRG LSRH
Backup LSP
MPLS TE Tunnel
MPLS TE Tunnel:
Tunnel Interface = Tunnel 0/0/0
Tunnel ID = 100
Primary LSP ID = 2
Backup LSP ID = 1024
 
