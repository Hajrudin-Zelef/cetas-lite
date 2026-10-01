---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-213
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [30919, 31072]
sha256: fabca65a52aaf9111e3dadbaeafe0ac27b8abc42574ee8c9d2a15b228d3d24cf
---

                 # Configure the backup CR-LSP on LSR1.
                 [LSR1] interface tunnel 1
                 [LSR1-Tunnel1] mpls te backup ordinary
                 [LSR1-Tunnel1] mpls te path explicit-path backup-path secondary
                 [LSR1-Tunnel1] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      511
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration


         Step 8 Configure synchronization between TE FRR and CR-LSP backup on the ingress of
                the primary CR-LSP.
                 # Configure LSR1.
                 [LSR1] interface tunnel 1
                 [LSR1-Tunnel1] mpls te backup frr-in-use
                 [LSR1-Tunnel1] quit

                 # After the configuration is complete, check information about the primary CR-
                 LSP.
                 [LSR1] display mpls te tunnel-interface Tunnel 1
                    ----------------------------------------------------------------
                                          Tunnel1
                    ----------------------------------------------------------------
                    Tunnel State Desc : UP
                    Active LSP          : Primary LSP
                    Session ID          : 100
                    Ingress LSR ID        : 1.1.1.9        Egress LSR ID: 4.4.4.9
                    Admin State           : UP             Oper State : UP
                    Primary LSP State : UP
                    Main LSP State         : READY            LSP ID : 40

                 ----End

Verifying the Configuration
                 # Disable the protected outbound interface on LSR2.
                 [LSR2] interface vlanif 200
                 [LSR2-Vlanif200] shutdown
                 [LSR2-Vlanif200] quit

                 # Check tunnel states on the ingress LSR1.
                 [LSR1] display mpls te tunnel-interface
                    ----------------------------------------------------------------
                                          Tunnel1
                    ----------------------------------------------------------------
                    Tunnel State Desc : UP
                    Active LSP          : Ordinary LSP
                    Session ID          : 100
                    Ingress LSR ID        : 1.1.1.9        Egress LSR ID: 4.4.4.9
                    Admin State           : UP             Oper State : UP
                    Primary LSP State : UP
                    Main LSP State         : READY             LSP ID : 40
                    Modify LSP State : SETTING UP
                    Ordinary LSP State : UP
                    Main LSP State         : READY              LSP ID : 32774

                 The command output shows that the tunnel state is up, indicating that the
                 primary tunnel is in the FRR-in-use state and is setting up an ordinary backup CR-
                 LSP and restoring the primary CR-LSP. If the primary CR-LSP is faulty, the system
                 starts the TE FRR bypass tunnel (the primary CR-LSP is in the FRR-in-use state)
                 and attempts to restore the primary CR-LSP. At the same time, the system
                 attempts to set up a backup CR-LSP.

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 600


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                              512
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        mpls lsr-id 1.1.1.9
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        explicit-path backup-path
                         next hop 10.1.6.1
                         next hop 10.1.7.2
                         next hop 10.1.3.2
                         next hop 4.4.4.9
                        #
                        explicit-path pri-path
                         next hop 10.1.1.2
                         next hop 10.1.2.2
                         next hop 10.1.3.2
                         next hop 4.4.4.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif600
                         ip address 10.1.6.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 600
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 4.4.4.9
                         mpls te tunnel-id 100
                         mpls te record-route label
                         mpls te path explicit-path pri-path
                         mpls te path explicit-path backup-path secondary
                         mpls te fast-reroute
                         mpls te backup ordinary
                         mpls te backup frr-in-use
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.6.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100 200 400
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      513
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

