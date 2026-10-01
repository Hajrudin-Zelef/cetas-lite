---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-50
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [6819, 6968]
sha256: 323bea6de878198929557c101ca095a0c56a0ff2c64f58fb75801ac475adbfce
---

                 # Configure LSRC.
                 [LSRC] mpls lsr-id 3.3.3.9
                 [LSRC] mpls
                 [LSRC-mpls] quit
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] quit
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] mpls
                 [LSRC-Vlanif100] mpls ldp
                 [LSRC-Vlanif100] quit
                 [LSRC] interface vlanif 200
                 [LSRC-Vlanif200] mpls
                 [LSRC-Vlanif200] mpls ldp
                 [LSRC-Vlanif200] quit
                 [LSRC] interface vlanif 300
                 [LSRC-Vlanif300] mpls
                 [LSRC-Vlanif300] mpls ldp
                 [LSRC-Vlanif300] quit

                 # Configure LSRD.
                 [LSRD] mpls lsr-id 4.4.4.9
                 [LSRD] mpls
                 [LSRD-mpls] quit
                 [LSRD] mpls ldp
                 [LSRD-mpls-ldp] quit
                 [LSRD] interface vlanif 100
                 [LSRD-Vlanif100] mpls
                 [LSRD-Vlanif100] mpls ldp
                 [LSRD-Vlanif100] quit

                 # After completing the preceding configuration, run the display mpls lsp
                 command on LSRA to check information about established LSPs.
                 [LSRA] display mpls lsp
                 Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                 Flag after LDP FRR: (L) - Logic FRR LSP
                 -------------------------------------------------------------------------------
                              LSP Information: LDP LSP
                 -------------------------------------------------------------------------------
                 FEC              In/Out Label In/Out IF                        Vrf Name
                 1.1.1.9/32         3/NULL          -/-                      --
                 2.2.2.9/32         NULL/3          -/Vlanif100                  --
                 2.2.2.9/32         1024/3         -/Vlanif100                  --
                 3.3.3.9/32         NULL/3          -/Vlanif200                  --
                 3.3.3.9/32         1025/3         -/Vlanif200                  --
                 4.4.4.9/32         NULL/1026         -/Vlanif200                  --
                 4.4.4.9/32         1026/1026        -/Vlanif200                  --

                 The command output shows that LDP uses host routes with 32-bit addresses to
                 trigger LSP establishment by default.
         Step 4 Enable IS-IS Auto FRR on LSRA and check routing information and backup LSP
                information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        116
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


                 # Enable IS-IS Auto FRR on LSRA.
                 [LSRA] isis
                 [LSRA-isis-1] frr
                 [LSRA-isis-1-frr] loop-free-alternate
                 [LSRA-isis-1-frr] quit
                 [LSRA-isis-1] quit

                 # Check information about the direct routes between LSRA and LSRC and between
                 LSRC and LSRD.
                 [LSRA] display ip routing-table 10.1.4.0 verbose
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                  ------------------------------------------------------------------------------
                  Routing Table : _public_
                  Summary Count : 1

                  Destination: 10.1.4.0/24
                     Protocol: ISIS        Process ID: 1
                   Preference: 15                Cost: 20
                      NextHop: 10.1.2.2        Neighbour: 0.0.0.0
                       State: Active Adv           Age: 00h05m38s
                        Tag: 0             Priority: low
                       Label: NULL             QoSInfo: 0x0
                   IndirectID: 0x0
                  RelayNextHop: 0.0.0.0         Interface: Vlanif200
                     TunnelID: 0x0              Flags: D
                    BkNextHop: 10.1.1.2         BkInterface: Vlanif100
                      BkLabel: NULL          SecTunnelID: 0x0
                  BkPETunnelID: 0x0        BkPESecTunnelID: 0x0
                  BkIndirectID: 0x0

                 The command output shows that a backup IS-IS route is generated after IS-IS
                 Auto FRR is enabled.

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

                 The command output shows that LDP uses backup routes with 32-bit addresses to
                 trigger backup LSP establishment by default.

         Step 5 Configure LDP to trigger LSP establishment based on all routes. Check the LSP
                information.

                 # Run the lsp-trigger command on LSRA to allow LDP to use all routes to trigger
                 LSP establishment. Check the LSP information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         117
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration

                 [LSRA] mpls
                 [LSRA-mpls] lsp-trigger all
                 [LSRA-mpls] quit

                 # Run the lsp-trigger command on LSRB to allow LDP to use all routes to trigger
                 LSP establishment.
                 [LSRB] mpls
                 [LSRB-mpls] lsp-trigger all
                 [LSRB-mpls] quit

                 # Run the lsp-trigger command on LSRC to allow LDP to use all routes to trigger
                 LSP establishment.
                 [LSRC] mpls
                 [LSRC-mpls] lsp-trigger all
                 [LSRC-mpls] quit

                 # Run the lsp-trigger command on LSRD to allow LDP to use all routes to trigger
                 LSP establishment.
                 [LSRD] mpls
                 [LSRD-mpls] lsp-trigger all
                 [LSRD-mpls] quit

