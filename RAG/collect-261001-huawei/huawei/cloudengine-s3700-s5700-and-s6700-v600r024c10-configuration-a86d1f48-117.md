---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-117
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [14706, 14835]
sha256: 5e82de25d3d1db88000356df53757e626226610ac94a5380ae58412696f4ed19
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  By default, the interval is 0 seconds, indicating that the SSH client does not send
                  keepalive packets.
                  If the interval for sending keepalive packets is set to 0 seconds, the configured
                  maximum number of keepalive packets does not take effect.
                  If the SSH client does not receive any data packets from the server within a
                  certain period, the client will keep sending keepalive packets to the server after
                  the period elapses, until the number of sent keepalive packets reaches the upper
                  limit. If the client does not receive any keepalive response packet from the server,
                  the client will disconnect from the server.
         Step 3 Set the maximum number of keepalive packets that can be sent by the SSH client.
                  ssh client keepalive-maxcount count

                  By default, the maximum number of keepalive packets that can be sent by an SSH
                  client is 3.
                  If the SSH client does not receive any data packets from the server within a
                  certain period, the client will keep sending keepalive packets to the server after
                  the period elapses, until the number of sent keepalive packets reaches the upper
                  limit. If the client does not receive any keepalive response packet from the server,
                  the client will disconnect from the server.
         Step 4 (Optional) Configure a key exchange algorithm list for the SSH client.
                  ssh client key-exchange { dh_group_exchange_sha256 | dh_group_exchange_sha1 | dh_group1_sha1 |
                  ecdh_sha2_nistp256 | ecdh_sha2_nistp384 | ecdh_sha2_nistp521 | sm2_kep | dh_group14_sha1 |
                  dh_group16_sha512 | curve25519_sha256 | sm2_sm3 } *


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    269
Security Configuration
Security Configuration                                                                             13 SSH Configuration


                  By default:
                  When a device starts with a configuration file that does not contain the ssh client
                  key-exchange command configuration, the SSH client uses the
                  dh_group_exchange_sha256, dh_group16_sha512, or curve25519_sha256 key
                  exchange algorithm.
                  When a device starts with factory settings, the SSH server uses the
                  dh_group_exchange_sha256, dh_group16_sha512, or curve25519_sha256 key
                  exchange algorithm.

                          NOTE

                         ● For security purposes, you are advised to use the curve25519_sha256,
                           ecdh_sha2_nistp521, ecdh_sha2_nistp384, and ecdh_sha2_nistp256 key exchange
                           algorithms.
                         ● Parameters dh_group_exchange_sha1, dh_group1_sha1, sm2_kep, and
                           dh_group14_sha1 in the command can be used only after the weak security algorithm/
                           protocol feature package is installed using the install feature-software WEAKEA
                           command.

         Step 5 (Optional) Configure the condition for renegotiating the SSH session key.
                  ssh client rekey { data-limit data-limit | max-packet max-packet | time minutes } *

                  To improve transmission security, the SSH client can initiate key renegotiation. If
                  renegotiation fails, the SSH connection is terminated. By default, the SSH client
                  triggers key renegotiation only when at least one of the following conditions is
                  met:
                  ●      The total data volume of packets transmitted using the current key reaches
                         1000 megabytes.
                  ●      The total number of sent and received packets reaches 2147483648.
                  ●      An SSH connection lasts for 60 minutes.
         Step 6 (Optional) Configure an encryption algorithm list for the SSH client.
                  ssh client cipher { des_cbc | 3des_cbc | aes128_cbc | aes256_cbc | aes128_ctr | aes256_ctr | arcfour128 |
                  arcfour256 | aes192_cbc | aes192_ctr | aes128_gcm | aes256_gcm | sm4_cbc | sm4_gcm | sm4_ctr } *

                  By default:
                  ●      When a device starts with no configuration, the SSH client supports the
                         following encryption algorithms: AES256_GCM, AES128_GCM, AES256_CTR,
                         AES192_CTR, and AES128_CTR.
                  ●      When a device functioning as an SSH client starts with a configuration file
                         which does not contain the ssh client cipher command configuration, the
                         SSH client supports the following encryption algorithms: AES128_CTR,
                         AES256_CTR, AES192_CTR, AES128_GCM, and AES256_GCM.

                          NOTE

                         ● For security purposes, you are advised to use the following encryption algorithms:
                           AES256_GCM, AES128_GCM, AES256_CTR, AES192_CTR, and AES128_CTR.
                         ● Parameters des_cbc, 3des_cbc, aes128_cbc, aes256_cbc, arcfour128, arcfour256,
                           aes192_cbc, and sm4_cbc in the command can be used only after the weak security
                           algorithm/protocol feature package is installed using the install feature-software
                           WEAKEA command.

         Step 7 (Optional) Configure an authentication algorithm list for the SSH client.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              270
Security Configuration
Security Configuration                                                                         13 SSH Configuration

                  ssh client hmac { md5 | md5_96 | sha1 | sha1_96 | sha2_256 | sha2_256_96 | sha2_512 | sm3 |
                  sha2_256_etm | sha2_512_etm } *

                  By default:
                  ●      When a device starts with factory settings, the SSH client supports the
                         following HMAC authentication algorithms: SHA2_512 and SHA2_256.
                  ●      When a device starts with a configuration file which does not contain the ssh
                         client hmac configuration, the SSH client supports the following HMAC
                         authentication algorithms: SHA2_512, SHA2_256_ETM, SHA2_512_ETM, and
                         SHA2_256.

                          NOTE

                         ● For security purposes, you are advised to use the following HMAC algorithms:
                           SHA2_256, SHA2_512_ETM, SHA2_256_ETM, and SHA2_512.
                         ● Parameters md5, md5_96, sha1, sha1_96, and sha2_256_96 in the command can be
                           used only after the weak security algorithm/protocol feature package is installed using
                           the install feature-software WEAKEA command.

         Step 8 (Optional) Configure the DSCP priority of SSH packets.
                  ssh client dscp value

                  By default, the DSCP priority of SSH packets is 48.

                  ----End

13.6.3 Applying SSH

Context
                  SSH is a security protocol, and an SSH policy takes effect only when it is
                  associated with an application.


Procedure
         Step 1 Apply an SSH policy. Table 13-8 lists the main applications of an SSH policy when
                the device functions as an SSH client.

                  Table 13-8 Main applications of an SSH policy when the device functions as an
                  SSH client

