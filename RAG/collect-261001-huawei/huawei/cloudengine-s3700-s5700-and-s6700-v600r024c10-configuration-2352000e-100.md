---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-100
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [13953, 14104]
sha256: be5c3d5ad19bf66ff57ef23ee09e65c72a6d90e935abffe9bc396f5081e28f4b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 5.5.5.5 255.255.255.255
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         221
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         222
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

