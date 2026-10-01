---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-92
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [12737, 12890]
sha256: 96407c0d5e94f451e44a9473dd51598ed2d19f6d1c95caafa8bf1398e7d05961
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 5.5.5.5 255.255.255.255
                        #
                        interface LoopBack2
                         ip binding vpn-instance vpna
                         ip address 55.55.55.55 255.255.255.255
                        #
                        bgp 100
                         router-id 5.5.5.5
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
                          auto-frr



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         204
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                          route-select delay 300
                          peer 10.4.1.2 as-number 65420
                        #
                        ospf 1
                         area 0.0.0.0
                          network 5.5.5.5 0.0.0.0
                          network 172.18.5.0 0.0.0.255
                          network 172.20.6.0 0.0.0.255
                        #
                        route-policy SPE1 permit node 10
                         apply local-preference 200
                        #
                        route-policy SPE2 permit node 10
                         apply local-preference 190
                        #
                        return

                    ●   UPE2
                        #
                        sysname UPE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:1
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.17.4.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 172.16.2.2 255.255.255.0
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
                         ip address 2.2.2.2 255.255.255.255
                        #
                        bgp 100
                         router-id 2.2.2.2
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         205
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

