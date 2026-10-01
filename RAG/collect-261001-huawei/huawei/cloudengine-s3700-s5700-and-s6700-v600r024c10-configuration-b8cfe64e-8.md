---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-8
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "latency", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [571, 669]
sha256: 6efc98c25269db5298f904e976b81ac37e57b4d63e7b74f8a3d3431e1a141bba
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    In order to ensure services remain stable even in scenarios where network
                    resources are limited, user traffic rates must be limited. This can be achieved by
                    using traffic policing, traffic shaping, and interface-based rate limiting, which
                    define basic bandwidth (rate limiting) of different services passing through
                    network devices and monitor the rate of services entering network devices (speed
                    testing). When the volume of service traffic exceeds the basic bandwidth
                    (speeding), excess traffic is discarded or buffered (punishment). In this way, traffic
                    is limited and resource utilization is improved, resulting in more stable services for
                    users.
                    ●   Traffic policing controls the traffic rate within a bandwidth limit. It does this
                        by discarding excess traffic when the service traffic exceeds the rate limit. This
                        prevents individual services or users from consuming excessive bandwidth
                        resources.
                    ●   Traffic shaping adjusts the rate of outgoing traffic to stabilize the rate at
                        which it is transmitted, thereby avoiding unnecessary packet loss and
                        congestion. Unlike traffic policing, which discards packets exceeding the rate
                        limit, traffic shaping buffers such packets and sends them out at an even rate.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  7
QoS Configuration
QoS Configuration                                                                      2 Overview of QoS


                    ●   Interface-based rate limiting limits the total rate of packets sent or received
                        on an interface and can be implemented through either traffic policing or
                        traffic shaping.
                    Congestion avoidance
                    Congestion avoidance is a congestion control mechanism that monitors network
                    resources such as queues and memory buffers, as well as discarding packets when
                    congestion occurs or worsens.
                    Congestion management
                    Congestion management is a queue-based technology that works by buffering
                    packets in queues upon network congestion occurs. Congestion management
                    schedules the packets based on a scheduling algorithm, ensuring that QoS-
                    demanding services, such as latency-sensitive services, are preferentially scheduled.

QoS Service Process
                    To sum up, traffic classification is the basis of differentiated services. Traffic
                    policing, traffic shaping, interface-based rate limiting, congestion avoidance, and
                    congestion management are techniques for controlling network traffic and
                    resource allocation, and ultimately implementing differentiated services.
                    Figure 2-1 shows the QoS service process on network devices.

                    Figure 2-1 QoS service process




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 8
QoS Configuration
QoS Configuration                                                                  3 MQC Configuration




                                                    3       MQC Configuration


                    3.1 Overview of MQC
                    3.2 Understanding MQC
                    3.3 Configuration Precautions for MQC
                    3.4 Configuring a Traffic Classifier
                    3.5 Configuring a Traffic Behavior
                    3.6 Configuring a Traffic Policy
                    3.7 Applying a Traffic Policy
                    3.8 Verifying the Configuration
                    3.9 Maintaining MQC


3.1 Overview of MQC
                    Modular QoS Command-Line Interface (MQC) allows the device to classify
                    packets based on their characteristics and provide the same QoS level for packets
                    of the same type. This enables the device to provide differentiated services for
                    packets of different types. As more diversified services are deployed on a network,
                    service deployment becomes increasingly complex, if differentiated services need
                    to be provided for traffic of different services or users. Leveraging MQC, you can
                    implement fine-grained classification of network traffic and specify the QoS levels
                    for traffic of different types according to your requirements, enhancing
                    serviceability of your network.

                    To implement MQC, you need to configure traffic classifiers, traffic behaviors, and
                    traffic policies, and apply the traffic policies. MQC is commonly used in the
                    following scenarios:
                    ●   4.3 Configuring MQC-based Packet Filtering
                    ●   Configuring MQC-based Traffic Statistics Collection
                    ●   6.4 Configuring MQC-based Priority Re-marking
                    ●   Configuring MQC-based Redirection


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                9
QoS Configuration
QoS Configuration                                                                       3 MQC Configuration




