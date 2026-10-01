---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-107
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [13297, 13448]
sha256: e507eb361e7bd53bb288116a1bcdda1289e902cc3bfd6fb18ae8044241738e8b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         ● For enhanced security, you are advised to use more secure certificates. To be specific,
                           RSA/DSA certificates have a key length of 3072 bits or greater, and ECC certificates have
                           a key length of 256 bits or greater. Additionally, the certificate's hash algorithm is
                           SHA-256 or a later version.
                         ● The install feature-software WEAKEA command needs to be run if the key length of
                           the RSA/DSA certificate is less than 2048 bits, the key length of the ECC certificate is less
                           than 256 bits, or the hash algorithm of the certificate is SHA1, SHA-224, MD4, or MD5.

                  ●      Load a PEM digital certificate for the SSL policy.
                         certificate load pem-cert certFile key-pair keyType key-file keyFile auth-code [ cipher authCode ]

                  ●      Load a PEM certificate chain for the SSL policy.
                         certificate load pem-chain certFile key-pair keyType key-file keyFile auth-code [ cipher authCode ]

                  ●      Load a PFX digital certificate for the SSL policy.
                         Format 1:
                         certificate load pfx-cert certFile key-pair keyType key-file keyFile auth-code [ cipher authCode ]

                         Format 2:
                         certificate load pfx-cert certFile key-pair keyType mac [ cipher macCode auth-code cipher
                         authCode ]

         Step 9 (Optional) Load a certificate revocation list (CRL) for the SSL policy.
                  crl load crlType crlFile

                  By default, no CRL is loaded for the SSL policy.

        Step 10 (Optional) Load a trusted-CA file for the SSL policy.

                  By default, no trusted-CA file is loaded for the SSL policy.

                  The trusted-CA file is used to verify validity of the digital certificate sent by the
                  server. A maximum of four trusted-CA files can be loaded for an SSL policy.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              245
Security Configuration
Security Configuration                                                                           12 SSL Configuration


                          NOTE

                         ● Perform this step if the identity of the peer needs to be authenticated.
                         ● For enhanced security, you are advised to use more secure certificates. To be specific,
                           RSA/DSA certificates have a key length of 3072 bits or greater, and ECC certificates have
                           a key length of 256 bits or greater. Additionally, the certificate's hash algorithm is
                           SHA-256 or a later version.
                         ● The install feature-software WEAKEA command needs to be run if the key length of
                           the RSA/DSA certificate is less than 2048 bits, the key length of the ECC certificate is less
                           than 256 bits, or the hash algorithm of the certificate is SHA1, SHA-224, MD4, or MD5.
                  ●      Load an ASN1 trusted-CA file for the SSL policy.
                         trusted-ca load asn1-ca caFile

                  ●      Load a PEM trusted-CA file for the SSL policy.
                         trusted-ca load pem-ca caFile

                  ●      Load a PFX trusted-CA file for the SSL policy.
                         trusted-ca load pfx-ca caFile auth-code [ cipher authCode ]

        Step 11 (Optional) Exclude key exchange algorithms from the cipher suite list.
                  cipher-suite exclude key-exchange { rsa | dhe } *
                  cipher-suite exclude cipher mode cbc
                  cipher-suite exclude hmac sha1

                  By default, an SSL policy does not support the RSA key exchange algorithm, CBC
                  encryption algorithm, or SHA1 digest algorithm.
        Step 12 (Optional) Bind a cipher suite to the SSL policy.
                  binding cipher-suite-customization customization-name

                  By default, no cipher suite is bound to an SSL policy. In this case, all encryption
                  algorithms can be used.
                  The cipher suite to be bound to the SSL policy must have been configured. For
                  details, see 12.5.1 (Optional) Configuring a Cipher Suite for an SSL Policy.
        Step 13 (Optional) Set the certificate expiration alarm threshold and the interval for
                checking certificate expiration alarms.
                  quit
                  ssl certificate alarm-threshold early-alarm time check-interval check-period

                  By default, the certificate expiration alarm threshold is 90 days, and the interval
                  for checking certificate expiration alarms is 24 hours.
        Step 14 (Optional) Enable the SSL renegotiation function.
                  ssl renegotiation enable

                  By default, the SSL renegotiation function is disabled.
                  When SSL is used for data transmission, you can enable the SSL renegotiation
                  function to periodically update the key and algorithm without interrupting the
                  connection, improving communication security.

                  ----End




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       246
Security Configuration
Security Configuration                                                                12 SSL Configuration


12.5.3 Configuring an SSL Policy (Loading a Certificate Using
PKI)

Prerequisites
                  A certificate has been loaded to the PKI realm to be bound. The loaded certificate
                  can be an initial device certificate or a digital certificate applied by a user. For
                  details on how to load a certificate using PKI, see CLI Configuration Guide > PKI
                  Configuration.


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
                  ssl verify version crl-version2 enable
                  ssl verify key-usage enable
                  ssl verify certificate-signature-algorithm enable

                  By default, the digital certificate verification function is disabled.

