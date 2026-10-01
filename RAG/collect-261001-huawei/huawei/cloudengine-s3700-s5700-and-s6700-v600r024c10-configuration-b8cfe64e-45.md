---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-45
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [6155, 6317]
sha256: 1b914f0dc99bcb1b39d774795b95abacd2215e88cf91164b874cc3fec8046c11
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration                                                  9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                  based Rate Limiting Configuration


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

9.5.2 Configuring Traffic Policing (Level-2 CAR)

Prerequisites
                    Traffic policing (level-1 CAR) has been configured. For details, see 9.5.1
                    Configuring MQC-based Traffic Policing (Level-1 CAR).


Context
                    The device supports 2-level traffic policing. After using MQC to implement traffic
                    policing (level-1 CAR) for service flows matching a traffic classifier in a traffic
                    policy, the device aggregates all the service flows matching the traffic classifiers
                    associated with the level-1 CAR in the same traffic policy and performs traffic
                    policing (level-2 CAR) for the aggregated flow. 2-level traffic policing implements
                    multiplexing of traffic statistics and provides fine-grained service control.


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                         112
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration


                          NOTE

                         Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                         S5755E-H, S5755-S and S5755-H series support this function.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 (Optional) Disable the device from counting the inter-frame gaps and preambles
                when the device calculates the traffic policing rate.
                    qos car ifg disable

         Step 3 Configure a CAR profile.
                    qos car car-name { percent percent-value | cir cir-value [ kbps | mbps | gbps ] [ cbs cbs-value [ bytes |
                    kbytes | mbytes ] [ pbs pbs-value [ bytes | kbytes | mbytes ] ] | pir pir-value [ kbps | mbps | gbps ] [ cbs
                    cbs-value [ bytes | kbytes | mbytes ] pbs pbs-value [ bytes | kbytes | mbytes ] ] ] }

                          NOTE

                    The CIR of the QoS CAR profile must be greater than the CIR of the CAR configured in level-1
                    CAR.

         Step 4 Configure aggregated CAR.
                    1.   Enter the traffic behavior view.
                         traffic behavior behavior-name

                    2.   Configure aggregated CAR.
                         car car-name share

                                NOTE

                               A traffic policy containing the aggregated CAR action can be applied only to the
                               inbound direction.

                    ----End

9.5.3 Verifying the Configuration

Procedure

                    Operation                                               Command

                    Check the configured traffic classifiers.               display traffic classifier [ classifier-
                                                                            name ]

                    Check the configured traffic behaviors.                 display traffic behavior [ behavior-
                                                                            name ]

                    Check the configured traffic policy.                    display traffic policy [ policy-name
                                                                            [ classifier classifier-name ] ]


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                113
QoS Configuration                                               9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                               based Rate Limiting Configuration


                    Operation                                         Command

                    Check the traffic policy application              display traffic-policy applied-record
                    records.

                    Check the traffic statistics after MQC-           display traffic-policy statistics
                    based traffic policing (level-1 CAR) is
                    enabled.

                    Check the CAR profile configuration.              display qos car [ name car-name ]



9.5.4 Example for Configuring MQC-based Traffic Policing
(Level-1 CAR)
Networking Requirements
                    In Figure 9-7, packets sent by Host1, Host2, and Host3 traverse DeviceA, DeviceB,
                    and DeviceC to reach the external network. Interface 1 (connected to Host1),
                    interface 2 (connected to Host2), and interface 3 (connected to Host3) join VLAN
                    10, VLAN 20, and VLAN 30, respectively.

                    Figure 9-7 Networking for configuring MQC to implement traffic policing
                         NOTE

                        In this example, interface 1, interface 2, interface 3, and interface 4 represent 10GE 1/0/1,
                        10GE 1/0/2, 10GE 1/0/3, and 10GE 1/0/4, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        114
QoS Configuration                                                9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                based Rate Limiting Configuration


