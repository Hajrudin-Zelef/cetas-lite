---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-64
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "latency", "parameters", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [9033, 9194]
sha256: d2cfa6eb2b602730c2335303b74897d05e2751043d8f2b71fc9fd04f4e86303b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                         When queue-based congestion monitoring is enabled, the lower buffer threshold is 10% of
                         the queue buffer space and the upper buffer threshold is 90% of the queue buffer space by
                         default.

         Step 4 (Optional) Configure the upper and lower buffer thresholds.
                    qos [ queue queue-index ] buffer-monitoring percent low low-percent high high-percent

                    ----End

Verifying the Configuration
                    ●    Run the display qos buffer-monitoring result interface { interface-type
                         interface-number | interface-name } command to check the real-time buffer
                         usage of queues.
                    ●    Run the display qos buffer-monitoring result history interface { interface-
                         type interface-number | interface-name } [ queue queue-index ] [ record-
                         number record-number ] command to check historical congestion monitoring
                         information about queues.



11.8 Maintaining Congestion Management
Context


                        NOTICE

                    Statistics cannot be restored after they are cleared. Exercise caution when clearing
                    the statistics.


Procedure
                    ●    Clear queue-based traffic statistics.
                         reset qos queue statistics { interface { interface-type interface-number | interface-name } | slot slot-
                         id }
                    ●    Clear statistics on the buffer usage.
                         reset qos buffer-usage [ slot slot-id | interface { interface-type interface-number | interface-name } ]

                    ----End


11.9 Example for Configuring Congestion Management
Networking Requirements
                    Host1 and Host2 provide voice, video, and data services. Traffic from these services
                    is transmitted through DeviceB and then DeviceA. To reduce the impact of

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                164
QoS Configuration
QoS Configuration                                                     11 Congestion Management Configuration


                    network congestion and guarantee high-priority services that require low latency,
                    set congestion management parameters according to Table 11-3.

                    Table 11-3 Congestion management parameters

                     Service Type        Color                  CoS            Scheduling          Scheduling
                                                                               Mode                Weight

                     Voice               Green                  EF             PQ                  -

                     Video               Yellow                 AF3            WDRR                100

                     Data                Red                    AF1            WDRR                50




                    Figure 11-4 Network diagram of congestion management
                          NOTE

                        In this example, interface 1, interface 2, and interface 3 represent 10GE 1/0/1, 10GE 1/0/2,
                        and 10GE 1/0/3, respectively.




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


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     165
QoS Configuration
QoS Configuration                                                     11 Congestion Management Configuration

                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] portswitch
                    [DeviceB-10GE1/0/3] port link-type trunk
                    [DeviceB-10GE1/0/3] port trunk allow-pass vlan 100 200
                    [DeviceB-10GE1/0/3] quit

         Step 2 Configure priority mapping to map 802.1p values in voice, video, and data packets
                to different CoS values and colors.
                    # Create DiffServ domain ds1, map 802.1p values 6, 5, and 2 to CoS values EF,
                    AF3, and AF1, respectively, and color the packets green, yellow, and red,
                    respectively.
                    [DeviceB] diffserv domain ds1
                    [DeviceB-dsdomain-ds1] 8021p-inbound 6 phb ef green
                    [DeviceB-dsdomain-ds1] 8021p-inbound 5 phb af3 yellow
                    [DeviceB-dsdomain-ds1] 8021p-inbound 2 phb af1 red
                    [DeviceB-dsdomain-ds1] quit

                    # Bind the DiffServ domain to the inbound interfaces 10GE 1/0/1 and 10GE 1/0/2
                    on DeviceB.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] trust 8021p outer
                    [DeviceB-10GE1/0/1] trust upstream ds1
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] trust 8021p outer
                    [DeviceB-10GE1/0/2] trust upstream ds1
                    [DeviceB-10GE1/0/2] quit

         Step 3 Configure congestion management. Set scheduling parameters such as the
                scheduling mode and weight to implement differentiated scheduling for queues
                with different priorities.
                    # Configure scheduling parameters for queues with different CoS values on the
                    outbound interface 10GE 1/0/3 of DeviceB.
                    The following uses the configuration of the S6780-H, S6750E-S, S6750-S, S6730E-
                    H-V2, S6730-H-V2, S6750-H, S5732-H-V2, S5755-S, S5755E-H and S5755-H series
                    as an example. For the configuration of other models, see 11.5 Configuring
                    Congestion Management.
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] qos pq 5 to 7 drr 0 to 4
                    [DeviceB-10GE1/0/3] qos queue 3 drr weight 100
                    [DeviceB-10GE1/0/3] qos queue 1 drr weight 50
                    [DeviceB-10GE1/0/3] quit
                    [DeviceB] quit

                    ----End

Verifying the Configuration
                    # Check the configuration of DiffServ domain ds1.
                    <DeviceB> display diffserv domain name ds1
                    Diffserv domain name:ds1
                     8021p-inbound 0 phb be green
                     8021p-inbound 1 phb af1 green
                     8021p-inbound 2 phb af1 red
                     8021p-inbound 3 phb af3 green
                     8021p-inbound 4 phb af4 green
                     8021p-inbound 5 phb af3 yellow


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                             166

