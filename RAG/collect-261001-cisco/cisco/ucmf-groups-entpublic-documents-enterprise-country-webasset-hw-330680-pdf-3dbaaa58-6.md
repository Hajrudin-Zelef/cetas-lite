---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-6
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [532, 741]
sha256: 480c90a186479d014bc302bbad3854c1f5fd16dbeb77e2a6d4385acb531c3734
---

# A line starting with the # sign is comments.

LSPs are classified into the following types:
l Static LSP: set up by the administrator.
l Dynamic LSP: set up using the routing protocols and label distribution protocols.
Establishing Static LSPs
You can manually allocate labels to set up static LSPs. The value of the outgoing label of the
upstream node is equal to the value of the incoming label of the downstream node.
The availability of a static LSP makes sense only for the local node that cannot detect the entire
LSP.
A static LSP is set up without label distribution protocols or the exchanging of control packets.
The static LSP costs little and is recommended for small-scale networks with the simple and
stable topology. The static LSP cannot change with the network topology. Instead, it needs to
be configured by an administrator.
Establishing Dynamic LSPs
Dynamic LSPs are established using label distribution protocols. As the control protocol or
signaling protocol for MPLS, a label distribution protocol defines FECs, distributes labels, and
establishes and maintains LSPs.
The following label distribution protocols apply to an MPLS network.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
8

l LDP
LDP is defined to distribute labels and used to dynamically establish LSPs. An LSR can
use LDP to map routing information on the network layer to the LSP on the data link layer.
For details about LDP, see MPLS LDP.
l RSVP-TE
RSVP-TE is an extension to RSVP and used to establish or delete constraint-based LSPs.
For details about RSVP-TE, see MPLS TE.
l MP-BGP
MP-BGP is an extension to BGP and allocates labels to MPLS VPN routes and inter-AS
VPN routes.
For details about MP-BGP, see Feature Description - IP Routing.
1.2.4 MPLS Forwarding
MPLS Forwarding Principle
The LSP that supports the PHP is used in the following example to describe how MPLS packets
are forwarded.
Figure 1-8 MPLS label distribution and packet forwarding
IP Packet
To 4.4.4.2
Label=Z Label=Y
PUSH SWAP
PHP
Label distributing
IP Packet
To 4.4.4.2
IP Packet
To 4.4.4.2
IP Packet
To 4.4.4.2
Ingress EgressTransit Transit
To 4.4.4.2/24
Label=Z
To 4.4.4.2/24
Label=Y
To 4.4.4.2/24
Label=3
4.4.4.2/24
Ingress Egress Transit Transit 4.4.4.2/24
IP Packet
To 4.4.4.2
Packet transmitting
POP
 
As shown in Figure 1-8, an LSP whose FEC is identified by the destination address 4.4.4.2/24
is set up on an MPLS network. MPLS packets are forwarded as follows:
1. The ingress node receives an IP packet destined for 4.4.4.2. Then, the ingress node adds
Label Z to the packet and forwards it.
2. The transit node receives the labeled packet and swaps labels by popping Label Z out and
pushing Label Y into the packet.
3. A transit node at the penultimate hop receives the packet with Label Y. The transit node
pops Label Y out because the label value is 3. The transit node then forwards the packet to
the egress node as an IP packet.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
9

4. The egress node receives the IP packet and forwards it to 4.4.4.2/24.
Process of MPLS Packet Forwarding
l NHLFE
The next hop label forwarding entry (NHLFE) can guide MPLS packet forwarding.
An NHLFE contains the following information:
– Tunnel ID
– Outbound interface
– Next hop
– Outgoing label
– Label operation
l FTN
FTN is a short form of FEC-to-NHLFE. The FTN indicates the mapping between a FEC
and a set of NHLFEs.
Details about the FTN can be obtained by searching for the Tunnel ID values that are not
0x0 in a FIB. The FTN is available on the ingress only.
l ILM
The incoming label map (ILM) indicates the mapping between an incoming label and a set
of NHLFEs.
The ILM contains the following information:
– Tunnel ID
– Incoming label
– Inbound interface
– Label operation
The ILM on a transit node can bind the labels to NHLFEs. The function of an ILM table
is similar to the FIB that is searched according to destination IP addresses. Therefore, you
can obtain all label forwarding information by searching an ILM table.
l Tunnel ID
To provide the same interface of a tunnel used by upper layer applications such as the VPN
and route management, the system automatically allocates an ID to each tunnel, referred
to as the tunnel ID. The tunnel ID is 32 bits long and is valid only on the local end.
When an IP packet enters an MPLS domain, the ingress node searches the FIB to check whether
the tunnel ID corresponding to the destination IP address is 0x0.
l If the tunnel ID is 0x0, the packet is forwarded along the IP link.
l If the tunnel ID is not 0x0, the packet is forwarded along an LSP.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
10

Figure 1-9 Process of MPLS packet forwarding
DEST
PUSH SWAP
Ingress Egress Transit Transit
4.4.4.2/24
0x11
Tunnel ID
FIB
GE1/0/1 0x11 PUSH 1.1.1.2 Z
NHLFE
GE1/0/1
OUT IF
0x22
Tunnel ID
POP
OPER
3.3.3.2
NEXTHOP
3
Out Label
GE1/0/1
OUT IF
0x15
Tunnel ID
SWAP
OPER
2.2.2.2
NEXTHOP
Y
Out Label
GE1/0/1
1.1.1.1/24
GE1/0/0
1.1.1.2/24
GE1/0/1
2.2.2.1/24
GE1/0/0
2.2.2.2/24
GE1/0/1
3.3.3.1/24
GE1/0/0
3.3.3.2/24
Z
In Label
GE1/0/0
In IF
ILM
0x15
Tunnel ID
Y
In Label
GE1/0/0
In IF
0x22
Tunnel ID
POP
GE1/0/1
4.4.4.1/24
4.4.4.0/24
PHP
OUT IF Tunnel ID OPER NEXTHOP Out Label
 
MPLS packets are forwarded as follows on nodes along an LSP:
l The ingress node searches the FIB and NHLFE tables.
l The transit node searches the ILM and NHLFE tables.
l The egress node searches the ILM table or RIB.
During MPLS forwarding, FIB entries, ILM entries, and NHLFEs are associated with each other
through the tunnel ID.
l Forwarding on the ingress node
The ingress node processes the forwarding of MPLS packets as follows:
1. Searches the FIB and finds the tunnel ID corresponding to the destination IP address.
2. Finds the NHLFE corresponding to the tunnel ID in the FIB and associates the FIB
entry with the NHLFE entry.
3. Checks the NHLFE for information about the outbound interface, next hop, outgoing
label, and label operation type. The label operation type is Push.
4. Pushes the obtained label into IP packets, processes the EXP field according to QoS
policy and TTL field, and sends the encapsulated MPLS packets to the next hop.
l Forwarding on the transit node
The transit node forwards the received MPLS packets as follows:
1. Checks the ILM table corresponding to an MPLS label and finds the Tunnel ID.
2. Finds the NHLFE corresponding to the Tunnel ID in the ILM table.
3. Checks the NHLFE for information about the outbound interface, next hop, outgoing
label, and label operation type.
4. Processes the MPLS packets according to the specific label value:
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
11

