---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-70
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [9940, 10045]
sha256: 7d8684eca8735c3e55834d3c3e7f6410c4f0bd8a7e9f7cdb010d6ecd148bcac1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                        In pipe mode, the ingress node does not copy the IP precedence, DSCP value,
                        or 802.1p value to the EXP field when a packet enters an MPLS network.
                        Likewise, when the packet leaves the MPLS network, the egress node does not
                        copy the EXP value to the IP precedence, DSCP value, or 802.1p value. If the
                        EXP value of a packet is changed on an MPLS network, the change is valid
                        only on the MPLS network. After the packet leaves the MPLS network, the
                        original CoS value carried in the packet remains valid.
                    ●   Short pipe mode
                        The short pipe mode is an enhancement of the pipe mode. In short pipe
                        mode, the ingress node of an MPLS network processes packets in the same
                        way as that in pipe mode. Where this mode differs is that the egress node of
                        an MPLS network removes the label before performing QoS scheduling. That
                        is, the egress node of an MPLS network schedules packets based on the
                        original CoS value carried in the packets, and performs QoS scheduling based
                        on your requirements from when the packets enter the MPLS network to the
                        penultimate hop.
                        On the L3VPN network shown in Figure 12-6, PE_1 sets the outer MPLS EXP
                        value to 1 and the inner MPLS EXP value to 2; P_2 removes the outer MPLS

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                         181
QoS Configuration
QoS Configuration                                                             12 MPLS QoS Configuration


                        label and sets the inner MPLS EXP value to the outer MPLS EXP value; PE_2
                        does not change the IP DSCP value of the original packet, and selects a PHB
                        based on the IP DSCP value of the packet.

                        Figure 12-6 Short pipe mode




                        In pipe or short pipe mode, you can plan QoS on the internal network based
                        on your requirements, without changing the QoS on the customer edge (CE)
                        side.
                        The difference between pipe mode and short pipe mode lies in how QoS
                        marking is implemented for outgoing traffic from a PE (PE_2 in this example)
                        to a CE. In pipe mode, QoS marking configured by you is used; in short pipe
                        mode, QoS marking configured on the CE side is used.
                        In Figure 12-7, CE_1 and CE_3 belong to VPN_1 and connect to two branches
                        of enterprise A; CE_2 and CE_4 belong to VPN_2 and connect to two branches
                        of enterprise B. When service flows from different VPNs enter the MPLS
                        network, devices on the MPLS network must differentiate priorities of the
                        services to ensure that service flows from enterprise A have higher priorities
                        than those from enterprise B. The devices then provide differentiated QoS
                        services to the service flows based on their priorities. In this case, you need to
                        specify the EXP values of packets in different VPNs when the packets enter
                        the MPLS network. Therefore, the pipe or short pipe mode is used.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             182
QoS Configuration
QoS Configuration                                                                12 MPLS QoS Configuration


                           Figure 12-7 Differentiating priorities of services in different VPNs




                    Table 12-1 Comparison between three tunnel modes
                     Tu      Whether the             Whether the      QoS               Wheth     Applica
                     nn      Ingress Node            Egress Node      Scheduling        er the    tion
                     el      Trusts the Original     Trusts the       Behavior on       Origin    Scenari
                     Mo      CoS Value Carried       EXP Value or     the Egress        al CoS    o
                     de      in an IP Packet         the Original     Node              Value
                                                     CoS Value                          Carried
                                                     Carried in an                      in an
                                                     IP Packet                          IP
                                                                                        Packet
                                                                                        Chang
                                                                                        es

                     Uni     Trusted. The CoS        The EXP value    Scheduling        Yes       Differen
                     for     value carried in an     is trusted and   based on                    tiating
                     m       IP packet is copied     mapped to        mapping of                  prioritie
                             to the EXP field in     the CoS value    the EXP value               s of
                             the outer MPLS          of an IP                                     differen
                             label.                  packet.                                      t
                                                                                                  services
                                                                                                  in a
                                                                                                  VPN

                     Pip     Untrusted. A new        The original     Scheduling        No        Differen
                     e       value can be            CoS value        based on                    tiating
                             assigned to the EXP     carried in an    mapping of                  prioritie
                             field in the outer      IP packet is     the EXP value               s of
                             MPLS label.             trusted and                                  services
                                                     retained.                                    in
                                                                                                  differen
                                                                                                  t VPNs




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                            183
QoS Configuration
QoS Configuration                                                            12 MPLS QoS Configuration


