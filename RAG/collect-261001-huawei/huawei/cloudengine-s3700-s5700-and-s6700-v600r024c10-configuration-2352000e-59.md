---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-59
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [7666, 7803]
sha256: c0c17af55b1ea00b060536710a61200ba92521d4550386abff23463e721cad55
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        10.1.1.0/24 Direct 0 0             D 10.1.1.1     Vlanif100
                        10.1.1.1/32 Direct 0 0             D 127.0.0.1    Vlanif100
                       10.1.1.255/32 Direct 0 0             D 127.0.0.1    Vlanif100
                      11.11.11.11/32 EBGP 255 1               RD 10.1.1.2      Vlanif100
                        10.2.1.0/24 Direct 0 0             D 10.2.1.1     Vlanif200
                        10.2.1.1/32 Direct 0 0             D 127.0.0.1    Vlanif200
                       10.2.1.255/32 Direct 0 0             D 127.0.0.1    Vlanif200
                    255.255.255.255/32 Direct 0 0             D 127.0.0.1     InLoopBack0

         Step 6 Configure VPN static routes on the PE.
                    # Configure the PE.
                    [PE] ip route-static vpn-instance vpna 11.11.11.11 255.255.255.255 10.2.1.2

         Step 7 Enable VPN IP FRR on the PE.
                    # Configure the PE.
                    [PE] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE1-vpn-instance-vpna-af-ipv4] ip frr
                    [PE1-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit

                    ----End

Verifying the Configuration
                    After completing the configuration, run the display ip routing-table vpn-
                    instance command on the PE. The command output shows that the next hop of
                    the route to 11.11.11.11/32 is 10.2.1.2, and the route has a backup next hop and a
                    backup outbound interface.
                    <PE> display ip routing-table vpn-instance vpna 11.11.11.11 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                    Summary Count : 2

                    Destination: 11.11.11.11/32
                       Protocol: Static        Process ID: 0
                     Preference: 60               Cost: 0
                        NextHop: 10.2.1.2        Neighbour: 0.0.0.0
                         State: Active Adv Relied      Age: 00h35m31s
                          Tag: 0             Priority: low
                         Label: NULL             QoSInfo: 0x0
                     IndirectID: 0x100015B            Instance:
                    RelayNextHop: 10.2.1.2         Interface: Vlanif200
                       TunnelID: 0x0              Flags: RD
                      BkNextHop: 10.1.1.2         BkInterface: Vlanif100
                        BkLabel: NULL            SecTunnelID: 0x0
                    BkPETunnelID: 0x0          BkPESecTunnelID: 0x0
                    BkIndirectID: 0x100015A

                    Run the shutdown command on VLANIF 200 of CE1 to simulate a link fault.
                    [CE1] interface Vlanif200
                    [CE1-Vlanif200] shutdown
                    [CE1] quit

                    Run the display ip routing-table vpn-instance command on the PE again. The
                    command output shows that the next hop to 11.11.11.11/32 is 10.1.1.2, and there
                    is no backup next hop or outbound interface.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         122
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    <PE> display ip routing-table vpn-instance vpna 11.11.11.11 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                    Summary Count : 1

                    Destination: 11.11.11.11/32
                       Protocol: EBGP            Process ID: 0
                     Preference: 255                  Cost: 1
                        NextHop: 10.1.1.2          Neighbour: 10.1.1.2
                         State: Active Adv Relied        Age: 00h37m18s
                          Tag: 0               Priority: low
                         Label: NULL               QoSInfo: 0x0
                     IndirectID: 0x100015A            Instance:
                    RelayNextHop: 10.1.1.2         Interface: Vlanif100
                       TunnelID: 0x0              Flags: RD
                     RouteColor: 0

                    The preceding information shows that VPN IP FRR has taken effect.

Configuration Scripts
                    ●     PE
                          #
                          sysname PE
                          #
                          vlan batch 100 200
                          #
                          ip vpn-instance vpna
                           ipv4-family
                            route-distinguisher 100:1
                            vpn-target 100:100 export-extcommunity
                            vpn-target 100:100 import-extcommunity
                          #
                          interface Vlanif100
                           ip binding vpn-instance vpna
                           ip address 10.1.1.1 255.255.255.0
                          #
                          interface Vlanif200
                           ip binding vpn-instance vpna
                           ip address 10.2.1.1 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 100
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 200
                          #
                          bgp 100
                           #
                           ipv4-family unicast
                            #
                           ipv4-family vpn-instance vpna
                            auto-frr
                            route-select delay 300
                            peer 10.1.1.2 as-number 65410
                          #
                          ip route-static vpn-instance vpna 11.11.11.11 255.255.255.255 10.2.1.2
                          #
                          return

                    ●     CE1
                          #
                          sysname CE1
                          #


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         123
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

