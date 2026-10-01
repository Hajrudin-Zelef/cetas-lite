---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-10
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [492, 611]
sha256: e1dd3ea01695abbf73e46ad9b6a8a7760d12bb253c5ecfa9c533f66ebacc3246
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

7.2 Understanding Tunnel Management...........................................................................................................................803
7.3 Configuration Precautions for Tunnel Management..............................................................................................804
7.4 Configuring a Tunnel Policy............................................................................................................................................ 804
7.4.1 Understanding Tunnel Policies................................................................................................................................... 805
7.4.2 Configuring a Tunnel Policy........................................................................................................................................ 806
7.4.3 Applying a Tunnel Policy.............................................................................................................................................. 808
7.4.4 Verifying the Configuration......................................................................................................................................... 809
7.5 Configuring a Tunnel Selector....................................................................................................................................... 810
7.5.1 Understanding Tunnel Selectors................................................................................................................................ 810
7.5.2 Configuring a Tunnel Selector.................................................................................................................................... 810
7.5.3 Applying a Tunnel Selector.......................................................................................................................................... 811
7.5.4 Verifying the Configuration......................................................................................................................................... 812
7.6 Maintaining Tunnel Management................................................................................................................................ 813
7.6.1 Monitoring the Running Status of Tunnels............................................................................................................ 813




Issue 01 (2025-03-03)                               Copyright © Huawei Technologies Co., Ltd.                                                                                xi
VPN Configuration
VPN Configuration                                                                1 About This Document




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
VPN Configuration
VPN Configuration                                                                     1 About This Document


                     Symbol                                          Description

                                                                     Supplements the important
                                                                     information in the main text.
                                                                     NOTE is used to address information
                                                                     not related to personal injury,
                                                                     equipment damage, and environment
                                                                     deterioration.




Command Conventions
                     Convention                        Description

                     Boldface                          The keywords of a command line are in boldfaces.

                     Italic                            Command arguments are in italic.

                     []                                Items (keywords or arguments) in square brackets
                                                       [ ] are optional.

                     { x | y | ... }                   Alternative items are grouped in braces and
                                                       separated by vertical bars. One is selected.

                     [ x | y | ... ]                   Optional alternative items are grouped in square
                                                       brackets and separated by vertical bars. One or none
                                                       is selected.

                     { x | y | ... } *                 Alternative items are grouped in braces and
                                                       separated by vertical bars. A minimum of one or a
                                                       maximum of all can be selected.

                     [ x | y | ... ] *                 Optional alternative items are grouped in square
                                                       brackets and separated by vertical bars. Many or
                                                       none can be selected.

                     &<1-n>                            This parameter before the & sign can be repeated 1
                                                       to n times.

                     #                                 This parameter before the # sign can be repeated 1
                                                       to n times.




Interface Numbering Conventions
                    Interface numbers used in this manual are examples. In device configuration, use
                    the existing interface numbers on devices.

