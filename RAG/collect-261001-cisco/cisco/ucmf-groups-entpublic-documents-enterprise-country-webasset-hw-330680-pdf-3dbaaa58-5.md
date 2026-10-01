---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-5
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution", "ethernet"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [388, 531]
sha256: 6269686e6c411aaa3ed303121a2bb6003b805c1c89f31b40ed182c899dea101e
---

# A line starting with the # sign is comments.

A label contains the following fields:
l Label: indicates the value field of a label. The length is 20 bits.
l Exp: indicates the bits used for extension. The length is 3 bits. Generally, this field is used
for the class of service (CoS) that serves in a manner similar to Ethernet 802.1p.
l S: identifies the bottom of a label stack. The length is 1 bit. MPLS supports multiple labels,
namely, the label nesting. When the S field is 1, the label is at the bottom of the label stack.
l TTL: indicates the time to live. The length is 8 bits. This field is the same as the TTL in IP
packets.
Labels are encapsulated between the data link layer and the network layer. Labels can be
supported by all data link layer protocols.
Figure 1-5 shows the position of the label in a packet.
Figure 1-5 Position of a label in a packet
Label Layer 3 payloadLink layer header Layer 3 header
 
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
5

Label Space
The label space is the value range of the label. The following describes the label space
classification:
l 0 to 15: indicates special labels. For details about special labels, see Table 1-1.
Table 1-1 Special labels
Label Value Label Description
0 IPv4 Explicit
NULL Label
The label must be popped out, and the packets must be
forwarded based on IPv4. If the egress node allocates
a label whose value is 0 to the LSR at the penultimate
hop, the LSR at the penultimate hop pushes label 0 to
the top of the label stack and forwards the packet to the
egress node. When the egress node recognizes that the
value of the label carried in the packet is 0, the egress
node pops it out. The label 0 is valid only at the bottom
of the label stack.
1 Router Alert
Label
A label that is only valid when it is not at the bottom of
a label stack. The label is similar to the Router Alert
Option field in IP packets. After receiving such a label,
the node sends it to a local software module for further
processing. Packet forwarding is determined by the
next-layer label. If the packet needs to be forwarded
continuously, the node pushes the Router Alert Label
to the top of the label stack again.
2 IPv6 Explicit
NULL Label
The label must be popped out, and the packets must be
forwarded based on IPv6. If the egress node allocates
a label with the value of 2 to the LSR at the penultimate
hop, the LSR pushes label 2 to the top of the label stack
and forwards the packet to the egress node. When the
egress node recognizes that the value of the label
carried in the packet is 2, the egress node immediately
pops it out. The label 2 is valid only at the bottom of
the label stack.
3 Implicit
NULL Label
When the label with the value of 3 is swapped on an
LSR at the penultimate hop, the LSR pops the label out
and forwards the packet to the egress node. Upon
receiving the packet, the egress node forwards the IP
or VPN packet.
4 to 13 Reserved None.
14 OAM Router
Alert Label
A label for operation, administration and maintenance
(OAM) packets over an MPLS network. MPLS OAM
sends OAM packets to monitor LSPs and notify faults.
OAM packets are transparent on transit nodes and the
penultimate LSR.
15 Reserved None.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
6

l 16 to 1023: indicates the label space shared by static LSPs and static constraint-based routed
LSPs (CR-LSPs).
l 1024 or above: indicates the label space for dynamic signaling protocols, such as Label
Distribution Protocol (LDP), Resource Reservation Protocol-Traffic Engineering (RSVP-
TE), and Multiprotocol Extensions for BGP (MP-BGP).
Label Stack
A label stack is a set of arranged labels. An MPLS packet can carry multiple labels at the same
time. The label next to the Layer 2 header is called the top label or the outer label. The label next
to the Layer 3 header is called the bottom label or inner label. Theoretically, MPLS labels can
be nested without any limit.
Figure 1-6 Label stack
Outer labelLink layer header Layer3 payloadLayer3 headerInner label
Label Stack
 
The label stack organizes labels according to the rule of Last-In, First-Out. The labels are
processed from the top of the stack.
Label Operations
Information about basic label operations is a part of the label forwarding table. The operations
are described as follows:
l Push: When an IP packet enters an MPLS domain, the ingress node adds a new label to the
packet between the Layer 2 header and the IP header. Alternatively, an LSR adds a new
label to the top of the label stack, namely, the label nesting.
l Swap: When a packet is transferred within the MPLS domain, a local node swaps the label
at the top of the label stack in the MPLS packet for the label allocated by the next hop
according to the label forwarding table.
l Pop: When a packet leaves the MPLS domain, the label is popped out of the MPLS packet.
Alternatively, the top label of the label stack is popped out at the penultimate hop on an
MPLS network to decrease the number of labels in the stack.
In fact, the label is useless at the last hop of an MPLS domain. The penultimate hop popping
(PHP) feature applies. On the penultimate node, the label is popped out of the packet to
reduce the size of the packet that is forwarded to the last hop. Then, the last hop directly
forwards the IP packet or forwards the packet by using the second label.
PHP is configured on the egress node. The egress node supporting PHP allocates the label
with the value of 3 to the penultimate hop.
The VPN Option C scenario supports the following action to process labels:
l Swappush: swaps an existing inner label for a new one and then pushes an outer label of
other tunnel into a packet.
Enterprise Data Communication Products
Feature Description - MPLS 1 MPLS Overview
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
7

l Popgo: pops out an inner label from a packet and then pushes an outer label of other tunnel
into the packet.
1.2.3 Establishing LSPs
Procedure for Establishing LSPs
Usually, MPLS allocates labels to packets and establishes an LSP through which MPLS forwards
packets.
The downstream LSR allocate labels to packets sent to the upstream LSR. As shown in Figure
1-7, the downstream LSR identifies FEC based on the destination address, allocates a label to
the specified FEC, and records the mapping between the label and FEC. The downstream LSR
then encapsulates the mapping relationship into a message and sends it to the upstream LSR. A
label forwarding table and an LSP are established.
Figure 1-7 Establishment of an LSP
To 3.3.3.3/24
Label=Z
To 3.3.3.3/24
Label=Y
To 3.3.3.3/24
Label=3
3.3.3.3/24Ingress Transit Transit Egress
LSP
 
