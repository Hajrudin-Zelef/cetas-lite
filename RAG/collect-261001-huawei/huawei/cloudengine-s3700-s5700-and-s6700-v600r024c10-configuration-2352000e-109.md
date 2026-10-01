---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-109
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [15223, 15329]
sha256: 69010358f2f48bd06597b41cf4d66b31b7401ef1c59c7ca97d3e0f16bd00ad46
---

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

                    After completing the configuration, run the display bgp peer or display bgp
                    vpnv4 all peer command on PEs. The command output shows that BGP peer
                    relationships have been established between PEs and are in the Established state.

                    ----End

Verifying the Configuration
                    After the configuration is complete, the Spoke-CEs can ping each other. Run the
                    tracert command. The command output shows that the traffic between the
                    Spoke-CEs is forwarded through the Hub-CE. You can also deduce the number of
                    forwarding devices between the Spoke-CEs based on the TTL displayed in the
                    command output.
                    The following example uses the command output on Spoke-CE1.
                    <Spoke-CE1> ping -a 11.11.11.11 22.22.22.22
                     PING 22.22.22.22: 56 data bytes, press CTRL_C to break
                       Reply from 22.22.22.22: bytes=56 Sequence=1 ttl=250 time=80 ms
                       Reply from 22.22.22.22: bytes=56 Sequence=2 ttl=250 time=129 ms


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   242
VPN Configuration
VPN Configuration                                                                                3 IPv4 L3VPN Configuration

                        Reply from 22.22.22.22: bytes=56 Sequence=3 ttl=250 time=132 ms
                        Reply from 22.22.22.22: bytes=56 Sequence=4 ttl=250 time=92 ms
                        Reply from 22.22.22.22: bytes=56 Sequence=5 ttl=250 time=126 ms
                      --- 22.22.22.22 ping statistics ---
                        5 packet(s) transmitted
                        5 packet(s) received
                        0.00% packet loss
                        round-trip min/avg/max = 80/111/132 ms
                    <Spoke-CE1> tracert -a 11.11.11.11 22.22.22.22
                    traceroute to 22.22.22.22(22.22.22.22), max hops: 64, packet length: 40, press CTRL_C to break
                     1 10.1.1.2 31 ms 12 ms 6 ms
                     2 10.3.1.2 9 ms 11 ms 24 ms
                     3 10.3.1.1 46 ms 27 ms 9 ms
                     4 10.2.1.2 23 ms 22 ms 28 ms
                     5 10.4.1.2 12 ms 34 ms 23 ms
                     6 22.22.22.22 46 ms 1 ms 9 ms

                    Run the display bgp routing-table command on each Spoke-CE. The command
                    output shows that there are repetitive AS numbers in the AS_Path attributes of the
                    BGP routes to the peer Spoke-CE.
                    The following example uses the command output on Spoke-CE1.
                    <Spoke-CE1> display bgp routing-table

                    BGP Local router ID is 11.11.11.11
                    Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                             h - history, i - internal, s - suppressed, S - Stale
                             Origin : i - IGP, e - EGP, ? - incomplete

                    Total Number of Routes: 3
                         Network         NextHop                    MED        LocPrf PrefVal Path/Ogn
                    *>    11.11.11.11/32  0.0.0.0                   0               0     i
                    *>    22.22.22.22/32   10.1.1.2                             0    100 65430 100 65420i
                    *>    33.33.33.33/32  10.1.1.2                                  0    100 65430i


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
                          peer 10.1.1.2 as-number 100
                          #
                          ipv4-family unicast
                           network 11.11.11.11 255.255.255.255
                           peer 10.1.1.2 enable
                         #
                         return
                    ●    Spoke-PE1
                         #
                         sysname Spoke-PE1
                         #
                         vlan batch 100 200


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          243
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

