---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-23
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [2919, 3058]
sha256: 893206ceb73b4eaf103769ddbd767178a667b7f7a5908c5555dc6b2bd78ab2b9
---

                    # Check the traffic policy application records.
                    <DeviceC> display traffic-policy applied-record
                    Total records : 2
                    --------------------------------------------------------------------------------
                    Policy Type/Name                  Apply Parameter                 Slot State
                    --------------------------------------------------------------------------------
                    p1                         10GE1/0/1(IN)                   slot_1     success

                                              10GE1/0/2(IN)                   slot_1    success
                    --------------------------------------------------------------------------------


Configuration Scripts
                    DeviceC
                    #
                    sysname DeviceC
                    #
                    traffic classifier c1 type or
                     if-match 8021p 2
                    #
                    traffic classifier c2 type or
                     if-match 8021p 5
                    #
                    traffic classifier c3 type or
                     if-match 8021p 6
                    #
                    traffic behavior b1
                     remark dscp 15
                    #
                    traffic behavior b2
                     remark dscp cs5
                    #
                    traffic behavior b3


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                               51
QoS Configuration
QoS Configuration                                                               6 Re-marking Configuration

                     remark dscp 50
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                     classifier c2 behavior b2 precedence 10
                     classifier c3 behavior b3 precedence 15
                    #
                    interface 10GE1/0/1
                     traffic-policy p1 inbound
                    #
                    interface 10GE1/0/2
                     traffic-policy p1 inbound
                    #
                    return




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          52
QoS Configuration
QoS Configuration                                                               7 Redirection Configuration




                                   7          Redirection Configuration


                    7.1 Overview of Redirection
                    7.2 Configuration Precautions for Redirection
                    7.3 Configuring MQC-based Redirection
                    7.4 Example for Configuring Redirection to an Interface
                    7.5 Example for Configuring Redirection to a Next-Hop Address
                    7.6 Example for Configuring Association Between Redirection to a Next-Hop
                    Address and NQA
                    7.7 Example for Configuring Redirection to Implement Route Selection


7.1 Overview of Redirection
                    MQC-based redirection redirects packets that match traffic classification rules to a
                    specific location for processing.
                    The device supports redirection to the following objects:
                    ●   Redirection to an interface: Packets are redirected to a specified interface for
                        processing or further forwarding.
                    ●   Redirection to a next-hop address: When the received packets need to be
                        processed by a downstream device, packets can be redirected to a next-hop
                        address. This type of redirection is also called policy-based routing (PBR). PBR
                        takes precedence over direct routes, static routes, and routes generated using
                        dynamic routing protocols. Redirection to a next-hop address is also called
                        policy-based routing (PBR).
                        This operation has no detection mechanism. As such, if the corresponding link
                        fails, the next-hop address to which packets are redirected does not change
                        automatically; rather, a network administrator needs to change it manually or
                        the device waits for the ARP entry corresponding to the next-hop address to
                        age out. As a result, services cannot be switched immediately and may be
                        interrupted for a long period of time. You can associate Network Quality
                        Analysis (NQA) with redirection to a next-hop address to detect the link
                        status of the next hop. If the link of the next hop fails, NQA detection will
                        also fail and this type of redirection immediately becomes unavailable,

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               53
QoS Configuration
QoS Configuration                                                                           7 Redirection Configuration


                         reducing the communication interruption time and improving service quality.
                         When associating NQA with redirection to a next-hop address, take into
                         account the following:
                         –    If an NQA test instance does not exist or it is not the Internet Control
                              Message Protocol (ICMP) type, link detection fails and redirection to a
                              next-hop address becomes unavailable.
                         –    If an NQA test instance is successfully associated with redirection to a
                              next-hop address and the link is detected to be normal, packets can be
                              redirected to a next-hop address.
                         –    If an NQA test instance is successfully associated but link detection
                              continuously fails and the number of detection times exceeds the preset
                              value, redirection to a next-hop address automatically becomes
                              unavailable.
                         –    If an NQA test instance is successfully associated with redirection to a
                              next-hop address and the link is detected to be normal, redirection to a
                              next-hop address automatically becomes available.


7.2 Configuration Precautions for Redirection

7.3 Configuring MQC-based Redirection
Context
                    After redirection is configured, the device can redirect packets matching traffic
                    classification rules to a specified interface and a specified next-hop address.

                          NOTE

                         ● For Layer 2 data traffic, you are advised to configure redirection to an interface. For
                           Layer 3 data traffic, you are advised to configure redirection to a next-hop address.
                         ● For details about MQC-related configuration precautions, see "Configuration Precautions
                           for MQC" in MQC Configuration.


