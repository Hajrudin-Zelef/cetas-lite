---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-140
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [15509, 15680]
sha256: b6cb4f150671a3d2e660dd0ecc15979701445156dee91f5bf3088c6cf0065471
---

                    # Configure 10GE1/0/1 as a trunk interface and add it to VLAN 30. Configure
                    10GE1/0/2 and 10GE1/0/3 as access interfaces and add them to VLAN 10 and
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

         Step 2 Enable application identification on inbound interfaces.
                    # Enable application identification on access-side interfaces of DeviceB.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] sa enable


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                             285
QoS Configuration
QoS Configuration                                                            13 Experience Assurance Configuration

                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] sa enable
                    [DeviceB-10GE1/0/3] quit

         Step 3 Configure a traffic policy to match a specific application, and configure a rule of
                re-marking the packet priority in the traffic behavior.
                    # Create and configure a traffic classifier c1 on DeviceB to match the application
                    type TencentMeeting.
                    [DeviceB] traffic classifier c1
                    [DeviceB-classifier-c1] if-match application TencentMeeting
                    [DeviceB-classifier-c1] quit

                    # Create and configure a traffic behavior b1 on DeviceB to re-mark 802.1p values
                    of packets to 5.
                    [DeviceB] traffic behavior b1
                    [DeviceB-behavior-b1] remark 8021p 5
                    [DeviceB-behavior-b1] quit

                    # Create and configure a traffic policy p1 on DeviceB, and bind the traffic classifier
                    to the traffic behavior in the traffic policy.
                    [DeviceB] traffic policy p1
                    [DeviceB-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceB-trafficpolicy-p1] quit

         Step 4 Apply the traffic policy to interfaces to ensure that the corresponding application
                traffic is preferentially forwarded.
                    # Apply the traffic policy p1 to inbound interfaces of DeviceB to ensure that the
                    corresponding application traffic is preferentially forwarded.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] traffic-policy p1 inbound
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] traffic-policy p1 inbound
                    [DeviceB-10GE1/0/3] quit

                    ----End

Verifying the Configuration
                    After the preceding configurations are complete, check the experience assurance
                    configuration on DeviceB.
                    # Check the traffic classifier configuration.
                    <DeviceB> display traffic classifier
                     Traffic Classifier Information:
                       Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match application TencentMeeting
                    Total classifier number is 1

                    # Check the traffic policy configuration.
                    <DeviceB> display traffic policy
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                 286
QoS Configuration
QoS Configuration                                                                        13 Experience Assurance Configuration

                        Behavior: b1
                         Remark:
                          Remark 8021p 5

                    Total policy number is 1

                    # Check the traffic policy application records.
                    <DeviceB> display traffic-policy applied-record
                    Total records : 2
                    --------------------------------------------------------------------------------
                    Policy Type/Name                     Apply Parameter              Slot State
                    --------------------------------------------------------------------------------
                    p1                             10GE1/0/2(IN)                1 success

                    p1                             10GE1/0/3(IN)                1 success
                    --------------------------------------------------------------------------------


Configuration Script
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    vlan batch 10 20 30
                    #
                    interface 10GE1/0/1
                     port link-type trunk
                     port trunk allow-pass vlan 30
                    #
                    interface 10GE1/0/2
                     port link-type access
                     port default vlan 10
                     sa enable
                     traffic-policy p1 inbound
                    #
                    interface 10GE1/0/3
                     port link-type access
                     port default vlan 20
                     sa enable
                     traffic-policy p1 inbound
                    #
                    interface vlanif 10
                     ip address 192.168.10.1 255.255.255.0
                    #
                    interface vlanif 20
                     ip address 192.168.20.1 255.255.255.0
                    #
                    interface vlanif 30
                     ip address 192.168.100.1 255.255.255.0
                    #
                    traffic classifier c1 type or
                     if-match application TencentMeeting
                    #
                    traffic behavior b1
                     remark 8021p 5
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                    #
                    return


13.5.6 Example for Configuring Experience Assurance for
Traffic Statistics Collection

Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                           287

