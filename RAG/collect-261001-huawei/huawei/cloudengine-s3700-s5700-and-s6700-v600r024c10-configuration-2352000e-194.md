---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-194
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "distribution", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [28365, 28494]
sha256: eef315231b803ba937876f13035ba41a3dac9da92a8cfa69e46b1c00b282d7bf
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


PWE3 VPWS
                    LDP VPWS can be either PWE3-compatible VPWS or PWE3 VPWS.

                    ●   PWE3-compatible VPWS does not use Label Notification messages.
                    ●   PWE3 VPWS uses Label Notification messages.

                    Pseudowire Emulation Edge-to-Edge (PWE3) VPWS is a Layer 2 service bearer
                    technology that simulates the basic behaviors and characteristics of services, such
                    as Ethernet, low-speed TDM circuit, and SONET/SDH, on a PSN.

                    PWE3 is an implementation of VPWS and an extension of LDP. PWE3 extends the
                    LDP signaling and reduces the signaling cost, improving networking flexibility.
                    Compared with LDP VPWS, PWE3 VPWS exchanges fewer packets when the
                    network is unstable, so there are fewer PW establishments and deletions.
                    ●   PWE3 networking modes
                        Currently, only the single-segment PW networking mode is supported. Single-
                        segment PW means that there is only one PW between every two PEs, and no
                        inner label switching is needed. Because the PW uses LDP as a signaling
                        protocol to transmit VC information, an LDP session must be established
                        between the PEs.
                        –    If there are Ps between the PEs, a remote LDP session must be created
                             between the PEs.
                        –    If the PEs are directly connected, a local LDP session can be established
                             between them.
                        Figure 5-20 shows the typical single-segment PWE3 networking with a PW
                        established using LDP signaling.

                        Figure 5-20 Single-segment PWE3 topology




                    ●   Dynamic PW
                        Dynamic PWs are established using signaling protocols. User-end provider
                        edges (UPEs) exchange VC labels using LDP and are bound to the
                        corresponding CEs based on the VC ID. A VC is established when all the
                        following conditions are satisfied: A tunnel has been established between two
                        PEs; the two PEs have exchanged labels and bound to the corresponding CEs
                        based on the VC ID; the ACs of the two PEs are up.
                        Messages used by a dynamic PW include:

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            453
VPN Configuration
VPN Configuration                                                                          5 VPWS Configuration


                        –    Label Request: requests the remote end to allocate labels.
                        –    Label Mapping: notifies the remote end of a label allocated by the local
                             end. Whether the Label Mapping message carries the Status field
                             depends on the default signaling. By default, LDP VPWS does not support
                             the Status field.
                        –    Label Notification: advertises and negotiates the PW status, reducing the
                             number of messages to be exchanged.
                        –    Label Withdraw: carries label and status information to instruct the
                             remote end to withdraw labels.
                        –    Label Release: responds to a Label Withdraw message, and notifies the
                             remote end that sends the Label Withdraw message of the label
                             withdrawal event.
                    ●   Establishment, maintenance, and deletion of dynamic PWs
                        Dynamic PWs are established using LDP, whose type-length-value (TLV) is
                        extended to carry VC information. Before a dynamic PW is established
                        between two PEs, an LDP session must be established between the two PEs.
                        During the establishment of a dynamic PW, the label distribution control
                        mode is downstream unsolicited (DU) and the label retention mode is liberal.
                              NOTE

                             If there are Ps between the PEs, a remote LDP session must be created between the
                             PEs. If the two PEs are directly connected, a local LDP session can be created between
                             them.
                        After PWE3 is configured on the two PEs (PE1 and PE2) and an LDP session is
                        established between the two PEs, the dynamic PW establishment starts.
                        Figure 5-21 shows the process of establishing a dynamic PW.
                        a.   PE1 sends a Label Request message and a Label Mapping message to
                             PE2.
                        b.   PE2 receives the Label Request message from PE1 and sends a Label
                             Mapping message to PE1.
                        c.   PE2 receives the Label Mapping message from PE1 and determines
                             whether its PW configuration (such as the VC ID, VC type, MTU, and
                             control word enabling status) is consistent with that on PE1. If so, PE2
                             sets the PW status to up.
                        d.   PE1 receives the Label Mapping message from PE2 and determines
                             whether its PW configuration is consistent with that on PE2. If so, PE1
                             sets the PW status to up. A dynamic PW is then established between PE1
                             and PE2.
                        e.   After the dynamic PW is established, PE1 and PE2 learn each other's
                             status by exchanging Label Notification messages.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      454
VPN Configuration
VPN Configuration                                                                5 VPWS Configuration


                        Figure 5-21 Process of establishing and maintaining a PWE3 VPWS PW




                        If the AC interface of a PW is down or the corresponding tunnel is down,
                        PWE3-compatible VPWS and PWE3 VPWS use different processing
                        mechanisms:
                        –   In PWE3-compatible mode, the local PE sends a Label Withdraw packet
                            to its peer to tear down the PW. After the AC interface or tunnel goes up,
                            another round of negotiation is required for the PEs to establish a PW.
                        –   In PWE3 mode, the local PE sends a Label Notification message to notify
                            its peer that data packets cannot be forwarded, but the PW is not torn
                            down. After the AC interface or tunnel goes up, the local PE sends a Label
                            Notification message to notify its peer that data packets can be
                            forwarded.
                        A PW is torn down only when the PW configuration is deleted or the LDP
                        session is interrupted. Label Notification messages prevent repeated PW
                        establishment and deletion caused by link flapping.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           455
VPN Configuration
VPN Configuration                                                                      5 VPWS Configuration


