---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-118
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [14836, 15004]
sha256: 4bc1df7cc216bf6c54d7470d4f1f44c00c36a69b05a9437f8adca75f9d0a0034
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Applicable Application                             Example

                   STelnet login                                      Example for Configuring a Device to
                                                                      Access Another Device as an STelnet
                                                                      Client (Password and RSA
                                                                      Authentication)

                   SFTP                                               Example for Configuring a Device as
                                                                      an SFTP Client (Password
                                                                      Authentication and RSA
                                                                      Authentication)


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     271
Security Configuration
Security Configuration                                                            13 SSH Configuration


                   Applicable Application                      Example

                   SCP                                         Example for Configuring a Device as
                                                               an SCP Client



                  ----End


13.7 Troubleshooting SSH

13.7.1 SSH Key Exchange Failure
Fault Symptom
                  When the device is used as an SSH server, the third-party client software fails to
                  exchange the key with the SSH server, causing SSH connection failure.

Possible Cause
                  Keys can be exchanged only after the client and server negotiate the key exchange
                  algorithm, encryption algorithm, public key algorithm, and HMAC algorithm. If
                  any algorithm fails to be negotiated, the key exchange will fail. The algorithm
                  negotiation fails because the algorithm supported by the client is not configured
                  on the SSH server.

Procedure
         Step 1 Use other methods to log in to the SSH server. For details, see "Logging In to the
                CLI" in CLI Configuration Guide - Basic Configuration.
         Step 2 Enable the debugging for the SSH server.
                  system-view
                  info-center enable
                  quit
                  terminal monitor
                  terminal debugging
                  debugging ssh server all

         Step 3 Connect the SSH server through third-party client software.
         Step 4 Check the algorithms supported by the client in the debug information displayed
                by the SSH server user terminal.
                  Copy the debug information to a TXT file, and use the following regular
                  expressions to search for the algorithms supported by the third-party client
                  software. For details, see Table 13-9.

                  Table 13-9 Searching for the algorithms supported by the third-party client
                  software
                   Algorithm                                   Regular Expression

                   Key exchange algorithm                      SSH protocol packet received.*key_ex


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           272
Security Configuration
Security Configuration                                                          13 SSH Configuration


                   Algorithm                                 Regular Expression

                   Encryption algorithm                      SSH protocol packet
                                                             received.*ciph_ctos

                   Public key algorithm                      SSH protocol packet
                                                             received.*ser_host_key

                   HMAC algorithm                            SSH protocol packet
                                                             received.*hmac_ctos


         Step 5 Configure the key exchange algorithm, encryption algorithm, public key algorithm,
                and HMAC algorithm supported by both the client and server on the SSH server.
                For details about the supported algorithms and configuration methods, see 13.5.1
                Configuring the SSH Server Function and Related Parameters.

                  ----End




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                        273
Security Configuration
Security Configuration                                                           14 HTTPS Configuration




                                      14                 HTTPS Configuration


                  14.1 Overview of HTTPS
                  14.2 Configuration Precautions for HTTP
                  14.3 Understanding HTTPS
                  14.4 Default Settings for HTTPS
                  14.5 Configuring an HTTPS Client


14.1 Overview of HTTPS
Definition
                  The Hypertext Transfer Protocol (HTTP) is an application-layer protocol that
                  transfers hypertext from WWW servers to local browsers. This protocol transmits a
                  variety of data, such as web pages, based on the TCP/IP protocol, and uses a
                  client-server model in which requests and responses are exchanged.

                  Hypertext Transfer Protocol Secure (HTTPS) is an HTTP protocol that runs on top
                  of Secure Sockets Layer (SSL) for secure transactions.

Purpose
                  The HTTP function provides a unified interface for users and features that use the
                  HTTP protocol to transmit data. HTTP, however, does not have any security
                  mechanisms; it transmits data in clear text and does not authenticate either
                  communication party. Therefore, data transmitted over such a protocol is
                  vulnerable to tampering, sacrificing transmission security. To overcome this, HTTPS
                  establishes an SSL encryption layer on HTTP, and is therefore more secure. HTTPS
                  improves device security in the following ways:
                  ●      The data exchanged between a client and a server is encrypted to ensure data
                         security and integrity, implementing secure device management.
                  ●      Certificate-based authentication is performed on a server and a client.
                  ●      The message authentication code (MAC) algorithm is used to verify message
                         integrity during message transmission.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         274
Security Configuration
Security Configuration                                                         14 HTTPS Configuration


                  The HTTPS client function is required when:
                  ●      The configuration data of a specified YANG file is loaded from an HTTPS
                         server to a configuration database through HTTPS in a Network Configuration
                         Protocol (NETCONF) scenario.


14.2 Configuration Precautions for HTTP

14.3 Understanding HTTPS
HTTP Message Format
                  HTTP is a request/response protocol and therefore involves request and response
                  messages.
                  Request message
                  An HTTP client sends a request message to an HTTP server. This request message
                  consists of three parts: request line, request header, and request body, as shown in
                  Figure 14-1.

                  Figure 14-1 Request format




                  Table 14-1 Request fields
                   Field                                      Description

