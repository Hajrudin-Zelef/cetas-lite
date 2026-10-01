---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-26
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [3356, 3538]
sha256: 9b0252fd6e63f815c62b66e5091a0e6ba58fbcd59ab1b5d3598e53cd20583a05
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration
QoS Configuration                                                                                      7 Redirection Configuration

                    <DeviceA> display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match any

                    Total classifier number is 1

                    # Check the traffic behavior configuration.
                    <DeviceA> display traffic behavior
                     Traffic Behavior Information:
                      Behavior: b1
                        Redirect:
                          Redirect interface 10GE1/0/3

                    Total behavior number is 1

                    # Check the traffic policy configuration.
                    <DeviceA> display traffic policy
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Redirect:
                           Redirect interface 10GE1/0/3

                    Total policy number is 1

                    # Check the traffic policy application records.
                    <DeviceA> display traffic-policy applied-record
                    Total records : 1
                    --------------------------------------------------------------------------------
                    Policy Type/Name                     Apply Parameter              Slot State
                    --------------------------------------------------------------------------------
                    p1                             10GE1/0/1(IN)              1    success
                    --------------------------------------------------------------------------------


Configuration Scripts
                    ●     DeviceA
                          #
                          sysname DeviceA
                          #
                          vlan batch 200 300
                          #
                          traffic classifier c1 type or
                           if-match any
                          #
                          traffic behavior b1
                           redirect interface 10GE 1/0/3
                          #
                          traffic policy p1
                           classifier c1 behavior b1 precedence 5
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 200 300
                           traffic-policy p1 inbound
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 200 300
                          #
                          interface 10GE1/0/3


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                                60
QoS Configuration
QoS Configuration                                                                     7 Redirection Configuration

                         port link-type trunk
                         port trunk allow-pass vlan 200 300
                         port-isolate enable group 1
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 200 300
                         port-isolate enable group 1
                         mac-address learning disable
                        #
                        return

                    ●   DeviceB
                        #
                        sysname DeviceB
                        #
                        vlan batch 200 300
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200 300
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type access
                         port default vlan 300
                        #
                        return



7.5 Example for Configuring Redirection to a Next-Hop
Address
Networking Requirements
                    In Figure 7-2, DeviceA functioning as a Layer 3 forwarding device is routable to
                    NetworkA and is connected to the Internet through two links. One uplink is a
                    high-speed link with the gateway at 10.1.20.1/24, and the other is a low-speed
                    link with the gateway at 10.1.30.1/24. The user requires that DeviceA forward
                    packets from network segments 192.168.100.0/24 and 192.168.101.0/24 to the
                    Internet through the high-speed link and low-speed link, respectively. This
                    requirement can be met by configuring redirection to a next-hop address, which is
                    sometimes also called policy-based routing (PBR).

                    Figure 7-2 Network diagram of redirection to a next-hop address
                         NOTE

                        In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                        and 10GE 1/0/3, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        61
QoS Configuration
QoS Configuration                                                              7 Redirection Configuration




Procedure
         Step 1 Create VLANs and configure interfaces.
                    # Create VLAN 10, VLAN 20, and VLAN 30 on DeviceA.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 10 20 30

                    # Configure 10GE 1/0/1, 10GE 1/0/2, and 10GE 1/0/3 on DeviceA as trunk
                    interfaces and add them to corresponding VLANs.
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] portswitch
                    [DeviceA-10GE1/0/1] port link-type trunk
                    [DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface 10ge 1/0/2
                    [DeviceA-10GE1/0/2] portswitch
                    [DeviceA-10GE1/0/2] port link-type trunk
                    [DeviceA-10GE1/0/2] port trunk allow-pass vlan 20
                    [DeviceA-10GE1/0/2] quit
                    [DeviceA] interface 10ge 1/0/3
                    [DeviceA-10GE1/0/3] portswitch
                    [DeviceA-10GE1/0/3] port link-type trunk
                    [DeviceA-10GE1/0/3] port trunk allow-pass vlan 30
                    [DeviceA-10GE1/0/3] quit

                    # Create VLANIF 10, VLANIF 20, and VLANIF 30 and configure IP addresses for
                    them.
                    [DeviceA] interface vlanif 10
                    [DeviceA-Vlanif10] ip address 10.1.10.2 24
                    [DeviceA-Vlanif10] quit
                    [DeviceA] interface vlanif 20
                    [DeviceA-Vlanif20] ip address 10.1.20.2 24
                    [DeviceA-Vlanif20] quit
                    [DeviceA] interface vlanif 30
                    [DeviceA-Vlanif30] ip address 10.1.30.2 24
                    [DeviceA-Vlanif30] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           62
QoS Configuration
QoS Configuration                                                                         7 Redirection Configuration


