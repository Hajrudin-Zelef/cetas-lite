---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-300
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [44386, 44553]
sha256: 86e2e14251eadba97039f7525ef5b4de6058849e21f18f1bf01a491aa1e6a251
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Verifying the Configuration
                    # Perform a ping test to check the connectivity.
                    [CE1] ping 192.168.10.2
                     PING 192.168.10.2: 56 data bytes, press CTRL_C to break
                      Reply from 192.168.10.2: bytes=56 Sequence=1 ttl=255 time=190 ms
                      Reply from 192.168.10.2: bytes=56 Sequence=2 ttl=255 time=190 ms
                      Reply from 192.168.10.2: bytes=56 Sequence=3 ttl=255 time=140 ms
                      Reply from 192.168.10.2: bytes=56 Sequence=4 ttl=255 time=140 ms
                      Reply from 192.168.10.2: bytes=56 Sequence=5 ttl=255 time=110 ms

                     --- 192.168.10.2 ping statistics ---


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                            713
VPN Configuration
VPN Configuration                                                               6 VPLS Configuration

                        5 packet(s) transmitted
                        5 packet(s) received
                        0.00% packet loss
                        round-trip min/avg/max = 110/154/190 ms

                    The command output shows that CE1 can ping CE2 successfully.

Configuration Scripts
                    ●     PE1
                          #
                          sysname PE1
                          #
                          vlan batch 10 20 40
                          #
                          mpls lsr-id 1.1.1.9
                          mpls
                          #
                          mpls l2vpn
                          #
                          vsi vsi1 static
                           pwsignal ldp
                            vsi-id 1
                            peer 2.2.2.9
                            peer 3.3.3.9
                          #
                          mpls ldp
                          #
                          interface Vlanif10
                           l2 binding vsi vsi1
                          #
                          interface Vlanif20
                           ip address 192.168.1.1 255.255.255.0
                           mpls
                           mpls ldp
                          #
                          interface Vlanif40
                           ip address 192.168.2.1 255.255.255.0
                           mpls
                           mpls ldp
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 20
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 40
                          #
                          interface 10GE1/0/3
                           port link-type trunk
                           port trunk allow-pass vlan 10
                          #
                          interface LoopBack1
                           ip address 1.1.1.9 255.255.255.255
                          #
                          ospf 1
                           area 0.0.0.0
                            network 1.1.1.9 0.0.0.0
                            network 192.168.1.0 0.0.0.255
                            network 192.168.2.0 0.0.0.255
                          #
                          return

                    ●     PE2
                          #
                          sysname PE2
                          #


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                   714
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        vlan batch 20 30
                        #
                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi vsi1 static
                         pwsignal ldp
                          vsi-id 1
                          peer 1.1.1.9
                          peer 3.3.3.9
                        #
                        mpls ldp
                        #
                        interface Vlanif20
                         ip address 192.168.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         ip address 192.168.3.1 255.255.255.0
                         mpls
                         mpls ldp
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
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 192.168.1.0 0.0.0.255
                          network 192.168.3.0 0.0.0.255
                        #
                        return
                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 30 40 50 60
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi vsi1
                         pwsignal ldp
                          vsi-id 1
                          peer 1.1.1.9 upe
                          peer 2.2.2.9 upe
                         bgp-ad
                          vpls-id 192.168.0.0:1
                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                        #
                        mpls ldp
                        #
                        interface Vlanif30
                         ip address 192.168.3.2 255.255.255.0
                         mpls


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   715
VPN Configuration
VPN Configuration                                                              6 VPLS Configuration

