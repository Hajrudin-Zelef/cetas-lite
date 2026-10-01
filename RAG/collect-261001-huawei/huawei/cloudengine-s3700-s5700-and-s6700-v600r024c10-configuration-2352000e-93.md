---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-93
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [12891, 13042]
sha256: 045cdd0ae27bdb191791129335c1f7948c0848772bd698c47866df388f7cc847
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                          peer 4.4.4.4 enable
                         #
                         ipv4-family vpn-instance vpna
                          auto-frr
                          route-select delay 300
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 172.16.2.0 0.0.0.255
                          network 172.17.4.0 0.0.0.255
                        #
                        return
                    ●   SPE2
                        #
                        sysname SPE2
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:1
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.17.4.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 172.18.4.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip address 172.19.6.1 255.255.255.0
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         206
VPN Configuration
VPN Configuration                                                                       3 IPv4 L3VPN Configuration

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
                          policy vpn-target
                          peer 1.1.1.1 enable
                          peer 1.1.1.1 ip-prefix default export
                          peer 1.1.1.1 upe
                          peer 2.2.2.2 enable
                          peer 2.2.2.2 ip-prefix default export
                          peer 2.2.2.2 upe
                          peer 5.5.5.5 enable
                          peer 5.5.5.5 route-policy NPE1 import
                          peer 6.6.6.6 enable
                          peer 6.6.6.6 route-policy NPE2 import
                         #
                         ipv4-family vpn-instance vpna
                          network 0.0.0.0 route-policy default
                          auto-frr
                          route-select delay 300
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 172.17.4.0 0.0.0.255
                          network 172.18.4.0 0.0.0.255
                          network 172.19.6.0 0.0.0.255
                        #
                        ip ip-prefix default index 10 permit 0.0.0.0 0
                        #
                        route-policy NPE1 deny node 10
                         if-match ip-prefix default
                        #
                        route-policy NPE1 permit node 20
                         apply local-preference 180
                        #
                        route-policy NPE2 deny node 10
                         if-match ip-prefix default
                        #
                        route-policy NPE2 permit node 20
                         apply local-preference 170
                        #
                        route-policy default permit node 10
                         apply local-preference 190
                        #
                        ip route-static vpn-instance vpna 0.0.0.0 0.0.0.0 66.66.66.66
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   207
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

