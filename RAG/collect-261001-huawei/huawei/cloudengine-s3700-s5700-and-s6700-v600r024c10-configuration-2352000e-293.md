---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-293
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [43267, 43437]
sha256: 6d6d1594febcd206b444b5885de31fa82ccbda7f36ce4edad7d6bb09b06ae9c9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         vlan 10
                         #
                         interface Vlanif10
                          ip address 10.1.1.1 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 10


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                        694
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 1.1.1.1
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi v1 auto
                         pwsignal bgp
                          route-distinguisher 100:1
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                          site 1 range 5 default-offset 0
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 10.0000.0000.0001.00
                        #
                        interface Vlanif10
                         l2 binding vsi v1
                        #
                        interface Vlanif20
                         ip address 100.1.1.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                         isis enable 1
                        #
                        bgp 100
                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 2.2.2.2 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 2.2.2.2 enable
                          peer 2.2.2.2 signaling vpls
                        #
                        return
                    ●   ASBR_PE1
                        #
                        sysname ASBR_PE1
                        #
                        vlan batch 20 30
                        #
                        mpls lsr-id 2.2.2.2
                        mpls


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   695
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        #
                        mpls l2vpn
                        #
                        vsi v1 auto
                         pwsignal bgp
                          route-distinguisher 100:2
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                          site 2 range 5 default-offset 0
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 10.0000.0000.0002.00
                        #
                        interface Vlanif20
                         ip address 100.1.1.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         l2 binding vsi v1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                         isis enable 1
                        #
                        bgp 100
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 1.1.1.1 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 1.1.1.1 enable
                          peer 1.1.1.1 signaling vpls
                        #
                        return
                    ●   ASBR_PE2
                        #
                        sysname ASBR_PE2
                        #
                        vlan batch 30 40
                        #
                        mpls lsr-id 3.3.3.3
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi v1 auto
                         pwsignal bgp
                          route-distinguisher 200:1
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                          site 1 range 5 default-offset 0
                        #
                        mpls ldp


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   696
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

