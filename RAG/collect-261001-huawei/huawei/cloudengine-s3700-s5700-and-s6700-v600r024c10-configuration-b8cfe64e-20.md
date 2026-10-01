---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-20
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [2418, 2572]
sha256: 5556bcfaf2447b78fd4a2306e23f1d726d6171f6a53633f67c112cd42ed92438
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                                For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2,
                                S5735R-L-V2, S3710-H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series:
                                remark 8021p 8021p-value

                         –      Re-marking DSCP values of IP packets
                                remark dscp { dscp-name | dscp-value }

                         –      Re-marking internal priorities
                                remark local-precedence { local-precedence-name | local-precedence-value } [ color ]

                         –      Re-marking local IDs
                                For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S,
                                S6730E-H-V2, S5755E-H, S5755-S and S5755-H series:
                                remark qos-local-id qos-local-id [ inbound-match ]

                                For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2,
                                S5735R-L-V2, S3710-H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series:
                                remark qos-local-id qos-local-id inbound-match

                    3.   Exit the traffic behavior view.
                         quit


         Step 3 Configure a traffic policy.
                    1.   Create a traffic policy and enter the traffic policy view, or enter the view of an
                         existing traffic policy.
                         traffic policy policy-name

                    2.   Bind a traffic behavior and a traffic classifier to the traffic policy.
                         classifier classifier-name behavior behavior-name [ precedence precedence-value ]

                    3.   Exit the traffic policy view.
                         quit

         Step 4 Apply the traffic policy.
                    ●    Apply a traffic policy to the system.
                         a.     Apply a traffic policy to the system.
                                traffic-policy policy-name global [ slot slot-id ] { inbound | outbound }

                    ●    Apply a traffic policy to an interface.
                         a.     Enter the interface view.
                                interface interface-type interface-number

                         b.     Apply a traffic policy to the interface.
                                traffic-policy policy-name { inbound | outbound }

                         c.     Exit the interface view.
                                quit

                    ●    Apply a traffic policy to a VLAN.
                         a.     Create a VLAN and enter the VLAN view.
                                vlan vlan-id

                         b.     Apply a traffic policy to the VLAN.
                                traffic-policy policy-name { inbound | outbound }

                         c.     Exit the VLAN view.
                                quit


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                       43
QoS Configuration
QoS Configuration                                                                         6 Re-marking Configuration


                    ●   Apply a traffic policy to a QoS group.
                        If the same traffic policy needs to be applied to multiple VLANs or interfaces,
                        you are advised to add the VLANs or interfaces to the same QoS group and
                        then apply the traffic policy to the QoS group.
                        a.    Create a QoS group and enter the QoS group view.
                              qos group group-name

                        b.    Add a specified interface or VLAN to the QoS group.
                              group-member { interface { interface-type interface-num | interface-name [ to interface-type
                              interface-num | interface-name ] } &<1-8> | vlan { vlanid [ to vlanid ] } &<1-8> }

                              Only one type of member can be specified.
                        c.    Apply a traffic policy to the QoS group.
                              traffic-policy policy-name { inbound | outbound }

                        d.    Exit the QoS group view.
                              quit

                    ----End


Verifying the Configuration

                    Operation                                           Command

                    Check the configured traffic classifiers.           display traffic classifier [ classifier-
                                                                        name ]

                    Check the configured traffic behaviors.             display traffic behavior [ behavior-
                                                                        name ]

                    Check the configured traffic policy.                display traffic policy [ policy-name
                                                                        [ classifier classifier-name ] ]

                    Check the traffic policy application                display traffic-policy applied-record
                    records.

                    Check TCAM delivery failures.                       display system tcam fail-record [ slot
                                                                        slot-id ]

                    Check the group indexes and rule                    display system tcam service brief
                    counts occupied by different services.              [ slot slot-id ]
                                                                        display system tcam service { cpcar
                                                                        slot slot-id | service-name slot slot-id
                                                                        [ chip chip-id ] }

                    Check the traffic policy application                display system tcam service traffic-
                    records.                                            policy

                    Check information about matched                     display system tcam match-rules slot
                    rules.                                              slot-id


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                           44
QoS Configuration
QoS Configuration                                                                     6 Re-marking Configuration


                    Operation                                        Command

                    Check statistics on packets that match           display traffic-policy statistics
                    a traffic policy.

                    Check matching fields and actions                display system tcam acl group-
                    supported by a traffic policy in each            information
                    view.

                    Check information about the resources display traffic-policy pre-state
                    occupied by the traffic policy to be
                    applied to determine whether the
                    traffic policy can be successfully applied
                    after the configuration is committed.




6.5 Example for Configuring Re-marking to Distinguish
Users
Networking Requirements
                    On the network shown in Figure 6-3, packets sent from Host1 and Host2 to
                    DeviceB are identified by different VLAN IDs (10 and 20, respectively). DeviceB re-
                    marks the VLAN packets received from Host1 and Host2 so that the internal
                    priority of the packets sent by Host1 is higher than that of the packets sent by
                    Host2 on DeviceA. This ensures the experience of services on Host1.

                    Figure 6-3 Network diagram for configuring re-marking
                         NOTE

