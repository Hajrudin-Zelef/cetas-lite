---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-115
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [16114, 16285]
sha256: 22e46620683cea5f3caac232145a5abc99ccf1ccafd4cd15a10102688f4e72f0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Hub-PE
                        #
                        sysname Hub-PE
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpnhub
                         ipv4-family
                          route-distinguisher 100:21
                          export route-policy policy_in
                          apply-label per-route pop-go
                          vpn-target 200:1 export-extcommunity
                          vpn-target 100:1 import-extcommunity
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 20.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 11.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpnhub
                         ip address 10.2.1.2 255.255.255.0



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         255
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpn-instance vpnhub
                          peer 10.2.1.1 as-number 65430
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 11.1.1.0 0.0.0.255
                          network 20.1.1.0 0.0.0.255
                        #
                        ip ip-prefix defaultip index 10 permit 0.0.0.0 0
                        #
                        route-policy policy_in permit node 1
                         if-match ip-prefix defaultip
                        #
                        route-policy policy_in deny node 2
                        #
                        return


3.15.9 Example 2 for Configuring Hub-Spoke (Single Link
Between the Hub-PE and Hub-CE)
Networking Requirements
                    On the network shown in Figure 3-50, the communication between the Spoke-
                    CEs is controlled by the Hub-CE at the central site. In other words, the traffic
                    between Spoke-CEs is forwarded also through the Hub-CE, not only through the
                    Hub-PE. The Hub-PE connects to the Hub-CE and Spoke-CE3 over one link each.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         256
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


                    Figure 3-50 Example 2 for configuring hub-spoke (single link between the Hub-PE
                    and Hub-CE)
                          NOTE

                    In this example, interface 1, interface 2, interface 3, and interface 4 represent VLANIF 100,
                    VLANIF 200, VLANIF 300, and VLANIF 400, respectively.




Precautions
                    Note the following during the configuration:
                    ●    The import and export VPN targets configured on a Spoke-PE are different.
                    ●    A VPN instance (vpnhub) is created on the Hub-PE. The VPN targets received
                         by vpnhub are the VPN targets advertised by the two Spoke-PEs; the VPN
                         targets advertised by vpnhub are the VPN targets received by the two Spoke-
                         PEs and are different from the VPN targets received by vpnhub.
                    ●    In this scenario, to avoid loops, ensure that connected interfaces have STP
                         disabled and are removed from VLAN 1. If STP is enabled and VLANIF
                         interfaces of devices are used to construct a Layer 3 ring network, an
                         interface on the network will be blocked. As a result, Layer 3 services on the
                         network cannot run normally.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         257
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Establish MP-IBGP peer relationships between the Hub-PE and Spoke-PEs.
                         (There is no need to establish an MP-IBGP peer relationship or exchange VPN
                         routing information between the two Spoke-PEs.)
                    2.   Create a VPN instance on each PE and configure pop-go on the Hub-PE.
                    3.   Configure EBGP connections between CEs and PEs. Configure the Hub-CE to
                         advertise the default route to the Hub-PE.

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

