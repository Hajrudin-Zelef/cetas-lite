---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-13
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1152, 1319]
sha256: f6bcd8519b3b676eda911001b9770bb4ffe81e932bca3166c1b9d489e7517672
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 By default, the MPLS MTU size is equal to the interface MTU size.

                 ----End


2.5 Configuring MPLS TTL Processing Modes

2.5.1 Understanding MPLS TTL Processing Modes
                 When IP packets in a VPN instance enter an MPLS tunnel, the MPLS TTL
                 processing modes for private network labels and public network labels are as
                 follows:

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   16
MPLS Configuration
MPLS Configuration                                                         2 Basic MPLS Configuration


                 ●      The MPLS TTL processing modes for private network labels and public
                        network labels are both uniform.
                        When an IP packet passes through an MPLS network, the IP TTL is
                        decremented by one on the ingress and is mapped to the MPLS TTL of the
                        private network label. The MPLS TTL of the private network label is mapped
                        to the MPLS TTL of the public network label. Then, the MPLS TTL in the
                        packet is processed in standard mode. The egress removes the public network
                        label, copies and pastes the MPLS TTL in the public network label to that in
                        the private network label, and removes the private network label. The egress
                        then decrements the MPLS TTL by one and changes the IP TTL to the smaller
                        value between the MPLS TTL and IP TTL. Figure 2-11 shows MPLS TTL
                        processing modes.

                        Figure 2-11 MPLS TTL processing modes (1)




                 ●      The MPLS TTL processing mode for private network labels is pipe, and that for
                        public network labels is uniform.
                        When an IP packet passes through an MPLS network, the ingress decrements
                        the IP TTL and copies and pastes the MPLS TTL that is fixed at 255 in the
                        private network label to the MPLS TTL in the public network label. Then, the
                        MPLS TTL in the packet is processed in standard mode. The egress removes
                        the public network label, copies and pastes the MPLS TTL to that in the
                        private network label, and removes the private network label. The IP TTL is
                        decremented by one only by the egress. 2.5.1 Understanding MPLS TTL
                        Processing Modes shows MPLS TTL processing modes.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           17
MPLS Configuration
MPLS Configuration                                                         2 Basic MPLS Configuration


                        Figure 2-12 MPLS TTL processing modes (2)




                 ●      The MPLS TTL processing mode for private network labels is uniform, and
                        that for public network labels is pipe.
                        When an IP packet passes through an MPLS network, the ingress decrements
                        the IP TTL by one and copies and pastes the IP TTL to the MPLS TTL of the
                        private network label. The MPLS TTL of the public network label is fixed at
                        255. Then, the MPLS TTL in the packet is processed in standard mode. The
                        egress removes the public network label and then the private network label,
                        decrements the MPLS TTL by one, and changes the IP TTL to the smaller
                        value between the MPLS TTL and the IP TTL. Figure 2-13 shows MPLS TTL
                        processing modes.

                        Figure 2-13 MPLS TTL processing modes (3)




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          18
MPLS Configuration
MPLS Configuration                                                                2 Basic MPLS Configuration


                 ●      The MPLS TTL processing modes for private network labels and public
                        network labels are both pipe.
                        When an IP packet passes through an MPLS network, the IP TTL is
                        decremented by one on the ingress. The MPLS TTLs in the private network
                        label and public network label are fixed at 255. Then, the MPLS TTL in the
                        packet is processed in standard mode. The egress removes the public network
                        label and private network label in sequence. The IP TTL is decremented by
                        one only by the egress. Figure 2-14 shows MPLS TTL processing modes.

                        Figure 2-14 MPLS TTL processing modes (4)




2.5.2 Configuring MPLS TTL Processing Modes

Context
                 You can configure MPLS TTL processing modes on the ingress PE or egress PE.

                         NOTE

                        After the TTL mode of an MPLS public network or VPN is changed, the new mode takes
                        effect only for new MPLS LDP sessions. To make the change take effect for previously
                        established MPLS LDP sessions, run the reset mpls ldp command to reestablish the
                        sessions.


Procedure
                 ●      Configure a mode for processing MPLS LDP TTLs in public network labels.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS view.
                             mpls


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     19
MPLS Configuration
MPLS Configuration                                                                     2 Basic MPLS Configuration


                        c.   Configure a mode for processing MPLS LDP TTLs in public network labels.
                             mpls ldp ttl-mode { pipe | uniform }

                             By default, MPLS LDP TTLs in public network labels are processed in
                             uniform mode.
                 ●      Configure a mode for processing MPLS TTLs in private network labels.
                        a.   Enter the system view.
                             system-view

                        b.   Create a VPN instance and enter the VPN instance view.
                             ip vpn-instance vpn-instance-name

                        c.   Enable the VPN instance IPv4 address family and enter the VPN instance
                             IPv4 address family view.
                             ipv4-family unicast

                        d.   Configure a mode for processing MPLS TTLs in private network labels.
                             ttl-mode { pipe | uniform }

                             By default, MPLS TTLs in private network labels are processed in pipe
                             mode.
                 ----End


2.6 Configuring Alarm Thresholds for MPLS Resources

2.6.1 Configuring Alarm Thresholds for LDP LSPs
Context
                 To facilitate device operation and maintenance, configure alarm thresholds for
                 LDP LSPs. This enables the device to report an alarm when the LDP LSP usage
                 reaches the upper threshold and to clear the alarm when the usage falls below
                 the lower threshold.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure alarm thresholds for LDP LSPs.
                 mpls ldp-lsp-number threshold-alarm upper-limit upper-limit-value lower-limit lower-limit-value

