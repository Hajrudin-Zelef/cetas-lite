---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-109
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [13626, 13788]
sha256: 5621bf7e0edc59f350053bda62005d5be6da3221d4a884bf5773ae7ac91f5fcc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  13.1 Overview of SSH
                  13.2 Configuration Precautions for SSH
                  13.3 Understanding SSH
                  13.4 Default Settings for SSH
                  13.5 Configuring an SSH Server
                  13.6 Configuring an SSH Client
                  13.7 Troubleshooting SSH


13.1 Overview of SSH
Definition
                  Secure Shell (SSH) is a cryptographic network protocol for transmitting network
                  services (such as access and file transfer) securely over an unsecured network.
                  SSH versions are classified as SSH1.X and SSH2.0. SSH2.0 has an extended
                  structure and features both security and function improvements over SSH1.X,
                  including support for SFTP and more authentication and key exchange methods.

Purpose
                  Telnet lacks a secure authentication mode and uses TCP to transmit data in clear
                  text, which brings great security risks. As a result, the system is vulnerable to
                  attacks such as denial of service (DoS), IP address spoofing, and route spoofing.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  251
Security Configuration
Security Configuration                                                              13 SSH Configuration


                  Telnet alone is no longer viable with the increasing importance of network
                  security. SSH addresses Telnet's shortcomings by providing secure login and other
                  secure network services on an insecure network.
                  SSH provides a secure channel over TCP for data exchange using the well-known
                  port 22 (this can be changed as required for security purposes).


13.2 Configuration Precautions for SSH

13.3 Understanding SSH
                  This section uses SSH2.0 as an example to describe how SSH works.

                  Table 13-1 SSH working mechanism
                   Stage                Description

                   Connection           The SSH server listens to port 22 for SSH connections. After
                   setup                the client sends a connection request to the server, a TCP
                                        connection is set up between the client and server.

                   Version              The server and client determine which SSH version to use
                   negotiation          through version negotiation.

                   Algorithm            SSH supports multiple algorithms. Based on their supported
                   negotiation          algorithms, the server and client negotiate the following
                                        algorithms: key exchange algorithm for generating a session
                                        key, encryption algorithm for encrypting data, public key
                                        algorithm for digital signature and authentication, and hash-
                                        based message authentication code (HMAC) algorithm for
                                        data integrity protection.

                   Key exchange         The server and client dynamically generate a session key to
                                        protect data transmission and a session ID to identify the SSH
                                        connection through Diffie-Hellman key exchange. The client
                                        also authenticates the server during this stage.

                   Client               The client sends an authentication request to the server, and
                   authentication       the server authenticates the client.

                   Session request      After the authentication succeeds, the client sends a session
                                        request to the server, requesting the server to provide a
                                        certain type of service (STelnet, SFTP, or SCP). That is, the
                                        client requests to establish a session with the server.

                   Session              After a session is established, the server and client exchange
                   interaction          data.




13.4 Default Settings for SSH
                  13.4 Default Settings for SSH describes the default settings for SSH.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              252
Security Configuration
Security Configuration                                                           13 SSH Configuration


                  Table 13-2 Default settings for SSH

                   Parameter                                 Default Setting

                   STelnet server                            Disabled

                   SSH server port number                    22

                   Interval for updating the SSH server      0 hours, indicating that the server key
                   key pair                                  pair is never updated

                   SSH authentication timeout interval       60s

                   Maximum number of SSH                     3
                   authentication retries

                   Virtual type terminal (VTY) user          No authentication mode configured
                   interface authentication mode

                   Protocol supported by a VTY user          All protocols
                   interface

                   SSH user authentication mode              No authentication mode supported

                   Service type for the SSH user             No service type supported

                   Whether the SSH server assigns a          No public key assigned
                   public key to a user

                   User privilege level                      The default command privilege level
                                                             for a VTY user interface is 0.

                   First login for the SSH client            Disabled

                   Public key assigned by the SSH client     None
                   to the SSH server (Rivest-Shamir-
                   Adleman [RSA], Digital Signature
                   Algorithm [DSA], or Elliptic-curve
                   cryptography [ECC] public key)




13.5 Configuring an SSH Server

13.5.1 Configuring the SSH Server Function and Related
Parameters

Context
                  Configuring the SSH server function and related parameters includes the following
                  tasks: generating a local key pair for the server; enabling the SSH server; and
                  setting server parameters (such as the SSH server port number, interval for
                  updating the SSH server key pair, SSH authentication timeout interval, and
                  maximum number of SSH authentication retries).


Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           253
Security Configuration
Security Configuration                                                                       13 SSH Configuration


                          NOTE

                         ● To ensure security, you are advised to periodically change the key.
                         ● To ensure that the SSH algorithm negotiation is successful, the SSH client must support
                           the key exchange algorithm, encryption algorithm, public key algorithm, and HMAC
                           algorithm configured on the SSH server.
                         ● The SSH server does not support SSH1.X.
                         ● For security purposes, do not use the RSA algorithm whose length is less than 3072
                           digits. You are advised to use the ECC authentication algorithm instead.


Procedure
         Step 1 Enter the system view.
                  system-view

