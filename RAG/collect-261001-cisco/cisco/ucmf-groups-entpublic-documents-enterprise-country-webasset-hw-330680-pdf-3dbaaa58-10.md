---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-10
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["advisory", "copyright", "distribution"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [1184, 1309]
sha256: 9f4cfeacc18e3b3fb2f0760fbf4ac3914d634844917a34c52b8c077024bd08d7
---

# A line starting with the # sign is comments.

2.1 Introduction to MPLS LDP
Definition
The Label Distribution Protocol (LDP) is a control protocol of Multiprotocol Label Switching
(MPLS), which functions similarly to a signaling protocol on a traditional network. It classifies
FECs, distributes labels, and establishes and maintains LSPs. LDP defines messages in the label
distribution process as well as procedures for processing these messages.
Purpose
MPLS supports multiple labels and its forwarding plane is connection-oriented, and thus this
excellent scalability enables the MPLS/IP-based network to provide various services. Through
LDP, Label Switching Routers (LSRs) directly map routing information at the network layer to
the switched paths at the data link layer, and thus establish LSPs at the network layer.
Currently, LDP is widely used to provide VPN services because it features simple networking
and configurations, supports route-based establishment of LSPs, and supports high-capacity
LSPs.
2.2 Principles
2.2.1 Basic Concepts
LDP Adjacency
When an LSR receives a Hello message from a peer, an LDP peer may exist. An LDP adjacency
can be created to maintain the presence of the peer. There are two types of LDP adjacencies:
l Local adjacency: The adjacency is discovered by exchanging Link Hello messages.
l Remote adjacency: The adjacency is discovered by exchanging Target Hello messages.
LDP Peers
LDP peers refer to two LSRs that use LDP to set up an LDP session and then exchange label
messages.
LDP peers learn labels from each other using the LDP session between them.
LDP Session
LSRs in an LDP session exchange messages such as label mapping messages and label release
messages. LDP sessions are classified into the following types:
l Local LDP session: The LDP session is set up between local adjacencies. The two LSRs
setting up the local LDP session are directly connected.
l Remote LDP session: The LDP session is set up between remote adjacencies. The two
LSRs setting up the remote LDP session can be either directly or indirectly connected.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
22

NOTE
LDP maintains the presence of peers using adjacencies. The type of peers depends on the type of
adjacencies. A pair of peers can be maintained by multiple adjacencies. If a pair of peers is maintained by
both local and remote adjacencies, the peers support coexistence of the local and remote adjacencies. An
LDP session can only be established if such pairs of peers exist.
A local and a remote LDP session can be set up simultaneously.
The principle is that the local and remote LDP adjacencies can be connected to the same peer
so that the peer is maintained by both the local and remote LDP adjacencies.
As shown in Figure 2-1, when the local LDP adjacency is deleted due to a failure on the link to
which the adjacency is connected, the peer's type may change without affecting its presence or
status. (The peer type is determined by the adjacency type. The types of adjacencies include
local, remote, and coexistent local and remote.)
If the link becomes faulty or is recovering from a fault, the peer type may change while the type
of the session associated with the peer changes accordingly. However, the session is not deleted
and does not become Down. Instead, the session remains Up.
Figure 2-1 Networking diagram for a coexistent local and remote LDP session
Remote Adjacency
CE1 CE2 PE1 PE2Local
 Adjacency
P
 
A coexistent local and remote LDP session is typically applied to L2VPN. As shown in Figure
2-1, L2VPN services are transmitted between PE1 and PE2. When the directly-connected link
between PE1 and PE2 recovers after being disconnected, the processing is as follows:
1. MPLS LDP is enabled on the directly-connected PE1 and PE2, and a local LDP session is
set up between PE1 and PE2. PE1 and PE2 are configured as the remote peer of each other,
and a remote LDP session is set up between PE1 and PE2. Local and remote adjacencies
are then set up between PE1 and PE2. Since now, both local and remote LDP sessions exist
between PE1 and PE2. L2VPN signaling messages are transmitted through the compatible
local and remote LDP session.
2. When the physical link between PE1 and PE2 becomes Down, the local LDP adjacency
also goes Down. The route between PE1 and PE2 is still reachable through the P, indicating
that the remote LDP adjacency remains Up. The session changes to a remote session so
that it can remain Up. The L2VPN does not detect the change in session status and therefore
does not delete the session. This prevents the L2VPN from having to disconnect and recover
services, and shortens service interruption time.
3. When the fault is rectified, the link between PE1 and PE2 as well as the local LDP adjacency
can go Up again. The session changes to the compatible local and remote LDP session and
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
23

remains Up. Again, the L2VPN will not detect the change in session status and therefore
does not delete the session. This shortens service interruption time.
Type of LDP Messages
LDP messages are classified into the following types:
l Discovery message: used to notify and maintain the existence of an LSR on a network.
l Session message: used to establish, maintain, and terminate sessions between LDP peers.
l Advertisement message: used to create, modify, and delete label mappings for FECs.
l Notification message: used to provide advisory and error information.
To ensure the reliability of message transmission, LDP uses the TCP transport for Session,
Advertisement, and Notification messages. LDP uses the UDP transport only for transmitting
the Discovery message.
Label space
A label space is a range of labels allocated between LDP peers, which can be categorized as
follows:
l Per-platform label space: An entire LSR uses one label space. Currently, per-platform label
space is mostly used.
l Per-interface label space: Each interface of an LSR is assigned a label space.
LDP identifier
An LDP identifier identifies the label space used by a specified LSR. An LDP identifier is 6
bytes in the format <LSR ID>:<Label space ID>.
l LSR ID: indicates the 4-byte LSR identifier.
l Label space ID: indicates the 2-byte label space identifier. The value 0 indicates the per-
platform label space, while the value non-0 indicates the per-interface label space.
For example, the LDP ID is 192.168.1.1:0, indicating that the LSR ID is 192.168.1.1 and per-
platform label space is used.
2.2.2 LDP Working Mechanism
LDP defines the label distribution process and messages transmitted during label distribution.
An LSR can use LDP to map routing information on the network layer to on the data link layer,
setting up an LSP. LDP working process goes through the following phases:
1. After discovering a neighbor, an LSR sets up an LDP session.
2. After the session is established, LDP notifies LDP adjacencies of the mappings between
FECs and labels and sets up an LSP.
RFC 5036 defines the label advertisement mode, label distribution control mode, and label
retention mode to determine how the LSR advertises and manages labels.
LDP Session
LDP Discovery Mechanisms
LDP discovery mechanisms are used by LSRs to discover potential LDP peers. LDP discovery
mechanisms are classified into the following types:
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
24

