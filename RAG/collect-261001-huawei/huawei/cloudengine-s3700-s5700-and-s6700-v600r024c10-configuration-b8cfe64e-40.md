---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-40
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5534, 5645]
sha256: 8b01de505593e79d3a5c05fa26e6738db65677c47f31e34702b19afe3705e711
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                         insufficient; therefore, the third packet is marked red. The numbers of tokens
                         in buckets C and E remain unchanged.
                    ●    Assuming that the fourth packet arriving at the interface after a delay of 20
                         ms is 1500 bytes long, as with the preceding packets, additional 2500-byte
                         tokens are placed into bucket C (CIR x time period = 1 Mbit/s x 20 ms =
                         20000 bits = 2500 bytes). Bucket C now has 3250-byte tokens. The 1250-byte
                         tokens exceeding the CBS (2000 bytes) are placed into bucket E. Now, bucket
                         E has 1750-byte tokens. The packet is marked green because the number of
                         tokens in bucket C is greater than the packet length. The number of tokens in
                         bucket C decreases by 1500 bytes, with 500 bytes remaining. The number of
                         tokens in bucket E remains unchanged.
                    Table 9-1 describes packet processing.

                    Table 9-1 Packet processing in single-rate-two-bucket mode
                                                                  Number of          Number of
                                        Pack             Toke     Tokens             Tokens After
                                        et               n        Before             Packet
                     Pack                                         Packet             Processing
                              Time      Lengt   Delay    Addit                                         Mark
                     et                                           Processing         (Bytes)
                              (ms)      h       (ms)     ion                                           ing
                     No.                                          (Bytes)
                                        (Byte            (Byte
                                        s)               s)       Buck      Buck     Buck     Buck
                                                                  et C      et E     et C     et E

                     -        -         -       -        -        2000      2000     2000     2000     -

                     1        0         1500    0        0        2000      2000     500      2000     Green

                     2        1         1500    1        125      625       2000     625      500      Yello
                                                                                                       w

                     3        2         1000    1        125      750       500      750      500      Red

                     4        22        1500    20       2500     2000      1750     500      1750     Green




Single-Rate-Single-Bucket Mechanism
                    If burst traffic is not allowed, the EBS must be set to 0 in single-rate-two-bucket
                    mode. In this case, only one token bucket is used because there are always 0
                    tokens in bucket E.
                    As shown in Figure 9-2, bucket C contains Tc tokens. The single-rate-single-bucket
                    mechanism uses two parameters:
                    ●    CIR: indicates the rate at which tokens are placed into bucket C, that is, the
                         average traffic rate that bucket C allows.
                    ●    CBS: indicates the capacity of bucket C, that is, the maximum volume of burst
                         traffic that bucket C allows.
                    The system places tokens into the bucket at the CIR. If Tc is less than the CBS, Tc
                    increases. If Tc is greater than or equal to the CBS, Tc remains unchanged.
                    B indicates the size of an arriving packet:

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                100
QoS Configuration                                               9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                               based Rate Limiting Configuration


                    ●    If B is less than or equal to Tc, the packet is marked green, and Tc decreases
                         by B.
                    ●    If B is greater than Tc, the packet is marked red, and Tc remains unchanged.
                    The single-rate-single-bucket mechanism does not allow burst traffic. When the
                    traffic rate is lower than the CIR, packets are marked green. When the rate is
                    higher than the CIR, packets are marked red.

                    Figure 9-2 Single-rate-single-bucket mechanism




                    This example uses the CIR of 1 Mbit/s and the CBS of 2000 bytes. Bucket C is
                    initially full of tokens. In single-rate-single-bucket mode, the token bucket
                    processes packets as follows:

                          NOTE

                    Here, 1 Mbit/s is equal to 1 x 106 bit/s.
                    ●    Assuming that the first packet arriving at the interface is 1500 bytes long, the
                         packet is marked green because the number of tokens in bucket C is greater
                         than the packet length. The number of tokens in bucket C then decreases by
                         1500 bytes, with 500 bytes remaining.
                    ●    Assuming that the second packet arriving at the interface after a delay of 1
                         ms is 1500 bytes long, additional 125-byte tokens are placed into bucket C
                         (CIR x time period = 1 Mbit/s x 1 ms = 1000 bits = 125 bytes). Bucket C now
                         has 625-byte tokens. Tokens in bucket C are insufficient, so the second packet
                         is marked red.
                    ●    Assuming that the third packet arriving at the interface after a delay of 1 ms
                         is 1000 bytes long, additional 125-byte tokens are placed into bucket C.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   101
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration


                         Bucket C now has 750-byte tokens. Tokens in bucket C are insufficient, so the
                         third packet is marked red.
                    ●    Assuming that the fourth packet arriving at the interface after a delay of 20
                         ms is 1500 bytes long, additional 2500-byte tokens are placed into bucket C
                         (CIR x time period = 1 Mbit/s x 20 ms = 20000 bits = 2500 bytes). Bucket C
                         now has 3250-byte tokens. However, bucket C can have a maximum of 2000-
                         byte tokens. Therefore, the CBS is 2000 bytes. The packet is marked green
                         because the number of tokens in bucket C is greater than the packet length.
                         The number of tokens in bucket C decreases by 1500 bytes, with 500 bytes
                         remaining.

                    Table 9-2 describes packet processing.

                    Table 9-2 Packet processing in single-rate-single-bucket mode

