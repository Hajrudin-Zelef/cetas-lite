---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-103
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [12854, 12951]
sha256: 608a402ceaa8f764834681a2ad5901785fb0e47709fe1223925fb5198916ed35
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                                   tls1_ck_rsa_with_aes_128     In this algorithm, RSA is used for key
                                   _sha                         exchange and signature, AES_128 is
                                                                used for data encryption, and SHA is
                                                                used for message integrity check.

                                   tls1_ck_dhe_rsa_with_aes     In this algorithm, Diffie-Hellman and
                                   _256_sha                     RSA are used for key exchange, RSA
                                                                is used for signature, AES_256 is used
                                                                for data encryption, and SHA is used
                                                                for message integrity check.

                                   tls1_ck_dhe_dss_with_aes     In this algorithm, Diffie-Hellman is
                                   _256_sha                     used for key exchange, DSS is used
                                                                for signature, AES_256 is used for
                                                                data encryption, and SHA is used for
                                                                message integrity check.

                                   tls1_ck_dhe_rsa_with_aes     In this algorithm, Diffie-Hellman is
                                   _128_sha                     used for key exchange, RSA is used
                                                                for signature, AES_128 is used for
                                                                data encryption, and SHA is used for
                                                                message integrity check.

                                   tls1_ck_dhe_dss_with_aes     In this algorithm, Diffie-Hellman is
                                   _128_sha                     used for key exchange, DSS is used
                                                                for signature, AES_128 is used for
                                                                data encryption, and SHA is used for
                                                                message integrity check.

                   TLS1.2 and      tls12_ck_rsa_aes_128_cbc     In this algorithm, RSA is used for key
                   TLS1.3          _sha                         exchange and signature,
                                                                AES_128_CBC (with a key length of
                                                                128 bits and encryption mode of
                                                                CBC) is used for data encryption, and
                                                                SHA is used for message integrity
                                                                check.

                                   tls12_ck_rsa_aes_256_cbc     In this algorithm, Diffie-Hellman is
                                   _sha                         used for key exchange, DSS is used
                                                                for signature, AES_256_CBC (with a
                                                                key length of 256 bits and encryption
                                                                mode of CBC) is used for data
                                                                encryption, and SHA is used for
                                                                message integrity check.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           238
Security Configuration
Security Configuration                                                         12 SSL Configuration


                   TLS Version   Encryption Algorithm        Description
                                 Supported by a Cipher
                                 Suite

                                 tls12_ck_rsa_aes_128_cbc    In this algorithm, Diffie-Hellman is
                                 _sha256                     used for key exchange, RSA is used
                                                             for signature, AES_128_CBC (with a
                                                             key length of 256 bits and encryption
                                                             mode of CBC) is used for data
                                                             encryption, and SHA256 is used for
                                                             message integrity check.

                                 tls12_ck_dhe_rsa_aes_128    In this algorithm, Diffie-Hellman is
                                 _cbc_sha                    used for key exchange, RSA is used
                                                             for signature, AES_128_CBC (with a
                                                             key length of 128 bits and encryption
                                                             mode of CBC) is used for data
                                                             encryption, and SHA is used for
                                                             message integrity check.

                                 tls12_ck_dhe_dss_aes_128    In this algorithm, Diffie-Hellman is
                                 _cbc_sha                    used for key exchange, DSS is used
                                                             for signature, AES_128_CBC (with a
                                                             key length of 128 bits and encryption
                                                             mode of CBC) is used for data
                                                             encryption, and SHA is used for
                                                             message integrity check.

                                 tls12_ck_dhe_dss_aes_256    In this algorithm, Diffie-Hellman is
                                 _cbc_sha                    used for key exchange, DSS is used
                                                             for signature, AES_256_CBC (with a
                                                             key length of 256 bits and encryption
                                                             mode of CBC) is used for data
                                                             encryption, and SHA is used for
                                                             message integrity check.

                                 tls12_ck_dhe_rsa_aes_256    In this algorithm, Diffie-Hellman is
                                 _cbc_sha                    used for key exchange, RSA is used
                                                             for signature, AES_256_CBC (with a
                                                             key length of 256 bits and encryption
                                                             mode of CBC) is used for data
                                                             encryption, and SHA is used for
                                                             message integrity check.

