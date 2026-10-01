---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-17
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [2105, 2245]
sha256: d6755c12cb8f4f37432f9da3c51bc2390e178e83b31675e457cc369a08f843c2
---

# A line starting with the # sign is comments.

LDP establishes liberal LSPs, not inter-area LDP LSPs, for aggregated routes. In this situation,
LDP cannot provide required backbone network tunnels for VPN services.
Therefore, in the situation shown in Figure 2-10, configure LDP to search for routes according
to the longest match rule for establishing LSPs. There is already an aggregated route to 1.3.0.0/24
in the routing table of LSRA. When LSRA receives a Label Mapping message (such as the
carried FEC is 1.3.0.1/32) from Area 10, LSRA searches for a route according to the longest
match rule defined in RFC 5283. Then, LSRA finds information about the aggregated route to
1.3.0.0/24, and uses the outbound interface and next hop of this route as those of the route to
1.3.0.1/32. In this manner, LDP can establish inter-area LDP LSPs.
2.2.12 LDP over GRE/mGRE
GRE provides a mechanism to encapsulate packets of a protocol into packets of another protocol.
This allows packets to be transmitted over heterogeneous networks. A channel for transmitting
heterogeneous packets is called a tunnel.
A GRE tunnel can be established using the following tunnel interfaces:
l GRE tunnel interface
A GRE tunnel interface is a point-to-point virtual interface used to encapsulate packets,
and has the source address, destination address, tunnel interface IP address, and
encapsulation type.
l mGRE tunnel interface
An mGRE tunnel interface is a point-to-multipoint virtual interface used in DSVPN
applications, and has the source address, destination address, and tunnel interface IP
address.
The destination IP address of a GRE tunnel interface is manually configured, whereas the
destination IP address of an mGRE tunnel is resolved by the Next Hop Resolution Protocol
(NHRP). An mGRE tunnel interface has multiple remote ends because there are multiple
GRE tunnels on the interface.
If backbone network devices are not enabled with MPLS or do not support MPLS, LDP LSPs
cannot be established. As a result, L2VPN or L3VPN services cannot be deployed. LDP over
GRE/mGRE addresses the preceding problem.
LDP over GRE
LDP over GRE technology establishes an LDP LSP on a GRE tunnel interface configured with
MPLS LDP to transmit MPLS LDP packets.
As shown in Figure 2-11, L2VPN or L3VPN services are deployed between PE1 and PE2 of
an enterprise. Because backbone network devices may be not enabled with MPLS or do not
support MPLS, an LDP LSP across a GRE tunnel needs to be set up between PE1 and PE2.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
42

Figure 2-11 Deploying LDP over GRE on L2VPN/L3VPN networking (all P devices do not
support MPLS)
PE1 PE2
L2VPN
Site1
L2VPN
Site2
L3VPN
Site1
L3VPN
Site2
GRE Tunnel
P1 P2
IP/MPLS  
backbone
LDP LSP for L3VPN
LDP LSP for L2VPN
CE1 CE2
CE4 CE3
 
As shown in Figure 2-12, backbone network device P2 supports MPLS, whereas P1 does not
support MPLS. A GRE tunnel can be established between PE1 and P2 so that an LDP LSP is
set up across the GRE tunnel.
Figure 2-12 Deploying LDP over GRE on L2VPN/L3VPN networking (some P devices do not
support MPLS)
PE1 PE2
L2VPN
Site1
L2VPN
Site2
L3VPN
Site1
L3VPN
Site2
P1 P2
IP/MPLS
backbone
LDP LSP for L3VPN
LDP LSP for L2VPN
CE1 CE2
CE4 CE3
GRE Tunnel
 
LDP over mGRE
As shown in Figure 2-13, the IP/MPLS backbone network is established in the enterprise
headquarters. Enterprise branches in other areas need to connect to the IP/MPLS backbone
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
43

network through public network devices. If an L3VPN network needs to be established between
the enterprise headquarters and branches, LDP over GRE can be used. There are the following
problems:
l Public addresses of branch devices are not fixed.
The branch Spoke uses a dynamically allocated public address. When LDP over GRE is
used, GRE cannot be used to establish a GRE tunnel because public addresses of devices
in multiple branches are not fixed. As a result, L3VPN services cannot be deployed.
l There are many branches.
When there are many branches, a large number of Spokes exist. If GRE is used to establish
GRE tunnels (assume that devices use fixed public addresses), many GRE interfaces need
to be created on Hubs. The configuration is complex and maintenance is difficult.
Dynamic Smart Virtual Private Network (DSVPN) technology solves the preceding problem.
NHRP solves problems caused by non-fixed public addresses of branch devices and dynamically
establishes a tunnel between the headquarters and a branch. mGRE allows multiple GRE tunnels
to be set up on a tunnel interface, simplifying the Hub configuration.Similar to LDP over GRE,
LDP over mGRE technology establishes an LDP LSP on an mGRE tunnel interface configured
with MPLS LDP to transmit MPLS LDP packets.
As shown in Figure 2-13, the enterprise establishes a backbone network. The headquarters Hub-
P uses a static address to connect to the public network and branch Spoke-PE uses dynamic
addresses to connect to the public network. The enterprise requires that L3VPN services be
deployed between all PEs. Because branch devices connect to the enterprise IP/MPLS backbone
network through the public network and Spoke-PEs' public addresses are not fixed, LDP over
mGRE is used to establish LDP LSPs between PEs across GRE tunnels.
Figure 2-13 Deploying LDP over mGRE in Hub-Spoke networking
Spoke-PE2 Spoke-PE3
Site2 Site3
GRE Tunnel
Internet
Hub-P1
Site1
PE1
GRE Tunnel
mGRE
tunnel interface
IP/MPLS
backbone
VPN1
VPN2 VPN1
VPN2
VPN1
VPN2
 
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
44

