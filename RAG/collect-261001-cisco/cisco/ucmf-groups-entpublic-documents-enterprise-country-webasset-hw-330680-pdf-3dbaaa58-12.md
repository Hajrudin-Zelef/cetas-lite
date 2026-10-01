---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-12
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution", "memory"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [1501, 1672]
sha256: 214ed14f9b3f41803880cacd5bfc3f09d9d67c217209e17ff3229b36f329534b
---

# A line starting with the # sign is comments.

Label Distribution
Control Modes
Definition Description
Ordered mode An LSR advertises the
mapping between a label and
an FEC to its upstream LSR
only when this LSR is the
outgoing node of the FEC or
receives the Label Mapping
message of the next hop for
the FEC.
l As shown in Figure 2-3,
the label distribution
mode is DU and the label
distribution control mode
is ordered. Consequently,
the LSR (the transit LSR
in the diagram) must
receive a Label Mapping
message from the
downstream LSR (the
egress node in the
diagram). Then, it can
distribute a label to the
ingress node in the
diagram.
l As shown in Figure 2-3,
if the label distribution
mode is DoD and the label
distribution control mode
is Ordered, the directly-
connected transit of the
ingress node that sends
the Label Request
message must receive a
Label Mapping message
from the downstream (the
egress node in the
diagram). Then, it can
distribute a label to the
ingress node in the
diagram.
 
Label Retention Modes
The label retention mode refers to the way an LSR processes the label mapping that it receives
but does not immediately use.
The label mapping that an LSR receives may or may not originate at the next hop.
As described in Table 2-3, two label retention modes are available.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
29

Table 2-3 Label retention modes
Label Retention Modes Definition Description
Liberal mode When receiving a Label
Mapping message from a
neighbor LSR, an LSR
retains the message
regardless of whether the
neighbor LSR is its next hop.
When the next hop of an LSR
changes due to a change in
network topology, note that:
l In Liberal mode, the LSR
can use the previous label
sent by a non-next hop to
quickly reestablish an
LSP. This requires more
memory and label space
than in conservative
mode.
l In Conservative mode,
the LSR only retains
labels sent by the next
hop. This saves memory
and label space but slows
down the reestablishment
of an LSP.
Conservative mode and
DoD mode are used
together to set up LSRs
with limited label space.
Conservative mode When receiving a Label
Mapping message from a
neighbor LSR, an LSR
retains the message only
when the neighbor LSR is its
next hop.
 
Currently, the combination of the following modes is supported:
l Combination of the DU label advertisement mode, ordered label control mode, and liberal
label retention mode
l Combination of the DoD label advertisement mode, ordered label control mode, and
conservative label retention mode
NOTE
On the device, LDP by default works in the DU label advertisement mode, ordered label control mode,
and liberal label retention mode.
2.2.3 LDP Label Filtering Mechanism
By default, an LSR receives and sends Label Mapping messages for all FECs, resulting in the
establishment of a large number of LDP LSPs. The establishment of a large number of LDP
LSPs consumes a great deal of LSR resources. As a result, the LSR may be overburdened. An
outbound or inbound LDP policy needs to be configured to reduce the number of Label Mapping
messages to be sent or received, reducing the number of LSPs to be established and saving
memory.
Outbound LDP Policy
LDP outbound policies are used to filter out Label Mapping messages sent to peers. If a FEC
matches no outbound policy, neither a transit LSP nor an egress LSP can be established. If a pair
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
30

of or all peers have the same restriction on the FEC range when sending Label Mapping
messages, the same outbound policy can be configured for the pair of or all peers.
An LDP outbound policy filters out Label Mapping messages only for the FEC, but not those
for L2VPN. Meanwhile, the LDP outbound policy specifies the FEC range.
In addition, the outbound LDP policy supports split horizon. After split horizon is configured,
an LSR distributes labels only to its upstream LDP peers.
Before sending Label Mapping messages only for the FEC to a peer, an LSR checks whether an
outbound policy is configured.
l If no outbound policy is configured, the LSR sends the Label Mapping message.
l If an outbound policy is configured, the LSR checks whether the FEC in the Label Mapping
message is within the range defined in the outbound policy. If the FEC is within the FEC
range, the LSR sends a Label Mapping message for the FEC; if the FEC is not within the
FEC range, the LSR does not send a Label Mapping message.
Inbound LDP Policy
LDP inbound policies are used to filter out Label Mapping messages received from peers. If a
FEC matches no inbound policy, Label Mapping messages are not accepted. If a pair of or all
peers have the same restriction on the FEC range when receiving Label Mapping messages, the
same inbound policy can be configured for the pair of or all peers.
An LDP inbound policy filters out Label Mapping messages only for the FEC, but not those for
L2VPN. Meanwhile, the LDP inbound policy specifies the FEC range for non-BGP routes.
An LSR checks whether an inbound policy mapped to a FEC is configured before receiving a
Label Mapping message for the FEC.
l If no inbound policy is configured, the LSR receives the Label Mapping message.
l If an inbound policy is configured, the LSR checks whether the FEC in the Label Mapping
message is within the range defined in the inbound policy. If the FEC is within the FEC
range, the LSR receives the Label Mapping message for the FEC; if the FEC is not in the
FEC range, the LSR does not receive the Label Mapping message.
If the FEC fails to pass an outbound policy on an LSR, the LSR receives no Label Mapping
message for the FEC.
One of the following results may occur:
l If a DU LDP session is established between an LSR and its peer, a liberal LSP is established.
This liberal LSP cannot function as a backup LSP after LDP FRR is enabled.
l If a DoD LDP session is established between an LSR and its peer, the LSR sends a Release
message to tear down label-based bindings.
NOTE
An LSP that is distributed with a label but is not successfully established called a liberal LSP.
2.2.4 Synchronization Between LDP and Static Routes
Synchronization between LDP and static routes applies to MPLS networks where primary and
backup LSPs exist. LSPs are established between LSRs based on static routes. When the LDP
session on the primary LSP fails (not due to a link failure) or the primary LSP is restored, MPLS
traffic is interrupted for a short time.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
31

As shown in Figure 2-4, LSRA and LSRD are connected using static routes. LDP establishes
primary and backup LSPs between LSRA and LSRD based on static routes, and LinkA is the
primary path.
Figure 2-4 LSP switchover based on synchronization between LDP and static routes
LSRA LSRD
LSRC
LSRB
LinkA
LinkB
 
