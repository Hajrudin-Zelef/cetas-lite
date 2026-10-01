---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-48
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [6647, 6791]
sha256: 893982c5b2ac572f47bbe9ea6698d30911a1a6f1e1063f4c882110ec1f066bde
---

                    # On DeviceB, create VLANs 10, 20, and 30, and add interfaces to the
                    corresponding VLANs.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10 20 30
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 10 20 30
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type trunk
                    [DeviceB-10GE1/0/2] port trunk allow-pass vlan 10 20 30
                    [DeviceB-10GE1/0/2] quit

         Step 2 Configure a CAR profile.
                    [DeviceB] qos car car1 cir 12000

         Step 3 Configure traffic classifiers.
                    # On DeviceB, create traffic classifiers c1, c2, and c3 to classify different service
                    flows based on their VLAN IDs.
                    [DeviceB] traffic classifier c1
                    [DeviceB-classifier-c1] if-match vlan 10
                    [DeviceB-classifier-c1] quit
                    [DeviceB] traffic classifier c2
                    [DeviceB-classifier-c2] if-match vlan 20
                    [DeviceB-classifier-c2] quit
                    [DeviceB] traffic classifier c3
                    [DeviceB-classifier-c3] if-match vlan 30
                    [DeviceB-classifier-c3] quit

         Step 4 Configure traffic behaviors and define traffic policing.
                    # On DeviceB, create traffic behaviors b1, b2, and b3 to perform traffic policing
                    and priority re-marking for different service flows.
                    [DeviceB] traffic behavior b1
                    [DeviceB-behavior-b1] car cir 2000 pir 10000 green pass
                    [DeviceB-behavior-b1] car car1 share
                    [DeviceB-behavior-b1] remark dscp 46
                    [DeviceB-behavior-b1] statistics enable
                    [DeviceB-behavior-b1] quit
                    [DeviceB] traffic behavior b2
                    [DeviceB-behavior-b2] car cir 4000 pir 10000 green pass
                    [DeviceB-behavior-b2] remark dscp 14
                    [DeviceB-behavior-b2] statistics enable
                    [DeviceB-behavior-b2] quit
                    [DeviceB] traffic behavior b3
                    [DeviceB-behavior-b3] car cir 4000 pir 10000 green pass
                    [DeviceB-behavior-b3] remark dscp 30
                    [DeviceB-behavior-b3] statistics enable
                    [DeviceB-behavior-b3] quit

         Step 5 Configure a traffic policy and apply it to an inbound interface.
                    # On DeviceB, create traffic policy p1, bind traffic classifiers to traffic behaviors in
                    the traffic policy, and apply the traffic policy to the inbound direction of 10GE
                    1/0/1 to perform traffic policing and re-marking on packets.
                    [DeviceB] traffic policy p1
                    [DeviceB-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceB-trafficpolicy-p1] classifier c2 behavior b2
                    [DeviceB-trafficpolicy-p1] classifier c3 behavior b3
                    [DeviceB-trafficpolicy-p1] quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       120
QoS Configuration                                                   9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                   based Rate Limiting Configuration

                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] traffic-policy p1 inbound
                    [DeviceB-10GE1/0/1] quit

                    ----End

Verifying the Configuration
                    # Check the traffic classifier configuration.
                    [DeviceB] display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match vlan 10

                      Classifier: c2
                       Type: OR
                       Rule(s):
                         if-match vlan 20

                      Classifier: c3
                       Type: OR
                       Rule(s):
                         if-match vlan 30

                    Total classifier number is 3

                    # Check the configuration of traffic policy p1.
                    [DeviceB] display traffic policy p1
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Committed Access Rate:
                           CIR 2000 (Kbps), PIR 10000 (Kbps), CBS 16000 (Bytes), PBS 80000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard
                          Remark:
                           Remark dscp ef
                          Statistics: enable

                        Classifier: c2
                         Type: OR
                        Behavior: b2
                         Committed Access Rate:
                           CIR 4000 (Kbps), PIR 10000 (Kbps), CBS 32000 (Bytes), PBS 80000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard
                         Remark:
                           Remark dscp af13
                         Statistics: enable

                        Classifier: c3
                         Type: OR
                        Behavior: b3
                         Committed Access Rate:
                           CIR 4000 (Kbps), PIR 10000 (Kbps), CBS 32000 (Bytes), PBS 80000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      121
QoS Configuration                                                        9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                        based Rate Limiting Configuration

                         Remark:
                          Remark dscp af33
                         Statistics: enable

