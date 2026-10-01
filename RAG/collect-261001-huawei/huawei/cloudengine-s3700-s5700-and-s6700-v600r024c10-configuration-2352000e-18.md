---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-18
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [1584, 1717]
sha256: 8a48fde6cfb4d4a247b58a5a9819ad15fe22849dd54888bdf39301f722204e2b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After the configuration is complete, run the display ip routing-table command on
                    DeviceA and DeviceC. The command output shows that DeviceA and DeviceC can
                    learn OSPF routes to the network segments of the remote interfaces.
                    # The following example uses the command output on DeviceA.
                    [DeviceA] display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : _public_
                           Destinations : 8        Routes : 8

                    Destination/Mask     Proto Pre Cost        Flags NextHop         Interface

                       172.20.1.0/24 Direct 0 0             D 172.20.1.1    Vlanif1
                       172.20.1.1/32 Direct 0 0             D 127.0.0.1     Vlanif1
                      172.20.1.255/32 Direct 0 0             D 127.0.0.1     Vlanif1
                       172.21.1.0/24 OSPF 10 2               D 172.20.1.2     Vlanif1
                        127.0.0.0/8 Direct 0 0             D 127.0.0.1    InLoopBack0
                        127.0.0.1/32 Direct 0 0            D 127.0.0.1     InLoopBack0
                    127.255.255.255/32 Direct 0 0             D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0             D 127.0.0.1      InLoopBack0

         Step 3 Configure tunnel interfaces.
                    # Configure DeviceA.
                    [DeviceA] interface tunnel 1
                    [DeviceA-Tunnel1] tunnel-protocol gre
                    [DeviceA-Tunnel1] ipv6 enable
                    [DeviceA-Tunnel1] ipv6 address 2001:DB8:3::1 64
                    [DeviceA-Tunnel1] source 172.20.1.1
                    [DeviceA-Tunnel1] destination 172.21.1.2
                    [DeviceA-Tunnel1] quit

                    # Configure DeviceC.
                    [DeviceC] interface tunnel 1
                    [DeviceC-Tunnel1] tunnel-protocol gre
                    [DeviceC-Tunnel1] ipv6 enable
                    [DeviceC-Tunnel1] ipv6 address 2001:DB8:3::2 64
                    [DeviceC-Tunnel1] source 172.21.1.2
                    [DeviceC-Tunnel1] destination 172.20.1.1
                    [DeviceC-Tunnel1] quit

                    After the configuration is complete, the IPv6 status of tunnel interfaces becomes
                    up and the tunnel interfaces can ping each other successfully.
                    # The following example uses the command output on DeviceA.
                    [DeviceA] ping ipv6 -a 2001:DB8:3::1 2001:DB8:3::2
                     PING 2001:DB8:3::2 : 56 data bytes, press CTRL_C to break
                      Reply from 2001:DB8:3::2
                      bytes=56 Sequence=1 hop limit=64 time=11 ms
                      Reply from 2001:DB8:3::2
                      bytes=56 Sequence=2 hop limit=64 time=5 ms
                      Reply from 2001:DB8:3::2


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                21
VPN Configuration
VPN Configuration                                                                                    2 GRE Configuration

                       bytes=56 Sequence=3 hop limit=64 time=5 ms
                       Reply from 2001:DB8:3::2
                       bytes=56 Sequence=4 hop limit=64 time=5 ms
                       Reply from 2001:DB8:3::2
                       bytes=56 Sequence=5 hop limit=64 time=5 ms

                     --- 2001:DB8:3::2 ping statistics---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max=5/6/11 ms

         Step 4 Configure static routes.
                    # Configure DeviceA.
                    [DeviceA] ipv6 route-static 2001:DB8:2:: 64 tunnel1

                    # Configure DeviceC.
                    [DeviceC] ipv6 route-static 2001:DB8:1:: 64 tunnel1

                    ----End

Verifying the Configuration
                    After the configuration is complete, run the display ipv6 routing-table command
                    on DeviceA and DeviceC. The command output shows the IPv6 static route with
                    the local tunnel interface as the outbound interface and destined for the user-side
                    network segment of the remote end.
                    # The following example uses the command output on DeviceA.
                    [DeviceA] display ipv6 routing-table
                    Route Flags: R - relay, D - download to fib, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : _public_
                           Destinations : 8        Routes : 8

                    Destination : ::1                              PrefixLength : 128
                    NextHop      : ::1                             Preference : 0
                    Cost      :0                                 Protocol    : Direct
                    RelayNextHop : ::                               TunnelID      : 0x0
                    Interface : InLoopBack0                            Flags       :D

                    Destination : ::FFFF:127.0.0.0                     PrefixLength : 104
                    NextHop      : ::FFFF:127.0.0.1                    Preference : 0
                    Cost      :0                                 Protocol    : Direct
                    RelayNextHop : ::                               TunnelID      : 0x0
                    Interface : InLoopBack0                           Flags        :D

                    Destination : ::FFFF:127.0.0.1                     PrefixLength : 128
                    NextHop      : ::1                             Preference : 0
                    Cost      :0                                 Protocol    : Direct
                    RelayNextHop : ::                               TunnelID      : 0x0
                    Interface : InLoopBack0                            Flags       :D

                    Destination : 2001:DB8:3::                         PrefixLength : 64
                    NextHop      : 2001:DB8:3::1                        Preference : 0
                    Cost      :0                                 Protocol    : Direct
                    RelayNextHop : ::                               TunnelID      : 0x0
                    Interface : Tunnel1                             Flags       :D

                    Destination : 2001:DB8:3::1                        PrefixLength : 128
                    NextHop      : ::1                             Preference : 0
                    Cost      :0                                 Protocol    : Direct
                    RelayNextHop : ::                               TunnelID      : 0x0


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                      22
VPN Configuration
VPN Configuration                                                                    2 GRE Configuration

                    Interface   : Tunnel1                       Flags   :D

                    Destination : 2001:DB8:1::                  PrefixLength : 64
                    NextHop      : ::1                      Preference : 0
                    Cost      :0                          Protocol    : Direct
                    RelayNextHop : ::                        TunnelID      : 0x0
                    Interface : Vlanif2                    Flags       :D

