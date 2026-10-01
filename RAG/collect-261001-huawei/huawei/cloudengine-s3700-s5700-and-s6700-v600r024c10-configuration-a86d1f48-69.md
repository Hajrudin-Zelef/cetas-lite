---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-69
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [8365, 8484]
sha256: 0fe3413d7ef9e6395513a0e6d366385ffe14b47ecde67f9f17f769172bbbc6c7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  1.     User A uses the public key of user B to encrypt clear text, generating
                         ciphertext.
                  2.     User A performs hash on the clear text, generating a digital fingerprint.
                  3.     User A uses its own private key to encrypt the digital fingerprint, generating a
                         digital signature.
                  4.     User A sends both the ciphertext and digital signature to user B.
                  5.     User B uses the public key of user A to decrypt the digital signature, obtaining
                         the digital fingerprint.
                  6.     After receiving the ciphertext from user A, user B uses its own private key for
                         decryption, obtaining the original clear text.
                  7.     User B performs hash on the clear text, generating a digital fingerprint.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             155
Security Configuration
Security Configuration                                                               10 PKI Configuration


                  8.     User B compares the generated fingerprint with that received from user A. If
                         the two fingerprints match, user B accepts the clear text. Otherwise, user B
                         discards it.
                  The digital signature proves that information has not been tampered with and
                  verifies the sender's identity, and can be used in conjunction with the digital
                  envelope.
                  Despite these advanced security features, a vulnerability still exists. If an attacker
                  modifies user B's public key, user A will obtain the attacker's public key. The
                  attacker can obtain information sent from user B to user A, sign forged
                  information using its own private key, encrypts forged information using user A's
                  public key, and sends this forged information to user A. Upon receipt, user A
                  decrypts the information, verifies that the information has not been tampered
                  with, and considers it to have originated from user B. However, the digital
                  certificate can address this vulnerability.

10.2.1.3 Digital Certificate
                  A digital signature cannot determine whether a public key belongs to a specific
                  owner, as any entity can generate public and private keys. Therefore, a secure and
                  reliable carrier is required to exchange public keys. This carrier is a digital
                  certificate.
                  A digital certificate is a digitally signed file issued by a CA, containing the owner's
                  public key and identity information.
                  A digital certificate, which is similar to an electronic copy of a passport or an ID
                  card, is typically used for identity verification on the network. It ensures that one
                  public key is possessed by only one owner.

Digital Certificate Structure
                  A simplest digital certificate contains mandatory information such as public key,
                  name, and digital signature of the CA. In most cases, the certificate also includes
                  information such as the public key validity period, issuer name (CA), and
                  certificate serial number. The certificate structure complies with X.509 v3. Figure
                  10-5 shows the typical structure of a digital certificate.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             156
Security Configuration
Security Configuration                                                                  10 PKI Configuration


                  Figure 10-5 Digital certificate structure diagram




                  The meaning of each field in the digital certificate is as follows:

                  ●      Version: version of X.509. Generally, the v3 (0x2) is used.
                  ●      Serial Number: a positive and unique integer assigned by the issuer to the
                         certificate. Each certificate is uniquely identified by the issuer name and the
                         serial number.
                  ●      Signature Algorithm: signature algorithm used by the issuer to sign the
                         certificate.
                  ●      Issuer: name of the device that has issued a certificate. It must be the same as
                         the subject name in the digital certificate. Generally, the issuer name is the CA
                         server's name.
                  ●      Validity: time interval during which a digital certificate is valid, including the
                         start and end dates. The expired certificates are invalid.
                  ●      Subject: name of the entity that possesses a digital certificate. In a self-signed
                         certificate, the issuer name is the same as the subject name.
                  ●      Subject Public Key Info: public key and the algorithm with which the key is
                         generated.
                  ●      Extensions: a sequence of optional fields such as certificate usage, CRL
                         address (URL), and OCSP server URL.
                  ●      Signature: signature signed on a digital certificate by the issuer using the
                         private key.

                  The process of generating a certificate signature is as follows: The CA first uses a
                  cryptographic hash algorithm to generate digest information of the certificate, and
                  then uses a public-key cryptography algorithm and a private key of the CA to
                  encrypt the digest information and finally generate a signature. These operations
                  are performed on the CA before the certificate is issued.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               157
Security Configuration
Security Configuration                                                                 10 PKI Configuration


Digital Certificate Types
                  There are three types of certificates, as described in Table 10-2.

                  Table 10-2 Certificate types
                   Type                          Description                    Description

                   CA certificate                CA's own certificate. If a     An applicant trusts a CA
                                                 PKI system does not            by verifying its digital
                                                 have a hierarchical CA         signature. Any applicant
                                                 structure, the CA              can obtain the CA's
                                                 certificate is the self-       certificate (including the
                                                 signed certificate. If a PKI   public key) to verify the
                                                 system has a hierarchical      local certificate issued by
                                                 CA structure, the top CA       the CA.
                                                 is the root CA, which
                                                 owns a self-signed
                                                 certificate.

                   Local certificate             A certificate issued by a      -
                                                 CA to the applicant.

