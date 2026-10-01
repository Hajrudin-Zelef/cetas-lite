---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-111
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [13916, 14032]
sha256: d7472df267533ae97d951f2fe08a6e92c01862bfab291a5cd23b544b7a87f96a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                              –    For security purposes, you are advised to use the curve25519_sha256,
                                   ecdh_sha2_nistp521, ecdh_sha2_nistp384, and ecdh_sha2_nistp256 key exchange
                                   algorithms.
                              –    Parameters dh_group_exchange_sha1, dh_group1_sha1, sm2_kep, and
                                   dh_group14_sha1 in the command can be used only after the weak security
                                   algorithm/protocol feature package is installed using the install feature-software
                                   WEAKEA command.
                  2.     Configure the condition for renegotiating the SSH session key.
                         ssh server rekey { data-limit data-limit | max-packet max-packet | time minutes } *

                         To improve transmission security, the SSH server can initiate key
                         renegotiation. If renegotiation fails, the SSH connection is terminated. By
                         default, the SSH server triggers key renegotiation only when at least one of
                         the following conditions is met:
                         –    The total data volume of packets transmitted using the current key
                              reaches 1000 megabytes.
                         –    The total number of sent and received packets reaches 2147483648.
                         –    An SSH connection lasts for 60 minutes.
                  3.     Set the minimum key length supported during diffie-hellman-group-exchange
                         key exchange between the SSH server and client.
                         ssh server dh-exchange min-len min-len

                         By default, the minimum supported key length is 3072 bits when the SSH
                         server uses the diffie-hellman-group-exchange key to exchange with clients.
                         You are advised to set the minimum key length to 3072 bits to improve
                         security.
                                  NOTE

                              Security risks exist if the minimum key length of the diffie-hellman-group-exchange
                              algorithm is less than or equal to 2048 bits. In this case, you need to run the install
                              feature-software WEAKEA command to install the weak security algorithm/protocol
                              feature package (WEAKEA). You are advised to set the minimum key length to 3072
                              bits. This command takes effect for both IPv4 and IPv6 SSH servers.
                  4.     Configure an encryption algorithm list for the SSH server.
                         ssh server cipher { des_cbc | 3des_cbc | aes128_cbc | aes256_cbc | aes128_ctr | aes256_ctr |
                         arcfour128 | arcfour256 | aes192_cbc | aes192_ctr | aes128_gcm | aes256_gcm | blowfish_cbc |
                         sm4_cbc | sm4_gcm | sm4_ctr } *

                         By default, an SSH server supports the following encryption algorithms:
                         AES128_CTR, AES256_CTR, AES192_CTR, AES128_GCM, and AES256_GCM.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                           256
Security Configuration
Security Configuration                                                                          13 SSH Configuration


                                  NOTE

                              –    For security purposes, you are advised to use the following encryption algorithms:
                                   AES256_GCM, AES128_GCM, AES256_CTR, AES192_CTR, and AES128_CTR.
                              –    Parameters blowfish_cbc, des_cbc, 3des_cbc, aes128_cbc, aes256_cbc,
                                   arcfour128, arcfour256, aes192_cbc, and sm4_cbc in the command can be used
                                   only after the weak security algorithm/protocol feature package is installed using
                                   the install feature-software WEAKEA command.
                  5.     Configure an authentication algorithm list for the SSH server.
                         ssh server hmac { md5 | md5_96 | sha1 | sha1_96 | sha2_256 | sha2_256_96 | sha2_512 | sm3 |
                         sha2_256_etm | sha2_512_etm } *

                         By default:
                         When a device starts with factory settings, the SSH server uses the HMAC
                         authentication algorithm SHA2_512 or SHA2_256.
                         When a device starts with a configuration file which does not contain the ssh
                         client hmac configuration, the SSH server uses the HMAC authentication
                         algorithm SHA2_512, SHA2_256_ETM, SHA2_512_ETM, or SHA2_256.
                                  NOTE

                              –    For security purposes, you are advised to use the following HMAC algorithms:
                                   SHA2_512, SHA2_512_ETM, SHA2_256_ETM, and SHA2_256.
                              –    Parameters md5, md5_96, sha1, sha1_96, and sha2_256_96 in the command can
                                   be used only after the weak security algorithm/protocol feature package is
                                   installed using the install feature-software WEAKEA command.
                  6.     Disable the function that triggers a warning when the SSH server uses an
                         insecure algorithm.
                         ssh server security-banner disable

                         By default, the function that triggers a warning when the SSH server uses an
                         insecure algorithm is enabled.
                  7.     Set the SSH authentication timeout interval.
                         ssh server timeout seconds

                         By default, the SSH authentication timeout interval is 60 seconds.
                         If you have not logged in successfully within the SSH authentication timeout
                         interval, the current connection is terminated to ensure security.
                  8.     Set the maximum number of SSH authentication retries.
                         ssh server authentication-retries times

                         By default, a maximum of three SSH authentication retries are supported.
                         The number of SSH authentication retries is limited to prevent unauthorized
                         access.
                  9.     Configure an ACL.
                         ssh [ ipv6 ] server acl { acl-number | acl-name }

                         By default, no ACL is configured.
                         An ACL is configured to determine which clients can log in to the current
                         device through SSH.
                  10. Set the minimum length of the RSA public key allowed for SSH server
                      authentication.
                         ssh server rsa-key min-length min-length-val

                         By default, the minimum length of the RSA public key allowed for SSH server
                         authentication is 512 bits.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          257
Security Configuration
Security Configuration                                                                            13 SSH Configuration


         Step 8 Configure the source interface or source IP address for the SSH server.

                  By default, no source interface or source IPv6 address is specified for an SSH
                  server.

                  ●      Configure the SSH server to use a specified interface as the source interface.
                         ssh server-source -i interface-type interface-number

