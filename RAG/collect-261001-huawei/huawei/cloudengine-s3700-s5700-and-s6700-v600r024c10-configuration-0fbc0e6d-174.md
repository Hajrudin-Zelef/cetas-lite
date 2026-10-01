---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-174
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [25267, 25400]
sha256: 2cc095a44433e5b1b2c5fe22a891e65508c5f3d9c71407b29071587ace0c8bdd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 3 Enable the egress to passively create BFD sessions.
                 # Configure LSR3.
                 [LSR3] bfd
                 [LSR3-bfd] mpls-passive
                 [LSR3-bfd] quit

                 ----End

Verifying the Configuration
                 # Check the BFD session states on LSR1.
                 [LSR1] display bfd session mpls-te interface Tunnel 1 te-lsp
                 --------------------------------------------------------------------------------
                 Local Remote        PeerIpAddr       State     Type        InterfaceName
                 --------------------------------------------------------------------------------
                 8192 8192          3.3.3.9       Up        D_TE_LSP Tunnel1
                 --------------------------------------------------------------------------------
                     Total UP/DOWN Session Number : 1/0

                 # Check the states of BFD sessions passively established on LSR3.
                 [LSR3] display bfd session passive-dynamic
                 --------------------------------------------------------------------------------
                 Local Remote        PeerIpAddr       State     Type        InterfaceName
                 --------------------------------------------------------------------------------
                 8192 8192          1.1.1.9       Up        E_Dynamic           -
                 --------------------------------------------------------------------------------
                     Total UP/DOWN Session Number : 1/0

                 Connect two ports on a tester (such as Port1 and Port2) to LSR1 and LSR3,
                 respectively. Inject MPLS traffic from Port1 to Port2. Ensure that label values are
                 set correctly. When the cable connected to 10GE1/0/1 on LSR1 or LSR2 is removed,
                 traffic is rapidly switched to the backup CR-LSP. The fault convergence time is at
                 the millisecond level.
                 To simulate a scenario in which the backup CR-LSP fails within the switchback
                 delay (15s) after the primary CR-LSP recovers and BFD rapidly detects the fault
                 and switches traffic back to the primary CR-LSP, re-insert the cable to 10GE1/0/1.
                 Run the display mpls te tunnel-interface tunnel 1 command repeatedly on LSR1
                 to check tunnel information until the primary CR-LSP is set up. Remove the cable
                 from 10GE1/0/2 on LSR1 or LSR4 within 15s to make the backup CR-LSP faulty. In
                 this case, traffic can be quickly switched back to the primary CR-LSP, and the fault
                 convergence time is at the millisecond level.

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 500
                        #
                        bfd
                        #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           421
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration

                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        explicit-path backup-path
                         next hop 10.1.5.2
                         next hop 10.1.3.1
                         next hop 3.3.3.9
                        #
                        explicit-path pri-path
                         next hop 10.1.1.2
                         next hop 10.1.2.2
                         next hop 3.3.3.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
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
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te tunnel-id 1
                         mpls te bfd enable
                         mpls te bfd min-tx-interval 500 min-rx-interval 500 detect-multiplier 3
                         mpls te record-route
                         mpls te path explicit-path pri-path
                         mpls te path explicit-path backup-path secondary
                         mpls te backup hot-standby mode revertive wtr 15
                         mpls te backup ordinary best-effort
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          mpls-te enable
                        #
                        return

                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100 200 400
                        #
                        mpls lsr-id 2.2.2.9


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                           422
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

