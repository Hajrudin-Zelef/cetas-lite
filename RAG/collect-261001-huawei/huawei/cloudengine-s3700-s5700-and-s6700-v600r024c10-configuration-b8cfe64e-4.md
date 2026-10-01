---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-4
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [200, 291]
sha256: 3b1c86522009fb3ae3f7d60d855ace148142607906033e93ce541527d35d1c0a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

11.7.2 Configuring Congestion Monitoring...................................................................................................................... 163
11.8 Maintaining Congestion Management.....................................................................................................................164
11.9 Example for Configuring Congestion Management............................................................................................ 164
11.10 Example for Configuring Congestion Avoidance and Congestion Management (PQ+WDRR
Scheduling and WRED Profile)..............................................................................................................................................168
11.11 Example for Configuring Congestion Monitoring.............................................................................................. 172
11.12 Microburst Detection.................................................................................................................................................... 174
11.12.1 Configuring Microburst Detection........................................................................................................................ 174
11.12.2 Verifying the Configuration.................................................................................................................................... 175

12 MPLS QoS Configuration.................................................................................................177
12.1 Overview of MPLS QoS.................................................................................................................................................. 177
12.2 Understanding MPLS QoS.............................................................................................................................................178
12.3 Configuration Precautions for MPLS QoS................................................................................................................184
12.4 Default Settings for MPLS QoS................................................................................................................................... 184
12.5 Configuring a DiffServ Mode....................................................................................................................................... 186
12.6 Configuring Priority Mapping...................................................................................................................................... 188
12.6.1 Configuring a DiffServ Domain............................................................................................................................... 188
12.6.2 Applying a DiffServ Domain..................................................................................................................................... 189
12.7 Checking the MPLS QoS Configuration.................................................................................................................... 190
12.8 Example for Configuring MPLS QoS..........................................................................................................................190

13 Experience Assurance Configuration............................................................................200
13.1 Overview of Experience Assurance............................................................................................................................ 200
13.2 Understanding Experience Assurance....................................................................................................................... 200
13.3 Configuration Precautions for Experience Assurance.......................................................................................... 203
13.4 Default Settings for Experience Assurance..............................................................................................................276
13.5 Configuring Experience Assurance............................................................................................................................. 276
13.5.1 Updating a Service Awareness Signature Database......................................................................................... 276
13.5.2 Configuring MQC-based Experience Assurance................................................................................................. 279
13.5.3 (Optional) Configuring Global Parameters for Experience Assurance.......................................................280
13.5.4 Verifying the Configuration....................................................................................................................................... 283
13.5.5 Example for Configuring MQC-based Experience Assurance........................................................................ 284
13.5.6 Example for Configuring Experience Assurance for Traffic Statistics Collection.....................................287




Issue 01 (2025-03-03)                                Copyright © Huawei Technologies Co., Ltd.                                                                                  v
QoS Configuration
QoS Configuration                                                                1 About This Document




                                            1         About This Document

Intended Audience
                    This document is intended for network engineers responsible for switch
                    management and maintenance. You should be familiar with basic Ethernet
                    knowledge and have extensive network management experience. In addition, you
                    should understand your network well, including the network topology and
                    deployed network services.

Symbol Conventions
                    The symbols used in this document are described in the following table. They are
                    defined as follows.

                     Symbol                                   Description

                                                              Indicates a hazard with a high level of
                                                              risk which, if not avoided, will result in
                                                              death or serious injury.

                                                              Indicates a hazard with a medium
                                                              level of risk which, if not avoided,
                                                              could result in death or serious injury.

                                                              Indicates a hazard with a low level of
                                                              risk which, if not avoided, could result
                                                              in minor or moderate injury.

                                                              Indicates a potentially hazardous
                                                              situation which, if not avoided, could
                                                              result in equipment damage, data loss,
                                                              performance deterioration, or
                                                              unanticipated results.
                                                              NOTICE is used to address practices
                                                              not related to personal injury.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  1
QoS Configuration
QoS Configuration                                                                     1 About This Document


                     Symbol                                          Description

