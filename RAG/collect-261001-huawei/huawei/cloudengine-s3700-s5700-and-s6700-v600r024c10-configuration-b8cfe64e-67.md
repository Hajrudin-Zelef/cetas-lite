---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-67
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2015-11-04", "2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [9508, 9678]
sha256: 962361055b779132627abc65aa254c4f01d4b3ed0f45b67a50ac117c0dc57721
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     8021p-inbound 5 phb af3 yellow
                     8021p-inbound 6 phb ef green
                    #
                    drop-profile wred1
                     color green low-limit 80 high-limit 100 discard-percentage 10
                     color yellow low-limit 60 high-limit 80 discard-percentage 20
                     color red low-limit 40 high-limit 60 discard-percentage 40
                    #
                    interface 10GE1/0/2
                     port link-type trunk
                     port trunk allow-pass vlan 2 5 to 6
                     qos drr 0 to 4
                     qos queue 1 drr weight 50
                     qos queue 3 drr weight 100
                     qos queue 1 wred wred1
                     qos queue 3 wred wred1
                     qos queue 5 wred wred1
                    #
                    interface 10GE1/0/3
                     port link-type trunk
                     port trunk allow-pass vlan 2 5 to 6
                     qos drr 0 to 4
                     qos queue 1 drr weight 50
                     qos queue 3 drr weight 100
                     qos queue 1 wred wred1
                     qos queue 3 wred wred1
                     qos queue 5 wred wred1
                    #
                    interface 10GE1/0/1
                     port link-type trunk
                     port trunk allow-pass vlan 2 5 to 6
                     trust upstream ds1
                     trust 8021p inner
                    #
                    return



11.11 Example for Configuring Congestion Monitoring

Networking Requirements
                    Servers store information about voice, video, and data services, and users obtain
                    data from the servers through DeviceB. When a large amount of data is
                    transmitted from the servers to hosts, the total rate of inbound interfaces is higher
                    than the rate of the outbound interface. As a result, congestion may occur on the
                    outbound interface of DeviceB.
                    When packet loss occurs on the outbound interface of DeviceB due to queue
                    congestion, the administrator needs to know the buffer usage of each queue and
                    information about the packets that cause congestion to determine subsequent
                    network traffic planning and adjustment.

                    Figure 11-6 Network diagram for configuring congestion monitoring
                          NOTE

                         In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                         and 10GE 1/0/3, respectively.




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                     172
QoS Configuration
QoS Configuration                                                      11 Congestion Management Configuration




                          NOTE

                        This example is supported only by the S6750-H and S6780-H.


Procedure
         Step 1 Configure VLANs for each interface so that devices can communicate with each
                other at the link layer.

                    # Configure 10GE 1/0/3 on DeviceB as a trunk interface. Add 10GE 1/0/1 to VLAN
                    100, 10GE 1/0/2 to VLAN 200, and 10GE 1/0/3 to VLAN 100 and VLAN 200.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 100 200
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type access
                    [DeviceB-10GE1/0/1] port default vlan 100
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type access
                    [DeviceB-10GE1/0/2] port default vlan 200
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] portswitch
                    [DeviceB-10GE1/0/3] port link-type trunk
                    [DeviceB-10GE1/0/3] port trunk allow-pass vlan 100 200
                    [DeviceB-10GE1/0/3] quit

         Step 2 Enable congestion monitoring. Use the default lower and upper buffer thresholds.
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] qos buffer-monitoring enable
                    [DeviceB-10GE1/0/3] quit

                    ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                              173
QoS Configuration
QoS Configuration                                                         11 Congestion Management Configuration


Verifying the Configuration
                    # Check real-time congestion monitoring information about queues on 10GE
                    1/0/3.
                    <DeviceB> display qos buffer-monitoring result interface 10ge 1/0/3
                    Queue Time                       BufferUsage(Bytes)         Percent(%)
                    --------------------------------------------------------------------
                    0 2015-11-04 09:46:27.095                       0            0
                    1 2015-11-04 09:46:27.101                       0            0
                    2 2015-11-04 09:46:27.107                       0            0
                    3 2015-11-04 09:46:27.113                   1598064            100
                    4 2015-11-04 09:46:27.118                       0            0
                    5 2015-11-04 09:46:27.124                       0            0
                    6 2015-11-04 09:46:27.130                       0            0
                    7 2015-11-04 09:46:27.136                       0            0
                    --------------------------------------------------------------------


Configuration Scripts
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    vlan batch 100 200
                    #
                    interface 10GE1/0/1
                     port link-type access
                     port default vlan 100
                    #
                    interface 10GE1/0/2
                     port link-type access
                     port default vlan 200
                    #
                    interface 10GE1/0/3
                     port link-type trunk
                     port trunk allow-pass vlan 100 200
                     qos buffer-monitoring enable
                    #
                    return



11.12 Microburst Detection
                          NOTE

                         Only the S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S5755E-H, S5755-H, and S5732-H-
                         V2 series support this function.


11.12.1 Configuring Microburst Detection
Context
                    Microbursts are short spikes in network traffic, typically lasting several
                    milliseconds. During a microburst, the instantaneous data rate is tens or hundreds
                    times higher than the average rate, and in some cases even exceeds the interface
                    bandwidth. When a microburst exceeds the forwarding capability of a device, the
                    device buffers the burst data for later transmission. If the device does not have
                    sufficient available buffer, the excess data is discarded. This is known as packet
                    loss due to congestion. Such issues cannot be effectively handled using traditional

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                174

