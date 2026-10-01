---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-48
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [6236, 6349]
sha256: d6c71d58fbffdbeb2b0c6f1cf47ac7a79c34221ba2f8c71fc948518ab315196e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   S5732-H44S4X6QZ-V2/S5732-                  ● Interfaces on the HSIC-X08S000
                   H44S4X6QZ-TV2                                card
                                                              ● Interfaces on the HSIC-Y08S000
                                                                card
                                                              ● GE Interfaces on the panel 1 to 44,
                                                                40GE interfaces 5 and 6, and 10GE
                                                                interfaces 1 to 4




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                              113
Security Configuration
Security Configuration                                                      8 MACsec Configuration


                   Product                                 Interface Supporting MACsec

                   S5732-H48UM4Y2CZ-V2/S5732-              ● Interfaces on the S7X08000 card
                   H48UM4Y2CZ-TV2                          ● Interfaces on the S7C02000 card
                                                           ● MultiGE interfaces 1 to 48 on the
                                                             panel

                   S6750-H36C                              100GE interfaces 1 to 16 and 25 to 36
                                                           on the panel.
                                                           100GE interfaces 1 to 4 are mutually
                                                           exclusive with interfaces 35 and 36,
                                                           100GE interfaces 5 to 8 are mutually
                                                           exclusive with interfaces 33 and 34,
                                                           100GE interfaces 9 to 12 are mutually
                                                           exclusive with interfaces 29 to 32, and
                                                           100GE interfaces 13 to 16 are mutually
                                                           exclusive with interfaces 25 to 28.
                                                           100GE interfaces 1 to 16 by default

                   S6750-H48Y8C                            25GE interfaces 1 to 32 and 100GE
                                                           interfaces 1 to 8
                                                           25GE interfaces 1 to 8 are mutually
                                                           exclusive with 100GE interfaces 7 and
                                                           8, 25GE interfaces 9 to 16 are mutually
                                                           exclusive with 100GE interfaces 5 and
                                                           6, 25GE interfaces 17 to 24 are
                                                           mutually exclusive with 100GE
                                                           interfaces 3 and 4, and 25GE interfaces
                                                           25 to 32 are mutually exclusive with
                                                           100GE interfaces 1 and 2.
                                                           25GE interfaces 1 to 24 and 100GE
                                                           interfaces 1 and 2 by default

                   S5755-H24HB2Y2CZ/S5755-                 ● Interfaces on the HSIC-X08S000
                   H24P4Y2CZ/S5755-H24T4Y2CZ/S5755-          card
                   H24U4Y2CZ/S5755-H24UM4Y2CZ/             ● Interfaces on the HSIC-Y08S000
                   S5755-H24UM4Y2CZ-T/S5755-                 card
                   H24UN4Y2CZ/S5755-H48P4Y2CZ/
                   S5755-H48T4Y2CZ/S5755-H48T4Y2CZ-        ● All interfaces on the panel
                   B/S5755-H48U4Y2CZ/S5755-
                   H48UN4Y2CZ/S5755-H48UM4Y2CZ/
                   S5755-H48UM4Y2CZ-T

                   S5755-H24N4Y-A/S5755-                   All interfaces on the panel
                   H24UTM4X4Y2C/S5755-
                   H24UTM4X4Y2C-T/S5755-H48N4Y-A/
                   S5755-H48UTM4X4Y2C/S5755-
                   H48UTM4X4Y2C-T




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                           114
Security Configuration
Security Configuration                                                      8 MACsec Configuration


                   Product                                 Interface Supporting MACsec

                   S6780-H                                 ● Interfaces on the XSIC-C10H000
                                                             card: In mode A, 10GE interfaces 1
                                                             and 2 and interfaces 9 and 10 in all
                                                             slots support MACsec. In mode B,
                                                             10GE interfaces 1 and 2 and
                                                             interfaces 9 and 10 in slot 1 support
                                                             MACsec, and all interfaces in slot 3
                                                             support MACsec.
                                                           ● Interfaces on the XSIC-C16H000
                                                             card: In mode A, 100GE interfaces 1
                                                             to 6 in all slots support MACsec. In
                                                             mode B, 100GE interfaces 1 to 6 in
                                                             slot 1 and all interfaces in slot 3
                                                             support MACsec.
                                                           ● Interfaces on the XSIC-D12B000
                                                             card: In mode A, 400GE interfaces 1
                                                             and 2 in all slots support MACsec.
                                                             In mode B, 400GE interfaces 1 and
                                                             2 in slot 1 and all interfaces in slot
                                                             3 support MACsec.
                                                           ● Interfaces on the XSIC-L16Q000
                                                             card: In mode A, 40GE interfaces 1
                                                             to 6 in all slots support MACsec. In
                                                             mode B, 40GE interfaces 1 to 6 in
                                                             slot 1 and all interfaces in slot 3
                                                             support MACsec.
                                                           ● Interfaces on the XSIC-Y26B000,
                                                             XSIC-X26B000, and XSIC-M26B000
                                                             cards: In mode A, 100GE interfaces
                                                             1 and 2 in all slots support MACsec.
                                                             In mode B, 100GE interfaces 1 and
                                                             2 in slot 1 and all interfaces in slot
                                                             3 support MACsec.

                   S5735-L16LP2UM2X-QA-V2/S5735-           16 x GE electrical interfaces and 2 x
                   L16P2UM2X-QA-V2                         MultiGE interfaces on the panel

