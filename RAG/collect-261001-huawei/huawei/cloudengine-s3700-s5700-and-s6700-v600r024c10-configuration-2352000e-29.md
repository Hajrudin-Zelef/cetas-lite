---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-29
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3191, 3358]
sha256: f3d7c23d585c84ffc4b8dd21bee3dd184ddfba78ff3c52ddf631a5ff45fd1e47
---

                    # Configure an IP address for a related interface on CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 100
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [CE1-10GE1/0/1] quit
                    [CE1] interface Vlanif 100
                    [CE1-Vlanif100] ip address 10.1.1.1 24
                    [CE1-Vlanif100] quit


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   49
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration


                    # Configure an IP address for the related interface on CE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] vlan batch 100
                    [CE2] interface 10ge 1/0/1
                    [CE2-10GE1/0/1] port link-type trunk
                    [CE2-10GE1/0/1] port trunk allow-pass vlan 100
                    [CE2-10GE1/0/1] quit
                    [CE2] interface Vlanif 100
                    [CE2-Vlanif100] ip address 10.2.1.1 24
                    [CE2-Vlanif100] quit

                    The PE can ping its connected CE. The following uses the ping between PE1 and
                    CE1 as an example.
                    [PE1] ping -vpn-instance vpna 10.1.1.1
                     PING 10.1.1.1: 56 data bytes, press CTRL_C to break
                       Reply from 10.1.1.1: bytes=56 Sequence=1 ttl=255 time=2 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=2 ttl=255 time=2 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=3 ttl=255 time=2 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=4 ttl=255 time=2 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=5 ttl=255 time=2 ms

                     --- 10.1.1.1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 2/2/2 ms

         Step 2 Configure BGP and import the local direct routes destined for CEs to the VPN
                routing table on PE1.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] ipv4-family vpn-instance vpna
                    [PE1-bgp-vpna] import-route direct
                    [PE1-bgp-vpna] quit
                    [PE1-bgp] ipv4-family vpn-instance vpnb
                    [PE1-bgp-vpnb] import-route direct
                    [PE1-bgp-vpnb] quit
                    [PE1-bgp] quit

         Step 3 Configure a static route on each CE.
                    # Configure CE1.
                    [CE1] ip route-static 10.2.1.0 24 10.1.1.2

                    # Configure CE2.
                    [CE2] ip route-static 10.1.1.0 24 10.2.1.2

                    ----End

Verifying the Configuration
                    After the configuration is complete, run the display ip routing-table vpn-
                    instance command on PE1. The VPNs have imported routes of each other. The
                    following uses vpna as an example.
                    [PE1] display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          50
VPN Configuration
VPN Configuration                                                                          3 IPv4 L3VPN Configuration

                         Destinations : 7    Routes : 7

                    Destination/Mask   Proto Pre Cost        Flags NextHop     Interface

                         10.1.1.0/24 Direct 0 0           D 10.1.1.2     Vlanif100
                         10.1.1.2/32 Direct 0 0           D 127.0.0.1    Vlanif100
                       10.1.1.255/32 Direct 0 0            D 127.0.0.1    Vlanif100
                         10.2.1.0/24 BGP    255 0          RD 10.2.1.2     Vlanif200
                         10.2.1.2/32 BGP    255 0          RD 127.0.0.1    Vlanif200
                        127.0.0.0/8 Direct 0 0            D 127.0.0.1    InLoopBack0
                    255.255.255.255/32 Direct 0 0            D 127.0.0.1     InLoopBack0

                    CE1 and CE2 can ping each other.
                    [CE1] ping 10.2.1.1
                     PING 10.2.1.1: 56 data bytes, press CTRL_C to break
                      Reply from 10.2.1.1: bytes=56 Sequence=1 ttl=254 time=8 ms
                      Reply from 10.2.1.1: bytes=56 Sequence=2 ttl=254 time=3 ms
                      Reply from 10.2.1.1: bytes=56 Sequence=3 ttl=254 time=2 ms
                      Reply from 10.2.1.1: bytes=56 Sequence=4 ttl=254 time=3 ms
                      Reply from 10.2.1.1: bytes=56 Sequence=5 ttl=254 time=2 ms

                     --- 10.2.1.1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 2/3/8 ms


Configuration Scripts
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 100 200
                         #
                         ip vpn-instance vpna
                          ipv4-family
                           route-distinguisher 100:1
                           vpn-target 111:1 export-extcommunity
                           vpn-target 111:1 import-extcommunity
                           vpn-target 222:2 import-extcommunity
                         #
                         ip vpn-instance vpnb
                          ipv4-family
                           route-distinguisher 100:2
                           vpn-target 222:2 export-extcommunity
                           vpn-target 222:2 import-extcommunity
                           vpn-target 111:1 import-extcommunity
                         #
                         interface Vlanif100
                          ip binding vpn-instance vpna
                          ip address 10.1.1.2 255.255.255.0
                         #
                         interface Vlanif200
                          ip binding vpn-instance vpnb
                          ip address 10.2.1.2 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 100 200
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 100 200
                         #
                         bgp 100
                          #
                          ipv4-family unicast


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      51
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         #
                         ipv4-family vpn-instance vpna
                          import-route direct
                         #
                         ipv4-family vpn-instance vpnb
                          import-route direct
                        #
                        return

