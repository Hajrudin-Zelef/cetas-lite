---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-32
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [3999, 4140]
sha256: 745c50666dcf0c140d7dfe2520a923475da054d61ab4805d07fbb339906b1258
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
         Step 1 Configure LDP LSPs.
                 After you complete the task described in 3.6.4 Example for Configuring Local
                 LDP Sessions, each LSR uses the default LDP LSP triggering policy. That is, each
                 LSR uses host IP routes with 32-bit addresses to trigger LDP LSP establishment.
                 # Run the display mpls ldp lsp command on each LSR. The command output
                 shows the LSR has successfully established LDP LSPs for all host routes.
                 The following example uses the command output on LSRA.
                 [LSRA] display mpls ldp lsp

                   LDP LSP Information
                   -------------------------------------------------------------------------------
                   Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                   -------------------------------------------------------------------------------
                   DestAddress/Mask In/OutLabel UpstreamPeer NextHop                              OutInterface
                   -------------------------------------------------------------------------------
                   1.1.1.9/32        3/NULL          2.2.2.9        127.0.0.1        LoopBack1
                  *1.1.1.9/32         Liberal/3                  DS/2.2.2.9
                   2.2.2.9/32        NULL/3          -            10.1.1.2        Vlanif100
                   2.2.2.9/32        1024/3          2.2.2.9       10.1.1.2        Vlanif100
                   3.3.3.9/32        NULL/1025         -            10.1.1.2        Vlanif100
                   3.3.3.9/32        1025/1025         3.3.3.9       10.1.1.2        Vlanif100
                   ------------------------------------------------------------------------------
                   TOTAL: 5 Normal LSP(s) Found.
                   TOTAL: 1 Liberal LSP(s) Found.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                68
MPLS Configuration
MPLS Configuration                                                                                    3 MPLS LDP Configuration

                  TOTAL: 0 Frr LSP(s) Found.
                  An asterisk (*) before an LSP means the LSP is not established
                  An asterisk (*) before a Label means the USCB or DSCB is stale
                  An asterisk (*) before an UpstreamPeer means the session is stale
                  An asterisk (*) before a DS means the session is stale
                  An asterisk (*) before a NextHop means the LSP is FRR LSP

                         NOTE

                        The default triggering policy is recommended, as this allows a device to use host IP routes
                        with 32-bit addresses to trigger LDP LSP establishment. You can also perform the following
                        steps to change the policy for triggering LDP LSP establishment as required.

         Step 2 Change the policy for triggering LDP LSP establishment.

                 Change the policy for triggering LDP LSP establishment to all on each LSR so that
                 the LSR uses all static routes and IGP routes in the routing table to trigger LDP
                 LSP establishment.

                 # Configure LSRA.
                 [LSRA] mpls
                 [LSRA-mpls] lsp-trigger all
                 [LSRA-mpls] quit

                 # Configure LSRB.
                 [LSRB] mpls
                 [LSRB-mpls] lsp-trigger all
                 [LSRB-mpls] quit

                 # Configure LSRC.
                 [LSRC] mpls
                 [LSRC-mpls] lsp-trigger all
                 [LSRC-mpls] quit

         Step 3 Verify the configuration.

                 # After completing the configuration, run the display mpls ldp lsp command on
                 each node to check information about LDP LSPs. The following example uses the
                 command output on LSRA.
                 [LSRA] display mpls ldp lsp
                  LDP LSP Information
                   -------------------------------------------------------------------------------
                   Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                   -------------------------------------------------------------------------------
                   DestAddress/Mask In/OutLabel UpstreamPeer NextHop                         OutInterface
                   -------------------------------------------------------------------------------
                   1.1.1.9/32        3/NULL          2.2.2.9       127.0.0.1 LoopBack1
                  *1.1.1.9/32         Liberal/3                 DS/2.2.2.9
                   2.2.2.9/32        NULL/3          -           10.1.1.2 Vlanif100
                   2.2.2.9/32        1024/3         2.2.2.9       10.1.1.2 Vlanif100
                   3.3.3.9/32        NULL/1025          -          10.1.1.2 Vlanif100
                   3.3.3.9/32        1025/1025         2.2.2.9      10.1.1.2 Vlanif100
                   10.1.1.0/30        3/NULL          2.2.2.9       10.1.1.1 Vlanif100
                  *10.1.1.0/30        Liberal/3                  DS/2.2.2.9
                   10.2.1.0/30        NULL/3          -           10.1.1.2 Vlanif100
                   10.2.1.0/30        1026/3        2.2.2.9        10.1.1.2 Vlanif100
                   -------------------------------------------------------------------------------
                   TOTAL: 8 Normal LSP(s) Found.
                   TOTAL: 2 Liberal LSP(s) Found.
                   TOTAL: 0 Frr LSP(s) Found.
                   An asterisk (*) before an LSP means the LSP is not established
                   An asterisk (*) before a Label means the USCB or DSCB is stale
                   An asterisk (*) before an UpstreamPeer means the session is stale


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                               69
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                  An asterisk (*) before a DS means the session is stale
                  An asterisk (*) before a NextHop means the LSP is FRR LSP

                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                          lsp-trigger all
                        #
                        mpls ldp
                         #
                          ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.3
                        #
                        return

