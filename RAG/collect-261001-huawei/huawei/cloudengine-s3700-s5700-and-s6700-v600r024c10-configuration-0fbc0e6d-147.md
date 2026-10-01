---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-147
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [21081, 21233]
sha256: 9f86081ec2c68d4cb96595ef00c0c9404f50631e6f4ce8301bec4f5b9b4f5416
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Context
                 An explicit path is a vector path comprised of a series of nodes that are arranged
                 in the configuration sequence. The IP address on an explicit path is the IP address
                 of an interface on a node. Generally, the IP address of the loopback interface on
                 the egress is used as the destination address of the explicit path.
                 Two adjacent nodes on an explicit path are connected in either of the following
                 modes:
                 ●      Strict: The two nodes are directly connected.
                 ●      Loose: Other nodes may exist between the two nodes.
                 The strict and loose modes can be used independently or in combination.
                 TE tunnels are classified into intra-area tunnels and inter-area tunnels. Here, areas
                 indicate OSPF and IS-IS areas, but not autonomous systems (ASs) running the
                 Border Gateway Protocol (BGP). OSPF areas are differentiated by area IDs,
                 whereas IS-IS areas are differentiated by levels.
                 ●      An intra-area TE tunnel's ingress and egress are located in the same area. An
                        intra-area tunnel can be established over a strict or loose explicit path.
                 ●      An inter-area TE tunnel traverses multiple areas. It can only be established
                        over an explicit path, with the area border router (ABR) and autonomous
                        system boundary router (ASBR) specified on the explicit path.
                 Configure an explicit path on the ingress of an MPLS TE tunnel and specify the
                 nodes that the tunnel must traverse or bypass.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Create an explicit path and enter the explicit path view.
                 explicit-path path-name

         Step 3 Configure an explicit path as required. For details, see Table 4-15.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              352
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


                 Table 4-15 Explicit path configuration

                  Operation                         Command                  Purpose

                  Specify a next hop for an         next hop ip-address      By default, the include
                  explicit path.                    [ include [ [ strict |   and strict parameters
                                                    loose ] | [ incoming |   are configured, meaning
                                                    outgoing ] ] * |         that an LSP must pass
                                                    exclude ]                through a specified node.
                                                                             The include parameter
                                                                             indicates that an LSP
                                                                             must pass through a
                                                                             specified node, whereas
                                                                             the exclude parameter
                                                                             indicates that an LSP
                                                                             cannot pass through a
                                                                             specified node.

                  Specify an intermediate           add hop ip-address1      By default, the include
                  for an explicit path.             [ include [ [ strict |   and strict parameters
                                                    loose ] | [ incoming |   are configured, meaning
                                                    outgoing ] ] * |         that an LSP must pass
                                                    exclude ] { after |      through a specified node.
                                                    before } ip-address2

                  Change the address of a           modify hop ip-address1   By default, the include
                  node on an explicit path.         ip-address2 [ include    and strict parameters
                                                    [ [ strict | loose ] |   are configured, meaning
                                                    [ incoming |             that an LSP must pass
                                                    outgoing ] ] * |         through a specified node.
                                                    exclude ]

                  Remove a node from an             delete hop ip-address    -
                  explicit path.



         Step 4 Return to the system view.
                 quit

         Step 5 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 6 Configure an explicit path for a tunnel.
                 mpls te path explicit-path path-name [ secondary ]

                 secondary indicates that an explicit path is configured for a backup tunnel.

                 ----End


Verifying the Configuration
                 ●      Run the list hop [ ip-address ] command to check node information of an
                        explicit path.



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           353
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration




4.18 Configuring a Hop Limit for an MPLS TE Tunnel
Prerequisites
                 Before configuring a hop limit for an MPLS TE tunnel, complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 Similar to the link administrative group and affinity attributes, a hop limit is a
                 condition for path selection during dynamic CR-LSP establishment. A hop limit
                 defines the maximum number of hops that a CR-LSP can traverse.

                 Set a hop limit on the ingress of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 3 Limit the number of hops on a CR-LSP of the tunnel.
                 mpls te hop-limit hop-limit-value [ best-effort | secondary ]

                 best-effort indicates that the hop limit is set for a best-effort path of the tunnel.
                 secondary indicates that a hop limit is set for a backup path of the tunnel.

                 ----End


4.19 Configuring an MPLS TE SRLG

4.19.1 Understanding MPLS TE SRLGs
                 An SRLG is mainly used in hot standby (HSB) and TE fast reroute (FRR)
                 networking scenarios to restrict path calculation for a backup tunnel. An SRLG
                 prevents the backup path from being established on a link with the same risk level
                 as the primary path of the tunnel, further enhancing the reliability of the TE
                 tunnel.

Context
                 Network administrators generally use CR-LSP hot standby or TE FRR to improve
                 MPLS TE tunnel reliability. However, in real-world situations, protection failures
                 often occur, as the following TE FRR example demonstrates.

                 On the network shown in Figure 4-30, the primary CR-LSP path is Path1. The link
                 P1 -> P2 needs to be protected using TE FRR. The bypass CR-LSP path is Path2.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             354

