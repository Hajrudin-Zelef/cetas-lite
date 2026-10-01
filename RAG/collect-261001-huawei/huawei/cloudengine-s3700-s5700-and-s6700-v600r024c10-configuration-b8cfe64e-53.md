---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-53
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [7400, 7521]
sha256: 2b366f5ffe332d9ebd510ab64eaceef1b947eabd629f89a3431483864bef2317
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

         Step 2 Configure priority mapping.
                    # Create DiffServ domain ds1 and map 802.1p priorities 6, 5, and 2 to PHBs CS7,
                    EF, and AF2, respectively.
                    [DeviceB] diffserv domain ds1
                    [DeviceB-dsdomain-ds1] 8021p-inbound 6 phb cs7 //Map 802.1p priorities of different service packets to
                    different PHBs to ensure that the service packets enter different queues.
                    [DeviceB-dsdomain-ds1] 8021p-inbound 5 phb ef
                    [DeviceB-dsdomain-ds1] 8021p-inbound 2 phb af2
                    [DeviceB-dsdomain-ds1] quit
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] trust upstream ds1
                    [DeviceB-10GE1/0/1] quit

         Step 3 Configure traffic shaping on an interface.
                    # Configure traffic shaping on an interface of DeviceB to limit the rate of the
                    interface to 10000 kbit/s.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] qos lr cir 10000 outbound //Configure interface-based rate limiting in the outbound
                    direction of the interface to limit the total bandwidth.

         Step 4 Configure queue-based traffic shaping on an interface.
                    # Configure queue-based traffic shaping on an interface of DeviceB. Set the CIR
                    values of voice, video, and data service packets to 3000 kbit/s, 5000 kbit/s, and
                    2000 kbit/s, respectively, and their PIR values to 5000 kbit/s, 8000 kbit/s, and 3000
                    kbit/s, respectively.
                    [DeviceB-10GE1/0/2] qos queue 7 shaping cir 3000 pir 5000 //Set the CIR value of voice packets entering
                    queue 7 to 3000 kbit/s according to the default mappings between PHBs and local priorities.
                    [DeviceB-10GE1/0/2] qos queue 5 shaping cir 5000 pir 8000
                    [DeviceB-10GE1/0/2] qos queue 2 shaping cir 2000 pir 3000
                    [DeviceB-10GE1/0/2] quit

                    ----End

Verifying the Configuration
                    # Check queue-based traffic statistics in the outbound direction of 10GE 1/0/2.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           132
QoS Configuration                                                        9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                        based Rate Limiting Configuration

                    [DeviceB] display qos queue statistics interface 10ge 1/0/2
                     Queue CIR/PIR                 Passed      Pass Rate            Dropped       Drop Rate Drop Time
                           (% or kbps) (Packets/Bytes) (pps/bps)                (Packets/Bytes) (pps/
                    bps)
                     ----------------------------------------------------------------------------------------------
                        0        0               0           0               0           0         -
                           10000000                  0           0               0           0
                     ----------------------------------------------------------------------------------------------
                        1        0               0           0               0           0         -
                           10000000                  0           0               0           0
                     ----------------------------------------------------------------------------------------------
                        2     2000              54584            0                0           0        -
                              3000            5676736             0                0          0
                     ----------------------------------------------------------------------------------------------
                        3        0               0           0               0           0         -
                           10000000                  0           0               0           0
                     ----------------------------------------------------------------------------------------------
                        4        0               0           0               0           0         -
                           10000000                  0           0               0           0
                     ----------------------------------------------------------------------------------------------
                        5     5000              49648            0                0           0        -
                              8000            5163392             0                0          0
                     ----------------------------------------------------------------------------------------------
                        6        0               0           0               0           0         -
                           10000000                  0           0               0           0
                     ----------------------------------------------------------------------------------------------
                        7     3000              49998            0                0           0        -
                              5000            5199792             0                0          0
                     ----------------------------------------------------------------------------------------------


Configuration Scripts
                    ●     DeviceB
                          #
                          sysname DeviceB
                          #
                          vlan batch 10
                          #
                          diffserv domain ds1
                           8021p-inbound 6 phb cs7 green
                           8021p-inbound 5 phb ef green
                           8021p-inbound 2 phb af2 green
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 10
                           trust upstream ds1
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 10
                           qos lr cir 10000 outbound
                           qos queue 2 shaping cir 2000 kbps pir 3000 kbps
                           qos queue 5 shaping cir 5000 kbps pir 8000 kbps
                           qos queue 7 shaping cir 3000 kbps pir 5000 kbps
                          #
                          return


9.6.5 Example for Configuring Traffic Shaping Based on
Trusted 802.1p Priorities

Networking Requirements
                    In Figure 9-11, packets of voice, video, and data services from the user side
                    traverse DeviceA, DeviceB, and DeviceC to reach the external network.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                          133
QoS Configuration                                                9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                based Rate Limiting Configuration


                    Figure 9-12 Network diagram for configuring traffic shaping
                          NOTE

                        In this example, interface 1 and interface 2 represent 10GE 1/0/1 and 10GE 1/0/2,
                        respectively.




