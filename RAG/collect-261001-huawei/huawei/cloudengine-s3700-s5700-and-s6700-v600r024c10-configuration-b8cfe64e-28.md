---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-28
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [3699, 3855]
sha256: dbb8e5bdbe6b187348d8d68cd337751fbaa3b3928900ea0175495471dd44ff92
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

7.6 Example for Configuring Association Between
Redirection to a Next-Hop Address and NQA
Networking Requirements
                    In Figure 7-3, DeviceA is the upper-layer device of DeviceB and DeviceB is the user
                    gateway. There are reachable routes between DeviceA and DeviceB. DeviceA is
                    connected to the Internet through two links: high-speed link with the gateway at
                    10.1.20.1/24 and low-speed link with the gateway at 10.1.30.1/24. A default route
                    has been configured on DeviceA to ensure that traffic is transmitted through the
                    high-speed link by default. The customer requirements are as follows:
                    ●    Packets from the network segment 192.168.101.0/24 are redirected to the
                         low-speed link for transmission, alleviating the bandwidth pressure of the
                         high-speed link.
                    ●    If the low-speed link fails, packets from the network segment
                         192.168.101.0/24 can be rapidly switched back to the high-speed link to
                         minimize communication interruption caused by the link fault.

                    Figure 7-3 Network diagram of association between redirection to a next-hop
                    address and NQA
                          NOTE

                         In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                         and 10GE 1/0/3, respectively.




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       65
QoS Configuration
QoS Configuration                                                             7 Redirection Configuration




Procedure
         Step 1 Create VLANs and configure interfaces.
                    # Configure DeviceA.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 10 20 30
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
                    [DeviceA] interface 10ge1/0/3
                    [DeviceA-10GE1/0/3] portswitch
                    [DeviceA-10GE1/0/3] port link-type trunk
                    [DeviceA-10GE1/0/3] port trunk allow-pass vlan 30
                    [DeviceA-10GE1/0/3] quit
                    [DeviceA] interface vlanif 10
                    [DeviceA-Vlanif10] ip address 10.1.20.2 24
                    [DeviceA-Vlanif10] quit
                    [DeviceA] interface vlanif 20
                    [DeviceA-Vlanif20] ip address 10.1.30.2 24
                    [DeviceA-Vlanif20] quit
                    [DeviceA] interface vlanif 30
                    [DeviceA-Vlanif30] ip address 10.1.10.2 24
                    [DeviceA-Vlanif30] quit

                    # Configure DeviceC.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceC
                    [DeviceC] vlan batch 10
                    [DeviceC] interface 10ge 1/0/1
                    [DeviceC-10GE1/0/1] portswitch


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           66
QoS Configuration
QoS Configuration                                                                          7 Redirection Configuration

                    [DeviceC-10GE1/0/1] port link-type trunk
                    [DeviceC-10GE1/0/1] port trunk allow-pass vlan 10
                    [DeviceC-10GE1/0/1] quit
                    [DeviceC] interface vlanif 10
                    [DeviceC-Vlanif10] ip address 10.1.20.1 24
                    [DeviceC-Vlanif10] quit

                    # Configure DeviceD.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceD
                    [DeviceD] vlan batch 20
                    [DeviceD] interface 10ge 1/0/1
                    [DeviceD-10GE1/0/1] portswitch
                    [DeviceD-10GE1/0/1] port link-type trunk
                    [DeviceD-10GE1/0/1] port trunk allow-pass vlan 20
                    [DeviceD-10GE1/0/1] quit
                    [DeviceD] interface vlanif 20
                    [DeviceD-Vlanif20] ip address 10.1.30.1 24
                    [DeviceD-Vlanif20] quit

         Step 2 On DeviceA, configure an NQA test instance.
                    [DeviceA] nqa test-instance user test
                    [DeviceA-nqa-user-test] test-type icmp
                    [DeviceA-nqa-user-test] destination-address ipv4 10.1.30.1
                    [DeviceA-nqa-user-test] frequency 11
                    [DeviceA-nqa-user-test] probe-count 2
                    [DeviceA-nqa-user-test] interval seconds 5
                    [DeviceA-nqa-user-test] timeout 4
                    [DeviceA-nqa-user-test] start now
                    [DeviceA-nqa-user-test] quit

         Step 3 Configure an ACL rule.
                    # Create advanced ACL 3001 on DeviceA to allow packets from the network
                    segment 192.168.101.0/24 to pass through.
                    [DeviceA] acl 3001
                    [DeviceA-acl4-advance-3001] rule permit ip source 192.168.101.0 0.0.0.255
                    [DeviceA-acl4-advance-3001] quit

         Step 4 Configure a traffic classifier.
                    # Create a traffic classifier c1 on DeviceA and reference ACL 3001.
                    [DeviceA] traffic classifier c1
                    [DeviceA-classifier-c1] if-match acl 3001
                    [DeviceA-classifier-c1] quit

         Step 5 Configure a traffic behavior.
                    # Create a traffic behavior b1 on DeviceA to redirect packets to the IP address
                    10.1.30.1, and associate NQA with redirection to a next-hop address.
                    [DeviceA] traffic behavior b1
                    [DeviceA-behavior-b1] redirect nexthop 10.1.30.1 track nqa user test
                    [DeviceA-behavior-b1] quit

         Step 6 Configure a traffic policy and apply it to an interface.
                    # Create a traffic policy p1 on DeviceA, and bind the traffic classifier c1 and traffic
                    behavior b1 to the traffic policy.
                    [DeviceA] traffic policy p1
                    [DeviceA-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceA-trafficpolicy-p1] quit

                    # Apply traffic policy p1 to the inbound direction of 10GE 1/0/3.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      67
QoS Configuration
QoS Configuration                                                                                      7 Redirection Configuration

                    [DeviceA] interface 10ge 1/0/3
                    [DeviceA-10GE1/0/3] traffic-policy p1 inbound
                    [DeviceA-10GE1/0/3] quit
                    [DeviceA] quit

                    ----End

Verifying the Configuration
                    # Check the ACL configuration.
                    <DeviceA> display acl 3001
                    Advanced ACL 3001, 1 rule
                    ACL's step is 5
                     rule 5 permit ip source 192.168.101.0 0.0.0.255 (0 times matched)

