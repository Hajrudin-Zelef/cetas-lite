---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-113
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [14167, 14300]
sha256: 80d268d867777ff70e060c27df1f68aaacd62a2468e11d3758ffc562af3ec8b4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Context
                  Configuring an SSH user includes the following tasks: creating an SSH user and
                  configuring an authentication mode for the SSH user. The authentication modes
                  supported by the device include RSA, password, password-rsa, DSA, password-dsa,
                  ECC, password-ecc, password-x509v3-rsa, x509v3-rsa, sm2, password-sm2,
                  password-x509v3-ecdsa-sha2, password-sm2-sm3, and all.
                  ●      password-rsa: The password authentication and RSA authentication
                         requirements must be met.
                  ●      password-dsa: The password authentication and DSA authentication
                         requirements must be met.
                  ●      password-ecc: The password authentication and ECC authentication
                         requirements must be met.
                  ●      password-x509v3-rsa: The password authentication and X509V3-SSH-RSA
                         authentication requirements must be met.
                  ●      password-sm2: The password authentication and SM2 authentication
                         requirements must be met.
                  ●      password-x509v3-ecdsa-sha2: Sets the SSH user authentication mode to
                         password and X509V3-ECDSA-SHA2.
                  ●      password-sm2-sm3: Both password authentication and SM2-SM3
                         authentication need to be performed.
                  ●      all: The requirements of any one of the authentication modes must be met.

                          NOTE

                         For security purposes, do not use the RSA algorithm whose length is less than 3072 digits.
                         You are advised to use the ECC authentication algorithm instead.


Procedure
         Step 1 Enter the system view.
                  system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      260
Security Configuration
Security Configuration                                                                            13 SSH Configuration


         Step 2 Create an SSH user.
                  ssh user user-name

                  By default, no SSH user is created.

         Step 3 Configure an authentication mode for the SSH user.
                  ssh user user-name authentication-type { password | rsa | password-rsa | all | dsa | password-dsa | ecc |
                  password-ecc | sm2 | password-sm2 | password-x509v3-rsa | x509v3-rsa | password-x509v3-ecdsa-sha2
                  | x509v3-ecdsa-sha2 | sm2-sm3 | password-sm2-sm3 }

                  By default, no authentication type is configured for an SSH user.

                  If no SSH user is configured using the ssh user user-name command, run the ssh
                  authentication-type default password command to configure password
                  authentication as the default authentication mode. In this case, you only need to
                  configure AAA users. This helps to simplify the configuration if there are a large
                  number of users.

                  ●      The password authentication mode is implemented based on AAA. When the
                         password, password-rsa, password-x509v3-rsa, password-dsa, password-ecc,
                         password-x509v3-ecdsa-sha2, password-sm2-sm3, or password-sm2 mode is
                         used to log in to the device, you need to create a local user in the AAA view
                         with the same name as the SSH user.
                  ●      If an SSH user is authenticated using the RSA, DSA, SM2, or ECC
                         authentication mode, both the server and client need to generate the local
                         RSA, DSA, SM2, or ECC key pair (for details, see 13.5.1 Configuring the SSH
                         Server Function and Related Parameters) and have each other's public key
                         configured locally.

                  Configure the authentication mode based on the preceding configuration. For
                  details, see Table 13-3.


                  Table 13-3 Configuration in different authentication modes

                   Authentication Mode                                  Configuration Notes

                   password                                             Create an AAA user with the same
                                                                        username as the SSH user. For details,
                                                                        see Table 13-4.

                   RSA, DSA, ECC, SM2-SM3, or SM2                       Configure the device to generate a
                   authentication                                       local RSA, DSA, SM2, or ECC key pair.
                                                                        For details, see Table 13-5.
                                                                        When SM2-SM3 authentication is
                                                                        used, the type of the key pair used is
                                                                        also SM2. You need to run the sm2
                                                                        key-pair label label-name command
                                                                        to generate a local key pair.

                   password-rsa, password-dsa,                          Create an AAA user with the same
                   password-sm2, password-sm2-sm3, or                   username as the SSH user and
                   password-ecc authentication                          generate a local RSA, DSA, SM2, or
                                                                        ECC key pair. For details, see Table
                                                                        13-4 and Table 13-5.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                             261
Security Configuration
Security Configuration                                                           13 SSH Configuration


                   Authentication Mode                       Configuration Notes

                   x509v3-rsa or x509v3-ecdsa-sha2           Bind the SSH user to the PKI realm.
                                                             For details, see Table 13-6.

                   password-x509v3-rsa or password-          Create an AAA user with the same
                   x509v3-ecdsa-sha2                         username as the SSH user and bind
                                                             the AAA user to the PKI realm. For
                                                             details, see Table 13-4 and Table
                                                             13-6.




                  Table 13-4 Creating a local user with the same name as the SSH user in the AAA
                  view
                              Step                     Command                   Description

                   Enter the system view.      system-view                 -

                   Enter the AAA view.         aaa                         -

                                               local-user user-name        For security purposes,
                   Configure the local
                                               password irreversible-      change the password
                   username and password.
                                               cipher password             periodically.

                   Configure a service type    local-user user-name
                                                                           -
                   for the local user.         service-type ssh

