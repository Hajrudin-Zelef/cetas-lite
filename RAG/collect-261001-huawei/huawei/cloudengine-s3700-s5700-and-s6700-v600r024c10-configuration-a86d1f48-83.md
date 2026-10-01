---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-83
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10148, 10261]
sha256: f76bf77c90cc1ba1698a7f2cede7270cde51be2f98ccbfb9e394ac13cff4130b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                              If an RA exists and the certificate enrollment request is issued by the RA, cert-file-
                              name must be set to the CA certificate of the RA. Otherwise, the certificate verification
                              fails, which in turn causes the certificate application to fail.
                  ●      If this command is not configured and the CMPv2 server signs its certificate
                         response, the device uses the certificates on the device and in the CMPv2
                         server's response to build a certificate chain, and then verifies the CMPv2
                         server's response signature based on the certificate chain.

        Step 11 Configure the mode for applying for a local certificate based on the site
                requirements.
                  ●      Initial local certificate application (IR) using a message authentication code
                         a.   Configure the authentication mode of CMPv2-based initial local
                              certificate application.
                              cmp-request origin-authentication-method message-authentication-code

                              By default, the authentication mode of CMPv2-based initial local
                              certificate application is message authentication code.
                         b.   Configure the reference value and secret value of the message
                              authentication code.
                              cmp-request message-authentication-code reference-value [ secret-value ]
                              quit


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       188
Security Configuration
Security Configuration                                                                     10 PKI Configuration


                                    NOTE

                                  The secret value of the message authentication code is the value of CMP secret
                                  key on the web UI of the CMPv2 server. The reference value is configurable, and
                                  must be a string of case-sensitive characters excluding question marks (?). The
                                  value is a string of 1 to 128 characters in clear text or a string of 48 to 188
                                  characters in ciphertext.
                         c.   In the system view, submit an initial certificate enrollment request to the
                              CMPv2 server based on the CMP session configuration.
                              pki cmp initial-request session session-name

                              After this command is configured, the system first checks the CMP
                              session configuration to determine whether the certificate application
                              condition is met. If the condition is not met, an error message is
                              displayed. If the condition is met, the system initiates an initial certificate
                              enrollment request according to the configuration. The certificate
                              obtained is saved in a file to the device storage without being imported
                              to the memory. If the server provides a CA certificate in a response, the
                              CA certificate is also saved in a file.
                  ●      Initial local certificate application (IR) using a signature
                         a.   Configure the authentication mode of CMPv2-based initial local
                              certificate application.
                              cmp-request origin-authentication-method signature

                         b.   Configure the certificate carried in a CMPv2 request for identity
                              authentication.
                              cmp-request authentication-cert cert-name
                              quit

                              This certificate is an additional certificate (known as the external identity
                              certificate) and must be issued by another trusted CA.
                         c.   In the system view, submit an initial certificate enrollment request to the
                              CMPv2 server based on the CMP session configuration.
                              pki cmp initial-request session session-name

                              After this command is configured, the system first checks the CMP
                              session configuration to determine whether the certificate application
                              condition is met. If the condition is not met, an error message is
                              displayed. If the condition is met, the system initiates an initial certificate
                              enrollment request according to the configuration. The obtained
                              certificate is saved in a file to the flash:/pki/public directory without
                              being imported to the memory. If the server provides a CA certificate in a
                              response, the CA certificate is also saved in a file.
                  ●      Signature-based non-initial local certificate application (CR)
                         a.   Configure the certificate carried in a CMPv2 request for identity
                              authentication.
                              cmp-request authentication-cert cert-name
                              quit

                              This certificate is an additional certificate (known as the external identity
                              certificate) and must be issued by another trusted CA.
                         b.   In the system view, submit a certificate enrollment request to the CMPv2
                              server based on the CMP session configuration.
                              pki cmp certificate-request session session-name

                              After this command is configured, the system first checks the CMP
                              session configuration to determine whether the certificate application

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   189
Security Configuration
Security Configuration                                                                 10 PKI Configuration


                              condition is met. If the condition is not met, an error message is
                              displayed. If the condition is met, the system initiates a certificate
                              enrollment request according to the configuration. The certificate
                              obtained is saved in a file to the device storage without being imported
                              to the memory.

        Step 12 Configure the mode for updating the local certificate based on the site
                requirements.
                  ●      Manual update
                         a.   Enter the CMP session view and configure the certificate for identity
                              authentication in a CMPv2 request.
                              pki cmp session session-name
                              cmp-request authentication-cert cert-name
                              quit

                              This certificate is the local certificate that the CA has issued to the device
                              and needs to be replaced by a new local certificate.
                         b.   In the system view, submit a Key Update Request (KUR) to the CMPv2
                              server based on the CMP session configuration.
                              pki cmp keyupdate-request session session-name

