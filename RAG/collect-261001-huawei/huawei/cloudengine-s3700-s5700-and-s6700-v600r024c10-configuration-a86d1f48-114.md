---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-114
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [14301, 14419]
sha256: 835c651ea617076908251cb40ddf4a8b63ddd939428ed503400c59094ff6c2d4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Configure a privilege       local-user user-name        -
                   level for the local user.   privilege level level

                   Return to the system
                                               quit                        -
                   view.




                  Table 13-5 Configuring the local RSA, DSA, SM2, or ECC key for the SSH user
                              Step                     Command                   Description

                   Enter the system view.      system-view                 -




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            262
Security Configuration
Security Configuration                                                          13 SSH Configuration


                             Step                     Command                   Description

                                                                          By default, the
                                                                          authentication type of
                                                                          the SSH connection is
                                                                          AAA.
                                                                          When the authentication
                                                                          type is AAA, only the
                                                                          password authentication
                                                                          mode can be configured.
                                                                          If the public key
                   Configure an                                           authentication mode is
                                              ssh authorization-type
                   authentication type for                                used, perform either of
                                              default { aaa | root }
                   the SSH connection.                                    the following operations:
                                                                          ● Run this command
                                                                            with the
                                                                            authentication type
                                                                            set to root.
                                                                          ● In the AAA view,
                                                                            create a local user
                                                                            with the same name
                                                                            as the SSH user.

                                              rsa peer-public-key key-
                                              name [ encoding-type
                                              enc-type ]
                                              or
                                              dsa peer-public-key
                                              key-name encoding-          If SM2-SM3
                                              type enc-type               authentication is used,
                   Enter the RSA, SM2, DSA,                               you also need to run the
                                              or
                   or ECC public key view.                                sm2 peer-public-key
                                              ecc peer-public-key key-    command to enter the
                                              name [ encoding-type        SM2 public key view.
                                              enc-type ]
                                              or
                                              sm2 peer-public-key
                                              key-name [ encoding-
                                              type enc-type ]

                   Enter the public key
                                              public-key-code begin       -
                   editing view.




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                            263
Security Configuration
Security Configuration                                                           13 SSH Configuration


                              Step                     Command                    Description

                                                                           ● The public key must
                                                                             be a hexadecimal
                                                                             character string in the
                                                                             public key encoding
                                                                             format, and generated
                                                                             by SSH client
                                                                             software. For detailed
                                                                             operations, see the
                   Edit the public key.        hex-data                      help documentation
                                                                             for the SSH client
                                                                             software.
                                                                           ● You need to enter the
                                                                             RSA, DSA, SM2, or
                                                                             ECC public key on the
                                                                             device functioning as
                                                                             the SSH server.

                                                                           ● If hex-data is invalid,
                                                                             the key cannot be
                                                                             generated after you
                                                                             run this command.
                                                                           ● If the key specified by
                                                                             key-name has been
                   Exit the public key                                       deleted in another
                                               public-key-code end
                   editing view.                                             view, the system
                                                                             displays a message
                                                                             indicating that the key
                                                                             does not exist and
                                                                             directly returns to the
                                                                             system view when you
                                                                             run this command.

                   Return to the system
                   view from the public key    peer-public-key end         -
                   view.

