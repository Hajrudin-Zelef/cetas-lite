---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-17
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [1960, 2095]
sha256: d5f75a5e023f0c187125da8f3c6b68e0f0be621858c7342099ace30b1de00ceb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     Interface statistics       display interface          All packets on an      Packets sent to
                     collection                                            interface              the CPU are
                                                                                                  included.




5.2 Configuration Precautions for Traffic Statistics
Collection

5.3 Configuring MQC-based Traffic Statistics Collection
Context
                    Traffic statistics collection allows the device to collect statistics on packets that
                    match traffic classification rules. Statistics on forwarded and discarded packets
                    that match traffic policies help to locate faults and check whether traffic policies
                    are correctly applied.

                          NOTE

                         ● If a traffic policy contains a lot of rules and traffic statistics have been cleared, wait for a
                           period before querying traffic statistics. Otherwise, only some or no statistics are
                           displayed.
                         ● If traffic classifiers bound to a traffic policy define many matching rules and the traffic
                           statistics collection action is defined, the device may respond slowly when querying
                           statistics based on MQC instances or traffic classifiers.
                         ● For details about MQC-related configuration precautions, see "Configuration Precautions
                           for MQC" in MQC Configuration.


Procedure
         Step 1 Configure a traffic classifier.
                    For details about how to configure a traffic classifier, see 3.4 Configuring a Traffic
                    Classifier in "MQC Configuration".
         Step 2 Configure a traffic behavior.
                    1.   (Optional) Configure the device to count the inter-frame gap and preamble of
                         Ethernet frames when the device collects traffic statistics.
                         qos statistics ifg enable
                    2.   Create a traffic behavior and enter the traffic behavior view, or enter the view
                         of an existing traffic behavior.
                         traffic behavior behavior-name
                    3.   Enable the traffic statistics collection function.
                         statistics enable [ history-record interval interval-value ]
                    4.   Exit the traffic behavior view.
                         quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            34
QoS Configuration
QoS Configuration                                                           5 Traffic Statistics Collection Configuration


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
                    ●    Apply a traffic policy to a VPN instance.
                         a.     Create a VPN instance and enter the VPN instance view.
                                ip vpn-instance vpn-instance-name
                         b.     Apply a traffic policy to the VPN instance.
                                traffic-policy policy-name inbound

                                        NOTE

                                       A traffic policy can be applied to a VPN instance only in the inbound direction.
                                       Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S,
                                       S6730E-H-V2, S5755E-H, S5755-S and S5755-H series support this function.
                         c.     Exit the VPN instance view.
                                quit
                    ●    Apply a traffic policy to a QoS group.
                         If the same traffic policy needs to be applied to multiple VLANs or interfaces,
                         you are advised to add the VLANs or interfaces to the same QoS group and
                         then apply the traffic policy to the QoS group.
                         a.     Create a QoS group and enter the QoS group view.
                                qos group group-name
                         b.     Add a specified interface or VLAN to the QoS group.
                                group-member { interface { interface-type interface-num | interface-name [ to interface-type
                                interface-num | interface-name ] } &<1-8> | vlan { vlanid [ to vlanid ] } &<1-8> }
                                Only one type of member can be specified.

Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                           35
QoS Configuration
QoS Configuration                                                       5 Traffic Statistics Collection Configuration


                        c.    Apply a traffic policy to the QoS group.
                              traffic-policy policy-name { inbound | outbound }

                        d.    Exit the QoS group view.
                              quit

                    ----End

Verifying the Configuration

                    Operation                                          Command

                    Check the configured traffic classifiers.          display traffic classifier [ classifier-
                                                                       name ]

                    Check the configured traffic behaviors.            display traffic behavior [ behavior-
                                                                       name ]

                    Check the configured traffic policy.               display traffic policy [ policy-name
                                                                       [ classifier classifier-name ] ]

                    Check the traffic policy application               display traffic-policy applied-record
                    records.

                    Check TCAM delivery failures.                      display system tcam fail-record [ slot
                                                                       slot-id ]

