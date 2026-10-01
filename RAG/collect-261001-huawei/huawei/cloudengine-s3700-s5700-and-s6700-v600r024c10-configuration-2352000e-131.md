---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-131
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [18526, 18679]
sha256: faed0477a291d7b96cc6f3ac4ce6e40aa1111e077aba64887cd9f412b225ad4f
---

                    # On PE1, configure an IPv6-address-family-enabled VPN instance named vpnb.
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv6-family
                    [PE1-vpn-instance-vpnb-af-ipv6] route-distinguisher 100:3
                    [PE1-vpn-instance-vpnb-af-ipv6] vpn-target 44:44 export-extcommunity
                    [PE1-vpn-instance-vpnb-af-ipv6] vpn-target 55:55 import-extcommunity
                    [PE1-vpn-instance-vpnb-af-ipv6] quit
                    [PE1-vpn-instance-vpnb] quit


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                               294
VPN Configuration
VPN Configuration                                                                4 IPv6 L3VPN Configuration


                    # Bind the interface that directly connects PE1 to CE2 to the VPN instance named
                    vpnb.
                    [PE1] interface Vlanif 200
                    [PE1-Vlanif200] ip binding vpn-instance vpnb
                    [PE1-Vlanif200] ipv6 enable
                    [PE1-Vlanif200] ipv6 address 2001:db8:3::2 64
                    [PE1-Vlanif200] quit

                    The configuration of PE2 is similar to the configuration of PE1. For detailed
                    configurations, see Configuration Scripts.
                    After completing the configuration, run the display ip vpn-instance verbose
                    command on each PE to check VPN instance configuration. The command output
                    shows that each PE can ping its connected CE. The following example uses the
                    command output on PE1.
                    [PE1] display ip vpn-instance verbose
                    Total VPN-Instances configured : 2
                     Total IPv4 VPN-Instances configured : 0
                     Total IPv6 VPN-Instances configured : 2

                    VPN-Instance Name and ID : vpna, 1
                    Interfaces : Vlanif100
                    Address family ipv6
                    Create date : 2010/07/20 12:31:47
                    Up time : 0 days, 04 hours, 37 minutes and 05 seconds
                    Vrf Status : UP
                    Route Distinguisher : 100:1
                    Export VPN Targets : 22:22
                    Import VPN Targets : 33:33
                    Label Policy : label per route
                    The diffserv-mode Information is : uniform
                    The ttl-mode Information is : pipe

                     VPN-Instance Name and ID : vpnb, 2
                     Interfaces : Vlanif200
                     Address family ipv6
                     Create date : 2010/07/20 14:41:46
                     Up time : 0 days, 02 hours, 27 minutes and 06 seconds
                     Vrf Status : UP
                     Route Distinguisher : 100:3
                     Export VPN Targets : 44:44
                     Import VPN Targets : 55:55
                     Label Policy : label per route
                     The diffserv-mode Information is : uniform
                     The ttl-mode Information is : pipe
                    [PE1] ping ipv6 vpn-instance vpna 2001:db8:1::1
                     PING 2001:db8:1::1 : 56 data bytes, press CTRL_C to break
                       Reply from 2001:db8:1::1
                       bytes=56 Sequence=1 hop limit=64 time = 20 ms
                       Reply from 2001:db8:1::1
                       bytes=56 Sequence=2 hop limit=64 time = 30 ms
                       Reply from 2001:db8:1::1
                       bytes=56 Sequence=3 hop limit=64 time = 30 ms
                       Reply from 2001:db8:1::1
                       bytes=56 Sequence=4 hop limit=64 time = 1 ms
                       Reply from 2001:db8:1::1
                       bytes=56 Sequence=5 hop limit=64 time = 1 ms

                     --- 2001:db8:1::1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 1/16/30 ms

         Step 5 Establish a VPNv6 peer relationship between PEs.
                    # Configure PE1.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          295
VPN Configuration
VPN Configuration                                                                                4 IPv6 L3VPN Configuration

                    [PE1] bgp 100
                    [PE1-bgp] peer 3.3.3.9 as-number 100
                    [PE1-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [PE1-bgp] ipv6-family vpnv6
                    [PE1-bgp-af-vpnv6] peer 3.3.3.9 enable
                    [PE1-bgp-af-vpnv6] quit
                    [PE1-bgp] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.9 as-number 100
                    [PE2-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [PE2-bgp] ipv6-family vpnv6
                    [PE2-bgp-af-vpnv6] peer 1.1.1.9 enable
                    [PE2-bgp-af-vpnv6] quit
                    [PE2-bgp] quit

                    After completing the configuration, run the display bgp vpnv6 all peer command
                    on each PE to check the VPNv6 peer relationship status. The following example
                    uses the command output on PE1.
                    [PE1] display bgp vpnv6 all peer
                     BGP local router ID : 1.1.1.9
                     Local AS number : 100
                     Total number of peers : 1          Peers in established state : 1

                     Peer        V       AS MsgRcvd MsgSent OutQ Up/Down                 State    PrefRcv

                     3.3.3.9     4      100      4      3   0 00:01:50 Established       0

                    The command output shows that State is Established, indicating that the VPNv6
                    peer relationship between PE1 and PE2 has been established.
         Step 6 Configure BGP4+ on PE1 and CE1.
                    # Configure EBGP on PE1.
                    [PE1] bgp 100
                    [PE1-bgp] ipv6-family vpn-instance vpna
                    [PE1-bgp6-vpna] peer 2001:db8:1::1 as-number 65410
                    [PE1-bgp6-vpna] quit
                    [PE1-bgp] quit

                    # Configure EBGP on CE1.
                    [CE1] bgp 65410
                    [CE1-bgp] router-id 10.10.10.10
                    [CE1-bgp] peer 2001:db8:1::2 as-number 100
                    [CE1-bgp] ipv6-family unicast
                    [CE1-bgp-af-ipv6] network 2001:db8:8:: 64
                    [CE1-bgp-af-ipv6] peer 2001:db8:1::2 enable
                    [CE1-bgp-af-ipv6] import-route direct
                    [CE1-bgp-af-ipv6] quit
                    [CE1-bgp] quit

                    The configurations between PE2 and CE4 are similar to the configurations
                    between PE1 and CE1. For detailed configurations, see Configuration Scripts.
                    After completing the configuration, run the display bgp vpnv6 vpn-instance vpn-
                    instance-name peer command on each PE to check whether the peer relationship
                    is established. The following example uses the command output on PE1.
                    [PE1] display bgp vpnv6 vpn-instance vpna peer
                    BGP local router ID : 1.1.1.9
                     Local AS number : 100
                     Total number of peers : 1        Peers in established state : 1



Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                         296
VPN Configuration
VPN Configuration                                                                              4 IPv6 L3VPN Configuration

