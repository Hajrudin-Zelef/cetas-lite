---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-9
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "disclosure"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [535, 653]
sha256: ce7cf6dac3aeb681428fdcd9327b1d5267d04d800d302bcc7fc1d84f9a7b6b01
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         –   When configuring a cleartext password, do not start and end the
                             password with %+%# or %@%# because this will allow the password to
                             be considered as a valid ciphertext that can be decrypted by the device
                             and make it visible in the configuration file.
                         –   Multiple features cannot use the same ciphertext password. For example,
                             the ciphertext password set for the AAA feature cannot be used for other
                             features.
                  ●      Encryption algorithms
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

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              3
Security Configuration
Security Configuration                                                           1 About This Document


                  ●      Reliability design declaration
                         Network planning and site design must comply with reliability design
                         principles and provide device- and solution-level protection. Device-level
                         protection includes planning principles of dual-network and inter-card dual-
                         link to avoid single point or single link of failure. Solution-level protection
                         refers to fast convergence protection mechanisms such as FRR and VRRP. If
                         solution-level protection is used, ensure that the primary and backup paths do
                         not share links or transmission devices. Otherwise, solution-level protection
                         may fail to take effect.

Reference Standards and Protocols
                  To obtain reference standards and protocols, log in to Huawei official website,
                  search for "standard and protocol compliance list", and download the Huawei S-
                  Series Switch Standard and Protocol Compliance List. If you have not obtained the
                  access permission of the document, see Help on the website to find out how to
                  obtain it.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            4
Security Configuration
Security Configuration                                                              2 Overview of Security




                                                2         Overview of Security


Threats to Network Security
                  Existing or potential security events may threaten the confidentiality, integrity, or
                  availability of resources in a network system. Network security services can
                  mitigate network security threats to a certain extent and within a specified scope.
                  It is critical that a device is able to forward data without data interception or
                  tampering. Network security involves:
                  ●      Confidentiality: Data stored, processed, and transmitted by a device must not
                         be leaked to unauthorized users, entities, or processes. Data must be available
                         only to authorized users.
                  ●      Integrity: Data cannot be modified without authorization. That is, network
                         information must not be deleted, modified, forged, unsequenced, replayed, or
                         inserted during storage and transmission.
                  ●      Availability: As long as required external resources are available, functions of
                         a device can be executed under specific conditions and at a specific time or
                         within a specific time period. Sustainable services must be provided to meet
                         carrier-class quality of service (QoS) requirements.
                  To meet the preceding requirements, plan and deploy network security on three
                  planes: management plane, control plane, and forwarding plane.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  5
Security Configuration
Security Configuration                                                         2 Overview of Security


                  Figure 2-1 Security deployment planes




