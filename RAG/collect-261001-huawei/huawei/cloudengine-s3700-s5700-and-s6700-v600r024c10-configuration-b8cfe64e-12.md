---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-12
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [1120, 1275]
sha256: 40247ed6079f5aa99b5a4bf443fc1fdc062b107664ff3d324baa33986e8edb91
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    MQC-based            vlan-stacking vlan vlan-id
                    selective QinQ       For details, see "Configuring QinQ" in CLI Configuration
                                         Guide > Ethernet Switching Configuration > VLAN
                                         Configuration.

                    Disabling URPF       ip urpf disable
                    check                For details, see "URPF Configuration" in CLI Configuration
                                         Guide > Security Configuration.
                                         Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-
                                         V2, S6750E-S, S6730E-H-V2, S5755E-H, S5755-S and S5755-H
                                         series support this function.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             18
QoS Configuration
QoS Configuration                                                                                  3 MQC Configuration


                     Action                  Command

                     Specifying a            network-slice-instance netsliceinstid
                     network slice           For details, see "Network Slicing Configuration" in CLI
                     instance                Configuration Guide > Network Slicing Configuration.



                    ----End


3.6 Configuring a Traffic Policy
Prerequisites
                    Before configuring a traffic policy, complete the following tasks:
                    ●    Configure a traffic classifier.
                    ●    Configure a traffic behavior.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a traffic policy and enter the traffic policy view, or enter the view of an
                existing traffic policy.
                    traffic policy policy-name

                    For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-L-V2,
                    S3710-H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series, the device supports a
                    maximum of 1024 traffic policies.
                    For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-
                    H-V2, S5755E-H, S5755-S and S5755-H series, the device supports a maximum of
                    2048 traffic policies.
         Step 3 Bind a traffic behavior and a traffic classifier to the traffic policy.
                    classifier classifier-name behavior behavior-name [ precedence precedence-value ]

                    For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-L-V2,
                    S3710-H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series, a maximum of 1024
                    traffic classifiers can be bound to a traffic policy.
                    For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-
                    H-V2, S5755E-H, S5755-S and S5755-H series, a maximum of 2048 traffic
                    classifiers can be bound to a traffic policy.

                    ----End


3.7 Applying a Traffic Policy
Prerequisites
                    Before applying a traffic policy, configure the traffic policy.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      19
QoS Configuration
QoS Configuration                                                                                   3 MQC Configuration


                         NOTE

                        If a traffic policy fails to be applied due to insufficient ACL resources on the device, you are
                        advised to delete the traffic policy configuration that fails to be applied. Otherwise, after
                        configurations are saved and the device is restarted, the configuration of other services that
                        are running properly cannot be restored.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Apply a traffic policy.
                    ●   Apply a traffic policy to the system.
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


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                           20
QoS Configuration
QoS Configuration                                                                            3 MQC Configuration


                              Only one type of member can be specified.
                        c.    Apply a traffic policy to the QoS group.
                              traffic-policy policy-name { inbound | outbound }

                        d.    Exit the QoS group view.
                              quit

                    ----End


3.8 Verifying the Configuration
Procedure

                    Operation                                          Command

                    Check the configured traffic classifiers.          display traffic classifier [ classifier-
                                                                       name ]

                    Check the configured traffic behaviors.            display traffic behavior [ behavior-
                                                                       name ]

