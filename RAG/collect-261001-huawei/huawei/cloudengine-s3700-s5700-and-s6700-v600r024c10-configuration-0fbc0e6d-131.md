---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-131
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [18621, 18755]
sha256: f0f99ffec5c694530947c65e44604e9ff0cf82ec880bd41204b2a517e4986425
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
                 ●      Enable BFD for RSVP globally.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS view.
                             mpls

                        c.   Enable BFD for RSVP globally.
                             mpls rsvp-te bfd all-interfaces enable

                             After this command is run in the MPLS view, BFD for RSVP is enabled on
                             all RSVP interfaces except those having BFD for RSVP blocked.
                        d.   (Optional) Block BFD for RSVP.
                             quit
                             interface interface-type interface-number
                             mpls rsvp-te bfd block

                             If some RSVP interfaces do not require BFD for RSVP, block this function
                             on these interfaces.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         313
MPLS Configuration
MPLS Configuration                                                                              4 MPLS TE Configuration


                                      NOTE

                             For an Ethernet interface, you also need to run the undo portswitch command to
                             switch the interface to Layer 3 mode.
                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                             S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                             Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                             whether to run this command based on the current interface mode.
                        e.   (Optional) Set global BFD session parameters.
                             mpls rsvp-te bfd all-interfaces { min-tx-interval tx-interval | min-rx-interval rx-interval |
                             detect-multiplier multiplier }*

                             By default, the local detection multiplier of a BFD session is 3, and the
                             minimum intervals for sending and receiving BFD packets are both 1000
                             ms.

                                      NOTE

                                  The BFD parameters actually used may be different from the ones configured.
                                  ●     Actual interval for the local device to send BFD packets = Max {Locally
                                        configured min-tx-interval, Remotely configured min-rx-interval}
                                  ●     Actual interval for the local device to receive BFD packets = Max {Remotely
                                        configured min-tx-interval, Locally configured min-rx-interval}
                                  ●     Actual detection interval of the local device = Max {Remotely configured
                                        min-tx-interval, Locally configured min-rx-interval} x Remotely configured
                                        BFD detection multiplier
                 ●      Enable BFD for RSVP on an RSVP interface.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the view of the specified RSVP interface.
                             interface interface-type interface-number

                        c.   Switch the interface mode from Layer 2 to Layer 3.
                             undo portswitch

                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                             S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                             be switched from Layer 2 mode to Layer 3 mode using the undo
                             portswitch command.

                             Determine whether to perform this step based on the current interface
                             mode.
                        d.   Enable BFD for RSVP on the interface.
                             mpls rsvp-te bfd enable

                        e.   (Optional) Configure BFD session parameters for the interface.
                             mpls rsvp-te bfd { min-tx-interval tx-interval | min-rx-interval rx-interval | detect-multiplier
                             multiplier }*

                             By default, the local detection multiplier of a BFD session is 3, and the
                             minimum intervals for sending and receiving BFD packets are both 1000
                             ms.



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                                   314
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration


                                    NOTE

                                The BFD parameters actually used may be different from the ones configured.
                                ●     Actual interval for the local device to send BFD packets = Max {Locally
                                      configured min-tx-interval, Remotely configured min-rx-interval}
                                ●     Actual interval for the local device to receive BFD packets = Max {Remotely
                                      configured min-tx-interval, Locally configured min-rx-interval}
                                ●     Actual detection interval of the local device = Max {Remotely configured
                                      min-tx-interval, Locally configured min-rx-interval} x Remotely configured
                                      BFD detection multiplier

                 ----End

4.12.4 Verifying the Configuration

Procedure
                 ●      Run the display mpls rsvp-te command to check the BFD for RSVP
                        configuration.
                 ●      Run the display mpls rsvp-te interface [ interface-type interface-number ]
                        command to check the BFD for RSVP configuration on the specified interface.
                 ●      Run the display mpls rsvp-te bfd session { all | interface interface-type
                        interface-number | peer ip-address } [ verbose ] command to check BFD for
                        RSVP session information.
                 ●      Run the display mpls rsvp-te peer [ interface { interface-type interface-
                        number | interface-name } | peer-address ] command to check information
                        about RSVP-TE neighbors on RSVP-TE-enabled interfaces.
                 ●      Run the display mpls rsvp-te statistics { global | interface { interface-type
                        interface-number | interface-name } } command to check RSVP-TE statistics.

                 ----End

4.12.5 Example for Configuring Dynamic BFD for RSVP

Networking Requirements
                 On the MPLS network shown in Figure 4-22, a Layer 2 device (Switch) is deployed
                 between P1 and P2. An MPLS TE tunnel is established between PE1 and PE2. TE
                 FRR is configured, using P1 as a PLR and PE2 as an MP. The primary CR-LSP is PE1
                 -> P1 -> Switch -> P2 -> PE2, and the bypass CR-LSP is P1 -> P3 -> PE2.

                 If the link between Switch and P2 fails, P1 cannot quickly detect the link fault of
                 P2 because Switch isolates the fault. Instead, P1 can detect the fault only through
                 RSVP-TE Hello messages.

