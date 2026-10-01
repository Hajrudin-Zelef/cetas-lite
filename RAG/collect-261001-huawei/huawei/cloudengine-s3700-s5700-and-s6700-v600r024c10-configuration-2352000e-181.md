---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-181
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [26241, 26413]
sha256: 19bb81445915fac5c576a91aabefa85f0971688fbbf345dae985ccbdbad33cde
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Table 5-2 Default settings for VPWS
                     Parameter                                      Default Setting

                     VPWS connection                                Not established

                     Flow label function                            Disabled

                     L2VPN service name                             Not configured




5.5 Configuring CCC VPWS

5.5.1 Understanding CCC VPWS
                    CCC is a way to implement L2VPN through manual configuration.
                    CCC VPWS needs to be manually configured by administrators and applies to
                    small MPLS networks with simple topologies. It does not require signaling
                    negotiation or exchange of control packets. It consumes few resources and is easy
                    to configure, but its maintenance is inconvenient — requiring manual
                    configuration by administrators — and its scalability is poor.

Topology of Local CCC VPWS
                    A local CCC connection is established between two CEs connected to the same PE.
                    Similar to a Layer 2 switch, a PE can directly complete switching without the need
                    to configure a label switched path (LSP).
                    Figure 5-3 shows the topology when local CCC is used.

                    Figure 5-3 Topology of local CCC VPWS
                         NOTE

                        In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




                    On the network shown in Figure 5-3, the CEs and PE are connected through
                    VLANIF interfaces. A local CCC connection is established between CE1 and CE2.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     416
VPN Configuration
VPN Configuration                                                                          5 VPWS Configuration


                    The PE that CE1 and CE2 are connected to functions like a Layer 2 switch and can
                    transmit data of different link types, such as VLAN and Ethernet data.
                    The advantage of this mode is that no label signaling is required to transmit
                    L2VPN information, and the Internet service provider (ISP) network only needs to
                    support MPLS forwarding.

Topology of Remote CCC VPWS
                    A remote CCC connection is established between two CEs connected to different
                    PEs. Static constraint-based routed label switched paths (CR-LSPs) need to be
                    configured to transmit packets from one PE to the other PE. A static CR-LSP must
                    be configured on PEs and mapped to a CCC connection.
                    Figure 5-4 shows the topology when remote CCC is used.

                    Figure 5-4 Topology of remote CCC VPWS
                         NOTE

                        In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




                    In Figure 5-4, CE1 and CE2 are connected to different PEs. To allow the two CEs to
                    communicate, establish a remote CCC connection. When establishing the remote
                    CCC connection, configure two static CR-LSPs on the P to transmit packets in both
                    directions.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     417
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration


5.5.2 Configuring a Local CCC VPWS Connection

Prerequisites
                    Before configuring a local CCC VPWS connection, you have completed the
                    following tasks:

                    ●      Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                           network to ensure IP connectivity.
                    ●      Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                           network.
                    ●      Establish tunnels between PEs based on tunnel policies. If no tunnel policy is
                           configured, LDP tunnels are established by default.


Context
                    If two CEs are connected to the same PE, you can establish a local CCC connection
                    between the two CEs for them to communicate. In this case, the PE functions like
                    a Layer 2 switch and can directly complete label switching without having LSPs
                    configured.

                    Perform the following operations on the PE.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enable the MPLS L2VPN function.
                    mpls l2vpn

         Step 3 Return to the system view.
                    quit

         Step 4 Create a local CCC connection.
                    ccc ccc-connection-name interface { interface-name1 | interface-type1 interface-number1 } [ raw ] out-
                    interface { interface-name2 | interface-type2 interface-number2 } [ outraw ]

                    ●      Here, you only need to configure the inbound and outbound interfaces of the
                           local CCC connection on the PE. Because a local CCC connection is
                           bidirectional, only one connection is required.
                    ●      The raw parameter is available in this command only for Ethernet links.

         Step 5 (Optional) Configure a description for the local CCC connection.
                    ccc cccName description text

                    To configure a description for a local CCC connection, perform this step. By
                    default, no description is configured for a CCC connection.

                    ----End




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                             418
VPN Configuration
VPN Configuration                                                                                   5 VPWS Configuration


5.5.3 Configuring a Remote CCC VPWS Connection

Prerequisites
                    Before configuring a remote CCC VPWS connection, you have completed the
                    following tasks:

                    ●   Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                        network to ensure IP connectivity.
                    ●   Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                        network.
                    ●   Establish tunnels between PEs based on tunnel policies. If no tunnel policy is
                        configured, LDP tunnels are established by default.


Context
                    To create a remote CCC connection, configure the inbound interface, next-hop IP
                    address, incoming label, and outgoing label for the connection on the PEs, and
                    configure two bidirectional static CR-LSPs on the Ps. A remote CCC connection is
                    unidirectional, and therefore two such connections must be created.


Procedure
                    ●   Configure PEs.

                        Perform the following operations on the PEs at both ends of a VC.

                        a.   Enter the system view.
                             system-view

                        b.   Create a remote CCC connection.
                             ccc ccc-connection-name interface { interface-name1 | interface-type1 interface-number1 }
                             [ raw ] in-label in-label-value out-label out-label-value nexthop nexthop-address [ control-
                             word | no-control-word ]

