---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-33
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [4695, 4825]
sha256: 36fbb6ef500a136d046c6a8bf7437b8bc7b5992a73027c7d255d1b232e358156
---

# A line starting with the # sign is comments.

NOTE
FRR does not take effect if multiple nodes fail simultaneously. This means that after FRR switches data
from the primary CR-LSP to the bypass CR-LSP, all nodes on the bypass CR-LSP must be working properly
when transmitting data. If the bypass CR-LSP fails, the protected data cannot be forwarded, and the FRR
function fails. Even if the bypass CR-LSP is reestablished, it cannot forward data. Data will be restored
only after the primary CR-LSP is restored or reestablished.
Other Usage
l Board hot removal protection
Board hot removal protection protects traffic on the primary CR-LSP's outbound interface
on a PLR. If an interface board on which a protected outbound interface of a primary CR-
LSP resides is removed from a PLR, the PLR rapidly switches traffic to a bypass CR-LSP.
After the interface board is re-installed and the outbound interface of the primary CR-LSP
becomes available, traffic switches back to the primary CR-LSP.
Hot removal protection does not apply to an interface board, on which tunnel interfaces are
configured. If an interface board configured with a tunnel interface is removed, CR-LSP
information is lost and traffic is interrupted. The bypass CR-LSPs' tunnel interfaces and
the bypass CR-LSP's outbound interface must be configured on boards different from the
board configured with the bypass CR-LSP's outbound interface on the PLR.
Configuring tunnel interfaces on the main control board of the PLR is recommended. If an
interface board on which the primary CR-LSP's outbound interface is removed or fails, the
primary CR-LSP's tunnel interface enters the Stale state, and resources allocated to the
tunnel interface remain. After the interface board is re-installed, the tunnel interface
recovers and a primary CR-LSP is reestablished.
l N:1 protection
A single bypass CR-LSP can protect traffic over multiple primary CR-LSPs.
Coexistence of CR-LSP Backup and TE FRR
1. CR-LSP backup functions can be used together with TE FRR.
l Ordinary backup and TE FRR: If TE FRR detects a link fault, traffic switches to a TE
FRR bypass CR-LSP. If both the primary and TE FRR bypass CR-LSPs fail, an ordinary
backup CR-LSP is established and takes over traffic.
l Hot standby and TE FRR: If TE FRR detects a link fault, traffic switches to a TE FRR
bypass CR-LSP and then to a hot-standby CR-LSP.
2. CR-LSP backup can be associated with TE FRR.
The association improves tunnel security. The association provides the following functions
based on backup modes:
l Association between an ordinary backup CR-LSP and a TE FRR bypass CR-LSP
provides the following functions:
If a protected link or node fails, traffic switches to a bypass CR-LSP. The ingress
attempts to reestablish the primary CR-LSP, while attempting to establish an ordinary
backup CR-LSP.
If the ordinary backup CR-LSP is established successfully before the primary CR-LSP
is restored, traffic switches to the ordinary backup CR-LSP.
After the primary CR-LSP recovers, traffic switches back to the primary CR-LSP.
If the ordinary backup CR-LSP fails to be established and the primary CR-LSP does
not recover, traffic passes still through the bypass CR-LSP.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
95

l Association between a hot-standby CR-LSP and a TE FRR bypass CR-LSP provides
the following functions:
If a hot-standby CR-LSP is Up and a protected link or node fails, traffic switches to a
TE FRR bypass CR-LSP and then immediately switches to the hot-standby CR-LSP.
At the same time, the ingress attempts to restore the primary CR-LSP.
If the hot-standby CR-LSP is Down, the traffic switching procedure is the same as that
when the ordinary backup is used.
Association between ordinary backup CR-LSPs and TE FRR is recommended. An ordinary
backup CR-LSP without additional bandwidth needed is established only after the primary
CR-LSP enters the FRR-in-use state. Although the primary CR-LSP is Up, the system
attempts to establish a hot-standby CR-LSP with additional bandwidth needed.
3.2.9.6 SRLG
The shared risk link group (SRLG) functions as a constraint that is used to calculate a backup
path in the scenario where CR-LSP hot standby or TE FRR is used. This constraint helps prevent
backup and primary paths from overlapping over links with the same risk level, improving MPLS
TE tunnel reliability as a consequence.
Background
Network administrators use CR-LSP hot standby or TE FRR to improve MPLS TE tunnel
reliability. However, in real-world situations protection failures can occur, requiring the SRLG
technique to be configured as a preventative measure, as the following example demonstrates.
Figure 3-26 Networking diagram for an SRLG
PE1 PE2 P1 P2
P3
Path1
Path2
SRLG
PE1 PE2 P1 P2
P3
NE1
Logical topology
Physical topology Optical transport 
device
Shared link
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
96

The primary CR-LSP is established over the path Path1 on the network shown in Figure 3-26.
The link between P1 and P2 is protected by a TE FRR bypass CR-LSP established over the path
Path2.
In the lower part of Figure 3-26, core nodes P1, P2, and P3 are connected using a transport
network device. They share some transport network links marked in yellow. If a fault occurs on
a shared link, both the primary and bypass CR-LSPs are affected, causing an FRR protection
failure. An SRLG can be configured to prevent the bypass CR-LSP from sharing a link with the
primary CR-LSP, ensuring that FRR properly protects the primary CR-LSP.
An SRLG is a set of links at the same risk of faults. If a link in an SRLG fails, other links also
fail. If a link in this group is used by a hot-standby CR-LSP or bypass CR-LSP, the hot-standby
CR-LSP or bypass CR-LSP cannot provide protection.
Implementation
An SRLG link attribute is a number and links with the same SRLG number are in a single SRLG.
Interior Gateway Protocol (IGP) TE advertises SRLG information to all nodes in a single MPLS
TE domain. The constraint shortest path first (CSPF) algorithm uses the SRLG attribute together
with other constrains, such as bandwidth, to calculate a path.
The MPLS TE SRLG works in either of the following modes:
l Strict mode: The SRLG attribute is a necessary constraint used by CSPF to calculate a path
for a hot-standby CR-LSP or an bypass CR-LSP.
l Preferred mode: The SRLG attribute is an optional constraint used by CSPF to calculate a
path for a hot-standby CR-LSP or bypass CR-LSP. For example, if CSPF fails to calculate
a path for a hot-standby CR-LSP based on the SRLG attribute, CSPF recalculates the path,
regardless of the SRLG attribute.
Usage Scenario
The SRLG attribute is used in either the TE FRR or CR-LSP hot-standby scenario.
Benefits
The SRLG attribute limits the selection of a path for a hot-standby CR-LSP or bypass CR-LSP,
which prevents the primary and bypass CR-LSPs from sharing links with the same risk level.
3.2.9.7 TE Tunnel Protection Group
A tunnel protection group protects E2E MPLS TE tunnels. If a working tunnel in a protection
group fails, traffic switches to a protection tunnel, minimizing traffic interruptions.
Related Concepts
Concepts related to a tunnel protection group are as follows:
l Working tunnel: a tunnel to be protected.
l Protection tunnel: a tunnel that protects a working tunnel.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
97

