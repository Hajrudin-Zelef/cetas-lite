---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-64
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [8401, 8553]
sha256: b22bfe28ff9a838e1c676275b539d8d5c6420a26a531f452921e7c1f1537dbe9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         ipv4-family vpn-instance vpn1
                         import-route direct
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                          network 10.20.1.0 0.0.0.3
                          network 1.1.1.1 0.0.0.0
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpn1
                         ipv4-family
                          route-distinguisher 100:2
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                          ip frr
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ip address 192.168.1.1 255.255.255.252
                        #
                        interface Vlanif300
                         ip address 10.11.1.1 255.255.255.252
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
                         ip address 2.2.2.2 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         peer 3.3.3.3 as-number 100
                         peer 3.3.3.3 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                          peer 3.3.3.3 enable
                        #
                         ipv4-family vpnv4


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         133
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                          policy vpn-target
                          peer 1.1.1.1 enable
                          peer 3.3.3.3 enable
                        #
                         ipv4-family vpn-instance vpn1
                           import-route ospf 2
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                          network 10.11.1.0 0.0.0.3
                          network 2.2.2.2 0.0.0.0
                        #
                        ospf 2 vpn-instance vpn1
                         import-route bgp
                         area 0.0.0.1
                          network 192.168.1.0 0.0.0.3
                        #
                        return

                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpn1
                         ipv4-family
                          route-distinguisher 100:2
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.20.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ip address 192.168.2.1 255.255.255.252
                        #
                        interface Vlanif300
                         ip address 10.11.1.2 255.255.255.252
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
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         134
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

