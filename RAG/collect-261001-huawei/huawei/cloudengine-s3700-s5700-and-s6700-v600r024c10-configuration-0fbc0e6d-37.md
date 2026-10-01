---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-37
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [4746, 4874]
sha256: 5f829a1e08607014ad4c43015caa130ba22ac1722ca2a7aa2cee677a7e78634c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        After the label advertise { explicit-null | implicit-null | non-null } command is run, the
                        configuration takes effect for new LSPs. For an existing LDP LSP, you need to run the reset
                        mpls ldp command in the user view of the local node to make the configuration take
                        effect.

                 ----End

3.10.3 Configuring an LDP Label Advertisement Mode
Context
                 By default, a downstream device sends Label Mapping messages to an upstream
                 device. This means that if a fault occurs on the network, services can be rapidly
                 switched to the backup path, improving network reliability. Digital subscriber line
                 access multiplexers (DSLAMs) deployed on an MPLS network for user access,
                 however, have low performance. On a large-scale network, a DSLAM can be
                 configured to send Label Mapping messages to an upstream device only after
                 receiving requests for labels. This minimizes the number of unwanted MPLS
                 forwarding entries on the DSLAM and controls the establishment of LSPs.

Procedure
                 ●      Configure a label advertisement mode for the local LDP session.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       80
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Configure a label advertisement mode.
                             mpls ldp advertisement { dod | du }

                                   NOTE

                                 ● When multiple links exist between neighbors, all interfaces must use the
                                   same label advertisement mode.
                                 ● Modifying a configured label advertisement mode leads to the
                                   reestablishment of an LDP session, resulting in service interruptions.
                 ●      Configure a label advertisement mode for the remote LDP session.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the remote MPLS-LDP peer view.
                             mpls ldp remote-peer remote-peer-name
                        c.   Configure a label advertisement mode.
                             mpls ldp advertisement { dod | du }

                                   NOTE

                                 ● When the local and remote LDP sessions coexist, they must have the same
                                   label advertisement mode.
                                 ● Modifying a configured label advertisement mode leads to the
                                   reestablishment of an LDP session, resulting in service interruptions.

                 ----End

3.10.4 Configuring a Global LDP Label Distribution Control
Mode
Context
                 A label distribution control mode defines how an LSR distributes labels during the
                 establishment of an LSP.
                 There are two label distribution control modes:
                 ●      Label distribution control in independent mode
                        In independent label distribution control mode, a local LSR independently
                        distributes and binds a label to a FEC and notifies the upstream LSR of the
                        label without waiting for a label from the downstream LSR.
                        –    If the label advertisement mode is DU and the label distribution control
                             mode is independent, an LSR directly distributes a label to its upstream
                             LSR, without waiting for a label from the downstream LSR.
                        –    If the label advertisement mode is DoD and the distribution control mode
                             is independent, an LSR distributes a label to its upstream LSR after
                             receiving a label request from the upstream LSR, without waiting for a
                             label from the downstream LSR.
                 ●      Label distribution control in ordered mode
                        In ordered label distribution control mode, an LSR sends the label mapping of
                        a FEC to the upstream device only if the LSR has received the Label Mapping
                        message from the next hop of the FEC or if the LSR is the egress of the FEC.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    81
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


                        –   If the label advertisement mode is DU and the label distribution control
                            mode is ordered, an LSR distributes a label to its upstream device only
                            after receiving a Label Mapping message from the downstream device.
                        –   If the label advertisement mode is DoD and the label distribution control
                            mode is ordered, the downstream LSR of a directly connected LSR
                            distributes a label to the upstream LSR only after receiving a Label
                            Mapping message from the directly connected LSR.

                 An LDP label distribution control mode can be globally configured to enable a
                 local node to control the sequence of distributing labels to upstream nodes. The
                 default label distribution control mode is recommended.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Configure an LDP label distribution control mode.
                 label distribution control-mode { ordered | independent }

                 The default LDP label distribution control mode is ordered.

                 ----End

3.10.5 Configuring LDP Label Policies

Context
                 Generally, an LSR distributes labels to both upstream and downstream LDP peers,
                 speeding up LDP LSP convergence. If a device receives all Label Mapping messages
                 or sends Label Mapping messages to all peers, a large number of LSPs need to be
                 established. This wastes resources and potentially causes unstable device running.
                 To reduce the number of LSPs and memory consumption, configure the following
                 policies as needed:

