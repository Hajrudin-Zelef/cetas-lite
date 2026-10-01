---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-115
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [14420, 14579]
sha256: 1effe7507e7972f176310eb1efb20f86855d0faaf80c7d3e12c91b0e8926dcee
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                                               ssh user user-name
                   Assign an RSA, DSA,
                                               assign { rsa-key | dsa-
                   SM2, or ECC public key                                  -
                                               key | ecc-key | sm2-key}
                   to the SSH user.
                                               key-name



                  Table 13-6 Binding a PKI realm to the SSH user
                              Step                     Command                    Description

                   Enter the system view.      system-view                 -




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           264
Security Configuration
Security Configuration                                                                             13 SSH Configuration


                                 Step                             Command                          Description

                                                                                            Assign a PKI certificate to
                   Bind a PKI realm to the               ssh user user-name                 the SSH server.
                   SSH user.                             assign pki pki-name                The prerequisite is that
                                                                                            PKI has been configured.


         Step 4 Configure a service type for the SSH user.
                  ssh user user-name service-type { all | { sftp | stelnet | snetconf }*}

                  By default, no service type is configured for an SSH user.
         Step 5 (Optional) Configure SAN/CN verification for the SSH user.
                  ssh user user-name cert-verify-san enable

                  By default, the device does not check whether the CN or SAN in the certificate
                  contains the realm name of the authentication user.
                  After a PKI realm is bound to the SSH user, the system checks whether the Subject
                  Alternative Name (SAN) or common name (CN) in the PKI certificate contains the
                  realm name of the authentication user to enhance security.
         Step 6 (Optional) Configure the source interface that allows SSH users to establish
                connections.
                  ssh user user-name source -i { interface-type interface-num | interface-name }

                  By default, no source interface is specified for SSH users to establish connections.

                  ----End

13.5.4 Applying SSH
Context
                  SSH is a security protocol, and an SSH policy takes effect only when it is
                  associated with an application.

Procedure
         Step 1 Apply an SSH policy. Table 13-7 lists the main applications of an SSH policy when
                the device functions as an SSH server.

                  Table 13-7 Main applications of an SSH policy when the device functions as an
                  SSH server

                   Applicable Application                                  Example

                   STelnet login                                           Example for Configuring STelnet Login
                                                                           for IPv4 Users (Local Authentication)

                   SFTP                                                    Example for Configuring a Device as
                                                                           an SFTP Server (IPv4)

                   SCP                                                     Configuring a Device as an SCP Server


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         265
Security Configuration
Security Configuration                                                                       13 SSH Configuration


                   Applicable Application                             Example

                   NETCONF                                            Example for Configuring a Device to
                                                                      Communicate with ncclient Using
                                                                      NETCONF



                  ----End


13.6 Configuring an SSH Client

13.6.1 Configuring the Mode for Connecting a Device to the
SSH Server for the First Time

Context
                  When a device functioning as an SSH client connects to the SSH server for the first
                  time, the client cannot check the validity of the SSH server because the client does
                  not have the server's public key or has no related PKI certificate bound. As a
                  result, the connection fails.

                  You can use any of the following methods based on requirements:

                  ●      Enable the first login function for the SSH client. The client will not perform a
                         validity check on the SSH server, meaning that the first connection is set up
                         successfully. The client then automatically assigns a public key to the server
                         and saves it for subsequent authentication.
                  ●      Bind the public key of the SSH server to the SSH client. The public key
                         generated on the server will be saved on the client, ensuring that the client
                         can successfully validate the server upon the first connection.
                  ●      Bind the PKI realm used for authentication with the server to the client to
                         ensure that the certificate of the SSH server is valid when the client connects
                         to the server for the first time.

                  The first method is simple, and the other two methods are complex but more
                  secure.

                          NOTE

                         ● To ensure security, you are advised to periodically change the key.
                         ● For security purposes, do not use the RSA algorithm whose length is less than 3072
                           digits. You are advised to use the ECC authentication algorithm instead.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 (Optional) Generate a local key pair.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     266
Security Configuration
Security Configuration                                                                        13 SSH Configuration


                  Perform this step only when the device connects to the SSH server through RSA,
                  DSA, or ECC authentication. This step is not required if password authentication is
                  used.

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

