---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-113
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [15979, 16106]
sha256: 33e6298aeef78d81524c8e0b15284d10d3d7930314e55dc706db8f7724791e92
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Verifying the Configuration
                 ●      Run the display mpls rsvp-te command to check the RSVP-TE configuration.
                 ●      Run the display default-parameter mpls rsvp-te command to check default
                        MPLS RSVP-TE parameters.
                 ●      Run the display mpls rsvp-te session ingress-lsr-id tunnel-id egress-lsr-id
                        command to check all information about a specified RSVP-TE session.
                 ●      Run the display mpls rsvp-te psb-content [ ingress-lsr-id tunnel-id [ lsp-id ]]
                        command to check RSVP-TE PSB information.
                 ●      Run the display mpls rsvp-te rsb-content [ ingress-lsr-id tunnel-id lsp-id ]
                        command to check RSVP-TE RSB information.
                 ●      Run the display mpls rsvp-te statistics { global | interface { interface-type
                        interface-number | interface-name } } command to check RSVP-TE statistics.
                 ●      Run the display mpls rsvp-te peer [ interface { interface-type interface-
                        number | interface-name } | peer-address ] command to check information
                        about RSVP-TE neighbors on RSVP-TE-enabled interfaces.

4.9.4 Configuring RSVP-TE Srefresh
Prerequisites
                 Before configuring RSVP-TE Srefresh, complete the following task:
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 RSVP-TE Refresh messages can synchronize PSBs and RSBs between nodes, detect
                 reachability between RSVP-TE neighbors, and maintain RSVP-TE neighbor

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         269
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


                 relationships. Because the sizes of Path and Resv messages used in the soft state
                 mechanism are large, sending many RSVP Refresh messages to establish a large
                 number of CR-LSPs occupies excess network resources. RSVP-TE Srefresh can
                 resolve this problem.
                 RSVP-TE Srefresh uses new objects defined in the original RSVP protocol and is
                 implemented as follows:
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
                 Enabling Srefresh on interfaces connecting two RSVP neighboring nodes reduces
                 the network cost and improves network performance.
                 Perform the following configuration on each node of an MPLS TE tunnel.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            270
MPLS Configuration
MPLS Configuration                                                                         4 MPLS TE Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enable Srefresh.
                 mpls rsvp-te srefresh

                 Srefresh enabled in the MPLS view takes effect globally.

         Step 4 (Optional) Configure retransmission parameters.
                 mpls rsvp-te timer retransmission { increment-value incvalue | retransmit-value retvalue }*

                 By default, the retransmission increment is 1, and the retransmission timer interval
                 is 5000 ms.

         Step 5 (Optional) Enable Srefresh forward compatibility.

                 This function can be configured in the interface view of an MPLS TE tunnel link.
                 quit
                 interface interface-type interface-number
                 mpls rsvp-te srefresh compatible

                 When the primary and backup CR-LSPs share the same link, the node on both
                 ends of the link may run different versions. If the upstream node runs V200 but
                 the downstream node runs V600, Srefresh incompatibility occurs. To address this
                 problem, run the mpls rsvp-te srefresh compatible command to enable Srefresh
                 compatibility on the interface that connects to the device running the earlier
                 version.

                         NOTE

                        This command can only be run on a downstream node running V600 when its upstream
                        node runs V200, which ensures that Srefresh can be properly negotiated between the two
                        nodes.

                 ----End


