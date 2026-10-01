---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-46
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [6156, 6339]
sha256: ddf23d369199ec47f139b45f8ea77d781ad1d72a8c87082afc5905f99bf4c293
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The command output shows that the host routes to LSRB and LSRC have been
                 summarized.
         Step 4 Configure MPLS and MPLS LDP globally and on interfaces on each node so that
                the network can forward MPLS traffic. Then, check information about established
                LSPs.
                 # Configure LSRA.
                 [LSRA] mpls lsr-id 10.10.1.1
                 [LSRA] mpls
                 [LSRA-mpls] quit
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] quit
                 [LSRA] interface Vlanif100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] mpls ldp
                 [LSRA-Vlanif100] quit

                 # Configure LSRD.
                 [LSRD] mpls lsr-id 10.10.2.2
                 [LSRD] mpls
                 [LSRD-mpls] quit
                 [LSRD] mpls ldp
                 [LSRD-mpls-ldp] quit
                 [LSRD] interface Vlanif100
                 [LSRD-Vlanif100] mpls
                 [LSRD-Vlanif100] mpls ldp
                 [LSRD-Vlanif100] quit
                 [LSRD] interface Vlanif200
                 [LSRD-Vlanif200] mpls
                 [LSRD-Vlanif200] mpls ldp
                 [LSRD-Vlanif200] quit
                 [LSRD] interface Vlanif300
                 [LSRD-Vlanif300] mpls
                 [LSRD-Vlanif300] mpls ldp
                 [LSRD-Vlanif300] quit

                 # Configure LSRB.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         104
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration

                 [LSRB] mpls lsr-id 10.10.3.1
                 [LSRB] mpls
                 [LSRB-mpls] quit
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] quit
                 [LSRB] interface Vlanif100
                 [LSRB-Vlanif100] mpls
                 [LSRB-Vlanif100] mpls ldp
                 [LSRB-Vlanif100] quit

                 # Configure LSRC.
                 [LSRC] mpls lsr-id 10.10.3.2
                 [LSRC] mpls
                 [LSRC-mpls] quit
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] quit
                 [LSRC] interface Vlanif100
                 [LSRC-Vlanif100] mpls
                 [LSRC-Vlanif100] mpls ldp
                 [LSRC-Vlanif100] quit

                 # After completing the preceding configuration, run the display mpls lsp
                 command on LSRA to check information about established LSPs.
                 [LSRA] display mpls lsp
                 Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                  Flag after LDP FRR: (L) - Logic FRR LSP
                  -------------------------------------------------------------------------------
                               LSP Information: LDP LSP
                  -------------------------------------------------------------------------------
                  FEC              In/Out Label In/Out IF                       Vrf Name
                  10.10.2.2/32         NULL/3          -/Vlanif100
                  10.10.2.2/32         1024/3         -/Vlanif100

                 The preceding command output shows that by default, LDP does not establish
                 inter-area LSPs from LSRA to LSRB or from LSRA to LSRC.

         Step 5 Configure LDP extension for inter-area LSPs.

                 # Run the longest-match command on LSRA to enable LDP to use the longest
                 match rule to search for routes to establish LSPs.
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] longest-match
                 [LSRA-mpls-ldp] quit

         Step 6 Verify the configuration.

                 # After completing the preceding configuration, run the display mpls lsp
                 command on LSRA to check established LSPs.
                 [LSRA] display mpls lsp
                 Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                  Flag after LDP FRR: (L) - Logic FRR LSP
                  -------------------------------------------------------------------------------
                               LSP Information: LDP LSP
                  -------------------------------------------------------------------------------
                  FEC              In/Out Label In/Out IF                      Vrf Name
                  10.10.2.2/32         NULL/3          -/Vlanif100
                  10.10.2.2/32         1024/3         -/Vlanif100
                  10.10.3.1/32        NULL/1025           -/Vlanif100
                  10.10.3.1/32        1025/1025          -/Vlanif100
                  10.10.3.2/32        NULL/1026           -/Vlanif100
                  10.10.3.2/32        1026/1026          -/Vlanif100



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         105
MPLS Configuration
MPLS Configuration                                                                                    3 MPLS LDP Configuration


                 The preceding command output shows that LDP has established inter-area LSPs
                 from LSRA to LSRB and from LSRA to LSRC.
                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                         sysname LSRA
                        #
                        vlan batch 100
                        #
                         mpls lsr-id 10.10.1.1
                         mpls
                        #
                        mpls ldp
                         longest-match
                        #
                        isis 1
                         is-level level-2
                         network-entity 20.0010.0100.0001.00
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack0
                         ip address 10.10.1.1 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      LSRD
                        #
                         sysname LSRD
                        #
                        vlan batch 100 200 300
                        #
                         mpls lsr-id 10.10.2.2
                         mpls
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 10.0010.0200.0001.00
                         import-route isis level-1 into level-2 filter-policy ip-prefix permit-host
                         summary 10.10.3.0 255.255.255.0 avoid-feedback
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         isis enable 1
                         isis circuit-level level-2
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         isis enable 1
                         isis circuit-level level-1
                         mpls
                         mpls ldp


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              106
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration

