---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-51
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [6969, 7087]
sha256: 69cd6922c6603639e792bd9a3b1e3dcba25097cba61b19a10e4a4151aaef18b8
---

                 # Run the display mpls lsp command on LSRA to check information about
                 established LSPs.
                 [LSRA] display mpls lsp
                 Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                 Flag after LDP FRR: (L) - Logic FRR LSP
                 -------------------------------------------------------------------------------
                              LSP Information: LDP LSP
                 -------------------------------------------------------------------------------
                 FEC              In/Out Label In/Out IF                        Vrf Name
                 1.1.1.9/32         3/NULL          -/-                      --
                 2.2.2.9/32         NULL/3          -/Vlanif100                  --
                 2.2.2.9/32         23/3          -/Vlanif100                  --
                   **LDP FRR**       NULL/17          -/Vlanif200                   --
                   **LDP FRR**       23/17          -/Vlanif200                   --
                 3.3.3.9/32         NULL/18          -/Vlanif200                   --
                 3.3.3.9/32         24/18          -/Vlanif200                  --
                   **LDP FRR**       NULL/18          -/Vlanif100                   --
                   **LDP FRR**       24/18          -/Vlanif100                   --
                 4.4.4.9/32         NULL/3          -/Vlanif200                  --
                 4.4.4.9/32         25/3          -/Vlanif200                  --
                   **LDP FRR**       NULL/19          -/Vlanif100                   --
                   **LDP FRR**       25/19          -/Vlanif100                   --
                 10.1.1.0/24        3/NULL           -/-                      --
                 10.1.2.0/24        3/NULL           -/-                      --
                 10.1.3.0/24        NULL/3           -/Vlanif100                   --
                 10.1.3.0/24        28/3           -/Vlanif100                  --
                 10.1.3.0/24        NULL/3           -/Vlanif200                   --
                 10.1.3.0/24        28/3           -/Vlanif200                  --
                 10.1.4.0/24        NULL/3           -/Vlanif200                   --
                 10.1.4.0/24        29/3           -/Vlanif200                  --

                 The command output shows that LSPs are triggered by routes with 24-bit
                 addresses.
         Step 6 Configure a policy for triggering backup LSP establishment based all routes.
                 # Run the auto-frr lsp-trigger command on LSRA to allow LDP to establish
                 backup LSPs for all backup routes.
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] auto-frr lsp-trigger all
                 [LSRA-mpls-ldp] quit


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        118
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration


         Step 7 Verify the configuration.

                 # After completing the preceding configuration, run the display mpls lsp
                 command on LSRA to check information about backup LSPs.
                 [LSRA] display mpls lsp
                 Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                 Flag after LDP FRR: (L) - Logic FRR LSP
                 -------------------------------------------------------------------------------
                              LSP Information: LDP LSP
                 -------------------------------------------------------------------------------
                 FEC              In/Out Label In/Out IF                        Vrf Name
                 1.1.1.9/32         3/NULL          -/-                      --
                 2.2.2.9/32         NULL/3          -/Vlanif100                  --
                 2.2.2.9/32         23/3          -/Vlanif100                  --
                   **LDP FRR**       NULL/17          -/Vlanif200                   --
                   **LDP FRR**       23/17          -/Vlanif200                   --
                 3.3.3.9/32         NULL/18          -/Vlanif200                   --
                 3.3.3.9/32         24/18          -/Vlanif200                  --
                   **LDP FRR**       NULL/18          -/Vlanif100                   --
                   **LDP FRR**       24/18          -/Vlanif100                   --
                 4.4.4.9/32         NULL/3          -/Vlanif200                  --
                 4.4.4.9/32         25/3          -/Vlanif200                  --
                   **LDP FRR**       NULL/19          -/Vlanif100                   --
                   **LDP FRR**       25/19          -/Vlanif100                   --
                 10.1.1.0/24        3/NULL           -/-                      --
                 10.1.2.0/24        3/NULL           -/-                      --
                 10.1.3.0/24        NULL/3           -/Vlanif100                   --
                 10.1.3.0/24        28/3           -/Vlanif100                  --
                 10.1.3.0/24        NULL/3           -/Vlanif200                   --
                 10.1.3.0/24        28/3           -/Vlanif200                  --
                 10.1.4.0/24        NULL/3           -/Vlanif200                   --
                 10.1.4.0/24        29/3           -/Vlanif200                  --
                   **LDP FRR**       NULL/26          -/Vlanif100                   --
                   **LDP FRR**       29/26          -/Vlanif100                   --

                 The command output shows that a backup LSP is established for the primary LSP
                 that is on the path LSRA -> LSRC -> LSRD.

                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         lsp-trigger all
                        #
                        mpls ldp
                         #
                         ipv4-family
                          auto-frr lsp-trigger all
                        #
                        isis 1
                         frr
                          loop-free-alternate level-1
                          loop-free-alternate level-2
                         network-entity 10.0000.0000.0001.00
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        119
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

