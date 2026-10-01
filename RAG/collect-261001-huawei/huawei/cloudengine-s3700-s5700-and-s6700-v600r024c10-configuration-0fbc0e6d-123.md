---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-123
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17357, 17466]
sha256: bde9c7f828b4cf8cf0a49fb1697a76b5940074221030afc9a08c4868bf14e21b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                  Configure the message        mpls rsvp-te                 By default, the size of
                  window function.             authentication               the RSVP-TE
                                               window-size window-          authentication message
                                               size                         window is 1. In this case,
                                                                            the local device can save
                                                                            only the latest, largest
                                                                            sequence number of an
                                                                            RSVP message sent from
                                                                            the neighbor.
                                                                            NOTE
                                                                             If the RSVP-TE interface
                                                                             type is Eth-Trunk, only one
                                                                             neighbor relationship can
                                                                             be established between
                                                                             RSVP-TE neighbors over a
                                                                             trunk link. RSVP-TE
                                                                             messages can be received
                                                                             from any member
                                                                             interface of the trunk.
                                                                             Because the messages are
                                                                             not received in sequence
                                                                             on the member interfaces,
                                                                             they may be disordered. To
                                                                             prevent this issue,
                                                                             configure an RSVP-TE
                                                                             message window. You are
                                                                             advised to set the window
                                                                             size to a value greater than
                                                                             32. If the size is too small,
                                                                             some received RSVP-TE
                                                                             messages may be beyond
                                                                             the window and are
                                                                             discarded. As a result, the
                                                                             RSVP-TE neighbor
                                                                             relationship is terminated.



         Step 5 Enter the MPLS view.
                 quit
                 mpls

         Step 6 (Optional) Set Challenge message parameters as required. For details, see Table
                4-14.

                 To improve the success rate of RSVP-TE authentication in different network
                 conditions, you can perform this step to adjust the interval for retransmitting
                 Challenge messages and the maximum number of retransmission times. If the
                 authentication messages between two nodes are out of order, one node sends a
                 Challenge message to the other node to request connection restoration. If the
                 node does not receive any Response message, it retransmits the Challenge
                 message after a retransmission interval expires. If the node still fails to receive any
                 Response message from the other node when the maximum number of
                 retransmission times is reached, the authentication fails. If the authentication
                 succeeds before the maximum number of retransmission times is reached, the
                 number of retransmission times is cleared for Challenge messages.

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                                293
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Table 4-14 Configuring Challenge message parameters

                  Operation                     Command                     Purpose

                  Configure the maximum         mpls rsvp-te challenge-     By default, the maximum
                  number of Challenge           lost max-miss-times         number of Challenge
                  messages allowed to be                                    messages allowed to be
                  dropped when being                                        dropped is 3.
                  sent from the
                  authenticated end to the
                  authenticator end during
                  RSVP-TE authentication.

                  Configure an interval for     mpls rsvp-te retrans-       By default, the interval
                  retransmitting Challenge      timer challenge             for retransmitting
                  messages.                     retransmission-interval     Challenge messages is
                                                                            1000 ms.



                 ----End


Verifying the Configuration
                 ●      Run the display mpls rsvp-te command to check interface-specific RSVP-TE
                        configurations.
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

4.10.3 Example for Configuring RSVP-TE Authentication

Networking Requirements
                 On the network shown in Figure 4-17, the member interfaces of Eth-Trunk 1
                 between LSR1 and LSR2 are 10GE1/0/1, 10GE1/0/2, and 10GE1/0/3. An MPLS TE
                 tunnel is established between LSR1 and LSR3 through RSVP-TE.

