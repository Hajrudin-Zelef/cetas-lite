---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-329
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [49102, 49304]
sha256: 3e8b47002ad600bec39a5569cdbd45845c49fc06cb2f70d3327715e9dd7a3631
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 10 100
                        #
                        stp region-configuration
                         instance 1 vlan 10 100
                        #
                        erps ring 1
                         control-vlan 100
                         protected-instance 1
                         version v2
                         sub-ring
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi s1 static
                         pwsignal ldp
                          vsi-id 10
                          peer 3.3.3.3
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         l2 binding vsi s1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         undo port trunk allow-pass vlan 1
                         port trunk allow-pass vlan 10 100
                         stp disable
                         erps ring 1
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        return

                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 10 100
                        #
                        stp region-configuration
                         instance 1 vlan 10 100
                        #
                        erps ring 1
                         control-vlan 100
                         protected-instance 1
                         version v2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   789
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                         sub-ring
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi s1 static
                         pwsignal ldp
                          vsi-id 10
                          peer 3.3.3.3
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         l2 binding vsi s1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         undo port trunk allow-pass vlan 1
                         port trunk allow-pass vlan 10 100
                         stp disable
                         erps ring 1
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.2.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                        #
                        return
                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 10
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi s1 static
                         pwsignal ldp
                          vsi-id 10
                          peer 1.1.1.1
                          peer 2.2.2.2
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         l2 binding vsi s1
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   790
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.2.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.2.1.0 0.0.0.255
                        #
                        return

                    ●   CE1
                        #
                        sysname CE1
                        #
                        vlan batch 10 100
                        #
                        stp region-configuration
                         instance 1 vlan 10 100
                        #
                        erps ring 1
                         control-vlan 100
                         protected-instance 1
                         version v2
                         sub-ring
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         undo port trunk allow-pass vlan 1
                         port trunk allow-pass vlan 10 100
                         stp disable
                         erps ring 1
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         undo port trunk allow-pass vlan 1
                         port trunk allow-pass vlan 10 100
                         stp disable
                         erps ring 1
                        #
                        return

