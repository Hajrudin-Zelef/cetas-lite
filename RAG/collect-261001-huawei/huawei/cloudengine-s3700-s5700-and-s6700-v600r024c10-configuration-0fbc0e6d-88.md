---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-88
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [12265, 12384]
sha256: 367bbb8347f6555590c38ad007652f2373ed06d4bfa5ad6023e3b2b6f96a8e48
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


                        Each Srefresh message carries a series of Message_ID objects to identify the
                        Path and Resv states to be refreshed. Srefresh requires the use of the
                        Message_ID extension. Only the states that have been advertised through
                        Path and Resv messages carrying Message_ID can be refreshed through the
                        Srefresh mechanism.
                        When a node receives an Srefresh message, it matches the message with the
                        local state block (PSB or RSB). If they match, the node does not change the
                        state. If the Message_ID is greater than that saved in the local state block, the
                        node sends a NACK message to the sender, refreshes the PSB or RSB based on
                        the Path or Resv-TE message, and updates the Message_ID.
                        Message_ID objects contain Message_ID sequence numbers. If an LSP
                        changes, the associated Message_ID sequence number increases. When
                        receiving a Path message, the node compares the sequence number of the
                        Message_ID with the sequence number of the Message_ID saved in its local
                        state block. If they are the same, the node does not change the state; if the
                        received sequence number is greater than the local one, the state has been
                        updated.

Error Notification
                 RSVP-TE uses the following messages to notify LSP errors.
                 ●      PathErr message: sent upstream by an RSVP-TE node if an error occurs during
                        the processing of a Path message. A PathErr message is forwarded upstream
                        by each transit node until it arrives at the ingress.
                 ●      ResvErr message: sent downstream by an RSVP-TE node if an error occurs
                        during the processing of a Resv message. After receiving the ResvErr message,
                        the transit node continues to forward the message to the downstream node
                        until the message reaches the egress.

Path Teardown
                 After a user instructs an ingress to delete a CR-LSP or the ingress receives a
                 PathErr message, the ingress sends a PathTear message to a downstream node.
                 The downstream node receives this message, tears down the CR-LSP, and replies
                 to the ingress with a ResvTear message.
                 Here:
                 ●      PathTear message: instructs a node to remove path information, performing
                        the reverse action of a Path message.
                 ●      ResvTear message: instructs a node to remove resource reservation status
                        information, performing the reverse action of a Resv message.

4.2.5 MPLS TE Make-Before-Break Mechanism
                 The make-before-break mechanism prevents traffic loss during a traffic switchover
                 between CR-LSPs when link or tunnel attributes change, improving MPLS TE
                 tunnel reliability.

Context
                 Any change in link or tunnel attributes causes a CR-LSP to be reestablished using
                 new attributes. Traffic is then switched from the previous CR-LSP to the new CR-

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            209
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 LSP. If a traffic switchover is triggered before the new CR-LSP is set up, some
                 traffic is lost. The make-before-break mechanism prevents traffic loss.

Implementation
                 The make-before-break mechanism sets up a new CR-LSP (also called modify LSP)
                 and switches traffic to the new CR-LSP before the original CR-LSP is torn down.
                 This mechanism helps minimize data loss and additional bandwidth consumption.
                 Make-before-break is implemented using the SE resource reservation style.
                 During establishment, the new LSP may compete with the original LSP for
                 bandwidth on some shared links. The new LSP cannot be established if it fails the
                 competition. The make-before-break mechanism allows the system to reserve
                 bandwidth used by the original CR-LSP for the new CR-LSP, without calculating the
                 reserved bandwidth on shared links. Additional bandwidth resources are consumed
                 only when some links on the new LSP do not overlap with those on the original
                 LSP.

                 Figure 4-4 Make-before-break mechanism




                 Assume that the maximum reservable bandwidth of all links is 60 Mbit/s, as
                 shown in Figure 4-4. A CR-LSP (Path1 in the figure) is established, with the
                 ingress being LSR1 and the egress being LSR4. The bandwidth of the CR-LSP is 40
                 Mbit/s.
                 Assume that traffic needs to be switched to Path2 in order to forward data traffic
                 through lightly loaded LSR5. The remaining reservable bandwidth on the path
                 LSR3 -> LSR4 is only 20 Mbit/s. The make-before-break mechanism can be used in
                 this situation to allow the new CR-LSP Path2 to use the LSR3 -> LSR4 link
                 bandwidth reserved for the original path. After the new path is established and
                 traffic is switched to it, the original path is torn down.
                 In addition to the preceding method, another method of increasing the tunnel
                 bandwidth can be used. If the reservable bandwidth of a shared link increases to a
                 certain extent, a new CR-LSP can be established.
                 Assume that the maximum reservable bandwidth of all links is 60 Mbit/s, as
                 shown in Figure 4-4. A CR-LSP (Path1 in the figure) is established, with the
                 ingress being LSR1 and the egress being LSR4. The bandwidth of the CR-LSP is 30
                 Mbit/s.
                 A new CR-LSP needs to be established along Path2 to forward data through lightly
                 loaded LSR5, and the path bandwidth needs to increase to 40 Mbit/s. The
                 remaining reservable bandwidth on the path LSR3 -> LSR4 is only 30 Mbit/s. The

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                             210
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 make-before-break mechanism can be used in this situation. This mechanism
                 allows Path2 (the new CR-LSP path) to use the bandwidth of the link LSR3 ->
                 LSR4 reserved for the original CR-LSP, and reserves an additional bandwidth of 10
                 Mbit/s for the new path. After the new CR-LSP is established and traffic is
                 switched to it, the original CR-LSP is torn down.


Delayed Switchover and Deletion
                 In real-world applications, the service status of each node on an MPLS network
                 varies. If an upstream node on an MPLS network is busy but its downstream node
                 is idle or an upstream node is idle but its downstream node is busy, a CR-LSP may
                 be torn down before the new CR-LSP is established, causing a temporary traffic
                 interruption.

