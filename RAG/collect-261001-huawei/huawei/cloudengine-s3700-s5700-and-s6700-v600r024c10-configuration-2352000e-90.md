---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-90
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [12462, 12584]
sha256: d0ac49e5ce9eb0e842a48d68a25c10b583d0969d5a7e2b4f24d497af93272edc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Run the display bgp vpnv4 all routing-table command on UPE1. The command
                    output shows information about a default route in VPN instance vpna. The next
                    hop of the default route is SPE1.
                    <UPE1> display bgp vpnv4 all routing-table
                     BGP Local router ID is 1.1.1.1
                     Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                              h - history, i - internal, s - suppressed, S - Stale
                              Origin : i - IGP, e - EGP, ? - incomplete


                    Total number of routes from all PE: 4
                    Route Distinguisher: 100:1


                        Network           NextHop        MED        LocPrf       PrefVal Path/Ogn

                    *>i 0.0.0.0      3.3.3.3      0           200        0       i
                    *i             4.4.4.4      0           200      0       i
                    *> 10.1.1.0/24      0.0.0.0     0                    0       ?
                    *> 10.1.1.2/32      0.0.0.0     0                    0       ?

                    VPN-Instance vpna, router ID 1.1.1.1:

                    Total Number of Routes: 4
                       Network        NextHop            MED        LocPrf       PrefVal Path/Ogn

                    *>i 0.0.0.0      3.3.3.3      0           200        0       i
                    *i             4.4.4.4      0           200      0       i
                    *> 10.1.1.0/24      0.0.0.0     0                    0       ?
                    *> 10.1.1.2/32      0.0.0.0     0                    0       ?

                    Run the display ip routing-table vpn-instance vpna 10.3.1.1 verbose command
                    on UPE1. The command output shows information about the backup label and
                    backup tunnel ID of the route to the EPC side.
                    <UPE1> display ip routing-table vpn-instance vpna 10.3.1.1 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                    Summary Count : 1


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                            200
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration


                    Destination: 0.0.0.0/0
                       Protocol: IBGP         Process ID: 0
                     Preference: 255               Cost: 0
                        NextHop: 3.3.3.3        Neighbour: 0.0.0.0
                         State: Active Adv Relied     Age: 00h15m22s
                          Tag: 0            Priority: low
                         Label: 16            QoSInfo: 0x0
                     IndirectID: 0x5200006A
                    RelayNextHop: 3.3.3.3         Interface: Vlanif100
                       TunnelID: 0x0000000001004c4b44 Flags: RD
                      BkNextHop: 4.4.4.4       BkInterface: Vlanif200
                        BkLabel: 16         SecTunnelID: 0x0
                    BkPETunnelID: 0x0000000001004c4b62 BkPESecTunnelID: 0x0
                    BkIndirectID: 0x5200006C


Configuration Scripts
                    ●    UPE1
                         #
                         sysname UPE1
                         #
                         vlan batch 100 200 300
                         #
                         ip vpn-instance vpna
                          ipv4-family
                           route-distinguisher 100:1
                           vpn-target 1:1 import-extcommunity
                           vpn-target 1:1 export-extcommunity
                         #
                         mpls lsr-id 1.1.1.1
                         #
                         mpls
                         #
                         mpls ldp
                         #
                         interface Vlanif100
                          ip address 172.16.3.1 255.255.255.0
                          mpls
                          mpls ldp
                         #
                         interface Vlanif200
                          ip address 172.16.2.1 255.255.255.0
                          mpls
                          mpls ldp
                         #
                         interface Vlanif300
                          ip binding vpn-instance vpna
                          ip address 10.1.1.2 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 100
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
                          router-id 1.1.1.1
                          peer 3.3.3.3 as-number 100
                          peer 3.3.3.3 connect-interface LoopBack1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         201
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

