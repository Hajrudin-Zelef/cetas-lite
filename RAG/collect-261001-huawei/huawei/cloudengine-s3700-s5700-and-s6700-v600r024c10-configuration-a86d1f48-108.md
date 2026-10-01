---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-108
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [13449, 13625]
sha256: 005b2b775857fd8e1f8ff9f160e91ba6bf0771a8b5999267dd6ba82d35c94d97
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 7 (Optional) Configure a minimum SSL version for the SSL policy.
                  ssl minimum version { tls1.1 | tls1.2 | tls1.3 }

                  By default, the minimum version used by an SSL policy is TLS1.2.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           247
Security Configuration
Security Configuration                                                                       12 SSL Configuration


                          NOTE

                         ● SSL policies support three SSL versions: TLS1.1, TLS1.2, and TLS1.3. TLS1.3 ensures the
                           highest security, followed by TLS1.2 and TLS1.1. TLS1.2 and TLS1.3 are recommended.
                         ● The tls1.1 parameter in this command can be used only after the weak security
                           algorithm/protocol feature package (WEAKEA) has been installed using the install
                           feature-software WEAKEA command.

         Step 8 Bind a PKI realm to the SSL policy. After a PKI realm is bound, the SSL policy uses
                the certificates and CRLs in the PKI realm.
                  pki-domain pki-domain

         Step 9 (Optional) Exclude key exchange algorithms from the cipher suite list.
                  cipher-suite exclude key-exchange { rsa | dhe } *
                  cipher-suite exclude cipher mode cbc
                  cipher-suite exclude hmac sha1

                  By default, an SSL policy does not support the RSA key exchange algorithm, CBC
                  encryption algorithm, or SHA1 digest algorithm.

        Step 10 (Optional) Bind a cipher suite to the SSL policy.
                  binding cipher-suite-customization customization-name

                  By default, no cipher suite is bound to an SSL policy. In this case, all encryption
                  algorithms can be used.

                  The cipher suite to be bound to the SSL policy must have been configured. For
                  details, see 12.5.1 (Optional) Configuring a Cipher Suite for an SSL Policy.

        Step 11 (Optional) Enable the SSL renegotiation function.
                  ssl renegotiation enable

                  By default, the SSL renegotiation function is disabled.

                  When SSL is used for data transmission, you can enable the SSL renegotiation
                  function to periodically update the key and algorithm without interrupting the
                  connection, improving communication security.

                  ----End

12.5.4 Configuring an SSL Policy Using an SM Cipher Suite
(Loading a Certificate Using PKI)
Prerequisites
                  An SM2-SM3 certificate has been loaded to the PKI realm to be bound. The
                  loaded certificate can be an initial device certificate or a digital certificate applied
                  by a user. For details on how to load a certificate using PKI, see CLI Configuration
                  Guide > PKI Configuration.

Context
                  If an application needs to use the SM cipher suites TLS_SM4_CCM_SM3 and
                  TLS_SM4_GCM_SM3, the matching ECDH group algorithm and signature
                  algorithm must be used. SM cipher suites TLS_SM4_CCM_SM3 and
                  TLS_SM4_GCM_SM3 cannot be used with other cipher suites.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     248
Security Configuration
Security Configuration                                                            12 SSL Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create a cipher suite for an SSL policy and enter the customized view of the cipher
                suite.
                  ssl cipher-suite-list customization-policy-name

                  By default, no cipher suite is created for an SSL policy.

         Step 3 Configure encryption algorithms supported by the cipher suite for the SSL policy.
                  set cipher-suite { tls13_sm4_ccm_sm3 | tls13_sm4_gcm_sm3 }

                  By default, no encryption algorithm is configured in the cipher suite for an SSL
                  policy.

         Step 4 Configure an SSL policy and enter the SSL policy view.
                  ssl policy policy-name

                  By default, no SSL policy is configured.

         Step 5 Bind the created SM cipher suite.
                  binding cipher-suite-customization cipher-suite-name

                  By default, no SSL cipher suite is configured.

         Step 6 Set the elliptic curve parameter for the ECDHE algorithm.
                  ecdh group curve-sm2

                  By default, the elliptic curve parameters of the ECDHE algorithm are Curve, Nist,
                  and Brainpool.

                  The SM cipher suite cannot use the default ECDHE elliptic curve algorithm.

         Step 7 Configure a signature algorithm.
                  signature algorithm-list sm2-sm3

                  By default, the following signature algorithms are configured: ed25519, ed448,
                  ecdsa-secp256r1-sha256, ecdsa-secp384r1-sha384, ecdsa-secp521r1-sha512, rsa-
                  pss-pss-sha256, rsa-pss-pss-sha384, rsa-pss-pss-sha512, rsa-pss-rsae-sha256, rsa-
                  pss-rsae-sha384, and rsa-pss-rsae-sha512.

                  The SM cipher suite cannot use the default signature algorithm.

         Step 8 Bind a PKI realm to the SSL policy. After a PKI realm is bound to the SSL policy,
                the SSL policy uses the certificates and CRLs in the PKI realm.
                  pki-domain pki-domain

                  By default, no PKI realm is bound to an SSL policy.

                  The certificates and CRLs in the PKI realm must be SM2-SM3 certificates.

                  ----End




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           249
Security Configuration
Security Configuration                                                             12 SSL Configuration


12.5.5 Applying an SSL Policy
Context
                  SSL is a security protocol, and an SSL policy takes effect only when it is associated
                  with an application.

Procedure
         Step 1 Apply an SSL policy. Table 12-3 lists the major applications in which an SSL policy
                can be used.

                  Table 12-3 Applicable applications of an SSL policy
                   Applicable Application      Example

                   Border Gateway Protocol     See "Configuring SSL/TLS Authentication for BGP" in
                   (BGP)                       CLI Configuration Guide > IP Routing.

                   HyperText Transfer          See "Example for Configuring RESTCONF-based
                   Protocol Secure (HTTPS)     Device Management" in CLI Configuration Guide >
                                               System Management Configuration.

                  ----End




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                             250
Security Configuration
Security Configuration                                                                    13 SSH Configuration




                                              13                   SSH Configuration


Context
                          NOTE

                         In SSH2.0, the symmetric encryption algorithm in CBC mode may encounter plaintext
                         recovery attacks and leak encrypted data. Therefore, the CBC mode is not recommended for
                         data encryption in SSH2.0.


