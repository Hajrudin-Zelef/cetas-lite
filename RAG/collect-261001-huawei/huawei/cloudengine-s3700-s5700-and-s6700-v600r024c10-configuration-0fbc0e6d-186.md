---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-186
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [27072, 27200]
sha256: e62bdfe2faf7a8b2b3cea0a84c5f083be9060a445eb335179f727d8eeec0d6e5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                  Protected       Node protection      On the network shown in Figure 4-47, the
                  object                               primary CR-LSP passes through LSR3, which
                                                       exists between the PLR (LSR2) and an MP
                                                       (LSR4). Bypass LSP 2 protects LSR3 and
                                                       therefore works in node protection mode.
                                                       Either node protection or link protection can
                                                       be prioritized on the ingress of a tunnel. If
                                                       the Node protection flag in the Flag field of
                                                       the SESSION_ATTRIBUTE object in a Path
                                                       message is set to 1, node protection is
                                                       prioritized. Otherwise, link protection takes
                                                       precedence.

                                  Link protection      On the network shown in Figure 4-47, the
                                                       PLR (LSR2) and an MP (LSR3) are directly
                                                       connected, and the primary CR-LSP passes
                                                       through the direct link between them. Bypass
                                                       LSP 1 protects this link and therefore works
                                                       in link protection mode.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             450
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration


                  Classified         Type                  Description
                  By

                  Bandwidth          Bandwidth             When bandwidth of a bypass CR-LSP is
                  guarantee          protection            higher than or equal to bandwidth of the
                                                           primary CR-LSP, the bypass CR-LSP protects
                                                           the path and bandwidth of the primary CR-
                                                           LSP.

                                     Non-bandwidth         If no bandwidth is configured for a bypass
                                     protection            CR-LSP, it only protects the path of the
                                                           primary CR-LSP.

                  Implementa         Manual mode           A bypass CR-LSP is manually configured and
                  tion                                     bound to a primary CR-LSP.

                                     Automatic mode        An auto FRR-enabled node automatically
                                                           establishes a bypass CR-LSP. A node
                                                           automatically establishes a bypass CR-LSP
                                                           and binds it to a primary CR-LSP only if the
                                                           primary CR-LSP requires FRR and the
                                                           topology meets FRR requirements.




                 Figure 4-47 TE FRR link and node protection




                         NOTE

                        A bypass CR-LSP supports a protection mode combination. For example, manual protection,
                        node protection, and bandwidth protection are implemented together on a bypass CR-LSP.


Implementation
                 TE FRR is implemented as follows:

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 451
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 1.     Primary CR-LSP establishment
                        A primary CR-LSP is established in a way similar to that of an ordinary CR-LSP.
                        The difference is that the ingress appends the following flags into the
                        Session_Attribute object in a Path message: Local protection, Label recording,
                        and SE style.

                        Table 4-27 Session_Attribute object
                         Attribute Name                         Description

                         Local protection                       Local protection flag, indicating that
                                                                the primary CR-LSP needs to be
                                                                bound to a bypass CR-LSP.

                         Label recording                        Route recording flag, indicating that
                                                                routes and labels are recorded
                                                                during tunnel establishment.

                         SE style                               SE style flag, indicating that the
                                                                primary tunnel supports the SE style.

                         Bandwidth protection                   Bandwidth protection that a bypass
                                                                CR-LSP provides for the primary CR-
                                                                LSP.


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
                        in Figure 4-48 provide link protection and node protection, respectively.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           452
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration


                        Figure 4-48 Binding between bypass and primary CR-LSPs in TE FRR




