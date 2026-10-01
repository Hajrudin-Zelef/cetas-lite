---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-132
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [18680, 18815]
sha256: cf2155a6c04e8d742c37b4774e033e8e6421a62f6105dcb5aead70fc1c2c1d50
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     Peer         V       AS MsgRcvd MsgSent OutQ Up/Down             State PrefRcv

                     2001:DB8:1::1 4       65410      3      3   0 00:00:37 Established    1

         Step 7 Configure static routes between PE1 and CE2.
                    # Configure an IPv6 static route for the VPN instance named vpnb on PE1, and
                    import the route into the routing table of the BGP VPN instance IPv6 address
                    family.
                    [PE1] ipv6 route-static vpn-instance vpnb 2001:db8:8:: 64 2001:db8:3::1
                    [PE1] bgp 100
                    [PE1-bgp] ipv6-family vpn-instance vpnb
                    [PE1-bgp6-vpnb] import-route static
                    [PE1-bgp6-vpnb] quit
                    [PE1-bgp] quit

                    # Configure an IPv6 default route on CE2.
                    [CE2] ipv6 route-static :: 0 2001:db8:3::2

         Step 8 Configure OSPFv3 between PE2 and CE3.
                    # Configure OSPFv3 on PE2.
                    [PE2] ospfv3 1 vpn-instance vpna
                    [PE2-ospfv3-1] router-id 10.10.11.11
                    [PE2-ospfv3-1] area 0.0.0.0
                    [PE2-ospfv3-1-area 0.0.0.0] quit
                    [PE2-ospfv3-1] import-route bgp
                    [PE2-ospfv3-1] quit
                    [PE2] interface Vlanif 200
                    [PE2-Vlanif200] ospfv3 1 area 0
                    [PE2-Vlanif200] quit

                    # Import OSPFv3 routes into BGP on PE2.
                    [PE2] bgp 100
                    [PE2-bgp] ipv6-family vpn-instance vpna
                    [PE2-bgp6-vpna] import-route ospfv3 1
                    [PE2-bgp6-vpna] quit
                    [PE2-bgp] quit

                    # Configure OSPFv3 on CE3.
                    [CE3] ospfv3 1
                    [CE3-ospfv3-1] router-id 22.22.22.22
                    [CE3-ospfv3-1] area 0.0.0.0
                    [CE3-ospfv3-1-area 0.0.0.0] quit
                    [CE3-ospfv3-1] quit
                    [CE3] interface Vlanif100
                    [CE3-Vlanif100] ospfv3 1 area 0
                    [CE3-Vlanif100] quit
                    [CE3] interface LoopBack 1
                    [CE3-LoopBack1] ospfv3 1 area 0
                    [CE3-LoopBack1] quit

                    ----End

Verifying the Configuration
                    After the configuration is complete, the ping operations (with the source address
                    specified in the ping command) between CE1 and CE3 and between CE2 and CE4
                    can succeed. The following example uses the command output on CE1.
                    [CE1] ping ipv6 -a 2001:db8:8::1 2001:db8:9::1
                     PING 2001:db8:9::1 : 56 data bytes, press CTRL_C to break


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                      297
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration

                        Reply from 2001:db8:9::1
                        bytes=56 Sequence=1 hop limit=62 time = 170 ms
                        Reply from 2001:db8:9::1
                        bytes=56 Sequence=2 hop limit=62 time = 140 ms
                        Reply from 2001:db8:9::1
                        bytes=56 Sequence=3 hop limit=62 time = 150 ms
                        Reply from 2001:db8:9::1
                        bytes=56 Sequence=4 hop limit=62 time = 140 ms
                        Reply from 2001:db8:9::1
                        bytes=56 Sequence=5 hop limit=62 time = 170 ms

                     --- 2001:db8:9::1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 140/154/170 ms

                    The address 2001:db8:9::1/64 also exists on CE4. To determine whether the
                    forwarding path is the desired one, you only need to run the display ipv6
                    statistics interface command on PE2 to check if the number of ICMPv6 packets
                    sent/received on the interface changes.
                    Run the ping ipv6 -a 2001:db8:8::1 -c 100 2001:db8:9::1 command on CE1 to
                    send 100 IPv6 packets (with the source address specified in the command). Then,
                    repeatedly run the display ipv6 statistics interface Vlanif100 or display ipv6
                    statistics interface Vlanif200 command on PE2 to check the number of ICMPv6
                    packets received/sent on each interface. Data on VLANIF 200 keeps changing,
                    meaning that IPv6 data is forwarded to CE3 that is on the same VPN as CE1 and
                    the VPNs are isolated from each other.

Configuration Scripts
                    ●     PE1
                          #
                          sysname PE1
                          #
                          vlan batch 100 200 300
                          #
                          ip vpn-instance vpna
                           ipv6-family
                            route-distinguisher 100:1
                            vpn-target 22:22 export-extcommunity
                            vpn-target 33:33 import-extcommunity
                          #
                          ip vpn-instance vpnb
                           ipv6-family
                            route-distinguisher 100:3
                            vpn-target 44:44 export-extcommunity
                            vpn-target 55:55 import-extcommunity
                          #
                          mpls lsr-id 1.1.1.9
                          #
                          mpls
                          #
                          mpls ldp
                          #
                          isis 1
                           network-entity 10.1111.1111.1111.00
                          #
                          interface Vlanif100
                           ip binding vpn-instance vpna
                           ipv6 enable
                           ipv6 address 2001:db8:1::2/64
                          #
                          interface Vlanif200
                           ip binding vpn-instance vpnb


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                         298
VPN Configuration
VPN Configuration                                                                            4 IPv6 L3VPN Configuration

