---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-56
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [7775, 7944]
sha256: 4fd3a22796b2acae585a17522b82913cc0bec5bd9ccf9f9f640921ca4da2a19d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

9.7.3 Configuring Rate Limiting on the Management Interface
Context
                    If there is heavy traffic on the management interface due to malicious attacks or
                    network exceptions, the CPU of the device becomes overloaded, and system
                    operations are impacted. You can configure rate limiting on the management
                    interface to limit the rate of traffic entering the device through the management
                    interface, thereby ensuring the system runs properly.

                          NOTE

                         This function is supported only by the S6780-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-
                         V2, S6750-H, S5735E-L-V2, S5735-L-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735R-L-
                         V2, S5755-S, S5755E-H, S5755-H, and S5732-H-V2 series.
                         This function can be configured only on the management interface.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the management interface view.
                    interface meth 0/0/0

         Step 3 Configure rate limiting on the management interface.
                    qos lr pps packets

                    By default, the rate limit on the management interface is 3000 pps.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                138
QoS Configuration                                              9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                              based Rate Limiting Configuration


                          NOTE

                        A small rate limit may impact the FTP, Telnet, SFTP, STelnet, and SSH functions.

                    ----End

Verifying the Configuration
                    ●    Run the display this interface command in the management interface view.
                         The Over-car-pps field in the command output indicates the packets that are
                         discarded due to rate limiting on the management interface.
                    ●    Run the display interface meth 0/0/0 command in any view. The Over-car-
                         pps field in the command output indicates the packets that are discarded due
                         to rate limiting on the management interface.

9.7.4 Example for Configuring Traffic Policing to Limit the
Rate on an Interface

Networking Requirements
                    In Figure 9-13, the host sends packets through DeviceA. It is required that the
                    bandwidth of the packets sent by the host should not exceed 100 Mbit/s.

                    Figure 9-13 Networking of interface-based rate limiting
                          NOTE

                        In this example, interface 1 represents 10GE 1/0/1.




Procedure
         Step 1 Configure a CAR profile.

                    # On DeviceA, create a CAR profile named car1 to police traffic from the host.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] qos car car1 cir 100000

         Step 2 Apply the CAR profile.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                  139
QoS Configuration                                                           9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                           based Rate Limiting Configuration


                    # On DeviceA, apply CAR profile car1 to the inbound direction of 10GE 1/0/1 to
                    police the traffic from the host.
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] qos car inbound car1
                    [DeviceA-10GE1/0/1] quit

                    ----End

Verifying the Configuration
                    # Check the CAR profile configuration.
                    [DeviceA] display qos car name car1
                     ----------------------------------------------------------------
                      CAR Name : car1
                      CAR Index : 0
                       car cir 100000 kbps cbs 800000 bytes
                      Applied number on behavior : 0
                      Applied number on interface inbound : 1
                       10GE1/0/1
                      Applied number on Eth-Trunk inbound : 0
                      Applied number on interface outbound : 0
                      Applied number on Eth-Trunk outbound : 0

                    # Send packets to 10GE 1/0/1 at the rates of 60000 kbit/s and 110000 kbit/s,
                    respectively, and then run the display qos car statistics command to check the
                    traffic statistics. If the configuration is successful, all packets are successfully
                    forwarded when they are sent to 10GE 1/0/1 at 60000 kbit/s; however, some
                    packets are discarded when packets are sent to 10GE 1/0/1 at 110000 kbit/s.

Configuration Scripts
                    DeviceA
                    #
                    sysname DeviceA
                    #
                    qos car car1 cir 100000 kbps
                    #
                    interface 10GE1/0/1
                     qos car inbound car1
                    #
                    return



9.8 Configuring a QoS Profile
Context
                    You can configure traffic policing and packet processing priorities in the QoS
                    profile view. Currently, a QoS profile can be bound only to an AAA service scheme
                    for user authorization.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a QoS profile and enter the QoS profile view.

Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                            140
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration

                    qos-profile profile-name

         Step 3 Configure traffic policing and packet processing priorities in the QoS profile view.
                    ●    Configure traffic policing.
                         car cir cir-value [ pir pir-value ] [ cbs cbs-value pbs pbs-value ] { inbound | outbound }

                    ●    Configure the device to re-mark the DSCP priority of IP packets.
                         remark dscp dscp-value { inbound | outbound }

                    ●    Configure the device to re-mark the 802.1p priority of VLAN packets.
                         remark 8021p 8021p-value { inbound | outbound }

                    ----End

Verifying the Configuration
                    ●    Run the display qos-profile name profile-name command to check the
                         configuration of a specified QoS profile.
                    ●    Run the display qos-profile all command to check the configuration of all
                         QoS profiles.


9.9 Maintaining Traffic Policing, Traffic Shaping, and
Interface-based Rate Limiting
Context


                        NOTICE

                    Flow-based traffic statistics cannot be restored after they are cleared. Exercise
                    caution when running the below commands.


