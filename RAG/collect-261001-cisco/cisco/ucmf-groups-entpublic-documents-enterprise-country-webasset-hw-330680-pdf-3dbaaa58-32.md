---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-32
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [4535, 4694]
sha256: a73c29e2ad20e47970985dc95d3b13e66dbd3de32832eb99206dc25e27c5a722
---

# A line starting with the # sign is comments.

NOTE
A bypass CR-LSP supports the combination of protection types. For example, manual protection, node
protection, and bandwidth protection can be implemented together on a bypass CR-LSP.
Implementation
The PLR implements TE FRR as follows:
1. Establishes a primary CR-LSP.
The establishment of a primary CR-LSP is the same as that of a common CR-LSP. The
only difference is that the tunnel ingress adds SESSION_ATTRIBUTE related flags to the
Path message when establishing a primary CR-LSP. For example, a local protection flag
indicates a bypass CR-LSP that needs to be bound to the primary CR-LSP. A bandwidth
protection flag indicates that bandwidth protection is required.
2. Binds a bypass CR-LSP to the primary CR-LSP.
Searching for a suitable bypass CR-LSP is also called bypass CR-LSP binding. This process
is completed before a CR-LSP switchover is performed. A bypass CR-LSP can be bound
to a primary CR-LSP only with the local protection flag.
Before binding two CR-LSPs, a node must obtain the following information according to
the RRO field in the Resv message. Some of them are listed as follow:
l Outbound interface
l Next Hop Label Forwarding Entry (NHLFE)
l Label switching router (LSR) ID of the MP
l Label allocated by the MP
l Protection type
The PLR node on a primary CR-LSP already obtains information about the next hop
(NHOP) or next NHOP (NNHOP). If the egress LSR ID of the bypass CR-LSP is equal to
the LSR ID of the NHOP, link protection is provided. If the egress LSR ID of the bypass
CR-LSP is equal to the LSR ID of the NNHOP, node protection is provided. As shown in
Figure 3-24, bypass CR-LSP 1 provides link protection and bypass CR-LSP 2 provides
node protection.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
91

Figure 3-24 Binding between bypass and primary CR-LSPs
Primary CR-LSP
LSRA
LSRB LSRC LSRD
LSRE
LSRF LSRG LSRH
PLR
Bypass CR-LSP 1
Link protection
Bypass CR-LSP 2 
Node protection
NHOP NNHOP
Link Fault
Node Fault
 
If multiple bypass CR-LSPs are established, the PLR selects the one with the highest
priority. The PLR prioritizes bypass CR-LSPs in the following order:
l Bandwidth protection
l Non-bandwidth protection
l Manual protection
l Auto FRR protection
l Node protection
l Link protection
Both bypass CR-LSPs 1 and 2 shown in Figure 3-24 are manually configured and provide
bandwidth protection. Bypass CR-LSP 1, which protects a link, has a lower priority than
bypass CR-LSP 2, which protects a node. In such a scenario, bypass CR-LSP 2 is then
bound to a primary CR-LSP. If bypass CR-LSP 1 only protects bandwidth and bypass CR-
LSP 2 only protects a link, bypass CR-LSP 1 is then bound to the primary CR-LSP.
After the binding is complete, the primary CR-LSP NHLFE records the bypass CR-LSP
NHLFE index and an inner label that the MP allocates for the primary CR-LSP. The label
is used to forward traffic from the MP to the next hop along the primary CR-LSP.
3. Performs fault detection.
l Link protection directly uses a data link layer protocol to detect and report faults. The
speed of fault detection at the data link layer depends on the link type.
l Node protection uses the link layer protocol to detect link faults. If no fault occurs on
a link, the RSVP Hello mechanism is used to detect faults of the protected node or it
is used with the BFD for RSVP mechanism.
After a link or node fault is detected, FRR switching triggers immediately.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
92

NOTE
l If node protection is enabled, only the link between the protected node and PLR is protected. The
PLR cannot detect faults in the link between the protected node and MP.
l Link fault detection, BFD detection, and RSVP Hello detection detect faults at a speed in
descending order.
4. Performs a traffic switchover.
If the primary CR-LSP fails, both data traffic and RSVP messages switch to the bypass
CR-LSP, and the switchover event is reported upstream. The PLR pushes both an inner
label that the MP assigns for the primary CR-LSP and an outer label assigned for the bypass
CR-LSP into a packet. The outer label is removed at the penultimate hop of the bypass CR-
LSP, and the packet, only with the inner label, arrives at the MP. The MP forwards the
packet to the next hop along the primary CR-LSP.
Figure 3-25 shows nodes on the primary and bypass CR-LSPs and their allocated labels
and forwarding behaviors. The bypass CR-LSP provides node protection. If the link
between LSRB and LSRC fails or LSRC fails, LSRB (PLR) swaps an inner label 1024 for
an inner label 1022, pushes an outer label 34 into the packet, and forwards the packet over
the bypass CR-LSP. After the packet arrives at LSRD, the LSRD forward the packet to the
next hop LSRE. Packet forwarding after TE FRR switching in Figure 3-25 shows the
detailed forwarding process.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
93

Figure 3-25 Schematic diagram of packet forwarding before and after TE FRR
Swap Pop 36
PLR MP
LSRA LSRB LSRC LSRD LSRE
IP
1022
IP
34
1022
IP
35
1022
IP
36
1022
IP
Swap 1024→1022
Push  34
1024
IP
label assigned for the 
Primary CR-LSP
label assigned for the 
Bypass CR-LSP
Swap
Link Fault
Node Fault
PLR MP
LSRA LSRB LSRC LSRD LSRE
1025
IP
1024
IP
1022
IP
Primary CR-LSP
Bypass CR-LSP
Swap Swap Pop
IP
Packet forwarding before 
TE FRR switching
Packet forwarding after 
TE FRR switching
 
5. Performs a traffic switchback.
After TE FRR switching is complete, the PLR (ingress) attempts to reestablish the primary
CR-LSP using the Make-Before-Break mechanism. Service traffic and RSVP messages
switch from the bypass CR-LSP back to the primary CR-LSP after the primary CR-LSP is
successfully reestablished. The reestablished CR-LSP is called a modified CR-LSP. The
Make-Before-Break mechanism allows the original primary CR-LSP to be torn down only
after the modified CR-LSP is set up successfully.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
94

