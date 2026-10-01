---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-28
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [3895, 4033]
sha256: f42c1d60a60f0136f4732bce3f3a3ee4818d5236544c168e1fa4e12796478c5b
---

# A line starting with the # sign is comments.

– When no CR-LSP exists between RSVP neighbors, the RSVP adjacency remains until
the RSVP authentication lifetime expires. The configuration of the RSVP authentication
time does not affect the status of existing CR-LSPs.
– This function can avoid continuous RSVP authentication. For example, when RSVP
authentication is enabled between RTA and RTB, but the key is damaged because the
RSVP messages sent from RTA to RTB are incorrect, RTB receives and discards the
messages. This can cause RTA to continuously send RTB the faulty RSVP messages
and RTB to continuously discard these RSVP messages. The authentication relationship
between the neighbors, however, cannot be torn down. In this case, the authentication
lifetime needs to be configured. When a neighbor is able to receive a valid RSVP
message within the lifetime, the RSVP authentication lifetime resets. Otherwise, the
authentication relationship between RSVP neighbors is deleted after the authentication
lifetime expires.
l Handshake mechanism
The handshake mechanism maintains the RSVP authentication status. After RSVP
neighboring nodes authenticate each other, they exchanged handshake packets. If they
accept the packets, they record a successful handshake. If a local node receives a packet
with the sequence number less than the local maximum sequence number, the local node
processes the packet as follows:
– Discards the packet if the packet shows that the handshake mechanism is not enabled
on the remote node.
– Discards the packet if the packet shows that the handshake mechanism is enabled on
the remote node and the local node has a record about a successful handshake. If the
local node does not have a record about a successful handshake, this packet is the first
one arrives at the local node and the local node starts a handshake process.
l Message window
A message window saves sequence numbers of received RSVP messages. The number of
sequence numbers that can be saved ranges from 1 to 64. When the window size is 1, only
the largest sequence number of RSVP messages from neighbors can be saved. When the
window size is not 1, multiple sequence numbers of RSVP messages from neighbors can
be saved. For example, a window size is set to 10, and the largest sequence number of a
received RSVP message is 80. The sequence numbers between 71 and 80 can be saved if
there is no packet mis-sequence. If a packet mis-sequence problem occurs, the local node
arranges the messages and records the 10 largest sequence numbers.
NOTE
By default, the window size is 1. Packet processing in the handshake mechanism is based on the
prerequisite that the the window size is 1. A non-1 window-size affects packet processing in the
handshake mechanism.
RSVP Key Management Modes
RSVP keys can be managed in either of the following modes:
l MD5 key
An MD5 key is entered in either ciphertext or plaintext on an RSVP interface or node. An
MD5 key has the following characteristics:
– A key cannot be shared. Each protocol is configured with a separate key.
– An interface or a node is assigned only one key. The key can be reconfigured but cannot
be changed.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
80

l Keychain key
Keychain is an enhanced encryption algorithm. A group of passwords are defined in the
format of a password string during keychain authentication, and each password is assigned
a specified encryption and decryption algorithm and configured with a validity period.
When the system sends or receives a packet, the system selects a valid password. Within
the validity period of the password, the system uses the encryption algorithm matching the
password to encrypt the packet before sending it out, or uses the decryption algorithm
matching the password to decrypt the packet before accepting it. In addition, the system
automatically uses a new password after the previous password expires, minimizing
password decryption risks.
Keychain management has the following characteristics:
– A keychain authentication password and the encryption and decryption algorithms must
be configured. A password validity period can also be configured.
– Keychain settings can be shared by separate protocols and features and can be managed
uniformly.
Keychain can be used on an RSVP interface and node and support HMAC-MD5.
Leveled RSVP Authentication
Leveled RSVP authentication is supported
l Neighbor-oriented authentication
You can configure authentication information, such as authentication keys, based on
different neighbor addresses. RSVP then authenticates each neighbor separately.
The following configuration options available:
– The IP address of an interface on an RSVP neighboring node as an RSVP neighbor
address.
– The LSR ID of an RSVP neighboring node is used as an RSVP neighbor address.
l Interface-oriented authentication
Authentication is configured on interfaces, and RSVP authenticates messages based on
inbound interfaces.
Neighbor-oriented authentication has a higher priority than interface-oriented authentication. A
node discards messages if neighbor-oriented authentication fails and performs interface-oriented
authentication only if neighbor-oriented authentication is not enabled.
3.2.9 MPLS TE Reliability
3.2.9.1 Reliability Overview
MPLS TE reliability techniques need to prevent or minimize packet loss that occurs in one of
the following situations:
l If attributes, such as bandwidth, are modified when an MPLS TE tunnel is transmitting
services, the tunnel is reestablished using new attributes, and services switch to the new
path.
l If a node or link fails while an MPLS TE tunnel is transmitting services, a backup CR-LSP
is established and takes over traffic.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
81

l An MPLS TE tunnel has been established and is transmitting services. A fault occurs in
the control plane but not the forwarding plane of a node along the path. Traffic needs to be
forwarded uninterruptedly before the control plane recovers.
MPLS TE tunnels that transmit mission-critical services require high reliability. Table 3-14 lists
MPLS TE reliability functions.
Table 3-14 MPLS TE reliability functions
Technique
Classification
Description Functions
Reliability
mechanism for
updating MPLS
TE attributes
Ensures reliable traffic transmission after attributes
are updated and a new CR-LSP is established using
the updated attribute and takes over traffic.
l Make-
Before-
Break
Fault detection Rapidly detects MPLS TE network faults to speed up
a protection switchover.
l RSVP Hello
l BFD for
MPLS TE
Traffic
protection
Supports network-level reliability, including E2E
path protection and local protection.
l CR-LSP
Backup
l TE FRR
l SRLG
l TE Tunnel
Protection
Group
Supports device-level reliability, including
uninterrupted traffic transmission on the forwarding
plane while a fault occurs on the control plane of a
node.
l RSVP GR
 
