---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-54
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [7522, 7633]
sha256: 6232e2d2f89f5ed9f54916700afe621fc5f6f46d3927a3faae445f626037d935
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    Packets of voice, video, and data services are identified by 802.1p priorities 6, 5,
                    and 2, respectively. The interface bandwidth is limited to 10000 kbit/s. However,
                    jitter may occur when packets from interface 2 on DeviceB reach DeviceC. To
                    reduce jitter and ensure the bandwidth for various services, the following
                    bandwidth requirements must be met:

                    Table 9-11 Bandwidth provided for each service on DeviceB
                     Service Type                     CIR (kbit/s)                    PIR (kbit/s)

                     Voice                            3000                            5000

                     Video                            5000                            8000

                     Data                             2000                            3000




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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     134
QoS Configuration                                                        9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                        based Rate Limiting Configuration

                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type trunk
                    [DeviceB-10GE1/0/2] port trunk allow-pass vlan 10
                    [DeviceB-10GE1/0/2] quit

         Step 2 Configure the packet priority trusted by an interface.
                    # Configure an interface to trust 802.1p priorities of packets.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] trust 8021p outer //Configure the interface to trust 802.1p priorities so that packets
                    enter different queues based on the default mappings between 802.1p priorities, local priorities, and queues.
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
                    [DeviceB-10GE1/0/2] qos queue 6 shaping cir 3000 pir 5000 //Set the CIR value of voice packets entering
                    queue 6 to 3000 kbit/s according to the default mappings between PHBs and local priorities.
                    [DeviceB-10GE1/0/2] qos queue 5 shaping cir 5000 pir 8000
                    [DeviceB-10GE1/0/2] qos queue 2 shaping cir 2000 pir 3000
                    [DeviceB-10GE1/0/2] quit

                    ----End

Verifying the Configuration
                    # Check queue-based traffic statistics in the outbound direction of 10GE 1/0/2.
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
                              3000           5676736              0                0          0
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
                        6     3000              49998            0                0           0        -


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                              135
QoS Configuration                                                           9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                           based Rate Limiting Configuration

                             5000            5199792             0                0          0
                    ----------------------------------------------------------------------------------------------
                       7        0               0           0               0           0         -
                          10000000                  0           0               0           0
                    ----------------------------------------------------------------------------------------------


