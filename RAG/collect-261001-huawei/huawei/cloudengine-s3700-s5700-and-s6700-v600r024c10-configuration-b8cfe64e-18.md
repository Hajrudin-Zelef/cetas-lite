---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-18
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [2096, 2281]
sha256: f0da619fd19bbe23b7f355b770041a6f1a3a4f48ad78b0e2a0047cabc54401ec
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    Check the group indexes and rule                   display system tcam service brief
                    counts occupied by different services.             [ slot slot-id ]
                                                                       display system tcam service { cpcar
                                                                       slot slot-id | service-name slot slot-id
                                                                       [ chip chip-id ] }

                    Check the traffic policy application               display system tcam service traffic-
                    records.                                           policy

                    Check information about matched                    display system tcam match-rules slot
                    rules.                                             slot-id

                    Check statistics on packets that match             display traffic-policy statistics
                    a traffic policy.

                    Check matching fields and actions                  display system tcam acl group-
                    supported by a traffic policy in each              information
                    view.




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                    36
QoS Configuration
QoS Configuration                                                     5 Traffic Statistics Collection Configuration


                    Operation                                         Command

                    Check information about the resources display traffic-policy pre-state
                    occupied by the traffic policy to be
                    applied to determine whether the
                    traffic policy can be successfully applied
                    after the configuration is committed.




5.4 Example for Configuring MQC-based Traffic
Statistics Collection
Networking Requirements
                    In Figure 5-1, Host1 sends packets with the 802.1p value of 6 to DeviceA.
                    Statistics on service packets need to be collected to properly allocate bandwidth
                    resources.

                    Figure 5-1 Network diagram of traffic statistics collection
                          NOTE

                         In this example, interface 1 represents 10GE 1/0/1.




Procedure
         Step 1 Set the host name of the device to DeviceA.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA

         Step 2 Configure an ACL rule.

                    # On DeviceA, create Layer 2 ACL 4000 to match packets with the 802.1p value of
                    6.
                    [DeviceA] acl 4000
                    [DeviceA-acl-L2-4000] rule permit 8021p 6
                    [DeviceA-acl-L2-4000] quit

         Step 3 Configure a traffic classifier.

                    # On DeviceA, create traffic classifier c1 and match ACL 4000.
                    [DeviceA] traffic classifier c1
                    [DeviceA-classifier-c1] if-match acl 4000
                    [DeviceA-classifier-c1] quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                   37
QoS Configuration
QoS Configuration                                                               5 Traffic Statistics Collection Configuration


         Step 4 Configure a traffic behavior.
                    # On DeviceA, create traffic behavior b1 and define the traffic statistics collection
                    action in the traffic behavior.
                    [DeviceA] traffic behavior b1
                    [DeviceA-behavior-b1] statistics enable
                    [DeviceA-behavior-b1] quit

         Step 5 Configure a traffic policy and apply it to the interface.
                    # On DeviceA, create traffic policy p1, in which traffic classifier c1 is associated
                    with traffic behavior b1.
                    [DeviceA] traffic policy p1
                    [DeviceA-trafficpolicy-p1] classifier c1 behavior b1
                    [DeviceA-trafficpolicy-p1] quit

                    # Apply traffic policy p1 to the inbound direction of 10GE 1/0/1.
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] traffic-policy p1 inbound
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] quit

                    ----End

Verifying the Configuration
                    # Check the ACL configuration.
                    <DeviceA> display acl 4000
                    L2 ACL 4000, 1 rule
                    ACL's step is 5
                     rule 5 permit 8021p 6 (0 times matched)

                    # Check the traffic classifier configuration.
                    <DeviceA> display traffic classifier c1
                     Traffic Classifier Information:
                      Classifier: c1
                        Type: OR
                        Rule(s):
                          if-match acl 4000

                    # Check the traffic policy configuration.
                    <DeviceA> display traffic policy p1
                     Traffic Policy Information:
                      Policy: p1
                        Classifier: c1
                          Type: OR
                        Behavior: b1
                          Statistics: enable

                    # Check traffic statistics.
                    <DeviceA> display traffic-policy statistics interface 10ge 1/0/1 inbound
                    Traffic policy: p1, inbound
                    --------------------------------------------------------------------------------
                     Slot: 1
                     Item                Packets              Bytes          pps          bps
                     -------------------------------------------------------------------------------
                     Matched                212185             22067448           1600       1379215
                      Passed              212185             22067448           1600        1379215
                      Dropped                   0                0           0          0
                       Filter               0                 0          0           0


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                            38
QoS Configuration
QoS Configuration                                                                  5 Traffic Statistics Collection Configuration

                      CAR                    0                0           0          0
                    -------------------------------------------------------------------------------

                    You can view the statistics on service packets on 10GE 1/0/1.

Configuration Scripts
                    DeviceA
                    #
                    sysname DeviceA
                    #
                    acl number 4000
                     rule 5 permit 8021p 6
                    #
                    traffic classifier c1 type or
                     if-match acl 4000
                    #
                    traffic behavior b1
                     statistics enable
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                    #
                    interface 10GE1/0/1
                     traffic-policy p1 inbound
                    #
                    return




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                               39
QoS Configuration
QoS Configuration                                                              6 Re-marking Configuration




                                   6          Re-marking Configuration


