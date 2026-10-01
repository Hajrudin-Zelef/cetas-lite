---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-9
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["asic", "copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [573, 718]
sha256: 60dc436d1f952f0809a5d239a91108e6650fdb62115d7d758506077532b4ffbc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 2.1 Overview of MPLS
                 2.2 Understanding MPLS
                 2.3 Configuration Precautions for MPLS Basics
                 2.4 Configuring the MPLS MTU
                 2.5 Configuring MPLS TTL Processing Modes
                 2.6 Configuring Alarm Thresholds for MPLS Resources


2.1 Overview of MPLS
Definition
                 Multiprotocol Label Switching (MPLS) operates between the data link layer and
                 network layer in the TCP/IP protocol stack; it provides the IP layer with the
                 connectivity service and obtains services from the link layer. MPLS uses label
                 switching in place of IP forwarding. A label is a short, fixed-length connection
                 identifier, which is locally significant. It is similar to a data link connection
                 identifier (DLCI) in Frame Relay and a virtual path identifier (VPI)/virtual channel
                 identifier (VCI) in Asynchronous Transfer Mode (ATM). A label is encapsulated
                 between the data link layer and network layer.

Purpose
                 Asynchronous Transfer Mode (ATM) uses fixed-length labels (cells) and maintains
                 a label table much smaller than a routing table, enabling it to provide much
                 higher forwarding performance than IP routing. However, ATM is a complex
                 protocol and its deployment costs are high. This has hindered its widespread
                 popularity and growth. The traditional IP technology is simple, and its deployment
                 costs are low. Against this backdrop, MPLS was developed to combine the
                 advantages of IP and ATM.
                 MPLS was originally proposed to increase the forwarding speed of devices. It does
                 this by analyzing packet headers only on the edges of a network, not at each hop.
                 This means it requires less processing time than IP.

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                                  5
MPLS Configuration
MPLS Configuration                                                         2 Basic MPLS Configuration


                 With the emergence of ASIC, the route lookup speed is no longer a bottleneck for
                 network development. Consequently, MPLS lost its advantage in accelerating
                 forwarding. However, as MPLS supports multi-layer labels and its forwarding plane
                 is connection-oriented, MPLS is still widely used in some scenarios such as virtual
                 private network (VPN), traffic engineering (TE), and quality of service (QoS).

Benefits
                 MPLS is independent of data link layer protocols; that is, it can use any Layer 2
                 media to transmit packets.
                 MPLS is derived from the Internet Protocol version 4 (IPv4), but its core
                 technologies can be extended to support multiple network protocols, such as the
                 Internet Protocol version 6 (IPv6), Internet Packet Exchange (IPX), Appletalk,
                 DECnet, and Connectionless Network Protocol (CLNP). In MPLS, "multiprotocol"
                 refers to the fact that the protocol supports multiple network protocols.
                 MPLS is a tunneling technology, not a service or application. This technology
                 supports multiple upper-layer protocols and services and improves data
                 transmission security.


2.2 Understanding MPLS

2.2.1 Basic Concepts of MPLS
MPLS Network Structure
                 Figure 2-1 shows the typical MPLS network structure: Label switching routers
                 (LSRs) are the basic elements, and multiple LSRs form an MPLS domain. An LSR
                 that resides at the edge of an MPLS domain and connects to a non-MPLS network
                 is called a label edge router (LER). An LSR that resides inside an MPLS domain is
                 called a core LSR. Specifically, if an LSR connects to one or more MPLS-incapable
                 nodes, the LSR is an LER. If all adjacent nodes of an LSR run MPLS, the LSR is a
                 core LSR.

                 Figure 2-1 MPLS network structure




                 MPLS forwards packets based on labels. Specifically, when an IP packet enters an
                 MPLS network, the ingress LER analyzes the packet content and adds an

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                               6
MPLS Configuration
MPLS Configuration                                                           2 Basic MPLS Configuration


                 appropriate label to the IP packet. All nodes on the MPLS network forward the
                 packet based on labels. When the IP packet leaves the MPLS network, the egress
                 pops (removes) the label in the packet.

                 The path through which IP packets are transmitted on an MPLS network is called
                 a label switched path (LSP). An LSP is a unidirectional path that is in the same
                 direction as the data flow.

                 Figure 2-2 MPLS LSP




                 On an LSP, the start node is referred to as the ingress, intermediate nodes as
                 transit nodes, and the end node as the egress. An LSP has one ingress, one egress,
                 and zero, one, or multiple transit nodes.

Forwarding Equivalence Class
                 A forwarding equivalence class (FEC) refers to a group of data flows forwarded in
                 the same manner. Data flows in the same FEC are processed by LSRs in the same
                 way.

                 FECs can be classified based on addresses, service types, or QoS policies, among
                 others. For example, in IP forwarding, packets matching the same route based on
                 the longest match rule belong to an FEC.

Label
                 A label is a short, fixed-length identifier that has local significance. It is used to
                 uniquely identify the FEC to which a packet belongs. In some cases, for example,
                 when load balancing is required, one FEC may have multiple incoming labels.
                 However, a label can represent only one FEC on the same device.

                 Figure 2-3 illustrates the structure of an MPLS header.

                 Figure 2-3 Structure of an MPLS label header




                 An MPLS label header contains the following fields:

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                                   7
MPLS Configuration
MPLS Configuration                                                               2 Basic MPLS Configuration


                 ●      Label: a 20-bit field that identifies a label value.
                 ●      Exp: a 3-bit field used for extensions. This field is generally used in the class of
                        service (CoS) to provide a similar function as Ethernet 802.1p.
                 ●      S: a 1-bit field that identifies the bottom of a label stack. MPLS supports
                        multiple labels for label nesting. If the S field of a label is set to 1, the label is
                        at the bottom of the label stack.
                 ●      TTL: an 8-bit field indicating a time to live (TTL) value. This field has the
                        same meaning as the TTL field in IP packets.
                 Labels are encapsulated between the data link layer and network layer, and are
                 supported by all data link layer protocols. Figure 2-4 shows the position of a label
                 in a packet.

                 Figure 2-4 Position of a label in a packet




