---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-44
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [6011, 6154]
sha256: abe6f92cceea5fe81bffbbcc7052014b64500ed52d8be0c82c3a14ad0da543a2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    1.   If there are sufficient tokens in the bucket, the system sends the packets and
                         decreases the number of tokens accordingly.
                    2.   If tokens in the bucket are insufficient for packet forwarding, the system
                         places the packets into the buffer queue. If the buffer queue is full, the
                         system discards the packets.
                    3.   When there are packets in the buffer queue, the system compares the number
                         of packets with the number of tokens in the token bucket. If there are
                         sufficient tokens, the system forwards packets until all the packets in the
                         buffer queue are sent.


9.3 Configuration Precautions for Traffic Policing,
Traffic Shaping, and Interface-based Rate Limiting

9.4 Default Settings for Traffic Policing, Traffic Shaping,
and Interface-based Rate Limiting
                    Table 9-6 describes the default settings for traffic policing, traffic shaping, and
                    interface-based rate limiting.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                109
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration


                    Table 9-6 Default settings for traffic policing, traffic shaping, and interface-based
                    rate limiting

                     Parameter                                  Default Setting

                     MQC-based traffic policing (level-1        Disabled
                     CAR)

                     Logical operator between rules in a        OR
                     traffic classifier

                     Queue-based traffic shaping                Disabled

                     Queue-based traffic shaping rate           Maximum interface bandwidth

                     Interface-based rate limiting through      Disabled
                     traffic policing

                     Interface-based rate limiting through      Disabled
                     traffic shaping

                     Rate of interface-based rate limiting      Maximum interface bandwidth
                     through traffic shaping

                     Rate limit on the management               3000 pps
                     interface




9.5 Configuring Traffic Policing

9.5.1 Configuring MQC-based Traffic Policing (Level-1 CAR)
Context
                    There are various services on a network. When a large amount of service traffic
                    enters the network side, congestion may occur due to insufficient bandwidth. To
                    control a specific type of service traffic in the inbound or outbound direction on an
                    interface, MQC-based traffic policing (level-1 CAR) can be configured. MQC-based
                    traffic policing uses traffic classifiers to implement differentiated services. When
                    the receive or transmit rate of packets matching traffic classification rules exceeds
                    the specified limit, the device discards the packets.
                    When both QoS CAR and MQC are used on an interface for traffic statistics
                    collection, statistics about packets to which a QoS CAR profile is applied contain
                    only the packets that do not match the traffic policy.
                    The device performs the CAR action on interfaces of the same chip centrally, and
                    on interfaces of different chips separately. If a traffic policy containing traffic
                    policing is applied to an Eth-Trunk, a VLAN, or the system, and interfaces of the
                    object to which the traffic policy is applied belong to N chips, the actual rate limit
                    is N times as large as the CAR value.
                    For example, assume that interface 1 and interface 2 belong to chip 0 and
                    interface 3 and interface 4 belong to chip 1. If interfaces 1 to 4 join VLAN 10 and

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                110
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration


                    a traffic policy containing the rate limit of 1 Mbit/s is applied to VLAN 10, the rate
                    limit of interface 1 and interface 2 is 1 Mbit/s and the rate limit of interface 3 and
                    interface 4 is 1 Mbit/s. That is, the total rate limit is 2 Mbit/s after the traffic policy
                    is applied to VLAN 10.

Procedure
         Step 1 Configure a traffic classifier.

                    For details about how to configure a traffic classifier, see 3.4 Configuring a Traffic
                    Classifier in "MQC Configuration".

         Step 2 Configure a traffic behavior.
                    1.   (Optional) Disable the device from counting the inter-frame gaps and
                         preambles when the device calculates the traffic policing rate.
                         qos car ifg disable

                    2.   Create a traffic behavior and enter the traffic behavior view, or enter the view
                         of an existing traffic behavior.
                         traffic behavior behavior-name

                    3.   Configure a CAR action.
                         car cir cir-value [ kbps | mbps | gbps ] [ pir pir-value [ kbps | mbps | gbps ] ] [ cbs cbs-value [ bytes
                         | kbytes | mbytes ] pbs pbs-value [ bytes | kbytes | mbytes ] ] [ share ] [ mode { color-blind | color-
                         aware } ] [ green pass [ service-class class color color ] | yellow { discard | pass [ service-class
                         class color color ] } | red { discard | pass [ service-class class color color ] } ] *
                    4.   Exit the traffic behavior view.
                         quit

         Step 3 Configure a traffic policy.
                    1.   Create a traffic policy and enter the traffic policy view, or enter the view of an
                         existing traffic policy.
                         traffic policy policy-name

                    2.   Bind a traffic behavior and a traffic classifier to the traffic policy.
                         classifier classifier-name behavior behavior-name [ precedence precedence-value ]

                    3.   Exit the traffic policy view.
                         quit

         Step 4 Apply a traffic policy.
                    ●    Apply a traffic policy to the system.
                         a.     Apply a traffic policy to the system.
                                traffic-policy policy-name global [ slot slot-id ] { inbound | outbound }

                    ●    Apply a traffic policy to an interface.
                         a.     Enter the interface view.
                                interface interface-type interface-number

                         b.     Apply a traffic policy to the interface.
                                traffic-policy policy-name { inbound | outbound }

                         c.     Exit the interface view.
                                quit

                    ●    Apply a traffic policy to a VLAN.
                         a.     Create a VLAN and enter the VLAN view.
                                vlan vlan-id


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                             111

