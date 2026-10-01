---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-177
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [25746, 25865]
sha256: f7a20db343c36599e2d88b1b3e2ae2f4bfaad2d4d4c6a67c0b6b030102d973a0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 428
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 2.     Binding between a bypass CR-LSP and the primary CR-LSP
                        The process of searching for a proper bypass CR-LSP for a primary CR-LSP is
                        called binding. Only the primary CR-LSP with the "Local protection" flag can
                        trigger a binding process. The binding must be complete before a primary/
                        bypass CR-LSP switchover is performed. Before the binding, the node needs to
                        calculate the outbound interface of the bypass CR-LSP, next hop label
                        forwarding entry (NHLFE), LSR ID of the MP, label allocated by the MP, and
                        protection type based on the RRO in the Resv message.
                        The PLR already obtains the next hop (NHOP) and next NHOP (NNHOP) of
                        the primary CR-LSP. If the egress LSR ID of the bypass CR-LSP is the same as
                        the NHOP LSR ID, the bypass CR-LSP provides link protection. If the egress
                        LSR ID of the bypass CR-LSP is the same as the NNHOP LSR ID, the bypass
                        CR-LSP provides node protection. For example, bypass LSP 1 and bypass LSP 2
                        in Figure 4-43 provide link protection and node protection, respectively.

                        Figure 4-43 Binding between bypass and primary CR-LSPs in TE FRR




                        If multiple bypass CR-LSPs are available on a node, the node selects a bypass
                        CR-LSP based on the following factors in sequence: bandwidth/non-
                        bandwidth protection, implementation mode, and protected object.
                        Bandwidth protection takes precedence over non-bandwidth protection,
                        manual protection takes precedence over automatic protection, and node
                        protection takes precedence over link protection. In Figure 4-43, two bypass
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

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          429
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration


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
                        TE FRR switching. Figure 4-44 shows the labels assigned by each node on the
                        primary CR-LSP and forwarding actions. The bypass CR-LSP provides node
                        protection. If LSR3 or the link between LSR2 and LSR3 fails, traffic is switched
                        to the bypass CR-LSP. During the switchover, the PLR (LSR2) swaps label 1024
                        with label 1022 that the MP allocates to the previous node as the inner label,
                        and pushes label 34 that the next node of the bypass CR-LSP allocates to the
                        PLR as the outer label for forwarding. This ensures that the packet can still be
                        forwarded to the next hop when it reaches LSR4. For details about the
                        forwarding process, see "Packet forwarding after the TE FRR switchover" in
                        Figure 4-44.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                   430
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                        Figure 4-44 Packet forwarding before and after a TE FRR switchover




                 5.     Switchback
                        After the switchover, the ingress of the primary CR-LSP attempts to reestablish
                        the primary CR-LSP. After the primary CR-LSP is successfully reestablished,
                        service traffic and RSVP messages are switched back from the bypass CR-LSP
                        to the primary CR-LSP. The reestablished CR-LSP is called a modified CR-LSP.
                        In this process, TE FRR uses the make-before-break mechanism, which
                        ensures that the original primary CR-LSP is deleted only after the modified
                        CR-LSP is set up.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           431
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                         NOTE

