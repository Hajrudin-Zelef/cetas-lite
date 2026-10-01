---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-33
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [4469, 4642]
sha256: ed625bf93a103ea89fd06b092135f33598458c23b913cb528420820d6e447d70
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     DiffServ domain applied to an                   DiffServ domain default
                     interface

                     Interface priority                              0




                    In addition, the device defines default packet priority mappings, which are
                    referenced during priority mapping.


Default Priority Mappings for Incoming Packets on an Interface
                    Table 8-2 lists the default mappings from 802.1p values of VLAN packets to
                    internal priorities or drop priorities.

                         NOTE

                    The mappings of interface priorities to PHBs and colors are similar to the mappings of 802.1p
                    values to PHBs and colors.


                    Table 8-2 Mapping between 802.1p values and internal priorities or drop priorities
                    in the inbound direction on the device

                     802.1p Value                      Internal Priority (CoS)        Drop Priority (Color)

                     0                                 BE                             Green

                     1                                 AF1                            Green

                     2                                 AF2                            Green

                     3                                 AF3                            Green

                     4                                 AF4                            Green


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         80
QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


                     802.1p Value                 Internal Priority (CoS)     Drop Priority (Color)

                     5                            EF                          Green

                     6                            CS6                         Green

                     7                            CS7                         Green




                    Table 8-3 lists the default mappings from DSCP values in IP packets to internal
                    priorities or drop priorities.

                    Table 8-3 Mappings from DSCP values in IP packets to internal priorities or drop
                    priorities in the inbound direction

                     DSCP         Internal Priority    Drop      DSCP       Internal Priority   Drop
                     Value        (CoS)                Priorit   Value      (CoS)               Priorit
                                                       y                                        y
                                                       (Color                                   (Color
                                                       )                                        )

                     0            BE                   Green     32         AF4                 Green

                     1            BE                   Green     33         AF4                 Green

                     2            BE                   Green     34         AF4                 Green

                     3            BE                   Green     35         AF4                 Green

                     4            BE                   Green     36         AF4                 Yellow

                     5            BE                   Green     37         AF4                 Green

                     6            BE                   Green     38         AF4                 Red

                     7            BE                   Green     39         AF4                 Green

                     8            AF1                  Green     40         EF                  Green

                     9            AF1                  Green     41         EF                  Green

                     10           AF1                  Green     42         EF                  Green

                     11           AF1                  Green     43         EF                  Green

                     12           AF1                  Yellow    44         EF                  Green

                     13           AF1                  Green     45         EF                  Green

                     14           AF1                  Red       46         EF                  Green

                     15           AF1                  Green     47         EF                  Green

                     16           AF2                  Green     48         CS6                 Green

                     17           AF2                  Green     49         CS6                 Green

                     18           AF2                  Green     50         CS6                 Green


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              81
QoS Configuration
QoS Configuration                                                      8 Priority Mapping Configuration


                     DSCP        Internal Priority   Drop      DSCP         Internal Priority   Drop
                     Value       (CoS)               Priorit   Value        (CoS)               Priorit
                                                     y                                          y
                                                     (Color                                     (Color
                                                     )                                          )

                     19          AF2                 Green     51           CS6                 Green

                     20          AF2                 Yellow    52           CS6                 Green

                     21          AF2                 Green     53           CS6                 Green

                     22          AF2                 Red       54           CS6                 Green

                     23          AF2                 Green     55           CS6                 Green

                     24          AF3                 Green     56           CS7                 Green

                     25          AF3                 Green     57           CS7                 Green

                     26          AF3                 Green     58           CS7                 Green

                     27          AF3                 Green     59           CS7                 Green

                     28          AF3                 Yellow    60           CS7                 Green

                     29          AF3                 Green     61           CS7                 Green

                     30          AF3                 Red       62           CS7                 Green

                     31          AF3                 Green     63           CS7                 Green




Mappings Between Internal Priorities and Inbound Queue Indexes
                    The internal priority on a device determines the queue from which packets are
                    forwarded. Table 8-4 describes the mappings between internal priorities and
                    queues.

                    Table 8-4 Mappings between internal priorities and queue indexes
                     Internal Priority (CoS)                   Queue Index

                     BE                                        0

                     AF1                                       1

                     AF2                                       2

                     AF3                                       3

                     AF4                                       4

                     EF                                        5

                     CS6                                       6




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             82

