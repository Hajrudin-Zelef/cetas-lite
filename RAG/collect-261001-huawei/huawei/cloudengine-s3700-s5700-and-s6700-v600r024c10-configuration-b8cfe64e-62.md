---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-62
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [8730, 8865]
sha256: 308cbc5d6c845de7d4bee8606afb7b686ca51f214958989b6d0930d4d2ba1cb3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Procedure
                    ●   For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S,
                        S6730E-H-V2, S5755E-H, S5755-S and S5755-H series:
                        a.   Enter the system view.
                             system-view


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             158
QoS Configuration
QoS Configuration                                                        11 Congestion Management Configuration


                        b.    Enter the interface view.
                              interface { interface-type interface-number | interface-name }

                        c.    Configure a scheduling mode for queues on the interface.
                              qos { pq { start-queue-index [ to end-queue-index ] } &<1-8> | drr { start-queue-index [ to end-
                              queue-index ] } &<1-8> }*

                              By default, the queue scheduling mode of an interface is PQ.
                        d.    (Optional) Configure the queue weight in WDRR scheduling mode.
                              qos queue queue-index drr weight weight-value

                              By default, the queue weight in WDRR scheduling mode is 1.
                    ●   For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-
                        L-V2, S3710-H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series:
                        a.    Enter the system view.
                              system-view

                        b.    Create a global scheduling profile and enter the scheduling profile view.
                              qos schedule-profile profile-name

                              By default, the device predefines a global scheduling profile named
                              default.

                                      NOTE

                                     In addition to the default global scheduling profile, a maximum of 11 global
                                     scheduling profiles can be created on the device. For the default global
                                     scheduling profile, you can only modify its queue scheduling mode and queue
                                     scheduling weight, but cannot delete it.
                        c.    Configure a scheduling mode for queues on an interface.
                              qos { pq { start-queue-index [ to end-queue-index ] } &<1-8> | drr { start-queue-index [ to end-
                              queue-index ] } &<1-8> }*

                              By default, the queue scheduling mode of an interface is PQ.
                        d.    (Optional) Configure the queue weight in WDRR scheduling mode.
                              qos queue queue-index drr weight weight-value

                              By default, the queue weight in WDRR scheduling mode is 1.
                        e.    Exit the scheduling profile view.
                              quit

                        f.    Enter the interface view.
                              interface { interface-type interface-number | interface-name }

                        g.    Apply the scheduling profile.
                              qos schedule-profile profile-name

                              By default, the default scheduling profile is applied to an interface.

                    ----End


Verifying the Configuration
                    Run the display qos queue statistics { slot slotid | interface { interface-type
                    interface-number | interface-name } } command to check queue-based traffic
                    statistics.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                            159
QoS Configuration
QoS Configuration                                               11 Congestion Management Configuration




11.6 (Optional) Configuring Buffer Management

11.6.1 Understanding Buffer Management
                    When packets are sent from an interface, the data buffer caches them to prevent
                    packet loss caused by congestion due to traffic bursts. When the data buffer is full,
                    the device stops caching packets. Instead, it discards the packets that do not enter
                    the buffer. If only queue scheduling is used, high-priority services cannot be
                    preferentially forwarded during heavy traffic bursts. Buffer management allows
                    the device to properly allocate the buffer to adjust and improve device
                    performance.
                    The buffer space available to a chip is shared by all interfaces of the chip, and the
                    buffer space available to an interface is shared by all queues on the interface. The
                    buffer can be classified into the chip, interface, and queue levels.

11.6.2 Configuring a Burst Traffic Buffering Mode
Context
                    The data buffer is allocated in static+dynamic mode. By default, each interface is
                    allocated with some static buffer to ensure the basic forwarding capability of
                    queues, while the remaining buffer is used as a dynamic buffer to ensure the
                    capability of forwarding burst traffic in queues. When burst packets enter a queue,
                    this dynamic buffer can be utilized. Table 11-2 lists the buffering modes
                    supported by the device.
                    When multiple interfaces send traffic to an interface, if there is burst traffic with a
                    volume exceeding the maximum allocated buffer, the device discards excess
                    packets. In this case, run the qos burst-mode enhanced command to set the
                    burst traffic buffering mode to enhanced to improve the device's capability to
                    forward burst traffic.

                    Table 11-2 Device buffering modes
                     Buffering Mode               Description                  Application Scenario

                     Standard                     Each interface reserves      Light burst traffic exists
                                                  some static buffer. When     on an interface.
                                                  traffic bursts, each
                                                  interface queue can
                                                  preempt a small part of
                                                  the global dynamic
                                                  buffer.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              160
QoS Configuration
QoS Configuration                                                        11 Congestion Management Configuration


                     Buffering Mode                    Description                             Application Scenario

                     Enhanced                          Each interface reserves                 Heavy burst traffic exists
                                                       some static buffer. When                on an interface.
                                                       traffic bursts, each
                                                       interface queue can
                                                       preempt a large part of
                                                       the global dynamic
                                                       buffer.

