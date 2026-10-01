---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-39
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5427, 5533]
sha256: da13a7299d617c55aaf260869abe9a6e6b87e66820ff54923a755de3074b9bee
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

9.2.1 Traffic Metering and Token Bucket Mechanism
Overview
                    Traffic metering is the prerequisite for implementing traffic policing, traffic
                    shaping, and interface-based rate limiting. In traffic metering, network devices
                    determine whether the rate of incoming traffic exceeds a defined limit and take
                    appropriate actions. Generally, the token bucket mechanism is used to measure
                    traffic.
                    A token bucket is a container that stores a specific number of tokens. The system
                    places tokens into a token bucket at the configured rate. If the token bucket is full,
                    excess tokens overflow. The system determines whether the bucket has enough
                    tokens for packet forwarding. Otherwise, the traffic rate exceeds or violates the
                    rate limit.
                    RFC standards define the following token bucket algorithms:
                    ●   The single rate three color marker (srTCM) algorithm determines traffic bursts
                        based on the length of packets.
                    ●   The two rate three color marker (trTCM) algorithm determines traffic bursts
                        based on the rate of packets.
                    The token bucket algorithms mark packets red, yellow, or green based on the
                    result of traffic metering. The system then processes the packets according to their
                    marked color. The two aforementioned algorithms can work in color-aware and
                    color-blind modes. The color-blind mode is used as an example in the following
                    section.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 97
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration


Single-Rate-Two-Bucket Mechanism
                    The single-rate-two-bucket mechanism uses the srTCM algorithm defined in RFC
                    2697 to measure traffic and marks packets green, yellow, or red based on the
                    metering result.
                    As shown in Figure 9-1, bucket C (also called CIR token bucket) and bucket E
                    (also called EIR token bucket) contain Tc and Te tokens, respectively. The single-
                    rate-two-bucket mechanism uses three parameters:
                    ●   Committed information rate (CIR): indicates the rate at which tokens are
                        placed into bucket C, that is, the average traffic rate that bucket C allows.
                    ●   Committed burst size (CBS): indicates the capacity of bucket C, that is, the
                        maximum volume of burst traffic that bucket C allows.
                    ●   Excess burst size (EBS): indicates the capacity of bucket E, that is, the
                        maximum volume of excess burst traffic that bucket E allows.
                    The system places tokens into bucket C at the CIR:
                    ●   If Tc is less than the CBS, Tc increases.
                    ●   If Tc is equal to the CBS but Te is less than the EBS, Te increases.
                    ●   If Tc is equal to the CBS and Te is equal to the EBS, Tc and Te do not increase.
                    B indicates the size of an arriving packet:
                    ●   If B is less than or equal to Tc, the packet is marked green, and Tc decreases
                        by B.
                    ●   If B is greater than Tc and less than or equal to Te, the packet is marked
                        yellow and Te decreases by B.
                    ●   If B is greater than Te, the packet is marked red, and Tc and Te remain
                        unchanged.
                    The single-rate-two-bucket mechanism allows burst traffic. When the traffic rate is
                    lower than the CIR, packets are marked green. When the burst traffic volume is
                    greater than the CBS but lower than the EBS, packets are marked yellow. When
                    the volume is greater than the EBS, packets are marked red.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  98
QoS Configuration                                               9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                               based Rate Limiting Configuration


                    Figure 9-1 Single-rate-two-bucket mechanism




                    The preceding example uses the CIR of 1 Mbit/s and the CBS and EBS of 2000
                    bytes each. Buckets C and E are initially full of tokens. In single-rate-two-bucket
                    mode, the token buckets process packets as follows:

                          NOTE

                    Here, 1 Mbit/s is equal to 1 x 106 bit/s.
                    ●    Assuming that the first packet arriving at the interface is 1500 bytes long, the
                         packet is marked green because the number of tokens in bucket C is greater
                         than the packet length. The number of tokens in bucket C then decreases by
                         1500 bytes, with 500 bytes remaining. The number of tokens in bucket E
                         remains unchanged.
                    ●    Assuming that the second packet arriving at the interface after a delay of 1
                         ms is 1500 bytes long, additional 125-byte tokens are placed into bucket C
                         (CIR x time period = 1 Mbit/s x 1 ms = 1000 bits = 125 bytes). Bucket C now
                         has 625-byte tokens, which are insufficient for the 1500-byte second packet.
                         Bucket E has 2000-byte tokens, which are sufficient for the second packet.
                         Therefore, the second packet is marked yellow. Subsequently, the number of
                         tokens in bucket E decreases by 1500 bytes, with 500 bytes remaining. The
                         number of tokens in bucket C remains unchanged.
                    ●    Assuming that the third packet arriving at the interface after a delay of 1 ms
                         is 1000 bytes long, as with the second packet, additional 125-byte tokens are
                         placed into bucket C. Bucket C now has 750-byte tokens, which are
                         insufficient for the 1000-byte third packet. The tokens in bucket E are also


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    99
QoS Configuration                                            9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                            based Rate Limiting Configuration


