---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-29
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2020-04-09", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [3856, 4027]
sha256: de06bb0382f36c7f46a1541d03e818352820366ab48cf41f9000a29dda0bce9f
---

                    # Check the traffic classifier configuration.
                    <DeviceA> display traffic classifier
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match acl 3001

                    Total classifier number is 1

                    # Check the traffic policy configuration.
                    <DeviceA> display traffic policy
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Redirect:
                           Redirect nexthop
                           10.1.30.1 track nqa user test

                    Total policy number is 1

                    # Check the traffic policy application records.
                    <DeviceA> display traffic-policy applied-record
                    Total records : 1
                    --------------------------------------------------------------------------------
                    Policy Type/Name                     Apply Parameter               Slot State
                    --------------------------------------------------------------------------------
                    p1                             10GE1/0/3(IN)                1 success
                    --------------------------------------------------------------------------------

                    # Check the NQA test result. If "Completion:success" and "Lost packet ratio: 0 %"
                    are displayed, the NQA test succeeds and the link is normal.
                    <DeviceA> display nqa results test-instance user test

                    NQA entry(user, test) :test flag is active ,test type is ICMP
                    1 . Test 1 result The test is finished
                     Send operation times: 2              Receive response times: 2
                     Completion:success                  RTD over thresholds number: 0
                     Attempts number:1                    Drop operation number:0
                     Disconnect operation number:0           Operation timeout number:0
                     System busy operation number:0           Connection fail number:0
                     Operation sequence errors number:0 RTT Status errors number:0
                     Destination IP address:10.1.30.1
                     Min/Max/Average completion time: 3/4/3
                     Sum/Square-Sum completion time: 7/25
                     Last response packet receiving tim: 2020-04-09 09:55:38.2
                     Lost packet ratio: 0 %


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                                68
QoS Configuration
QoS Configuration                                                             7 Redirection Configuration


Configuration Scripts
                    ●   DeviceA
                        #
                        sysname DeviceA
                        #
                        vlan batch 10 20 30
                        #
                        acl number 3001
                         rule 5 permit ip source 192.168.101.0 0.0.0.255
                        #
                        traffic classifier c1 type or
                         if-match acl 3001
                        #
                        traffic behavior b1
                         redirect nexthop 10.1.30.1 track nqa user test
                        #
                        traffic policy p1
                         classifier c1 behavior b1 precedence 5
                        #
                        interface Vlanif10
                         ip address 10.1.20.2 255.255.255.0
                        #
                        interface Vlanif20
                         ip address 10.1.30.2 255.255.255.0
                        #
                        interface Vlanif30
                         ip address 10.1.10.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 30
                         traffic-policy p1 inbound
                        #
                        nqa test-instance user test
                         test-type icmp
                         destination-address ipv4 10.1.30.1
                         interval seconds 5
                         timeout 4
                         probe-count 2
                         frequency 11
                         start now
                        #
                        return

                    ●   DeviceC
                        #
                        sysname DeviceC
                        #
                        vlan batch 10
                        #
                        interface Vlanif10
                         ip address 10.1.20.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        return

                    ●   DeviceD


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           69
QoS Configuration
QoS Configuration                                                                     7 Redirection Configuration

                        #
                        sysname DeviceD
                        #
                        vlan batch 20
                        #
                        interface Vlanif20
                         ip address 10.1.30.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        return



7.7 Example for Configuring Redirection to Implement
Route Selection
Networking Requirements
                    In Figure 7-4, the PC needs to access the server through link A, link B, and link C
                    (is used by default) in descending order of priority.

                    Figure 7-4 Network diagram for configuring redirection to implement route
                    selection
                         NOTE

                        In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                        and 10GE 1/0/3, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        70
QoS Configuration
QoS Configuration                                                              7 Redirection Configuration




