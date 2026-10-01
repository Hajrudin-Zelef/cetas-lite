---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-92
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [12832, 12976]
sha256: ea471768d29ade5f23ff60875b81c8ada7c7b955fd6ca5cbca7b4f4c111a55af
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                            218
MPLS Configuration
MPLS Configuration                                                                                4 MPLS TE Configuration

                 [LSR3] static-cr-lsp egress Tunnel10 incoming-interface vlanif 200 in-label 30

         Step 5 Configure a static CR-LSP from LSR3 to LSR1.
                 # Configure LSR3 as the ingress of the static CR-LSP.
                 [LSR3] static-cr-lsp ingress tunnel-interface Tunnel20 destination 1.1.1.9 nexthop 10.1.2.1 out-label
                 120

                 # Configure LSR2 as the transit node of the static CR-LSP.
                 [LSR2] static-cr-lsp transit Tunnel20 incoming-interface vlanif 200 in-label 120 nexthop 10.1.1.1 out-
                 label 130

                 # Configure LSR1 as the egress of the static CR-LSP.
                 [LSR1] static-cr-lsp egress Tunnel20 incoming-interface vlanif 100 in-label 130

                 ----End

Verifying the Configuration
                 # Check the tunnel interface state on LSR1.
                 [LSR1] display interface tunnel 10
                 Tunnel10 current state : UP
                 Line protocol current state : UP
                 ...

                 The command output shows that the tunnel interface state is up.
                 # Check MPLS TE tunnel information on LSR1.
                 [LSR1] display mpls te tunnel
                 * means the LSP is detour LSP
                 ------------------------------------------------------------------------------
                 Ingress LsrId Destination           LSPID In/Out Label         R Tunnel-name
                 ------------------------------------------------------------------------------
                 1.1.1.9         3.3.3.9        0      --/20          I Tunnel10
                 -             -            -      130/--          E Tunnel20
                 ------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress

                 # Check MPLS TE tunnel information on LSR2.
                 [LSR2] display mpls te tunnel
                 * means the LSP is detour LSP
                 ------------------------------------------------------------------------------
                 Ingress LsrId Destination           LSPID In/Out Label         R Tunnel-name
                 ------------------------------------------------------------------------------
                 -             -            -      20/30           T Tunnel10
                 -             -            -      120/130          T Tunnel20
                 ------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress

                 # Check MPLS TE tunnel information on LSR3.
                 [LSR3] display mpls te tunnel
                 * means the LSP is detour LSP
                 ------------------------------------------------------------------------------
                 Ingress LsrId Destination           LSPID In/Out Label         R Tunnel-name
                 ------------------------------------------------------------------------------
                 -             -            -      30/--          E Tunnel10
                 3.3.3.9         1.1.1.9        1      --/120          I Tunnel20
                 ------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress

                 # Check static CR-LSP information on LSR1.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          219
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration

                 [LSR1] display mpls static-cr-lsp
                 TOTAL        :2    STATIC CRLSP(S)
                 UP         :2     STATIC CRLSP(S)
                 DOWN          :0     STATIC CRLSP(S)
                 Name             FEC         I/O Label       I/O If        Status
                 Tunnel10        3.3.3.9/32      NULL/20        -/Vlanif100        Up
                 Tunnel20        -/32          130/NULL        Vlanif100/-       Up

                 # Check static CR-LSP information on LSR2.
                 [LSR2] display mpls static-cr-lsp
                 TOTAL        :2    STATIC CRLSP(S)
                 UP         :2     STATIC CRLSP(S)
                 DOWN          :0    STATIC CRLSP(S)
                 Name          FEC          I/O Label      I/O If       Status
                 Tunnel10      -/-           20/30    Vlanif100/Vlanif200      Up
                 Tunnel20      -/-         120/130     Vlanif200/Vlanif100      Up

                 # Check static CR-LSP information on LSR3.
                 [LSR3] display mpls static-cr-lsp
                 TOTAL        :2    STATIC CRLSP(S)
                 UP         :2     STATIC CRLSP(S)
                 DOWN          :0     STATIC CRLSP(S)
                 Name          FEC          I/O Label         I/O If        Status
                 Tunnel20      1.1.1.9/32      NULL/120         -/Vlanif200        Up
                 Tunnel10      -/-          30/NULL          Vlanif200/-        Up

                 When a static CR-LSP is used to establish an MPLS TE tunnel, packets on the
                 transit nodes and egress are forwarded directly based on the configured incoming
                 and outgoing labels. As such, no FEC information is displayed on LSR2 or LSR3
                 used in this example.

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                        #
                        static-cr-lsp ingress tunnel-interface Tunnel10 destination 3.3.3.9 nexthop 10.1.1.2 out-label 20
                        bandwidth ct0 0
                        #
                        static-cr-lsp egress Tunnel20 incoming-interface Vlanif100 in-label 130
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel10
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te signal-protocol cr-static


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                 220
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration

                         mpls te tunnel-id 100
                        #
                        return

