---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-34
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [4643, 4834]
sha256: c5f20b581c84889433f3ec7dd6fc973ec1fdb4602cd75a8b32907d719f5cbd25
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


                     Internal Priority (CoS)                    Queue Index

                     CS7                                        7




Default Priority Mappings for Outgoing Packets on an Interface
                    Table 8-5 lists the default mappings from internal priorities or drop priorities to
                    802.1p values of outgoing VLAN packets.


                    Table 8-5 Mappings from internal priorities or drop priorities to 802.1p values in
                    the outbound direction

                     Internal Priority (CoS)       Drop Priority (Color)       802.1p Value

                     BE                            Green                       0

                     BE                            Yellow                      0

                     BE                            Red                         0

                     AF1                           Green                       1

                     AF1                           Yellow                      1

                     AF1                           Red                         1

                     AF2                           Green                       2

                     AF2                           Yellow                      2

                     AF2                           Red                         2

                     AF3                           Green                       3

                     AF3                           Yellow                      3

                     AF3                           Red                         3

                     AF4                           Green                       4

                     AF4                           Yellow                      4

                     AF4                           Red                         4

                     EF                            Green                       5

                     EF                            Yellow                      5

                     EF                            Red                         5

                     CS6                           Green                       6

                     CS6                           Yellow                      6

                     CS6                           Red                         6

                     CS7                           Green                       7


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 83
QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


                     Internal Priority (CoS)       Drop Priority (Color)       802.1p Value

                     CS7                           Yellow                      7

                     CS7                           Red                         7




                    Table 8-6 lists the default mappings from internal priorities or drop priorities to
                    DSCP values of outgoing IP packets.


                    Table 8-6 Mappings from internal priorities or drop priorities to DSCP values of IP
                    packets in the outbound direction

                     Internal Priority (CoS)       Drop Priority (Color)       DSCP Value

                     BE                            Green                       0

                     BE                            Yellow                      0

                     BE                            Red                         0

                     AF1                           Green                       10

                     AF1                           Yellow                      12

                     AF1                           Red                         14

                     AF2                           Green                       18

                     AF2                           Yellow                      20

                     AF2                           Red                         22

                     AF3                           Green                       26

                     AF3                           Yellow                      28

                     AF3                           Red                         30

                     AF4                           Green                       34

                     AF4                           Yellow                      36

                     AF4                           Red                         38

                     EF                            Green                       46

                     EF                            Yellow                      46

                     EF                            Red                         46

                     CS6                           Green                       48

                     CS6                           Yellow                      48

                     CS6                           Red                         48

                     CS7                           Green                       56


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 84
QoS Configuration
QoS Configuration                                                                  8 Priority Mapping Configuration


                     Internal Priority (CoS)                   Drop Priority (Color)    DSCP Value

                     CS7                                       Yellow                   56

                     CS7                                       Red                      56




8.5 Configuring DiffServ Domain-based Priority
Mapping
Context
                    The process of configuring DiffServ Domain-based priority mapping is described as
                    follows:
                    1.    Specify the packet priority trusted by an interface so that the device performs
                          priority mapping based on the trusted priority.
                    2.    Configure a DiffServ domain to specify the mapping between external and
                          internal priorities. The device can then provide differentiated QoS services
                          based on internal priorities.
                    3.    Apply the DiffServ domain to an object for the mapping in the DiffServ
                          domain to take effect. By doing this, the device can re-mark priorities of
                          packets according to these mappings.
                    4.    Verify the configuration of the DiffServ domain.

8.5.1 Specifying the Packet Priority Trusted on an Interface
Context
                    The priority trusted on an interface determines the type of priority to be mapped
                    for packets on the interface.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface interface-type interface-number

         Step 3 Specify the priority to be mapped for packets.
                    trust { 8021p { inner | outer } | dscp }

                           NOTE

                         The Ethernet interface working in Layer 3 mode trusts DSCP values by default, and the
                         priority to be mapped for packets cannot be specified.
                         Layer 3 sub-interfaces do not support the trust 8021p inner command.
                         The inner parameter, that is, performing mapping for packets based on the 802.1p priority
                         in the inner VLAN tag, is not supported by the S5755-S series.

                    ----End

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                    85

