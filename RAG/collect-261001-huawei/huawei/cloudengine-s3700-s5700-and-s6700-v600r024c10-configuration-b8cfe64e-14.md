---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-14
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [1435, 1599]
sha256: 53a9e31d47606af400d99c040a1b2e47b93c8935c2ec0471b4518d4f4fab9a06
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                                In a traffic behavior, the deny action can be used with only traffic
                                statistics collection and flow mirroring. Even if other actions except traffic
                                statistics collection and flow mirroring are configured, they do not take
                                effect.
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

         Step 4 Apply a traffic policy.
                    ●    Apply a traffic policy to the system.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        24
QoS Configuration
QoS Configuration                                                                      4 Packet Filtering Configuration


                        a.    Apply a traffic policy to the system.
                              traffic-policy policy-name global [ slot slot-id ] { inbound | outbound }

                    ●   Apply a traffic policy to an interface.
                        a.    Enter the interface view.
                              interface interface-type interface-number

                        b.    Apply a traffic policy to the interface.
                              traffic-policy policy-name { inbound | outbound }

                        c.    Exit the interface view.
                              quit

                    ●   Apply a traffic policy to a VLAN.
                        a.    Create a VLAN and enter the VLAN view.
                              vlan vlan-id

                        b.    Apply a traffic policy to the VLAN.
                              traffic-policy policy-name { inbound | outbound }

                        c.    Exit the VLAN view.
                              quit

                    ●   Apply a traffic policy to a VPN instance.
                        a.    Create a VPN instance and enter the VPN instance view.
                              ip vpn-instance vpn-instance-name

                        b.    Apply a traffic policy to the VPN instance.
                              traffic-policy policy-name inbound

                                      NOTE

                                     A traffic policy can be applied to a VPN instance only in the inbound direction.
                                     Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S,
                                     S6730E-H-V2, S5755E-H, S5755-S and S5755-H series support this function.
                        c.    Exit the VPN instance view.
                              quit

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




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                           25
QoS Configuration
QoS Configuration                                                          4 Packet Filtering Configuration


Verifying the Configuration

                    Operation                                   Command

                    Check the configured traffic classifiers.   display traffic classifier [ classifier-
                                                                name ]

                    Check the configured traffic behaviors.     display traffic behavior [ behavior-
                                                                name ]

                    Check the configured traffic policy.        display traffic policy [ policy-name
                                                                [ classifier classifier-name ] ]

                    Check the traffic policy application        display traffic-policy applied-record
                    records.

                    Check TCAM delivery failures.               display system tcam fail-record [ slot
                                                                slot-id ]

                    Check the group indexes and rule            display system tcam service brief
                    counts occupied by different services.      [ slot slot-id ]
                                                                display system tcam service { cpcar
                                                                slot slot-id | service-name slot slot-id
                                                                [ chip chip-id ] }

                    Check the traffic policy application        display system tcam service traffic-
                    records.                                    policy

                    Check information about matched             display system tcam match-rules slot
                    rules.                                      slot-id

                    Check statistics on packets that match      display traffic-policy statistics
                    a traffic policy.

                    Check matching fields and actions           display system tcam acl group-
                    supported by a traffic policy in each       information
                    view.

                    Check information about the resources display traffic-policy pre-state
                    occupied by the traffic policy to be
                    applied to determine whether the
                    traffic policy can be successfully applied
                    after the configuration is committed.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  26
QoS Configuration
QoS Configuration                                                                    4 Packet Filtering Configuration




4.4 Example for Configuring MQC-based Packet
Filtering
Networking Requirements
                    In Figure 4-1, Host1, Host2, and Host3 communicate with each other through
                    DeviceA. For specific reasons, Host1 is allowed to receive traffic from Host2
                    through DeviceA but is not allowed to receive traffic from Host3.

                    Figure 4-1 Network diagram of packet filtering
                          NOTE

                         In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                         and 10GE 1/0/3, respectively.




