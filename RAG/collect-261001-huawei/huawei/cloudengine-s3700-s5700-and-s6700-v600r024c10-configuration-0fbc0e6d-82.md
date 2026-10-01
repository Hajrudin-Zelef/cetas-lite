---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-82
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11550, 11637]
sha256: 24fa74bf51c4d6b49a875715808e86ef4ef83e99f06e8442c29eb2f49f7331a4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

CR-LSP
                 The principle of traffic forwarding in a CR-LSP is identical to that of a common
                 LSP (such as an LDP LSP), with both types utilizing label switching rather than IP
                 forwarding for data transport. However, unlike a common LSP, a CR-LSP
                 established by MPLS TE is required to adhere to specific tunnel attributes, making
                 it a constraint-based LSP. Through the establishment of CR-LSPs, MPLS TE
                 facilitates path planning and optimization for network traffic.

                 CR-LSPs can be established either statically or dynamically, leading to the
                 categorization into static CR-LSPs and dynamic CR-LSPs.
                 ●      Static CR-LSP: For a static CR-LSP, label forwarding and resource reservation
                        are configured manually, without the involvement of any signaling protocol.
                        Establishing a static CR-LSP requires minimal resources since it does not
                        involve the exchange of MPLS control packets between the two ends of the
                        CR-LSP. However, static CR-LSPs lack the flexibility to adjust dynamically to
                        changes in the network topology. Consequently, they are typically suited for
                        smaller networks with simple topologies.
                 ●      Dynamic CR-LSP: For a dynamic CR-LSP, resources and labels are
                        automatically reserved and allocated using Resource Reservation Protocol-
                        Traffic Engineering (RSVP-TE). Manual configuration for each hop along a
                        dynamic CR-LSP is unnecessary. The dynamic setup of CR-LSPs is particularly
                        suited for large-scale networks.


MPLS TE Tunnel
                 MPLS TE generally associates multiple CR-LSPs with a virtual tunnel interface to
                 form an MPLS TE tunnel. An MPLS TE tunnel involves the following terms:

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               196
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 ●      Tunnel interface: a point-to-point (P2P) virtual interface designed to
                        encapsulate packets. Like a loopback interface, a tunnel interface serves as a
                        logical interface.
                 ●      Tunnel ID: a decimal number that uniquely identifies an MPLS TE tunnel,
                        simplifying the planning and management of MPLS TE tunnels. A tunnel ID
                        must be specified when an MPLS TE tunnel interface is configured.
                 ●      LSP ID: a decimal number that uniquely identifies an LSP, facilitating LSP
                        planning and management.
                 On the network shown in Figure 4-1, there are two CR-LSPs: one primary and one
                 backup. The primary CR-LSP is numbered 2 and established over the path LSR1 ->
                 LSR2 -> LSR3 -> LSR4 ->LSR5. The backup CR-LSP is numbered 1024 and
                 established over the path LSR1 -> LSR6 -> LSR7 -> LSR8 -> LSR5. Both CR-LSPs are
                 part of the same MPLS TE tunnel whose tunnel interface is Tunnel1 and ID is 100.

                 Figure 4-1 Relationship between an MPLS TE tunnel and CR-LSPs




Link Attribute
                 MPLS TE link attributes identify the bandwidth allocation, route cost, and link
                 reliability of a physical link. MPLS TE link attributes include:
                 ●      Total link bandwidth: total bandwidth of a physical link.
                 ●      Maximum reservable bandwidth: maximum bandwidth of a link that can be
                        reserved for an MPLS TE tunnel. The maximum reservable bandwidth of a link
                        must be less than or equal to the total bandwidth of the link.
                 ●      TE metric of a link: used in TE tunnel path calculation, allowing the
                        calculation process to be independent from IGP route-based path calculation.
                        By default, a link uses the IGP metric as the TE metric.
                 ●      Link administrative group (also called link color): represented by a 32-bit
                        vector. Each bit in this vector can be assigned a specific meaning or left
                        unassociated, depending on the configuration, such as the link bandwidth or
                        performance. A link administrative group can also be used for link
                        management. For example, it can determine whether a particular link is used
                        by an MPLS TE tunnel and whether it carries multicast services.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            197
MPLS Configuration
MPLS Configuration                                                                  4 MPLS TE Configuration


                        The link administrative group attribute needs to be utilized alongside the
                        affinity attribute to manage path selection. For details about link
                        administrative groups, see Link Administrative Group.
                 ●      Shared risk link group (SRLG): a group of links that share a public physical
                        resource, such as an optical fiber. Links within an SRLG share the same
                        vulnerability to faults. Specifically, if one link fails, the other links in the SRLG
                        fail as well.
                        SRLGs are mainly used in CR-LSP hot standby and TE FRR scenarios to
                        enhance TE tunnel reliability. For details about SRLGs, see 4.19.1
                        Understanding MPLS TE SRLGs.


