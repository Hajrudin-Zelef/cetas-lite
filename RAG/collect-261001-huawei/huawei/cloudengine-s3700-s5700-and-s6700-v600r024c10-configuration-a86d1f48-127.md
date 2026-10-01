---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-127
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2019-12-10", "2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [16250, 16397]
sha256: b535560b9473b42cbdbe5b3bdd36ae07c827b8a242bc292a13b9e601c2f3a3e8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        298
Security Configuration
Security Configuration                                                                  15 Keychain Configuration

                  [DeviceA] keychain huawei
                  [DeviceA-keychain-huawei] key-id 1
                  [DeviceA-keychain-huawei-keyid-1] algorithm hmac-sha-256
                  [DeviceA-keychain-huawei-keyid-1] key-string cipher YsHsjx_202206
                  [DeviceA-keychain-huawei-keyid-1] send-time 12:00 2019-12-10 to 18:00 2019-12-10
                  [DeviceA-keychain-huawei-keyid-1] receive-time 12:00 2019-12-10 to 18:00 2019-12-10
                  [DeviceA-keychain-huawei-keyid-1] default send-key-id
                  [DeviceA-keychain-huawei-keyid-1] quit
                  [DeviceA-keychain-huawei] quit

                  # Configure DeviceB.
                  [DeviceB] keychain huawei
                  [DeviceB-keychain-huawei] key-id 1
                  [DeviceB-keychain-huawei-keyid-1] algorithm hmac-sha-256
                  [DeviceB-keychain-huawei-keyid-1] key-string cipher YsHsjx_202206
                  [DeviceB-keychain-huawei-keyid-1] send-time 12:00 2019-12-10 to 18:00 2019-12-10
                  [DeviceB-keychain-huawei-keyid-1] receive-time 12:00 2019-12-10 to 18:00 2019-12-10
                  [DeviceB-keychain-huawei-keyid-1] default send-key-id
                  [DeviceB-keychain-huawei-keyid-1] quit
                  [DeviceB-keychain-huawei] quit

                  # Configure DeviceC.
                  [DeviceC] keychain huawei
                  [DeviceC-keychain-huawei] key-id 1
                  [DeviceC-keychain-huawei-keyid-1] algorithm hmac-sha-256
                  [DeviceC-keychain-huawei-keyid-1] key-string cipher YsHsjx_202206
                  [DeviceC-keychain-huawei-keyid-1] send-time 12:00 2019-12-10 to 18:00 2019-12-10
                  [DeviceC-keychain-huawei-keyid-1] receive-time 12:00 2019-12-10 to 18:00 2019-12-10
                  [DeviceC-keychain-huawei-keyid-1] default send-key-id
                  [DeviceC-keychain-huawei-keyid-1] quit
                  [DeviceC-keychain-huawei] quit

         Step 4 Configure keychain authentication for IS-IS.
                  # Configure DeviceA.
                  [DeviceA] interface 10ge 1/0/1
                  [DeviceA-10GE1/0/1] isis authentication-mode keychain huawei
                  [DeviceA-10GE1/0/1] quit
                  [DeviceA] quit

                  # Configure DeviceB.
                  [DeviceB] interface 10ge 1/0/1
                  [DeviceB-10GE1/0/1] isis authentication-mode keychain huawei
                  [DeviceB-10GE1/0/1] quit
                  [DeviceB] interface 10ge 1/0/2
                  [DeviceB-10GE1/0/2] isis authentication-mode keychain huawei
                  [DeviceB-10GE1/0/2] quit
                  [DeviceB] quit

                  # Configure DeviceC.
                  [DeviceC] interface 10ge 1/0/2
                  [DeviceC-10GE1/0/2] isis authentication-mode keychain huawei
                  [DeviceC-10GE1/0/2] quit
                  [DeviceC] quit

                  ----End

Verifying the Configuration
                  Using DeviceA as an example, check whether keychain authentication is
                  successfully configured for IS-IS.
                  ●      Run the display keychain keychain-name command to check the key ID in
                         the Active state.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                    299
Security Configuration
Security Configuration                                                                                15 Keychain Configuration

                         <DeviceA> display keychain huawei
                          Keychain Information:
                          ----------------------
                          Keychain Name                 : huawei
                            Timer Mode                : Absolute
                            Receive Tolerance(min) : 10
                            Digest Length            : 32
                            Time Zone               : LMT
                            TCP Kind              : 254
                            TCP Algorithm IDs           :
                             HMAC-MD5                   :5
                             HMAC-SHA1-12                 :2
                             HMAC-SHA1-20                 :6
                             MD5                 :3
                             SHA1                :4
                             HMAC-SHA-256                 :7
                             SHA-256               :8
                             SM3                 :9
                             HMAC-SHA-384                 : 11
                             HMAC-SHA-512                 : 12
                          Number of Key ID              :1
                          Active Send Key ID            :1
                          Active Receive Key ID : 01
                          Default send Key ID           :1

                         Key ID Information:
                           SHA1               :4
                           HMAC-SHA-256            :7
                           SHA-256             :8
                           SM3               :9
                         Number of Key ID        :1
                         Active Send Key ID      :1
                         Active Receive Key ID : 01
                         Default send Key ID      :1

                         Key ID Information:
                         ----------------------
                         Key ID                  :1
                           Key string             : ******
                           Algorithm                : HMAC-SHA-256
                           SEND TIMER                  :
                            Start time            : 2019-12-10 12:00
                            End time               : 2019-12-10 18:00
                            Status              : Active
                           RECEIVE TIMER                 :
                            Start time            : 2019-12-10 12:00
                            End time               : 2019-12-10 18:00
                            Status              : Active
                  ●      Run the display isis lsdb verbose command to check the details of the IS-IS
                         link state database (LSDB).
                         <DeviceA> display isis lsdb verbose
                                        Database information for ISIS(1)
                                        -----------------------------------


                                            Level-1 Link State Database

                         LSPID                Seq Num Checksum HoldTime                    Length ATT/P/OL
                         -----------------------------------------------------------------------------
                         0000.0000.0001.00-00* 0x0000020a 0x94e6                409           68       0/0/0
                          SOURCE         0000.0000.0001.00
                          NLPID        IPV4
                          AREA ADDR 10
                          INTF ADDR 192.168.1.1
                          NBR ID        0000.0000.0002.01 COST: 10
                          IP-Internal 192.168.1.0        255.255.255.0 COST: 10

                         0000.0000.0002.00-00 0x00000219 0xfa60             431          95      0/0/0
                          SOURCE     0000.0000.0002.00


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              300
Security Configuration
Security Configuration                                                                         15 Keychain Configuration

