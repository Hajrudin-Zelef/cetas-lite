---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-107
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [14944, 15077]
sha256: 27a1b310ff4804a9a63fd207d6a6f22009adc655f2aec5829511f731fb263436
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        –     If the Hub-PE and Hub-CE are connected over a single link, the Hub-CE
                              needs to use the default route to access the Hub-PE. In this case, perform
                              the following operations in the BGP view of the Hub-CE:
                              i.     Enter the system view.
                                     system-view

                              ii.    Enter the BGP view.
                                     bgp as-number

                              iii.   Configure a peer.
                                     peer ipv4-address as-number as-number

                              iv.    Configure the Hub-CE to send the default route to the Hub-PE.
                                     peer ipv4-address default-route-advertise [ route-policy route-policy-name ]
                                     [ conditional-route-match-all { ipv4-address1 { mask1 | mask-length1 } } &<1-4> |
                                     conditional-route-match-any { ipv4-address2 { mask2 | mask-length2 } } &<1-4> ]

                    ----End

3.15.6 Verifying the Configuration

Procedure
                    ●   Run the display ip routing-table vpn-instance vpn-instance-name command
                        on the Hub-PE to check the routing information of vpn_in and vpn_out.


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                           237
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


                    ●    Run the display ip routing-table command on the Hub-CE and all Spoke-CEs
                         to check routing information.
                    ----End

3.15.7 Example for Configuring Hub-Spoke (Double Links
Between the Hub-PE and Hub-CE)
Networking Requirements
                    On the network shown in Figure 3-48, the communication between the Spoke-
                    CEs is controlled by the Hub-CE at the central site. In other words, the traffic
                    between Spoke-CEs is forwarded also through the Hub-CE, not only through the
                    Hub-PE. The Hub-CE accesses the Hub-PE over double links.

                    Figure 3-48 Hub-spoke networking
                          NOTE

                    In this example, interface 1, interface 2, interface 3, and interface 4 represent VLANIF 100,
                    VLANIF 200, VLANIF 300, and VLANIF 400, respectively.




Precautions
                    Note the following during the configuration:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         238
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                    ●    The import and export VPN targets configured on a Spoke-PE are different.
                    ●    Two VPN instances (vpn_in and vpn_out) are created on the Hub-PE. The
                         VPN targets received by vpn_in are the VPN targets advertised by the two
                         Spoke-PEs; the VPN targets advertised by vpn_out are the VPN targets
                         received by the two Spoke-PEs and are different from the VPN targets
                         received by vpn_in.
                    ●    The Hub-PE is configured to accept the routes with AS numbers repeated
                         once in the AS_Path attribute.
                    ●    In this scenario, to avoid loops, ensure that connected interfaces have STP
                         disabled and are removed from VLAN 1. If STP is enabled and VLANIF
                         interfaces of devices are used to construct a Layer 3 ring network, an
                         interface on the network will be blocked. As a result, Layer 3 services on the
                         network cannot run normally.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Enable OSPF on the backbone network to ensure that PEs can communicate.
                    2.   Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                         backbone network.
                    3.   Configure a VPN instance on each PE, enable the IPv4 address family for the
                         instance, and bind the interface that connects each PE to a CE to the VPN
                         instance on that PE.
                    4.   Enable MP-IBGP on PEs to exchange VPN routing information.
                    5.   Configure EBGP between CEs and PEs to exchange VPN routing information.

Procedure
         Step 1 Configure IGP on the backbone network for the Hub-PE and Spoke-PEs to
                communicate.
                    OSPF is used as IGP in this example. For detailed configurations, see Configuration
                    Scripts.
                    After the configuration is complete, OSPF neighbor relationships are established
                    between the Hub-PE and Spoke-PEs. Run the display ospf peer command. The
                    command output shows that the neighbor status is Full. Run the display ip
                    routing-table command. The command output shows that the Hub-PE and
                    Spoke-PEs have learned the routes to each other's loopback interface.
         Step 2 Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                backbone network.
                    For detailed configurations, see Configuration Scripts.
                    After the configuration is complete, LDP peer relationships are established
                    between the Hub-PE and Spoke-PEs. Run the display mpls ldp session command
                    on each device. The command output shows that Session State is Operational.
         Step 3 Configure a VPN instance on each PE, enable the IPv4 address family for the
                instance, and bind the interface that connects each PE to a CE to the VPN instance
                on that PE.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           239
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


                          NOTE

                        The import VPN target list of a VPN instance on the Hub-PE must contain the export VPN
                        targets of all Spoke-PEs.
                        The export VPN target list of the other VPN instance on the Hub-PE must contain the
                        import VPN targets of all Spoke-PEs.

                    # Configure Spoke-PE1.
                    <Spoke-PE1> system-view
                    [Spoke-PE1] ip vpn-instance vpna
                    [Spoke-PE1-vpn-instance-vpna] ipv4-family
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] vpn-target 100:1 export-extcommunity
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] vpn-target 200:1 import-extcommunity
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] quit
                    [Spoke-PE1-vpn-instance-vpna] quit
                    [Spoke-PE1] interface Vlanif 100
                    [Spoke-PE1-Vlanif100] ip binding vpn-instance vpna
                    [Spoke-PE1-Vlanif100] ip address 10.1.1.2 24
                    [Spoke-PE1-Vlanif100] quit

