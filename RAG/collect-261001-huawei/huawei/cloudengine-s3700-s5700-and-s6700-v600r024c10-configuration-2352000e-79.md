---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-79
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [10834, 10979]
sha256: 45698a4e46bfb47e21082c82527e5cd021eee6dbe9255971d3ba48f0588f5589
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 3 Establish an MP-IBGP peer relationship between the PE and ASBR in the same AS.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.9 as-number 100
                    [PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [PE1-bgp] ipv4-family vpnv4


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          172
VPN Configuration
VPN Configuration                                                                            3 IPv4 L3VPN Configuration

                    [PE1-bgp-af-vpnv4] peer 2.2.2.9 enable
                    [PE1-bgp-af-vpnv4] quit

                    The configuration of PE2 is similar to the configuration of PE1. For detailed
                    configurations, see Configuration Scripts.
                    # Configure ASBR1.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 1.1.1.9 as-number 100
                    [ASBR1-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [ASBR1-bgp] ipv4-family vpnv4
                    [ASBR1-bgp-af-vpnv4] peer 1.1.1.9 enable
                    [ASBR1-bgp-af-vpnv4] quit

                    The configuration of ASBR2 is similar to the configuration of ASBR1. For detailed
                    configurations, see Configuration Scripts.
                    After the configuration is complete, run the display bgp vpnv4 all peer command
                    on PEs or ASBRs. The command output shows that an MP-IBGP peer relationship
                    has been established between the PE and ASBR in the same AS. The following
                    example uses the command output on PE1.
                    <PE1> display bgp vpnv4 all peer
                     BGP local router ID : 1.1.1.9
                     Local AS number : 100
                     Total number of peers : 1     Peers in established state : 1

                     Peer         V       AS MsgRcvd MsgSent OutQ Up/Down               State PrefRcv

                     2.2.2.9      4      100      54     59    0 00:45:03 Established   2

         Step 4 Create VPN instances on PEs and ASBRs and bind PE interfaces connected to CEs
                to the corresponding VPN instances.
                    # Configure PE1.
                    [PE1] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [PE1-vpn-instance-vpna-af-ipv4] vpn-target 1:1 both
                    [PE1-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] vlan batch 100 200
                    [PE1] interface 10GE1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 200
                    [PE1-10GE1/0/2] quit
                    [PE1] interface Vlanif 200
                    [PE1-Vlanif200] ip binding vpn-instance vpna
                    [PE1-Vlanif200] ip address 10.1.1.2 24
                    [PE1-Vlanif200] quit

                    The configurations of PE2, ASBR1, and ASBR2 are similar to the configuration of
                    PE1. For detailed configurations, see Configuration Scripts.
                    After completing the configuration, run the display ip vpn-instance verbose
                    command on PEs or ASBRs to check the VPN instance configuration. The following
                    example uses the command output on PE1.
                    <PE1> display ip vpn-instance verbose
                     Total VPN-Instances configured : 1
                     Total IPv4 VPN-Instances configured : 1
                     Total IPv6 VPN-Instances configured : 0

                    VPN-Instance Name and ID : vpna, 1
                    Interfaces : Vlanif200


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                    173
VPN Configuration
VPN Configuration                                                                        3 IPv4 L3VPN Configuration

                    Address family ipv4
                    Create date : 2009/09/18 11:30:35
                    Up time : 0 days, 00 hours, 05 minutes and 19 seconds
                    Vrf Status : UP
                    Route Distinguisher : 100:1
                    Export VPN Targets : 1:1
                    Import VPN Targets : 1:1
                    Label policy: label per route
                    The diffserv-mode Information is : uniform
                    The ttl-mode Information is : pipe

         Step 5 Establish EBGP peer relationships between PEs and CEs and between ASBRs and
                CEs, and import loopback routes from CEs.
                    # Configure CE1.
                    [CE1] interface loopback 1
                    [CE1-Loopback1] ip address 11.11.11.11 32
                    [CE1-Loopback1] quit
                    [CE1] bgp 65001
                    [CE1-bgp] peer 10.1.1.2 as-number 100
                    [CE1-bgp] network 11.11.11.11 32
                    [CE1-bgp] quit

                    The configurations of CE2, CE3, and CE4 are similar to the configuration of CE1.
                    For detailed configurations, see Configuration Scripts.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] ipv4-family vpn-instance vpna
                    [PE1-bgp-vpna] peer 10.1.1.1 as-number 65001
                    [PE1-bgp-vpna] quit
                    [PE1-bgp] quit

                    The configurations of PE2, ASBR1, and ASBR2 are similar to the configuration of
                    PE1. For detailed configurations, see Configuration Scripts.
                    After completing the configuration, run the display bgp vpnv4 vpn-instance peer
                    command on PEs or ASBRs to check whether a BGP peer relationship has been
                    established between PEs or ASBRs and their connected CEs. The command output
                    shows that BGP peer relationships have been established between the PEs or
                    ASBRs and their connected CEs and are in the Established state. The following
                    example uses the peer relationship between PE1 and CE1.
                    <PE1> display bgp vpnv4 vpn-instance vpna peer
                     BGP local router ID : 10.1.1.2
                     Local AS number : 100
                     Total number of peers : 1        Peers in established state : 1
                      Peer        V AS MsgRcvd MsgSent OutQ Up/Down State                PrefRcv
                      10.1.1.1    4 65001 11        9      0    00:06:37 Established 1

         Step 6 Enable MPLS on ASBR interfaces connected to each other.
                    [ASBR1] interface 10GE1/0/2
                    [ASBR1-10GE1/0/2] port link-type trunk
                    [ASBR1-10GE1/0/2] port trunk allow-pass vlan 200
                    [ASBR1-10GE1/0/2] quit
                    [ASBR1] interface Vlanif 200
                    [ASBR1-Vlanif200] ip address 10.12.12.1 24
                    [ASBR1-Vlanif200] quit

         Step 7 On ASBR1, establish an MP-EBGP peer relationship with ASBR2, and disable ASBR1
                from filtering received VPNv4 routes based on VPN targets.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 10.12.12.2 as-number 200
                    [ASBR1-bgp] ipv4-family vpnv4


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   174
VPN Configuration
VPN Configuration                                                                                   3 IPv4 L3VPN Configuration

