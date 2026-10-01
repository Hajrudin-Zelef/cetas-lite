---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-39
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [4544, 4700]
sha256: e76b635429718e0310732357929a3b14d0bf6ea00147e7e519019394d5bbfda8
---

                    # Configure PE2.
                    [PE2] ip vpn-instance vpna
                    [PE2-vpn-instance-vpna] ipv4-family
                    [PE2-vpn-instance-vpna-af-ipv4] route-distinguisher 200:1
                    [PE2-vpn-instance-vpna-af-ipv4] vpn-target 111:1 both
                    [PE2-vpn-instance-vpna-af-ipv4] quit
                    [PE2-vpn-instance-vpna] quit
                    [PE2] ip vpn-instance vpnb
                    [PE2-vpn-instance-vpnb] ipv4-family
                    [PE2-vpn-instance-vpnb-af-ipv4] route-distinguisher 200:2
                    [PE2-vpn-instance-vpnb-af-ipv4] vpn-target 222:2 both
                    [PE2-vpn-instance-vpnb-af-ipv4] quit
                    [PE2-vpn-instance-vpnb] quit
                    [PE2] interface 10GE1/0/1
                    [PE2-10GE1/0/1] port link-type trunk
                    [PE2-10GE1/0/1] port trunk allow-pass vlan 100
                    [PE2-10GE1/0/1] quit
                    [PE2] interface Vlanif 100
                    [PE2-Vlanif100] ip binding vpn-instance vpna
                    [PE2-Vlanif100] ip address 10.3.1.2 24
                    [PE2-Vlanif100] quit
                    [PE2] interface 10GE
                    [PE2-10GE1/0/3] port link-type trunk
                    [PE2-10GE1/0/3] port trunk allow-pass vlan 300
                    [PE2-10GE1/0/3] quit
                    [PE2] interface Vlanif 300
                    [PE2-Vlanif300] ip binding vpn-instance vpnb
                    [PE2-Vlanif300] ip address 10.4.1.2 24
                    [PE2-Vlanif300] quit

                    # Configure IP addresses for interfaces on CEs, as shown in Figure 3-11. For
                    detailed configurations, see Configuration Scripts.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           72
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


                    After completing the configuration, run the display ip vpn-instance verbose
                    command on each PE to check the VPN instance configurations. The command
                    output shows that each PE can ping its connected CE.

                          NOTE

                         If a PE has multiple interfaces bound to the same VPN instance, use the -a source-ip-
                         address parameter to specify a source IP address when running the ping -vpn-instance
                         vpn-instance-name -a source-ip-address dest-ip-address command to ping the CE
                         connected to the remote PE. If the source IP address is not specified, the ping operation
                         may fail.

                    The following example uses the command output on PE1.
                    [PE1] display ip vpn-instance verbose
                     Total VPN-Instances configured : 2
                    Total IPv4 VPN-Instances configured : 2
                    Total IPv6 VPN-Instances configured : 0


                    VPN-Instance Name and ID : vpna, 1
                    Interfaces : Vlanif100
                    Address family ipv4
                    Create date : 2009/01/21 11:30:35
                    Up time : 0 days, 00 hours, 05 minutes and 19 seconds
                    Vrf Status : UP
                    Route Distinguisher : 100:1
                    Export VPN Targets : 111:1
                    Import VPN Targets : 111:1
                    Label policy: label per route
                    The diffserv-mode Information is : uniform
                    The ttl-mode Information is : pipe


                     VPN-Instance Name and ID : vpnb, 2
                     Interfaces : Vlanif200
                     Address family ipv4
                     Create date : 2009/01/21 11:31:18
                     Up time : 0 days, 00 hours, 04 minutes and 36 seconds
                     Vrf Status : UP
                     Route Distinguisher : 100:2
                     Export VPN Targets : 222:2
                     Import VPN Targets : 222:2
                     Label policy: label per route
                     The diffserv-mode Information is : uniform
                     The ttl-mode Information is : pipe
                    [PE1] ping -vpn-instance vpna 10.1.1.1
                     PING 10.1.1.1: 56 data bytes, press CTRL_C to break
                       Reply from 10.1.1.1: bytes=56 Sequence=1 ttl=255 time=56 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=2 ttl=255 time=4 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=3 ttl=255 time=4 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=4 ttl=255 time=52 ms
                    Reply from 10.1.1.1: bytes=56 Sequence=5 ttl=255 time=3 ms


                     --- 10.1.1.1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 3/23/56 ms

         Step 4 Establish an EBGP peer relationship between each PE and its connected CE.
                    # Configure CE1.
                    [CE1] interface loopback 1
                    [CE1-LoopBack1] ip address 11.11.11.11 32
                    [CE1-LoopBack1] quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                        73
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration

                    [CE1] bgp 65410
                    [CE1-bgp] peer 10.1.1.2 as-number 100
                    [CE1-bgp] network 11.11.11.11 32
                    [CE1-bgp] quit

                          NOTE

                        The configurations of CE2, CE3, and CE4 are similar to the configuration of CE1. For
                        detailed configurations, see Configuration Scripts.

                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] ipv4-family vpn-instance vpna
                    [PE1-bgp-vpna] peer 10.1.1.1 as-number 65410
                    [PE1-bgp-vpna] import-route direct
                    [PE1-bgp-vpna] quit
                    [PE1-bgp] ipv4-family vpn-instance vpnb
                    [PE1-bgp-vpnb] peer 10.2.1.1 as-number 65420
                    [PE1-bgp-vpnb] import-route direct
                    [PE1-bgp-vpnb] quit
                    [PE1-bgp] quit

                          NOTE

                        The configuration of PE2 is similar to the configuration of PE1. For detailed configurations,
                        see Configuration Scripts.

                    After the configuration is complete, run the display bgp vpnv4 vpn-instance vpn-
                    instance-name peer command on PEs to check BGP peer relationships between
                    PEs and CEs. The command output shows that the BGP peer relationships are in
                    the Established state.
                    The following example uses the peer relationship between PE1 and CE1.
                    [PE1] display bgp vpnv4 vpn-instance vpna peer
                     Status codes: * - Dynamic
                     BGP local router ID : 1.1.1.9
                     Local AS number : 100

                    VPN-Instance vpna, Router ID 1.1.1.9:
                    Total number of peers : 1
                    Peers in established state : 1
                    Total number of dynamic peers : 0
                     Peer         V AS MsgRcvd MsgSent        OutQ Up/Down State       PrefRcv
                     10.1.1.1     4 65410 11       9      0   00:06:37 Established 1


