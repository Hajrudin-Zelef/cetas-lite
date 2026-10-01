---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-69
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [9811, 9939]
sha256: ce8368aaeb6bfd8efcc2eb9ef186c9ac31f79648c82b65c1f9937d83730128d4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

E-LSP Solution
                    IP QoS implements differentiated services by classifying traffic at the network
                    edge. Traffic is classified into multiple priorities or service classes. If packets are
                    marked using the IP precedence in the IP header, they can be classified into a
                    maximum of eight classes; if the DSCP value is used, they can be classified into a
                    maximum of 64 classes. On each node through which packets pass, the IP
                    precedence or DSCP value is checked to determine QoS requirements of the
                    packets.

                    On an MPLS network, however, a device does not check the contents of IP
                    headers. Therefore, traffic classification cannot be implemented based on the IP
                    precedence or DSCP value. Currently, the device uses the E-LSP solution to classify
                    traffic on an MPLS network.

                    The E-LSP solution uses the EXP field in the MPLS header (shown in Figure 12-1)
                    to carry differentiated service information. The device forwards packets based on
                    their labels and determines forwarding behaviors such as scheduling and policing
                    — otherwise known as per-hop behaviors (PHBs) — based on the EXP field. The
                    EXP value can be copied from the DSCP value or IP precedence of IP packets
                    transmitted over a label switched path (LSP), or alternatively be configured.

                    Figure 12-1 MPLS packet encapsulation format




                    In Figure 12-2, when MPLS packets enter LSR2, the packets are classified and EXP
                    values in the packets are mapped to internal class of service (CoS) values and
                    drop priorities of the device. After traffic classification is performed, QoS
                    operations such as traffic shaping, traffic policing, and congestion avoidance are
                    implemented on the MPLS network in the same manner as those on an IP
                    network. When packets leave LSR2, the EXP values carried in the packets are
                    retained.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               178
QoS Configuration
QoS Configuration                                                           12 MPLS QoS Configuration


                    Figure 12-2 E-LSP solution




MPLS DiffServ Tunnel Modes
                    The DiffServ model allows intermediate nodes in a DiffServ domain to check and
                    modify the IP precedence, DSCP value, 802.1p value, or EXP value, which are called
                    CoS values. As a result, the CoS values of packets may change during packet
                    transmission on both IP and MPLS networks. Therefore, when a packet enters an
                    MPLS network or passes from an MPLS network to an IP network, the MPLS edge
                    router needs to determine whether to trust the CoS information carried in the IP/
                    MPLS packet based on the configuration. Related standards define three MPLS
                    DiffServ tunnel modes: uniform, pipe, and short pipe.
                    ●   Uniform mode
                        When the CoS value (IP precedence, DSCP value, or 802.1p value) carried in a
                        packet from an IP network can be trusted, you can use the uniform mode. In
                        this mode, the MPLS ingress node copies the CoS value carried in the packet
                        to the EXP field in the outer MPLS label. In this manner, the same QoS
                        guarantee is provided on the MPLS network. When the packet leaves the
                        MPLS network, the egress node maps the EXP value to the IP precedence,
                        DSCP value, or 802.1p value in the IP packet.
                        On the L3VPN network shown in Figure 12-3, PE_1 maps the DSCP value of
                        an IP packet to the outer and inner MPLS EXP values (5); P_1 changes the
                        outer MPLS EXP value to 6; P_2 removes the outer MPLS label and sets the
                        inner MPLS EXP value to the outer MPLS EXP value; PE_2 changes the DSCP
                        value of the packet to 48 based on the mapping.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         179
QoS Configuration
QoS Configuration                                                             12 MPLS QoS Configuration


                        Figure 12-3 Uniform mode




                        As its name implies, uniform mode ensures the same priority of packets on
                        the IP and MPLS networks. That is, priority mapping is performed for packets
                        entering and leaving an MPLS network. However, if the EXP value of a packet
                        changes on an MPLS network, the PHB for the packet also changes after the
                        packet leaves the MPLS network. As a result, the original CoS value of the
                        packet cannot be retained.
                        On the network shown in Figure 12-4, two sites belong to different branches
                        of the same enterprise. The enterprise network transmits voice, video, and
                        data services, listed in descending order of priority. When traffic of different
                        VPN services enters the MPLS network, the priorities of the three types of
                        services need to be differentiated on the MPLS network to ensure services are
                        correctly prioritized. In addition, different QoS services are provided for the
                        three types of services based on their priorities. In this case, the priorities
                        carried in the original packets can be trusted, and so the uniform mode is
                        used.


                        Figure 12-4 Differentiating priorities of different services in a VPN




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            180
QoS Configuration
QoS Configuration                                                          12 MPLS QoS Configuration


                    ●   Pipe mode
                        When the CoS value carried in a packet from an IP network cannot be trusted,
                        you can use the pipe mode. In this mode, the CoS value carried in the packet
                        is ignored, and the ingress node of an MPLS network assigns a new value to
                        the EXP field in the outer MPLS label. The packet is scheduled from the
                        ingress node to the egress node according to your requirements. After leaving
                        the MPLS network, the packet is forwarded according to the original CoS
                        value.
                        On the L3VPN network shown in Figure 12-5, PE_1 sets the outer MPLS EXP
                        value to 1 and the inner MPLS EXP value to 2; P_2 removes the outer MPLS
                        label and sets the inner MPLS EXP value to the outer MPLS EXP value; PE_2
                        does not change the IP DSCP value of the original packet, and selects a PHB
                        based on the inner MPLS EXP value.

                        Figure 12-5 Pipe mode




