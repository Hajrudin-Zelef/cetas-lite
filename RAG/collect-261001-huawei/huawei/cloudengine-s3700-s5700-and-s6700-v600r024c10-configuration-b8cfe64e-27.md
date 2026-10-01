---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-27
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [3539, 3698]
sha256: 069e325059a39c1f85ce37d9ee91ec90d7d583a5c9aa5f22edd57521b69145d5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

         Step 2 Configure ACL rules.
                    # Create advanced ACLs 3001 and 3002 on DeviceA to allow packets on network
                    segments 192.168.100.0/24 and 192.168.101.0/24 to pass through.
                    [DeviceA] acl 3001
                    [DeviceA-acl4-advance-3001] rule permit ip source 192.168.100.0 0.0.0.255
                    [DeviceA-acl4-advance-3001] quit
                    [DeviceA] acl 3002
                    [DeviceA-acl4-advance-3002] rule permit ip source 192.168.101.0 0.0.0.255
                    [DeviceA-acl4-advance-3002] quit

         Step 3 Configure traffic classifiers.
                    # Create traffic classifiers c1 and c2 on DeviceA, and bind c1 to ACL 3001 and c2
                    to ACL 3002.
                    [DeviceA] traffic classifier c1
                    [DeviceA-classifier-c1] if-match acl 3001
                    [DeviceA-classifier-c1] quit
                    [DeviceA] traffic classifier c2
                    [DeviceA-classifier-c2] if-match acl 3002
                    [DeviceA-classifier-c2] quit

         Step 4 Configure traffic behaviors.
                    # Create traffic behaviors b1 and b2 on DeviceA and configure actions that
                    redirect packets to IP addresses 10.1.20.1 and 10.1.30.1.
                    [DeviceA] traffic behavior b1
                    [DeviceA-behavior-b1] redirect nexthop 10.1.20.1
                    [DeviceA-behavior-b1] quit
                    [DeviceA] traffic behavior b2
                    [DeviceA-behavior-b2] redirect nexthop 10.1.30.1
                    [DeviceA-behavior-b2] quit

         Step 5 Configure a traffic policy and apply it to an interface.
                    # Create a traffic policy p1 on DeviceA, and bind the traffic classifier c1 to the
                    traffic behavior b1 and the traffic classifier c2 to the traffic behavior b2.
                    [DeviceA] traffic policy p1
                    [DeviceA-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceA-trafficpolicy-p1] classifier c2 behavior b2
                    [DeviceA-trafficpolicy-p1] quit

                    # Apply traffic policy p1 to the inbound direction of 10GE 1/0/1.
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] traffic-policy p1 inbound
                    [DeviceA-10GE1/0/1] quit

                    ----End

Verifying the Configuration
                    # Check the traffic classifier configuration.
                    <DeviceA> display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match acl 3001

                      Classifier: c2
                       Type: OR


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                  63
QoS Configuration
QoS Configuration                                                                                      7 Redirection Configuration

                        Rule(s):
                         if-match acl 3002

                    Total classifier number is 2

                    # Check the traffic policy configuration.
                    <DeviceA> display traffic policy
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Redirect:
                           Redirect nexthop
                           10.1.20.1

                        Classifier: c2
                         Type: OR
                        Behavior: b2
                         Redirect:
                           Redirect nexthop
                           10.1.30.1

                    Total policy number is 2

                    # Check the traffic policy application records.
                    <DeviceA> display traffic-policy applied-record
                    Total records : 1
                    --------------------------------------------------------------------------------
                    Policy Type/Name                     Apply Parameter              Slot State
                    --------------------------------------------------------------------------------
                    p1                             10GE1/0/1(IN)               1 success
                    --------------------------------------------------------------------------------


Configuration Scripts
                    DeviceA
                    #
                    sysname DeviceA
                    #
                    vlan batch 10 20 30
                    #
                    acl number 3001
                     rule 5 permit ip source 192.168.100.0 0.0.0.255
                    #
                    acl number 3002
                     rule 5 permit ip source 192.168.101.0 0.0.0.255
                    #
                    traffic classifier c1 type or
                     if-match acl 3001
                    #
                    traffic classifier c2 type or
                     if-match acl 3002
                    #
                    traffic behavior b1
                     redirect nexthop 10.1.20.1
                    #
                    traffic behavior b2
                     redirect nexthop 10.1.30.1
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                     classifier c2 behavior b2 precedence 10
                    #
                    interface Vlanif10
                     ip address 10.1.10.2 255.255.255.0


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                                64
QoS Configuration
QoS Configuration                                                                      7 Redirection Configuration

                    #
                    interface Vlanif20
                     ip address 10.1.20.2 255.255.255.0
                    #
                    interface Vlanif30
                     ip address 10.1.30.2 255.255.255.0
                    #
                    interface 10GE1/0/1
                     port link-type trunk
                     port trunk allow-pass vlan 10
                     traffic-policy p1 inbound
                    #
                    interface 10GE1/0/2
                     port link-type trunk
                     port trunk allow-pass vlan 20
                    #
                    interface 10GE1/0/3
                     port link-type trunk
                     port trunk allow-pass vlan 30
                    #
                    return



