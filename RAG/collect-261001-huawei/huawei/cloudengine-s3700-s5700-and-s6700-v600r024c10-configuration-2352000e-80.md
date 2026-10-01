---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-80
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [10980, 11123]
sha256: 5b2e74bb219af943dfb7e38bff793e699f396c96922af3d9fe98f6703bf69a5f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    [ASBR1-bgp-af-vpnv4] peer 10.12.12.2 enable
                    [ASBR1-bgp-af-vpnv4] undo policy vpn-target

                    The configuration of ASBR2 is similar to the configuration of ASBR1. For detailed
                    configurations, see Configuration Scripts.

                           NOTE

                         ASBRs do not filter received VPNv4 routes based on VPN targets. Instead, an ASBR directly
                         sends the VPNv4 routes to the peer ASBR or the PE in the local AS. The VPN instance
                         routing table configured on an ASBR, however, checks the VPN targets of these routes to
                         determine whether to import these routes.

                    ----End

Verifying the Configuration
                    After completing the configuration, run the display ip routing-table command on
                    each CE. The command output shows that the local CE has learned the routes
                    advertised by other CEs. The following example uses the command output on CE3.
                    <CE3> display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table: _public_
                           Destinations : 11         Routes : 11
                    Destination/Mask Proto Pre Cost               Flags NextHop            Interface
                          10.3.1.0/24 Direct 0 0                D 10.3.1.1        Vlanif100
                          10.3.1.1/32 Direct 0 0                D 127.0.0.1        Vlanif100
                        10.3.1.255/32 Direct 0 0                 D 127.0.0.1         Vlanif100
                         11.11.11.11/32 EBGP 255 0                   RD 10.3.1.2          Vlanif100
                         22.22.22.22/32 EBGP 255 0                   RD 10.3.1.2          Vlanif100
                       33.33.33.33/32 Direct 0 0                  D 127.0.0.1        LoopBack1
                         44.44.44.44/32 EBGP 255 0                   RD 10.3.1.2          Vlanif100
                         127.0.0.0/8 Direct 0 0                 D 127.0.0.1        InLoopBack0
                         127.0.0.1/32 Direct 0 0                D 127.0.0.1         InLoopBack0
                    127.255.255.255/32 Direct 0 0                   D 127.0.0.1         InLoopBack0
                    255.255.255.255/32 Direct 0 0                   D 127.0.0.1         InLoopBack0

                    Run the display bgp vpnv4 all routing-table command on an ASBR. The
                    command output shows the VPNv4 routes on the ASBR.
                    The following example uses the command output on ASBR1.
                    <ASBR1> display bgp vpnv4 all routing-table
                     BGP Local router ID is 2.2.2.9
                     Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                               h - history, i - internal, s - suppressed, S - Stale
                               Origin : i - IGP, e - EGP, ? - incomplete
                     RPKI validation codes: V - valid, I - invalid, N - not-found


                    Total number of routes from all PE: 4
                    Route Distinguisher: 100:1


                        Network           NextHop        MED        LocPrf       PrefVal Path/Ogn

                    *>i 11.11.11.11/32    1.1.1.9                100         0     65001i
                    Route Distinguisher: 200:1


                        Network           NextHop        MED        LocPrf       PrefVal Path/Ogn

                    *> 22.22.22.22/32       3.3.3.9                     0        200 65002i


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                            175
VPN Configuration
VPN Configuration                                                                              3 IPv4 L3VPN Configuration

                    Route Distinguisher: 100:3


                        Network          NextHop      MED          LocPrf   PrefVal Path/Ogn
                    *>i 33.33.33.33/32    0.0.0.0    0                 0     65003i
                     Route Distinguisher: 200:4


                        Network        NextHop       MED        LocPrf      PrefVal Path/Ogn
                    *>i 44.44.44.44/32    3.3.3.9     0         100         0    200 65004i


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         vlan batch 100
                         #
                         interface Vlanif100
                          ip address 10.1.1.1 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 100
                         #
                         interface Loopback1
                          ip address 11.11.11.11 255.255.255.255
                         #
                         bgp 65001
                          peer 10.1.1.2 as-number 100
                          #
                          ipv4-family unicast
                           network 11.11.11.11 255.255.255.255
                           peer 10.1.1.2 enable
                         #
                         return
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 100 200
                         #
                         ip vpn-instance vpna
                          ipv4-family
                           route-distinguisher 100:1
                           vpn-target 1:1 export-extcommunity
                           vpn-target 1:1 import-extcommunity
                         #
                         mpls lsr-id 1.1.1.9
                         #
                         mpls
                         #
                         mpls ldp
                         #
                         interface Vlanif100
                          ip address 10.10.1.2 255.255.255.0
                          mpls
                          mpls ldp
                         #
                         interface Vlanif200
                          ip binding vpn-instance vpna
                          ip address 10.1.1.2 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 100
                         #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                         176
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

