---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-26
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [2924, 3063]
sha256: a20c6d32ceb07a6628a1663b5840dff9e2c333809e164113eb9646da23baf918
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         Improper CPCAR settings will impact services on your network. To adjust the CPCAR value,
                         contact technical support personnel.


3.15.5 How Can I Identify and Prevent Common Attacks?

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     49
Security Configuration
Security Configuration                                              3 Local Attack Defense Configuration


                  You can detect common attacks as follows:
                  1.     Run the reset cpu-defend statistics command to clear statistics on the
                         packets sent to the CPU.
                  2.     Wait for one minute, and then run the display cpu-defend statistics
                         command to check the number of protocol packets sent to the CPU and the
                         number of discarded protocol packets, such as ICMP, TTL Expired, SSH, and
                         FTP. If a lot of packets are sent to the CPU or discarded, an attack, such as
                         ICMP attack, TTL Expired attack, SSH attack, or FTP attack, may occur.
                  3.     Identify the attack source through attack source tracing.
                  After identifying the attack source, you can run the cpu-defend policy command
                  to configure a filter or punishment to discard attack packets.
                  In addition, the device can rate-limit ICMP packets from this source, or use a
                  traffic policy to discard SSH and FTP attack packets based on your configuration.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               50
Security Configuration
Security Configuration                                                 4 Storm Suppression Configuration




                  4           Storm Suppression Configuration


                  4.1 Overview of Storm Suppression
                  4.2 Configuration Precautions for Storm Suppression
                  4.3 Default Settings for Storm Suppression
                  4.4 Configuring Traffic Suppression
                  4.5 Configuring Storm Control
                  4.6 Troubleshooting Storm Suppression


4.1 Overview of Storm Suppression
Definition
                  Storm suppression is used to control broadcast packets, unknown multicast
                  packets, and unknown unicast packets, thereby preventing broadcast storms that
                  they may cause.
                  Storm suppression includes traffic suppression and storm control.
                  ●      Traffic suppression limits the rate of broadcast packets, unknown multicast
                         packets, or unknown unicast packets by setting a threshold. When the traffic
                         exceeds this threshold, the system discards excess traffic, allowing only
                         packets within the threshold to pass through. In this way, the volume of traffic
                         is kept within an appropriate range. Note that traffic suppression can also
                         block outgoing packets on interfaces.
                  ●      Storm control blocks broadcast packets, unknown multicast packets, or
                         unknown unicast packets by discarding packets or shutting down an interface.
                         Storm control can also control the average rate of packets by suppressing
                         them. When traffic exceeds the specified threshold, the system performs a
                         predefined storm control action.
                  Table 4-1 compares traffic suppression and storm control.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            51
Security Configuration
Security Configuration                                               4 Storm Suppression Configuration


                  Table 4-1 Comparison between traffic suppression and storm control
                   Item         Traffic Suppression                  Storm Control

                   Traffic      ● If traffic suppression is          ● If the storm control action is
                   control        configured in the outbound           configured to suppress
                                  direction of an interface, the       packets, when the average
                                  system blocks all traffic of the     rate of packets received on
                                  corresponding packet type.           the interface exceeds the
                                ● In other cases, the system           configured upper threshold,
                                  discards the traffic exceeding       the system discards excess
                                  the threshold and allows the         traffic until the average rate
                                  packets within the threshold         of packets drops below the
                                  to pass through.                     threshold.
                                                                     ● In other cases, when traffic
                                                                       exceeds the specified
                                                                       threshold, the system blocks
                                                                       the traffic received by the
                                                                       interface or shuts down the
                                                                       interface.

                   Traffic      ● Chip-based detection               ● Software-based detection
                   detection    ● If traffic exceeds the             ● If the average packet rate
                                  threshold, traffic suppression       exceeds the threshold within
                                  takes effect immediately.            the detection interval, storm
                                                                       control takes effect.




Purpose
                  When a Layer 2 Ethernet interface on a device receives broadcast, unknown
                  multicast, or unknown unicast packets, the device forwards these packets to other
                  Layer 2 Ethernet interfaces in the same VLAN if the outbound interfaces cannot be
                  determined based on the destination MAC addresses of these packets. As a result,
                  a broadcast storm may be generated, degrading forwarding performance of the
                  device. This problem also occurs on a Virtual eXtensible Local Area Network
                  (VXLAN) network.
                  Traffic suppression and storm control can effectively control the traffic of these
                  types of packets.


4.2 Configuration Precautions for Storm Suppression

4.3 Default Settings for Storm Suppression
                  Table 4-2 describes the default settings for storm suppression.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                                52
Security Configuration
Security Configuration                                               4 Storm Suppression Configuration


                  Table 4-2 Default settings for storm suppression
                   Parameter                                  Default Setting

                   Traffic suppression in the inbound         Enabled
                   direction of an interface

                   Traffic suppression mode on an             Percentage of bandwidth occupied
                   interface

