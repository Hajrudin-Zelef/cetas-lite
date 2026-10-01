---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-188
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [27333, 27484]
sha256: 2c209a87290939fe8d89bdc61788a7f20edf9541024c8d4d7b2f38e80cb47207
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        is met. PE1 also re-determines whether this label block meets the
                        requirement: LOn ≤ m < LOn + LRn. Because LOn is 0, m is 1, and LRn is 4,
                        this requirement is met.
                        Similarly, PE3 receives the two label blocks (LB/LR/LO = 1000/5/0) and
                        (LB/LR/LO = 1055/10/5) allocated by PE1 to CE1. Only the label block
                        (LB/LR/LO = 1055/10/5) meets the requirement.
                        PE1 then determines that labels must be allocated from the second label
                        block of CE1. The outgoing VC label (CE13) calculated by PE1 is 1001: LBn +
                        m - LOn = 1000 + 1 - 0 = 1001. The incoming VC label (CE1) calculated by
                        PE1 is 1063: LBm + n - LOm = 1055 + 13 - 5 = 1063.
                        PE3 calculates the VC labels based on this formula. The outgoing VC label
                        calculated by PE3 is 1063: LBn + m - LOn = 1055 + 13 - 5 = 1063; the
                        incoming label calculated by PE3 is 1001: LBm + n - LOm - LOn = 1000 + 1 -
                        0.
                        This example shows that though the label allocation method used by BGP
                        VPWS uses a large number of labels, the number of VCs established on a PE is
                        limited. Therefore, the label consumption can be ignored.
                        In an actual network deployment, the network administrator is used to
                        identifying the location of a CE by its CE ID. If large CE IDs are used, a lot of
                        label spaces will be consumed. In extreme cases, the valid label range will be
                        insufficient. To solve this problem, use CE names to describe CE locations and
                        create a table to record actual CE IDs.

Signaling for Transmitting VC Labels
                    BGP VPWS extends the NLRI of MP-BGP to transmit VC information. Like L3VPN,
                    BGP VPWS also uses the route distinguisher (RD) and VPN target. Because VPWS
                    is a P2P technology, a CE needs to use multiple interfaces or sub-interfaces to
                    establish VCs with multiple other CEs. Even if two CEs are in the same VPN, they
                    can directly communicate with each other only after VCs are established between
                    them.

                    Figure 5-13 describes label block information in the NLRI. The circuit status vector
                    (CSV) in the TLV with variable length is used to describe the LR and tunnel status
                    of a label block.

                    Figure 5-13 MP-BGP extension




                    An extended community attribute is defined to carry more L2VPN information, as
                    shown in Figure 5-14.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             436
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


                    Figure 5-14 Extended community attribute for Layer 2 information




                    Table 5-4 describes each field shown in Figure 5-14.

                    Table 5-4 Description of each field in the extended community attribute

                     Field                  Description      Number of      Description
                                                             Bits

                     Extended               Extended         16             Extended community type
                     Community Type         information
                                            type

                     Encaps Type            Encapsulation    8              Layer 2 encapsulation
                                            type                            type

                     Control Flags          Control word     8              Control word

                     Layer-2 MTU            Layer 2 MTU      16             -

                     Reserved               Reserved         16             Reserved




Application Scenario
                    BGP VPWS applies to networks with dense Layer 2 connections, such as a network
                    with the mesh topology.

Benefits
                    BGP VPWS does not directly perform operations on connections between CEs.
                    Instead, it partitions the entire ISP network into different VPNs and numbers the
                    CEs in each VPN. Similar to BGP/MPLS VPN, BGP VPWS uses VPN targets to
                    control the advertisement or acceptance of VPN routes, which improves
                    networking flexibility.

5.6.2 Configuring a Local BGP VPWS Connection
Prerequisites
                    Before configuring a local BGP VPWS connection, you have completed the
                    following tasks:
                    ●   Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                        network to ensure IP connectivity.
                    ●   Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                        network.
                    ●   Establish tunnels between PEs based on tunnel policies.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           437
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration


Context
                    In Figure 5-15, CE1 and CE2 are connected to the same PE. A local BGP VPWS
                    connection needs to be established between CE1 and CE2 for them to
                    communicate. In this case, the PE functions like a Layer 2 switch and can complete
                    label switching without having BGP configured.

                    Figure 5-15 Network diagram of a local BGP VPWS connection




Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure MPLS L2VPN.
                    mpls l2vpn

         Step 3 Return to the system view.
                    quit

         Step 4 Create a BGP VPWS instance and enter the MPLS L2VPN instance view.
                    mpls l2vpn l2vpn-name [ encapsulation { ethernet | vlan } [ control-word | no-control-word ] ]

         Step 5 Configure an RD for the MPLS L2VPN instance.
                    route-distinguisher route-distinguisher

         Step 6 (Optional) Configure an MTU for the MPLS L2VPN instance.
                    mtu mtu-value

                    The MTU determines the maximum packet size allowed by a VPWS network. If the
                    MTU exceeds the maximum packet size allowed by a VPWS network or an
                    intermediate node (P), there will be packet fragmentation or even dropped
                    packets, which will increase the network transmission load. The MTU is one of
                    VPWS negotiation parameters. If the MTUs of the same VPN instance on the PEs
                    at both ends are different, the two PEs cannot exchange reachability information
                    or establish a PW. An appropriate MTU must be configured for an MPLS L2VPN
                    instance based on the MTU of the interface bound to the L2VPN instance.
                    Specifically, the MTU of an MPLS L2VPN instance cannot exceed the MTU of the
                    interface bound to the L2VPN instance. By default, the MTU of an MPLS L2VPN
                    instance is 1500 bytes.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                        438
VPN Configuration
VPN Configuration                                                                                   5 VPWS Configuration


