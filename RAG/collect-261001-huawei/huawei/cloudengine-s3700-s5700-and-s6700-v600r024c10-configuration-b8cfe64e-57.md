---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-57
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory", "throughput"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [7945, 8117]
sha256: ed586fe3628ae34079c0269f3e0ac2b1275f066bca6e8acdf6dcc00e0f1edca0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Procedure
                    ●    Clear the statistics about packets forwarded and discarded on a specified
                         interface for which traffic policing is configured to implement interface-based
                         rate limiting.
                         reset qos car statistics interface { interface-type interface-number | interface-name } inbound

                    ●    Clear queue-based traffic statistics.
                         reset qos queue statistics { interface { interface-type interface-number | interface-name } | slot slot-
                         id }


                    ----End




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                 141
QoS Configuration
QoS Configuration                                                 10 Congestion Avoidance Configuration




10                    Congestion Avoidance Configuration


                    10.1 Overview of Congestion Avoidance
                    10.2 Understanding Congestion Avoidance
                    10.3 Configuration Precautions for Congestion Avoidance
                    10.4 Default Settings for Congestion Avoidance
                    10.5 Configuring WRED
                    10.6 Configuring the CFI as the Internal Drop Priority
                    10.7 Verifying the Configuration
                    10.8 Maintaining Congestion Avoidance
                    10.9 Example for Configuring WRED


10.1 Overview of Congestion Avoidance
Definition
                    Congestion avoidance is a mechanism for controlling congestion. It monitors
                    network resources such as queues and memory buffers and discards packets when
                    congestion occurs or worsens.

Purpose
                    On traditional networks, quality of service (QoS) issues are mainly caused by
                    network congestion that arises from insufficient network resources. Congestion
                    extends the delay of packet transmission, lowers the throughput, and consumes
                    many resources. However, network congestion is very common in a complex
                    environment with a variety of services. Congestion avoidance uses specific packet
                    drop algorithms to prevent congestion from intensifying and make full use of
                    network bandwidth.


10.2 Understanding Congestion Avoidance

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          142
QoS Configuration
QoS Configuration                                                 10 Congestion Avoidance Configuration


10.2.1 WRED Fundamentals
                    Tail drop and Weighted Random Early Detection (WRED) are common methods
                    for congestion avoidance.

Tail Drop
                    Tail drop is the conventional method for discarding packets. When the length of a
                    queue reaches the maximum value, the device enabled with tail drop discards all
                    new packets buffered at the tail of the queue.
                    The tail drop method may cause global TCP synchronization, preventing TCP
                    connections from being set up. In the following figure, the three colors represent
                    three TCP connections. When packets from multiple TCP connections are
                    discarded, these TCP connections enter the congestion avoidance and slow start
                    state. Traffic volume decreases, before peaking once again. This is repeated over
                    and over, resulting in unstable traffic that is constantly changing in volume.

                    Figure 10-1 Tail drop




WRED
                    Random Early Detection (RED) is used to avoid global TCP synchronization that
                    occurs with tail drop. It does this by randomly discarding packets so that the
                    transmission speed of multiple TCP connections is not reduced simultaneously.
                    This results in more stable rates of TCP traffic and other network traffic.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           143
QoS Configuration
QoS Configuration                                                10 Congestion Avoidance Configuration


                    Figure 10-2 RED




                    However, RED does not accommodate QoS differentiation. Based on RED, WRED
                    uses drop profiles to implement congestion avoidance. A drop profile defines the
                    upper drop threshold, lower drop threshold, and drop probability. When a drop
                    profile is applied to an interface or an interface queue, packets are discarded
                    based on settings in the drop profile. When the length of a queue is smaller than
                    the lower drop threshold, the device does not discard packets. When the length of
                    a queue is between the lower drop threshold and the upper drop threshold, the
                    device randomly discards new packets, with a larger drop probability for longer
                    queues. When the length of a queue exceeds the upper drop threshold, the device
                    discards all new packets in the queue.
                    Figure 10-3 shows the curve of the WRED drop probability.

                    Figure 10-3 Curve of the WRED drop probability




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         144
QoS Configuration
QoS Configuration                                                  10 Congestion Avoidance Configuration


10.2.2 Fundamentals of CFI (Used as the Internal Drop
Priority)
                    Layer 2 devices exchange VLAN frames. As defined in IEEE 802.1Q, the Canonical
                    Format Indicator (CFI), also known as the Drop Eligible Indicator (DEI), in the
                    VLAN frame header identifies the drop priority of packets. Figure 10-4 shows the
                    CFI field in a VLAN frame.

                    Figure 10-4 CFI field in a VLAN frame




                    In the VLAN tag, the value of the CFI field is 0 or 1. When the rate of packets on
                    the device exceeds the committed information rate (CIR), the value of the DEI
                    field is set to 1, indicating a high drop priority. When congestion occurs, the device
                    first discards the packets with the DEI field value of 1.


10.3 Configuration Precautions for Congestion
Avoidance

10.4 Default Settings for Congestion Avoidance
                    Table 10-1 Default settings for congestion avoidance
                     Parameter                                  Default Setting

                     WRED                                       Disabled

                     Lower drop threshold, in percentage        100

                     Upper drop threshold, in percentage        100

                     Maximum drop probability, in               100
                     percentage

                     Absolute values of the upper and           Not configured
                     lower drop thresholds




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            145
QoS Configuration
QoS Configuration                                                               10 Congestion Avoidance Configuration




