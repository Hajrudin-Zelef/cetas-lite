---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-58
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "latency", "parameters", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [8118, 8281]
sha256: 7e667223312a37eb12bb4ae20c12fe9ead0bdbd0255ff8d5d32bb9e166616998
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

10.5 Configuring WRED
Context
                    WRED randomly discards packets based on WRED parameter settings, preventing
                    global TCP synchronization and ensuring packets with higher priorities are less
                    likely to be discarded. According to a packet's color (drop priority), a WRED drop
                    profile defines absolute values of upper and lower drop thresholds, upper and
                    lower drop thresholds in percentage and the maximum drop probability.
                    Colors are used to determine whether packets are discarded during congestion
                    avoidance implementation and are independent of the mapping between internal
                    priorities and queues.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a WRED drop profile and enter the WRED drop profile view.
                    drop-profile drop-profile-name

                    By default, a WRED drop profile named default exists on the device.
                    A maximum of 63 WRED drop profiles, including the default drop profile, can be
                    configured on the device For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-
                    H-V2, S6750E-S, S6730E-H-V2, S5755E-H, S5755-S and S5755-H series. The default
                    drop profile cannot be deleted, and only parameters in the profile can be
                    modified.
                    A maximum of 30 WRED drop profiles, including the default drop profile, can be
                    configured on the deviceFor S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2,
                    S5735I-H-V2, S5735R-L-V2, S3710-H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2
                    series. The default drop profile cannot be deleted, and only parameters in the
                    profile can be modified.
         Step 3 Set WRED parameters.
                    color { green | red | yellow } { buffer-size low-limit low-buffer-size high-limit high-buffer-size | buffer-
                    size cell low-limit low-buffer-size-cell high-limit high-buffer-size-cell | low-limit low-limit-percentage
                    high-limit high-limit-percentage } discard-percentage discard-percentage

                    By default, the upper drop threshold in percentage, lower drop threshold in
                    percentage, and maximum drop probability in a WRED drop profile are all 100,
                    and absolute values of upper and lower drop thresholds are not configured.

                            NOTE

                           Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                           S5755E-H, S5755-S and S5755-H series support setting of the buffer-size parameter, that is,
                           setting of absolute values of upper and lower drop thresholds.

         Step 4 Exit the WRED drop profile view.
                    quit

         Step 5 Enter the interface view.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                 146
QoS Configuration
QoS Configuration                                                              10 Congestion Avoidance Configuration

                    interface { interface-type interface-number | interface-name }

         Step 6 Apply the WRED drop profile to a port queue.
                    qos queue queue-index wred drop-profile-name

         Step 7 Exit the interface view.
                    quit

                    ----End


10.6 Configuring the CFI as the Internal Drop Priority
Context
                    The Canonical Format Indicator (CFI) field, also known as the Drop Eligible
                    Indicator (DEI), in a VLAN tag identifies the drop priority of a packet. When the
                    rate of packets on the device exceeds the CIR, the value of the DEI field is set to 1.
                    In this case, the drop priority of the packets is high. When congestion occurs, the
                    device first discards the packets with the DEI field value of 1.
                    To configure the device to discard packets whose rate exceeds the CIR, configure
                    the CFI as the internal drop priority.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface { interface-type interface-number | interface-name }

         Step 3 Configure the CFI as the internal drop priority.
                    dei enable

         Step 4 Exit the interface view.
                    quit

                    ----End


10.7 Verifying the Configuration
Procedure
                    ●      Run the display drop-profile [ name drop-profile-name | brief ] command to
                           check the configuration of a WRED drop profile.
                    ●      Run the display qos configuration interface [ { interface-type interface-
                           number | interface-name } ] command to check all QoS configurations on an
                           interface.
                    ●      Run the display qos queue statistics { slot slotid | interface { interface-type
                           interface-number | interface-name } } command to check queue-based traffic
                           statistics.
                    ----End

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                   147
QoS Configuration
QoS Configuration                                                              10 Congestion Avoidance Configuration




10.8 Maintaining Congestion Avoidance
Context


                        NOTICE

                    Statistics cannot be restored after they are cleared. Exercise caution when clearing
                    the statistics.



Procedure
                    ●   Clear queue-based traffic statistics.
                        reset qos queue statistics { interface { interface-type interface-number | interface-name } | slot slot-
                        id }

                    ----End


10.9 Example for Configuring WRED
Networking Requirements
                    Host1 and Host2 provide voice, video, and data services, for which traffic is
                    transmitted through DeviceB and then DeviceA. To reduce the impact of network
                    congestion and guarantee high-priority, latency-sensitive services, set congestion
                    avoidance parameters according to Table 10-2.


                    Table 10-2 Congestion avoidance parameters

                     Service          Color             Lower             Upper              Drop              CoS
                     Type                               Drop              Drop               Probabilit
                                                        Threshold         Threshold          y (%)
                                                        (%)               (%)

                     Voice            Green             80                100                10                EF

                     Video            Yellow            60                80                 20                AF3

                     Data             Red               40                60                 40                AF1




                    Figure 10-5 Network diagram of congestion avoidance
                         NOTE

                        In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                        and 10GE 1/0/3, respectively.



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                 148

