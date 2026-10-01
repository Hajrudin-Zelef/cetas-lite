---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-11
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [1310, 1500]
sha256: 21df2d70ffedb9b05d5599525fb89ab7f8ee8a85d65a971e49a4d3691ac0b7cd
---

# A line starting with the # sign is comments.

l Basic discovery mechanism: used to discover directly-connected LSR peers on a link.
An LSR periodically sends LDP Hello messages to implement the mechanism and establish
a local LDP session.
The Hello messages are encapsulated in UDP packets with the multicast destination address
and sent through LDP port 646. A Hello message carries an LDP ID and other information
(such as the hello-hold time and the transport address). If an LSR receives an LDP Hello
message on a specified interface, a potential LDP peer is connected to the same interface.
l Extended discovery mechanism: used to discover the LSR peers that are not directly
connected on a link.
An LSR periodically sends Target Hello messages to a specified destination address
according to the mechanism to establish a remote LDP session.
The Target Hello messages are encapsulated in UDP packets and carry unicast destination
addresses, sent using LDP port 646. A Target Hello message carries an LDP ID and other
information (such as the hello-hold time and the transport address). If an LSR receives a
Target Hello message, the LSR has a potential LDP peer.
Process of Establishing an LDP Session
Two LSRs exchange Hello messages to trigger the establishment of an LDP session.
Figure 2-2 shows the process of LDP session establishment.
Figure 2-2 Process for establishing an LDP session
LSRB (Passive)
192.168.1.1/32
LSRA (Active)
192.168.1.2/32
Send Hello messages
LSRA sends an Initialization message to 
negotiate parameters.
If LSRB accepts all parameters, it 
sends an Initialization message and 
a Keepalive message.
If LSRA accepts all parameters, it 
sends a Keepalive message.
Step 1
 Step 2
 Step 3
 Step 4
 Step 5
Connect a TCP connection
 
1. Two LSRs send Hello messages to each other.
2. After receiving the Hello messages carrying the transport addresses, the two LSRs use the
transport addresses to establish an LDP session. The LSR with the larger transport address
serves as the active peer and initiates a TCP connection. As shown in Figure 2-2, LSRA
serves as the active peer to initiate a TCP connection and LSRB serves as the passive peer
to wait for the initiation of the TCP connection.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
25

3. After the TCP connection is successfully established, LSRA sends an Initialization message
to negotiate parameters used to establish an LDP session with LSRB. The main parameters
include the LDP version, label advertisement mode, the Keepalive hold timer value,
maximum PDU length, and label space.
4. If LSRB rejects some parameters, it sends a Notification message to terminate the
establishment of the LDP session. If LSRB accepts all parameters, it sends an Initialization
message carrying the LDP version, label advertisement mode, the Keepalive hold timer
value, maximum PDU length, and label space, and sends a Keepalive message to LSRA.
5. If LSRA rejects certain parameters after receiving the Initialization message, it sends a
Notification message to terminate LDP session establishment. If LSRA accepts all
parameters, it sends a Keepalive message to LSRB.
After both LSRA and LSRB have accepted Keepalive messages from each other, the LDP session
is successfully established.
Advertising and Managing Labels
Label Advertisement Modes
An LSR on an MPLS network assigns a label to a specified FEC and notifies its upstream LSRs
of the label. This means that the label is specified by a downstream LSR, and is distributed from
downstream to upstream.
As described in Table 2-1, two label advertisement modes are available.
Table 2-1 Label advertisement modes
Label Advertisement
Modes
Definition Description
Downstream Unsolicited
(DU) mode
An LSR distributes labels to
a specified FEC without
having to receive Label
Request messages from its
upstream LSR.
As shown in Figure 2-3, the
downstream egress triggers
the establishment of an LSP
destined for the FEC
192.168.1.1/32 using a host
route and sends a Label
Mapping message to the
upstream transit node to
advertise the label of the host
route to 192.168.1.1/32.
Downstream on Demand
(DoD) mode
An LSR distributes labels to
a specified FEC only after
receiving Label Request
messages from its upstream
LSR.
As shown in Figure 2-3, the
downstream egress triggers
the establishment of an LSP
destined for the FEC
192.168.1.1/32 in host mode.
The upstream ingress sends a
Label Request message to the
downstream egress. After
receiving the message, the
downstream egress sends a
Label Mapping message to
the upstream LSR.
 
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
26

The label advertisement modes on upstream and downstream LSRs must be the same.
NOTE
When DU is used, LDP supports label distribution for all peers by default. Each node can send Label
Mapping messages to all peers without distinguishing upstream and downstream nodes. If an LSR
distributes labels only for upstream peers when it sends Label Mapping messages, the LSR checks the
upstream/downstream relationship of the session in routing information. An upstream node cannot send
Label Mapping messages to its downstream node along a route. If the route changes and the upstream/
downstream relationship is switched, the new downstream node resends Label Mapping messages. In this
process, the convergence is slow.
Figure 2-3 DU and DoD
Ingress Transit Egress
192.168.1.1/32
Allocate labels to 
the upstream node
Allocate labels to 
the upstream node
Request a label from 
the downstream node
Request a label from 
the downstream node
Send a label upon 
receiving the request
DU
DOD
Send a label upon 
receiving the request
 
Label Distribution Control Modes
The label distribution control mode refers to a method of label distribution on the LSR during
LSP establishment.
As described in Table 2-2, two label distribution control modes are available.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
27

Table 2-2 Label distribution control modes
Label Distribution
Control Modes
Definition Description
Independent mode A local LSR can distribute a
label bound to an FEC and
then inform the upstream
LSR, without waiting for the
label distributed by the
downstream LSR.
l As shown in Figure 2-3,
if the label advertisement
mode is DU and the label
distribution control mode
is Independent, a transit
LSR can assign a label to
the ingress node without
waiting for the label
assigned by the egress
node.
l As shown in Figure 2-3,
if the label advertisement
mode is DoD and the label
distribution control mode
is Independent, the
directly-connected
ingress transit node that
sends a Label Request
message replies with a
label without waiting for
the label assigned by the
egress node.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
28

