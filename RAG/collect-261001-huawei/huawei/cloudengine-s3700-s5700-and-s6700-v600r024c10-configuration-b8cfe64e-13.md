---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-13
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [1276, 1434]
sha256: a4a8776c0fe50b371ca7e1954193414b39ba509cf1553753aee24b202ad1cbe3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    Check the configured traffic policy.               display traffic policy [ policy-name
                                                                       [ classifier classifier-name ] ]

                    Check the traffic policy application               display traffic-policy applied-record
                    records.

                    Check TCAM delivery failures.                      display system tcam fail-record [ slot
                                                                       slot-id ]

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




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                    21
QoS Configuration
QoS Configuration                                                                                       3 MQC Configuration


                    Operation                                               Command

                    Check matching fields and actions                       display system tcam acl group-
                    supported by a traffic policy in each                   information
                    view.

                    Check information about the resources display traffic-policy pre-state
                    occupied by the traffic policy to be
                    applied to determine whether the
                    traffic policy can be successfully applied
                    after the configuration is committed.


                    For details about the display commands, see the Command Reference.


3.9 Maintaining MQC
Context
                    Before re-collecting statistics on packets matching a traffic policy, clear all existing
                    statistics.


                        NOTICE

                    Traffic statistics cannot be restored after being cleared. Exercise caution when you
                    use this command.


Procedure
                    ●   Clear statistics on packets matching a traffic policy.
                        reset traffic-policy statistics { global [ slot slot-id ] | interface { interface-type interface-number |
                        interface-name } | vlan vlan-id | qos group group-id | vpn-instance vpn-instance-name } [ policy-
                        name ] [ inbound | outbound ] [ classifier-base classifier-name ] [ history ]

                               NOTE

                              Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-
                              V2, S5755E-H, S5755-S and S5755-H series support the vpn-instance parameter.

                    ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                         22
QoS Configuration
QoS Configuration                                                          4 Packet Filtering Configuration




                          4         Packet Filtering Configuration


                    4.1 Overview of Packet Filtering
                    4.2 Configuration Precautions for Packet Filtering
                    4.3 Configuring MQC-based Packet Filtering
                    4.4 Example for Configuring MQC-based Packet Filtering
                    4.5 Example for Configuring Access Control Based on Source MAC Addresses


4.1 Overview of Packet Filtering
                    Untrusted packets are those that present potential security risks or that users do
                    not wish to receive, and are quite common on modern networks. The packet
                    filtering function allows a device to immediately discard such untrusted packets in
                    order to improve overall network security.
                    With Modular QoS Command-Line Interface (MQC), a device is configured to
                    identify untrusted packets and discard them, as well as identify trusted packets
                    and permit them to pass through.
                    MQC-based packet filtering classifies packets more precisely and is more flexible
                    to deploy.


4.2 Configuration Precautions for Packet Filtering

4.3 Configuring MQC-based Packet Filtering
Context
                    A packet filtering-enabled device filters packets that match traffic classification
                    rules.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 23
QoS Configuration
QoS Configuration                                                                     4 Packet Filtering Configuration


                          NOTE

                         ● In a traffic behavior, if the permit action is configured in combination with other traffic
                           actions, these actions are taken on packets in the order in which they were configured; if
                           the deny action is configured in combination with other traffic actions (excluding traffic
                           statistics collection and flow mirroring), only the deny action is taken on packets.
                         ● If you specify a packet filtering action for packets matching an ACL rule, the system first
                           checks the action defined in the ACL rule. If the ACL rule defines permit, the action
                           specified in the traffic behavior is taken on the packets. If the ACL rule defines deny, the
                           packets are discarded regardless of the action specified in the traffic behavior.
                         ● For details about MQC-related configuration precautions, see "Configuration Precautions
                           for MQC" in MQC Configuration.


Procedure
         Step 1 Configure a traffic classifier.

                    For details about how to configure a traffic classifier, see 3.4 Configuring a Traffic
                    Classifier in "MQC Configuration".

         Step 2 Configure a traffic behavior.
                    1.   Create a traffic behavior and enter the traffic behavior view, or enter the view
                         of an existing traffic behavior.
                         traffic behavior behavior-name

                    2.   Configure a traffic behavior as needed.
                         –      Configure a traffic behavior so that packets matching the associated
                                traffic classifier are forwarded based on the original policy.
                                permit

                                If both permit and other actions are configured in a traffic behavior, these
                                actions are performed in sequence.
                         –      Configure a traffic behavior to deny packets matching the associated
                                traffic classifier.
                                deny

