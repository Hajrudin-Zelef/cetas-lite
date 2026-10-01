---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-121
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [15334, 15494]
sha256: 47830e31d098af7f2b98aff9abf92b02829f90fa1ecce49d2c8ac1397cc318c4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Procedure
         Step 1 Configure an SSL policy for the HTTPS client.
                  # Configure a PKI realm. Import the local certificates and private key file.
                  <HUAWEI> system-view
                  [HUAWEI] pki realm domain1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                   281
Security Configuration
Security Configuration                                                                        14 HTTPS Configuration

                  [HUAWEI-pki-realm-domain1] quit
                  [HUAWEI] pki import-certificate local realm domain1 pem filename https_local.pem
                  [HUAWEI] pki import rsa-key-pair http-key pem restconf.pem //Skip this step if the private key file is
                  generated on the local device.

                  # Import the CA certificate of the HTTP client to verify the validity of the local
                  certificate of the HTTP client.
                  [HUAWEI] pki import-certificate ca realm domain1 pem filename https_ca.pem

                  # Configure an SSL policy and bind it to the PKI realm.
                  [HUAWEI] ssl policy policy1
                  [HUAWEI-ssl-policy-policy1] pki-domain domain1
                  [HUAWEI-ssl-policy-policy1] quit

         Step 2 Configure an HTTPS client.
                  [HUAWEI] http
                  [HUAWEI-http] client ssl-policy policy1
                  [HUAWEI-http] client ssl-verify peer
                  [HUAWEI-http] quit

                  ----End

Verifying the Configuration
                  Run the display ssl policy command to check whether the HTTPS client is
                  configured successfully.
                  [HUAWEI] display ssl policy
                        SSL Policy Name: policy1
                            PKI domain: domain_name
                      Policy Applicants: HTTP-CLIENT
                         Key-pair Type:
                   Certificate File Type:
                       Certificate Type:
                    Certificate Filename:
                      Key-file Filename:
                             CRL File:
                        Trusted-CA File:


Configuration Scripts
                  #
                  ssl policy policy1
                   pki-domain domain1
                  #
                  http
                   client ssl-policy policy1
                   client ssl-verify peer
                  #
                  pki realm domain1
                  #
                  return




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           282
Security Configuration
Security Configuration                                                       15 Keychain Configuration




                                15                 Keychain Configuration

                  A keychain itself manages just the encryption and authentication keys, and only
                  takes effect when used in applications.
                  15.1 Overview of Keychains
                  15.2 Understanding Keychains
                  15.3 Configuration Precautions for Keychain
                  15.4 Default Settings for Keychains
                  15.5 Configuring a Keychain


15.1 Overview of Keychains
Definition
                  A keychain, as its name implies, is a chain of keys used to open the encryption
                  lock that is constantly changed by an application.
                  A key in a keychain is not an algorithm or a key string; rather, it is a set of
                  encryption and authentication rules. A keychain centrally controls and flexibly
                  manages a series of its own keys to provide dynamic security authentication
                  services for applications.

Purpose
                  Before an application that runs a routing protocol (for example, RIP, IS-IS, OSPF, or
                  BGP) establishes a session with the peer end, it needs to set up a transport-layer
                  connection.
                  To ensure the security of an application's session connections and exchanged data,
                  MD5 can be used to authenticate packets; however, this has the following
                  disadvantages:
                  ●      The MD5 algorithm is relatively simple and cannot meet high security
                         requirements of networks.
                  ●      MD5 keys must be manually updated at intervals to ensure key security. MD5
                         algorithms and keys are configured in applications, and are statically bound

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           283
Security Configuration
Security Configuration                                                               15 Keychain Configuration


                         to applications in one-to-one mappings. Therefore, you need to manually
                         update the keys configured on the applications of the devices at both ends
                         one by one.
                  To eliminate these disadvantages, a keychain for application authentication is
                  introduced:
                  ●      In each key of a keychain, an algorithm that is more secure than MD5 can be
                         selected. In the future, more highly secure algorithms will be available.
                  ●      Each key in a keychain has an independent algorithm, key string, and lifetime.
                         Applications on devices at both ends use keychain authentication; that is, they
                         need to match multiple keys. Therefore, the authentication algorithms and
                         key strings can be automatically and periodically updated on multiple
                         applications of the devices at both ends based on the lifetimes of the keys.
                  ●      When keys in a keychain are dynamically updated, the transport-layer
                         connections in use do not need to be disconnected and reconnected,
                         maintaining the stability of session connections and service continuity.


15.2 Understanding Keychains

15.2.1 Related Concepts of Keychains
                  A keychain is a chain of keys, known as a series of encryption and authentication
                  rules.

Three Elements of a Key
                  Each key in a keychain consists of three elements:
                  ●      Authentication algorithm: supports MD5, SHA-1, HMAC-MD5, HMAC-
                         SHA1-12, HMAC-SHA1-20, HMAC-SHA-256, SHA-256, SM3, HMAC-SHA-384,
                         and HMAC-SHA-512.
                               NOTE

                              MD5, HMAC-MD5, and SHA-1 algorithms are not recommended since they are less
                              secure.
                  ●      Authentication key string: a character string used for encryption. The same
                         clear text can be encrypted using different key strings to obtain different
                         ciphertexts. The same ciphertext can be obtained only when the same key
                         string is used for encryption.
                  ●      Lifetime: indicates the time interval within which a key is valid. If the lifetime
                         of a key expires, the key is replaced by another active key.
                          NOTE

                         When an authentication algorithm and an authentication key string are used for encryption
                         calculation of packets, a string of fixed-length message authentication code (MAC) is
                         generated.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   284

