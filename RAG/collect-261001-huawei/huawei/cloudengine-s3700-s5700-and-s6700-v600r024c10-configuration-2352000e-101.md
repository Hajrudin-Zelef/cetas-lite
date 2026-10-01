---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-101
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [14105, 14275]
sha256: 3c41e684091040ddf183be776fad135d3c44d7def4373de8ac6f57d0a43e9e89
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                         ip address 4.4.4.4 255.255.255.255
                        #
                        bgp 100
                         router-id 4.4.4.4
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
                          network 4.4.4.4 0.0.0.0
                          network 172.17.4.0 0.0.0.255
                          network 172.18.4.0 0.0.0.255
                          network 172.19.6.0 0.0.0.255
                        #
                        route-policy NPE1 permit node 10
                         apply local-preference 180
                        #
                        route-policy NPE2 permit node 10
                         apply local-preference 170
                        #
                        route-policy pref permit node 10



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         223
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         apply local-preference 50
                        #
                        return
                    ●   NPE2
                        #
                        sysname NPE2
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:1
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                        #
                        mpls lsr-id 6.6.6.6
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.19.6.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 172.20.6.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
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
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 6.6.6.6 255.255.255.255
                        #
                        bgp 100
                         peer 3.3.3.3 as-number 100
                         peer 3.3.3.3 connect-interface LoopBack1
                         peer 4.4.4.4 as-number 100
                         peer 4.4.4.4 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.3 enable
                          peer 4.4.4.4 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 3.3.3.3 enable
                          peer 3.3.3.3 route-policy SPE1 import
                          peer 4.4.4.4 enable
                          peer 4.4.4.4 route-policy SPE2 import
                         #
                         ipv4-family vpn-instance vpna
                          import-route direct


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         224
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                          auto-frr
                          route-select delay 300
                          peer 10.2.1.2 as-number 65420
                        #
                        ospf 1
                         area 0.0.0.0
                          network 6.6.6.6 0.0.0.0
                          network 172.19.6.0 0.0.0.255
                          network 172.20.6.0 0.0.0.255
                        #
                        route-policy SPE1 permit node 10
                         apply local-preference 180
                        #
                        route-policy SPE2 permit node 10
                         apply local-preference 170
                        #
                        return

