---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-16
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [1779, 1959]
sha256: c3210db6e275e12fd4500529817eed0bbcf8fbefc2c98254f79af43f467f5783
---

                    # On DeviceA, create a traffic behavior b2 and configure the deny action in the
                    traffic behavior.
                    [DeviceA] traffic behavior b2
                    [DeviceA-behavior-b2] deny
                    [DeviceA-behavior-b2] quit

         Step 5 Configure a traffic policy and apply it to the inbound direction of an interface.

                    # On DeviceA, create a traffic policy p1 and bind traffic classifiers to traffic
                    behaviors in the traffic policy.
                    [DeviceA] traffic policy p1
                    [DeviceA-trafficpolicy-p1] classifier c1 behavior b1 precedence 5
                    [DeviceA-trafficpolicy-p1] classifier c2 behavior b2 precedence 10
                    [DeviceA-trafficpolicy-p1] classifier c3 behavior b2 precedence 15
                    [DeviceA-trafficpolicy-p1] classifier c4 behavior b2 precedence 20
                    [DeviceA-trafficpolicy-p1] quit

                    # Apply the traffic policy p1 to the inbound direction of an interface.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         30
QoS Configuration
QoS Configuration                                                             4 Packet Filtering Configuration

                    [DeviceA] interface vlanif 10
                    [DeviceA-Vlanif10] traffic-policy p1 inbound
                    [DeviceA-Vlanif10] quit

                    ----End

Verifying the Configuration
                    # Check the ACL rule configuration.
                    <DeviceA> display acl 3001
                    Advanced ACL 3001, 1 rule
                    ACL's step is 5
                     rule 1 permit ip destination 10.1.1.0 0.0.0.255
                     (0 times matched)

                    # Check the traffic classifier configuration.
                    <DeviceA> display traffic classifier c1
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: AND
                        Rule(s):
                          if-match acl 3001
                    <DeviceA> display traffic classifier c2
                     Traffic Classifier Information:
                      Classifier: c2
                        Type: AND
                        Rule(s):
                          if-match source-mac 00e0-fc0d-0001 ffff-ffff-ffff

                    # Check the traffic policy configuration.
                    <DeviceA> display traffic policy p1
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: AND
                        Behavior: b1

                        Classifier: c2
                         Type: AND
                        Behavior: b2
                         Deny

                        Classifier: c3
                         Type: AND
                        Behavior: b2
                         Deny

                        Classifier: c4
                         Type: AND
                        Behavior: b2
                         Deny


Configuration Scripts
                    DeviceA
                    #
                    sysname DeviceA
                    #
                    vlan batch 10
                    #
                    interface 10GE1/0/1
                     port link-type trunk
                     port trunk allow-pass vlan 10


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                             31
QoS Configuration
QoS Configuration                                                           4 Packet Filtering Configuration

                    #
                    acl 3001
                     rule 1 permit ip destination 10.1.1.0 0.0.0.255
                    #
                    traffic classifier c1 type and
                     if-match acl 3001
                    #
                    traffic classifier c2 type and
                     if-match source-mac 00e0-fc0d-0001
                    #
                    traffic classifier c3 type and
                     if-match source-mac 00e0-fc0d-0002
                    #
                    traffic classifier c4 type and
                     if-match source-mac 00e0-fc0d-0003
                    #
                    traffic behavior b1
                     permit
                    #
                    traffic behavior b2
                     deny
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                     classifier c2 behavior b2 precedence 10
                     classifier c3 behavior b2 precedence 15
                     classifier c4 behavior b2 precedence 20
                    #
                    interface Vlanif10
                     ip address 10.1.1.1 255.255.255.0
                     traffic-policy p1 inbound
                    #
                    return




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                            32
QoS Configuration
QoS Configuration                                               5 Traffic Statistics Collection Configuration




                                   5         Traffic Statistics Collection
                                                           Configuration

                    5.1 Overview of Traffic Statistics Collection
                    5.2 Configuration Precautions for Traffic Statistics Collection
                    5.3 Configuring MQC-based Traffic Statistics Collection
                    5.4 Example for Configuring MQC-based Traffic Statistics Collection


5.1 Overview of Traffic Statistics Collection
                    Traffic statistics collection allows the device to collect statistics on packets that
                    match traffic classification rules. With this function enabled, the device can collect
                    statistics on forwarded and discarded packets matching traffic policies, helping
                    administrators to locate faults and determine whether traffic policies are correctly
                    applied. You can run the display traffic-policy statistics command to view
                    statistics on forwarded and discarded packets matching traffic policies only when
                    MQC is used to implement traffic statistics collection.
                    Table 5-1 describes the differences between traffic statistics collection and
                    interface statistics collection.

                    Table 5-1 Differences between traffic statistics collection and interface statistics
                    collection
                     Statistics            Command                  Scope                  Remarks
                     Collection Mode

                     Traffic statistics    display traffic-         Packets matching       Packets sent to
                     collection            policy statistics        traffic                the CPU are
                                                                    classification rules   excluded.
                                                                    after a traffic
                                                                    policy is applied




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  33
QoS Configuration
QoS Configuration                                                          5 Traffic Statistics Collection Configuration


                     Statistics                 Command                    Scope                  Remarks
                     Collection Mode

