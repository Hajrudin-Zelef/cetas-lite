---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-15
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [1600, 1778]
sha256: 378a80ea215a1334ffbbbaf36a9a4c412a7448724903933a5b16eb7757238ac0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Procedure
         Step 1 Configure an ACL rule.
                    # On DeviceA, create ACL 3001 to match the traffic with source IP address
                    192.168.3.1 and destination IP address 192.168.1.1 (traffic from Host3 to Host1).
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] acl 3001
                    [DeviceA-acl4-advance-3001] rule permit ip source 192.168.3.1 0 destination 192.168.1.1 0
                    [DeviceA-acl4-advance-3001] quit

         Step 2 Configure a traffic classifier.
                    # On DeviceA, create traffic classifier c1 and reference ACL 3001 in the traffic
                    classifier.
                    [DeviceA] traffic classifier c1
                    [DeviceA-classifier-c1] if-match acl 3001
                    [DeviceA-classifier-c1] quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       27
QoS Configuration
QoS Configuration                                                                         4 Packet Filtering Configuration


         Step 3 Configure a traffic behavior.
                    # On DeviceA, create traffic behavior b1 and define the deny action in the traffic
                    behavior.
                    [DeviceA] traffic behavior b1
                    [DeviceA-behavior-b1] deny
                    [DeviceA-behavior-b1] quit

         Step 4 Configure a traffic policy and apply it to the outbound direction of 10GE 1/0/1.
                    # On DeviceA, create a traffic policy p1 and bind a traffic classifier to a traffic
                    behavior in the traffic policy.
                    [DeviceA] traffic policy p1
                    [DeviceA-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceA-trafficpolicy-p1] quit

                    # Apply traffic policy p1 to the outbound direction of 10GE 1/0/1.
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] traffic-policy p1 outbound
                    [DeviceA-10GE1/0/1] quit

                    ----End

Verifying the Configuration
                    # Check the ACL configuration.
                    <DeviceA> display acl 3001
                    Advanced ACL 3001, 1 rule
                    ACL's step is 5
                     rule 5 permit ip source 192.168.3.0 0 destination 192.168.1.0 0
                     (0 times matched)

                    # Check the traffic classifier configuration.
                    <DeviceA> display traffic classifier c1
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match acl 3001

                    # Check the traffic policy configuration.
                    <DeviceA> display traffic policy p1
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Deny


Configuration Scripts
                    DeviceA
                    #
                    sysname DeviceA
                    #
                    acl number 3001
                     rule 5 permit ip source 192.168.3.0 0.0.0.255 destination 192.168.1.0 0.0.0.255
                    #
                    traffic classifier c1 type or


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          28
QoS Configuration
QoS Configuration                                                              4 Packet Filtering Configuration

                     if-match acl 3001
                    #
                    traffic behavior b1
                     deny
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                    #
                    interface 10GE1/0/1
                     traffic-policy p1 outbound
                    #
                    return



4.5 Example for Configuring Access Control Based on
Source MAC Addresses
Networking Requirements
                    In Figure 4-2, users of an enterprise access the Internet through DeviceA. The
                    enterprise does not allow some hosts on the LAN to access the Internet. However,
                    users can still access the Internet from these hosts by changing host IP addresses,
                    and firewalls cannot prevent such unauthorized access based on IP addresses.
                    Access control based on source MAC addresses can be configured to solve this
                    problem. In this example, some hosts can be prevented from accessing the
                    Internet but can access DeviceA.

                    Figure 4-2 Network diagram
                          NOTE

                         In this example, interface 1 represents 10GE 1/0/1.




Procedure
         Step 1 Create a VLAN and configure interfaces.
                    # On DeviceA, create VLAN 10, configure VLANIF 10, and add 10GE 1/0/1 to the
                    VLAN.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan 10
                    [DeviceA-vlan10] quit
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] port link-type trunk
                    [DeviceA-10GE1/0/1] port trunk allow-pass vlan 10


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                               29
QoS Configuration
QoS Configuration                                                                        4 Packet Filtering Configuration

                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface vlanif 10
                    [DeviceA-Vlanif10] ip address 10.1.1.1 255.255.255.0

         Step 2 Configure an ACL rule.

                    # On DeviceA, create ACL 3001 to match the traffic with the destination IP
                    address 10.1.1.1/24.
                    [DeviceA] acl 3001
                    [DeviceA-acl4-advance-3001] rule 1 permit ip destination 10.1.1.0 0.0.0.255
                    [DeviceA-acl4-advance-3001] quit

         Step 3 Configure traffic classifiers.

                    # On DeviceA, create a traffic classifier c1 and reference ACL 3001 in the traffic
                    classifier.
                    [DeviceA] traffic classifier c1 type and
                    [DeviceA-classifier-c1] if-match acl 3001
                    [DeviceA-classifier-c1] quit

                    # On DeviceA, create traffic classifiers c2 to c4 to match MAC addresses of user
                    hosts.
                    [DeviceA] traffic classifier c2 type and
                    [DeviceA-classifier-c2] if-match source-mac 00e0-fc0d-0001
                    [DeviceA-classifier-c2] quit
                    [DeviceA] traffic classifier c3 type and
                    [DeviceA-classifier-c3] if-match source-mac 00e0-fc0d-0002
                    [DeviceA-classifier-c3] quit
                    [DeviceA] traffic classifier c4 type and
                    [DeviceA-classifier-c4] if-match source-mac 00e0-fc0d-0003
                    [DeviceA-classifier-c4] quit

         Step 4 Configure traffic behaviors.

                    # On DeviceA, create a traffic behavior b1 and configure the permit action in the
                    traffic behavior.
                    [DeviceA] traffic behavior b1
                    [DeviceA-behavior-b1] permit
                    [DeviceA-behavior-b1] quit

