---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-67
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [8104, 8231]
sha256: ace973b51ecce7b36e4e8c996df3818b6223c9082b51adad57d924018694b9ac
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

10.2.1.1 Cryptography
                  Cryptography is the basis for secure information transmission on networks. It relies
                  on mathematical techniques to transform cleartext into unreadable ciphertext,
                  ensuring the integrity and confidentiality of communication. Encryption
                  algorithms are classified into reversible encryption algorithms and irreversible
                  encryption algorithms. Reversible encryption algorithms are further classified into
                  symmetric encryption algorithms and asymmetric encryption algorithms.

Irreversible Encryption
                  Irreversible encryption is usually used as the basis of encryption to generate a
                  message digest. Irreversible encryption compares two encrypted ciphertexts to
                  check whether data is modified during transmission. The original non-cipher text
                  information cannot be inferred from the ciphertext. Common irreversible
                  encryption algorithms include Message Digest algorithm 5 (MD5), secure hash
                  algorithm (SHA), and Hash-based Message Authentication Code (HMAC) that
                  requires a key. SHA includes SHA1, SHA256, and SHA512.

Symmetric Key Cryptography
                  Symmetric key cryptography, also known as shared key cryptography, utilizes the
                  same key to encrypt and decrypt data. It includes stream ciphers and block
                  ciphers.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            150
Security Configuration
Security Configuration                                                              10 PKI Configuration


                  Figure 10-1 shows the symmetric key encryption and decryption process.

                  Figure 10-1 Symmetric key encryption and decryption process




                  Users A and B have negotiated the symmetric key. The encryption and decryption
                  process is as follows:
                  1.     User A uses the symmetric key to encrypt data and sends the encrypted data
                         to user B.
                  2.     User B decrypts the data using the symmetric key and gets the original data.
                         Symmetric key cryptography features high efficiency, simple algorithm, and
                         low cost. It is suitable for encrypting a large amount of data. However, it is
                         difficult to implement because the two parties must exchange their keys
                         securely before communication. Besides, it is difficult to expand because each
                         pair of communicating parties needs to negotiate keys, and n users need to
                         negotiate n x (n – 1)/2 different keys.
                  The algorithms commonly used in symmetric key cryptography include Data
                  Encryption Standard (DES), Triple Data Encryption Standard (3DES), and Advanced
                  Encryption Standard (AES). These algorithms are block cryptographic algorithms.

Public Key Cryptography
                  Public key cryptography, also known as asymmetric key cryptography, employs a
                  pair of different keys — a public key and a private key — for data encryption and
                  decryption. The public key is freely accessible, whereas the private key is kept
                  confidential by its sole owner.
                  Public key cryptography mitigates the security risks associated with sharing and
                  managing a symmetric key between parties. In an asymmetric key pair, the public
                  key is used to encrypt data and the private key is used to decrypt data. The sender
                  uses the public key of the receiver to encrypt the data, and the receiver uses its
                  own private key to decrypt data. The receiver's private key is only known by the
                  receiver, so the data is secure.
                  Figure 10-2 shows the public key encryption and decryption process.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           151
Security Configuration
Security Configuration                                                              10 PKI Configuration


                  Figure 10-2 Public key encryption and decryption process




                  Assume that user A has the public key of user B. The encryption and decryption
                  process is as follows:
                  1.     User A uses the public key of user B to encrypt data and sends the encrypted
                         data to user B.
                  2.     User B decrypts the data using its own private key and gets the original data.

                  Attackers cannot use one key in a key pair to figure out the other key. The data
                  encrypted by a public key can only be decrypted by the private key of the same
                  user. However, the public key cryptography requires a long time to encrypt a large
                  amount of data, and the encrypted data is too long, consuming much bandwidth.

                  Public key cryptography is suitable for encrypting sensitive information such as
                  keys and identities to provide higher security.

                  The algorithms commonly used in public key cryptography include Diffie-Hellman
                  (DH), Ron Rivest, Adi Shamirh, LenAdleman (RSA), Digital Signature Algorithm
                  (DSA), and Elliptic Curve Cryptography (ECC).


Chinese Cryptographic Algorithms
                  Chinese cryptographic algorithms refer to the commercial cryptography approved
                  by China's State Cryptography Administration and are used to encrypt information
                  that does not involve state secrets. To ensure the security of commercial
                  cryptography, China has formulated a series of cryptographic standards, including
                  symmetric encryption algorithms, elliptic curve asymmetric encryption algorithms,
                  and hash algorithms. The following table lists the algorithms contained in the
                  Chinese cryptographic algorithm suite. The SM2, SM3, and SM4 algorithms are
                  mainly used in PKI.

                  Table 10-1 Overview of the Chinese cryptographic algorithm suite

                   Algori     Type          Function               Corresponding      Disclosed or
                   thm                                             International      Not
                                                                   Algorithm

                   SM1        Symmetric     Block cipher           AES128             No, it exists in
                              encryption    algorithm                                 the hardware
                                                                                      only in the form
                                                                                      of an IP core.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            152
Security Configuration
Security Configuration                                                           10 PKI Configuration


                   Algori   Type         Function                Corresponding     Disclosed or
                   thm                                           International     Not
                                                                 Algorithm

