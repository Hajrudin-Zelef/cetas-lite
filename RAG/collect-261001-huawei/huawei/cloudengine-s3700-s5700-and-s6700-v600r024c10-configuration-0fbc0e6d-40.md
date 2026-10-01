---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-40
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [5162, 5325]
sha256: 4a6a683ce9689336fd1b51d179dd772eb8ec72b158118d7be1c56b2d2b5bb392
---

                 # Configure LSRC.
                 [LSRC] mpls lsr-id 3.3.3.3
                 [LSRC] mpls
                 [LSRC-mpls] quit
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] quit
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] mpls
                 [LSRC-Vlanif100] mpls ldp
                 [LSRC-Vlanif100] quit

                 # Configure LSRD.
                 [LSRD] mpls lsr-id 4.4.4.4
                 [LSRD] mpls
                 [LSRD-mpls] quit


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        87
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration

                 [LSRD] mpls ldp
                 [LSRD-mpls-ldp] quit
                 [LSRD] interface vlanif 100
                 [LSRD-Vlanif100] mpls
                 [LSRD-Vlanif100] mpls ldp
                 [LSRD-Vlanif100] quit

                 # After completing the preceding configuration, run the display mpls lsp
                 command on LSRD to check information about established LSPs.
                 [LSRD] display mpls lsp
                 Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                 Flag after LDP FRR: (L) - Logic FRR LSP
                 -------------------------------------------------------------------------------
                              LSP Information: LDP LSP
                 -------------------------------------------------------------------------------
                 FEC              In/Out Label In/Out IF                        Vrf Name
                 1.1.1.1/32         NULL/32829         -/Vlanif100
                 1.1.1.1/32         32828/32829        -/Vlanif100
                 2.2.2.2/32         NULL/3          -/Vlanif100
                 2.2.2.2/32         32829/3         -/Vlanif100
                 3.3.3.3/32         NULL/32830         -/Vlanif100
                 3.3.3.3/32         32830/32830        -/Vlanif100
                 4.4.4.4/32         3/NULL          -/-

                 The command output shows that LSPs to LSRA, LSRB, and LSRC have been
                 established on LSRD.
         Step 3 Configure an inbound LDP policy.
                 # Configure an IP prefix list on LSRD to permit only the routes to LSRC.
                 [LSRD] ip ip-prefix prefix1 permit 3.3.3.3 32

                 # Configure an inbound policy on LSRD to allow LSRD to receive Label Mapping
                 messages only from LSRC.
                 [LSRD] mpls ldp
                 [LSRD-mpls-ldp] ipv4-family
                 [LSRD-mpls-ldp-ipv4] inbound peer 2.2.2.2 fec ip-prefix prefix1
                 [LSRD-mpls-ldp-ipv4] quit
                 [LSRD-mpls-ldp] quit

         Step 4 Verify the configuration.
                 After completing the preceding configuration, run the display mpls lsp command
                 on LSRD. The command output shows that only an LSP to LSRC is established.
                 [LSRD] display mpls lsp
                 Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                 Flag after LDP FRR: (L) - Logic FRR LSP
                 -------------------------------------------------------------------------------
                              LSP Information: LDP LSP
                 -------------------------------------------------------------------------------
                 FEC              In/Out Label In/Out IF                        Vrf Name
                 3.3.3.3/32         NULL/32830         -/Vlanif100
                 3.3.3.3/32         32830/32830        -/Vlanif100
                 4.4.4.4/32         3/NULL          -/-

                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         88
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        vlan batch 100
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
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
                 ●      LSRB
                        #
                        sysname LSRB
                        #
                        vlan batch 100 200 300
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type access
                         port default vlan 300
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        89
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

