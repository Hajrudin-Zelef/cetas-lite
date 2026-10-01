---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-74
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [10022, 10118]
sha256: ee5cec7a9be55d2ea5788f48d1d20dce33a7945a4cc285b0078b13079440fe8f
---

                    # On ASBR1, enable MPLS on VLANIF 200 connected to ASBR2.
                    [ASBR1] interface 10GE1/0/2
                    [ASBR1-10GE1/0/2] port link-type trunk
                    [ASBR1-10GE1/0/2] port trunk allow-pass vlan 200
                    [ASBR1-10GE1/0/2] quit
                    [ASBR1] interface Vlanif 200
                    [ASBR1-Vlanif200] ip address 192.168.1.1 24
                    [ASBR1-Vlanif200] mpls
                    [ASBR1-Vlanif200] quit

                    # On ASBR1, establish an MP-EBGP peer relationship with ASBR2, and disable
                    ASBR1 from filtering received VPNv4 routes based on VPN targets.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 192.168.1.2 as-number 200
                    [ASBR1-bgp] ipv4-family vpnv4
                    [ASBR1-bgp-af-vpnv4] peer 192.168.1.2 enable
                    [ASBR1-bgp-af-vpnv4] undo policy vpn-target



                    ----End

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                    159
VPN Configuration
VPN Configuration                                                                                        3 IPv4 L3VPN Configuration


Verifying the Configuration
                    After the configuration is complete, the CEs can learn routes to the loopback
                    interface of each other, and can ping each other successfully.
                    The following example uses the command output on CE1.
                    <CE1> display ip routing-table
                    Proto: Protocol         Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table: _public_
                             Destinations : 9      Routes : 9
                    Destination/Mask Proto Pre Cost               Flags NextHop            Interface
                            5.5.5.5/32 Direct 0 0              D 127.0.0.1        LoopBack1
                           6.6.6.6/32 EBGP 255 0                 D 10.1.1.2         Vlanif100
                           10.1.1.0/24 Direct 0 0               D 10.1.1.1        Vlanif100
                           10.1.1.1/32 Direct 0 0               D 127.0.0.1        Vlanif100
                         10.1.1.255/32 Direct 0 0                D 127.0.0.1         Vlanif100
                          127.0.0.0/8 Direct 0 0                D 127.0.0.1        InLoopBack0
                          127.0.0.1/32 Direct 0 0               D 127.0.0.1         InLoopBack0
                    127.255.255.255/32 Direct 0 0                   D 127.0.0.1         InLoopBack0
                    255.255.255.255/32 Direct 0 0                   D 127.0.0.1         InLoopBack0
                    <CE1> ping -a 5.5.5.5 6.6.6.6
                      PING 6.6.6.6: 56 data bytes, press CTRL_C to break
                        Reply from 6.6.6.6: bytes=56 Sequence=1 ttl=252 time=120 ms
                        Reply from 6.6.6.6: bytes=56 Sequence=2 ttl=252 time=73 ms
                        Reply from 6.6.6.6: bytes=56 Sequence=3 ttl=252 time=111 ms
                        Reply from 6.6.6.6: bytes=56 Sequence=4 ttl=252 time=86 ms
                        Reply from 6.6.6.6: bytes=56 Sequence=5 ttl=252 time=110 ms
                      --- 6.6.6.6 ping statistics ---
                        5 packet(s) transmitted
                        5 packet(s) received
                        0.00% packet loss
                        round-trip min/avg/max = 73/100/120 ms

                    Run the display bgp vpnv4 all routing-table command on an ASBR. The
                    command output shows the VPNv4 routes on the ASBR.
                    The following example uses the command output on ASBR1.
                    <ASBR1> display bgp vpnv4 all routing-table
                     BGP Local router ID is 2.2.2.2
                     Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                              h - history, i - internal, s - suppressed, S - Stale
                              Origin : i - IGP, e - EGP, ? - incomplete RPKI validation codes: V - valid, I - invalid, N - not-found

                    Total number of routes from all PE: 2
                    Route Distinguisher: 100:1

                        Network            NextHop           MED         LocPrf       PrefVal Path/Ogn

                    *>i 5.5.5.5/32      1.1.1.1          0         100       0         65001i
                    Route Distinguisher: 200:1

                        Network            NextHop           MED         LocPrf       PrefVal Path/Ogn

                    *> 6.6.6.6/32          192.168.1.2                            0    200 65002i


Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan batch 100
                          #


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                                    160
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

