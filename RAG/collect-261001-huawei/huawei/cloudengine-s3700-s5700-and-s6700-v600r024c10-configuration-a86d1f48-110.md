---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-110
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [13789, 13915]
sha256: 63970ca1c3ef3a1a3bfbd69d789a2798c745ee9443decaf01ae00c77988a629d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 2 Generate a local key pair.

                  Method 1: Generate a local RSA, DSA, or ECC key pair.

                  ●      Generate a local RSA key pair.
                         rsa local-key-pair create

                  ●      Generate a DSA key pair.
                         dsa local-key-pair create

                  ●      Generate an ECC key pair.
                         ecc local-key-pair create

                  After a key pair is generated, you can run the display rsa local-key-pair public,
                  display dsa local-key-pair public, or display ecc local-key-pair public command
                  to view information about the RSA, DSA, or ECC public key in the local key pair.

                  If you no longer need the local DSA or ECC key pairs, run the dsa local-key-pair
                  destroy or ecc local-key-pair destroy command to destroy all the local DSA or
                  ECC key pairs. After this command is run, the file that stores the corresponding
                  keys on the device is cleared. Exercise caution when running this command.

                  Destroy all the local DSA keys.

                  Method 2: Generate a labeled SM2, RSA, DSA, or ECC key pair.

                          NOTE

                  You can generate up to 20 key pairs by using this method. To enhance communication security,
                  rotate between these key pairs at different periods. To limit the maximum number of key pairs
                  that the device can generate, you can run the rsa key-pair maximum, dsa key-pair maximum,
                  or ecc key-pair maximum command.

                  1.     Generate a labeled RSA, DSA, SM2, or ECC key pair.
                         –    Generate a labeled RSA key pair.
                              rsa key-pair label label-name [ modulus modulus-bits ]

                         –    Generate a labeled DSA key pair.
                              dsa key-pair label label-name [ modulus modulus-bits ]

                         –    Generate a labeled ECC key pair.
                              ecc key-pair label label-name [ modulus modulus-bits ]

                         –    Generate a labeled SM2 key pair.
                              sm2 key-pair label label-name

                  2.     Assign a host key or PKI certificate to the SSH server.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    254
Security Configuration
Security Configuration                                                                            13 SSH Configuration

                         ssh server assign { rsa-host-key key-name | dsa-host-key key-name | ecc-host-key key-name | sm2-
                         host-key key-name | pki key-name }

                  By default, no key or PKI certificate is assigned to an SSH server.
                  After the key pair is generated, you can run the display rsa key-pair [ brief |
                  label label-name ], display dsa key-pair [ brief | label label-name ], display ecc
                  key-pair [ brief | label label-name ], display sm2 key-pair [ brief | label label-
                  name ] commands to view information about the labeled RSA, DSA, SM2, or ECC
                  key pair.
         Step 3 Enable the public key algorithm for the SSH server.
                  ssh server publickey { dsa | ecc | rsa | x509v3-ssh-rsa | rsa_sha2_256 | rsa_sha2_512 | sm2 | x509v3-
                  rsa2048-sha256 | x509v3-ecdsa-sha2 | sm2-sm3 } *

                  By default, the ECC, RSA_SHA2_256, and RSA_SHA2_512 public key algorithms are
                  enabled.

                          NOTE

                         Parameters dsa, rsa, and x509v3-ssh-rsa in the command can be used only after the weak
                         security algorithm/protocol feature package is installed by running the install feature-
                         software WEAKEA command.
                         To log in to the device using public key authentication, ensure that the public key
                         algorithms enabled for the SSH server are the same as those configured for SSH users using
                         the ssh user authentication-type command. Otherwise, login to the device fails.

         Step 4 Enable the SSH server function.
                  ●      Enable the STelnet server function on the device.
                         stelnet [ ipv4 | ipv6 ] server enable
                         By default, the STelnet service is disabled.
                  ●      Enable the SFTP server function on the device.
                         sftp [ ipv4 | ipv6 ] server enable
                         By default, the SFTP service is disabled.
                  ●      Enable the SCP server function on the device.
                         scp [ ipv4 | ipv6 ] server enable
                         By default, the SCP service is disabled.
         Step 5 Configure the port number of the SSH server.
                  ssh [ ipv4 | ipv6 ] server port port-number

                  By default, the port number of an SSH server is 22.
                  If a new port number is configured, the SSH server disconnects from all SSH
                  clients and uses the new port number to establish connections. This is more secure
                  as it prevents attackers from connecting to the server straight through the
                  standard SSH port.
         Step 6 Enable the keepalive function on the SSH server.
                  undo ssh server keepalive disable

                  By default, the keepalive function is enabled on an SSH server.
                  When this function is enabled, the SSH server will respond to keepalive packets
                  sent from the SSH client. If the SSH client does not receive any keepalive response
                  packets from the SSH server, the client will disconnect from the server. This
                  ensures that server resources are not wasted due to the client repeatedly
                  attempting to reconnect.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              255
Security Configuration
Security Configuration                                                                             13 SSH Configuration


         Step 7 (Optional) Configure extended attributes for the SSH server.
                  1.     Configure a key exchange algorithm list for the SSH server.
                         ssh server key-exchange { dh_group_exchange_sha256 | dh_group_exchange_sha1 |
                         dh_group1_sha1 | ecdh_sha2_nistp256 | ecdh_sha2_nistp384 | ecdh_sha2_nistp521 | sm2_kep |
                         dh_group14_sha1 | dh_group16_sha512 | curve25519_sha256 | sm2_sm3 } *

                         By default, the SSH server uses dh_group_exchange_sha256,
                         dh_group16_sha512, and curve25519_sha256.
                         The client and server negotiate the key exchange algorithm for packet
                         transmission. During this negotiation, the server compares the key exchange
                         algorithm list sent by the client with its own, and selects the first algorithm
                         that appears on both lists for packet transmission. If there are no matches,
                         the negotiation fails.
                                  NOTE

