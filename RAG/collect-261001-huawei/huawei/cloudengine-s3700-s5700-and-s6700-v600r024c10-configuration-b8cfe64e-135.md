---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-135
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [14949, 15062]
sha256: 8298415f5c9fa0f415618da3ebf829566989075c25bd5c81b6155ec239723885
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                                                                                         L48T4XE-
                                                                                         A-V2,
                                                                                         S5735E-
                                                                                         L24ST4XE
                                                                                         -A-V2,
                                                                                         S5735E-
                                                                                         L24P4S-
                                                                                         A-V2,
                                                                                         S5735E-
                                                                                         L16LP2U
                                                                                         M2X-QA-
                                                                                         V2,
                                                                                         S5735E-
                                                                                         L8T4X-
                                                                                         QA-V2,
                                                                                         S5735E-
                                                                                         L8P4X-
                                                                                         QA-V2,
                                                                                         S5735E-
                                                                                         L24P4XE-
                                                                                         A-V2,
                                                                                         S5735E-
                                                                                         L48LP4S-
                                                                                         A-V2
                                                                                         S5755-
                                                                                         S48U8YZ,
                                                                                         S5755-
                                                                                         S24T8J8Y
                                                                                         Z, S5755-
                                                                                         S48U8Y,
                                                                                         S5755-
                                                                                         S24U8Y,
                                                                                         S5755-
                                                                                         S48T8YZ,
                                                                                         S5755-
                                                                                         S24P8J8Y
                                                                                         Z, S5755-
                                                                                         S48P8Y,
                                                                                         S5755-
                                                                                         S24T8Y,
                                                                                         S5755-
                                                                                         S24U8J8Y
                                                                                         Z, S5755-
                                                                                         S48T8Y,
                                                                                         S5755-
                                                                                         S24P8Y,
                                                                                         S5755-
                                                                                         S48P8YZ




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                           275
QoS Configuration
QoS Configuration                                                  13 Experience Assurance Configuration


                     Featu   Feature Requirements                                Series      Models
                     re

                                                                                             S6750-
                                                                                             S16X10Y2
                                                                                             CZ,
                                                                                             S6750-
                                                                                             S24T16X8
                                                                                             Y2CZ,
                                                                                             S6750-
                                                                                             S16X8YZ




13.4 Default Settings for Experience Assurance
                    Table 13-3 describes the default settings for experience assurance.

                    Table 13-3 Default settings for experience assurance

                     Parameter                                 Default Setting

                     Status of the SA function on an           Disabled
                     interface

                     Application identification statistics     Disabled
                     collection

                     Aging period of the application           By default, the aging period of the
                     identification flow table                 application identification flow table is
                                                               300 seconds.




13.5 Configuring Experience Assurance

13.5.1 Updating a Service Awareness Signature Database

Prerequisites
                    ●   For information about how to prepare for the update of a signature database,
                        see "Preparations for Signature Database Update" in CLI Configuration Guide
                        > System Management > Signature Database Update Configuration.
                    ●   Before updating the service awareness signature database using the online
                        update function, you need to configure the device to communicate with the
                        Huawei Security Center. For details, see "Configuring the Device to
                        Communicate with the Huawei Security Center" in CLI Configuration Guide >
                        System Management > Signature Database Update Configuration >
                        Configuring Online Signature Database Update.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            276
QoS Configuration
QoS Configuration                                                  13 Experience Assurance Configuration


