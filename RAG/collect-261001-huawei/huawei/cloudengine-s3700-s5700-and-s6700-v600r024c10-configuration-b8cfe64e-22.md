---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-22
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [2732, 2918]
sha256: b8234d3fb45048ccd62775c9c2d90d2b52f04a745ff7ab7e511782d8d6ec6fb5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Configuration Scripts
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    vlan batch 10 20 30
                    #
                    traffic classifier c1 type or
                     if-match vlan 10
                    #
                    traffic classifier c2 type or
                     if-match vlan 20
                    #
                    traffic behavior b1
                     remark 8021p 4
                    #
                    traffic behavior b2
                     remark 8021p 2
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                    #
                    traffic policy p2
                     classifier c2 behavior b2 precedence 5
                    #
                    interface Vlanif10
                     ip address 192.168.10.1 255.255.255.0
                    #
                    interface Vlanif20
                     ip address 192.168.20.1 255.255.255.0
                    #
                    interface Vlanif30
                     ip address 192.168.100.1 255.255.255.0
                    #
                    interface 10GE1/0/1


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                               48
QoS Configuration
QoS Configuration                                                                      6 Re-marking Configuration

                     port link-type trunk
                     port trunk allow-pass vlan 30
                    #
                    interface 10GE1/0/2
                     port link-type access
                     port default vlan 10
                     traffic-policy p1 inbound
                    #
                    interface 10GE1/0/3
                     port link-type access
                     port default vlan 20
                     traffic-policy p2 inbound
                    #
                    return



6.6 Example for Configuring Re-marking to Distinguish
Services

Networking Requirements
                    In Figure 6-4, voice, video, and data terminals on the LAN of an enterprise are
                    connected to interface 1 and interface 2 of DeviceC through DeviceA and DeviceB,
                    and are connected to the WAN through DeviceC and DeviceD.

                    Packets of different services are identified by 802.1p priorities on the LAN side.
                    When packets reach the WAN side through interface 3, it is required that
                    differentiated services be provided based on DSCP priorities of packets.

                    Figure 6-4 Network diagram for configuring re-marking
                          NOTE

                         In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                         and 10GE 1/0/3, respectively.




Procedure
         Step 1 Set the device name to DeviceC.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       49
QoS Configuration
QoS Configuration                                                                  6 Re-marking Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceC

         Step 2 Configure traffic classifiers.
                    # On DeviceC, create and configure traffic classifiers c1, c2, and c3 to classify
                    packets based on 802.1p priorities.
                    [DeviceC] traffic classifier c1
                    [DeviceC-classifier-c1] if-match 8021p 2
                    [DeviceC-classifier-c1] quit
                    [DeviceC] traffic classifier c2
                    [DeviceC-classifier-c2] if-match 8021p 5
                    [DeviceC-classifier-c2] quit
                    [DeviceC] traffic classifier c3
                    [DeviceC-classifier-c3] if-match 8021p 6
                    [DeviceC-classifier-c3] quit

         Step 3 Configure traffic behaviors.
                    # On DeviceC, create and configure traffic behaviors b1, b2, and b3 to re-mark
                    priorities of user packets.
                    [DeviceC] traffic behavior b1
                    [DeviceC-behavior-b1] remark dscp 15
                    [DeviceC-behavior-b1] quit
                    [DeviceC] traffic behavior b2
                    [DeviceC-behavior-b2] remark dscp 40
                    [DeviceC-behavior-b2] quit
                    [DeviceC] traffic behavior b3
                    [DeviceC-behavior-b3] remark dscp 50
                    [DeviceC-behavior-b3] quit

         Step 4 Configure a traffic policy and apply it to interfaces.
                    # On DeviceC, create a traffic policy p1, bind traffic classifiers to traffic behaviors
                    in the traffic policy, and apply the traffic policy to the inbound direction of
                    interface 1 and interface 2 to re-mark packets.
                    [DeviceC] traffic policy p1
                    [DeviceC-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceC-trafficpolicy-p1] classifier c2 behavior b2
                    [DeviceC-trafficpolicy-p1] classifier c3 behavior b3
                    [DeviceC-trafficpolicy-p1] quit
                    [DeviceC] interface 10ge 1/0/1
                    [DeviceC-10GE1/0/1] traffic-policy p1 inbound
                    [DeviceC-10GE1/0/1] quit
                    [DeviceC] interface 10ge 1/0/2
                    [DeviceC-10GE1/0/2] traffic-policy p1 inbound
                    [DeviceC-10GE1/0/2] quit

                    ----End

Verifying the Configuration
                    # Check the traffic classifier configuration.
                    <DeviceC> display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match 8021p 2

                      Classifier: c2
                       Type: OR
                       Rule(s):


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                          50
QoS Configuration
QoS Configuration                                                                                      6 Re-marking Configuration

                          if-match 8021p 5

                       Classifier: c3
                        Type: OR
                        Rule(s):
                          if-match 8021p 6

                    Total classifier number is 3

                    # Check the traffic policy configuration.
                    <DeviceC> display traffic policy
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Remark:
                           Remark dscp 15

                        Classifier: c2
                         Type: OR
                        Behavior: b2
                         Remark:
                           Remark dscp cs5

                        Classifier: c3
                         Type: OR
                        Behavior: b3
                         Remark:
                           Remark dscp 50

                    Total policy number is 3

