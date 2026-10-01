---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-14
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [1798, 1912]
sha256: 8ca595fecbfe1678457ffb8f5391cca3174a8f9a9396b3334c2b50b4d80a57f1
---

# A line starting with the # sign is comments.

l If the LSP status changes between Up and Down frequently, BFD sends two messages of
LSP changes successively. Therefore, the detection can be performed flexibly.
l If the reverse link of the BFD control packets sent by the egress node to the ingress node
fails, the BFD session is Down.
NOTE
BFD is a bidirectional detection mechanism, but BFD for LSP is unidirectional. BFD for LSP sends BFD
control packets through LSPs on the ingress node and through IP links on the egress node. As a result,
when the ingress node does not receive BFD control packets sent through the reverse path from the egress
node, the system considers that the LSP fails no matter the fault occurs on LSP or on the reverse link.
BFD Session Setup
To check MPLS LSP connectivity, negotiation on a BFD session can be performed in the
following modes:
l Static: The negotiation on a BFD session is performed using the local discriminator (LD)
and remote discriminator (RD) that are manually configured.
l Dynamic: The negotiation on a BFD session is performed using the BFD discriminator
TLV in an LSP ping packet.
BFD detects the following types of LSPs:
l Static BFD for static LSP
l Static BFD for LDP LSP
l Dynamic BFD for LDP LSP
Figure 2-6 shows the establishment of dynamic BFD sessions that detect LDP LSPs.
1. The ingress node sends an MPLS echo request packet that carries the type-length-value
(TLV) with the type as 15 along an LSP. The packet contains an LD that the ingress node
allocates to the BFD session.
2. The egress node receives the MPLS echo request packet sent from the ingress node and
takes the contained LD as its own RD.
3. The egress node sends an MPLS echo reply packet to the ingress node. The packet contains
an LD that the egress node allocates to the BFD session.
4. The ingress node receives the MPLS echo reply packet sent by the egress node and takes
the contained LD as its own RD.
The dynamic BFD session that detects the LDP LSP is created successfully.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
35

Figure 2-6 Establishing a session of dynamic BFD for LDP LSP
Ingress Egress
Echo request packet
Ingress LD = 1
Egress RD =
Ingress LD = 1
Automatic allocation
Ingress LD = 1
Echo reply packet
Egress LD = 2 Automatic allocation
Egress LD = 2
Ingress RD =
Egress LD = 2
 
2.2.7 LDP FRR
LDP Fast Reroute (FRR) provides the fast reroute function for MPLS networks by backing up
local interfaces.
LDP FRR, in liberal label retention mode of LDP, obtains a liberal label, applies a forwarding
entry for the label, and then forwards the forwarding entry to the forwarding plane as the backup
forwarding entry for the primary LSP. When the interface is faulty (detected by the interface
itself or according to BFD detection) or the primary LSP fails (according to BFD detection),
LDP FRR fast switches traffic to the backup LSP to protect the primary LSP.
l Manually configured LDP FRR needs to be specified with the outbound interface and next
hop of the backup LSP by running a command. When the source of the liberal label matches
the outbound interface and next hop, a backup LSP can be established and its forwarding
entries can be delivered.
l LDP auto FRR depends on the implementation of IP FRR. When the source of the preserved
liberal label matches the outbound interface and next hop of the backup route, the
requirement for the policy for establishing the backup LSP is met, and no backup LSP
manually configured according to the backup route exists, a backup LSP can be established
and its forwarding entries can be delivered. The default policy of LDP auto FRR is that
LDP can use the 32-bit backup routes to establish backup LSPs. When both the manually
configured LDP FRR and LDP auto FRR meet the establishment conditions, the manually
configured LDP FRR is established preferentially.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
36

Applicable Environment
Figure 2-7 A typical applicable environment of LDP FRR (triangle topology)
LSRC
LSRA LSRB
Primary LSP
Backup LSP
 
Figure 2-7 shows a typical applicable environment of LDP FRR. The optimal route from LSRA
to LSRB is LSRA -> LSRB and the less optimal route is LSRA -> LSRC -> LSRB. A primary
LSP along the path LSRA -> LSRB is established on LSRA, and a backup LSP along the path
LSRA -> LSRC -> LSRB is established to protect the primary LSP. After receiving a label from
LSRC, LSRA compares the label with the route from LSRA to LSRB and finds that LSRC is
not the next hop of the route. LSRA preserves the label as a liberal label and applies for a
forwarding entry as the backup forwarding entry of the primary LSP. LSRA forwards the
forwarding entries of both the primary and backup LSPs to the forwarding plane. In this manner,
the primary LSP is associated with the backup LSP.
When the interface detects faults by itself, BFD detects faults on the interface, or BFD detects
that the primary LSP fails, LDP FRR is triggered. After LSP FRR is complete, traffic is switched
to the backup LSP according to the backup forwarding entry. In this manner, LSP FRR takes
effect. Then, the route is converged from LSRA-LSRB to LSRA-LSRC-LSRB. An LSP is
established on the new LSP (the original backup LSP), and the original primary LSP is deleted,
and then the traffic is forwarded along the new LSP LSRA -> LSRC -> LSRB.
Figure 2-8 A typical applicable environment of LDP FRR (rectangle topology)
N1 N2
D SPrimary LSP
Backup LSP
 
As shown in Figure 2-7, all nodes in the triangle topology supports LDP FRR, but only parts
of nodes in the rectangle topology supports LDP FRR. As shown in Figure 2-8, if the optimal
route from N1 to D is N1 -> N2 -> D (load balancing is unavailable), S receives a liberal label
from N1 and is configured with LDP FRR. When the link between S and D is faulty, traffic is
switched to the route of S -> N1 -> N2 -> D without forming a loop.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
37

