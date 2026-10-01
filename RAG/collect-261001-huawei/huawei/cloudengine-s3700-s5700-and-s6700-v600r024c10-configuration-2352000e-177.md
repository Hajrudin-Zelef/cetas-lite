---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-177
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [25593, 25781]
sha256: d6a1d862b312b5bdafca54061afd6994c360622945b0c91e2c7e5ae67c0e734f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●   Spoke-CE1
                        #
                        sysname Spoke-CE1
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:1::1/64
                        #
                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:db8:11::1/128
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        bgp 65410
                         router-id 1.1.1.1
                         peer 2001:db8:1::2 as-number 100
                         #
                         ipv6-family unicast
                          network 2001:db8:11::1 128
                          peer 2001:db8:1::2 enable
                        #
                        return
                    ●   Spoke-PE1
                        #
                        sysname Spoke-PE1
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpna
                         ipv6-family
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
                         ipv6 enable
                         ipv6 address 2001:db8:1::2/64
                        #
                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         405
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

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
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 2.2.2.9 enable
                         #
                         ipv6-family vpn-instance vpna
                          peer 2001:db8:1::1 as-number 65410
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                        #
                        return
                    ●   Spoke-PE2
                        #
                        sysname Spoke-PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpna
                         ipv6-family
                          route-distinguisher 100:3
                          vpn-target 100:1 export-extcommunity
                          vpn-target 200:1 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpna
                         ipv6 enable
                         ipv6 address 2001:db8:2::2/64
                        #
                        interface Vlanif200
                         ip address 10.1.1.1 255.255.255.0
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         406
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                          peer 2.2.2.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 2.2.2.9 enable
                         #
                         ipv6-family vpn-instance vpna
                          peer 2001:db8:2::1 as-number 65420
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        return

                    ●   Spoke-CE2
                        #
                        sysname Spoke-CE2
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:2::1/64
                        #
                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:db8:12::2/128
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        bgp 65420
                         router-id 3.3.3.3
                         peer 2001:db8:2::2 as-number 100
                         #
                         ipv6-family unicast
                          network 2001:db8:12::2 128
                          peer 2001:db8:2::2 enable
                        #
                        return

