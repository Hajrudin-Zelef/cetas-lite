---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-87
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [12128, 12264]
sha256: 77b165893b9e264e9575e803a08a7cf0c56c2081e38682b00c341537d9be7c85
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 4.     After receiving the Path message, PE2 learns from the Destination field in the
                        Session object that it is the egress of the CR-LSP to be established. PE2 then
                        allocates a label and reserves resources, and generates a Resv message based
                        on the local PSB. The Resv message carries the label allocated by PE2 and is
                        sent to P2.
                        PE2 extracts the address in the RSVP_HOP field of the received Path message
                        and uses the address as the destination IP address of the Resv message. The
                        Resv message does not carry the ERO field because it is forwarded along the
                        reverse path. Table 4-8 describes the objects in the Resv message.
                             NOTE

                            If a Resv message carries the RESV_CONFIRM object, the receiver needs to send a
                            ResvConf message to the sender to confirm the resource reservation request.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    206
MPLS Configuration
MPLS Configuration                                                                 4 MPLS TE Configuration


                        Table 4-8 Resv message on PE2

                         Object                                  Value

                         SESSION                                 Source: PE2-if0; Destination: PE1-if1

                         RSVP_HOP                                PE2-if0

                         LABEL                                   3

                         RECORD_ROUTE                            PE2-if0


                 5.     When receiving the Resv message, P2 creates an RSB according to the Resv
                        message, allocates a new label, updates the Resv message, and sends the
                        message to P1. Table 4-9 describes the objects in the Resv message.

                        Table 4-9 Resv message on P2

                         Object                                  Value

                         SESSION                                 Source: PE2-if0; Destination: PE1-if1

                         RSVP_HOP                                P2-if0

                         LABEL                                   17

                         RECORD_ROUTE                            P2-if0; PE2-if0


                 6.     After receiving the Resv message from P2, P1 creates an RSB according to the
                        Resv message, allocates a new label, updates the Resv message, and sends
                        the message to PE1. Table 4-10 describes the objects in the Resv message.
                        After receiving the Resv message from P1, PE1 obtains the label allocated by
                        P1, indicating that resources are successfully reserved. A CR-LSP is then
                        successfully established.

                        Table 4-10 Resv message on P1

                         Object                                  Value

                         SESSION                                 Source: PE2-if0; Destination: PE1-if1

                         RSVP_HOP                                P1-if0

                         LABEL                                   18

                         RECORD_ROUTE                            P1-if0; P2-if0; PE2-if0




Path Maintenance
                 Soft state

                 RSVP-TE is a soft-state protocol, which periodically updates RSVP-TE messages to
                 maintain the resource reservation states on nodes.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            207
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Resource reservation states include the path state and reservation state. RSVP
                 nodes along an established CR-LSP periodically send Path and Resv messages
                 (collectively called RSVP-TE Refresh messages) to maintain the path and
                 reservation states. RSVP Refresh messages are used to synchronize PSBs and RSBs
                 between RSVP-TE neighbors. If a node does not receive any Refresh message
                 about the path state or reservation state within a specified period, the node
                 deletes the state.
                 RSVP-TE refresh
                 RSVP-TE sends messages as IP datagrams. Therefore, the transmission of RSVP-TE
                 messages is unreliable. After a CR-LSP is established, the soft state mechanism
                 synchronizes the PSB and RSB between RSVP neighbors. Each node periodically
                 sends RSVP-TE Refresh messages to its upstream and downstream nodes.
                 A Refresh message is not a new message. It retransmits a previously advertised
                 message. The interval at which Refresh messages are sent is determined by the
                 TIME_VALUE Value field of a message.
                 If no Refresh message is received within the specified refresh interval, the state is
                 deleted.
                 A node can send Path and Resv messages to its neighbors in any sequence.
                 RSVP-TE Srefresh
                 In addition to state synchronization, RSVP-TE Refresh messages can also be used
                 to detect reachability between RSVP neighbors and maintain RSVP-TE neighbor
                 relationships. Because the sizes of Path and Resv messages used in the "soft state"
                 mechanism are large, sending many RSVP Refresh messages to establish a large
                 number of CR-LSPs occupies excess network resources. RSVP-TE Srefresh can
                 address this problem.
                 RSVP-TE Srefresh uses new objects defined in the original RSVP protocol and is
                 implemented depending on the following mechanisms:
                 ●      Message_ID extension and retransmission mechanisms
                        According to the Message_ID extension mechanism, RSVP-TE messages carry
                        extended objects, , including Message_ID and Message_ID_ACK objects. The
                        two objects are used to confirm RSVP-TE messages and support reliable RSVP-
                        TE message delivery.
                        The Message_ID object can also be used to provide the RSVP-TE
                        retransmission mechanism. For example, a node initializes a retransmission
                        interval as Rf seconds after it sends an RSVP-TE message carrying the
                        Message_ID object. If the node does not receive ACK message within Rf
                        seconds, the node retransmits the message after (1 + Delta) x Rf seconds.
                        Delta is determined by the rate at which the sender increases the
                        retransmission interval. The node keeps retransmitting the message until it
                        receives an ACK message or the retransmission times reach the threshold
                        (called a retransmission increment value).
                 ●      Srefresh
                        Srefresh messages can be sent instead of standard Path or Resv messages to
                        update RSVP-TE states. These messages reduce the amount of information
                        that must be transmitted and processed for maintaining RSVP-TE states.
                        When Srefresh messages are sent to update the RSVP-TE states, standard
                        Refresh messages are suppressed.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            208

