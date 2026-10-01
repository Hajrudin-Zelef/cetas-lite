---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-50
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [6926, 7101]
sha256: 9f8dad481bebd41bb6682f6d65011d7dbf6f9f089189e8d68f8bb52ca627638b
---

                    # On DeviceB, create ACL rules to match packets from enterprise users, and create
                    traffic classifiers c1 and c2 to classify service flows from different enterprise users
                    based on their IP addresses.
                    [DeviceB] acl 2001
                    [DeviceB-acl4-basic-2001] rule permit source 192.168.1.0 0.0.0.255
                    [DeviceB-acl4-basic-2001] quit
                    [DeviceB] acl 2002
                    [DeviceB-acl4-basic-2002] rule permit source 192.168.2.0 0.0.0.255
                    [DeviceB-acl4-basic-2002] quit
                    [DeviceB] traffic classifier c1
                    [DeviceB-classifier-c1] if-match acl 2001
                    [DeviceB-classifier-c1] quit
                    [DeviceB] traffic classifier c2
                    [DeviceB-classifier-c2] if-match acl 2002
                    [DeviceB-classifier-c2] quit

         Step 3 Configure traffic behaviors and define traffic policing.

                    # On DeviceB, create traffic behaviors b1 and b2 to perform traffic policing for
                    packets from different enterprise users.
                    [DeviceB] traffic behavior b1
                    [DeviceB-behavior-b1] car cir 64
                    [DeviceB-behavior-b1] quit
                    [DeviceB] traffic behavior b2
                    [DeviceB-behavior-b2] car cir 128
                    [DeviceB-behavior-b2] quit

         Step 4 Configure a traffic policy and apply it to the inbound direction of an inbound
                interface.

                    # On DeviceB, create traffic policy p1, bind traffic classifiers to traffic behaviors in
                    the traffic policy, and apply the traffic policy to the inbound direction of 10GE
                    1/0/1 to perform traffic policing for packets from two different network segments.
                    [DeviceB] traffic policy p1
                    [DeviceB-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceB-trafficpolicy-p1] classifier c2 behavior b2
                    [DeviceB-trafficpolicy-p1] quit
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] traffic-policy p1 inbound
                    [DeviceB-10GE1/0/1] quit

                    ----End



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       124
QoS Configuration                                                  9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                  based Rate Limiting Configuration


Verifying the Configuration
                    After the preceding configurations are complete, check the traffic policing
                    configuration on DeviceB.
                    # Check the traffic classifier configuration.
                    [DeviceB] display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match acl 2001

                      Classifier: c2
                       Type: OR
                       Rule(s):
                         if-match acl 2002

                    Total classifier number is 2

                    # Check the traffic policy configuration.
                    [DeviceB] display traffic policy p1
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Committed Access Rate:
                           CIR 64 (Kbps), PIR 64 (Kbps), CBS 10000 (Bytes), PBS 10000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard

                        Classifier: c2
                         Type: OR
                        Behavior: b2
                         Committed Access Rate:
                           CIR 128 (Kbps), PIR 128 (Kbps), CBS 10000 (Bytes), PBS 10000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard


Configuration Scripts
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    vlan batch 10 20
                    #
                    acl 2001
                     rule permit source 192.168.1.0 0.0.0.255
                    #
                    acl 2002
                     rule permit source 192.168.2.0 0.0.0.255
                    #
                    traffic classifier c1 type or
                     if-match acl 2001
                    #
                    traffic classifier c2 type or
                     if-match acl 2002
                    #
                    traffic behavior b1
                     car cir 64


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                     125
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration

                    #
                    traffic behavior b2
                     car cir 128
                    #
                    traffic policy p1
                     classifier c1 behavior b1
                     classifier c2 behavior b2
                    #
                    interface Vlanif10
                     ip address 192.168.1.1 255.255.255.0
                    #
                    interface Vlanif20
                     ip address 192.168.2.1 255.255.255.0
                    #
                    interface 10GE1/0/1
                     port link-type trunk
                     port trunk allow-pass vlan 10 20
                     traffic-policy p1 inbound
                    #
                    return



9.6 Configuring Traffic Shaping
Context
                    Traffic accessing a network is discarded due to the bandwidth limit on the
                    network. In contrast with traffic policing, which discards packets exceeding the
                    rate limit, traffic shaping buffers excess packets and sends them out at an even
                    rate.

9.6.1 Configuring Traffic Shaping for a Queue

Context
                    Packets received on an interface enter queues based on priority mapping. Then
                    the device uses different traffic shaping parameter settings for queues of different
                    priorities, providing differentiated services.

                    Before configuring traffic shaping for queues on an interface, configure priority
                    mapping to map packet priorities to per hop behaviors (PHBs) so that packets of
                    different services enter different queues.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface { interface-type interface-number | interface-name }

         Step 3 Configure the traffic shaping rate for queues on an interface.
                    qos queue queue-index shaping { percent cir cir-percent-value [ pir pir-percent-value ] | cir cir-value
                    [ kbps | mbps | gbps ] [ cbs cbs-value [ bytes | kbytes | mbytes ] | pir pir-value [ kbps | mbps | gbps ]
                    [ cbs cbs-value [ bytes | kbytes | mbytes ] pbs pbs-value [ bytes | kbytes | mbytes ] ] ] }

