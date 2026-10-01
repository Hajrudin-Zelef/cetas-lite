---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-187
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [27201, 27309]
sha256: af60313e5537f61e1afebe07f428e2d5e36e353a2587a205c857b587c10831b7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        If multiple bypass CR-LSPs are available on a node, the node selects a bypass
                        CR-LSP based on the following factors in sequence: bandwidth/non-
                        bandwidth protection, implementation mode, and protected object.
                        Bandwidth protection takes precedence over non-bandwidth protection,
                        manual protection takes precedence over automatic protection, and node
                        protection takes precedence over link protection. In Figure 4-48, two bypass
                        CR-LSPs are available. If both of them provide bandwidth protection and are
                        manually configured, bypass LSP 2 is bound to the primary CR-LSP. (Bypass
                        LSP 2 provides node protection, and bypass LSP 2 provides link protection.) If
                        bypass CR-LSP 1 protects bandwidth and bypass CR-LSP 2 does not, only
                        bypass CR-LSP 1 can be bound to the primary CR-LSP.
                        If a bypass CR-LSP is successfully bound to the primary CR-LSP, the NHLFE
                        entry of the primary CR-LSP records the index of the NHLFE entry of the
                        bypass CR-LSP and the label (inner label) allocated by the MP to the previous
                        node. The inner label is used to guide traffic forwarding during FRR switching.
                 3.     Fault detection
                        –   In link protection, a data link layer protocol is used to detect and notify
                            faults. The fault detection speed at the data link layer depends on link
                            types.
                        –   In node protection, a data link layer protocol is used to detect link faults.
                            If no link fault exists, the RSVP-TE Hello or BFD for RSVP mechanism or
                            both mechanisms are used to detect faults on protected nodes.
                        If a link or node fault is detected, FRR switching is triggered immediately.
                             NOTE

                            ● In node protection, only the protected node and the link between the protected
                              node and the PLR are protected. The PLR cannot detect faults in the link between
                              the protected node and MP.
                            ● Link fault detection, BFD, and RSVP-TE Hello mechanisms detect a failure at
                              descending speeds.
                 4.     Switchover

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                   453
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


                        A switchover is a process that switches both service traffic and RSVP messages
                        to a bypass CR-LSP and notifies the upstream node of the switchover when a
                        primary CR-LSP fails. During the switchover, the MPLS label nesting
                        mechanism is used. The PLR pushes the label that the MP assigns for the
                        primary CR-LSP as the inner label, and then pushes a backup label allocated
                        by the next node of the bypass CR-LSP as an outer label. The penultimate hop
                        along the bypass CR-LSP removes the outer label from the packet and
                        forwards the packet only with the inner label to the MP. As the inner label is
                        assigned by the MP to the previous node, the MP can forward the packet to
                        the next hop on the primary CR-LSP.
                        Assume that a primary CR-LSP and a bypass CR-LSP have been set up before
                        TE FRR switching. Figure 4-49 shows the labels assigned by each node on the
                        primary CR-LSP and forwarding actions. The bypass CR-LSP provides node
                        protection. If LSR3 or the link between LSR2 and LSR3 fails, traffic is switched
                        to the bypass CR-LSP. During the switchover, the PLR (LSR2) swaps label 1024
                        with label 1022 that the MP allocates to the previous node as the inner label,
                        and pushes label 34 that the next node of the bypass CR-LSP allocates to the
                        PLR as the outer label for forwarding. This ensures that the packet can still be
                        forwarded to the next hop when it reaches LSR4. For details about the
                        forwarding process, see "Packet forwarding after the TE FRR switchover" in
                        Figure 4-49.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            454
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                        Figure 4-49 Packet forwarding before and after a TE FRR switchover




                 5.     Switchback
                        After the switchover, the ingress of the primary CR-LSP attempts to reestablish
                        the primary CR-LSP. After the primary CR-LSP is successfully reestablished,
                        service traffic and RSVP messages are switched back from the bypass CR-LSP
                        to the primary CR-LSP. The reestablished CR-LSP is called a modified CR-LSP.
                        In this process, TE FRR uses the make-before-break mechanism, which
                        ensures that the original primary CR-LSP is deleted only after the modified
                        CR-LSP is set up.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           455
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                         NOTE

                        FRR does not support multi-point failures. If FRR switching occurs, data is switched from
                        the primary CR-LSP to the bypass CR-LSP. During data forwarding through the bypass CR-
                        LSP, the bypass CR-LSP must remain up. If the bypass CR-LSP fails during this period, the
                        protected data cannot be forwarded through MPLS. As a result, traffic is interrupted and
                        FRR fails. Even if the bypass CR-LSP is reestablished, it cannot forward data. Data
                        forwarding will be restored only after the primary CR-LSP recovers or is reestablished.


Other Functions
                 When TE FRR is in the FRR-in-use state, the RSVP messages sent by the transmit
                 interface do not carry the interface authentication TLV, and the receive interface
                 does not perform interface authentication on the RSVP messages that do not carry
                 the authentication TLV and are in the FRR-in-use state. In this case, you can
                 configure neighbor authentication.

