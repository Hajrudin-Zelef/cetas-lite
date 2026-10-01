---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-24
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [3210, 3372]
sha256: f39a64eee9701295704c7808167698f4fbf23220cf078b6b724070172c947617
---

# A line starting with the # sign is comments.

l CSPF calculates the shortest path between the ingress and egress, and SPF calculates the
shortest path between a node and each of other nodes on a network.
l CSPF uses tunnel constraints but not link costs between neighboring nodes as the metric.
l CSPF does not support load balancing and uses three tie-breaking policies to determine a
path if multiple paths have the same metric.
3.2.5 Path Establishment
3.2.5.1 Path Establishment Modes
CR-LSP Establishment Modes
A CR-LSP can be established statically or dynamically.
Establishment of a static CR-LSP relies on manual configuration by a network administrator.
This section describes how a dynamic CR-LSP is established using the RSVP-TE signaling
protocol.
RSVP-TE Overview
The Resource Reservation Protocol (RSVP) is designed for the integrated service model and
used on each node along a path for resource reservation. The bandwidth reservation capability
makes it a suitable signaling protocol for establishing an MPLS TE tunnel.
RSVP-TE is an extension of RSVP to meet MPLS TE requirements. RSVP-TE extends RSVP
in the following aspects:
l RSVP-TE appends Label Request objects to Path messages to request labels. Resv
messages carry Label objects that are used to allocate labels.
l The extended RSVP messages can carry information about path constraint parameters, in
addition to label binding information.
l RSVP-TE provides the resource reservation function by supporting MPLS TE bandwidth
constraints through the extended objects.
RSVP Messages
RSVP has the following message types:
l Path message: This message is sent from a sender to receivers to collect path information
of the passing nodes.
l Resv message: This message is sent upstream by the receiver hop-by-hop to respond to the
Path message, require resource reservation.
l PathErr message: This message is sent upstream by an node to report errors in processing
of the Path messages.
l ResvErr message: This message is sent downstream by an node if errors occur during the
processing of the Resv messages.
l PathTear message: This message is sent to remove path state of the passing nodes.
l ResvTear message: This message is sent to remove resource reservation state on the node.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
66

l ResvConf message: This message is sent downstream by the sender hop-by-hop to confirm
the resource reservation requests. It is sent only when the Resv message contains the
RESV_CONFIRM object.
l Srefresh message: This message refreshes the states of RSVP neighbors.
RSVP-TE Principles
Table 3-4 lists RSVP-TE principles.
Table 3-4 RSVP-TE principles
Function Module Description
3.2.5.2
Establishment of
Dynamic CR-LSPs
A CR-LSP is established over a path calculated by CSPF or an
explicit path on the ingress.
3.2.5.3 Maintenance
of Dynamic CR-
LSPs
l Path Status Maintenance
RSVP-TE sends messages to maintain path status on each node.
l Fault Advertisement
RSVP nodes send advertisements to notify upstream and
downstream nodes of faults that occur during path establishment
or maintenance.
l Path Teardown
A CR-LSP is torn down and releases labels and bandwidths on
each node. The ingress initiates the request for a teardown.
 
3.2.5.2 Establishment of Dynamic CR-LSPs
To establish dynamic CR-LSPs, an ingress node sends Path messages to an egress node and the
egress node sends Resv messages to the ingress node. Path messages are used to create RSVP
sessions and maintain path states, so each node along the path that receives a Path message
creates a path state block (PSB). Resv messages carry resource reservation information, so each
node along the path that receives a Resv message creates a reserved state block (RSB) and the
allocated label.
Figure 3-16 shows the process of establishing an RSVP-TE CR-LSP.
Figure 3-16 Process of establishing an RSVP-TE CR-LSP
Path Path Path
ResvResv
if1 if0 if1 if0 if1 if0
PE1 P1 P2 PE2 Resv
1 2 3
4 5 6
 
1. PE1 uses CSPF to calculate a path between PE1 and PE2. The IP address of every hop on
this path has been specified. PE1 generates a Path message and creates a PSB. PE1 then
adds the explicit route object (ERO) field containing a list of IP addresses calculated by
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
67

CSPF, and sends the Path message to P1 along the path specified by the ERO. Table 3-5
lists information carried in a Path message.
Table 3-5 Path message on PE1
Object Value
SESSION Source: PE1-if1; Destination: PE2-if0
RSVP_HOP PE1-if1
EXPLICIT_ROUTE P1-if0; P2-if0; PE2-if0
LABEL LABEL_REQUEST
 
2. After P1 receives the Path message, P1 parses the message and creates PSB based on the
Path message. P1 then generates a new Path message and sends it to P2 based on the ERO.
Table 3-6 lists information carried in a Path message.
l In the previous step, PE1 updates the RSVP_HOP field in the Path message to the IP
address of the outbound interface when the message is transmitted from PE1 to P1.
Similarly, P1 updates the RSVP_HOP field in the Path message to the IP address of the
outbound interface when the message is transmitted from P1 to P2.
l P1 deletes the local LSR ID and IP addresses of the inbound and outbound interfaces
from the ERO field in the Path message.
Table 3-6 Path message on P1
Object Value
SESSION Source: PE1-if1; Destination: PE2-if0
RSVP_HOP P1-if1
EXPLICIT_ROUTE P2-if0: PE2-if0
LABEL LABEL_REQUEST
 
3. P2 deals with the received Path message in the same process as that on P1. P2 creates a
PSB based on the Path message, updates the new Path message, and sends it to PE2. Table
3-7 lists information carried in a Path message.
Table 3-7 Path message on P2
Object Value
SESSION Source: PE1-if1; Destination: PE2-if0
RSVP_HOP P2-if1
EXPLICIT_ROUTE PE2-if0
LABEL LABEL_REQUEST
 
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
68

4. After PE2 receives a Path message, PE2 knows that itself is the egress of the CR-LSP to
be set up based on the Tunnel Address field in the Session object. PE2 then allocates a label
and bandwidth resources, and generates an RSB based on the Resv message. The Resv
message is sent to P2 and carries the label which is allocated by PE2.
PE2 extracts an IP address from the RSVP_HOP field of the received Path message and
uses it as the destination IP address of the Resv message. The Resv message is forwarded
along the reverse path. Therefore, the Resv message does not carry the ERO field. Table
3-8 lists information carried in a Resv message.
NOTE
If the Resv message contains the RESV_CONFIRM object, nodes receiving the Resv message must
send a ResvConf message to the generator of the Resv message to confirm the request for resource
reservation.
Table 3-8 Resv message on PE2
Object Value
SESSION Source: PE2-if0; Destination: PE1-if1
RSVP_HOP PE2-if0
LABEL 3
RECORD_ROUTE PE2-if0
 
5. When P2 receives the Resv message, P2 create an RSB based on the Resv message, allocates
a new label, updates the Resv message, and sends the message to P1. Table 3-9 lists
information carried in a Resv message.
Table 3-9 Resv message on P2
Object Value
SESSION Source: PE2-if0; Destination: PE1-if1
RSVP_HOP P2-if0
LABEL 17
RECORD_ROUTE P2-if0; PE2-if0
 
