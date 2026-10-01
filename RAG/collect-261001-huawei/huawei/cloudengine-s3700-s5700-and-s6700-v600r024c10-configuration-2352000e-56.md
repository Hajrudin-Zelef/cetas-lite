---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-56
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [7207, 7325]
sha256: 489067e9c999746cd3807e45580663f1ee30ee36e16eb67f7512d5d4ef253007
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Run the display bgp vpnv4 all peer command on each PE. The command output
                    shows that the status of the MP-IBGP peer relationship between PEs is
                    Established.
                    The following example uses the command output on PE1.
                    <PE1> display bgp vpnv4 all peer
                     BGP local router ID : 1.1.1.1
                     Local AS number : 100
                     Total number of peers : 2            Peers in established state : 2

                     Peer         V        AS MsgRcvd MsgSent OutQ Up/Down        State PrefRcv
                     2.2.2.2      4       100    43   30   0 00:21:55 Established    1
                     3.3.3.3      4       100    36   25   0 00:18:12 Established    1

         Step 7 Enable VPN FRR.
                    [PE1] bgp 100
                    [PE1-bgp] ipv4-family vpn-instance vpn1
                    [PE1-bgp-vpn1] route-select delay 300
                    [PE1-bgp-vpn1] quit
                    [PE1-bgp] quit
                    [PE1] ip vpn-instance vpn1
                    [PE1-vpn-instance-vpn1] ipv4-family
                    [PE1-vpn-instance-vpn1-af-ipv4] vpn frr
                    [PE1-vpn-instance-vpn1-af-ipv4] quit
                    [PE1-vpn-instance-vpn1] quit

                    ----End

Verifying the Configuration
                    After completing the configuration, check information about the backup next hop,
                    backup label, and backup tunnel ID on each PE.
                    The following example uses the command output on PE1.
                    <PE1> display ip routing-table vpn-instance vpn1 11.11.11.11 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       115
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration

                    Summary Count : 1

                    Destination: 11.11.11.11/32
                       Protocol: IBGP         Process ID: 0
                     Preference: 255                Cost: 0
                        NextHop: 2.2.2.2        Neighbour: 0.0.0.0
                         State: Active Adv Relied      Age: 00h08m28s
                          Tag: 0             Priority: low
                         Label: 4098            QoSInfo: 0x0
                     IndirectID: 0x6400006D
                    RelayNextHop: 10.10.1.2         Interface: Vlanif200
                       TunnelID: 0x0000000001004c4b42 Flags: RD
                      BkNextHop: 3.3.3.3        BkInterface: Vlanif300
                        BkLabel: 4098         SecTunnelID: 0x0
                    BkPETunnelID: 0x0000000001004c4b43 BkPESecTunnelID: 0x0
                    BkIndirectID: 0x6400006F


Configuration Scripts
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 200 300
                         #
                         ip vpn-instance vpn1
                          ipv4-family
                           route-distinguisher 100:1
                           vpn-target 111:1 export-extcommunity
                           vpn-target 111:1 import-extcommunity
                           vpn frr
                         #
                         mpls lsr-id 1.1.1.1
                         #
                         mpls
                         #
                         mpls ldp
                         #
                         interface Vlanif200
                          ip address 10.10.1.1 255.255.255.252
                          mpls
                          mpls ldp
                         #
                         interface Vlanif300
                          ip address 10.20.1.1 255.255.255.252
                          mpls
                          mpls ldp
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 200
                         #
                         interface 10GE1/0/3
                          port link-type trunk
                          port trunk allow-pass vlan 300
                         #
                         interface LoopBack1
                          ip address 1.1.1.1 255.255.255.255
                         #
                         bgp 100
                          peer 2.2.2.2 as-number 100
                          peer 2.2.2.2 connect-interface LoopBack1
                          peer 3.3.3.3 as-number 100
                          peer 3.3.3.3 connect-interface LoopBack1
                          #
                          ipv4-family unicast
                           peer 2.2.2.2 enable
                           peer 3.3.3.3 enable
                          #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         116
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

