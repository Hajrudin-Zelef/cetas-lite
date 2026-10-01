---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-52
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [7253, 7399]
sha256: e4f9d8300ae617bacf8cc73a874efd766571b5077875e5c5e1a3c981250022e5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                          129
QoS Configuration                                                           9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                           based Rate Limiting Configuration

                    ----------------------------------------------------------------------------------------------
                       1        0               0           0               0           0          -
                          10000000                  0           0               0           0
                    ----------------------------------------------------------------------------------------------
                       2     2000                 54584           0               0           0        -
                             3000                5676736          0               0           0
                    ----------------------------------------------------------------------------------------------
                       3     5000                 49648           0               0           0        -
                             8000                5163392            0               0           0
                    ----------------------------------------------------------------------------------------------
                       4        0               0           0               0           0          -
                          10000000                  0           0               0           0
                    ----------------------------------------------------------------------------------------------
                       5     3000                49998            0               0           0        -
                             5000               5199792              0               0           0
                    ----------------------------------------------------------------------------------------------
                       6        0               0           0               0           0          -
                          10000000                  0           0               0           0
                    ----------------------------------------------------------------------------------------------
                       7        0               0           0               0           0          -
                          10000000                  0           0               0           0
                    ----------------------------------------------------------------------------------------------


Configuration Scripts
                    ●    DeviceB
                         #
                         sysname DeviceB
                         #
                         vlan batch 10 20 30
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 10 20 30
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 10 20 30
                          qos queue 2 shaping cir 2000 kbps pir 3000 kbps
                          qos queue 3 shaping cir 5000 kbps pir 8000 kbps
                          qos queue 5 shaping cir 3000 kbps pir 5000 kbps
                         #
                         return

                    ●    DeviceA
                         #
                         sysname DeviceA
                         #
                         vlan batch 10 20 30
                         #
                         interface 10GE1/0/1
                          port link-type access
                          port default vlan 10
                          port priority 5
                         #
                         interface 10GE1/0/2
                          port link-type access
                          port default vlan 20
                          port priority 3
                         #
                         interface 10GE1/0/3
                          port link-type access
                          port default vlan 30
                          port priority 2
                         #
                         interface 10GE1/0/4
                          port link-type trunk
                          port trunk allow-pass vlan 10 20 30


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                             130
QoS Configuration                                             9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                             based Rate Limiting Configuration

                        #
                        return


9.6.4 Example for Configuring Traffic Shaping Based on
Priority Mapping in a DiffServ Domain

Networking Requirements
                    In Figure 9-11, packets of voice, video, and data services from the user side
                    traverse DeviceA, DeviceB, and DeviceC to reach the external network.


                    Figure 9-11 Network diagram for configuring traffic shaping
                         NOTE

                        In this example, interface 1 and interface 2 represent 10GE 1/0/1 and 10GE 1/0/2,
                        respectively.




                    Packets of voice, video, and data services are identified by 802.1p priorities 6, 5,
                    and 2, respectively. The interface bandwidth is limited to 10000 kbit/s. However,
                    jitter may occur when packets from interface 2 on DeviceB reach DeviceC. To
                    reduce jitter and ensure the bandwidth for various services, the following
                    bandwidth requirements must be met:


                    Table 9-10 Bandwidth provided for each service on DeviceB

                     Service Type                   CIR (kbit/s)                    PIR (kbit/s)

                     Voice                          3000                            5000

                     Video                          5000                            8000

                     Data                           2000                            3000




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   131
QoS Configuration                                                 9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                 based Rate Limiting Configuration


Procedure
         Step 1 On DeviceB, create a VLAN and configure interfaces so that users can access the
                network through DeviceB.
                    # Create VLAN 10.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10

                    # Add 10GE 1/0/1 and 10GE 1/0/2 to VLAN 10 as trunk interfaces.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 10
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type trunk
                    [DeviceB-10GE1/0/2] port trunk allow-pass vlan 10
                    [DeviceB-10GE1/0/2] quit

