---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-120
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [15154, 15333]
sha256: 4596c79d0d44e5c30eb3b2f7a7276750e4fb356bbdaed0b39dfd50f431c029b2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  A CRL is a list of digital certificates that have been revoked by the issuing CA
                  before their scheduled expiration date and should no longer be trusted.

                  If a CA revokes a digital certificate, the declaration on authorized key pairs is
                  revoked before the certificate expires. After a certificate in a CRL expires, the
                  certificate is deleted in order to shorten the CRL.


14.4 Default Settings for HTTPS
                  Table 14-3 describes the default settings for HTTPS.


                  Table 14-3 Default settings for HTTPS

                   Parameter                                     Default Setting

                   HTTP                                          Disabled




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 278
Security Configuration
Security Configuration                                                                           14 HTTPS Configuration




14.5 Configuring an HTTPS Client

14.5.1 Configuring an SSL Policy

Context
                  Before configuring HTTPS, you need to deploy an SSL policy on the device, and
                  load the corresponding digital certificate. An SSL policy contains the SSL
                  parameters used during device startup, and only takes effect after it is associated
                  with an application layer protocol (such as HTTP).

                          NOTE

                         ● For enhanced security, you are advised to use more secure certificates. To be specific,
                           RSA/DSA certificates have a key length of 3072 bits or greater, and ECC certificates have
                           a key length of 256 bits or greater. Additionally, the certificate's hash algorithm is
                           SHA-256 or a later version.
                         ● The install feature-software WEAKEA command needs to be run if the key length of
                           the RSA/DSA certificate is less than 2048 bits, the key length of the ECC certificate is less
                           than 256 bits, or the hash algorithm of the certificate is SHA1, SHA-224, MD4, or MD5.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure an SSL policy and enter the SSL policy view.
                  ssl policy policy-name

         Step 3 Load a certificate for the SSL policy. The format must be the same as that of the
                certificate loaded on the HTTPS server.
                  ●      Load a PEM digital certificate for the SSL policy.
                         certificate load pem-cert certFile key-pair keyType key-file keyFile auth-code [ cipher authCode ]

                  ●      Load a PFX digital certificate for the SSL policy.
                         certificate load pfx-cert certFile key-pair keyType mac [ cipher macCode auth-code cipher
                         authCode ]
                         certificate load pfx-cert certFile key-pair keyType key-file keyFile auth-code [ cipher authCode ]

                  ●      Load a PEM certificate chain for the SSL policy.
                         certificate load pem-chain certFile key-pair keyType key-file keyFile auth-code [ cipher authCode ]

         Step 4 Load a trusted-CA file for the SSL policy. The format must be the same as that of
                the trusted-CA file loaded on the HTTPS server.
                  ●      Load a PEM trusted-CA file for the SSL policy.
                         trusted-ca load pem-ca caFile

                  ●      Load a PFX trusted-CA file for the SSL policy.
                         trusted-ca load pfx-ca caFile auth-code [ cipher authCode ]

                  ●      Load an ASN1 trusted-CA file for the SSL policy.
                         trusted-ca load asn1-ca caFile

                  ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              279
Security Configuration
Security Configuration                                                                             14 HTTPS Configuration


14.5.2 Configuring an HTTPS Client

Prerequisites
                  Before configuring an HTTPS client, you have completed the following task:
                  ●      There are reachable routes between the device and HTTPS server.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable HTTP and enter the HTTP view.
                  http

         Step 3 Configure an SSL policy for the HTTPS client.
                  client ssl-policy policy-name

         Step 4 Configure the HTTPS client to authenticate the server.
                  client ssl-verify peer

         Step 5 (Optional) Configure the source interface bound to the HTTPS client.
                  client source-interface { interface-name | interface-type interface-number }

                  By default, no source interface is bound to the HTTPS client.

         Step 6 (Optional) Configure the source IPv6 address and VPN for the HTTP client.
                  client ipv6 source-address ipv6-address [ vpn-instance ipv6-vpn-instance-name ]

                  ----End

14.5.3 Configuring HTTPS for System Software Download

Prerequisites
                  Before configuring HTTPS for system software download, you have completed the
                  following task:
                  ●      Configure an SSL policy.


Context
                  The device can download system software using HTTPS. If no SSL policy is
                  specified, the SSL policy configured on the HTTPS client is used.


Procedure
         Step 1 Download a file.
                  download file-url [ save-as file-path | [ ssl-policy policy-name [ ssl-verify peer [ verify-dns ] ] | verify-
                  dns ] | vpn-instance vpn-name | source-ip ip-address ] *

                  ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                   280
Security Configuration
Security Configuration                                                                              14 HTTPS Configuration


14.5.4 Configuring HTTPS for Local File Upload
Prerequisites
                  Before configuring HTTPS for local file upload, you have completed the following
                  task:
                  ●      Configure an SSL policy.

Context
                  The device can upload a local file to a server using HTTPS. If no SSL policy is
                  specified, the SSL policy configured on the HTTPS client is used. You can analyze
                  the device running status on the server based on the local file.

Procedure
         Step 1 Upload a file.
                  upload file-url local-file file-path [ [ ssl-policy policy-name [ ssl-verify peer [ verify-dns ] ] | verify-dns ]
                  | user-name name-value password password-value | vpn-instance vpn-name | source-ip ip-address ] *

                  ----End

14.5.5 Example for Configuring a Device as an HTTPS Client
Networking Requirements
                  As shown in Figure 14-4, before configuring the HTTPS client, you need to
                  configure an SSL policy on the device and load the corresponding digital
                  certificate. After the SSL policy is associated with HTTP, you can log in to the
                  HTTPS server from the HTTPS client.

                  Figure 14-4 Network diagram for accessing files on another device using HTTPS




Configuration Roadmap
                  The configuration roadmap is as follows:
                  1.     Configure an SSL policy for the HTTPS client.
                  2.     Configure an HTTPS client.

