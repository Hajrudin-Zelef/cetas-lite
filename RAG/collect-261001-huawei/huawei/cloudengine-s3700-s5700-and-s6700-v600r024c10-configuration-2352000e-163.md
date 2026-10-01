---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-163
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [23431, 23557]
sha256: b85edba7a5141a8d8908f17aa9f40fea3b997f83bd659535868204b5cdd5334d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Destination : 2001:db8:3::                PrefixLength : 64
                    NextHop      : 2001:db8:1::2               Preference : 255
                    Cost      :0                        Protocol     : EBGP
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif100                   Flags       :D

                     Destination : FE80::                 PrefixLength : 10
                     NextHop       : ::                 Preference : 0
                     Cost       :0                    Protocol    : Direct
                     RelayNextHop : ::                    TunnelID     : 0x0
                     Interface : NULL0                     Flags     :D
                    <CE1> ping ipv6 2001:db8:2::1
                     PING 2001:db8:2::1 : 56 data bytes, press CTRL_C to break
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=1 hop limit=60 time = 94 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=2 hop limit=60 time = 109 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=3 hop limit=60 time = 110 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=4 hop limit=60 time = 94 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=5 hop limit=60 time = 110 ms

                     --- 2001:db8:2::1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 94/103/110 ms

                    Run the display ipv6 routing-table vpn-instance command on an ASBR. The
                    command output shows the routing table of the VPN instance IPv6 address
                    maintained on the ASBR.

                    The following example uses the command output on ASBR1.
                    <ASBR1> display ipv6 routing-table vpn-instance vpn1
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                           Destinations : 5        Routes : 5


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         370
VPN Configuration
VPN Configuration                                                                                 4 IPv6 L3VPN Configuration


                    Destination : 2001:db8:1::                PrefixLength : 64
                    NextHop      : ::FFFF:1.1.1.9            Preference : 255
                    Cost      :0                        Protocol     : IBGP
                    RelayNextHop : ::                      TunnelID        : 0xa0010082
                    Interface : NULL0                       Flags        : RD

                    Destination : 2001:db8:2::                PrefixLength : 64
                    NextHop      : 2001:db8:3::2               Preference : 255
                    Cost      :0                        Protocol     : EBGP
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif200                   Flags       :D

                    Destination : 2001:db8:3::                PrefixLength : 64
                    NextHop      : 2001:db8:3::1               Preference : 0
                    Cost      :0                        Protocol     : Direct
                    RelayNextHop : ::                      TunnelID       : 0x0
                    Interface : Vlanif200                   Flags       :D

                    Destination : 2001:db8:3::1               PrefixLength : 128
                    NextHop      : ::1                    Preference : 0
                    Cost      :0                        Protocol    : Direct
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif200                   Flags      :D

                    Destination : FE80::                   PrefixLength : 10
                    NextHop      : ::                    Preference : 0
                    Cost      :0                        Protocol   : Direct
                    RelayNextHop : ::                      TunnelID     : 0x0
                    Interface : NULL0                       Flags     :D

                    Run the display bgp vpnv6 all routing-table command on an ASBR. The
                    command output shows the VPN IPv6 routes on the ASBR.
                    The following example uses the command output on ASBR1.
                    <ASBR1> display bgp vpnv6 all routing-table

                    BGP Local router ID is 2.2.2.9
                     Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                              h - history, i - internal, s - suppressed, S - Stale
                              Origin : i - IGP, e - EGP, ? - incomplete


                    Total number of routes from all PE: 4
                    Route Distinguisher: 100:1


                    *>i Network : 2001:db8:1::                       PrefixLen : 64
                       NextHop : ::FFFF:1.1.1.9                     LocPrf : 100
                       MED     :0                               PrefVal : 0
                       Label : 105472
                       Path/Ogn : ?

                    Route Distinguisher: 100:2


                    *> Network : 2001:db8:2::                        PrefixLen : 64
                       NextHop : 2001:db8:3::2                       LocPrf :
                       MED    :                                 PrefVal : 0
                       Label : NULL
                       Path/Ogn : 200 ?
                    *> Network : 2001:db8:3::                        PrefixLen : 64
                       NextHop : ::                             LocPrf :
                       MED    :0                                PrefVal : 0
                       Label : NULL
                       Path/Ogn : ?
                    *
                       NextHop : 2001:db8:3::2                       LocPrf     :
                       MED    :0                                PrefVal : 0


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                          371
VPN Configuration
VPN Configuration                                                                  4 IPv6 L3VPN Configuration

                        Label : NULL
                        Path/Ogn : 200 ?

                    VPN-Instance vpn1 :

