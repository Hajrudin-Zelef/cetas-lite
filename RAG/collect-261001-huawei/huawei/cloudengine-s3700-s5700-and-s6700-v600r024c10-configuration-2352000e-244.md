---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-244
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [35770, 35973]
sha256: a8640bd07a2c01977411af82d1c760c77a1a72c1dcf93b1f27d79cb1137c1b7f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.2.1.0 0.0.0.3
                          network 10.4.1.0 0.0.0.3
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   567
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration

                          network 10.5.1.0 0.0.0.3
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
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
                         bfd-detect min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
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
                         bfd-detect min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                        #
                        mpls ldp
                        #
                         mpls ldp remote-peer 1.1.1.1
                         remote-ip 1.1.1.1
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.5.1.2 255.255.255.252


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              568
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration

                         mpls
                         mpls ldp
                         mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4 track-interface
                        #
                        interface 10GE1/0/2
                         undo portswitch
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



5.12 Configuring PW Redundancy

5.12.1 Understanding PW Redundancy
PW Redundancy Signaling
                    The introduction of the PW protection mechanism changes the model of one-to-
                    one mapping between ACs and PWs in PWE3. To retain the original forwarding
                    behavior, only one PW in a PW protection group works as the primary PW to
                    transmit data, while other PWs work as secondary PWs.

                    Relevant standards in LDP PW signaling specify the PW Status TLV to transmit the
                    PW forwarding status. The PW Status TLV, a 32-bit status code field, is carried in a
                    Label Mapping or Notification message. PW redundancy introduces a new PW
                    status code of 0x00000020 indicating PW forwarding standby to indicate that a
                    PW is a secondary PW.

                         NOTE

                        PW redundancy is supported only in PWE3 VPWS.


