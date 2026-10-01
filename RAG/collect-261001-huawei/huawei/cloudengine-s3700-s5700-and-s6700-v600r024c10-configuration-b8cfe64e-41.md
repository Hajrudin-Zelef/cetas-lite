---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-41
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5646, 5760]
sha256: 7afdb6027020689c2ace2cb99a7989ee060fe51b1fa7dab0e57d80bbdb3fd884
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     Packet     Time       Packet    Delay      Token        Numbe      Numbe      Markin
                     No.        (ms)       Length    (ms)       Additio      r of       r of       g
                                           (Bytes)              n            Tokens     Tokens
                                                                (Bytes)      Before     After
                                                                             Packet     Packet
                                                                             Process    Process
                                                                             ing        ing
                                                                             (Bytes)    (Bytes)

                     -          -          -         -          -            2000       2000       -

                     1          0          1500      0          0            2000       500        Green

                     2          1          1500      1          125          625        625        Red

                     3          2          1000      1          125          750        750        Red

                     4          22         1500      20         2500         2000       500        Green




Two-Rate-Two-Bucket Mechanism
                    The two-rate-two-bucket mechanism uses the trTCM algorithm defined in RFC
                    2698 to measure traffic and marks packets green, yellow, or red based on the
                    metering result.

                    As shown in Figure 9-3, buckets P and C contain Tp and Tc tokens respectively.
                    The two-rate-two-bucket mechanism uses four parameters:
                    ●    Peak information rate (PIR): indicates the rate at which tokens are placed into
                         bucket P, that is, the maximum traffic rate that bucket P allows. The PIR is
                         greater than the CIR.
                    ●    CIR: indicates the rate at which tokens are placed into bucket C, that is, the
                         average traffic rate that bucket C allows.
                    ●    Peak burst size (PBS): indicates the capacity of bucket P, that is, the maximum
                         volume of burst traffic that bucket P allows.
                    ●    CBS: indicates the capacity of bucket C, that is, the maximum volume of burst
                         traffic that bucket C allows.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                102
QoS Configuration                                               9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                               based Rate Limiting Configuration


                    The system places tokens into bucket P at the PIR and places tokens into bucket C
                    at the CIR:
                    ●    If Tp is less than the PBS, Tp increases. If Tp is greater than or equal to the
                         PBS, Tp remains unchanged.
                    ●    If Tc is less than the CBS, Tc increases. If Tc is greater than or equal to the
                         CBS, Tp remains unchanged.
                    B indicates the size of an arriving packet:
                    ●    If B is greater than Tp, the packet is marked red.
                    ●    If B is greater than Tc and less than or equal to Tp, the packet is marked
                         yellow and Tp decreases by B.
                    ●    If B is less than or equal to Tc, the packet is marked green, and Tp and Tc
                         decrease by B.
                    The two-rate-two-bucket mechanism enables burst traffic rates. When the traffic
                    rate is lower than the CIR, packets are marked green. When the rate is higher than
                    the CIR but less than the PIR, packets are marked yellow. When the rate is higher
                    than the PIR, packets are marked red.

                    Figure 9-3 Two-rate-two-bucket mechanism




                    This example uses the CIR of 1 Mbit/s, the PIR of 2 Mbit/s, the CBS of 2000 bytes,
                    and the PBS of 3000 bytes. Buckets C and P are initially full of tokens. In two-rate-
                    two-bucket mode, the token buckets process packets as follows:

                          NOTE

                    Here, 1 Mbit/s is equal to 1 x 106 bit/s.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   103
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration


                    ●    If the first packet arriving at the interface is 1500 bytes long, the packet is
                         marked green because the numbers of tokens in both buckets P and C are
                         greater than the packet length. The numbers of tokens in both buckets P and
                         C then decrease by 1500 bytes, with 500 and 1500 bytes remaining in bucket
                         C and bucket P, respectively.
                    ●    Assume that the second packet arriving at the interface after a delay of 1 ms
                         is 1800 bytes long. Additional 250-byte tokens are placed into bucket P (PIR x
                         time period = 2 Mbit/s x 1 ms = 2000 bits = 250 bytes). Bucket P now has
                         1750-byte tokens, and is smaller than the packet length. Additional 125-byte
                         tokens are placed into bucket C (CIR x time period = 1 Mbit/s x 1 ms = 1000
                         bits = 125 bytes). Bucket C now has 625-byte tokens. Therefore, the second
                         packet is marked red, and the numbers of tokens in buckets P and C remain
                         unchanged.
                    ●    Assume that the third packet arriving at the interface after a delay of 1 ms is
                         1000 bytes long. Additional 250-byte tokens are placed into bucket P. Bucket
                         P now has 2000-byte tokens, and is larger than the packet length. Additional
                         125-byte tokens are placed into bucket C. Bucket C now has 750-byte tokens,
                         and is still smaller than the packet length. The packet is marked yellow. The
                         number of tokens in bucket P decreases by 1000 bytes, with 1000 bytes
                         remaining. The number of tokens in bucket C remains unchanged.
                    ●    Assume that the fourth packet arriving at the interface after a delay of 20 ms
                         is 1500 bytes long. Additional 5000-byte tokens are placed into bucket P (PIR
                         x time period = 2 Mbit/s x 20 ms = 40000 bits = 5000 bytes), but tokens that
                         exceed the PBS (3000 bytes) are dropped. Bucket P has 3000-byte tokens,
                         which are sufficient for the 1500-byte fourth packet. Additional 2500-byte
                         tokens are placed into bucket C (CIR x time period = 1 Mbit/s x 20 ms =
                         20000 bits = 2500 bytes), but tokens that exceed the CBS (2000 bytes) are
                         dropped. Bucket C then has 2000-byte tokens, which are sufficient for the
                         1500-byte fourth packet. Therefore, the fourth packet is marked green. The
                         number of tokens in bucket P decreases by 1500 bytes, with 1500 bytes
                         remaining. The number of tokens in bucket C decreases by 1500 bytes, with
                         500 bytes remaining.

