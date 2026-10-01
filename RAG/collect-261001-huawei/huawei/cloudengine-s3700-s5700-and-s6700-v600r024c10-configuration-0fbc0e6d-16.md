---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-16
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1596, 1732]
sha256: f235c086af3508ca52c637908d0dfc53810804bd80550d2c87cac980454fe6a5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 3.1 Overview of MPLS LDP
                 3.2 Understanding MPLS LDP
                 3.3 Configuration Precautions for MPLS LDP
                 3.4 Default Settings for MPLS LDP
                 3.5 Configuring Static LSPs
                 3.6 Configuring LDP Sessions
                 3.7 Adjusting LDP Sessions
                 3.8 Configuring the Dynamic LDP Advertisement Capability
                 On devices enabled with global LDP, the dynamic LDP advertisement capability
                 allows extended LDP functions to be dynamically enabled or disabled when the
                 LDP session is working properly, ensuring stable LSP operation.
                 3.9 Configuring LDP LSPs
                 3.10 Configuring LDP Label Advertisement and Management Modes
                 3.11 Configuring IGP-based LDP Automatic Deployment
                 IGP-based automatic LDP deployment reduces the configuration workload and
                 ensures configuration correctness.
                 3.12 Disabling a Device from Forwarding Unknown TLVs
                 3.13 Configuring LDP Extension for Inter-Area LSPs
                 3.14 Configuring LDP Auto FRR
                 3.15 Configuring Static BFD for LDP LSP
                 3.16 Configuring Dynamic BFD for LDP LSP
                 3.17 Configuring LDP Session Protection
                 3.18 Configuring LDP-IGP Synchronization
                 3.19 Configuring the LDP GR Helper
                 3.20 Configuring the Uniform/Pipe Mode for the MPLS Penultimate Hop

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                          26
MPLS Configuration
MPLS Configuration                                                          3 MPLS LDP Configuration


                 3.21 Disabling LDP Session Flapping Suppression
                 3.22 Disabling LDP Interface Flapping Suppression
                 3.23 Configuring LDP Security Features
                 3.24 Checking MPLS Network Connectivity Using Ping/Tracert
                 3.25 Maintaining MPLS LDP
                 3.26 Troubleshooting MPLS LDP


3.1 Overview of MPLS LDP
Definition
                 The Label Distribution Protocol (LDP) is a Multiprotocol Label Switching (MPLS)
                 control protocol that functions similarly to a signaling protocol on a traditional
                 network. LDP classifies packets into forwarding equivalence classes (FECs),
                 distributes labels, and establishes and maintains label switched paths (LSPs). LDP
                 defines messages in the label distribution process as well as procedures for
                 processing these messages.

Purpose
                 On an MPLS network, LDP distributes label mappings and establishes LSPs based
                 on routing information. Data packets can then be transmitted along LSPs over the
                 MPLS network. Just like with most routing protocols, LDP uses multicast Hello
                 messages to automatically discover local neighbors and establish local
                 adjacencies, or uses unicast Hello messages to discover remote neighbors and
                 establish remote adjacencies. LDP establishes a TCP connection between neighbors
                 and negotiates parameters to establish a session between them. They can then
                 exchange messages over the LDP session to set up an LSP. LDP networking is
                 simple to construct and configure. LDP establishes LSPs using routing information
                 and supports a large number of LSPs. The main LDP applications are as follows:
                 ●      LDP establishes full-mesh LSPs to transmit IP data over the MPLS network,
                        implementing a core network free without Border Gateway Protocol (BGP).
                 ●      LDP establishes end-to-end intra-domain tunnels to transmit Layer 3 virtual
                        private network (L3VPN) services.

                 Figure 3-1 LDP networking




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              27
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration




3.2 Understanding MPLS LDP

3.2.1 Basic LDP Concepts
                 MPLS supports multiple label distribution protocols, among which LDP is widely
                 used.
                 LDP defines messages to be used in the label distribution process and procedures
                 for processing the messages. Label switching routers (LSRs) obtain information
                 about incoming labels, next-hop nodes, and outgoing labels for specified FECs
                 from local forwarding information bases (LFIBs) to establish LSPs.
                 For detailed information about LDP, see relevant standards (LDP Specification).

LDP Adjacencies
                 When an LSR receives a Hello message from a peer, the LSR establishes an LDP
                 adjacency with the peer. There are two types of LDP adjacencies:
                 ●      Local adjacency: established by exchanging Link Hello messages.
                 ●      Remote adjacency: established by exchanging Targeted Hello messages.

LDP Peers
                 LDP peers refer to LSRs that establish an LDP session and use LDP to exchange
                 label messages.
                 LDP peers learn each other's labels through the LDP session established between
                 them.

LDP Sessions
                 An LDP session is used by LSRs to exchange messages, such as Label Mapping
                 messages and Label Release messages. LDP sessions are classified into the
                 following types:
                 ●      Local LDP session: created over a local adjacency. The two LSRs in a local
                        session are directly connected. LSRs can exchange labels and establish LDP
                        LSPs only after local LDP sessions are established.
                 ●      Remote LDP session: created over a remote adjacency. The two LSRs of a
                        remote session can be directly or indirectly connected. A remote LDP session
                        provides the following functions: An L2VPN needs to use LDP extensions to
                        transmit its own protocol packets. When devices at both ends of an L2VPN
                        are indirectly connected, a remote LDP session needs to be established. Also,
                        LDP-based remote peers directly distribute LDP labels to each other, which
                        can be used in scenarios similar to LDP over TE.
                 Local and remote LDP sessions can be established at the same time.

Differences and Relationships Between LDP Adjacencies, Peers, and Sessions
                 Differences

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             28
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


                 The differences between LDP adjacencies, peers, and sessions are as follows:

                 ●      An LDP adjacency is a TCP connection established after two devices exchange
                        Hello messages. It is used for the link between two interconnected interfaces.
                 ●      LDP peers refer to devices that exchange label messages through LDP after a
                        TCP connection is set up.
                 ●      An LDP session refers to a series of processes in which two LDP peers
                        exchange label messages.

                 Relationships

