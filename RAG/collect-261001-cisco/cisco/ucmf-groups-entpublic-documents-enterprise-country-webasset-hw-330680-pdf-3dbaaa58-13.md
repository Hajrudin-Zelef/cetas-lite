---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-13
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [1673, 1797]
sha256: 53d798ffaf064b7904f401db35d6a923a446ed6b9ec8307eb2048ab3c0744c9a
---

# A line starting with the # sign is comments.

Synchronization between LDP and static routes implements LSP switchover in the following
scenarios:
l The LDP session on the primary LSP fails (not due to a link failure).
When an LDP session is established, MPLS traffic is forwarded through LinkA. If LDP is
disabled or faulty on LSRB, the LDP session between LSRA and LSRB fails. However,
the link between LSRA and LSRB is running properly and static routes are active. MPLS
traffic is interrupted between LSRA and LSRD during LSP switchover to LinkB.
After synchronization between LDP and static routes is enabled on LSRA, static routes
automatically switch to LinkB when the LDP session is Down. This ensures uninterrupted
MPLS traffic during an LSP switchover.
l The primary LSP recovers from a fault.
If the link between LSRA and LSRB fails, the LSP switches to LinkB. When the link
between LSRA and LSRB recovers, the LSP switches back to LinkA. At this time, the
backup LSP cannot be used, but the new LSP has not been established. MPLS traffic
between LSRA and LSRD is interrupted during this period.
After synchronization between LDP and static routes is enabled on LSRA, static routes
become active only when the LDP session is Up, which ensures uninterrupted traffic.
2.2.5 Synchronization Between LDP and IGP
Background
The LDP convergence speed depends on the convergence speed of IGP routes, which indicates
IGP convergence is faster.
l On an MPLS network with the primary and backup links, the following problems occur:
1. When the primary link fails, an IGP route of the backup link becomes reachable and
a backup LSP over the backup link takes over traffic. After the primary link recovers,
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
32

the IGP route of the primary link becomes reachable before an LDP session is
established over the primary link. As a result, traffic is dropped when being transmitted
using the reachable IGP route along the unreachable LSP.
2. When the IGP route of the primary link is reachable and an LDP session between
nodes on the primary link fails, traffic is directed using the IGP route of the primary
link, whereas the LSP over the primary link is torn down. Because a preferred IGP
route of the backup link is unavailable, an LSP over the backup link cannot be
established, causing traffic loss.
l When the active/standby switchover occurs on a node, the LDP session establishment is
later than the IGP GR completion. IGP advertises the maximum cost of the link, causing
route flapping.
Related Concepts
Synchronization between LDP and IGP is implemented by suppressing IGP from advertising
normal routes to ensure convergence performed by synchronization between LDP and IGP.
Synchronization between LDP and IGP involves three timers:
l Hold-down timer: used to control the period for establishing the IGP neighbor relationship.
l Hold-max-cost timer: used to control the period for advertising the maximum cost of the
link.
l Delay timer: used to control the period for waiting for the LSP establishment.
Implementation
Figure 2-5 Synchronization between LDP and IGP for revertive switchover
LSR1 LSR2
LSR3
LSR4
LSR5 LSR6
1
2
The primary LSP recovers 
from a physical fault.
The IGP status of the 
primary LSP is normal and 
the LDP session is faulty.
Primary LSP
Backup LSP
 
l During active/standby link switchover, synchronization between LDP and IGP takes effect.
As shown in Figure 2-5, the processes of synchronization between LDP and IGP differ in
the following scenarios:
1. The primary link recovers from a physical fault.
a. The faulty link recovers.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
33

b. An LDP session is set up between LSR2 and LSR3. IGP suppresses the
establishment of the neighbor relationship and starts the Hold-down timer as
required.
c. Traffic keeps traveling through the LSP over the backup link.
d. After the LDP session is set up, Label Mapping messages are exchanged and
then synchronization between IGP and LDP starts.
e. The IGP establishes a neighbor relationship and switches traffic back to the
primary link, and the LSP is reestablished and its route converges on the primary
link (in milliseconds).
2. IGP on the primary link is normal and the LDP session is faulty.
a. An LDP session between nodes along the primary link becomes defective.
b. LDP notifies the IGP primary link of the session fault. IGP starts the Hold-max-
cost timer and advertises the maximum cost on the primary link.
c. The IGP route of the backup link becomes reachable.
d. An LSP is established over the backup link and the LDP module on LSR2 delivers
forwarding entries.
The Hold-max-cost timer can be configured to always advertise the maximum cost of
the primary link. This setting allows traffic to keep traveling through the backup link
before the LDP session over the primary link is reestablished.
l During active/standby system switchover, the procedure for synchronization between LDP
and IGP is as follows:
1. An IGP on the Restarter advertises a normal cost value and starts a Delay timer, waiting
for an LDP session to be set up. Then IGP ends the GR process.
2. If the Delay timer expires before the LDP session is set up, IGP starts a Hold-max-
cost timer, and advertises the maximum cost value of the link.
3. After the LDP session is established or the Hold-max-cost timer expires, IGP
advertises the actual link cost and updates the IGP route.
4. The helper retains the IGP route and LSP. After the LDP session on the helper goes
Down, the LDP module does not notify the IGP module of the session status change.
This indicates that IGP keeps advertising the actual link cost, preventing traffic or LSP
switchover.
2.2.6 BFD for LSP
A Bidirectional Forwarding Detection (BFD) session is established on an LSP. BFD is used to
quickly detect faults on the LSP, providing end-to-end protection.
BFD is used to detect faults on the data plane of the MPLS LSP, and the format of BFD packets
is fixed. When a BFD session is associated with a unidirectional LSP, the reverse link can be an
IP link, an LSP, or a TE tunnel.
Implementation
BFD detects LSPs in asynchronous mode. The ingress and the egress nodes send BFD control
packets to each other periodically.
l If any of the ingress and the egress nodes does not receive BFD control packets sent by the
peer within a detection period, LSP status is considered to be Down and a message that the
LSP is Down is sent to the LSP Management (LSPM) module.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
34

