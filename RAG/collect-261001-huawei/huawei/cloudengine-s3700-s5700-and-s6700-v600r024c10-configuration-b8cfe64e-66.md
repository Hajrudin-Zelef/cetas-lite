---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-66
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [9376, 9507]
sha256: bac84155b9f06028844fa9b70e399c3ccd2c985c3e037fa548d910da42e50b25
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] portswitch
                    [DeviceB-10GE1/0/3] port link-type trunk
                    [DeviceB-10GE1/0/3] port trunk allow-pass vlan 2 5 6
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

                    # Bind the DiffServ domain to the inbound interface 10GE 1/0/1 of DeviceB.
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] trust 8021p inner
                    [DeviceB-10GE1/0/1] trust upstream ds1
                    [DeviceB-10GE1/0/1] quit

         Step 3 Configure congestion avoidance.
                    # On DeviceB, create WRED drop profile wred1 and set parameters for green,
                    yellow, and red packets in the WRED drop profile.
                    [DeviceB] drop-profile wred1
                    [DeviceB-drop-wred1] color green low-limit 80 high-limit 100 discard-percentage 10 //Configure the
                    WRED drop profile and set the upper and lower drop thresholds and maximum drop probability for green
                    packets.
                    [DeviceB-drop-wred1] color yellow low-limit 60 high-limit 80 discard-percentage 20 //Configure the
                    device to discard packets with the maximum drop probability of 20% when the percentage of the yellow
                    packet length to the queue length reaches 60%. Configure the device to discard all newly arrived packets
                    when the percentage of the yellow packet length to the queue length reaches 80%.
                    [DeviceB-drop-wred1] color red low-limit 40 high-limit 60 discard-percentage 40
                    [DeviceB-drop-wred1] quit

                    # Apply WRED drop profile wred1 to outbound interfaces 10GE 1/0/2 and 10GE
                    1/0/3 on DeviceB.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] qos queue 5 wred wred1
                    [DeviceB-10GE1/0/2] qos queue 3 wred wred1
                    [DeviceB-10GE1/0/2] qos queue 1 wred wred1
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] qos queue 5 wred wred1
                    [DeviceB-10GE1/0/3] qos queue 3 wred wred1
                    [DeviceB-10GE1/0/3] qos queue 1 wred wred1
                    [DeviceB-10GE1/0/3] quit

         Step 4 Configure congestion management. Set scheduling parameters such as the
                scheduling mode and weight to implement differentiated scheduling for queues
                with different priorities.
                    # Set scheduling parameters for queues with different CoS values on outbound
                    interfaces 10GE 1/0/2 and 10GE 1/0/3 on DeviceB.
                    The following uses the configuration of the S6780-H, S6750E-S, S6750-S, S6730E-
                    H-V2, S6730-H-V2, S6750-H, S5732-H-V2, S5755-S, S5755E-H and S5755-H series

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              170
QoS Configuration
QoS Configuration                                                                 11 Congestion Management Configuration


                    as an example. For the configuration of other models, see 11.5 Configuring
                    Congestion Management.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] qos pq 5 to 7 drr 0 to 4 //Configure PQ scheduling for queues 5 to 7 and WDRR
                    scheduling for queues 0 to 4.
                    [DeviceB-10GE1/0/2] qos queue 3 drr weight 100 //Set the WDRR scheduling weight of queue 3 to 100.
                    [DeviceB-10GE1/0/2] qos queue 1 drr weight 50 //Set the WDRR scheduling weight of queue 1 to 50.
                    According to the preceding configurations, packets in queue 1 and queue 3 are scheduled based on the
                    ratio of 1:2.
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] qos pq 5 to 7 drr 0 to 4 //Configure PQ scheduling for queues 5 to 7 and WDRR
                    scheduling for queues 0 to 4.
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
                     8021p-inbound 6 phb ef green
                     8021p-inbound 7 phb cs7 green
                     8021p-outbound be green map 0
                     8021p-outbound be yellow map 0
                     8021p-outbound be red map 0
                     ...

                    In the DiffServ domain, 802.1p values 6, 5, and 2 are mapped to CoS values EF,
                    AF3, and AF1, respectively, and packets are colored green, yellow, and red,
                    respectively.
                    # Check the WRED drop profile configuration.
                    [DeviceB] display drop-profile name wred1
                    Drop-profile[7]: wred1
                    Color     Mode       Low-limit High-limit Unit Discard(%)
                    -----------------------------------------------------------------
                    Green Percentage 80               100         %       10
                    Yellow Percentage 60               80        %       20
                    Red      Percentage 40            60        %       40
                    -----------------------------------------------------------------


Configuration Scripts
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    vlan batch 2 5 to 6
                    #
                    diffserv domain ds1
                     8021p-inbound 2 phb af1 red


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       171
QoS Configuration
QoS Configuration                                                          11 Congestion Management Configuration

