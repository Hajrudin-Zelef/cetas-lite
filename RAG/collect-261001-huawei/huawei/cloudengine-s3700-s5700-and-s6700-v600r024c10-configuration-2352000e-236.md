---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-236
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [34652, 34859]
sha256: c05683941edb1e1e3ad61ae3ae9aa22561fd9b37f75debfeb908405f641f894f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        return
                    ●   P1
                        #
                        sysname P1
                        #
                         mpls lsr-id 2.2.2.2
                         mpls
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.3.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/2
                         ip address 10.2.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.2.1.0 0.0.0.3
                          network 10.3.1.0 0.0.0.3
                        #
                        return
                    ●   P2
                        #
                        sysname P2
                        #
                         mpls lsr-id 3.3.3.3
                         mpls
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.5.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.4.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.4.1.0 0.0.0.3
                          network 10.5.1.0 0.0.0.3
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   551
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration

                         bfd
                        #
                         mpls lsr-id 4.4.4.4
                         mpls
                        #
                         mpls l2vpn
                        #
                        pw-template 2to1
                         peer-address 1.1.1.1
                         control-word
                        #
                        mpls ldp
                        #
                         mpls ldp remote-peer 1.1.1.1
                         remote-ip 1.1.1.1
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         mpls l2vc pw-template 2to1 100
                         mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4 track-interface
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.3.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.3.1.0 0.0.0.3
                        #
                        bfd 2to1 bind pw interface 10GE1/0/1
                         discriminator local 21
                         discriminator remote 12
                        #
                        return
                    ●   PE3
                        #
                        sysname PE3
                        #
                         bfd
                        #
                         mpls lsr-id 5.5.5.5
                         mpls
                        #
                         mpls l2vpn
                        #
                        pw-template 3to1
                         peer-address 1.1.1.1
                         control-word
                        #
                        mpls ldp
                        #
                         mpls ldp remote-peer 1.1.1.1
                         remote-ip 1.1.1.1
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.5.1.2 255.255.255.252
                         mpls
                         mpls ldp
                         mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4 track-interface
                        #
                        interface 10GE1/0/2
                         undo portswitch


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              552
VPN Configuration
VPN Configuration                                                                          5 VPWS Configuration

                         mpls l2vc pw-template 3to1 200
                        #
                        interface LoopBack1
                         ip address 5.5.5.5 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 5.5.5.5 0.0.0.0
                          network 10.5.1.0 0.0.0.3
                        #
                        bfd 3to1 bind pw interface 10GE1/0/2
                         discriminator local 31
                         discriminator remote 13
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.1.1.2 255.255.255.252
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.1.2.2 255.255.255.252
                        #
                        return



5.11.5 Example for Configuring Dynamic BFD for VPWS
Networking Requirements
                    On the MPLS L2VPN networking:
                    ●   PW1 is established between PE1 and PE2 as the primary PW.
                    ●   PW2 is established between PE1 and PE3 as the secondary PW.
                    On the network shown in Figure 5-38, BFD needs to be configured to detect the
                    connectivity of the primary and secondary PWs, so that services can be switched
                    to the secondary PW within 50 ms if the primary PW fails.

                    Figure 5-38 Network diagram for configuring dynamic BFD for PWs
                         NOTE

                        In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                        10GE1/0/3, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    553
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration




