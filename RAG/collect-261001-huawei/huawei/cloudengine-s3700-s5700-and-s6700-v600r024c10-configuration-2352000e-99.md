---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-99
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [13801, 13952]
sha256: a037af8b0376134de3681260cc02bfd48ad6d908f4f35c11542f107f7ab2a941
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        return
                    ●   SPE1
                        #
                        sysname SPE1
                        #
                        vlan batch 100 200 300
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.16.3.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 172.18.4.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip address 172.18.5.1 255.255.255.0
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
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        bgp 100
                         router-id 3.3.3.3
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         peer 5.5.5.5 as-number 100
                         peer 5.5.5.5 connect-interface LoopBack1
                         peer 6.6.6.6 as-number 100
                         peer 6.6.6.6 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                          peer 2.2.2.2 enable
                          peer 5.5.5.5 enable
                          peer 6.6.6.6 enable
                         #
                         ipv4-family vpnv4
                          undo policy vpn-target
                          auto-frr
                          route-select delay 300
                          bestroute nexthop-resolved tunnel
                          peer 1.1.1.1 enable
                          peer 1.1.1.1 route-policy pref export
                          peer 1.1.1.1 reflect-client


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         219
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                          peer 1.1.1.1 next-hop-local
                          peer 2.2.2.2 enable
                          peer 2.2.2.2 route-policy pref export
                          peer 2.2.2.2 reflect-client
                          peer 2.2.2.2 next-hop-local
                          peer 5.5.5.5 enable
                          peer 5.5.5.5 route-policy NPE1 import
                          peer 5.5.5.5 reflect-client
                          peer 5.5.5.5 next-hop-local
                          peer 6.6.6.6 enable
                          peer 6.6.6.6 route-policy NPE2 import
                          peer 6.6.6.6 reflect-client
                          peer 6.6.6.6 next-hop-local
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 172.16.3.0 0.0.0.255
                          network 172.18.4.0 0.0.0.255
                          network 172.18.5.0 0.0.0.255
                        #
                        route-policy NPE1 permit node 10
                         apply local-preference 200
                        #
                        route-policy NPE2 permit node 10
                         apply local-preference 190
                        #
                        route-policy pref permit node 10
                         apply local-preference 150
                        #
                        return
                    ●   NPE1
                        #
                        sysname NPE1
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:1
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                        #
                        mpls lsr-id 5.5.5.5
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.18.5.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 172.20.6.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpna
                         ip address 10.4.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         220
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

