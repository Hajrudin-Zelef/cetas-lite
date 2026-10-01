---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-43
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5877, 6010]
sha256: a81f6a2cb1570d08555f417315671fe9e976f73dc371c86646c21e8ccd3ec634
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     Single-rate-single-bucket    Limits bandwidth.             Discards low-priority
                                                                                services, such as extranet
                                                                                HTTP traffic and excess
                                                                                traffic.

                     Single-rate-two-bucket       Limits bandwidth, allows      Reserves bandwidth for
                                                  certain traffic bursts, and   important services or
                                                  distinguishes burst and       burst traffic (for
                                                  normal services.              example, email data).

                     Two-rate-two-bucket          Limits bandwidth,             Recommended for
                                                  allocates bandwidth, and      important services
                                                  determines whether the        because it better
                                                  bandwidth is less than        monitors burst traffic
                                                  the CIR or is in the range    and guides traffic
                                                  of the CIR and PIR.           analysis.




Color-Aware Mode
                    In color-aware mode, the color of an arriving packet affects the metering results
                    of the token bucket mechanism in the following ways:
                    ●   If the packet has been marked green, the metering mechanism is the same as
                        that in color-blind mode.
                    ●   If the packet has been marked yellow, the system marks the packet yellow if
                        it conforms to the limit or marks the packet red if it violates the limit. This
                        depends on the packet length and the number of tokens. In single-rate-single-
                        bucket mode, the packet is marked red directly.
                    ●   If the packet has been marked red, it is also marked red in the token bucket.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                               106
QoS Configuration                                         9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                         based Rate Limiting Configuration


9.2.2 Traffic Policing
                    Traffic policing controls the rate of traffic entering a network within a specified
                    range by metering traffic and taking punitive actions on excess traffic to conserve
                    network resources.

                    Figure 9-4 Traffic policing components




                    As shown in Figure 9-4, traffic policing involves the following components:
                    ●   Meter: uses the token bucket mechanism to measure network traffic and
                        sends the result to the marker.
                    ●   Marker: colors packets green, yellow, or red based on the metering result
                        received from the meter.
                    ●   Action: performs actions based on the color of the packet. The actions are
                        defined as follows:
                        –    Pass: forwards the packets that conform to the limit.
                        –    Re-mark + pass: changes the local priorities of packets that exceed the
                             limit and forwards the packets.
                        –    Discard: drops packets that exceed the limit.
                    If the rate of a packet stream exceeds the limit, the system lowers the priority of
                    the excessive packets in the stream before forwarding them or discards the
                    packets. By default, the system forwards green and yellow packets, and discards
                    red packets.

9.2.3 Traffic Shaping
                    Traffic shaping adjusts the rate of outgoing traffic to reduce traffic bursts so that
                    outgoing packets can be transmitted at a stable rate. Traffic shaping uses a buffer
                    and token buckets to control the traffic rate. When packets are sent at a high rate,
                    the system buffers packets and then sends them evenly under the control of the
                    token buckets.
                    Figure 9-5 shows an example of traffic shaping, using flow-based queue shaping
                    in single-rate-single-bucket mode.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                107
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration


                    Figure 9-5 Traffic shaping




                    The process of traffic shaping is as follows:

                    1.   When packets arrive, the system classifies packets and places them into
                         different queues.
                    2.   If a queue is not configured with traffic shaping, packets placed in this queue
                         are immediately sent. For the queues configured with traffic shaping, the
                         system proceeds to the next step.
                    3.   The system places tokens in the bucket at the specified rate (CIR):
                         –   If there are sufficient tokens in the bucket, the system sends the packets
                             and decreases the number of tokens accordingly.
                         –   If tokens in the bucket are insufficient for packet forwarding, the system
                             places the packets into the buffer queue. If the buffer queue is full, the
                             system discards the packets.
                    4.   When there are packets in the buffer queue, the system compares the number
                         of packets with the number of tokens in the token bucket. If there are
                         sufficient tokens, the system forwards packets until all the packets in the
                         buffer queue are sent.

9.2.4 Interface-based Rate Limiting
                    Interface-based rate limiting controls the total rate of all packets sent or received
                    on an interface.

                    It uses the token bucket mechanism to control traffic rates. If rate limiting is
                    configured on an interface, all packets passing through this interface must be
                    processed by the token bucket. If there are sufficient tokens in the token bucket
                    for packet forwarding, packets are sent out from the interface. If the number is
                    insufficient, packets are discarded or buffered.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                108
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration


                    Interface-based rate limiting can be configured in the inbound or outbound
                    direction. Figure 9-6 illustrates interface-based rate limiting in the outbound
                    direction using the single-rate-single-bucket mechanism.

                    Figure 9-6 Interface-based rate limiting




                    The process of interface-based rate limiting is as follows:

