---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-11
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [851, 1004]
sha256: b3a07cac8721fd8cbe4f991a85d02b006e6e31754e229801e3345e83ef002a40
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

LSRs
                 A label switching router (LSR), also called an MPLS node, swaps labels and
                 forwards MPLS packets. As the fundamental elements of an MPLS network, all
                 LSRs must support MPLS.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            10
MPLS Configuration
MPLS Configuration                                                          2 Basic MPLS Configuration


LERs
                 An LSR that resides on the edge of an MPLS domain is called a label edge router
                 (LER). When an LSR connects to a node that does not run MPLS, the LSR functions
                 as an LER.
                 An ingress LER classifies the packets that enter an MPLS domain into FECs, pushes
                 labels into them, and then forwards them accordingly. An egress LER pops the
                 labels from the packets that leave an MPLS domain, and then forwards them
                 based on the original packet types (the type before labels are added).

LSPs
                 The path through which the packets of a FEC pass on an MPLS network is called a
                 label switched path (LSP).
                 An LSP is a unidirectional path from the ingress to the egress.

Ingress, Transit, and Egress LSRs
                 The LSRs on an LSP are categorized as follows:
                 ●      Ingress: the start node on an LSP. An LSP can have only one ingress.
                        An ingress pushes a label into an IP packet to encapsulate the IP packet into
                        an MPLS packet for forwarding.
                 ●      Transit node: an intermediate node on an LSP. An LSP may have multiple
                        transit LSRs.
                        A transit node searches the LFIB and swaps labels to implement MPLS packet
                        forwarding.
                 ●      Egress: the end node on an LSP. An LSP can have only one egress.
                        An egress pops the label from an MPLS packet and restores the original
                        packet before forwarding it.
                 Ingress and egress nodes function as both LSRs and LERs. Transit nodes functions
                 LSRs.

Upstream and Downstream
                 LSRs include upstream and downstream LSRs:
                 ●      Upstream: All LSRs that send MPLS packets to the local LSR are called
                        upstream LSRs.
                 ●      Downstream: All LSRs that receive MPLS packets from the local LSR are called
                        downstream LSRs.
                 In Figure 2-6, for the data flow destined for 192.168.1.0/24, LSR-A is the upstream
                 node of LSR-B, and LSR-B is the downstream node of LSR-A. Similarly, LSR-B is the
                 upstream node of LSR-C, and LSR-C is the downstream node of LSR-B.

                 Figure 2-6 Upstream and downstream




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            11
MPLS Configuration
MPLS Configuration                                                         2 Basic MPLS Configuration


Label Distribution
                 Packets with the same destination address are part of the same FEC, which is
                 allocated a label in the label resource pool. An LSR records the mapping between
                 a label and an FEC, encapsulates the mapping into a message, and notifies the
                 upstream LSR of the mapping. This process is called label distribution.

                 Figure 2-7 Label distribution




                 On the network shown in Figure 2-7, packets with the destination address
                 192.168.1.0/24 are part of the same FEC. LSR-B and LSR-C allocate a label for the
                 FEC and notify the upstream nodes of the mapping between the FEC and label.
                 Labels are allocated by downstream devices.

Label Distribution Protocol
                 Label distribution protocols, also called signaling protocols, are MPLS control
                 protocols used to identify FECs, distribute labels, and create and maintain LSPs.
                 MPLS supports multiple label distribution protocols, such as Label Distribution
                 Protocol (LDP), Resource Reservation Protocol Traffic Engineering (RSVP-TE), and
                 Multiprotocol Extensions for Border Gateway Protocol (MP-BGP).

MPLS Architecture
                 The MPLS architecture consists of a control plane and a forwarding plane.
                 Figure 2-8 shows the MPLS architecture.

                 Figure 2-8 MPLS architecture




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                               12
MPLS Configuration
MPLS Configuration                                                           2 Basic MPLS Configuration



                 ●      The control plane uses IP routes to transmit packets. It is mainly responsible
                        for allocating labels, creating label forwarding tables, and establishing and
                        tearing down LSPs.
                 ●      The forwarding plane, also called the data plane, does not use IP routes.
                        Instead, it uses a Layer 2 network, such as an Ethernet network, to transmit
                        packets. The forwarding plane is mainly responsible for adding and removing
                        labels and forwarding packets based on the LFIB.

2.2.2 LSP Establishment
LSP Establishment Process
                 MPLS needs to distribute labels to packets and establish LSPs before it can
                 forward packets.
                 Labels are assigned and distributed by a downstream LSR to an upstream LSR. A
                 downstream LSR divides packets into different FECs based on the IP routing table,
                 and assigns a label to each FEC. It then notifies the upstream LSR of the mapping
                 between FECs and labels to establish an LFIB and LSPs. Figure 2-9 shows the LSP
                 establishment process.

                 Figure 2-9 LSP establishment




                 LSPs can be either static or dynamic. Static LSPs are manually configured by the
                 administrator, and dynamic LSPs are dynamically established using label
                 distribution protocols based on routing information.

Dynamic LSP Establishment
                 Dynamic LSPs are dynamically established using label distribution protocols. MPLS
                 supports multiple label distribution protocols:
                 ●      LDP
                        LDP is dedicated for label distribution. When LDP sets up an LSP in hop-by-
                        hop mode, LDP determines next hops based on the forwarding information
                        base (FIB) on each LSR. Information contained in a FIB is generally collected
                        using routing protocols, such as Interior Gateway Protocol (IGP) and BGP. LDP
                        utilizes routing information, but is not dependent on any routing protocol.
                        LDP is not the only label distribution protocol. BGP and RSVP can also be
                        extended to distribute MPLS labels.
                 ●      MP-BGP
                        MP-BGP is an extension of BGP. MP-BGP introduces the community attribute
                        and allocates labels to private network routes and inter-AS VPN labeled
                        routes of MPLS VPN services.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                13
MPLS Configuration
MPLS Configuration                                                            2 Basic MPLS Configuration


2.2.3 MPLS Forwarding

