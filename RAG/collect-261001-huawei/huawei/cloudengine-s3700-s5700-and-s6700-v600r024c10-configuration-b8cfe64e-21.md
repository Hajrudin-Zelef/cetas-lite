---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-21
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [2573, 2731]
sha256: 62d37a7741cfd7814372049c160f84b99e336415278255a5edb05b6f03716622
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                        In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                        and 10GE 1/0/3, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        45
QoS Configuration
QoS Configuration                                                              6 Re-marking Configuration




Procedure
         Step 1 Create VLANs and configure interfaces so that DeviceB can communicate with
                Host1, Host2, and DeviceA.
                    # Create VLAN 10, VLAN 20, and VLAN 30 on DeviceB.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10 20 30

                    # Configure 10GE 1/0/1 as a trunk interface and add it to VLAN 30. Configure
                    10GE 1/0/2 and 10GE 1/0/3 as access interfaces and add them to VLAN 10 and
                    VLAN 20, respectively.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 30
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type access
                    [DeviceB-10GE1/0/2] port default vlan 10
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] portswitch
                    [DeviceB-10GE1/0/3] port link-type access
                    [DeviceB-10GE1/0/3] port default vlan 20
                    [DeviceB-10GE1/0/3] quit

                    # Create VLANIF 10, VLANIF 20, and VLANIF 30, and configure IP addresses for
                    them.
                    [DeviceB] interface vlanif 10
                    [DeviceB-Vlanif10] ip address 192.168.10.1 24
                    [DeviceB-Vlanif10] quit
                    [DeviceB] interface vlanif 20
                    [DeviceB-Vlanif20] ip address 192.168.20.1 24
                    [DeviceB-Vlanif20] quit
                    [DeviceB] interface vlanif 30
                    [DeviceB-Vlanif30] ip address 192.168.100.1 24
                    [DeviceB-Vlanif30] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          46
QoS Configuration
QoS Configuration                                                               6 Re-marking Configuration


         Step 2 Configure traffic classifiers.

                    # On DeviceB, create and configure traffic classifiers c1 and c2 to classify packets
                    based on VLAN IDs.
                    [DeviceB] traffic classifier c1
                    [DeviceB-classifier-c1] if-match vlan 10
                    [DeviceB-classifier-c1] quit
                    [DeviceB] traffic classifier c2
                    [DeviceB-classifier-c2] if-match vlan 20
                    [DeviceB-classifier-c2] quit

         Step 3 Configure traffic behaviors.

                    # Create and configure traffic behaviors b1 and b2 on DeviceB to re-mark the
                    802.1p value of VLAN packets sent from Host1 into 4 and the 802.1p value of
                    VLAN packets sent from Host2 into 2. In this manner, the priority of packets from
                    Host1 is higher than that of packets from Host2.
                    [DeviceB] traffic behavior b1
                    [DeviceB-behavior-b1] remark 8021p 4
                    [DeviceB-behavior-b1] quit
                    [DeviceB] traffic behavior b2
                    [DeviceB-behavior-b2] remark 8021p 2
                    [DeviceB-behavior-b2] quit

         Step 4 Configure traffic policies and apply them to interfaces.

                    # Create traffic policies p1 and p2 on DeviceB, bind the traffic classifiers and
                    traffic behaviors to the traffic policies, and apply the traffic policy p1 to 10GE
                    1/0/2 in the inbound direction and the traffic policy p2 to 10GE 1/0/3 in the
                    inbound direction to re-mark packet priorities.
                    [DeviceB] traffic policy p1
                    [DeviceB-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceB-trafficpolicy-p1] quit
                    [DeviceB] traffic policy p2
                    [DeviceB-trafficpolicy-p2] classifier c2 behavior b2
                    [DeviceB-trafficpolicy-p2] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] traffic-policy p1 inbound
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] traffic-policy p2 inbound
                    [DeviceB-10GE1/0/3] quit

                    ----End

Verifying the Configuration
                    # Check the traffic classifier configuration.
                    <DeviceB> display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match vlan 10

                      Classifier: c2
                       Type: OR
                       Rule(s):
                         if-match vlan 20

                    Total classifier number is 2


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                            47
QoS Configuration
QoS Configuration                                                                                      6 Re-marking Configuration


                    # Check the traffic policy configuration.
                    <DeviceB> display traffic policy
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Remark:
                           Remark 8021p 4

                       Policy: p2
                        Classifier: c2
                         Type: OR
                        Behavior: b2
                         Remark:
                           Remark 8021p 2

                    Total policy number is 2

                    # Check the traffic policy application records.
                    <DeviceB> display traffic-policy applied-record
                    Total records : 2
                    --------------------------------------------------------------------------------
                    Policy Type/Name                     Apply Parameter               Slot State
                    --------------------------------------------------------------------------------
                    p1                             10GE1/0/2(IN)                  1 success
                    --------------------------------------------------------------------------------
                    p2                             10GE1/0/3(IN)                  1 success
                    --------------------------------------------------------------------------------


