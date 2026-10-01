---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-75
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [10119, 10318]
sha256: 92f2b01ccc6d1c182647a395bf69f1e71ecaae1e990270738a2af793fe112d73
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 5.5.5.5 255.255.255.255
                        #
                        bgp 65001
                         peer 10.1.1.2 as-number 100
                         network 5.5.5.5 255.255.255.255
                         #
                         ipv4-family unicast
                          peer 10.1.1.2 enable
                        #
                        ip route-static 10.2.1.0 255.255.255.0 10.1.1.2
                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv4-family
                          route-distinguisher 100:1
                          vpn-target 1:1 export-extcommunity
                          vpn-target 1:1 import-extcommunity
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.16.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ip address 10.1.1.2 255.255.255.0
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
                         ip address 1.1.1.1 255.255.255.255
                        #
                        bgp 100
                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 2.2.2.2 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 2.2.2.2 enable
                         #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         161
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         ipv4-family vpn-instance vpn1
                          peer 10.1.1.1 as-number 65001
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.16.1.0 0.0.0.255
                        #
                        return

                    ●   ASBR1
                        #
                        sysname ASBR1
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.16.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 192.168.1.1 255.255.255.0
                         mpls
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
                         peer 192.168.1.2 as-number 200
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 192.168.1.2 enable
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpnv4
                          undo policy vpn-target
                          peer 1.1.1.1 enable
                          peer 192.168.1.2 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.16.1.0 0.0.0.255
                        #
                        return

                    ●   ASBR2
                        #
                        sysname ASBR2
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 3.3.3.3


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         162
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.17.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 192.168.1.2 255.255.255.0
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
                         ip address 3.3.3.3 255.255.255.255
                        #
                        bgp 200
                         peer 192.168.1.1 as-number 100
                         peer 4.4.4.4 as-number 200
                         peer 4.4.4.4 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 192.168.1.1 enable
                          peer 4.4.4.4 enable
                         #
                         ipv4-family vpnv4
                          undo policy vpn-target
                          peer 4.4.4.4 enable
                          peer 192.168.1.1 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.17.1.0 0.0.0.255
                        #
                        return

