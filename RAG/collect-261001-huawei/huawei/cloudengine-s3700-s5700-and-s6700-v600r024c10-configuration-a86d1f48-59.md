---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-59
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [7085, 7205]
sha256: 66a462e34e50b39e4ebf8ef2b23905d5566833159a5544ce6a04680d662dbb86
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   GE interfaces 1 to 24, 40GE interfaces 5 to 6,                S5755-S48P8Y,
                   and 10GE interfaces 1 to 4.                                   S5755-S24T8Y,
                   For the S6730-H28X6CZ-V2, S6730-H28X6CZ-                      S5755-S24U8J8YZ,
                   TV2:                                                          S5755-S48T8Y,
                                                                                 S5755-S24P8Y,
                   10GE interfaces 1 to 28, 40GE/100GE interfaces                S5755-S48P8YZ
                   5 to 6 and converted interfaces
                                                                                 S5735R-
                   For the S6750-H48Y8C:                                         L16LP2UM2X-QA-
                   25GE ports 1 to 32 and 100GE ports 1 to 8                     V2, S5735R-
                   support this function and are mutually                        L16LP2S-QA-V2,
                   exclusive. 25GE ports 1 to 8 are mutually                     S5735R-L16LP2X-
                   exclusive with 100GE ports 7 to 8, 25GE ports 9               QA-V2
                   to 16 are mutually exclusive with 100GE ports                 S6750E-
                   5 to 6, 25GE ports 17 to 24 are mutually                      S24T16X8Y2CZ,
                   exclusive with 100GE ports 3-4, and 25GE ports                S6750E-
                   25 to 32 are mutually exclusive with 100GE                    S16X10Y2CZ
                   ports 1-2. By default, 25GE ports 1 to 24 and
                   100GE ports 1 to 2 support this function.                     S5735-L14P2S-
                                                                                 QA-V2, S5735-
                   For the S5735R-L16LP2X-QA-V2, S5735R-                         L16LP2UM2X-QA-
                   L16LP2S-QA-V2, S5735-L16LP2X-QA-V2:                           V2, S5735-
                   16*GE electrical ports support MACsec.                        L16LP2X-QA-V2,
                   For the S5735I-S8T8P2S4XN-V2, S5735I-                         S5735-
                   S16T2S4XN-V2:                                                 L16P2UM2X-QA-
                                                                                 V2, S5735-
                   M0 mode: GE electrical ports 9 to 16, and GE                  L14P2S-TQA-V2
                   optical ports 17 and 18 M1 mode: GE electrical
                   ports 9 to 12, and 10GE optical ports 5 and 6                 S5735I-
                                                                                 S8T8P2S4XN-V2,
                   For the S5735R-L16LP2UM2X-QA-V2, S5735-                       S5735I-
                   L16LP2UM2X-QA-V2, S5735-L16P2UM2X-QA-                         S16T2S4XN-V2
                   V2:
                                                                                 S5732-
                   Supported by 16 x GE electrical ports and 2 x                 H48UM4Y2CZ-
                   MultiGE ports                                                 TV2, S5732-
                   For the S6750-H36C:                                           H24S4X6QZ-V2,
                   100G ports 1 to 16 and 25 to 36 support this                  S5732-
                   function and are mutually exclusive. 100GE                    H44S4X6QZ-V2,
                   ports 1 to 4 are mutually exclusive with ports                S5732-
                   35 to 36, 100GE ports 5 to 8 are mutually                     H24UM4Y2CZ-V2,
                   exclusive with ports 33 to 34, 100GE ports 9 to               S5732-
                   12 are mutually exclusive with ports 29 to 32,                H48UM4Y2CZ-V2,
                   and 100GE ports 13 to 16 are mutually                         S5732-
                   exclusive with ports 25 to 28, By default, 100G               H24S4X6QZ-TV2,
                   ports 1 to 16 support this function.                          S5732-
                                                                                 H44S4X6QZ-TV2,
                                                                                 S5732-
                                                                                 H24UM4Y2CZ-
                                                                                 TV2
                                                                                 S6730-H28X6CZ-
                                                                                 TV2, S6730-
                                                                                 H48X6CZ-TV2,


Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                          131
Security Configuration
Security Configuration                                                          8 MACsec Configuration


                   Feature Requirements                                Series       Models

                                                                                    S6730-H48X6CZ-
                                                                                    V2, S6730-
                                                                                    H28X6CZ-V2
                                                                                    S6750-
                                                                                    S16X10Y2CZ,
                                                                                    S6750-
                                                                                    S24T16X8Y2CZ,
                                                                                    S6750-S16X8YZ




8.4 Default Settings for MACsec

                  Table 8-4 Default settings for MACsec

                   Parameter                                  Default Setting

                   MACsec on interfaces                       Disabled

                   Key server priority                        16

                   CAK                                        Not configured

                   Encryption mode                            normal

                   Encryption offset                          0

                   Whether the MACsec frame header            SCI contained
                   contains the Secure Channel Identifier
                   (SCI)

                   SAK timeout period                         3600 seconds

                   Replay protection window size              0

                   MKA session timeout period                 6 seconds

                   MACsec capability value                    3




8.5 Enabling the MACsec Function

Context
                  To enable the MACsec function on a device, you need to create and configure a
                  MACsec profile, apply the profile to an interface, and configure a CAK.
                  After the MACsec function is enabled on interfaces of two connected devices, they
                  elect a key server based on their priorities, which can be configured in a MACsec
                  profile. A smaller priority value indicates a higher priority. The device with a higher

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             132
Security Configuration
Security Configuration                                                                  8 MACsec Configuration


