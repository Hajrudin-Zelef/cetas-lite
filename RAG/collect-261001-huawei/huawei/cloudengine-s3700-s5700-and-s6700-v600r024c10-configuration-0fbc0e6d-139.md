---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-139
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [19914, 20073]
sha256: 75a6cac8628ac306058dabfd1836a436621c891da035a8a282cf57064d178d76
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
                 ●      Configure a metric type for tunnels.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS TE tunnel interface view.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      333
MPLS Configuration
MPLS Configuration                                                                4 MPLS TE Configuration

                             interface tunnel tunnel-number

                        c.   Configure a metric type for tunnel path selection.
                             mpls te path metric-type { igp | te }

                             By default, the metric type configured in the MPLS view is used for tunnel
                             path selection.
                        d.   Return to the system view.
                             quit

                        e.   (Optional) Enter the MPLS view.
                             mpls

                        f.   (Optional) Configure a metric type for path selection.
                             mpls te path metric-type { igp | te }

                             If the mpls te path metric-type command is not run in the tunnel
                             interface view, the metric type configured in the MPLS view is used.
                             Otherwise, the metric type configured in the tunnel interface view is
                             used.

                             By default, the TE metric type is used.
                 ●      (Optional) Configure a link-specific TE metric.

                        If the TE metric type is specified, perform the following configuration on the
                        outbound interfaces of the ingress and transit node of the MPLS TE tunnel:

                        a.   Enter the system view.
                             system-view

                        b.   Enter the view of an MPLS TE-enabled interface.
                             interface interface-type interface-number

                        c.   Switch the interface mode from Layer 2 to Layer 3.
                             undo portswitch

                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                             S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                             be switched from Layer 2 mode to Layer 3 mode using the undo
                             portswitch command.

                             Determine whether to perform this step based on the current interface
                             mode.
                        d.   Configure a link-specific TE metric.
                             mpls te metric metric-value

                             By default, the IGP metric is used as the TE metric of a link.

                 ----End


Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information on the local node.




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                        334
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration




4.15 Configuring MPLS TE Tunnel Priorities
Prerequisites
                 Before configuring MPLS TE tunnel priorities, you have completed the following
                 task:
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 Priorities and preemption are used to prioritize the establishment of TE tunnels for
                 transmitting critical services, thereby preventing competition for resources during
                 the tunnel setup process. If there is no path meeting the bandwidth requirement
                 of a desired tunnel, a device tears down an existing path and uses bandwidth
                 resources assigned to that path to establish the desired tunnel. This is called
                 preemption.
                 Currently, the device supports only hard preemption. When a tunnel with a higher
                 priority competes for resources with a tunnel with a lower priority, the tunnel with
                 a higher priority directly preempts resources assigned to the tunnel with a lower
                 priority. As a result, some traffic on the tunnel with a lower priority is lost.
                 Tunnels use setup and holding priorities to determine whether to preempt
                 resources. Both the setup and holding priority values range from 0 to 7. The
                 smaller the value, the higher the priority. If only the setup priority is configured,
                 the value of the holding priority is equal to that of the setup priority. The setup
                 priority of a tunnel cannot be higher than the holding priority of this tunnel.
                 The setup and holding priorities are used in the following scenarios:
                 ●      If multiple LSPs are to be established, LSPs with higher setup priorities are
                        preferentially established by preempting resources.
                 ●      If bandwidth or other resources are insufficient, an LSP with a higher setup
                        priority but insufficient resources may preempt the resources of an established
                        LSP with a lower holding priority.
                 Figure 4-25 shows the bandwidth of each link. The following two TE tunnels are
                 established:
                 ●      Tunnel 1: established over the path LSR1 -> LSR6 -> LSR4. The required
                        bandwidth is 155 Mbit/s, and the setup and holding priorities are both 0.
                 ●      Tunnel 2: established over the path LSR2 -> LSR6 -> LSR3. The required
                        bandwidth is 155 Mbit/s, and the setup and holding priority values are both 7.
                 If the link between LSR6 and LSR4 fails after tunnel establishment, LSR1
                 recomputes a new path LSR1 -> LSR6 -> LSR3 -> LSR5 -> LSR4 for tunnel 1.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               335
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                 Figure 4-25 Preemption based on priorities




                 In hard preemption mode, LSR6 directly sends an RSVP-TE message to tear down
                 tunnel 2 because tunnel 1 has a higher priority than tunnel 2. As a result, some
                 traffic on tunnel 2 is dropped if this tunnel is transmitting traffic.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 3 Configure tunnel priorities.
                 mpls te priority setup-priority [ hold-priority ]

                 Both the setup and holding priority values range from 0 to 7. The smaller the
                 value, the higher the priority.

                 The default setup and holding priority values are both 7. If only the setup priority
                 is configured, the value of the holding priority is equal to that of the setup priority.

                         NOTE

                        The value of the setup priority must be greater than or equal to that of the holding priority.
                        This means that the setup priority cannot be higher than the holding priority.

                 ----End


