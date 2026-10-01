---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-16
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [2008, 2104]
sha256: 5146e22e02c9bbb4b4ede08d80d8c829ad1f790ac1c80f3bed6a5ee74a6cdc60
---

# A line starting with the # sign is comments.

message digest is a unique result calculated by an irreversible character string conversion. If a
message is modified during transmission, a different digest is generated. After the message
arrives at the receiver, the receiver can determine whether the packet is modified by comparing
the received digest with the pre-calculated digest.
LDP MD5 authentication prevents LDP packets from being modified by generating a unique
digest for an information segment. This authentication is stricter than the common checksum
verification of TCP connections.
Before an LDP message is sent over a TCP connection, LDP MD5 authentication is performed
by padding the TCP header with a unique digest. This digest is a result calculated by MD5 based
on the TCP header, LDP session message, and password set by the user.
When receiving this TCP packet, the receiver obtains the TCP header, digest, and LDP session
message, and then uses MD5 to calculate a digest based on the received TCP header, received
LDP session message, and locally stored password. The receiver compares the calculated digest
with the received one to check whether the packet is modified.
A password can be set in either cipher text or plain text. The plain-text password is directly
recorded in the configuration file. The cipher-text password is recorded in the configuration file
after being encrypted using a special algorithm.
During the calculation of a digest, the manually entered character string is used regardless of
whether the password is in plain text or cipher text. This indicates that a password calculated
using an encryption algorithm does not participate in MD5 calculation, ensuring that LDP MD5
authentication implemented on Huawei devices is transparent to non-Huawei devices.
Keychain Authentication
Keychain, an enhanced encryption algorithm to MD5, calculates a message digest for the same
LDP message to prevent the message from being modified.
During keychain authentication, a group of passwords are defined to form a password string.
Each password is specified with encryption and decryption algorithms such as MD5 algorithm
and SHA-1, and is configured with the validity period. When sending or receiving a packet, the
system selects a valid password based on the user's configuration. Within the valid period of the
password, the system uses the encryption algorithm matching the password to encrypt the packet
before sending it out, or uses the decryption algorithm matching the password to decrypt the
packet before accepting it. In addition, the system automatically uses a new password after the
previous password expires, preventing the password from being decrypted.
The keychain authentication password, the encryption and decryption algorithms, and the
password validity period that construct a keychain configuration node are configured using
different commands. A keychain configuration node requires at least one password and
encryption and decryption algorithms.
LDP GTSM
Generalized TTL Security Mechanism (GTSM) is a mechanism that protects the service by
checking whether the TTL value in the IP header is within the pre-defined range. The
prerequisites for using GTSM are as follows:
l The TTL of normal packets between routers is determined.
l The TTL value of packets can hardly be modified.
LDP GTSM refers to GTSM implementation over LDP.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
40

To protect the router against attacks, GTSM checks the TTL in a packet to verify it. GTSM for
LDP is applied to LDP packets between neighbor or adjacent (based on a fixed number of hops)
routers. The TTL range is preset on each router for packets from other routers and GTSM is
enabled. If the TTL of an LDP packet received by a router configured with LDP is out of the
TTL range, the packet is considered invalid and is discarded. This protects the upper-layer
protocols.
2.2.11 LDP Extension for Inter-Area LSP
This feature enables LDP to establish inter-area LDP LSPs to provide tunnels that traverse the
public network.
Figure 2-10 Networking topology for LDP extension for inter-area LSP
IS-IS 
Area10
IS-IS 
Area20
LSRA
LSRB
LSRC
LSRD
Loopback0
1.1.0.1/32
Loopback0
1.3.0.1/32
Loopback0
1.3.0.2/32
Loopback0
1.2.0.1/32
As shown in Figure 2-10, there are two IGP areas: Area 10 and Area 20.
In the routing table of LSRD at the edge of Area 10, two host routes are reachable to LSRB and
LSRC. You can use IS-IS to aggregate the two routes to one route to 1.3.0.0/24 and send this
route to Area 20 to prevent a large number of routes from occupying too many resources on the
LSRD. Consequently, there is only one aggregated route (1.3.0.0/24) but not 32-bit host routes
in LSRA's routing table. By default, when establishing LSPs, LDP searches the routing table for
the route that exactly matches the FEC in the received Label Mapping message. Figure 2-10
shows routing entry information of LSRA and routing information carried in the FEC, as shown
in Table 2-4.
Enterprise Data Communication Products
Feature Description - MPLS 2 MPLS LDP
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
41

Table 2-4 Routing entry information of LSRA and routing information carried in the FEC
Routing Entry Information
of LSRA
FEC
1.3.0.0/24 1.3.0.1/32
1.3.0.2/32
 
