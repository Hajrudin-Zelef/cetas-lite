---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-61
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [8585, 8729]
sha256: c11318fa31cf40d5e74534d99c2acb4135090e373ee573bfa86ba6907e0d9939
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration
QoS Configuration                                             11 Congestion Management Configuration


                        Consequently, queue 7 is not full and the scheduler can process packets in
                        queues with lower priorities.

                        Figure 11-1 PQ scheduling




                    ●   Weighted Deficit Round Robin (WDRR) scheduling
                        WDRR schedules packets by taking the packet length into account, ensuring
                        that packets in all the queues are scheduled in turn.
                        In WDRR scheduling, the deficit indicates the bandwidth deficit of each queue.
                        The initial value is 0, and the system allocates bandwidth to each queue
                        based on the weight, with the deficit calculated as follows: If the deficit of a
                        queue is greater than 0, the queue participates in scheduling. The device
                        sends a packet and calculates the deficit based on that packet's length. If the
                        deficit of a queue is less than 0, the queue does not participate in scheduling,
                        and the current deficit is used as the initial value in the next round of
                        scheduling.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            155
QoS Configuration
QoS Configuration                                               11 Congestion Management Configuration


                        Figure 11-2 Queue weights




                        In Figure 11-2, the weights of queues 7, 6, 5, 4, 3, 2, 1, and 0 are set to 40,
                        30, 20, 10, 40, 30, 20, and 10, respectively. During scheduling, queues 7, 6, 5,
                        4, 3, 2, 1, and 0 obtain 20%, 15%, 10%, 5%, 20%, 15%, 10%, and 5% of the
                        bandwidth, respectively. Queues 7 and 6 are used as examples to describe
                        WDRR scheduling. For this example, assume that queue 7 obtains 400 byte/s
                        bandwidth and queue 6 obtains 300 byte/s bandwidth.
                        –   First round of scheduling
                            Deficit[7][1] = 0 + 400 = 400
                            Deficit[6][1] = 0 + 300 = 300
                            After a packet of 900 bytes in queue 7 and a packet of 400 bytes in
                            queue 6 are sent, the values are as follows:
                            Deficit[7][1] = 400 – 900 = –500
                            Deficit[6][1] = 300 – 400 = –100
                        –   Second round of scheduling
                            Deficit [7][2] = –500 + 400 = –100
                            Deficit [6][2] = –100 + 300 = 200
                            The packet in queue 7 is not scheduled because the deficit of queue 7 is
                            negative. After a packet of 300 bytes in queue 6 is sent, the value is as
                            follows:
                            Deficit [6][2] = 200 – 300 = –100

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             156
QoS Configuration
QoS Configuration                                             11 Congestion Management Configuration


                        –    Third round of scheduling
                             Deficit[7][3] = –100 + 400 = 300
                             Deficit[6][3] = –100 + 300 = 200
                             After a packet of 600 bytes in queue 7 and a packet of 500 bytes in
                             queue 6 are sent, the values are as follows:
                             Deficit[7][3] = 300 – 600 = –300
                             Deficit[6][3] = 200 – 500 = –300
                             This process is repeated, leading to queue 7 and queue 6 respectively
                             obtaining 20% and 15% of the bandwidth. As such, you can obtain the
                             required bandwidth by setting proper weights.
                        WDRR scheduling prevents packets in queues with lower priorities from being
                        starved out for a long time in PQ scheduling and uneven bandwidth
                        allocation when the packet lengths of queues are different or vary greatly.
                        However, WDRR scheduling has a disadvantage that services requiring a short
                        delay (such as voice services) cannot be scheduled in a timely manner.

Scheduling Sequence in Different Scheduling Modes
                    All eight interface queues can be configured with one scheduling mode or a
                    combination of different scheduling modes. PQ+WDRR is the most commonly
                    used scheduling mode.
                    If only PQ scheduling is used, the packets in lower priority queues may fail to
                    obtain bandwidth for long periods. If only WDRR scheduling is used, short-delay
                    services such as voice services cannot be scheduled preferentially. PQ+WDRR
                    scheduling offers the advantages of both PQ and WDRR scheduling while
                    offsetting their disadvantages.
                    Eight queues on the device interface are classified into two groups, and you can
                    specify PQ scheduling for certain queues and WDRR scheduling for others.

                    Figure 11-3 PQ+WDRR scheduling




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          157
QoS Configuration
QoS Configuration                                              11 Congestion Management Configuration


                    In Figure 11-3, the device first schedules traffic in queues 7, 6, and 5 in PQ mode,
                    before then scheduling traffic in queues 4, 3, 2, 1, and 0 in WDRR mode. Queues 4,
                    3, 2, 1, and 0 have their own weights.
                    Important protocol packets or short-delay service packets must be placed in
                    queues using PQ scheduling in order to be scheduled first. Other packets are
                    placed in queues using WDRR scheduling.


11.3 Configuration Precautions for Congestion
Management

11.4 Default Settings for Congestion Management
                    Table 11-1 Default settings for congestion management
                     Parameter                                 Default Setting

                     Congestion management                     Enabled

                     Interface queue scheduling                PQ scheduling

                     Weight in WDRR scheduling mode            1

                     Burst traffic buffering mode              Standard

                     Manually configured buffer                Not configured




11.5 Configuring Congestion Management
Context
                    Congestion management is a queue-based technology. With congestion
                    management configured, if packets are buffered in queues due to congestion on a
                    network, the device determines the sequence at which packets are forwarded
                    according to the defined scheduling policy, thereby preferentially scheduling high-
                    priority services.
                    There are eight queues on each interface, and each queue is able to use a
                    different scheduling mode. During queue scheduling, PQ queues are scheduled
                    first. Multiple PQ queues are scheduled in descending order of priority, with a
                    larger queue index indicating a higher priority. After completing PQ scheduling,
                    the device schedules the queues using WDRR scheduling.

