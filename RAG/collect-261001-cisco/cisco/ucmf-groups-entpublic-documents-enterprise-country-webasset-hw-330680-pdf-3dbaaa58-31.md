---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-31
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [4368, 4534]
sha256: 4a864c6ff920a6187e199d99a35111129558257eaaf8ee74780140168a1b5ed4
---

# A line starting with the # sign is comments.

reestablished, traffic on the original backup CR-LSP (if it is transmitting traffic) switches
to this new backup CR-LSP, and the original backup CR-LSP is torn down.
4. Fault detection is implemented.
CR-LSP backup supports the following fault detection functions:
l The RSVP-TE fault advertisement mechanism sends signaling packets to detect faults
at a low speed.
l Bidirectional forwarding detection (BFD) for CR-LSP rapidly detects faults. This is a
recommended function.
5. A traffic switchover is implemented.
If a primary CR-LSP fails, the ingress attempts to switch traffic from the primary CR-LSP
to a hot-standby CR-LSP. If the hot-standby CR-LSP is unavailable, the ingress attempts
to switch traffic to an ordinary backup CR-LSP. If the ordinary backup CR-LSP is
unavailable, the ingress attempts to switch traffic to a best-effort path.
6. A traffic switchback is implemented.
Traffic switches back to a path based on the available CR-LSPs. Traffic will switch first to
the primary CR-LSP, which has the highest priority. If the primary CR-LSP is unavailable,
traffic will switch to the hot-standby CR-LSP. The ordinary CR-LSP has the lowest priority.
Dynamic Bandwidth Protection for Hot-standby CR-LSPs
Hot-standby CR-LSPs support dynamic bandwidth protection. The dynamic bandwidth
protection function allows a hot-standby CR-LSP to obtain bandwidth resources only after the
hot-standby CR-LSP takes over traffic from a faulty primary CR-LSP. This function uses
network resources efficiently and reduces network costs.
Dynamic bandwidth protection ensures that the hot-standby CR-LSP does not use bandwidth,
while the primary CR-LSP is transmitting traffic. The dynamic bandwidth protection process is
as follows:
1. If the primary CR-LSP fails, traffic immediately switches to the hot-standby CR-LSP with
0 bit/s bandwidth. The ingress uses the Make-Before-Break mechanism to establish a hot-
standby CR-LSP.
2. After the new hot-standby CR-LSP has been successfully established, the ingress switches
traffic to this CR-LSP and tears down the hot-standby CR-LSP with 0 bit/s bandwidth.
3. After the primary CR-LSP recovers, traffic switches back to the primary CR-LSP. The hot-
standby CR-LSP then releases the bandwidth it uses and the ingress establishes another
hot-standby CR-LSP with no bandwidth.
Overlapping Path for a Hot-standby CR-LSP
The path overlapping function can be configured for hot-standby CR-LSPs. This function allows
the path of a hot-standby CR-LSP partially overlaps the path of the primary CR-LSP. After the
hot-standby CR-LSP is established, it can protect traffic on the primary CR-LSP.
3.2.9.5 TE FRR
TE FRR protects links and nodes on CR-LSPs bound to an MPLS TE tunnel. If a link or node
fails, TE FRR rapidly switches traffic to a backup path, minimizing traffic loss.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
88

Background
A link or node failure triggers a primary/backup CR-LSP switchover. IGP routes of the backup
path need to converge, and CSPF recalculates a path over which a CR-LSP is established. Traffic
is dropped during this process.
TE FRR can be used to prevent traffic loss. After a link or node fails, TE FRR establishes a
bypass CR-LSP, which excludes the faulty link or node. The bypass CR-LSP can rapidly take
over traffic, minimizing traffic loss. The ingress can reestablish a primary CR-LSP.
Related Concepts
Figure 3-22 Local protection
PLR Primary CR-LSP
Bypass CR-LSP 
LSRA LSRB LSRC LSRD
LSRE
MP
 
Table 3-16 describes TE FRR concepts.
Table 3-16 TE FRR concepts
Concept Description
Primary CR-LSP A CR-LSP that is protected.
Bypass CR-LSP A CR-LSP that protects the primary CR-LSP. The bypass CR-LSP
is usually in the idle state and transmits few data. If the bypass CR-
LSP needs to forward service data when it protects the primary CR-
LSP, sufficient bandwidth must be allocated to the bypass CR-LSP.
Point of Local Repair
(PLR)
The ingress of the bypass CR-LSP. It must be on the path of the
primary CR-LSP. The PLR can be the ingress, not the egress of the
primary CR-LSP.
Merge point (MP) The egress of the bypass CR-LSP. It must be on the path of the
primary CR-LSP. The MP cannot be the ingress of the primary CR-
LSP.
 
Table 3-16 describes TE FRR protection functions.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
89

Table 3-17 TE FRR protection functions
Cla
ssif
ied
By
Type Description
Obj
ect
to
be
prot
ecte
d
Link
protectio
n
As shown in Figure 3-23, the PLR (LSRB) and MP (LSRC) are directly
connected, and the primary CR-LSP passes through the direct link. Bypass
CR-LSP 1 protects the direct link.
Node
protectio
n
As shown in Figure 3-23, a primary CR-LSP between the PLR (LSRB)
and MP (LSRD) passes through LSRC. Bypass CR-LSP 2 protects LSRC
on the primary CR-LSP.
Ban
dwi
dth
Bandwid
th
protectio
n
The bandwidth of a bypass CR-LSP is higher than or equal to that of the
primary CR-LSP. The bypass CR-LSP protects the primary CR-LSP and
its bandwidth.
Non-
bandwidt
h
protectio
n
No bandwidth is assigned to a bypass CR-LSP. The bypass CR-LSP
protects only the path of the primary CR-LSP.
Imp
lem
enta
tion
Manual
protectio
n
A manually configured bypass CR-LSP is established and bound to a CR-
LSP that is to be protected. If a link or node on the protected CR-LSP fails,
traffic automatically switches to the bypass CR-LSP.
Auto
FRR
protectio
n
An Auto FRR-enabled node automatically establishes a bypass CR-LSP.
The node binds the bypass CR-LSP to a primary CR-LSP if the node
receives an FRR protection request and the FRR topology requirements
are met.
 
Figure 3-23 TE FRR link and node protection
MP
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
Link Fault
MP 
Node Fault
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
90

