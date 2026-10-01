---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-81
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [11124, 11277]
sha256: a3adacd7e99f5a6063a1fa85e8a16680aaa4f94fd49f534efd301a07e2627d40
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                          peer 10.1.1.1 as-number 65001
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return
                    ●   ASBR1
                        #
                        sysname ASBR1
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:3
                          vpn-target 1:1 export-extcommunity
                          vpn-target 1:1 import-extcommunity
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.10.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.12.12.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpna
                         ip address 10.3.1.2 255.255.255.0
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         177
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        bgp 100
                         peer 10.12.12.2 as-number 200
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 10.12.12.2 enable
                          peer 1.1.1.9 enable
                         #
                         ipv4-family vpnv4
                          undo policy vpn-target
                          peer 1.1.1.9 enable
                          peer 10.12.12.2 enable
                         #
                         ipv4-family vpn-instance vpna
                          peer 10.3.1.1 as-number 65003
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return

                    ●   CE3
                        #
                        sysname CE3
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.3.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface Loopback1
                         ip address 33.33.33.33 255.255.255.255
                        #
                        bgp 65003
                         peer 10.3.1.2 as-number 100
                         #
                         ipv4-family unicast
                          network 33.33.33.33 255.255.255.255
                          peer 10.3.1.2 enable
                        return

                    ●   ASBR2
                        #
                        sysname ASBR2
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 200:4
                          vpn-target 1:1 export-extcommunity
                          vpn-target 1:1 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         178
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

