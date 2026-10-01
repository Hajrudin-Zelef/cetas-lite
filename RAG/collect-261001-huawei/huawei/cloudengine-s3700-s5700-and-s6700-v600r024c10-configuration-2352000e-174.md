---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-174
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [25180, 25336]
sha256: 475e55550ffdce3cf23f5b5c713398eebf582a537c7b32ac1f419b8ba5f620e0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

4.14.6 Configuring Route Exchange Between the PE and CE

Context
                    The routing protocol run between Spoke-PEs and Spoke-CEs is related to the
                    routing protocol run between the Hub-PE and Hub-CE. EBGP, IGP, and static routes

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                        397
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration


                    (including default routes) can be used between the Hub-PE and Hub-CE.
                    Determine which one to use as required.


Procedure
                    ●   Configure EBGP between the Hub-PE and Hub-CE.

                        For configuration details, see Configuring an IPv6 VPN Instance.

                        In this mode, EBGP, IGP, or static routes (including default routes) can be used
                        between Spoke-PEs and Spoke-CEs.

                               NOTE

                              If EBGP runs both between Spoke-PEs and Spoke-CEs and between the Hub-PE and
                              Hub-CE, you must run the peer ip-address allow-as-loop [ number ] command in the
                              BGP-VPN instance IPv4 address family view of the Hub-PE to allow routing loops. If
                              number is set to 1, routes with the local AS number repeated once in the AS_Path list
                              are allowed.
                    ●   Configure IGP between the Hub-PE and Hub-CEs.

                        For configuration details, see Configuring an IPv6 VPN Instance.

                        In this mode, only IGP or static routes (including default routes) can be used
                        between Spoke-PEs and Spoke-CEs.
                    ●   Configuring static routes (including default routes) between the Hub-PE and
                        Hub-CEs.

                        For configuration details, see Configuring an IPv6 VPN Instance.

                        In this mode, EBGP, IGP, or static routes (including default routes) can be used
                        between Spoke-PEs and Spoke-CEs.

                        –     In a CE dual-link access scenario, if a Hub-CE adopts the default route to
                              access the Hub-PE, you need to run the following commands on the Hub-
                              PE to enable it to advertise the default route to all the Spoke-PEs.
                              i.     Enter the system view.
                                     system-view

                              ii.    Configure the default route.
                                     ipv6 route-static vpn-instance vpn-instance-name :: 0 nexthop-address

                                     In this example, vpn-instance-name specifies VPN-out and nexthop-
                                     address specifies the IPv6 address of the Hub-CE interface that
                                     connects to the Hub-PE interface bound to VPN-out.
                              iii.   Enter the BGP view.
                                     bgp as-number

                              iv.    Enter the BGP-VPN instance IPv6 address family view.
                                     ipv6-family vpn-instance vpn-instance-name

                              v.     Import the default route of the local site.
                                     network :: 0

                    ----End

4.14.7 Verifying the Configuration

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                  398
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration


Procedure
                    ●    Run the display ipv6 routing-table vpn-instance vpn-instance-name
                         command on the Hub-PE to check the routing information of VPN-in and
                         VPN-out.
                    ●    Run the display ipv6 routing-table command on the Hub-CE and all Spoke-
                         CEs to check routing information.
                    ----End

4.14.8 Example for Configuring IPv6 L3VPN over MPLS Hub-
Spoke
Networking Requirements
                    On the network shown in Figure 4-11, the communication between the Spoke-
                    CEs is controlled by the Hub-CE at the central site. In other words, the traffic
                    between Spoke-CEs is forwarded also through the Hub-CE, not only through the
                    Hub-PE.

                    Figure 4-11 Hub-spoke networking
                          NOTE

                    In this example, interface 1, interface 2, interface 3, and interface 4 represent VLANIF 100,
                    VLANIF 200, VLANIF 300, and VLANIF 400, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         399
VPN Configuration
VPN Configuration                                                           4 IPv6 L3VPN Configuration




Precautions
                    During the configuration, note the following:
                    ●    The import and export VPN targets configured on a Spoke-PE are different.
                    ●    Two VPN instances (vpn_in and vpn_out) are created on the Hub-PE. The
                         VPN targets received by vpn_in are the VPN targets advertised by the two
                         Spoke-PEs; the VPN targets advertised by vpn_out are the VPN targets
                         received by the two Spoke-PEs and are different from the VPN targets
                         received by vpn_in.
                    ●    The Hub-PE is configured to accept the routes with AS numbers repeated
                         once in the AS_Path attribute.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Establish MP-IBGP peer relationships between the Hub-PE and Spoke-PEs.
                         There is no need to establish an MP-IBGP peer relationship or exchange VPN
                         routing information between the two Spoke-PEs.
                    2.   Create VPN instances and VPN targets on PEs.
                    3.   Configure EBGP connections between CEs and PEs.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           400
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration


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
                    on each device. The command output shows that the Session State field displays
                    Operational.

