---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-60
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "latency", "throughput"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [8440, 8584]
sha256: 60ef9df0cd6390e504dbb6a04090bdef7d4fc2c3398a368905b17e7b8846f5e1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    --------------------------------------------------------------------------


Configuration Scripts
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    drop-profile wred1
                     color green low-limit 80 high-limit 100 discard-percentage 10
                     color yellow low-limit 60 high-limit 80 discard-percentage 20
                     color red low-limit 40 high-limit 60 discard-percentage 40
                    #
                    vlan batch 100 200
                    #
                    diffserv domain ds1
                     8021p-inbound 2 phb af1 red
                     8021p-inbound 5 phb af3 yellow
                     8021p-inbound 6 phb ef green
                    #
                    interface 10GE1/0/1
                     port link-type access
                     port default vlan 100
                     trust upstream ds1
                    #
                    interface 10GE1/0/2
                     port link-type access
                     port default vlan 200
                     trust upstream ds1
                    #


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                           151
QoS Configuration
QoS Configuration                                                   10 Congestion Avoidance Configuration

                    interface 10GE1/0/3
                     port link-type trunk
                     port trunk allow-pass vlan 100 200
                     qos queue 1 wred wred1
                     qos queue 3 wred wred1
                     qos queue 5 wred wred1
                    #
                    return




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         152
QoS Configuration
QoS Configuration                                             11 Congestion Management Configuration




                            11                Congestion Management
                                                        Configuration

                    11.1 Overview of Congestion Management
                    11.2 Understanding Congestion Management
                    11.3 Configuration Precautions for Congestion Management
                    11.4 Default Settings for Congestion Management
                    11.5 Configuring Congestion Management
                    11.6 (Optional) Configuring Buffer Management
                    11.7 (Optional) Configuring Congestion Monitoring
                    11.8 Maintaining Congestion Management
                    11.9 Example for Configuring Congestion Management
                    11.10 Example for Configuring Congestion Avoidance and Congestion
                    Management (PQ+WDRR Scheduling and WRED Profile)
                    11.11 Example for Configuring Congestion Monitoring
                    11.12 Microburst Detection


11.1 Overview of Congestion Management
Definition
                    When a network is intermittently congested and latency-sensitive services require
                    higher bandwidth than others, congestion management adjusts the order in which
                    packets are scheduled.

Purpose
                    On a traditional network, QoS issues are primarily caused by network congestion
                    due to insufficient network resources. Congestion extends the delay of packet
                    transmission, lowers the throughput, and consumes many resources. However,

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                        153
QoS Configuration
QoS Configuration                                              11 Congestion Management Configuration


                    network congestion is very common in a complex environment with a variety of
                    services. When network congestion occurs, packets are buffered in queues.
                    Congestion management schedules the packets based on a scheduling algorithm,
                    ensuring that QoS-demanding services (such as latency-sensitive services) are
                    preferentially scheduled.


11.2 Understanding Congestion Management
                    As network services continue to increase and users demand even higher network
                    quality, the limited bandwidth cannot meet such lofty requirements. As a result,
                    issues begin to arise due to congestion, such as longer delays and signal loss.
                    Congestion management is required when a network is intermittently congested,
                    and delay-sensitive services require higher QoS than delay-insensitive services. If
                    congestion persists on the network after congestion management is configured, it
                    is necessary to increase the bandwidth. Congestion management implements
                    queuing and scheduling when sending packet flows.

                    The device has eight queues on each interface in the outbound direction, and the
                    queues are identified by index numbers ranging from 0 to 7. The device
                    automatically sends classified packets to queues based on mappings between
                    local priorities and queues, and then schedules the packets using queue
                    scheduling mechanisms.

Common Queue Scheduling Mechanisms
                    ●   Priority Queuing (PQ) scheduling
                        PQ scheduling, also referred to as strict priority (SP) scheduling, schedules
                        packets in descending order of queue priority. This means that packets in
                        queues with a low priority can be scheduled only after all packets in high
                        priority queues have been scheduled. The device places core services into
                        high-priority queues and non-core services (such as email services) into low-
                        priority queues, ensuring that core services are processed preferentially. In this
                        case, non-core services are processed only when all core services are
                        processed.
                        In Figure 11-1, queues 7 to 0 are arranged in descending order of priority.
                        When the link transmits packets, queue 7 is preferentially processed, with
                        subsequent queues processed only when queue 7 becomes empty. Packets in
                        lower-priority queues are sent at the link rate when higher-priority queues are
                        empty. The packets in queue 6 are sent at the link rate when packets in queue
                        6 need to be sent and queue 7 is empty. The packets in queue 5 are sent at
                        the link rate when queue 6 and queue 7 are empty, and so on.
                        PQ scheduling is valid for short-delay services. If we assume that data flow X
                        is placed in queue 7 on each node, the packets of data flow X are processed
                        first when they reach a node.
                        PQ scheduling, however, may result in the starvation of packets in queues
                        with lower priorities. For example, if data flows placed into queue 7 arrive at
                        100% link rate in a period, the scheduler does not process flows in queues 0
                        to 6.
                        To prevent starvation of packets in some queues, upstream devices must
                        accurately define the service characteristics of data flows so that data flows
                        placed in queue 7 do not exceed a given percentage of the link capacity.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             154

