---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-89
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [12332, 12461]
sha256: c6b656f71553783023aaebeaf1266e3fbda96093fe8331bbd849942ca7badfa6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The configuration of SPE2 is similar to that of SPE1. Note that on SPE2, the local
                    preference of routes advertised to UPEs is 190.
                    # Configure NPE1.
                    [NPE1] interface loopback 2
                    [NPE1-LoopBack2] ip binding vpn-instance vpna
                    [NPE1-LoopBack2] ip address 55.55.55.55 32
                    [NPE1-LoopBack2] quit

                    The configuration of NPE2 is similar to that of NPE1. For detailed configurations,
                    see Configuration Scripts.
         Step 7 Configure a route-policy to adjust the local preference of the primary and backup
                routes.
                    # Configure SPE1.
                    [SPE1] ip ip-prefix default permit 0.0.0.0 0
                    [SPE1] route-policy NPE1 deny node 10
                    [SPE1-route-policy] if-match ip-prefix default
                    [SPE1-route-policy] quit
                    [SPE1] route-policy NPE1 permit node 20
                    [SPE1-route-policy] apply local-preference 200
                    [SPE1-route-policy] quit
                    [SPE1] route-policy NPE2 deny node 10
                    [SPE1-route-policy] if-match ip-prefix default
                    [SPE1-route-policy] quit
                    [SPE1] route-policy NPE2 permit node 20
                    [SPE1-route-policy] apply local-preference 190
                    [SPE1-route-policy] quit
                    [SPE1] bgp 100
                    [SPE1-bgp] ipv4-family vpnv4


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     198
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    [SPE1-bgp-af-vpnv4] peer 5.5.5.5 route-policy NPE1 import
                    [SPE1-bgp-af-vpnv4] peer 6.6.6.6 route-policy NPE2 import
                    [SPE1-bgp-af-vpnv4] peer 1.1.1.1 ip-prefix default export
                    [SPE1-bgp-af-vpnv4] peer 2.2.2.2 ip-prefix default export
                    [SPE1-bgp-af-vpnv4] quit
                    [SPE1-bgp] quit

                    The configuration of SPE2 is similar to that of SPE1. Note that on SPE2, the local
                    preference of routes advertised to UPEs is 190.
                    # Configure NPE1.
                    [NPE1] route-policy SPE1 permit node 10
                    [NPE1-route-policy] apply local-preference 200
                    [NPE1-route-policy] quit
                    [NPE1] route-policy SPE2 permit node 10
                    [NPE1-route-policy] apply local-preference 190
                    [NPE1-route-policy] quit
                    [NPE1] bgp 100
                    [NPE1-bgp] ipv4-family vpnv4
                    [NPE1-bgp-af-vpnv4] peer 3.3.3.3 route-policy SPE1 import
                    [NPE1-bgp-af-vpnv4] peer 4.4.4.4 route-policy SPE2 import
                    [NPE1-bgp-af-vpnv4] quit
                    [NPE1-bgp] quit

                    The configuration of NPE2 is similar to that of NPE1. On NPE2, configure the local
                    preference as 180 for routes imported from SPE1 and 170 for routes imported
                    from SPE2.
         Step 8 Configure VPN FRR on UPEs, SPEs, and NPEs to enhance network reliability. The
                following uses UPE1 as an example.
                    [UPE1] bgp 100
                    [UPE1-bgp] ipv4-family vpn-instance vpna
                    [UPE1-bgp-vpna] auto-frr
                    [UPE1-bgp-vpna] route-select delay 300
                    [UPE1-bgp-vpna] quit
                    [UPE1-bgp] quit

                    ----End

Verifying the Configuration
                    After the configuration is complete, UPE1 does not have any route to the EPC side,
                    but has a default route with SPE1 as the next hop. NPE1 has a BGP route to the
                    eNodeB. The CE and UPE1 can ping each other.
                    <UPE1> display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 5        Routes : 5

                    Destination/Mask     Proto Pre Cost       Flags NextHop         Interface

                          0.0.0.0/0 IBGP 255 0            RD 3.3.3.3     Vlanif100
                         127.0.0.0/8 Direct 0 0           D 127.0.0.1     InLoopBack0
                         10.1.1.0/24 Direct 0 0           D 10.1.1.2     Vlanif300
                         10.1.1.2/32 Direct 0 0           D 127.0.0.1     Vlanif300
                        10.1.1.255/32 Direct 0 0           D 127.0.0.1     Vlanif300
                    255.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0
                    <UPE1> ping -vpn-instance vpna 10.3.1.1
                     PING 10.3.1.1: 56 data bytes, press CTRL_C to break
                       Reply from 10.3.1.1: bytes=56 Sequence=1 ttl=251 time=5 ms
                       Reply from 10.3.1.1: bytes=56 Sequence=2 ttl=253 time=3 ms
                       Reply from 10.3.1.1: bytes=56 Sequence=3 ttl=251 time=3 ms


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         199
VPN Configuration
VPN Configuration                                                                                   3 IPv4 L3VPN Configuration

                      Reply from 10.3.1.1: bytes=56 Sequence=4 ttl=253 time=2 ms
                      Reply from 10.3.1.1: bytes=56 Sequence=5 ttl=251 time=2 ms

                      --- 10.3.1.1 ping statistics ---
                        5 packet(s) transmitted
                        5 packet(s) received
                        0.00% packet loss
                        round-trip min/avg/max = 2/3/5 ms
                    <NPE1> display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 9        Routes : 9

                    Destination/Mask     Proto Pre Cost        Flags NextHop            Interface

                         0.0.0.0/0 IBGP 255 0               RD 3.3.3.3     Vlanif100
                        10.1.1.0/24 IBGP 255 0               RD 3.3.3.3     Vlanif100
                        10.4.1.0/24 Direct 0 0              D 10.4.1.1     Vlanif300
                        10.4.1.1/32 Direct 0 0              D 127.0.0.1    Vlanif300
                       10.4.1.255/32 Direct 0 0              D 127.0.0.1    Vlanif300
                        10.2.1.0/24 EBGP 255 0               RD 10.4.1.2     Vlanif300
                        10.3.1.0/24 EBGP 255 0               RD 10.4.1.2     Vlanif300
                         7.7.7.7/32 EBGP 255 0               RD 10.4.1.2     Vlanif300
                    255.255.255.255/32 Direct 0 0              D 127.0.0.1     InLoopBack0

