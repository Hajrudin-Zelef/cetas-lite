---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-123
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [17382, 17498]
sha256: daed215c592c5ac03412cc89e25196316156d88f6540cf27ec3959c6573a0964
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    PE1 can ping its connected CEs. The following example uses PE1 and CE1.
                    [PE1] ping ipv6 vpn-instance vpna 2001:DB8::11:1
                     PING 2001:DB8::11:1 : 56 data bytes, press CTRL_C to break
                       Reply from 2001:DB8::11:1
                       bytes=56 Sequence=1 hop limit=64 time=2 ms
                       Reply from 2001:DB8::11:1
                       bytes=56 Sequence=2 hop limit=64 time=2 ms
                       Reply from 2001:DB8::11:1
                       bytes=56 Sequence=3 hop limit=64 time=2 ms
                       Reply from 2001:DB8::11:1
                       bytes=56 Sequence=4 hop limit=64 time=2 ms
                       Reply from 2001:DB8::11:1
                       bytes=56 Sequence=5 hop limit=64 time=2 ms

                     --- 2001:DB8::11:1 ping statistics---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max=2/2/2 ms

         Step 2 Configure BGP and import the direct routes destined for local CEs to the VPN
                routing tables on PE1.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] ipv6-family vpn-instance vpna
                    [PE1-bgp-6-vpna] import-route direct
                    [PE1-bgp-6-vpna] quit
                    [PE1-bgp] ipv6-family vpn-instance vpnb


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                           276
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration

                    [PE1-bgp-6-vpnb] import-route direct
                    [PE1-bgp-6-vpnb] quit
                    [PE1-bgp] quit

         Step 3 Configure static routes on CEs.
                    # Configure CE1.
                    [CE1] ipv6 route-static 2001:DB8::12:1 112 2001:DB8::11:2

                    # Configure CE2.
                    [CE2] ipv6 route-static 2001:DB8::11:1 112 2001:DB8::12:2

                    ----End

Verifying the Configuration
                    After the configuration is complete, run the display ipv6 routing-table vpn-
                    instance command on PE1 to check route exchange between VPNs.
                    The following example uses the VPN instance vpna.
                    [PE1] display ipv6 routing-table vpn-instance vpna
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 5        Routes : 5

                    Destination : 2001:DB8::11:0                   PrefixLength : 112
                    NextHop      : 2001:DB8::11:2                  Preference : 0
                    Cost      :0                            Protocol    : Direct
                    RelayNextHop : ::                          TunnelID      : 0x0
                    Interface : Vlanif100                       Flags      :D

                    Destination : 2001:DB8::11:2                   PrefixLength : 128
                    NextHop      : ::1                        Preference : 0
                    Cost      :0                            Protocol    : Direct
                    RelayNextHop : ::                          TunnelID      : 0x0
                    Interface : Vlanif100                       Flags      :D

                    Destination : 2001:DB8::12:0                   PrefixLength : 112
                    NextHop      : 2001:DB8::12:2                  Preference : 255
                    Cost      :0                            Protocol    : BGP
                    RelayNextHop : 2001:DB8::12:2                     TunnelID    : 0x0
                    Interface : Vlanif200                       Flags      : RD

                    Destination : 2001:DB8::12:2                   PrefixLength : 128
                    NextHop      : ::1                        Preference : 255
                    Cost      :0                            Protocol    : BGP
                    RelayNextHop : ::1                          TunnelID      : 0x0
                    Interface : Vlanif200                       Flags      : RD

                    Destination : FE80::                       PrefixLength : 10
                    NextHop      : ::                        Preference : 0
                    Cost      :0                            Protocol   : Direct
                    RelayNextHop : ::                          TunnelID     : 0x0
                    Interface : NULL0                           Flags     : DB

                    CE1 and CE2 can ping each other. The following example uses the command
                    output on CE1.
                    [CE1] ping ipv6 2001:DB8::12:1
                     PING 2001:DB8::12:1 : 56 data bytes, press CTRL_C to break
                      Reply from 2001:DB8::12:1
                      bytes=56 Sequence=1 hop limit=63 time=3 ms
                      Reply from 2001:DB8::12:1
                      bytes=56 Sequence=2 hop limit=63 time=2 ms


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         277
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration

                        Reply from 2001:DB8::12:1
                        bytes=56 Sequence=3 hop limit=63 time=3 ms
                        Reply from 2001:DB8::12:1
                        bytes=56 Sequence=4 hop limit=63 time=3 ms
                        Reply from 2001:DB8::12:1
                        bytes=56 Sequence=5 hop limit=63 time=2 ms

                    --- 2001:DB8::12:1 ping statistics---
                      5 packet(s) transmitted
                      5 packet(s) received
                      0.00% packet loss
                      round-trip min/avg/max=2/2/3 ms


