---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-68
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [9679, 9810]
sha256: cd5ee95c53f10047893d239d33ef2b1319805d4fa870183d1e3b504a7c56626b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration
QoS Configuration                                                       11 Congestion Management Configuration


                    fault locating methods, which are complex and difficult. With such traditional
                    methods, after packets are lost due to congestion in the outbound direction of an
                    interface, outgoing packets are obtained to analyze the traffic trend and identify
                    the characteristics of burst traffic. Microburst detection can be used to detect
                    instantaneous (millisecond-level) burst traffic in the outbound direction of an
                    interface, helping maintenance personnel determine whether packet loss is caused
                    by microbursts. Microburst detection helps to identify potential congestion risks
                    before congestion occurs and, if congestion does occur, quickly locate abnormal
                    traffic.
                    The following microburst detection modes are supported:
                    ●   Default mode: Packets are sampled at an interval of 5 ms. In this mode,
                        microburst detection can be enabled on multiple interfaces.
                    ●   Enhanced mode: Packets are sampled at an interval of 1 ms. In this mode,
                        microburst detection can be enabled on only one interface.
                    The measurement period of microburst detection is 5 minutes. That is, the key
                    performance indicators of interfaces are collected every 5 minutes and related
                    entries are generated accordingly. The device can store statistics collected up to
                    300 minutes after microburst detection is enabled.
                    The key performance indicators of microburst detection are as follows:
                    ●   Average rate of burst traffic forwarded to an interface from any other
                        interface on the same device
                    ●   Peak rate of burst traffic forwarded to an interface from any other interface
                        on the same device
                    ●   Number of discarded packets on an interface
                    ●   Average buffer usage of an interface
                    ●   Peak buffer usage of an interface
                    ●   Buffer usage of interface queues when the interface buffer usage reaches the
                        peak value in a measurement period

Procedure
                    ●   Configuring microburst detection on the device
                        a.   Enter the system view.
                             system-view
                        b.   Configure the microburst detection mode.
                             qos micro-burst detection [ enhanced ] enable
                    ●   Configuring microburst detection on an interface
                        a.   Enter the system view.
                             system-view
                        b.   Enter the interface view.
                             interface { interface-type interface-number | interface-name }
                        c.   Enable the microburst detection function on the interface.
                             qos micro-burst detection enable


11.12.2 Verifying the Configuration
                    After the microburst detection function is enabled on the device, you can check
                    burst traffic and packet loss information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                175
QoS Configuration
QoS Configuration                                                        11 Congestion Management Configuration


                    ●   Check microburst detection statistics on an interface.
                        display qos micro-burst statistics interface { interface-name | interface-type interface-number }

                    ●   Check all interfaces enabled with microburst detection and packet loss
                        information on the interfaces.
                        display qos micro-burst status all

                    ●   Check the peak buffer usage on an interface and the buffer usage of queues
                        on the interface.
                        display qos micro-burst peak-buffer verbose interface { interface-name | interface-type interface-
                        number }




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                 176
QoS Configuration
QoS Configuration                                                            12 MPLS QoS Configuration




                             12                 MPLS QoS Configuration


                    12.1 Overview of MPLS QoS
                    12.2 Understanding MPLS QoS
                    12.3 Configuration Precautions for MPLS QoS
                    12.4 Default Settings for MPLS QoS
                    12.5 Configuring a DiffServ Mode
                    12.6 Configuring Priority Mapping
                    12.7 Checking the MPLS QoS Configuration
                    12.8 Example for Configuring MPLS QoS


12.1 Overview of MPLS QoS
Definition
                    MPLS QoS is a specific QoS service that works through the DiffServ model on an
                    MPLS network. MPLS QoS provides differentiated services for packets passing
                    through the MPLS network to meet diversified requirements.

Purpose
                    MPLS uses label-based forwarding in place of traditional route-based forwarding.
                    It provides powerful and flexible functions to meet network requirements of
                    various new applications, and supports multiple network protocols such as IPv4
                    and IPv6. Currently, MPLS is widely used for building large networks. Because IP
                    QoS cannot be used on an MPLS network, MPLS QoS is used instead.
                    Traditional IP QoS differentiates service classes based on priorities of IP packets.
                    Similarly, MPLS QoS differentiates data flows based on the EXP value of packets to
                    implement differentiated services. MPLS QoS ensures low delay and low packet
                    loss rate for voice and video data flows, as well as high network usage.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           177
QoS Configuration
QoS Configuration                                                               12 MPLS QoS Configuration




12.2 Understanding MPLS QoS
                    On an MPLS network, MPLS QoS distinguishes data flows based on the EXP field
                    in the MPLS header and provides differentiated services. This is known as the E-
                    LSP solution. For packets that enter an MPLS network or pass from an MPLS
                    network to an IP network, you need to configure the device to determine whether
                    to trust the original IP precedence, DSCP value, 802.1p value, or EXP value carried
                    in the packets. Based on this, related standards define three MPLS DiffServ tunnel
                    modes: uniform, pipe, and short pipe.


