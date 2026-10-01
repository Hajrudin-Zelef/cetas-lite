---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-62
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [8124, 8266]
sha256: 6c4879d6e99c65befb07a8c25947adc513befa886928cbcb0eb5551e9f7325f1
---

                    # Configure PE3.
                    [PE3] ip vpn-instance vpn1
                    [PE3-vpn-instance-vpn1] ipv4-family
                    [PE3-vpn-instance-vpn1-af-ipv4] route-distinguisher 100:2
                    [PE3-vpn-instance-vpn1-af-ipv4] vpn-target 111:1
                    [PE3-vpn-instance-vpn1-af-ipv4] quit
                    [PE3-vpn-instance-vpn1] quit
                    [PE3] interface Vlanif200
                    [PE3-Vlanif200] ip binding vpn-instance vpn1
                    [PE3-Vlanif200] ip address 192.168.2.1 30
                    [PE3-Vlanif200] quit

         Step 6 Configure an OSPF instance on PE2 and the CE and establish an EBGP peer
                relationship between PE3 and the CE.
                    # Configure PE2.
                    [PE2] ospf 2 vpn-instance vpn1
                    [PE2-ospf-2] import-route bgp
                    [PE2-ospf-2] area 1
                    [PE2-ospf-2-area-0.0.0.1] network 192.168.1.0 0.0.0.3
                    [PE2-ospf-2-area-0.0.0.1] quit
                    [PE2-ospf-2] quit
                    [PE2] bgp 100
                    [PE2-bgp] ipv4-family vpn-instance vpn1
                    [PE2-bgp-vpn1] import-route ospf 2
                    [PE2-bgp-vpn1] quit
                    [PE2-bgp] quit

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] ipv4-family vpn-instance vpn1
                    [PE3-bgp-vpn1] peer 192.168.2.2 as-number 65410
                    [PE3-bgp-vpn1] peer 192.168.2.2 preferred-value 600
                    [PE3-bgp-vpn1] bestroute as-path-ignore
                    [PE3-bgp-vpn1] quit
                    [PE3-bgp] quit

                    # Configure the CE.
                    [CE] bgp 65410
                    [CE-bgp] peer 192.168.2.1 as-number 100
                    [CE-bgp] network 22.22.22.22 32
                    [CE-bgp] quit
                    [CE] ospf 1
                    [CE-ospf-1] area 1
                    [CE-ospf-1-area-0.0.0.1] network 192.168.1.0 0.0.0.3
                    [CE-ospf-1-area-0.0.0.1] network 22.22.22.22 0.0.0.0
                    [CE-ospf-1-area-0.0.0.1] quit
                    [CE-ospf-1] quit

                    After completing the configuration, run the display ip routing-table vpn-
                    instance vpn1 22.22.22.22 verbose command on PE2. The command output
                    shows that PE2 has learned routes to the loopback interface on the CE.
                    <PE2> display ip routing-table vpn-instance vpn1 22.22.22.22 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         129
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    Summary Count : 2

                    Destination: 22.22.22.22/32
                       Protocol: OSPF         Process ID: 2
                     Preference: 10               Cost: 1
                        NextHop: 192.168.1.2      Neighbour: 0.0.0.0
                         State: Active Adv           Age: 00h11m08s
                          Tag: 0            Priority: medium
                         Label: NULL            QoSInfo: 0x0
                     IndirectID: 0x76
                    RelayNextHop: 0.0.0.0        Interface: Vlanif200
                       TunnelID: 0x0             Flags: D

                    Destination: 22.22.22.22/32
                       Protocol: IBGP         Process ID: 0
                     Preference: 255                Cost: 0
                        NextHop: 3.3.3.3         Neighbour: 0.0.0.0
                         State: Inactive Adv          Age: 00h13m25s
                          Tag: 0             Priority: low
                         Label: 0x23            QoSInfo: 0x0
                     IndirectID: 0xb7
                    RelayNextHop: 10.11.1.2         Interface: Vlanif300
                       TunnelID: 0x0000000001004c4c62 Flags: R

                    The command output shows that PE2 has learned from the CE through OSPF and
                    from PE3 through BGP the routes to the loopback interface on the CE. Because
                    OSPF takes precedence over BGP, PE2 preferentially selects the route learned
                    through OSPF.
         Step 7 Enable IP auto FRR for the VPN instance IPv4 address family on PE2.
                    # Configure PE2.
                    [PE2] ip vpn-instance vpn1
                    [PE2-vpn-instance-vpn1] ipv4-family
                    [PE2-vpn-instance-vpn1-af-ipv4] ip frr
                    [PE2-vpn-instance-vpn1-af-ipv4] quit
                    [PE2-vpn-instance-vpn1] quit

         Step 8 Enable BGP auto FRR for the BGP VPN instance IPv4 address family on PE3.
                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] ipv4-family vpn-instance vpn1
                    [PE3-bgp-vpn1] auto-frr
                    [PE3-bgp-vpn1] route-select delay 300
                    [PE3-bgp-vpn1] quit
                    [PE3-bgp] quit

                    ----End

Verifying the Configuration
                    After completing the configuration, run the display ip routing-table vpn-
                    instance verbose command on PE2 and PE3 to check the VPN instance routing
                    table.
                    <PE2> display ip routing-table vpn-instance vpn1 22.22.22.22 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                    Summary Count : 2

                    Destination: 22.22.22.22/32
                       Protocol: OSPF         Process ID: 2


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         130
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                     Preference: 10              Cost: 1
                        NextHop: 192.168.1.2       Neighbour: 0.0.0.0
                         State: Active Adv         Age: 00h26m40s
                          Tag: 0           Priority: medium
                         Label: NULL           QoSInfo: 0x0
                     IndirectID: 0x76
                    RelayNextHop: 0.0.0.0       Interface: Vlanif200
                       TunnelID: 0x0            Flags: D
                      BkNextHop: 3.3.3.3       BkInterface: Vlanif300
                        BkLabel: 0x23        SecTunnelID: 0x0
                    BkPETunnelID: 0x0000000001004c4c62 BkPESecTunnelID: 0x0
                    BkIndirectID: 0xb7

