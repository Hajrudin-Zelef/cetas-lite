---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-41
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [4827, 5022]
sha256: d5875f5f6470d5330b3a2814419601aa2e704bcce562ef898b11217a15a27c5b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 100 200 300
                         #
                         ip vpn-instance vpna
                          ipv4-family
                           route-distinguisher 100:1
                           vpn-target 111:1 export-extcommunity
                           vpn-target 111:1 import-extcommunity
                         #
                         ip vpn-instance vpnb
                          ipv4-family
                           route-distinguisher 100:2
                           vpn-target 222:2 export-extcommunity
                           vpn-target 222:2 import-extcommunity
                         #
                         mpls lsr-id 1.1.1.9
                         #
                         mpls
                         #
                         mpls ldp
                         #
                         interface Vlanif100
                          ip binding vpn-instance vpna
                          ip address 10.1.1.2 255.255.255.0
                         #
                         interface Vlanif200
                          ip binding vpn-instance vpnb
                          ip address 10.2.1.2 255.255.255.0
                         #
                         interface Vlanif300
                          ip address 11.11.11.1 255.255.255.0
                          mpls
                          mpls ldp


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   76
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

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
                         ip address 1.1.1.9 255.255.255.255
                        #
                        bgp 100
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.9 enable
                        #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpn-instance vpna
                          import-route direct
                          peer 10.1.1.1 as-number 65410
                        #
                         ipv4-family vpn-instance vpnb
                          import-route direct
                          peer 10.2.1.1 as-number 65420
                        #
                        ospf 1
                         area 0.0.0.0
                          network 11.11.11.0 0.0.0.255
                          network 1.1.1.9 0.0.0.0
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 200 300
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif300
                         ip address 11.11.11.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 12.12.12.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/2
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          77
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 11.11.11.0 0.0.0.255
                          network 12.12.12.0 0.0.0.255
                          network 2.2.2.9 0.0.0.0
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 200:1
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        ip vpn-instance vpnb
                         ipv4-family
                          route-distinguisher 200:2
                          vpn-target 222:2 export-extcommunity
                          vpn-target 222:2 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpna
                         ip address 10.3.1.2 255.255.255.0
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpnb
                         ip address 10.4.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 12.12.12.2 255.255.255.0
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
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          78
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

