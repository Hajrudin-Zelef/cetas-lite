---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-7
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [485, 570]
sha256: 5166ef91dce6ed134d6c6b106c250a2e6d6087a9591cc4b858c5738384a0422e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Packet Classification Methods
                    DiffServ classifies packets on a network into multiple classes and provides
                    differentiated processing for each class. The device supports behavior aggregate
                    (BA) classification and multi-field (MF) classification.
                    MF classification: MQC-based traffic classification
                    MF classification refers to the process of classifying packets in a refined manner
                    based on complex rules, such as 5-tuple (source IP address, source port number,
                    protocol number, destination IP address, and destination port number). Using MF
                    classification, the device classifies packets with the same characteristics into the
                    same type, and assigns them the same QoS level. MF classification can be
                    implemented through traffic classifiers in the Modular QoS Command-Line
                    Interface (MQC), which involves three elements: traffic classifier, traffic behavior,
                    and traffic policy.
                    1.   Traffic classifier: Defines matching rules.
                    2.   Traffic behavior: Specifies actions to be taken for packets. Different QoS
                         functions can be implemented based on traffic behaviors. This document
                         provides guidance for configuring MQC-based traffic statistics collection and
                         packet filtering.
                    3.   Traffic policy: Binds the configured traffic classifier and traffic behavior into a
                         traffic policy, which is then applied in a specified view.
                    BA classification: QoS priority-based classification
                    BA classification refers to the process of classifying packets based on simple rules
                    — a certain priority field in packets, to identify traffic with different priorities.
                    Priorities are described as follows:
                    ●    External priority
                         The external priority is also known as the packet priority or QoS priority. QoS
                         information is recorded by using certain fields in packets, such as the 802.1p
                         value of VLAN packets and the Differentiated Services Code Point (DSCP)
                         value of IP packets. A device can process received packets only based on
                         internal priorities to provide differentiated QoS levels for different services. As
                         such, external priorities are mapped to internal priorities after packets enter
                         the device.
                    ●    Internal priority
                         The internal priority is also known as the class of service (CoS), per-hop
                         behavior (PHB), or local priority. Internal priority values are CS7, CS6, EF, AF4,
                         AF3, AF2, AF1, and BE (in descending order of priority), and correspond to
                         queues 7, 6, 5, 4, 3, 2, 1, and 0, respectively. The internal priority determines
                         the queue into which packets are placed. When QoS services are configured
                         for a queue, the same QoS level is configured for all the packets forwarded
                         through that queue.
                    ●    Drop priority
                         The drop priority is also known as a color. It determines the sequence in
                         which packets are dropped when congestion occurs in a queue, without

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     6
QoS Configuration
QoS Configuration                                                                       2 Overview of QoS


                        affecting the mapping between internal priorities and queues. The drop
                        priorities defined by Institute of Electrical and Electronics Engineers (IEEE) are
                        green, yellow, and red in ascending order. By default, the device first discards
                        packets with a higher drop priority when congestion occurs in a queue.
                        Whether packets are discarded first is determined by QoS parameter settings.
                        For example, in a weighted random early detection (WRED) drop profile, if
                        green packets can use a maximum of 50% of the buffer and red packets can
                        use the entire buffer, the device first discards green packets when congestion
                        occurs in a queue.

                    When processing QoS services, in the inbound direction, the device maps external
                    priorities of packets to internal and drop priorities, while in the outbound
                    direction, it maps internal and drop priorities to external priorities. To centrally
                    schedule packets, you can use the following methods to adjust the mapping
                    between external priorities and internal priorities:

                    ●   Priority mapping: In the DiffServ model, the device defines DiffServ domains
                        to manage and record the mapping between external priorities and internal
                        or drop priorities, which can vary between multiple DiffServ domains. In this
                        way, differentiated QoS services are provided.

                    Relationship between MF classification and BA classification

                    MF classification can also be used to identify packets with specific QoS priorities,
                    and this is implemented through MQC. In summary, MF classification classifies
                    packets based on MQC traffic classifiers; BA classification classifies packets based
                    on the internal and drop priorities that the external priorities of packets are
                    mapped to, without using MQC traffic classifiers.

QoS Service Technologies
                    After packets are classified, differentiated QoS services can be provided for
                    different types of packets. Here, we describe the following QoS service
                    technologies:

                    Traffic policing, traffic shaping, and interface-based rate limiting

