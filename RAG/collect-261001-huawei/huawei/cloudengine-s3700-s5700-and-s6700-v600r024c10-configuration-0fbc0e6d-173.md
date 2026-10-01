---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-173
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [25126, 25266]
sha256: da571053a3ab5ef584671d83a595941e5cae8fc15ea941f4099374e472a769a9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        Actual local interval for sending BFD packets = max {Locally configured interval for sending
                        BFD packets, Remotely configured interval for receiving BFD packets}; actual local interval
                        for receiving BFD packets = max {Remotely configured interval for sending BFD packets,
                        Locally configured interval for receiving BFD packets}; actual local detection interval =
                        Actual local interval for receiving BFD packets x Remotely configured detection multiplier
                        On the egress that passively creates a BFD session, the BFD parameters cannot be adjusted,
                        because the default values are the smallest values that can be set on the TE ingress. As
                        such, the BFD detection interval on the ingress and that on the egress of a TE tunnel are as
                        follows:
                        ● Actual detection interval on the ingress = Configured interval at which BFD packets are
                          received on the ingress x 3
                        ● Actual detection interval on the egress = Configured local interval at which BFD packets
                          are sent on the ingress x Configured detection multiplier on the ingress

                 Perform the following configuration on the ingress of an MPLS TE tunnel.


Procedure
                 ●      Adjust BFD parameters globally.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS view.
                             mpls

                        c.   Set BFD-related time parameters.
                             mpls te bfd { min-tx-interval tx-interval | min-rx-interval rx-interval | detect-multiplier
                             multiplier }

                 ●      Adjust BFD parameters on a tunnel interface.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS TE tunnel interface view.
                             interface tunnel interface-number

                        c.   Set BFD-related time parameters.
                             mpls te bfd { min-tx-interval tx-interval | min-rx-interval rx-interval | detect-multiplier
                             multiplier }


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                                 418
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


                            If min-tx-interval tx-interval configured on the local end is different
                            from min-rx-interval rx-interval configured on the remote end, the
                            larger value is used as the actual session parameter.

                            The actual detection multiplier is the detect-multiplier multiplier value
                            set on the remote end.

                 ----End

4.25.6 Verifying the Configuration

Procedure
                 ●      Run the display bfd session dynamic [ verbose ] command to check
                        dynamic BFD session information on a tunnel ingress.
                 ●      Run the display bfd session passive-dynamic [ peer-ip peer-ip remote-
                        discriminator discr-value ] [ verbose ] command to check information about
                        passively created BFD sessions on an egress.
                 ●      Run the following commands to check BFD statistics:
                        –   Run the display bfd statistics command to check all BFD statistics.
                        –   Run the display bfd statistics session dynamic command to check
                            statistics about all dynamic BFD sessions.
                 ●      Run the display mpls bfd session [ fec ip-address | nexthop ip-address |
                        outgoing-interface interface-type interface-number ] [ verbose ] command
                        to check information about dynamic BFD sessions for detecting MPLS faults.
                 ●      Run the display mpls bfd session protocol rsvp-te [ verbose ] command to
                        check information about dynamic BFD sessions for detecting RSVP-TE faults.

                 ----End

4.25.7 Example for Configuring Dynamic BFD for CR-LSP

Networking Requirements
                 On the MPLS VPN network shown in Figure 4-40, establish a TE tunnel from LSR1
                 to LSR3 and configure CR-LSP hot standby and best-effort path creation. Here:

                 ●      The path of the primary CR-LSP is LSR1 -> LSR2 -> LSR3.
                 ●      The path of the backup CR-LSP is LSR1 -> LSR4 -> LSR3.

                 The requirements are as follows: If the primary CR-LSP fails, traffic is switched to
                 the backup CR-LSP. After the primary CR-LSP recovers, traffic is switched back to
                 the primary CR-LSP after a 15-second delay. If both the primary and backup CR-
                 LSPs fail, traffic is switched to the best-effort path. Explicit paths can be
                 configured for the primary and backup CR-LSPs. A best-effort path can be
                 generated automatically. In this example, the best-effort path is LSR1 -> LSR4 ->
                 LSR2 -> LSR3. The calculated best-effort path varies according to the faulty node.

                 Configure dynamic BFD for CR-LSP to monitor the primary and backup CR-LSPs.
                 After the configuration is complete, the following objects should be achieved:

                 ●      If the primary CR-LSP fails, traffic is rapidly switched to the backup CR-LSP.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                419
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 ●      If the backup CR-LSP fails within the switchback delay (15s) after the primary
                        CR-LSP recovers, traffic is switched back to the primary CR-LSP.

                 Figure 4-40 Network diagram of dynamic BFD for CR-LSP




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure CR-LSP hot standby and best-effort path creation.
                 2.     On the ingress, enable BFD, configure dynamic BFD for CR-LSP, and specify
                        the local intervals at which BFD packets are sent and received, and the local
                        BFD detection multiplier.
                 3.     Enable the egress to passively create BFD sessions.

Procedure
         Step 1 Configure CR-LSP hot standby and best-effort path creation.
                 Configure the primary CR-LSP, backup CR-LSP, and best-effort path according to
                 4.23.5 Example for Configuring CR-LSP Hot Standby.
         Step 2 Configure dynamic BFD for CR-LSP on the ingress.
                 On the ingress, configure dynamic BFD for CR-LSP. Set the local intervals at which
                 BFD packets are sent and received to 500 milliseconds. Set the local BFD detection
                 multiplier to 3.
                 # Configure LSR1.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           420
MPLS Configuration
MPLS Configuration                                                                                  4 MPLS TE Configuration

                 [LSR1] bfd
                 [LSR1-bfd] quit
                 [LSR1] interface tunnel 1
                 [LSR1-Tunnel1] mpls te bfd enable
                 [LSR1-Tunnel1] mpls te bfd min-tx-interval 500 min-rx-interval 500 detect-multiplier 3

