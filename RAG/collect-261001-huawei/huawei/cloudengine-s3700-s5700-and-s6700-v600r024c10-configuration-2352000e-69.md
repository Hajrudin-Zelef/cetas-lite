---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-69
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [9147, 9334]
sha256: ef9e97349c686504ee28b3f198e410f73bdaa861b924d738dcf5ff2e92085f6e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    *>i 11.11.11.11/32     1.1.1.9         0        100        0      65001i
                    Route Distinguisher: 100:2


                           Network           NextHop       MED        LocPrf       PrefVal Path/Ogn

                    *>      22.22.22.22/32    12.12.12.2                       0     200 65002i

                    VPN-Instance vpn1, Router ID 10.10.1.1:

                    Total Number of Routes: 2
                         Network        NextHop            MED        LocPrf       PrefVal Path/Ogn

                    *>i     11.11.11.11/32    1.1.1.9      0        100        0     65001i
                    *>      22.22.22.22/32    12.12.12.2                       0     200 65002i


Configuration Scripts
                    ●      CE1
                           #
                           sysname CE1
                           #
                           vlan batch 100
                           #
                           interface Vlanif100
                            ip address 10.1.1.1 255.255.255.0
                           #
                           interface 10GE1/0/1


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                              145
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 11.11.11.11 255.255.255.255
                        #
                        bgp 65001
                         peer 10.1.1.2 as-number 100
                         #
                         ipv4-family unicast
                          network 11.11.11.11 255.255.255.255
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
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.10.1.2 255.255.255.0
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
                         ipv4-family vpn-instance vpn1
                          peer 10.1.1.1 as-number 65001
                        #
                        ospf 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         146
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return

                    ●   ASBR1
                        #
                        sysname ASBR1
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv4-family
                          route-distinguisher 100:2
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
                         ip binding vpn-instance vpn1
                         ip address 12.12.12.1 255.255.255.0
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
                         ip address 2.2.2.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.9 enable
                         #
                         ipv4-family vpn-instance vpn1
                          peer 12.12.12.2 as-number 200
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return

                    ●   ASBR2
                        #
                        sysname ASBR2
                        #
                        vlan batch 100 200


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         147
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

