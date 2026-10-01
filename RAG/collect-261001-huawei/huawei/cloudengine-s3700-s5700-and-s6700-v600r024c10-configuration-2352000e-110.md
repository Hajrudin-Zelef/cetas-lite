---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-110
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [15330, 15515]
sha256: ac62ed1e7637263a37c8160cdb65ba2ffc3e1b543db7c1e0aa1924b9f9f9cef4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:1
                          vpn-target 100:1 export-extcommunity
                          vpn-target 200:1 import-extcommunity
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpna
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 20.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        bgp 100
                         peer 2.2.2.9 as-number 100
                         peer 2.2.2.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 2.2.2.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 2.2.2.9 enable
                        #
                         ipv4-family vpn-instance vpna
                          peer 10.1.1.1 as-number 65410
                        #
                        ospf 1
                         area 0.0.0.0
                          network 20.1.1.0 0.0.0.255
                          network 1.1.1.9 0.0.0.0
                        #
                        return
                    ●   Spoke-PE2
                        #
                        sysname Spoke-PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:3
                          vpn-target 100:1 export-extcommunity
                          vpn-target 200:1 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         244
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        mpls ldp
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpna
                         ip address 10.4.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 11.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 100
                         peer 2.2.2.9 as-number 100
                         peer 2.2.2.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 2.2.2.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 2.2.2.9 enable
                        #
                         ipv4-family vpn-instance vpna
                          peer 10.4.1.1 as-number 65420
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 11.1.1.0 0.0.0.255
                        #
                        return

                    ●   Spoke-CE2
                        #
                        sysname Spoke-CE2
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.4.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface Loopback 1
                         ip address 22.22.22.22 255.255.255.255
                        #
                        bgp 65420
                         peer 10.4.1.2 as-number 100
                         #
                         ipv4-family unicast
                          network 22.22.22.22 255.255.255.255
                          peer 10.4.1.2 enable
                        #
                        return

                    ●   Hub-CE



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         245
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        #
                        sysname Hub-CE
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 10.2.1.1 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 10.3.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface Loopback 1
                         ip address 33.33.33.33 255.255.255.255
                        #
                        bgp 65430
                         peer 10.2.1.2 as-number 100
                         peer 10.3.1.2 as-number 100
                         #
                         ipv4-family unicast
                          peer 10.3.1.2 enable
                          network 33.33.33.33 255.255.255.255
                          peer 10.2.1.2 enable
                        #
                        return

