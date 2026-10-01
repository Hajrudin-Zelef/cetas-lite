---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-113
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [15811, 15924]
sha256: 44ce255b0e9ba36db696084c6c366681d68f1f5d079e4c9b2ecb374bf445ab0a
---

                    # Configure the Hub-PE.
                    [Hub-PE] bgp 100
                    [Hub-PE-bgp] ipv4-family vpn-instance vpnhub
                    [Hub-PE-bgp-vpnhub] peer 10.2.1.1 as-number 65430
                    [Hub-PE-bgp-vpnhub] quit
                    [Hub-PE-bgp] quit

                    After completing the configuration, run the display bgp vpnv4 all peer command
                    on each PE. The command output shows that BGP peer relationships have been
                    established between the PEs and CEs and are in Established state.

         Step 5 Establish MP-IBGP peer relationships between the PEs.

                    # Configure Spoke-PE1.
                    [Spoke-PE1] bgp 100
                    [Spoke-PE1-bgp] peer 2.2.2.9 as-number 100
                    [Spoke-PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [Spoke-PE1-bgp] ipv4-family vpnv4
                    [Spoke-PE1-bgp-af-vpnv4] peer 2.2.2.9 enable
                    [Spoke-PE1-bgp-af-vpnv4] quit

                    # Configure Spoke-PE2.
                    [Spoke-PE2] bgp 100
                    [Spoke-PE2-bgp] peer 2.2.2.9 as-number 100
                    [Spoke-PE2-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [Spoke-PE2-bgp] ipv4-family vpnv4
                    [Spoke-PE2-bgp-af-vpnv4] peer 2.2.2.9 enable
                    [Spoke-PE2-bgp-af-vpnv4] quit

                    # Configure the Hub-PE.
                    [Hub-PE] bgp 100
                    [Hub-PE-bgp] peer 1.1.1.9 as-number 100
                    [Hub-PE-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [Hub-PE-bgp] peer 3.3.3.9 as-number 100
                    [Hub-PE-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [Hub-PE-bgp] ipv4-family vpnv4
                    [Hub-PE-bgp-af-vpnv4] peer 1.1.1.9 enable
                    [Hub-PE-bgp-af-vpnv4] peer 3.3.3.9 enable
                    [Hub-PE-bgp-af-vpnv4] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          251
VPN Configuration
VPN Configuration                                                                                3 IPv4 L3VPN Configuration


                    After completing the configuration, run the display bgp peer or display bgp
                    vpnv4 all peer command on PEs. The command output shows that BGP peer
                    relationships have been established between PEs and are in the Established state.

                    ----End

Verifying the Configuration
                    After completing the configuration, configure Spoke-CEs to ping each other. The
                    command output shows that the Spoke-CEs can ping each other.
                    The following example uses the command output on Spoke-CE1.
                    <Spoke-CE1> ping -a 11.11.11.11 22.22.22.22
                     PING 22.22.22.22: 56 data bytes, press CTRL_C to break
                        Reply from 22.22.22.22: bytes=56 Sequence=1 ttl=250ime=80 ms
                        Reply from 22.22.22.22: bytes=56 Sequence=2 ttl=250ime=129 ms
                        Reply from 22.22.22.22: bytes=56 Sequence=3 ttl=250 time=132 ms
                        Reply from 22.22.22.22: bytes=56 Sequence=4 ttl=250 time=92 ms
                        Reply from 22.22.22.22: bytes=56 Sequence=5 ttl=250 time=126 ms
                      --- 22.22.22.22 ping statistics ---
                        5 packet(s) transmitted
                        5 packet(s) received
                        0.00% packet loss
                        round-trip min/avg/max = 80/111/132 ms

                    Run the display bgp routing-table command on each Spoke-CE. The command
                    output shows BGP routing table information.
                    The following example uses the command output on Spoke-CE1.
                    <Spoke-CE1> display bgp routing-table

                    BGP Local router ID is 10.1.1.1
                    Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                              h - history, i - internal, s - suppressed, S - Stale
                              Origin : i - IGP, e - EGP, ? - incomplete
                    RPKI validation codes: V - valid, I - invalid, N - not-found


                    Total Number of Routes: 2
                         Network        NextHop                     MED        LocPrf     PrefVal Path/Ogn

                    *>   0.0.0.0/0      10.1.1.2                                  0       100 65430i
                    *>   11.11.11.11/32    0.0.0.0                   0                0      i


Configuration Scripts
                    ●    Spoke-CE1
                         #
                         sysname Spoke-CE1
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
                         interface Loopback 1
                          ip address 11.11.11.11 255.255.255.255
                         #
                         bgp 65410


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          252
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

