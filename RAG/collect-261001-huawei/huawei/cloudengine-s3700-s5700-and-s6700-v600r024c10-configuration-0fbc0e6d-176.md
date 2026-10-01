---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-176
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [25594, 25745]
sha256: 7132e45915c6eebfd229f86690a5a89e40e5c36a0ae370f3670a913af7d82764
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Related Concepts
                 On the network shown in Figure 4-41, TE FRR establishes a bypass CR-LSP for a
                 primary tunnel on each possible faulty link or node. One bypass CR-LSP can
                 protect multiple primary tunnels. This protection mode is also called facility
                 backup.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      425
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Figure 4-41 Networking diagram of TE FRR local protection




                 For details about TE FRR concepts involved in Figure 4-41, see Table 4-22.

                 Table 4-22 Concepts in TE FRR
                  Concept          Description

                  Primary CR-      Primary CR-LSP, which is the protected CR-LSP.
                  LSP

                  Bypass CR-       CR-LSP protecting the primary CR-LSP. A bypass CR-LSP and its
                  LSP              primary CR-LSP belong to different tunnels.
                                   A bypass CR-LSP is usually in idle state and does not forward
                                   service traffics. If a bypass CR-LSP needs to be used to
                                   independently forward service data while protecting the primary
                                   CR-LSP, sufficient bandwidth must be allocated to the bypass
                                   CR-LSP.

                  PLR              Point of local repair. It is the ingress of a bypass CR-LSP. It must
                                   reside on a primary CR-LSP, and can be the ingress or transit
                                   node of a primary CR-LSP, but cannot be the egress of a primary
                                   CR-LSP.

                  MP               Merge point. The egress of a bypass CR-LSP must be on the path
                                   of the primary CR-LSP and cannot be the ingress of the primary
                                   CR-LSP.




                 Table 4-23 describes TE FRR classification.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             426
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 Table 4-23 TE FRR classification
                  Classified      Type                 Description
                  By

                  Protected       Node protection      On the network shown in Figure 4-42, the
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

                                  Link protection      On the network shown in Figure 4-42, the
                                                       PLR (LSR2) and an MP (LSR3) are directly
                                                       connected, and the primary CR-LSP passes
                                                       through the direct link between them. Bypass
                                                       LSP 1 protects this link and therefore works
                                                       in link protection mode.

                  Bandwidth       Bandwidth            When bandwidth of a bypass CR-LSP is
                  guarantee       protection           higher than or equal to bandwidth of the
                                                       primary CR-LSP, the bypass CR-LSP protects
                                                       the path and bandwidth of the primary CR-
                                                       LSP.

                                  Non-bandwidth        If no bandwidth is configured for a bypass
                                  protection           CR-LSP, it only protects the path of the
                                                       primary CR-LSP.

                  Implementa      Manual mode          A bypass CR-LSP is manually configured and
                  tion                                 bound to a primary CR-LSP.

                                  Automatic mode       An auto FRR-enabled node automatically
                                                       establishes a bypass CR-LSP. A node
                                                       automatically establishes a bypass CR-LSP
                                                       and binds it to a primary CR-LSP only if the
                                                       primary CR-LSP requires FRR and the
                                                       topology meets FRR requirements.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            427
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration


                 Figure 4-42 TE FRR link and node protection




                         NOTE

                        A bypass CR-LSP supports a protection mode combination. For example, manual protection,
                        node protection, and bandwidth protection are implemented together on a bypass CR-LSP.


Implementation
                 TE FRR is implemented as follows:
                 1.     Primary CR-LSP establishment
                        A primary CR-LSP is established in a way similar to that of an ordinary CR-LSP.
                        The difference is that the ingress appends the following flags into the
                        Session_Attribute object in a Path message: Local protection, Label recording,
                        and SE style.

                        Table 4-24 Session_Attribute object

                         Attribute Name                              Description

                         Local protection                            Local protection flag, indicating that
                                                                     the primary CR-LSP needs to be
                                                                     bound to a bypass CR-LSP.

                         Label recording                             Route recording flag, indicating that
                                                                     routes and labels are recorded
                                                                     during tunnel establishment.

                         SE style                                    SE style flag, indicating that the
                                                                     primary tunnel supports the SE style.

                         Bandwidth protection                        Bandwidth protection that a bypass
                                                                     CR-LSP provides for the primary CR-
                                                                     LSP.



