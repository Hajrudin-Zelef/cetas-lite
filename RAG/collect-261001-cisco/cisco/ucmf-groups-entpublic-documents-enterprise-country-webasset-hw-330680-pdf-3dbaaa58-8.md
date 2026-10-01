---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-8
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [886, 1015]
sha256: e89c7d384d7b99cd016ecf55500d22e868b974aed449fe8e6545df759af29217
---

# A line starting with the # sign is comments.

node, each packet is assigned a label and the lower 3 bits in the DSCP field are mapped to
the EXP field. A change in the value of the EXP field on the MPLS network determines
the PHB used when the packet leaves the MPLS network. The egress node maps the EXP
field to the DSCP field.
On an L2VPN, the MPLS label is in the outer layer of an encapsulated packet. Therefore, the
802.1p field of VLAN packets needs to be mapped to the EXP field.
1.2.7 MPLS Ping/Tracert
Introduction to MPLS Ping/Tracert
On an MPLS network, the control panel used for setting up an LSP cannot detect the failure in
data forwarding of the LSP. This makes network maintenance difficult. The MPLS ping and
tracert mechanisms detect LSP errors and locate faulty nodes.
MPLS ping is used to check network connectivity and host reachability. MPLS tracert is used
to check the network connectivity and host reachability, and to locate network faults. Similar to
IP ping and tracert, MPLS ping and tracert use MPLS echo request packets and MPLS echo
reply packets to check LSP availability. MPLS echo request packets and echo reply packets are
both encapsulated into User Datagram Protocol (UDP) packets. The UDP port number of the
MPLS echo request packet is 3503, which can be identified only by MPLS-enabled devices.
An MPLS echo request packet carries FEC information to be detected, and is sent along the
same LSP as other packets with the same FEC. In this manner, the connectivity of the LSP is
checked. MPLS echo request packets are forwarded to the destination end using MPLS, while
MPLS echo reply packets are forwarded to the source end using IP. Routers set the destination
address in the IP header of the MPLS echo request packets to 127.0.0.1/8 (local loopback
address) and the TTL value is 1. In this way, MPLS echo request packets are not forwarded
using IP forwarding when the LSP fails so that the failure of the LPS can be detected.
MPLS Ping
Figure 1-12 MPLS network
LSP
5.5.5.5/32 4.4.4.4/32
RouterA RouterB RouterC RouterD
Loopback0Loopback0
As shown in Figure 1-12, RouterA establishes an LSP to RouterD. RouterA performs MPLS
ping on the LSP by performing the following steps:
1. RouterA checks whether the LSP exists. (On a TE tunnel, the router checks whether the
tunnel interface exists and the CR-LSP has been established.) If the LSP does not exist, an
error message is displayed and the MPLS ping stops. If the LSP exists, RouterA performs
the following operations.
2. RouterA creates an MPLS echo request packet and adds 4.4.4.4 to the destination FEC
stack in the packet. In the IP header of the MPLS echo request packet, the destination
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
15

address is 127.0.0.1/8 and the TTL value is 1. RouterA searches for the corresponding LSP,
adds the LSP label to the MPLS echo request packet, and sends the packet to RouterB.
3. Transit nodes RouterB and RouterC forward the MPLS echo request packet based on
MPLS. If MPLS forwarding on a transit node fails, the transit node returns an MPLS echo
reply packet carrying the error code to RouterA.
4. If no fault exists along the MPLS forwarding path, the MPLS echo request packet reaches
the LSP egress node RouterD. RouterD returns a correct MPLS echo reply packet after
verifying that the destination IP address 4.4.4.4 is the loopback interface address. MPLS
ping is complete.
MPLS Tracert
As shown in Figure 1-12, RouterA performs MPLS tracert on RouterD (4.4.4.4/32) by
performing the following steps:
1. RouterA checks whether an LSP exists to RouterD. (On a TE tunnel, the router checks
whether the tunnel interface exists and the CR-LSP has been established.) If the LSP does
not exist, an error message is displayed and the tracert stops. If the LSP exists, RouterA
performs the following operations.
2. RouterA creates an MPLS echo request packet and adds 4.4.4.4 to the destination FEC
stack in the packet. In the IP header of the MPLS echo request packet, the destination
address is 127.0.0.1/8. Then RouterA adds the LSP label to the packet, sets the TTL value
to 1, and sends the packet to RouterB. The MPLS echo request packet contains a
downstream mapping TLV that carries downstream information about the LSP at the
current node, such as next-hop address and outgoing label.
3. Upon receiving the MPLS echo request packet, RouterB decreases the TTL by one and
finds that TTL times out. RouterB then checks whether the LSP exists and the next-hop
address and whether the outgoing label of the downstream mapping TLV in the packet is
correct. If so, RouterB returns a correct MPLS echo reply packet that carries the downstream
mapping TLV of RouterB. If not, RouterB returns an incorrect MPLS echo reply packet.
4. After receiving the correct MPLS echo reply packet, RouterA resends the MPLS echo
request packet that is encapsulated in the same way as step 2 and sets the TTL value to 2.
The downstream mapping TLV of this MPLS echo request packet is replicated from the
MPLS echo reply packet. RouterB performs common MPLS forwarding on this MPLS
echo request packet. If TTL times out when RouterC receives the MPLS echo request
packet, RouterC processes the MPLS echo request packet and returns an MPLS echo reply
packet in the same way as step 3.
5. After receiving a correct MPLS echo reply packet, RouterA repeats step 4, sets the TTL
value to 3, replicates the downstream mapping TLV in the MPLS echo reply packet, and
sends the MPLS echo request packet. RouterB and RouterC perform common MPLS
forwarding on this MPLS echo request packet. Upon receiving the MPLS echo request
packet, RouterD repeats step 3 and verifies that the destination IP address 4.4.4.4 is the
loopback interface address. RouterD returns an MPLS echo reply packet that does not carry
the downstream mapping TLV. MPLS tracert is complete.
When routers return the MPLS echo reply packet that carries the downstream mapping TLV,
RouterA obtains information about each node along the LSP.
1.3 Applications
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
16

1.3.1 MPLS-based VPN
The traditional VPN transmits private network data over the public network using tunneling
protocols, such as the Generic Routing Encapsulation (GRE), Layer 2 Tunneling Protocol
(L2TP), and Point to Point Tunneling Protocol (PPTP).
An MPLS-based VPN has similar security as a frame relay (FR). Devices on the MPLS-based
VPN do not require the configuration of the GRE or L2TP tunnel. Network delay is minimized
because datagrams are not encapsulated or encrypted.
As shown in Figure 1-13, the MPLS-based VPN integrates private network branches through
an LSP to form a unified network. The MPLS-based VPN controls the interconnection between
VPNs. Figure 1-13 shows the devices on the MPLS-based VPN.
l Customer edge (CE) is an edge device on a customer network. The CE can be a router, a
switch, or a host.
l Provider edge (PE) is an edge device on a service provider network.
l Provider (P) is a backbone device on an SP network. A P is not directly connected to CEs.
Ps only need to possess basic MPLS forwarding capabilities and do not maintain
information about a VPN.
Figure 1-13 MPLS-based VPN
CE
CE
CE Service provider's 
backbone
CE
VPN 1
Site
Site
Site
Site
VPN 1
VPN 2
PE
PE
PE
P
P P
PVPN 2
 
