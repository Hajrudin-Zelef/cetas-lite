---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-25
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [3373, 3526]
sha256: 605c3e9a467a61889ac6241c7d560497d4ba5e8353110fd36851625c65b520cb
---

# A line starting with the # sign is comments.

6. P1 deals with the received Resv message in the same process as that on P2. P1 updates the
Resv message and sends the message to PE1. Table 3-10 lists information carried in a Resv
message.
PE1 obtains the label allocated by P1 based on the received Resv message. Resource
reservation succeeds and a CR-LSP is set up.
Table 3-10 Resv message on P1
Object Value
SESSION Source: PE2-if0; Destination: PE1-if1
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
69

Object Value
RSVP_HOP P1-if0
LABEL 18
RECORD_ROUTE P1-if0; P2-if0; PE2-if0
 
3.2.5.3 Maintenance of Dynamic CR-LSPs
Path Status Maintenance
Soft State
Software state indicates that RSVP-TE periodically refreshes RSVP messages to maintain the
resource reservation state.
The resource reservation state can be classified into two types: path state and reservation state.
Path and Resv messages are created and refreshed periodically to maintain the two states
respectively. These messages are called RSVP Refresh messages. RSVP Refresh messages
contain PSB and RSB and are used for state synchronization on neighboring RSVP nodes. If a
node does not receive any Refresh message about PSB or RSB within a specified time period,
the node deletes the path or reservation state.
RSVP Refresh
RSVP messages are transmitted as IP datagrams; therefore the transmission is unreliable. After
a CR-LSP is established, each node along the established CR-LSP periodically sends RSVP
Refresh messages to its upstream and downstream nodes to synchronize states (including PSB
and RSB) of neighboring RSVP nodes.
NOTE
A Refresh message is not a new type of message. Refresh messages are the messages that have already
been advertised.
The refreshing interval is specified in the Time Value field.
If the PSB or RSB does not receive any Refresh message about a certain state block after the
keep-multiplier refreshing intervals elapses, it deletes the state. keep-multiplier specifies the
number of dropped successive RSVP Refresh message on a node. The default keep-multiplier
is 3.
Sending of Path and Resv messages between neighboring RSVP nodes is independent to each
other.
RSVP Srefresh
RSVP Refresh messages are used to synchronize path state block (PSB) and reservation state
block (RSB) information between nodes. They can also be used to monitor the reachability
between RSVP neighbors and maintain RSVP neighbor relationships. As the sizes of Path and
Resv messages are larger, sending many messages to establish many CR-LSPs causes increased
consumption of network resources. RSVP Srefresh can be used to address this problem.
RSVP Srefresh defines new objects based on the existing RSVP protocol:
l Message_ID extension and retransmission extension
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
70

According to the Message_ID extension mechanism defined in RFC 2961, RSVP messages
carry extended objects, including Message_ID and Message_ID_ACK objects. The two
objects are used to confirm RSVP messages and support reliable RSVP message delivery.
The Message_ID object can also be used to provide the RSVP retransmission mechanism.
For example, a node initializes a retransmission interval as Rf seconds after it sends an
RSVP message carrying the Message_ID object. If the node receives no ACK message
within Rf seconds, the node retransmits an RSVP message after (1 + Delta) x Rf seconds.
The Delta determines the increased rate of the transmission interval set by the sender. The
node keeps retransmitting the message until it receives an ACK message or the
retransmission times reach the threshold (called a retransmission increment value). By
default, Rf is set to 500 milliseconds, Delta is set to 1, and the retransmission times is set
to 3.
l Summary Refresh extension
The Summary Refresh extension supports Srefresh messages to update the RSVP status,
without the transmission of standard Path or Resv messages. The Srefresh extension builds
on the Message_ID extension.
Each Srefresh message carries a Message_ID object. Each object contains multiple
messages IDs, each of which identifies a Path or Resv state to be refreshed. If a CR-LSP
changes, its message ID value increases.
Only the state that was previously advertised by Path and Resv messages containing
Message_ID objects can be refreshed using the Srefresh extension.
After a node receives an Srefresh message, the node compares the Message_ID with that
saved in a local state block. If they match, the node does not change the state. If the
Message_ID is greater than that saved in the local state block, the node sends a NACK
message to the sender, refreshes the PSB or RSB based on the Path or Resv message, and
updates the Message_ID.
Fault Advertisement
RSVP-TE uses the following messages to advertise LSP errors.
l PathErr message: sent upstream by an RSVP node if an error occurs while this node is
processing a Path message. A PathErr message is forwarded by consecutive transit nodes
and arrives at the ingress.
l ResvErr message: sent downstream by an RSVP node if an error occurs while this node is
processing a Resv message. A ResvErr message is forwarded by consecutive transit nodes
and arrives at the egress.
Path Teardown
After a user instructs an ingress to delete a CR-LSP or the ingress receives a ResvErr message,
the ingress sends a PathTear message to a downstream node. The downstream node receives this
message, tears down the CR-LSP, and replies to the ingress with a ResvTear message.
The functions of PathTear and ResvTear messages are as follows:
l A PathTear message instructs a node to remove saved path information. The PathTear
message functions in the opposite way to a Path message.
l A ResvTear message instructs a node to remove resource reservation status. The ResvTear
message functions in the opposite way to a Resv message.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
71

3.2.5.4 RSVP-TE Message
RSVP-TE messages exchange information between nodes during the MPLS TE implementation
process.
RSVP Message Format
Each type of RSVP messages contains a common header. The length and types of other fields
are not fixed. Figure 3-17 shows the format of RSVP messages.
Figure 3-17 RSVP message format
Objects ( Variable )
Length Class_Number C-Type
Object Content (Variable)
Version Flags RSVP Checksum
Send_TTL RSVP Length
Message Type
Reserved
0 4 8 16 31
0 16 31 24
Format of RSVP messages
Format of Objects
Table 3-11 describes each field in the format.
Table 3-11 Description of fields in RSVP messages
Field Length Description
Version 4 bits Indicates the RSVP version number. Currently, the version
is 1.
Flags 4 bits Indicates the flag bit. Commonly, the value is 0. In RFC
2961, it is extended to identify whether Summary Refresh
Extension (Srefresh) is supported. If Srefresh is supported,
the value of the flag field is 0x01.
Message
Type
8 bits Indicates an RSVP message type. For example, 1 represents
the Path message and 2 represents the Resv message.
RSVP
Checksum
16 bits Indicates the RSVP checksum. The value 0 indicates that no
checksum was transmitted.
Send_TTL 8 bits Indicates the TTL of the message. When a node receives an
RSVP message, it compares the Send_TTL and the TTL in
the IP header to calculate the hops that the message passes
in a non-RSVP area.
Reserved 8 bits Indicates that the field is reserved.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
72

