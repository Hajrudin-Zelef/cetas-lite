---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-94
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [13155, 13294]
sha256: b4de1e1249b8b627121596e7dd645131d3d6d705e09d94abc887f44d14dd3917
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Networking Requirements
                 On the network shown in Figure 4-7, there are two static CR-LSPs in opposite
                 directions: one from LSR1 to LSR3 and the other from LSR3 to LSR1. To implement
                 synchronous protection switching on the two ends, configure the two LSPs as
                 static bidirectional associated LSPs.

                 Figure 4-7 Network diagram of static bidirectional associated LSPs




Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Configure IP addresses for interfaces, including the loopback interfaces whose
                        addresses are to be used as MPLS LSR IDs.
                 2.     Configure forward and reverse MPLS TEs.
                 3.     Bind the static forward and reverse CR-LSPs.

Procedure
         Step 1 Configure interface IP addresses for the devices.

                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24
                 [LSR1-Vlanif100] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.9 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.

         Step 2 Configure LSR IDs and enable MPLS and MPLS TE globally on the nodes and on
                interfaces.

                 # Configure LSR1.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           224
MPLS Configuration
MPLS Configuration                                                                          4 MPLS TE Configuration

                 [LSR1] mpls lsr-id 1.1.1.9
                 [LSR1] mpls
                 [LSR1-mpls] mpls te
                 [LSR1-mpls] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls
                 [LSR1-Vlanif100] mpls te
                 [LSR1-Vlanif100] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.

         Step 3 Configure MPLS TE tunnel interfaces.

                 # On LSR1, configure an MPLS TE tunnel to LSR3.
                 [LSR1] interface Tunnel 10
                 [LSR1-Tunnel10] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel10] tunnel-protocol mpls te
                 [LSR1-Tunnel10] destination 3.3.3.9
                 [LSR1-Tunnel10] mpls te tunnel-id 100
                 [LSR1-Tunnel10] mpls te signal-protocol cr-static
                 [LSR1-Tunnel10] quit

                 # On LSR3, configure an MPLS TE tunnel to LSR1.
                 [LSR3] interface Tunnel 20
                 [LSR3-Tunnel20] ip address unnumbered interface loopback 1
                 [LSR3-Tunnel20] tunnel-protocol mpls te
                 [LSR3-Tunnel20] destination 1.1.1.9
                 [LSR3-Tunnel20] mpls te tunnel-id 200
                 [LSR3-Tunnel20] mpls te signal-protocol cr-static
                 [LSR3-Tunnel20] quit

         Step 4 Configure a static CR-LSP from LSR1 to LSR3.

                 # Configure LSR1 as the ingress of the static CR-LSP.
                 [LSR1] static-cr-lsp ingress tunnel-interface Tunnel10 destination 3.3.3.9 nexthop 10.1.1.2 out-label 20

                 # Configure LSR2 as the transit node of the static CR-LSP.
                 [LSR2] static-cr-lsp transit Tunnel10 incoming-interface vlanif 100 in-label 20 nexthop 10.1.2.2 out-
                 label 30

                 # Configure LSR3 as the egress of the static CR-LSP.
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

         Step 6 Bind the static forward and reverse CR-LSPs.

                 # Configure LSR1.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                            225
MPLS Configuration
MPLS Configuration                                                                                     4 MPLS TE Configuration

                 [LSR1] interface Tunnel 10
                 [LSR1-Tunnel10] mpls te reverse-lsp protocol static lsp-name Tunnel20

                 # Configure LSR3.
                 [LSR3] interface Tunnel 20
                 [LSR3-Tunnel20] mpls te reverse-lsp protocol static lsp-name Tunnel10

                 ----End

Verifying the Configuration
                 # Check detailed information about the reverse CR-LSP on LSR1.
                 [LSR1] display mpls te reverse-lsp verbose
                 -------------------------------------------------------------------------------
                              LSP Information: STATIC LSP
                 -------------------------------------------------------------------------------
                   Obverse Tunnel            : Tunnel10        //Tunnel interface of the forward LSP
                   Reverse LSP Name             : Tunnel20       //Name of the reverse LSP
                   Reverse LSP State          : Up          //State of the reverse LSP
                   Incoming Label            : 130
                   Incoming Interface          : Vlanif100

                 The command output shows that Tunnel10 and Tunnel20 have been bound to
                 each other.

