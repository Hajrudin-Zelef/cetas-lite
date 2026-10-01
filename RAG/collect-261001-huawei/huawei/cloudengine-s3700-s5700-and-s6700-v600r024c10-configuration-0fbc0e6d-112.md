---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-112
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [15821, 15978]
sha256: 61e62509f4d10f81a3cbac41bcf56dc5b9d40793317a613d8e55f9b1944612d0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Currently, the following reservation styles are supported:

                 ●      Fixed Filter (FF) style: reserves exclusive resources for each node that requests
                        resource reservation. This means that reserved resources vary according to CR-
                        LSPs on the same link.
                 ●      Shared Explicit (SE) style: reserves shared resources for a series of nodes that
                        request resource reservation. This means that reserved resources are shared
                        by certain CR-LSPs on the same link.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           266
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS TE tunnel interface view.
                 interface tunnel interface-number

         Step 3 Configure a resource reservation style.
                 mpls te resv-style { ff | se }

                 The SE style is used for tunnels established using the make-before-break
                 mechanism, whereas the FF style is seldom used.

                 ----End


Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface command on the ingress to check
                        tunnel interface information, including the resource reservation style.

4.9.2 Configuring RSVP-TE Resource Reservation Confirmation

Prerequisites
                 Before configuring RSVP-TE resource reservation confirmation, complete the
                 following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 RSVP TE supports diversified signaling parameters, which meet requirements of
                 reliability and network resources and some MPLS TE advanced features. For details
                 about RSVP-TE, see 4.2.3 RSVP-TE Message Format and 4.2.4 RSVP-TE
                 Fundamentals.

                 RSVP-TE resource reservation confirmation is used to check whether resources are
                 successfully reserved. Perform the following configuration on the egress of an
                 MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Confirm Reservation
                 mpls rsvp-te resvconfirm

                 After a node receives a Path message, it initiates reservation confirmation by
                 sending a Resv message carrying an object that requests reservation confirmation.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      267
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration


                         NOTE

                        Receiving ResvConf messages does not mean that resource reservation is successful. It
                        merely indicates that resources have been successfully reserved at the furthest upstream
                        node that receives this Resv message. These resources could be later preempted by other
                        applications.

                 ----End


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

4.9.3 Configuring the RSVP-TE State Timer

Prerequisites
                 Before configuring the RSVP-TE state timer, complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 In RSVP-TE, RSVP-TE messages are refreshed periodically to maintain the resource
                 reservation states on nodes, including path states and reservation states.

                 If no Refresh message is received in a consecutive period, the resource reservation
                 state is deleted. By configuring the RSVP-TE state timer, you can adjust the interval
                 at which RSVP-TE Refresh messages are sent, as well as modify the number of
                 retransmission attempts to change the timeout interval. The default settings are
                 recommended. Calculate the timeout interval using the following formula:

                 Timeout interval = (keep-multiplier-number + 0.5) x 1.5 x refresh-interval

                 keep-multiplier-number indicates the number of times that RSVP-TE Refresh
                 messages are retransmitted, and refresh-interval indicates the interval at which
                 RSVP-TE Refresh messages are sent.

                 Perform the following configuration on each node of an MPLS TE tunnel.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     268
MPLS Configuration
MPLS Configuration                                                                         4 MPLS TE Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Set the interval at which RSVP-TE Refresh messages are sent.
                 mpls rsvp-te timer refresh refresh-interval

                 The default interval is 30 seconds.

                         NOTE

                        If the interval is modified, the change will take effect only after the current timer expires.
                        You are advised not to set an excessively long interval or frequently change the interval.

         Step 4 Set the number of times for retransmitting RSVP-TE Refresh messages.
                 mpls rsvp-te keep-multiplier keep-multiplier-number

                 By default, the number of retransmission times for RSVP Refresh messages is 3.

                 ----End

