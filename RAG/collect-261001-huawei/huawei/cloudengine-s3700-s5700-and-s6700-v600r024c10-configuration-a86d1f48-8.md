---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-8
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [411, 534]
sha256: a835e179c9441845cb4fe2204500fe3cb8d64c74108e2ff181cfd749aef82a2b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

24 Trusted System Configuration....................................................................................... 336
24.1 Overview of Trusted Systems....................................................................................................................................... 336
24.2 Configuration Precautions for Trusted Systems.....................................................................................................337
24.3 Digital Signature of Software Packages................................................................................................................... 337
24.4 Secure Boot........................................................................................................................................................................ 338
24.5 Trusted Boot....................................................................................................................................................................... 339




Issue 01 (2025-03-03)                                   Copyright © Huawei Technologies Co., Ltd.                                                                                        ix
Security Configuration
Security Configuration                                                         1 About This Document




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




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                                  1
Security Configuration
Security Configuration                                                              1 About This Document


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

Security Conventions
                  ●      Password setting
                         –    Configuring a ciphertext password is recommended. For security
                              purposes, do not disable password complexity check, and change the
                              password periodically.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                              2
Security Configuration
Security Configuration                                                           1 About This Document


