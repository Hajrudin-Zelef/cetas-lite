---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-169
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [24475, 24589]
sha256: 22e5bd2410f9292598788d36eab053759e830690983547ace5cdbadf1a34db63
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The requirements are as follows: If the primary CR-LSP fails, traffic is switched to
                 the backup CR-LSP. After the primary CR-LSP recovers, traffic is switched back to
                 the primary CR-LSP after a 15-second delay. If both the primary and backup CR-
                 LSPs fail, traffic is switched to the best-effort path. Explicit paths can be
                 configured for the primary and backup CR-LSPs. A best-effort path can be
                 generated automatically. In this example, the best-effort path is LSR1 -> LSR4 ->
                 LSR2 -> LSR3. The calculated best-effort path varies according to the faulty node.

                 Configure two static BFD sessions to monitor the primary and backup CR-LSPs.
                 After the configuration, the following objects are achieved:

                 ●      If the primary CR-LSP fails, traffic is rapidly switched to the backup CR-LSP.
                 ●      If the backup CR-LSP fails within the switchback delay (15s) after the primary
                        CR-LSP recovers, traffic is switched back to the primary CR-LSP.


                 Figure 4-37 Network diagram of static BFD for CR-LSP




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                407
MPLS Configuration
MPLS Configuration                                                                         4 MPLS TE Configuration


Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure CR-LSP hot standby and best-effort path creation.
                 2.     Configure two BFD sessions on the ingress to monitor the primary and backup
                        CR-LSPs, respectively. Configure two BFD sessions on the egress to monitor IP
                        links (as long as routes from LSR3 -> LSR1 are reachable).

Procedure
         Step 1 Configure CR-LSP hot standby and best-effort path creation.
                 Configure the primary CR-LSP, backup CR-LSP, and best-effort path according to
                 4.23.5 Example for Configuring CR-LSP Hot Standby.
         Step 2 Configure static BFD for CR-LSP.
                 # Establish BFD sessions between LSR1 and LSR3 to monitor the primary and
                 backup CR-LSPs. Bind the BFD session on LSR1 to the CR-LSP. Bind the BFD session
                 on LSR3 to the IP link. Set the intervals for sending and receiving BFD packets to
                 500 milliseconds.
                 # Configure LSR1.
                 [LSR1] bfd
                 [LSR1-bfd] quit
                 [LSR1] bfd prilsp2lsrc bind mpls-te interface tunnel 1 te-lsp
                 [LSR1-bfd-lsp-session-prilsp2lsrc] discriminator local 139
                 [LSR1-bfd-lsp-session-prilsp2lsrc] discriminator remote 239
                 [LSR1-bfd-lsp-session-prilsp2lsrc] min-tx-interval 500
                 [LSR1-bfd-lsp-session-prilsp2lsrc] min-rx-interval 500
                 [LSR1-bfd-lsp-session-prilsp2lsrc] process-pst
                 [LSR1-bfd-lsp-session-prilsp2lsrc] quit
                 [LSR1] bfd backuplsp2lsrc bind mpls-te interface tunnel 1 te-lsp backup
                 [LSR1-bfd-lsp-session-backuplsp2lsrc] discriminator local 339
                 [LSR1-bfd-lsp-session-backuplsp2lsrc] discriminator remote 439
                 [LSR1-bfd-lsp-session-backuplsp2lsrc] min-tx-interval 500
                 [LSR1-bfd-lsp-session-backuplsp2lsrc] min-rx-interval 500
                 [LSR1-bfd-lsp-session-backuplsp2lsrc] process-pst
                 [LSR1-bfd-lsp-session-backuplsp2lsrc] quit

                 # Configure LSR3.
                 [LSR3] bfd
                 [LSR3-bfd] quit
                 [LSR3] bfd reversepri2lsra bind peer-ip 1.1.1.9
                 [LSR3-bfd-session-reversepri2lsra] discriminator local 239
                 [LSR3-bfd-session-reversepri2lsra] discriminator remote 139
                 [LSR3-bfd-session-reversepri2lsra] min-tx-interval 500
                 [LSR3-bfd-session-reversepri2lsra] min-rx-interval 500
                 [LSR3-bfd-session-reversepri2lsra] quit
                 [LSR3] bfd reversebac2lsra bind peer-ip 1.1.1.9
                 [LSR3-bfd-session-reversebac2lsra] discriminator local 439
                 [LSR3-bfd-session-reversebac2lsra] discriminator remote 339
                 [LSR3-bfd-session-reversebac2lsra] min-tx-interval 500
                 [LSR3-bfd-session-reversebac2lsra] min-rx-interval 500
                 [LSR3-bfd-session-reversebac2lsra] quit

                 ----End

Verifying the Configuration
                 # Check the BFD session states on LSR1.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                     408
MPLS Configuration
MPLS Configuration                                                                                  4 MPLS TE Configuration

                 [LSR1] display bfd session discriminator 139
                 --------------------------------------------------------------------------------
                 Local Remote        PeerIpAddr       State     Type        InterfaceName
                 --------------------------------------------------------------------------------
                 139 239          3.3.3.9        Up        S_TE_LSP Tunnel1
                 --------------------------------------------------------------------------------

                 [LSR1] display bfd session discriminator 339
                 --------------------------------------------------------------------------------
                 Local Remote        PeerIpAddr       State     Type        InterfaceName
                 --------------------------------------------------------------------------------
                 339 439          3.3.3.9        Up        S_TE_LSP Tunnel1
                 --------------------------------------------------------------------------------

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

