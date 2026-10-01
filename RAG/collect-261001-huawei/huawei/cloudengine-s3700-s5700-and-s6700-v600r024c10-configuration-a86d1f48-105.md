---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-105
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [13058, 13172]
sha256: 7cead26667736cb9da73c5846b75dda9008f55885e86dfbd318fd2cab7cc417a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                                 tls12_ck_rsa_aes_256_cbc    In this algorithm, RSA is used for key
                                 _sha256                     exchange and signature,
                                                             AES_256_CBC (with a key length of
                                                             256 bits and encryption mode of
                                                             CBC) is used for data encryption, and
                                                             SHA256 is used for message integrity
                                                             check.

                                 tls12_ck_ecdhe_rsa_with_    In this algorithm, ECDHE is used for
                                 aes_128_gcm_sha256          key exchange, RSA is used for
                                                             signature, AES_128_GCM (with a key
                                                             length of 128 bits and encryption
                                                             mode of GCM) is used for data
                                                             encryption, and SHA256 is used for
                                                             message integrity check.

                                 tls12_ck_ecdhe_rsa_with_    In this algorithm, ECDHE is used for
                                 aes_256_gcm_sha384          key exchange, RSA is used for
                                                             signature, AES_256_GCM (with a key
                                                             length of 256 bits and encryption
                                                             mode of GCM) is used for data
                                                             encryption, and SHA384 is used for
                                                             message integrity check.




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                          241
Security Configuration
Security Configuration                                                                12 SSL Configuration


                   TLS Version        Encryption Algorithm          Description
                                      Supported by a Cipher
                                      Suite

                                      tls12_ck_ecdhe_ecdsa_wit      In this algorithm, ECDHE is used for
                                      h_aes_128_gcm_sha256          key exchange, ECDSA is used for
                                                                    signature, AES_128_GCM (with a key
                                                                    length of 128 bits and encryption
                                                                    mode of GCM) is used for data
                                                                    encryption, and SHA256 is used for
                                                                    message integrity check.

                                      tls12_ck_ecdhe_ecdsa_wit      In this algorithm, ECDHE is used for
                                      h_aes_256_gcm_sha384          key exchange, ECDSA is used for
                                                                    signature, AES_256_GCM (with a key
                                                                    length of 256 bits and encryption
                                                                    mode of GCM) is used for data
                                                                    encryption, and SHA384 is used for
                                                                    message integrity check.

                   TLS1.3             tls13_aes_128_gcm_sha25       In this algorithm, AES_128_GCM
                                      6                             (with a key length of 128 bits and
                                                                    encryption mode of GCM) is used for
                                                                    data encryption and SHA256 is used
                                                                    for message integrity check.

                                      tls13_aes_256_gcm_sha38       In this algorithm, AES_256_GCM
                                      4                             (with a key length of 256 bits and
                                                                    encryption mode of GCM) is used for
                                                                    data encryption and SHA256 is used
                                                                    for message integrity check.

                                      tls13_chacha20_poly1305       In this algorithm, ChaCha20-
                                      _sha256                       Poly1305 is used for data encryption
                                                                    and SHA256 is used for message
                                                                    integrity check.

                                      tls13_aes_128_ccm_sha25       In this algorithm, AES_128_CCM
                                      6                             (with a key length of 128 bits and
                                                                    encryption mode of CCM) is used for
                                                                    data encryption and SHA256 is used
                                                                    for message integrity check.




Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create a cipher suite for an SSL policy and enter the customized view of the cipher
                suite.
                  ssl cipher-suite-list customization-policy-name

                  By default, no cipher suite is created for an SSL policy.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            242
Security Configuration
Security Configuration                                                                        12 SSL Configuration


         Step 3 Configure encryption algorithms supported by the cipher suite for the SSL policy.
                  set cipher-suite { tls1_ck_rsa_with_aes_256_sha | tls1_ck_rsa_with_aes_128_sha |
                  tls1_ck_dhe_rsa_with_aes_256_sha | tls1_ck_dhe_dss_with_aes_256_sha |
                  tls1_ck_dhe_rsa_with_aes_128_sha | tls1_ck_dhe_dss_with_aes_128_sha | tls12_ck_rsa_aes_128_cbc_sha |
                  tls12_ck_rsa_aes_256_cbc_sha | tls12_ck_rsa_aes_128_cbc_sha256 | tls12_ck_rsa_aes_256_cbc_sha256 |
                  tls12_ck_dhe_dss_aes_128_cbc_sha | tls12_ck_dhe_rsa_aes_128_cbc_sha |
                  tls12_ck_dhe_dss_aes_256_cbc_sha | tls12_ck_dhe_rsa_aes_256_cbc_sha |
                  tls12_ck_dhe_dss_aes_128_cbc_sha256 | tls12_ck_dhe_rsa_aes_128_cbc_sha256 |
                  tls12_ck_dhe_dss_aes_256_cbc_sha256 | tls12_ck_dhe_rsa_aes_256_cbc_sha256 |
                  tls12_ck_rsa_with_aes_128_gcm_sha256 | tls12_ck_rsa_with_aes_256_gcm_sha384 |
                  tls12_ck_dhe_rsa_with_aes_128_gcm_sha256 | tls12_ck_dhe_rsa_with_aes_256_gcm_sha384 |
                  tls12_ck_dhe_dss_with_aes_128_gcm_sha256 | tls12_ck_dhe_dss_with_aes_256_gcm_sha384 |
                  tls12_ck_ecdhe_rsa_with_aes_128_gcm_sha256 | tls12_ck_ecdhe_rsa_with_aes_256_gcm_sha384|
                  tls13_aes_128_gcm_sha256 | tls13_aes_256_gcm_sha384 | tls13_chacha20_poly1305_sha256 |
                  tls13_aes_128_ccm_sha256 | tls12_ck_ecdhe_ecdsa_with_aes_128_gcm_sha256 |
                  tls12_ck_ecdhe_ecdsa_with_aes_256_gcm_sha384 }

                  By default, no encryption algorithm is configured in the cipher suite for an SSL
                  policy.

