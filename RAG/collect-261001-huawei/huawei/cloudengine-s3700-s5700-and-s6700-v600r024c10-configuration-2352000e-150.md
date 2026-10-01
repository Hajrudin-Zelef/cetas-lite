---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-150
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [21400, 21547]
sha256: 67c4f4846ad1850e72e1f0377c943cdcbb5cea77f140419386751dd2d229786a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   CE
                        #
                        sysname CE
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:DB8:1::1/64
                        #
                        interface Vlanif200
                         ipv6 enable
                         ipv6 address 2001:DB8:3::1/64
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:DB8:2::1/128
                        #
                        bgp 65410
                         router-id 10.10.10.10
                         peer 2001:DB8:1::2 as-number 100
                         peer 2001:DB8:3::2 as-number 100
                         #
                         ipv4-family unicast
                         #
                         ipv6-family unicast
                          network 2001:DB8:0:1:2::1 128
                          peer 2001:DB8:1::2 enable
                          peer 2001:DB8:3::2 enable
                        #
                        return


4.10.6 Example for Configuring VPN IPv6 FRR

Networking Requirements
                    At a VPN site, different CEs use BGP to access the same PE. The PE learns multiple
                    IPv6 VPN routes with the same VPN route prefix from the CEs. To enable these
                    routes to form primary/backup relationships, configure VPN IPv6 FRR. After VPN
                    IPv6 FRR is configured, the PE generates a pair of primary and backup routes to
                    the VPN route prefix. After that, VPN IPv6 traffic can be quickly switched to the
                    backup route when the primary route fails.

                    As shown in Figure 4-5, an EBGP peer relationship is established between the PE
                    and CE1 and between the PE and CE2. There are two BGP routes from the PE to
                    Loopback 1 on DeviceA. The optimal route traverses Link_A, and the second-


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         339
VPN Configuration
VPN Configuration                                                                    4 IPv6 L3VPN Configuration


                    optimal route traverses Link_B. It is required that VPN IPv6 auto FRR be configured
                    on the PE so that IPv6 traffic can be quickly switched to Link_B if Link_A fails.

                    Figure 4-5 Configuring VPN IPv6 auto FRR
                         NOTE

                    In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200, respectively.




Precautions
                    In a VPN FRR scenario, traffic is switched back to the primary path after this path
                    recovers. Because the order in which nodes undergo IGP convergence differs,
                    packet loss may occur during the switchback. To resolve this problem, run the
                    route-select delay delay-value command to configure a route selection delay so
                    that traffic is switched back only after forwarding entries on the devices along the
                    primary path are updated. The delay specified using delay-value depends on
                    various factors, such as the number of routes on each device. Configure an
                    appropriate delay as needed.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure IGP at the VPN site, so that the routes to the loopback interface of
                         DeviceA can be advertised to CE1 and CE2.
                    2.   Configure a VPN instance named vpna that supports the IPv6 address family
                         on the PE, and bind VLANIF 100 and VLANIF 200 to vpna.
                    3.   Establish an EBGP peer relationship between the PE and CE1 and between the
                         PE and CE2. On CE1 and CE2, configure IGP and BGP to import routes from
                         each other.
                    4.   Enable VPN IPv6 auto FRR on the PE.

Procedure
         Step 1 Configure IPv6 addresses for the interfaces on the devices at the VPN site.
                    For detailed configurations, see Configuration Scripts.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     340
VPN Configuration
VPN Configuration                                                                    4 IPv6 L3VPN Configuration


         Step 2 Configure IGP at the VPN site, so that the routes to the loopback interface of
                DeviceA can be advertised to CE1 and CE2. This example uses OSPFv3 as IGP.
                    # Configure CE1.
                    [CE1] ospfv3 1
                    [CE1-ospfv3-1] router-id 2.2.2.2
                    [CE1-ospfv3-1] quit
                    [CE1] interface Vlanif200
                    [CE1-Vlanif200] ospfv3 1 area 0.0.0.0
                    [CE1-Vlanif200] quit

                    The configurations of CE2 and DeviceA are similar to the configuration of CE1. For
                    detailed configurations, see Configuration Scripts.
                    After completing the configuration, run the display ipv6 routing-table command
                    on the CEs. The command output shows that CE1 and CE2 have learned the route
                    to Loopback 1 on DeviceA. The following example uses the command output on
                    CE1.
                    [~CE1] display ipv6 routing-table
                    Routing Table : _public_
                          Destinations : 10    Routes : 10

                    Destination : ::1                      PrefixLength : 128
                    NextHop      : ::1                     Preference : 0
                    Cost      :0                         Protocol    : Direct
                    RelayNextHop : ::                       TunnelID      : 0x0
                    Interface : InLoopBack0                    Flags       :D

                    Destination : ::FFFF:127.0.0.0             PrefixLength : 104
                    NextHop      : ::FFFF:127.0.0.1            Preference : 0
                    Cost      :0                         Protocol    : Direct
                    RelayNextHop : ::                       TunnelID      : 0x0
                    Interface : InLoopBack0                   Flags        :D

                    Destination : ::FFFF:127.0.0.1             PrefixLength : 128
                    NextHop      : ::1                     Preference : 0
                    Cost      :0                         Protocol    : Direct
                    RelayNextHop : ::                       TunnelID      : 0x0
                    Interface : InLoopBack0                    Flags       :D

                    Destination : 2001:DB8:0::                  PrefixLength : 64
                    NextHop      : 2001:DB8:4::2                 Preference : 0
                    Cost      :0                          Protocol    : Direct
                    RelayNextHop : ::                        TunnelID      : 0x0
                    Interface : Vlanif100             Flags      :D

