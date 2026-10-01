---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-8
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "disclosure"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [450, 572]
sha256: e4a2df90d7368f4a914e68ebf28be051db9d8a0a013b256f015fb9be661ef31c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                  [ x | y | ... ]                   Optional alternative items are grouped in square
                                                    brackets and separated by vertical bars. One or none
                                                    is selected.

                  { x | y | ... } *                 Alternative items are grouped in braces and
                                                    separated by vertical bars. A minimum of one or a
                                                    maximum of all can be selected.

                  [ x | y | ... ] *                 Optional alternative items are grouped in square
                                                    brackets and separated by vertical bars. Many or
                                                    none can be selected.

                  &<1-n>                            This parameter before the & sign can be repeated 1
                                                    to n times.

                  #                                 This parameter before the # sign can be repeated 1
                                                    to n times.




Interface Numbering Conventions
                 Interface numbers used in this document are examples and must be replaced
                 according to configuration requirements.

Security Conventions
                 ●      Password setting
                        –    Configuring a ciphertext password is recommended. For security
                             purposes, do not disable password complexity check, and change the
                             password periodically.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                              2
MPLS Configuration
MPLS Configuration                                                              1 About This Document


                        –   When configuring a cleartext password, do not start and end the
                            password with %+%# or %@%# because this will allow the password to
                            be considered as a valid ciphertext that can be decrypted by the device
                            and make it visible in the configuration file.
                        –   Ciphertext passwords set for various features must be different. For
                            example, the ciphertext password set for the AAA feature cannot be used
                            for other features.
                 ●      Encryption algorithm
                        Currently, the device supports the following encryption algorithms: DES, 3DES,
                        AES, DSA, RSA, DH, ECDH, HMAC, SHA1, SHA2, and MD5. Select an
                        encryption algorithm according to the application scenario. Use the
                        recommended encryption algorithm; otherwise, security protection
                        requirements may not be met.
                        –   Recommended symmetric encryption algorithm: AES (with a 128-bit or
                            longer key).
                        –   Recommended asymmetric encryption algorithm: RSA (with a 3072-bit or
                            longer key). Use different key pairs for encryption and signature.
                        –   Recommended encryption algorithm for the digital signature: RSA (with a
                            3072-bit or longer key).
                        –   Recommended encryption algorithm for key negotiation: DH (with a
                            3072-bit or longer key) or ECDH (with a 256-bit or longer key).
                        –   Recommended hash algorithm: SHA2 (256-bit or higher).
                        –   Recommended hash-based message authentication code (HMAC)
                            algorithm: HMAC-SHA2.
                        –   The SHA1, SHA2, and MD5 encryption algorithms are irreversible, and the
                            DES, 3DES, RSA, and AES encryption algorithms are reversible.
                        –   In SSH2.0, when the symmetric encryption algorithm in CBC mode is
                            used, data may be subject to a plaintext-recovery attack, causing
                            disclosure of encrypted data. Therefore, you are not advised to use the
                            CBC mode for data encryption in SSH2.0.
                        –   SSL provides a handshake mechanism that allows a client and a server to
                            establish a session, authenticate each other's identity, and negotiate the
                            key and cipher suite. It is recommended that a cipher suite of TLS 1.2 or a
                            later version be used during communication. In TLS versions, when the
                            symmetric encryption algorithm in CBC mode is used, data may be
                            subject to a plaintext-recovery attack, causing disclosure of encrypted
                            data. Therefore, you are not advised to use the CBC mode for data
                            encryption in TLS versions.
                 ●      Personal data
                        Some personal data (such as MAC or IP addresses of terminals) may be
                        obtained or used during operation or fault locating of your purchased
                        products, services, or features, so you have an obligation to make privacy
                        policies and take proper measures according to applicable laws of the country
                        to fully protect personal data.
                 ●      The terms mirrored port, port mirroring, flow mirroring, and mirroring in this
                        document are mentioned only to describe the purpose of detecting faults and
                        errors in communication transmission. They do not involve collection or
                        processing of any personal information or communication data of users.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              3
MPLS Configuration
MPLS Configuration                                                                1 About This Document


                 ●      Reliability design
                        Reliability must be factored in during network planning and site design to
                        ensure device- and solution-level protection. Device-level protection refers to
                        adding redundancy, for example, duplicating networks, planes, devices, and
                        inter-board links, to prevent single points of failure. Solution-level protection
                        refers to fast convergence protection, such as FRR and VRRP. If solution-level
                        protection is used, ensure that the primary and backup paths do not share
                        links or transmission devices. Otherwise, solution-level protection may fail to
                        take effect.

Reference Standards and Protocols
                 To obtain reference standards and protocols, log in to Huawei official website,
                 search for "standard and protocol compliance list", and download the Huawei S-
                 Series Switch Standard and Protocol Compliance List. If you have not obtained the
                 access permission of the document, see Help on the website to find out how to
                 obtain it.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                   4
MPLS Configuration
MPLS Configuration                                                         2 Basic MPLS Configuration




                                 2          Basic MPLS Configuration


