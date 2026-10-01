---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-9
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [670, 805]
sha256: eef6d2f3aa257849bb4ac40702a23b4181e64a781cdfc8eef89308ee50513c68
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

3.2 Understanding MQC
MQC Entities
                    MQC involves three entities: traffic classifier, traffic behavior, and traffic policy.
                    Traffic classifier
                    A traffic classifier defines a group of matching rules for classifying packets. To
                    configure a traffic classifier, specify the following:
                    ●   Traffic classifier name
                    ●   Traffic classification rules: The device supports a broad range of traffic
                        classification rules, including: link-layer rules (Layer 2 rules), network-layer
                        rules (Layer 3 rules), transport-layer rules (Layer 4 rules), access control list
                        (ACL) rules, and other rules.
                    ●   Relationship between rules in a traffic classifier
                        –    AND: If a traffic classifier contains ACL rules, a packet matches the traffic
                             classifier only when it matches one ACL rule and all the non-ACL rules. If
                             a traffic classifier does not contain any ACL rules, a packet matches the
                             traffic classifier only when it matches all the rules in the traffic classifier.
                        –    OR: A packet matches a traffic classifier if it matches one or more rules.
                        The traffic classifier c1 is used as an example, which defines the following
                        rules:
                        –    ACL rules: ACL 3000 and ACL 3001
                        –    Non-ACL rules: The VLAN ID is 10, and the 802.1p value in the outer tag
                             of a packet is 3.
                        OR: A packet matches traffic classifier c1 if its VLAN ID is 10 or the 802.1p
                        value in the outer tag is 3, the packet matches ACL 3000, or the packet
                        matches ACL 3001.
                        AND: A packet matches traffic classifier c1 only when its VLAN ID is 10 and
                        the 802.1p value in the outer tag is 3, and the packet matches ACL 3000 or
                        3001.
                    Traffic behavior
                    A traffic behavior defines an action to be taken on packets of a specified type. To
                    configure a traffic behavior, specify the following:
                    ●   Traffic behavior name
                    ●   Actions: The device supports actions such as packet filtering and traffic
                        statistics collection. If a traffic behavior defines multiple non-conflicting
                        actions, all these actions are successfully configured and take effect. If a
                        traffic behavior defines conflicting actions, the outcome is one of the
                        following:
                        –    When conflicting actions are defined in the traffic behavior view, the
                             system displays an error message and the command fails to be executed.
                        –    When a traffic policy that contains a traffic behavior defining conflicting
                             actions is applied, the system displays an error message and the traffic
                             policy fails to be applied.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   10
QoS Configuration
QoS Configuration                                                                      3 MQC Configuration


                    Traffic policy
                    A traffic policy binds traffic classifiers and traffic behaviors, and then actions
                    defined in traffic behaviors are taken on packets of specific types. In Figure 3-1,
                    multiple traffic classifiers and traffic behaviors can be bound to one traffic policy.

                    Figure 3-1 Multiple pairs of traffic classifiers and traffic behaviors bound to a
                    traffic policy




MQC Configuration Process
                    Figure 3-2 shows the MQC configuration process in the following steps:
                    1.   Configure a traffic classifier. The traffic classifier defines a group of matching
                         rules to classify traffic and is the basis for providing differentiated services.
                    2.   Configure a traffic behavior. The traffic behavior defines actions for controlling
                         packets that match rules.
                    3.   Create a traffic policy, and bind the traffic classifier and traffic behavior to the
                         traffic policy.
                    4.   Apply the traffic policy in the required view.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 11
QoS Configuration
QoS Configuration                                                                              3 MQC Configuration


                    Figure 3-2 MQC configuration process




3.3 Configuration Precautions for MQC

3.4 Configuring a Traffic Classifier
Prerequisites
                    Before configuring a traffic classifier, complete the following task:
                    ●   Configure an ACL if it is required to classify packets.

Context
                    Multiple rules can be configured in a traffic classifier as long as they do not
                    conflict with each other. You can configure rules based on your requirements.

                         NOTE

                        ● A traffic classifier in which matching rules are ANDed cannot define duplicate matching
                          rules. For example, a traffic classifier cannot define both if-match source-mac and an
                          ACL rule that matches the same source MAC address.
                        ● To match multiple fields of packets of the same type (such as Layer 2/IPv6/IPv4
                          packets) in a view, apply one traffic policy in the view, and specify multiple traffic
                          classifiers and associated traffic behaviors in the traffic policy. If both IPv4 and IPv6
                          packets need to be matched, create one traffic policy for each type of packet.


Procedure
         Step 1 Enter the system view.
                    system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                           12
QoS Configuration
QoS Configuration                                                                             3 MQC Configuration


         Step 2 Create a traffic classifier and enter the traffic classifier view, or enter the view of
                an existing traffic classifier.
                    traffic classifier classifier-name [ type { and | or } ]

                    and is the logical operator between the rules in a traffic classifier, which means
                    that:
                    ●     If the traffic classifier contains ACL rules, packets must match one ACL rule
                          and all the non-ACL rules in order to match the traffic classifier.
                    ●     If the traffic classifier does not contain any ACL rules, packets must match all
                          rules in order to match the traffic classifier.
                    The logical operator or means that packets match a traffic classifier if they match
                    one or more rules in the traffic classifier.

                    By default, the relationship between rules in a traffic classifier is OR.

                    For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-L-V2,
                    S3710-H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series, the device supports a
                    maximum of 1024 traffic classifiers.

