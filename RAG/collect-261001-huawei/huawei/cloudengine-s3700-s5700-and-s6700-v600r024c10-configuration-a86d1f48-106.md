---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-106
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "disclosure", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [13173, 13296]
sha256: 11ae65bf61193c8136c328354516a54bec7b7c7dc5566d6f32f919179e028e44
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                          NOTE

                         For security purpose,you are not advised to use the weak security algorithm or weak
                         security protocols provided by this feature. If you need to use the weak security algorithm
                         or protocols, run the install feature-software WEAKEA command to install the weak
                         security algorithm or protocol feature package WEAKEA. By default, the device provides the
                         weak security algorithm or protocol feature package WEAKEA. For details about how to
                         install or uninstall the feature package, see "Upgrade Maintenance Configuration" in CLI
                         Configuration Guide > System Management Configuration.

                  The following table lists the commands that can be used only after the weak
                  security algorithm/protocol feature package is installed.
                   Command                              Parameters Available Only After Feature
                                                        Package Installation

                   set cipher-suite                     tls12_ck_dhe_dss_aes_128_cbc_sha,
                                                        tls12_ck_dhe_dss_aes_128_cbc_sha256,
                                                        tls12_ck_dhe_dss_aes_256_cbc_sha,
                                                        tls12_ck_dhe_dss_aes_256_cbc_sha256,
                                                        tls12_ck_dhe_rsa_aes_128_cbc_sha,
                                                        tls12_ck_dhe_rsa_aes_128_cbc_sha256,
                                                        tls12_ck_dhe_rsa_aes_256_cbc_sha,
                                                        tls12_ck_dhe_rsa_aes_256_cbc_sha256,
                                                        tls12_ck_rsa_aes_128_cbc_sha,
                                                        tls12_ck_rsa_aes_128_cbc_sha256,
                                                        tls12_ck_rsa_aes_256_cbc_sha,
                                                        tls12_ck_rsa_aes_256_cbc_sha256,
                                                        tls12_ck_rsa_with_aes_128_gcm_sha256,
                                                        tls12_ck_rsa_with_aes_256_gcm_sha384,
                                                        tls1_ck_dhe_dss_with_aes_128_sha,
                                                        tls1_ck_dhe_dss_with_aes_256_sha,
                                                        tls1_ck_dhe_rsa_with_aes_128_sha,
                                                        tls1_ck_dhe_rsa_with_aes_256_sha,
                                                        tls1_ck_rsa_with_aes_128_sha,
                                                        tls1_ck_rsa_with_aes_256_sha

                   ssl minimum version                  tls1.1



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       243
Security Configuration
Security Configuration                                                             12 SSL Configuration


                  SSL provides a handshake mechanism that allows a client and a server to establish
                  a session, authenticate each other's identity, and negotiate the key and cipher
                  suite. It is recommended that a cipher suite of TLS1.2 or a later version be used
                  during communication. In TLS versions, when the symmetric encryption algorithm
                  in CBC mode is used, data may be subject to a plaintext-recovery attack, causing
                  disclosure of encrypted data. Therefore, you are not advised to use the CBC mode
                  for data encryption in TLS versions.

                  ----End

12.5.2 Configuring an SSL Policy (Manually Loading a
Certificate)
Prerequisites
                  Before loading a trusted-CA file to an SSL policy, you have completed the
                  following task:
                  Apply for a certificate from a CA for the client or server and upload the certificate
                  to the security sub-directory of the system directory.

Context
                  SSL uses data encryption, identity authentication, and message integrity check
                  mechanisms to ensure security of TCP-based application layer protocols. An SSL
                  policy can be applied to application layer protocols to provide secure connections.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure an SSL policy and enter the SSL policy view.
                  ssl policy policy-name

                  By default, no SSL policy is configured.
         Step 3 (Optional) Set the elliptic curve parameter for the ECDHE algorithm.
                  ecdh group { nist | curve | brainpool | ffdhe } *

                  By default, the elliptic curve parameters of the ECDHE algorithm are Curve, Nist,
                  and Brainpool.
         Step 4 (Optional) Disable TLS 1.3 from using the brainpoolr1 curve.
                  ssl forbidden tls13-use-brainpoolr1

                  By default, TLS 1.3 allows the brainpoolr1 curve to be used.
         Step 5 (Optional) Configure the minimum path length of the digital certificate chain.
                  ssl verify certificate-chain minimum-path-length path-length

                  By default, the minimum path length of a digital certificate chain is 1.
         Step 6 (Optional) Configure the digital certificate verification function.
                  ssl verify basic-constrain enable
                  ssl verify version cert-version3 enable


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         244
Security Configuration
Security Configuration                                                                               12 SSL Configuration

                  ssl verify version crl-version2 enable
                  ssl verify key-usage enable
                  ssl verify certificate-signature-algorithm enable

                  By default, the digital certificate verification function is disabled.

         Step 7 (Optional) Configure a minimum SSL version for the SSL policy.
                  ssl minimum version { tls1.1 | tls1.2 | tls1.3 }

                  By default, the minimum version used by an SSL policy is TLS1.2.

                          NOTE

                         ● SSL policies support three SSL versions: TLS1.1, TLS1.2, and TLS1.3. TLS1.3 ensures the
                           highest security, followed by TLS1.2 and TLS1.1. TLS1.2 and TLS1.3 are recommended.
                         ● The tls1.1 parameter in this command can be used only after the weak security
                           algorithm/protocol feature package (WEAKEA) has been installed using the install
                           feature-software WEAKEA command.

         Step 8 Load a digital certificate for the SSL policy.

                  By default, no digital certificate is loaded for the SSL policy.
                          NOTE

