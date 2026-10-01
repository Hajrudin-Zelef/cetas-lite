---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-185
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [26912, 27094]
sha256: b825392352f8248f04a51aa287b878ab2051012da2d4aab9c2ae376f143dbbc6
---

                    # Ping CE1 and CE2 from each other. The ping operations are successful. The
                    following example uses the command output on CE1.
                    <CE1> ping 10.10.1.2
                     PING 10.10.1.2: 56 data bytes, press CTRL_C to break
                       Reply from 10.10.1.2: bytes=56 Sequence=1 ttl=255 time=58 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=2 ttl=255 time=67 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=3 ttl=255 time=52 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=4 ttl=255 time=69 ms
                       Reply from 10.10.1.2: bytes=56 Sequence=5 ttl=255 time=92 ms
                     --- 10.10.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                  427
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration

                        0.00% packet loss
                        round-trip min/avg/max = 52/67/92 ms


Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan batch 10
                          #
                          interface Vlanif10
                           ip address 10.10.1.1 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 10
                          #
                          return

                    ●     PE1
                          #
                          sysname PE1
                          #
                          vlan batch 10 20
                          #
                          mpls lsr-id 1.1.1.9
                          #
                          mpls
                           mpls te
                          #
                          mpls l2vpn
                          #
                          interface Vlanif10
                          #
                          interface Vlanif20
                           ip address 10.1.1.1 255.255.255.0
                           mpls
                           mpls te
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
                           ip address 1.1.1.9 255.255.255.255
                          #
                          ccc CE1-CE2 interface Vlanif10 in-label 100 out-label 200 nexthop 10.1.1.2
                          #
                          return

                    ●     P
                          #
                          sysname P
                          #
                          vlan batch 10 20
                          #
                          mpls lsr-id 2.2.2.9
                          #
                          mpls
                           mpls te
                          #
                          interface Vlanif10
                           ip address 10.2.2.2 255.255.255.0


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          428
VPN Configuration
VPN Configuration                                                                                    5 VPWS Configuration

                         mpls
                         mpls te
                        #
                        interface Vlanif20
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls te
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
                        static-cr-lsp transit PE1-PE2 incoming-interface Vlanif20 in-label 200 nexthop 10.2.2.1 out-label 201
                        static-cr-lsp transit PE2-PE1 incoming-interface Vlanif10 in-label 101 nexthop 10.1.1.1 out-label 100
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                        #
                        mpls l2vpn
                        #
                        interface Vlanif10
                         ip address 10.2.2.1 255.255.255.0
                         mpls
                         mpls te
                        #
                        interface Vlanif20
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
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ccc CE2-CE1 interface Vlanif20 in-label 201 out-label 101 nexthop 10.2.2.2
                        #
                        return
                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 20
                        #
                        interface Vlanif20
                         ip address 10.10.1.2 255.255.255.0
                        #
                        interface 10GE1/0/2
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                429
VPN Configuration
VPN Configuration                                                                   5 VPWS Configuration

                         port trunk allow-pass vlan 20
                        #
                        return



5.6 Configuring BGP VPWS

5.6.1 Understanding BGP VPWS

Definition
                    BGP VPWS is a type of MPLS L2VPN technology that uses BGP as the signaling
                    protocol to transmit Layer 2 information and VC labels between PEs.

