---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-51
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [7102, 7252]
sha256: 84ab905d29055f5dc317655c9d59dcde107e9c1db898ac898de42099d5c728ad
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    By default, the traffic shaping rate for a queue is the maximum bandwidth of the
                    interface.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                   126
QoS Configuration                                               9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                               based Rate Limiting Configuration


                         NOTE

                        Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                        S5755E-H, S5755-S and S5755-H series support the mbytes parameter.

                    ----End

9.6.2 Verifying the Configuration
Procedure
                    ●   Run the display qos configuration interface [ { interface-type interface-
                        number | interface-name } ] command to check all QoS configurations on an
                        interface.
                    ●   Run the display qos queue statistics { slot slotid | interface { interface-type
                        interface-number | interface-name } } command to check queue-based traffic
                        statistics.
                    ----End

9.6.3 Example for Configuring Traffic Shaping to Limit the
Rate of Different Services
Networking Requirements
                    In Figure 9-10, three servers are deployed to provide voice, video, and data
                    services, and service packets traverse DeviceA, DeviceB, and DeviceC to reach the
                    external network. The interface connected to the voice service host joins VLAN 10;
                    the interface connected to the video service host joins VLAN 20; the interface
                    connected to the data service host joins VLAN 30.

                    Figure 9-10 Network diagram for configuring traffic shaping
                         NOTE

                        In this example, interface 1, interface 2, interface 3, and interface 4 represent 10GE 1/0/1,
                        10GE 1/0/2, 10GE 1/0/3, and 10GE 1/0/4, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        127
QoS Configuration                                            9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                            based Rate Limiting Configuration




                    Packets of voice, video, and data services are identified by 802.1p priorities 5, 3,
                    and 2 respectively. However, jitter may occur when packets from interface 2 on
                    DeviceB reach DeviceC. Table 9-9 lists the bandwidth requirements to limit jitter
                    and ensure services.

                    Table 9-9 Bandwidth for each service on DeviceB
                     Service Type                    CIR (kbit/s)                 PIR (kbit/s)

                     Voice                           3000                         5000

                     Video                           5000                         8000

                     Data                            2000                         3000




Procedure
         Step 1 On DeviceB, create VLANs and add interfaces to these VLANs so that users can
                access the network through DeviceB.
                    # Create VLANs 10, 20, and 30.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10 20 30

                    # Set the access mode of interfaces 10GE 1/0/1 and 10GE 1/0/2 to trunk, and add
                    them to VLANs 10, 20, and 30.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                128
QoS Configuration                                                        9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                        based Rate Limiting Configuration

                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 10 20 30
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] portswitch
                    [DeviceB-10GE1/0/2] port link-type trunk
                    [DeviceB-10GE1/0/2] port trunk allow-pass vlan 10 20 30
                    [DeviceB-10GE1/0/2] quit

         Step 2 Set priorities for DeviceA's interfaces connected to the hosts to differentiate
                packets of different services.
                    # On DeviceA, set the priorities of 10GE 1/0/1, 10GE 1/0/2, and 10GE 1/0/3 to 5, 3,
                    and 2, respectively, and add 10GE 1/0/4 to VLANs 10, 20, and 30.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 10 20 30
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] portswitch
                    [DeviceA-10GE1/0/1] port link-type access
                    [DeviceA-10GE1/0/1] port default vlan 10
                    [DeviceA-10GE1/0/1] port priority 5
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface 10ge 1/0/2
                    [DeviceA-10GE1/0/2] portswitch
                    [DeviceA-10GE1/0/2] port link-type access
                    [DeviceA-10GE1/0/2] port default vlan 20
                    [DeviceA-10GE1/0/2] port priority 3
                    [DeviceA-10GE1/0/2] quit
                    [DeviceA] interface 10ge 1/0/3
                    [DeviceA-10GE1/0/3] portswitch
                    [DeviceA-10GE1/0/3] port link-type access
                    [DeviceA-10GE1/0/3] port default vlan 30
                    [DeviceA-10GE1/0/3] port priority 2
                    [DeviceA-10GE1/0/3] quit
                    [DeviceA] interface 10ge 1/0/4
                    [DeviceA-10GE1/0/4] portswitch
                    [DeviceA-10GE1/0/4] port link-type trunk
                    [DeviceA-10GE1/0/4] port trunk allow-pass vlan 10 20 30
                    [DeviceA-10GE1/0/4] quit

         Step 3 Configure queue-based traffic shaping to limit the bandwidth of voice, video, and
                data services.
                    # Configure queue-based traffic shaping on DeviceB. Set the CIR values of voice,
                    video, and data services to 3000 kbit/s, 5000 kbit/s, and 2000 kbit/s, respectively.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] qos queue 5 shaping cir 3000 pir 5000 kbps
                    [DeviceB-10GE1/0/2] qos queue 3 shaping cir 5000 pir 8000 kbps
                    [DeviceB-10GE1/0/2] qos queue 2 shaping cir 2000 pir 3000 kbps
                    [DeviceB-10GE1/0/2] quit

                    ----End

Verifying the Configuration
                    # Display statistics about queues in the outbound direction on 10GE 1/0/2.
                    [DeviceB] display qos queue statistics interface 10ge 1/0/2
                     Queue CIR/PIR                 Passed      Pass Rate            Dropped       Drop Rate Drop Time
                           (% or kbps) (Packets/Bytes) (pps/bps)                (Packets/Bytes) (pps/
                    bps)
                     ----------------------------------------------------------------------------------------------
                        0        0               0           0               0           0         -
                           10000000                  0           0               0           0


