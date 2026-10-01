---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-124
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [15776, 15928]
sha256: adb7b2af1c56c16e7fc3a0f86346d118479208993f581d24e258eea76237fbc9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   TCP authentication algorithm ID: TCP                   HMAC-SHA1-12: 2
                   algorithm-id                                           MD5: 3
                                                                          SHA-1: 4
                                                                          HMAC-MD5: 5
                                                                          HMAC-SHA1-20: 6
                                                                          HMAC-SHA-256: 7
                                                                          SHA-256: 8
                                                                          SM3: 9
                                                                          HMAC-SHA-384: 11
                                                                          HMAC-SHA-512: 12
                                                                          HMAC-SM3: 13
                                                                          NOTE
                                                                           MD5, HMAC-MD5, and SHA-1 algorithms
                                                                           are not recommended since they are less
                                                                           secure.

                   Time format                                            Local Mean Time (LMT)




15.5 Configuring a Keychain

15.5.1 Creating a Keychain
Prerequisites
                  You have configured NTP for time synchronization between the transmit end and
                  receive end.

Context
                  Before configuring a keychain, create one first. You can create one or more
                  keychains as required.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create a keychain and enter the keychain view.
                  keychain keychain-name mode { absolute | periodic { daily | weekly | monthly | yearly } }

                  When creating a keychain, the time mode is mandatory. After a keychain is
                  created, you can directly run keychain Keychain-name to enter the keychain view,
                  without specifying the time mode.
         Step 3 (Optional) Configure the acceptance tolerance of the keychain.
                  receive-tolerance { value | infinite | seconds secvalue }


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      291
Security Configuration
Security Configuration                                                                 15 Keychain Configuration


                  You are advised to set the acceptance tolerance time to prevent packet loss caused
                  by clock jitter.

                  The acceptance tolerance can be configured using either of the following methods:

                  ●      Set a specific time, in minutes or seconds. The maximum value is 14400
                         minutes (10 days) or 864000 seconds (10 days). By default, the acceptance
                         tolerance value is 0, which is no tolerance. You are advised to set a proper
                         acceptance tolerance to prevent packet loss caused by clock jitter.
                  ●      Set the infinite parameter so that an accept key is always valid.

         Step 4 (Optional) Set the time format of the keychain to LMT or UTC.
                  time mode { lmt | utc }

                  The default format is LMT.

         Step 5 To use a keychain in TCP applications, you also need to configure the TCP Kind
                and TCP Alg ID fields in the TCP Enhanced Authentication Option. Perform this
                step only when a keychain is used in TCP applications.
                  tcp-kind kind-value
                  tcp-algorithm-id { md5 | sha-1 | hmac-md5 | hmac-sha1-12 | hmac-sha1-20 | hmac-sha-256 | sha-256 |
                  sm3 | hmac-sha-384 | hmac-sha-512 | hmac-sm3 } algorithm-id

                          NOTE

                         MD5, HMAC-MD5, and SHA-1 algorithms are not recommended since they are less secure.

         Step 6 Exit the keychain view.
                  quit

                  ----End

15.5.2 Configuring a Key in a Keychain

Context
                  After a keychain is created, you need to create and configure one or more keys as
                  required for the keychain.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the view of a created keychain.
                  keychain keychain-name

         Step 3 Create a key and enter the key view.
                  key-id key-id

         Step 4 Configure the authentication algorithm of the key.
                  algorithm { md5 | sha-1 | hmac-md5 | hmac-sha1-12 | hmac-sha1-20 | hmac-sha-256 | sha-256 | sm3 |
                  hmac-sha-384 | hmac-sha-512 | hmac-sm3 }




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      292
Security Configuration
Security Configuration                                                                             15 Keychain Configuration


                          NOTE

                         MD5, HMAC-MD5, and SHA-1 algorithms are not recommended since they are less secure.
                         Parameters md5, sha-1, hmac-md5, hmac-sha1-12, and hmac-sha1-20 in this command
                         can be used only after the weak security algorithm/protocol feature package is installed.
                         For security purpose,you are not advised to use the weak security algorithm or weak
                         security protocols provided by this feature. If you need to use the weak security algorithm
                         or protocols, run the install feature-software WEAKEA command to install the weak
                         security algorithm or protocol feature package WEAKEA. By default, the device provides the
                         weak security algorithm or protocol feature package WEAKEA. For details about how to
                         install or uninstall the feature package, see "Upgrade Maintenance Configuration" in CLI
                         Configuration Guide > System Management Configuration.

         Step 5 Configure the authentication key string of the key.
                  key-string { plain-cipher-text | plain plain-text | cipher plain-cipher-text }

                          NOTE

                         It is recommended that the password complies with the password complexity rule: The
                         password is at least eight characters long and contains at least two of the following:
                         uppercase letters, lowercase letters, digits, and special characters (excluding question marks
                         and spaces).
                         For security purposes, the cipher mode is recommended to ensure that the configured key
                         string is displayed in ciphertext in the configuration file.

         Step 6 Configure the send lifetime of the key based on the configured keychain time
                mode, as described in Table 15-4.

                  The lifetime of a key depends on clock synchronization.


                  Table 15-4 Configuring the send lifetime of a key

                   Keychain Time Mode                                       Command

                   absolute                                                 send-time start-time start-date
                                                                            { duration { duration-value | infinite }
                                                                            | { to end-time end-date } }

