---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-42
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5761, 5876]
sha256: 0db9ef0fb760ea005aca21ada94e51ff7e3f7e1e57c56e8e52a3572887309e8c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    Table 9-3 describes packet processing.

                    Table 9-3 Packet processing in two-rate-two-bucket mode

                                                                     Number of       Number of
                                     Pack                            Tokens          Tokens
                                                    Token
                                     et                              Before          After
                     Pack                   Dela    Addition
                             Time    Leng                            Packet          Packet           Mar
                     et                     y       (Bytes)
                             (ms)    th                              Processing      Processing       king
                     No.                    (ms)                     (Bytes)         (Bytes)
                                     (Byt
                                     es)            Buck     Buck    Buck    Buck    Buck     Buck
                                                    et C     et P    et C    et P    et C     et P

                     -       -       -      -       -        -       2000    3000    2000     3000    -

                     1       0       1500   0       0        0       2000    3000    500      1500    Gree
                                                                                                      n

                     2       1       1800   1       125      250     625     1750    625      1750    Red


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                104
QoS Configuration                                            9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                            based Rate Limiting Configuration


                                                                       Number of          Number of
                                     Pack                              Tokens             Tokens
                                                      Token
                                     et                                Before             After
                     Pack                     Dela    Addition
                             Time    Leng                              Packet             Packet        Mar
                     et                       y       (Bytes)
                             (ms)    th                                Processing         Processing    king
                     No.                      (ms)                     (Bytes)            (Bytes)
                                     (Byt
                                     es)              Buck    Buck     Buck    Buck       Buck   Buck
                                                      et C    et P     et C    et P       et C   et P

                     3       2       1000     1       125     250      750     2000       750    1000   Yello
                                                                                                        w

                     4       22      1500     20      2500    5000     2000    3000       500    1500   Gree
                                                                                                        n




Difference and Application of Three Token Bucket Modes
                    Table 9-4 describes the difference and relationship between the three token
                    bucket modes.

                    Table 9-4 Difference and relationship of three token bucket modes

                     Difference             Single-Rate-            Single-Rate-Two-       Two-Rate-Two-
                                            Single-Bucket           Bucket                 Bucket

                     Parameters             CIR and CBS             CIR, CBS, and EBS      CIR, CBS, PIR, and
                                                                                           PBS

                     Mode in which          Tokens are placed       When bucket C is       Tokens are placed
                     tokens are placed      into bucket C at        full, excess tokens    into bucket C at
                                            the CIR. Excess         are placed into        the CIR and into
                                            tokens are              bucket E. When         bucket P at the
                                            dropped when            buckets C and E        PIR. Buckets C
                                            bucket C is full.       are not full,          and P are
                                                                    tokens are placed      independent.
                                                                    only into bucket       Excess tokens are
                                                                    C.                     dropped when
                                                                                           buckets C and P
                                                                                           are full.

                     Burst traffic          Traffic burst is not    Traffic burst is       Traffic burst is
                                            allowed. Packet         allowed. Tokens in     allowed. When
                                            processing is           bucket C are used      buckets C and P
                                            implemented only        first. When tokens     have sufficient
                                            when bucket C           in bucket C are        tokens, tokens in
                                            has a sufficient        insufficient,          both buckets are
                                            number of tokens.       tokens in bucket E     used. When
                                                                    are used.              tokens in bucket C
                                                                                           are insufficient,
                                                                                           tokens in only
                                                                                           bucket P are used.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                 105
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration


                     Difference            Single-Rate-         Single-Rate-Two-        Two-Rate-Two-
                                           Single-Bucket        Bucket                  Bucket

                     Marking result        Green or red         Green, yellow, or       Green, yellow, or
                                                                red                     red

                     Relationship          In single-rate-two-bucket mode, if the EBS is 0, the effect is
                                           the same as that in single-rate-single-bucket mode.
                                           In two-rate-two-bucket mode, if the PIR is equal to the CIR,
                                           the effect is the same as that in single-rate-two-bucket
                                           mode.




                    Table 9-5 describes the functions and scenarios of the three token bucket modes.

                    Table 9-5 Functions and application scenarios of the three token bucket modes
                     Token Bucket Mode            Function                      Application Scenario

