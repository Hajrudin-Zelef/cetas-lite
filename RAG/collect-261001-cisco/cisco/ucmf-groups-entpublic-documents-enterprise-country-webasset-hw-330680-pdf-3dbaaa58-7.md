---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-7
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution", "voice"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [742, 885]
sha256: 612737b20de725cca94a3c84574759961b024e07bef6512748b8f7646b3fac54
---

# A line starting with the # sign is comments.

– If the label value is equal to or greater than 16, a new label replaces the label in
the MPLS packet. At the same time, the EXP field and TTL field are processed.
The MPLS packet with the new label is forwarded to the next hop.
– If the label value is 3, the label is popped out of the MPLS packet. At the same
time, the EXP field and TTL field are processed. The packet is forwarded through
IP routes, or in accordance with its next layer label.
l Forwarding on the egress node
– When the egress node receives IP packets, it checks the FIB and performs IP forwarding.
– When the egress node receives MPLS packets, it checks the ILM table for the label
operation type. At the same time, the egress node processes the EXP field and TTL
field.
– When the S field in the label is equal to 1, the label is the stack's bottom label and
the packet is directly forwarded through IP routes.
– When the S field in the label is equal to 0, a next-layer label exists and the packet is
forwarded according to the next layer label.
1.2.5 MPLS TTL Processing
This section describes how MPLS processes the TTL and responds to TTL timeout.
MPLS TTL Processing Modes
The TTL field in an MPLS label is 8 bits long. The TTL field is the same as that in an IP packet
header. MPLS processes the TTL to prevent loops and implement traceroute.
RFC 3443 defines two modes to process the TTL in MPLS packets: Uniform mode and Pipe
mode. By default, MPLS processes the TTL in Uniform mode.
l Uniform mode
When IP packets enter an MPLS network, the ingress node decreases the IP TTL by one
and copies it to the MPLS TTL field. The TTL field in MPLS packets is processed in
standard mode. The egress node decreases the MPLS TTL by one and maps it to the IP
TTL field. Figure 1-10 shows how the TTL field is processed on the transmission path.
Figure 1-10 TTL processing in Uniform mode
CE CE PE P PE
MPLS
IP TTL
255
IP TTL
252
IP TTL
254
MPLS 
TTL 253
IP TTL
254
MPLS 
TTL 254
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
12

l Pipe mode
As shown in Figure 1-11, the ingress node decreases the IP TTL by one and the MPLS
TTL is constant. The TTL field in MPLS packets is processed in standard mode. The egress
node decreases the IP TTL by one. In Pipe mode, the IP TTL only decreases by one on the
ingress node and one on the egress node when packets travels across an MPLS network.
Figure 1-11 TTL processing in Pipe mode
CE CE PE P PE
IP TTL
255
IP TTL
253
MPLS
MPLS
TTL 254
IP TTL
254
MPLS
TTL 253
IP TTL
254
 
In MPLS VPN applications, the MPLS backbone network needs to be hidden to ensure network
security. The Pipe mode is recommended for private network packets.
TTL Timeout Responding
On an MPLS network, an LSR receives labeled MPLS packets. The LSR generates an ICMP
TTL-expired message when the TTL of an MPLS packet times out.
The LSR returns the TTL-expired message to the sender in the following ways:
l If the LSR has a reachable route to the sender, it directly sends the TTL-expired message
to the sender through the IP route.
l If the LSR has no reachable route to the sender, it forwards the TTL-expired message along
the LSP. The egress node forwards the TTL-expired message to the sender.
In most cases, the received MPLS packet contains only one label and the LSR responds to the
sender with the TTL-expired message using the first method. If the MPLS packet contains
multiple labels, the LSR uses the second method.
The MPLS VPN packets may contain only one label when they arrive at an autonomous system
boundary router (ASBR) on the MPLS VPN, a superstratum PE (SPE) device in HoVPN
networking, or a PE device in the VPN nesting networking. These devices have no IP routes to
the sender, so they use the second method to reply to the TTL-expired messages.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
13

1.2.6 MPLS QoS Implementation
MPLS QoS, an important part in the deployment of QoS services, implements QoS using the
Differentiated Services (DiffServ) model in actual MPLS networking. MPLS QoS differentiates
data flows based on the EXP field value, which ensures low delay and low packet loss ratio for
voice and video data streams and increases network resource efficiency.
MPLS DiffServ
In the DiffServ model, network edge nodes map a service to a service class based on the QoS
requirements of the service and use the DS field (ToS field) in IP packets to identify the service.
Nodes on the backbone network apply preset policies to the service based on the DS field to
ensure service quality. The service classification and label mechanism of DiffServ are similar
to label distribution of MPLS. MPLS DiffServ combines DS distribution and MPLS label
distribution.
MPLS DiffServ is implemented as the EXP field in an MPLS packet header carriers DiffServ
per-hop behavior (PHB). An LSR must consider the MPLS EXP value when determining the
forwarding policy. MPLS DiffServ provides the following plans for determining PHBs:
l E-LSP: an LSP whose PHB is determined by the EXP field. E-LSP applies to a network
with less than eight PHBs. In this plan, a differentiated services code point (DSCP) is
mapped to a specified EXP that identifies a PHB. Packets are forwarded based on labels,
while the EXP field determines the scheduling type and drop priority at each hop. An LSP
transmits a maximum of eight PHB flows that are differentiated based on the EXP field in
the MPLS packet header. The EXP field can be determined by the ISP or mapped from the
DSCP value in a packet. In this plan, PHB information does not need to be transmitted by
signaling protocols, the label efficiency is high, and the label status is easy to maintain.
l L-LSP: an LSP whose PHB is determined by both the label and EXP field. L-LSP applies
to a network with any number of PHBs. During packet forwarding, the label of a packet
determines the forwarding path and scheduling type, while the EXP field determines the
drop priority of the packet. Labels differentiate service flows, so multiple service flows can
be transmitted over one LSP. This plan requires more labels and so occupies a large number
of system resources.
NOTE
Currently, only the E-LSP plan is supported.
MPLS DiffServ Modes
An MPLS network provides tunnels for services. MPLS L3VPN DiffServ modes include: pipe,
short pipe, and uniform.
l Pipe: The EXP field value that the ingress node adds to the MPLS label of packets is
specified by the user. If the EXP field value of the packet is changed on the MPLS network,
the change is valid only on the MPLS network. The egress node selects the PHB according
to the EXP field value of the packet. When the packet leaves the MPLS network, the
previous DSCP value becomes effective again.
l Short pipe: The EXP field value that the ingress node adds to the MPLS label of packets
is specified by the user. If the EXP field value of the packet is changed on the MPLS
network, the change is valid only on the MPLS network. The egress node selects the PHB
according to the DSCP field value of the packet. When the packet leaves the MPLS network,
the previous DSCP value becomes effective again.
l Uniform: The priorities of packets on the IP network and the MPLS network are uniformly
defined, so the priorities of the packets on the two networks are globally valid. At the ingress
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
14

