---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-123
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [15612, 15775]
sha256: 88e4e5f70138bab7ba7b00014ad3b6db030a0d20d67d02e102d584e083f37fb2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Security Configuration
Security Configuration                                                     15 Keychain Configuration


                   Key             Authenticati     Authenticati     Lifetime        Key Set
                                   on Algorithm     on Key String

                   Key4            HMAC-            HeBgDfCa         Every           KeychainB
                                   SHA-256                           Tuesday,
                                                                     Thursday, and
                                                                     Saturday

                   Key5            HMAC-            DhAgBfCe         The third day   KeychainC
                                   SHA-256                           to the eighth
                                                                     day of every
                                                                     month

                   Key6            HMAC-            EaHgBcFd         June to         KeychainD
                                   SHA-256                           September
                                                                     every year




15.2.2 Keychain Fundamentals (for Non-TCP Applications)
                  A keychain itself manages just the encryption and authentication keys, and only
                  takes effect when used in applications.

                  An application needs to bind a keychain before using keychain authentication. For
                  example, if an application binds KeychainA, the application can use the keys in
                  KeychainA for encryption and decryption.


Encryption Process

                  Figure 15-1 Encryption process for a non-TCP application using keychain
                  authentication




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                          287
Security Configuration
Security Configuration                                                                  15 Keychain Configuration


                          NOTE

                         If an application does not obtain an active send key from the keychain, the application
                         cannot use keychain authentication when sending packets. That is, packets are sent without
                         being encrypted.


Decryption Process

                  Figure 15-2 Decryption process for a non-TCP application using keychain
                  authentication




                          NOTE

                         ● The purpose of the decryption process is not to decrypt packets; rather, the device re-
                           encrypts packets and checks whether the new encryption result is the same as the
                           received old decryption result. If so, the decryption is successful.
                         ● Keychain authentication in IS-IS is special. During the decryption, an IS-IS application
                           does not provide a key ID. Instead, the keychain searches all active accept keys and uses
                           the one with the same algorithm for decryption.


15.2.3 Keychain Fundamentals (for TCP Applications)
                  The fundamentals of keychain authentication for TCP applications are similar to
                  those for non-TCP applications. The only difference is that the TCP Enhanced
                  Authentication Option is added for TCP applications.


TCP Enhanced Authentication Option
                  Figure 15-3 depicts the format of the TCP Enhanced Authentication Option. The
                  TCP packet header carries this option to provide authentication protection
                  specifically for TCP connections.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      288
Security Configuration
Security Configuration                                                           15 Keychain Configuration


                  Figure 15-3 Format of TCP Enhanced Authentication Option




                  ●      Kind: 8 bits, identifies the TCP Enhanced Authentication Option. This value is
                         assigned by IANA.
                  ●      Length: 8 bits, specifies the length of the TCP Enhanced Authentication
                         Option, in octets.
                  ●      T: 1 bit, specifies whether the TCP Enhanced Authentication Option is included
                         in the TCP header for the purpose of TCP enhanced authentication
                         calculation. A value of 0 indicates that the TCP Enhanced Authentication
                         Option is included. The default value is 0.
                  ●      K: 1 bit, reserved for future enhancement. The current value is 0.
                  ●      Alg ID: 6 bits, identifies the TCP enhanced authentication algorithm.
                  ●      Res: 2 bits, reserved for future use. The current value is 0.
                  ●      Key ID: 6 bits, identifies the key for keychain authentication.
                  ●      Authentication Data: The length is variable. It contains at least the result of
                         TCP enhanced authentication calculation.

                  IANA does not define the values of the Kind and Alg ID fields in a unified manner.
                  Therefore, different vendors use different values. To enable devices of different
                  vendors to communicate, the keychain supports configuration of the TCP Kind and
                  TCP Alg ID fields.

Encryption Process

                  Figure 15-4 Encryption process for a TCP application using keychain
                  authentication




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             289
Security Configuration
Security Configuration                                                          15 Keychain Configuration


Decryption Process

                  Figure 15-5 Decryption process for a TCP application using keychain
                  authentication




15.3 Configuration Precautions for Keychain

15.4 Default Settings for Keychains
                  Table 15-3 describes the default settings for keychains.

                  Table 15-3 Default settings for keychains
                   Parameter                                  Default Setting

                   Digest length after encryption using       HMAC-SHA-256: 32 bytes
                   an authentication algorithm                SHA-256: 32 bytes
                                                              HMAC-SHA1-20: 20 bytes

                   Acceptance tolerance                       0: no tolerance

                   Type value in the TCP Enhanced             254
                   Authentication Option: TCP Kind




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             290
Security Configuration
Security Configuration                                                                     15 Keychain Configuration


                   Parameter                                              Default Setting

