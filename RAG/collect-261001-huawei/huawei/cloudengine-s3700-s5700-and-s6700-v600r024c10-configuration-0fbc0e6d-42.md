---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-42
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [5524, 5629]
sha256: 36cfaa63170c574802f9b35632ae8e47a5b555d0186394c1f3e5eab37c25d49d
---

                 # Configure LSRC.
                 <LSRC> system-view
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] mpls
                 [LSRC-Vlanif100] mpls ldp
                 [LSRC-Vlanif100] quit

                 # Configure the DSLAM.
                 <DSLAM> system-view
                 [DSLAM] interface vlanif 100
                 [DSLAM-Vlanif100] mpls
                 [DSLAM-Vlanif100] mpls ldp
                 [DSLAM-Vlanif100] quit

         Step 5 Verify the configuration.
                 After completing the preceding configuration, run the display mpls ldp lsp
                 command on the DSLAM. The command output shows that only an LSP to LSRC is
                 established.
                 [DSLAM] display mpls ldp lsp
                  LDP LSP Information
                   -------------------------------------------------------------------------------
                   Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                   -------------------------------------------------------------------------------
                   DestAddress/Mask In/OutLabel UpstreamPeer NextHop                             OutInterface
                   -------------------------------------------------------------------------------
                   1.1.1.9/32        NULL/1025         -            10.1.3.1       Vlanif100


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                93
MPLS Configuration
MPLS Configuration                                                                                     3 MPLS LDP Configuration

                  1.1.1.9/32        1024/1025         3.3.3.9       10.1.3.1       Vlanif100
                  4.4.4.9/32        3/NULL          3.3.3.9        127.0.0.1       LoopBack0
                  -------------------------------------------------------------------------------
                  TOTAL: 3 Normal LSP(s) Found.
                  TOTAL: 0 Liberal LSP(s) Found.
                  TOTAL: 0 Frr LSP(s) Found.
                  An asterisk (*) before an LSP means the LSP is not established
                  An asterisk (*) before a Label means the USCB or DSCB is stale
                  An asterisk (*) before an UpstreamPeer means the session is stale
                  An asterisk (*) before a DS means the session is stale
                  An asterisk (*) before a NextHop means the LSP is FRR LSP

                 If no outbound LDP policy is configured on LSRA, the LDP LSPs established on the
                 DSLAM are as follows:
                 LDP LSP Information
                   -------------------------------------------------------------------------------
                   Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                   -------------------------------------------------------------------------------
                   DestAddress/Mask In/OutLabel UpstreamPeer NextHop                             OutInterface
                   -------------------------------------------------------------------------------
                   1.1.1.9/32        NULL/1025         -            10.1.3.1       Vlanif100
                   1.1.1.9/32        1024/1025         3.3.3.9       10.1.3.1       Vlanif100
                   2.2.2.9/32        NULL/1024         -            10.1.3.1       Vlanif100
                   2.2.2.9/32        1027/1024         3.3.3.9       10.1.3.1       Vlanif100
                   3.3.3.9/32        NULL/3          -            10.1.3.1       Vlanif100
                   3.3.3.9/32        1028/3          3.3.3.9       10.1.3.1       Vlanif100
                   4.4.4.9/32        3/NULL          3.3.3.9        127.0.0.1       LoopBack0
                  *4.4.4.9/32         Liberal/1026                 DS/3.3.3.9
                   -------------------------------------------------------------------------------
                   TOTAL: 7 Normal LSP(s) Found.
                   TOTAL: 1 Liberal LSP(s) Found.
                   TOTAL: 0 Frr LSP(s) Found.
                   An asterisk (*) before an LSP means the LSP is not established
                   An asterisk (*) before a Label means the USCB or DSCB is stale
                   An asterisk (*) before an UpstreamPeer means the session is stale
                   An asterisk (*) before a DS means the session is stale
                   An asterisk (*) before a NextHop means the LSP is FRR LSP

                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                          outbound peer 4.4.4.9 fec ip-prefix prefix1
                        #
                        interface Vlanif100
                         ip address 10.1.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.3.1 255.255.255.0
                         mpls
                         mpls ldp
                        #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                94
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

