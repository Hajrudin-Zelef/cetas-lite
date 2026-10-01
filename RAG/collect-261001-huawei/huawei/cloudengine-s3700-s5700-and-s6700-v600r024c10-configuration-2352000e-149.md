---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-149
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [21237, 21399]
sha256: 1172e87c9d83b2c21fb28d542a78735e046abb982766eaac572b25ac8f430172
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 2.2.2.2 enable
                          peer 3.3.3.3 enable
                         #
                         ipv6-family vpn-instance vpn1
                          auto-frr
                          route-select delay 300
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                          network 10.20.1.0 0.0.0.3
                        #
                        bfd for_ldp_lsp bind peer-ip 2.2.2.2 interface Vlanif200
                         discriminator local 10
                         discriminator remote 20
                         process-pst
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv6-family
                          route-distinguisher 100:2
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        bfd
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.10.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ipv6 enable
                         ipv6 address 2001:DB8:1::2/64
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
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                              337
VPN Configuration
VPN Configuration                                                                  4 IPv6 L3VPN Configuration

                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 1.1.1.1 enable
                         #
                         ipv6-family vpn-instance vpn1
                          peer 2001:DB8:1::1 as-number 65410
                          import-route direct
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                        #
                        bfd for_ldp_lsp bind peer-ip 1.1.1.1 interface Vlanif100
                         discriminator local 20
                         discriminator remote 10
                        #
                        return

                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv6-family
                          route-distinguisher 100:3
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
                         ipv6 enable
                         ip binding vpn-instance vpn1
                         ipv6 address 2001:DB8:3::2/64
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
                        bgp 100
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 1.1.1.1 enable
                         #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                              338
VPN Configuration
VPN Configuration                                                            4 IPv6 L3VPN Configuration

                         ipv6-family vpn-instance vpn1
                          peer 2001:DB8:3::1 as-number 65410
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.20.1.0 0.0.0.3
                        #
                        return

