---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-114
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [16107, 16251]
sha256: dbc7060378d4cc89c92ba1b82588a6181d17fe75ee37865aedc5eea668539569
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

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     271
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 ●      Run the display mpls rsvp-te peer [ interface { interface-type interface-
                        number | interface-name } | peer-address ] command to check information
                        about RSVP-TE neighbors on RSVP-TE-enabled interfaces.

4.9.5 Configuring the RSVP-TE Hello Extension Function
Prerequisites
                 Before configuring RSVP-TE Hello extension, complete the following task:
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 RSVP-TE Refresh messages can synchronize PSBs and RSBs between nodes, detect
                 reachability between RSVP-TE neighbors, and maintain RSVP-TE neighbor
                 relationships. This "soft state" mechanism detects neighbor relationships using
                 Path and Resv messages. The detection speed is low and a link failure cannot
                 promptly trigger a service traffic switchover. RSVP-TE Hello extension can address
                 this problem.
                 RSVP-TE Hello extension is used to rapidly detect reachability between RSVP-TE
                 nodes. It is usually used to trigger TE FRR path protection. When an RSVP-TE
                 neighbor node is unreachable, the related MPLS TE tunnel is torn down. This
                 mechanism can also be used to detect whether a neighboring node is in the
                 restart state. During the detection, the local node functions as the RSVP-TE GR
                 Helper to help the neighboring node implement RSVP-TE graceful restart (GR).
                 RSVP-TE Hello is implemented as follows:
                 1.     Hello handshake

                        Figure 4-14 Hello handshake mechanism




                        As shown in Figure 4-14, LSR1 and LSR2 are directly connected.
                        –   After RSVP-TE Hello is enabled on LSR1's interface, LSR1 sends a Hello
                            Request message to LSR2.
                        –   If LSR2 receives the Hello message and is enabled with RSVP-TE Hello,
                            LSR2 returns a Hello ACK message to LSR1.
                        –   After receiving the Hello ACK message from LSR2, LSR1 confirms that
                            LSR2 is reachable.
                 2.     Neighbor loss detection
                        After LSR1 sends a Hello Request messages to LSR2 and the handshake
                        process succeeds, they start to exchange Hello messages. If LSR1 does not
                        receive a Hello ACK message from LSR2 after sending three consecutive Hello
                        Request messages, LSR1 determines that LSR2 is unreachable, initiates TE FRR
                        switching, and restarts the RSVP Hello process.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             272
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                 3.     Neighbor restart detection
                        When both LSR1 and LSR2 are enabled with RSVP-TE GR, LSR1 waits for LSR2
                        to send a Hello Request message containing a GR extension after recognizing
                        that LSR2 is unreachable. After receiving the message, LSR1 begins to help
                        LSR2 in recovering the RSVP-TE state and replies with a Hello ACK message to
                        LSR2. When LSR2 receives the Hello ACK message from LSR1, it understands
                        that LSR1 is assisting in the GR process. Then, LSR1 and LSR2 exchange Hello
                        messages to maintain the restored GR state.

                            NOTE

                        If LSR1 and LSR2 belong to the same CR-LSP:
                        ●     If GR is disabled but TE FRR is enabled on a node, it switches traffic to a bypass CR-
                              LSP to ensure uninterrupted traffic transmission when detecting loss of the neighbor
                              node.
                        ●     If both GR and FRR are enabled, GR takes precedence over FRR.

                 RSVP-TE Hello applies to TE FRR and RSVP-TE GR Helper scenarios. Perform the
                 following configuration on each node of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enable the RSVP-TE Hello extension mechanism on the local node.
                 mpls rsvp-te hello

         Step 4 (Optional) Set the maximum number of Hello message losses that indicates that a
                peer node or link is faulty.
                 mpls rsvp-te hello-lost times

                 By default, after the Hello extension mechanism is enabled, a node or link is
                 considered faulty if it fails to respond to Hello messages for three consecutive
                 times. Possible outcomes are as follows:

                 ●      If RSVP-TE GR is deployed, the RSVP-TE GR mechanism is started to restore
                        the TE tunnel.
                 ●      If RSVP-TE GR is not deployed:
                        –     If TE FRR is deployed, traffic is switched to a bypass tunnel.
                        –     If TE FRR is not deployed, the corresponding TE tunnel is torn down and
                              reestablished.

         Step 5 (Optional) Set the interval at which Hello messages are refreshed.
                 mpls rsvp-te timer hello interval

                 By default, Hello messages are refreshed at an interval of 3s.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      273
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                         NOTE

                        A modified refreshing interval takes effect after the previous timer expires.
                        The Hello interval should not be less than the time a device takes to perform an active/
                        standby switchover. Otherwise, an intermittent protocol interruption may occur during a
                        switchover. The default value is recommended.

         Step 6 Return to the system view.
                 quit

         Step 7 Enter the MPLS TE link interface view.
                 interface interface-type interface-number

         Step 8 Switch the interface working mode from Layer 2 to Layer 3.
                 undo portswitch

