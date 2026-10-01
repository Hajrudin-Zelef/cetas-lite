---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-15
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [1913, 2007]
sha256: 7e18baa35b4d0d42f3cb814573ad7514b8b33a4f38e6b690666fef1fb23d347c
---

# A line starting with the # sign is comments.

However, if the optimal route from N1 to D is load balanced between N1 -> N2 -> D and N1 -
> S -> D, S as the downstream neighbor of N1 does not necessarily receive the liberal label from
N1. In addition, even though S receives the liberal label (LDP distributes labels for each peer)
and is configured with LDP FRR, traffic may still go to S after traffic switches to N1, which
leads to a loop. This occurs till the route from N1 to D is converged to N1 -> N2 -> D.
2.2.8 LDP GR
LDP graceful restart (GR) ensures uninterrupted traffic forwarding on the restarter with the help
of a neighbor (Helper) when an active main board/standby main board (AMB/SMB) switchover
or a protocol restart occurs on the restarter.
When the AMB/SMB switchover occurs on a device that is not capable of GR, the neighbor
deletes the LSP because the LDP session becomes Down. As a result, traffic cannot be forwarded
and services are interrupted for a short period. To prevent service interruption, LDP GR can be
configured to keep labels consistent before and after the AMB/SMB switchover or the protocol
restart. LDP GR ensures uninterrupted MPLS forwarding. Figure 2-9 shows the detailed
process.
1. Before the AMB/SMB switchover, LDP neighbors negotiate the GR capability during the
LDP session establishment.
2. After the AMB/SMB switchover, the GR detects helper starts the LDP session failure and
starts the GR Reconnect timer. The GR helper retains the forwarding entries related to the
GR restarter and marks the entries with the stale tag.
3. After performing the AMB/SMB switchover, the GR restarter starts the Forwarding State
Holding timer. Before the Forwarding State Holding timer times out, the GR restarter
retains all MPLS forwarding entries before the restart and marks the entries with the stale
tag. The GR restarter then sends an LDP Initialization message to the GR helper. When the
Forwarding State Holding timer times out, the GR restarter performs step 6.
4. Before the GR Reconnect timer times out, the LDP session is reestablished. The GR helper
deletes the Forwarding State Holding timer, starts the GR Recovery timer, and retains the
forwarding entries with the stale tag.
5. Before the GR Recovery timer times out, the neighbors exchange Label Mapping messages
with each other and restore the label binding before the AMB/SMB switchover. When the
GR Recovery timer times out, the GR helper deletes all forwarding entries with the stale
tag.
6. The GR process ends. The GR restarter deletes all forwarding entries with the stale tag.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
38

Figure 2-9 LDP GR implementation
GR Restarter GR Helper
AMB/SMB 
switchover or 
protocol restart
Negotiate the GR capability
Exchange the Label Mapping 
messages
Start the Forwarding 
State Holding Timer
Label all MPLS 
forwarding 
entries as Stale
Send an LDP Initialization message
Reserve forwarding 
entries labeled Stale
Detects that the LDP 
session with the GR 
restarter fails and 
label forwarding 
entries related to the 
GR restarter as Stale
Reconnect the LDP session
Restore the 
label binding
Disable GR
Restore the 
label binding
Disable GR
 
2.2.9 LDP NSR
The non-stop routing (NSR) technology is an innovation based on non-stop forwarding (NSF)
technology. If a software or hardware fault occurs on the control plane, NSR ensures
uninterrupted forwarding and connection of the control plane. In addition, the control plane of
a neighbor will not detect any fault.
LDP NSR is implemented using the synchronization of the master and slave control boards.
During the startup, the slave control board backs up data of the master control board in batches
to ensure data consistency on both boards. LDP NSR simultaneously notifies the master and
slave control boards of receipt of packets and backs up these packets in real time. In this manner,
the slave control board synchronizes data with the master control board. NSR ensures that after
switchover, the slave control board can quickly take over services from the original master
control board, while the neighbor will not detect the fault on the local router.
LDP NSR synchronizes the following key data between the master and slave control boards:
l LSP forwarding entries
l Key resources such as labels and cross connections
l LDP protocol control blocks
2.2.10 LDP Security Mechanisms
MD5 Authentication
Message-digest algorithm 5 (MD5) is a standard digest algorithm defined in RFC 1321. A typical
application of MD5 is to calculate a message digest to prevent message spoofing. The MD5
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
39

