---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-301
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [44554, 44705]
sha256: 9ce6ebbbd1b90d753f217201a6b0ccbe42d64479e9e443b39e6abcc3e51d866d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         mpls ldp
                        #
                        interface Vlanif40
                         ip address 192.168.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif50
                         ip address 192.168.4.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif60
                         ip address 192.168.5.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 60
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 100
                         peer 4.4.4.9 as-number 100
                         peer 4.4.4.9 connect-interface LoopBack1
                         peer 5.5.5.9 as-number 100
                         peer 5.5.5.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 4.4.4.9 enable
                          peer 5.5.5.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 4.4.4.9 enable
                          peer 5.5.5.9 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 192.168.2.0 0.0.0.255
                          network 192.168.3.0 0.0.0.255
                          network 192.168.4.0 0.0.0.255
                          network 192.168.5.0 0.0.0.255
                        #
                        return
                    ●   PE4
                        #
                        sysname PE4
                        #
                        vlan batch 60 70
                        #
                        mpls lsr-id 4.4.4.9
                        mpls


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   716
VPN Configuration
VPN Configuration                                                              6 VPLS Configuration

                        #
                        mpls l2vpn
                        #
                        vsi vsi1
                         bgp-ad
                          vpls-id 192.168.0.0:1
                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                        #
                        mpls ldp
                        #
                        interface Vlanif60
                         ip address 192.168.5.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif70
                         ip address 192.168.6.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 70
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 60
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                        #
                        bgp 100
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         peer 5.5.5.9 as-number 100
                         peer 5.5.5.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 3.3.3.9 enable
                          peer 5.5.5.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 3.3.3.9 enable
                          peer 5.5.5.9 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 192.168.5.0 0.0.0.255
                          network 192.168.6.0 0.0.0.255
                        #
                        return
                    ●   PE5
                        #
                        sysname PE5
                        #
                        vlan batch 50 70 80
                        #
                        mpls lsr-id 5.5.5.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi vsi1
                         bgp-ad
                          vpls-id 192.168.0.0:1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   717
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

