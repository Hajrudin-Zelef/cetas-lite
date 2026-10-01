---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-36
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "cost", "preemption", "voice"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [5172, 5336]
sha256: ad3794c74e6edd462604b5a840b8736c83f57ec6e96ecec62d5eb8f7b6177ea8
---

# A line starting with the # sign is comments.

Deployment Scenarios
RSVP GR can be used on nodes that runs RSVP-TE to establish MPLS TE tunnels to improve
device reliability.
Benefits
RSVP GR ensures uninterrupted data service transmission when the control plane performs an
AMB/SMB switchover and supports device-level reliability for MPLS TE nodes.
3.2.10 DS-TE
3.2.10.1 Background
Traditional MPLS TE reserves resources for each node along the MPLS TE tunnel to ensure
QoS, but cannot use one TE tunnel to provide differentiated services. When a tunnel transmits
voice and data services, data services may be transmitted repeatedly. Therefore, data services
must have higher drop priority than voice services. However, MPLS TE allocates the same drop
priority to data and voice flows and cannot provide differentiated services.
The Diff-Serv model controls and forwards traffic according to specific service classes, meeting
different QoS requirements. The Diff-Serv model can reserve resources for a single node, but
cannot guarantee the QoS over an entire path.
In certain scenarios, Diff-Serv and MPLS TE must be used together to meet service requirements.
For example, a path may trasnmit both voice and data services. The total delay time of voice
flows needs to be reduced to ensure QoS guarantee of voice services.
Assume that the Diff-Serv model is used to classify service types and a single MPLS TE tunnel
is used to transmit a type of service. If the link or node becomes faulty, network topology changes,
or LSP preemption occurs, voice traffic on a link may exceed the bandwidth and voice services
are delayed.
For example, in Figure 3-33, the bandwidth of each link is 100 Mbit/s and the link cost is the
same. Voice flows pass through the links R1 -> R4 and R2 -> R4, and bandwidths of the links
are 60 Mbit/s and 40 Mbit/s respectively. When voice flows on the link R1 -> R4 are transmitted
through the TE tunnel of Path1, voice flows occupy 60% bandwidth on the link R3 -> R4. When
voice flows on the link R2 -> R4 are transmitted through the TE tunnel of Path2, voice flows
occupy 40% bandwidth on the link R7 -> R4.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
105

Figure 3-33 Networking where MPLS TE and DiffServ are used
R1
R2
R6
R4
R7
R5
R3
HSI:20M
HSI:20M
VoIP:40M
VoIP:60M
Voice service traffic
Data service traffic
Voice service traffic path
Path2
Path1
 
As shown in Figure 3-34, when the link between R3 and R4 becomes faulty, the CR-LSP
between R1 and R4 is changed to Path3. The reason is that Path3 is the shortest path with
sufficient bandwidth. In this case, voice flows on the link R7 -> R4 occupy 100% bandwidth,
causing a long delay in transmitting voice flows.
Figure 3-34 Link failure
R1
R2
R6
R4
R7
R5
R3
Path2
Path3
HSI:20M
HSI:20M
VoIP:40M
VoIP:60M
100%
Link fault
Voice service traffic
Data service traffic
Voice service traffic path
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
106

To solve the preceding problem, DiffServ-aware Traffic Engineering (DS-TE) is used. DS-TE
can efficiently use network resources and reserve resources for different service flows.
MPLS DS-TE combines MPLS TE and Diff-Serv to provide QoS guarantee.
MPLS DS-TE uses the Class Type (CT) so that MPLS TE can allocate resources based on the
type of traffic and provide differentiated services. To provide differentiated services, DS-TE
divides the LSP bandwidth into one to eight parts, each part corresponding to one Class of Service
(CoS). A set of bandwidth of an LSP or a group of LSPs with the same CoS are called a CT.
In Figure 3-33, multiple CT LSPs can be used. An LSP is divided into multiple CTs to transmit
traffic of different CoS values. VoIP and HSI services on the links R1 -> R4 and R2 -> R4 are
transmitted by different CTs of the same MPLS TE tunnel so that voice flows occupy a proper
bandwidth percentage, as shown in Figure 3-35.
Figure 3-35 MPLS DS-TE
R1
R2
R6
R4
R7
R5
R3
HSI:20M
HSI:20M
VoIP:40M
VoIP:60M CT for HSI:20%
CT for VoIP:60%
CT for VoIP:40%
CT for HSI:20%
Voice service traffic
Data service traffic
MPLS TE tunnel
 
When the link R3 -> R4 becomes faulty, VoIP and HSI services on the link R3 -> R4 are switched
to the link R1 -> R3 -> R5 -> R6 -> R4, as shown in Figure 3-36. After switching, voice flows
on the link R1 -> R4 still occupy a proper bandwidth percentage.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
107

Figure 3-36 Traffic switching after a link failure
R1
R2
R6
R4
R7
R5
R3
HSI:20M
HSI:20M
VoIP:40M
VoIP:60M
CT for HSI:20%
CT for VoIP:40%
CT for HSI:20%
CT for VoIP:60%
Voice service traffic
Data service traffic
MPLS TE tunnel
Link fault
 
3.2.10.2 Basic Concepts
DS Field
To carry out the Diff-Serv model, RFC 2474 redefines the ToS field in the IPv4 packet header
as the Differentiated Services (DS) field. The high-order 2 bits in the DS field are reserved, and
the low-order 6 bits specify the DS CodePoint (DSCP).
Per Hop Behavior
Per Hop Behavior (PHB) describes how the packets with the same DSCP value are forwarded
to the next hop.
The IETF defines three standardized PHBs: expedited forwarding (EF), assured forwarding
(AF), and best-effort (BE). BE is the default PHB.
CT
To carry out differentiated services, the DS-TE model divides the bandwidth of an LSP into one
to eight parts. Each part of bandwidth is allocated with a different service class. The set of
bandwidth of one LSP or a group of LSPs with the same service class is called a class type (CT).
One CT can transmit the traffic of a single service type.
As defined in the IETF, the DS-TE supports a maximum of eight CTs. CTs can be represented
as CTi. The value of "i" ranges from 0 to 7.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
108

