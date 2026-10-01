---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-176
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [25480, 25592]
sha256: 6d192e9a4c0afab52052ff7fbdcc4da511800ea0a3990088051af8e544f1c4c6
---

                    # Configure the Hub-PE.
                    [Hub-PE] bgp 100
                    [Hub-PE-bgp] ipv6-family vpn-instance vpn_in
                    [Hub-PE-bgp-6-vpn_in] peer 2001:db8:3::1 as-number 65430
                    [Hub-PE-bgp-6-vpn_in] quit
                    [Hub-PE-bgp] ipv6-family vpn-instance vpn_out
                    [Hub-PE-bgp-6-vpn_out] peer 2001:db8:4::1 as-number 65430
                    [Hub-PE-bgp-6-vpn_out] peer 2001:db8:4::1 allow-as-loop 1
                    [Hub-PE-bgp-6-vpn_out] quit
                    [Hub-PE-bgp] quit

                    After completing the configuration, run the display bgp vpnv6 all peer command
                    on each PE. The command output shows that BGP peer relationships have been
                    established between the PEs and CEs and are in Established state.
         Step 5 Establish MP-IBGP peer relationships between the PEs.
                    # Configure Spoke-PE1.
                    [Spoke-PE1] bgp 100
                    [Spoke-PE1-bgp] peer 2.2.2.9 as-number 100
                    [Spoke-PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [Spoke-PE1-bgp] ipv6-family vpnv6
                    [Spoke-PE1-bgp-af-vpnv6] peer 2.2.2.9 enable
                    [Spoke-PE1-bgp-af-vpnv6] quit

                    # Configure Spoke-PE2.
                    [Spoke-PE2] bgp 100
                    [Spoke-PE2-bgp] peer 2.2.2.9 as-number 100
                    [Spoke-PE2-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [Spoke-PE2-bgp] ipv6-family vpnv6
                    [Spoke-PE2-bgp-af-vpnv6] peer 2.2.2.9 enable
                    [Spoke-PE2-bgp-af-vpnv6] quit

                    # Configure the Hub-PE.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          403
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration

                    [Hub-PE] bgp 100
                    [Hub-PE-bgp] peer 1.1.1.9 as-number 100
                    [Hub-PE-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [Hub-PE-bgp] peer 3.3.3.9 as-number 100
                    [Hub-PE-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [Hub-PE-bgp] ipv6-family vpnv6
                    [Hub-PE-bgp-af-vpnv6] peer 1.1.1.9 enable
                    [Hub-PE-bgp-af-vpnv6] peer 3.3.3.9 enable
                    [Hub-PE-bgp-af-vpnv6] quit

                    After completing the configuration, run the display bgp peer or display bgp
                    vpnv6 all peer command on PEs. The command output shows that BGP peer
                    relationships have been established between PEs and are in Established state.

                    ----End

Verifying the Configuration
                    After completing the configuration, configure Spoke-CEs to ping each other. The
                    command output shows that the Spoke-CEs can ping each other.
                    The following example uses the command output on Spoke-CE1.
                    <Spoke-CE1> ping ipv6 -a 2001:db8:11::1 2001:db8:12::2
                     PING 2001:db8:12::2 : 56 data bytes, press CTRL_C to break
                      Reply from 2001:db8:12::2
                      bytes=56 Sequence=1 hop limit=59 time=7 ms
                      Reply from 2001:db8:12::2
                      bytes=56 Sequence=2 hop limit=59 time=3 ms
                      Reply from 2001:db8:12::2
                      bytes=56 Sequence=3 hop limit=59 time=3 ms
                      Reply from 2001:db8:12::2
                      bytes=56 Sequence=4 hop limit=59 time=3 ms
                      Reply from 2001:db8:12::2
                      bytes=56 Sequence=5 hop limit=59 time=3 ms

                     ---2001:db8:12::2 ping statistics---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max=3/3/7 ms

                    Run the display bgp ipv6 routing-table command on the Spoke-CEs. The
                    command output shows that there are repetitive AS numbers in the AS_Path
                    attributes of the BGP routes to the peer Spoke-CE.
                    The following example uses the command output on Spoke-CE1.
                    <Spoke-CE1> display bgp ipv6 routing-table
                     BGP Local router ID is 1.1.1.1
                     Status codes: * - valid, > - best, d - damped, x - best external,
                              h - history, i - internal, s - suppressed, S - Stale
                              Origin : i - IGP, e - EGP, ? - incomplete


                    Total Number of Routes: 3
                    *> Network : 2001:db8:11::1                         PrefixLen : 128
                        NextHop : ::                              LocPrf :
                        MED     :0                                PrefVal : 0
                        Label :
                        Path/Ogn : i
                    *> Network : 2001:db8:12::2                        PrefixLen : 128
                        NextHop : 2001:db8:1::2                       LocPrf :
                        MED     :                                PrefVal : 0
                        Label :
                        Path/Ogn : 100 65430 100 65420i
                    *> Network : 2001:db8:13::3                         PrefixLen : 128


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                  404
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        NextHop : 2001:db8:1::2                  LocPrf   :
                        MED     :                           PrefVal : 0
                        Label :
                        Path/Ogn : 100 65430i


