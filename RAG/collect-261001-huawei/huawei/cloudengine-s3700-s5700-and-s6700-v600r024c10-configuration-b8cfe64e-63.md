---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-63
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [8866, 9032]
sha256: 25c7f911d5a3791f2f81b5c3eabadce2048f4e457be580be7f3c9d84166dc8bd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     Extreme                           Each interface reserves                 Extremely heavy burst
                                                       some static buffer. When                traffic exists on an
                                                       traffic bursts, each                    interface.
                                                       interface queue can
                                                       preempt most of the
                                                       global dynamic buffer.




Procedure
                    ●   Configure a burst traffic buffering mode on the device.
                        a.    Enter the system view.
                              system-view

                        b.    Configure a burst traffic buffering mode.

                              For the S5732-H-V2, S6730E-H-V2, S6730-H-V2, S5735R-S-V2, S5735E-S-
                              V2, S5735-S-V2, S5735I-S-V2, S5735I-H-V2, S5735I-L-V2, S5735R-L-V2,
                              S5735E-L-V2 and S5735-L-V2:
                              qos burst-mode enhanced slot slot-id

                              For the S6780-H, S6750E-S, S6750-S, S6750-H, S5755-S S5755E-H and
                              S5755-H:
                              qos burst-mode { enhanced | extreme } slot slot-id

                    ●   Configure a burst traffic buffering mode on an interface.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the interface view.
                              interface { interface-type interface-number | interface-name }

                        c.    Configure a burst traffic buffering mode.

                              For the S5732-H-V2, S6730E-H-V2, S6730-H-V2, S5735R-S-V2, S5735E-S-
                              V2, S5735-S-V2, S5735I-S-V2, S5735I-H-V2, S5735I-L-V2, S5735R-L-V2,
                              S5735E-L-V2 and S5735-L-V2:
                              qos burst-mode enhanced

                              For the S6780-H, S6750E-S, S6750-S, S6750-H, S5755E-H, S5755-H,
                              S5755-S:
                              qos burst-mode { enhanced | extreme }

                    ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                            161
QoS Configuration
QoS Configuration                                                          11 Congestion Management Configuration


11.6.3 Configuring the Buffer Size
Context
                    If the device's current buffer cannot meet requirements, you can manually divide
                    the buffer.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface { interface-type interface-number | interface-name }

         Step 3 Configure the queue-level service buffer for outbound queues.
                    For the S6750E-S, S6750-S, S5755-S, S5755E-H, S5755-H, S6730E-H-V2, S6730-H-
                    V2, S5732-H-V2, S3710-H, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2,
                    S5735I-L-V2, S5735I-H-V2, S5735R-L-V2, S5735E-L-V2 and S5735-L-V2:
                    qos buffer queue queue-index shared-threshold dynamic dynamic-value

                    For the S6780-H and S6750-H:
                    qos buffer queue queue-index shared-threshold { static bytes-value { bytes | kbytes | mbytes } |
                    dynamic dynamic-value }

                          NOTE

                         Before running this command, ensure that the qos burst-mode command is not configured
                         in the system view or interface view.
                         For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-L-V2, S3710-
                         H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series, the qos buffer queue command can be
                         configured on a maximum of 28 interfaces on the device.

                    ----End

11.6.4 Configuring the Device to Generate an Alarm When the
Queue Buffer Usage Exceeds the Threshold
Context
                    The supported bandwidth on each interface is fixed. Once the traffic rate exceeds
                    the interface bandwidth and the used queue buffer exceeds the configured
                    threshold, the device starts discarding excess traffic. Run the qos buffer overrun
                    threshold command to set the threshold of the queue buffer usage, and run the
                    qos buffer overrun alarm enable command to enable the device to generate an
                    alarm when the queue buffer usage exceeds the threshold.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Set the threshold of the queue buffer usage.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          162
QoS Configuration
QoS Configuration                                                 11 Congestion Management Configuration

                    qos buffer overrun threshold percent

         Step 3 Configure the device to generate an alarm when the queue buffer usage exceeds
                the threshold.
                    qos buffer overrun alarm enable

                    ----End

11.6.5 Verifying the Configuration
Procedure
                    ●    Run the display qos buffer-usage command to check the buffer usage.

                    ----End


11.7 (Optional) Configuring Congestion Monitoring
                          NOTE

                        The congestion monitoring function is supported only by the S6750-H, S6780-H, S5755-S
                        series.


11.7.1 Understanding Congestion Monitoring
                    Queue
                    A queue is used to schedule packets and limit the bandwidth of an interface. Each
                    queue has a buffer space. If the device is congested, packets will go into the buffer
                    space of queues for scheduling to determine a packet forwarding sequence.
                    Buffer threshold
                    The buffer threshold measures the buffer usage of a queue, and includes the
                    lower and upper buffer thresholds. On a device enabled with queue-based
                    congestion monitoring, if the buffer usage of a queue reaches the upper buffer
                    threshold and then falls below the lower buffer threshold, the device saves
                    obtained historical congestion monitoring information so that you can check the
                    buffer usage of each queue.

11.7.2 Configuring Congestion Monitoring
Context
                    On an enterprise network, congestion mainly affects network performance,
                    leading to transmission delay and signal loss. Performing congestion monitoring
                    on each interface queue on a network device helps you to learn the buffer usage
                    of queues. It also provides information about packets that cause congestion,
                    guiding planning and adjusting network traffic.

Procedure
         Step 1 Enter the system view.
                    system-view


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                163
QoS Configuration
QoS Configuration                                                          11 Congestion Management Configuration


         Step 2 Enter the interface view.
                    interface { interface-type interface-number | interface-name }

         Step 3 Enable queue-based congestion monitoring.
                    qos [ queue queue-index ] buffer-monitoring enable

                          NOTE

