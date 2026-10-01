---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-281
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [41292, 41490]
sha256: f64ba847edaab39bead891290d635a5b3849287a83ce32cc5d99eda67c58b244
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        bgp 100
                         peer 2.2.2.9 as-number 100
                         peer 2.2.2.9 connect-interface LoopBack1
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 2.2.2.9 enable
                          peer 3.3.3.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 2.2.2.9 enable
                          peer 3.3.3.9 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 172.16.1.0 0.0.0.255
                          network 172.16.2.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 20 40 50
                        #
                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi vplsad1
                         bgp-ad
                          vpls-id 172.16.1.1:1
                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                        #
                        mpls ldp
                        #
                        interface Vlanif20
                         ip address 172.16.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif40
                         ip address 172.17.3.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif50
                         l2 binding vsi vplsad1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   663
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 172.16.2.0 0.0.0.255
                          network 172.17.3.0 0.0.0.255
                        #
                        return
                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 30 40 60
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi vplsad1
                         bgp-ad
                          vpls-id 172.16.1.1:1
                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                        #
                        mpls ldp
                        #
                        interface Vlanif30
                         ip address 172.16.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif40
                         ip address 172.17.3.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif60
                         l2 binding vsi vplsad1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 60


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   664
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration

                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         peer 2.2.2.9 as-number 100
                         peer 2.2.2.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 1.1.1.9 enable
                          peer 2.2.2.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 2.2.2.9 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 172.16.1.0 0.0.0.255
                          network 172.17.3.0 0.0.0.255
                        #
                        return



6.9 Configuring LDP HVPLS
Prerequisites
                    Before configuring LDP HVPLS, you have completed the following tasks:
                    ●   Configure LSR IDs on underlay provider edges (UPEs) and superstratum
                        provider edges (SPEs).
                    ●   Configure MPLS and MPLS LDP on UPEs and SPEs.
                    ●   Configure MPLS L2VPN on UPEs and SPEs.
                         NOTE

