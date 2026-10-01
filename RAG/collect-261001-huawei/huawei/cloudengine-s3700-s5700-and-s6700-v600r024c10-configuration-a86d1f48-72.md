---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-72
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [8745, 8836]
sha256: 8b93d0b792262bfbf05168fbe5d709045f987886ee855593353ca032e4f5d974
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  Certificate Update

                  When a certificate expires or if its private key is leaked, it must be replaced by the
                  PKI entity. You can manually apply for a new certificate or configure CMPv2 to
                  implement automatic certificate update.

                  Certificate Revocation

                  In scenarios involving a change of user identity, user information, or public key; or
                  due to user service suspension, the user must revoke the digital certificate, that is,
                  unbind the public key from the user's identity information. A CA provides the
                  certificate revocation function. When a PKI entity revokes its certificate in out-of-
                  band mode, the CA stores the certificate in the CRL database or OCSP server.

10.2.3 PKI Working Mechanism
                  On a PKI network, a PKI entity applies for a local certificate from the CA and the
                  applicant device authenticates the certificate. The processes for offline and online
                  certificate application are different, as described in the following sections.



Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                               163
Security Configuration
Security Configuration                                                                 10 PKI Configuration


Offline Certificate Application Process

                  Figure 10-9 PKI entity's offline application process




                  1.     A PKI entity sends a certificate enrollment request file to the CA in out-of-
                         band mode (for example, by disk or email), requesting the CA to create a
                         certificate.
                  2.     The CA checks the validity of the certificate enrollment request file. If the file
                         is valid, the CA creates a certificate based on this file.
                  3.     The PKI entity obtains the local certificate in out-of-band mode (for example,
                         by disk or email), and downloads the obtained local certificate.
                  4.     The PKI entity installs the local certificate to the device memory.
                  5.     Optional:
                         When PKI entities communicate with one another, they must obtain each
                         other's local certificate and CA certificate.
                  6.     The PKI entity uses CRL or OCSP to check whether the peer's local certificate
                         is valid.
                  7.     The PKI entity uses the public key in the peer's local certificate for encrypted
                         communication only after confirming that the peer's local certificate is valid.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               164
Security Configuration
Security Configuration                                                                 10 PKI Configuration


CMPv2-based Online Certificate Application Process

                  Figure 10-10 PKI entity's online certificate application process




                  1.     A PKI entity sends a certificate enrollment request (including the public key in
                         the RSA key pair and the PKI entity information) to the CA.
                         When a PKI entity applies for a local certificate using CMPv2, the PKI entity
                         can use a signature or message authentication code to send an identity
                         authentication request to the CA.
                         –   Signature: The PKI entity uses the CA certificate's public key to encrypt
                             the certificate enrollment request, and uses the private key corresponding
                             to its external identity certificate (local certificate issued by another CA)
                             for digital signature.
                         –   Message authentication code: The PKI entity uses the CA certificate's
                             public key to encrypt the certificate enrollment request, and the request
                             must contain the message authentication code's reference value and
                             secret value (the values must be the same as those of the CA).
                  2.     After receiving the certificate enrollment request from the PKI entity, the CA
                         verifies the request and issues a certificate.
                         –   Signature: The CA uses its own private key to decrypt the certificate
                             enrollment request, uses the public key of the PKI entity's external
                             identity certificate to decrypt the digital signature, and verifies the digital
                             fingerprint. When the fingerprint is the same as that of the CA, the CA

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               165
Security Configuration
Security Configuration                                                                  10 PKI Configuration


