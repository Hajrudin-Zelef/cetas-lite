---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-157
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [22468, 22664]
sha256: 83bafe2bf77f1f9202fcd8638a9d5334fdef40b844aa88ec8f638e7e94af36f7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                         ipv6 enable
                         ip binding vpn-instance vpn1
                         ipv6 address 2001:db8:1::2/64
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
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 1.1.1.1 enable
                          peer 3.3.3.3 enable
                        #
                         ipv6-family vpn-instance vpn1
                          preference 100 255 255
                          auto-frr
                          route-select delay 300
                          peer 2001:db8:1::1 as-number 65410
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                          network 10.11.1.0 0.0.0.3
                          network 2.2.2.2 0.0.0.0
                        #
                        return

                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpn1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         355
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

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
                         ipv6 address 2001:db8:3::2/64
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
                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                          peer 2.2.2.2 enable
                        #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 1.1.1.1 enable
                          peer 2.2.2.2 enable
                        #
                         ipv6-family vpn-instance vpn1
                          preference 100 255 255
                          auto-frr
                          peer 2001:db8:3::1 as-number 65410
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.20.1.0 0.0.0.3
                          network 10.11.1.0 0.0.0.3
                          network 3.3.3.3 0.0.0.0
                        #
                        Return

                    ●   CE


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         356
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        #
                        sysname CE
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:1::1/64
                        #
                        interface Vlanif200
                         ipv6 enable
                         ipv6 address 2001:db8:3::1/64
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
                         ipv6 enable
                         ipv6 address 2001:db8:0:1:2::1/128
                        #
                        bgp 65410
                         router-id 10.10.10.10
                         peer 2001:db8:1::2 as-number 100
                         peer 2001:db8:3::2 as-number 100
                         #
                         ipv4-family unicast
                         #
                         ipv6-family unicast
                          network 2001:db8:0:1:2::1 128
                          peer 2001:db8:1::2 enable
                          peer 2001:db8:3::2 enable
                        #
                        return



4.11 Configuring 6VPE

