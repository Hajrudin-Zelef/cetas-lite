---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-311
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [46140, 46354]
sha256: 2d3b9246133a633ea094398247db9ccab692eddafc45a4e6a0765281d6d0704d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan 50
                          #
                          interface Vlanif50
                           ip address 10.1.1.1 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 50
                          #
                          return

                    ●     CE2
                          #
                          sysname CE2
                          #
                          vlan 60
                          #
                          interface Vlanif60
                           ip address 10.1.1.2 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 60
                          #
                          return

                    ●     UPE1
                          #
                          sysname UPE1
                          #
                          vlan batch 30 50
                          #
                          mpls lsr-id 4.4.4.9
                          mpls
                          #
                          mpls l2vpn
                          #
                          mpls ldp
                          #
                          interface Vlanif30
                           ip address 10.2.3.2 255.255.255.0
                           mpls
                           mpls ldp
                          #
                          interface Vlanif50
                           mpls l2vc 1.1.1.9 100
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 50
                          #
                          interface 10GE1/0/2


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                        741
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 10.2.3.0 0.0.0.255
                        #
                        return
                    ●   SPE1
                        #
                        sysname SPE1
                        #
                        vlan batch 10 30
                        #
                        mpls lsr-id 1.1.1.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi v100 static
                         pwsignal ldp
                          vsi-id 100
                          peer 3.3.3.9
                          peer 4.4.4.9 upe
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 3.3.3.9
                         remote-ip 3.3.3.9
                        #
                        interface Vlanif10
                         ip address 10.2.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         ip address 10.2.3.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                          network 10.2.3.0 0.0.0.255
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 10 20
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   742
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         ip address 10.2.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif20
                         ip address 10.2.2.1 255.255.255.0
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
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                          network 10.2.2.0 0.0.0.255
                        #
                        return
                    ●   SPE2
                        #
                        sysname SPE2
                        #
                        vlan batch 20 40
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi v100 static
                         pwsignal ldp
                          vsi-id 100
                          peer 1.1.1.9
                          peer 5.5.5.9 upe
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 1.1.1.9
                         remote-ip 1.1.1.9
                        #
                        interface Vlanif20
                         ip address 10.2.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif40
                         ip address 10.2.4.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   743
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration

