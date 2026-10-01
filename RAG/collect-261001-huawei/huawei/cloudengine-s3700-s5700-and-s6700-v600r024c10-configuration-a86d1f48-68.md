---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-68
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [8232, 8364]
sha256: b6621d184cb5788c5728f0ffbb5dc23615d9ad5fe38e00ac9658330fc6d99ade
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   SM2      Elliptic     ECC encryption and      RSA, ECC          Yes, it has been
                            curve        decryption,                               incorporated
                            asymmetric   signature                                 into ISO
                            cryptograp   verification, and key                     international
                            hy           exchange                                  standards.

                   SM3      Hashing      Digital signature       SHA-256           Yes, it has been
                                         and verification in                       incorporated
                                         commercial                                into ISO
                                         cryptography                              international
                                         applications,                             standards.
                                         generation and
                                         verification of
                                         message
                                         authentication
                                         codes, and
                                         generation of
                                         random numbers

                   SM4      Block        Block encryption        DES, AES          Yes
                            cipher       and decryption. The
                                         structure about
                                         encryption and
                                         decryption is same.
                                         The only difference
                                         is that the
                                         decryption key is the
                                         reverse of the
                                         encryption key.

                   SM7      Block        Block cipher and        -                 No, it exists in
                            cipher       decipher                                  the hardware
                                                                                   only in the form
                                                                                   of an IP core.

                   SM9      ID-based     Signature               -                 Yes, it has been
                            asymmetric   verification, key                         incorporated
                            cryptograp   exchange, as well as                      into ISO
                            hy           key encapsulation,                        international
                                         encryption, and                           standards.
                                         decryption

                   ZUC      Stream       ZUC stream cipher       EEA3 & EIA3       Yes
                            cipher       algorithm and
                                         symmetric
                                         encryption
                                         algorithm

                   SSF33    Block        Block cipher and        -                 No
                            cipher       decipher



Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                           153
Security Configuration
Security Configuration                                                              10 PKI Configuration




10.2.1.2 Digital Envelope and Digital Signature

Digital Envelope
                  A digital envelope contains the data that combines the symmetric key encrypted
                  using the receiver's public key and the data encrypted using the symmetric key.
                  Upon receiving a digital envelope, the receiver uses its own private key to decrypt
                  the digital envelope and obtains the symmetric key. The digital envelope has the
                  advantages of both symmetric key cryptography and public key cryptography. It
                  speeds up key distribution and encryption, while improving key security,
                  extensibility, and as efficiency.

                  Figure 10-3 shows the encryption and decryption process for a digital envelope.


                  Figure 10-3 Digital envelope encryption and decryption process




                  Assume that user A has obtained the public key of user B. The encryption and
                  decryption process is as follows:

                  1.     User A uses a symmetric key to encrypt clear text, generating ciphertext.
                  2.     User A uses the public key of user B to encrypt the symmetric key, generating
                         a digital envelope.
                  3.     User A sends the digital envelope to user B.
                  4.     User B uses its own private key to decrypt the digital envelope, obtaining the
                         symmetric key.
                  5.     User B uses the symmetric key to decrypt the ciphertext, obtaining the
                         original clear text.

                  However, the following vulnerability should be noted regarding the digital
                  envelope: An attacker may intercept information from user A, use its own
                  symmetric key to encrypt forged information, use the public key of user B to
                  encrypt its own symmetric key, and send the information to user B. User B then
                  decrypts and considers this information to have been sent from user A. To address
                  this problem, the digital signature is utilized to ensure the credentials of the
                  sender.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           154
Security Configuration
Security Configuration                                                               10 PKI Configuration


Digital Signature
                  A digital envelope cannot determine if the information received actually
                  originated from the sender. Instead, a digital signature is utilized to verify the
                  identity of the sender as well as whether the information received has been
                  tampered with.

                  A digital signature is generated by the sender by encrypting the digital fingerprint
                  using its own private key. The receiver then uses the sender's public key to decrypt
                  the digital signature and obtain the digital fingerprint.

                  A digital fingerprint, also known as the information digest, is generated by the
                  sender using the hash algorithm on the clear text. The sender transmits both the
                  digital fingerprint and clear text to the receiver, which also performs hash on the
                  clear text to generate a digital fingerprint. If the two fingerprints match, the
                  receiver determines that the information has not been tampered with.

                  Figure 10-4 shows the encryption and decryption process for a digital signature.

                  Figure 10-4 Digital signature encryption and decryption process




                  Assume that user A has obtained the public key of user B. The encryption and
                  decryption process is as follows:

