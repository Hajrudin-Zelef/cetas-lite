---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-63
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [8267, 8400]
sha256: 88cc87e6e0a4d7e8f34ccaf81f107ba494f9caef444c0996e5e6bd166bf30276
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Destination: 22.22.22.22/32
                        Protocol: IBGP            Process ID: 0
                      Preference: 255                  Cost: 0
                         NextHop: 3.3.3.3           Neighbour: 0.0.0.0
                          State: Inactive Adv            Age: 00h28m57s
                            Tag: 0              Priority: low
                          Label: 0x23              QoSInfo: 0x0
                      IndirectID: 0xb7
                     RelayNextHop: 10.11.1.2           Interface: Vlanif300
                        TunnelID: 0x0000000001004c4c62 Flags: R
                    <PE3> display ip routing-table vpn-instance vpn1 22.22.22.22 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                    Summary Count : 1

                    Destination: 22.22.22.22/32
                       Protocol: IBGP         Process ID: 0
                     Preference: 255                Cost: 0
                        NextHop: 192.168.2.2       Neighbour: 0.0.0.0
                         State: Active Adv Relied      Age: 00h00m31s
                          Tag: 0             Priority: low
                         Label: NULL             QoSInfo: 0x0
                     IndirectID: 0xa9
                    RelayNextHop: 192.168.2.2        Interface: Vlanif200
                       TunnelID: 0x0              Flags: RD
                      BkNextHop: 2.2.2.2        BkInterface: Vlanif300
                        BkLabel: 0x27          SecTunnelID: 0x5000098
                    BkPETunnelID: 0x0         BkPESecTunnelID: 0x0
                    BkIndirectID: 0xaa

                    The command output shows that after IP FRR is enabled, both PE2 and PE3 have
                    the primary and backup routes to the loopback interface on the CE, and the
                    backup route recurses to an LDP LSP.

                    Run the shutdown and then display ip routing-table vpn-instance verbose
                    commands on VLANIF 200 of PE2. The command output shows that the next hop
                    to the loopback interface on the CE is changed to PE3.
                    <PE2> display ip routing-table vpn-instance vpn1 22.22.22.22 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                    Summary Count : 1

                    Destination: 22.22.22.22/32
                       Protocol: IBGP         Process ID: 0
                     Preference: 255                Cost: 0
                        NextHop: 3.3.3.3        Neighbour: 0.0.0.0
                         State: Active Adv Relied      Age: 00h33m16s
                          Tag: 0             Priority: low
                         Label: 0x23            QoSInfo: 0x0
                     IndirectID: 0xb7


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         131
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                    RelayNextHop: 10.11.1.2   Interface:Vlanif300
                      TunnelID: 0x0000000001004c4c62 Flags: RD

                    Perform the same operations on PE3. The command output shows similar
                    information.
                    It can be concluded that IP+VPNv4 hybrid FRR has taken effect on PE2 and PE3.

Configuration Scripts
                    ●   PE1
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
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.1.1 255.255.255.252
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
                        interface LoopBack2
                         ip binding vpn-instance vpn1
                         ip address 11.11.11.11 255.255.255.255
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
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 2.2.2.2 enable
                          peer 3.3.3.3 enable
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         132
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

