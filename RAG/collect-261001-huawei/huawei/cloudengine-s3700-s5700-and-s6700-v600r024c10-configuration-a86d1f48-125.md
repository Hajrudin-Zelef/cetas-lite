---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-125
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [15929, 16080]
sha256: cce77e9ef4e2c3969230b1e0a3b09a70cab088a09fa9faafcc08648f4cdf83a9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   periodic daily                                           send-time daily start-time to end-
                                                                            time
                   periodic weekly                                          send-time day { start-day to end-day
                                                                            | start-day &<1-7> }

                   periodic monthly                                         send-time date { start-date to end-
                                                                            date | start-date &<1-31> }
                   periodic yearly                                          send-time month { start-month to
                                                                            end-month | start-month &<1-12> }



         Step 7 Configure the accept lifetime of the key based on the configured keychain time
                mode, as described in Table 15-5.

                  The lifetime of a key depends on clock synchronization.



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            293
Security Configuration
Security Configuration                                                              15 Keychain Configuration


                  Table 15-5 Configuring the accept lifetime of a key

                   Keychain Time Mode                                Command

                   absolute                                          receive-time start-time start-date
                                                                     { duration { duration-value | infinite }
                                                                     | { to end-time end-date } }

                   periodic daily                                    receive-time daily start-time to end-
                                                                     time
                   periodic weekly                                   receive-time day { start-day to end-
                                                                     day | start-day &<1-7> }
                   periodic monthly                                  receive-time date { start-date to end-
                                                                     date | start-date &<1-31> }
                   periodic yearly                                   receive-time month { start-month to
                                                                     end-month | start-month &<1-12> }



         Step 8 (Optional) Configure the key as the default send key.
                  default send-key-id

                  Each keychain can have only one default send key.

         Step 9 (Optional) Configure the length of the digest after being encrypted using an
                authentication algorithm.
                  digest-length { hmac-sha1-20 | hmac-sha-256 | sha-256 } length

                  The default digest length after encryption varies depending on the authentication
                  algorithm:

                  ●      HMAC-SHA1-20: 20 bytes
                  ●      HMAC-SHA-256: 32 bytes
                  ●      SHA-256: 32 bytes

        Step 10 Exit the key view.
                  quit

        Step 11 Exit the keychain view.
                  quit

                  ----End

15.5.3 Applying a Keychain

Context
                  A keychain itself manages just the encryption and authentication keys, and only
                  takes effect when used in applications. Keychains can be used in applications
                  running various protocols, as described in Table 15-6.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 294
Security Configuration
Security Configuration                                                               15 Keychain Configuration


                  Table 15-6 Using keychains in applications

                   Trans      Applica      View           Applicati    Configuration Reference
                   port       tion                        on Scope
                   Layer
                   Proto
                   col

                   Non-       RIP          Interface      Interface    IP Routing Configuration > RIP
                   TCP                     view                        Configuration > Improving RIP
                                                                       Network Security > Configuring the
                                                                       Authentication Mode for RIP-2
                                                                       Packets

                              IS-          IS-IS          IS-IS area   IP Route Configuration > IS-IS
                              IS/IS-       view                        Configuration > Configuring IS-IS
                              ISv6                                     Authentication
                                           IS-IS          IS-IS
                                           view           routing      IP Route Configuration > IS-ISv6
                                                          domain       Configuration > Configuring IPv6 IS-IS
                                                                       Authentication
                                           Interface      Interface
                                           view

                              OSPF/        OSPF           OSPF         IP Route Configuration > OSPF
                              OSPFv3       area           area         Configuration > Configuring OSPF
                                           view                        Authentication

                                           Interface      Interface    IP Route Configuration > OSPFv3
                                           view                        Configuration > Configuring OSPFv3
                                                                       Authentication
                                           OSPF           Virtual
                                           area           link
                                           view

                   TCP          BGP/       BGP view       Peer or      IP Routing Configuration > BGP
                                BGP4+      and            peer         Configuration > Configuring BGP
                                           related        group        Authentication > Configuring
                                           views                       Keychain Authentication
                                                                       IP Route Configuration > BGP4+
                                                                       Configuration > Configuring BGP4+
                                                                       Authentication




                  The following uses RIP as an example of how to apply a keychain.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface interface-type interface-number

         Step 3 Change the working mode of the interface from Layer 2 to Layer 3.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  295
Security Configuration
Security Configuration                                                                 15 Keychain Configuration

                  undo portswitch

                  Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                  S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                  Layer 2 mode to Layer 3 mode using the undo portswitch command.

