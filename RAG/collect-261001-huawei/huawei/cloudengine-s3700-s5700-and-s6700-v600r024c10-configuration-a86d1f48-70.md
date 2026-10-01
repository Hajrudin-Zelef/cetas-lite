---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-70
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [8485, 8611]
sha256: 727c6c7ced1809be6e4e113e42cf5850063001ae8aa0942e3ac5d774a75e672c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Self-signed certificate       ● A self-signed                A device can generate a
                                                   certificate is issued by     self-signed or unsigned
                                                   a device to itself and       certificate for itself,
                                                   is signed by the initial     which is a simple
                                                   CA on the device.            certificate issuing
                                                   That is, the certificate     function.
                                                   issuer is the same as        The device does not
                                                   the certificate subject.     support lifecycle
                                                   This type of certificate     management (such as
                                                   contains signature           certificate update and
                                                   information, and it          revocation) of its self-
                                                   does not require             signed certificate. To
                                                   signature application.       ensure security of the
                                                 ● An unsigned                  device and certificate,
                                                   certificate, as its          you are advised to
                                                   name implies, is not         replace the self-signed
                                                   signed. It is issued by      certificate with a local
                                                   a device to itself. A        certificate.
                                                   signature needs to be
                                                   obtained from the CA,
                                                   and the certificate
                                                   issuer is the CA.




Certificate Formats
                  Three certificate formats are supported, as described in Table 10-3.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                                   158
Security Configuration
Security Configuration                                                              10 PKI Configuration


                  Table 10-3 Certificate formats
                   Format           Description       Description

                   PKCS#12          Saves             If the file name extension of a certificate
                                    certificate       is .CER or .CRT, use Notepad to open this
                                    files in binary   certificate and check its content to differentiate
                                    format,           the certificate format.
                                    including or      ● If the certificate starts with "-----BEGIN
                                    excluding the       CERTIFICATE-----" and ends with "-----END
                                    private key.        CERTIFICATE-----", the certificate format is
                                    Commonly            PEM.
                                    used file
                                    name              ● If the certificate content is displayed as
                                    extensions          garbled characters, the certificate format is
                                    include .P12        DER.
                                    and .PFX.

                   DER              Saves
                                    certificate
                                    files in binary
                                    format,
                                    excluding the
                                    private key.
                                    Commonly
                                    used file
                                    name
                                    extensions
                                    include .DER, .
                                    CER, and .CRT.

                   PEM              Saves
                                    certificate
                                    files in ASCII
                                    format,
                                    including or
                                    excluding the
                                    private key.
                                    Commonly
                                    used file
                                    name
                                    extensions
                                    include .PEM,
                                    .CER,
                                    and .CRT.



Chinese Cryptographic Algorithm–based Certificates
                  Only one single certificate is typically deployed on a device for both signature and
                  encryption. Both public and private keys are stored on the device. According to the
                  standards, Chinese cryptographic algorithms require a dual-certificate system.
                  Certificates can be classified into signature certificates and encryption certificates
                  based on their purposes and functions. Signature certificates are used only for

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                              159
Security Configuration
Security Configuration                                                                10 PKI Configuration


                  identity verification. Their public and private keys are generated and kept by
                  devices. Digital signature and non-repudiation are used for signature keys. The
                  encryption certificates are used during key negotiation. The public and private
                  keys are generated by the CA. Key encipherment, data encipherment, and key
                  agreement are used for encryption keys. Generally, a device generates a signature
                  key pair, generates a signature certificate request containing only public key
                  information, and sends the signature certificate to the CA. The CA verifies the
                  signature key pair of the device and generates an encryption key pair and an
                  encryption certificate. Figure 10-6 shows the process for the device to obtain the
                  encryption certificate and private key. SM2 key pairs are used for the signature
                  certificate and encryption certificate.

                  Figure 10-6 SM digital envelope encryption and decryption process




                  1.     The CA generates a symmetric key, uses the signature public key of the device
                         to encrypt the symmetric key, and generates the ciphertext of the symmetric
                         key, that is, digital envelope.
                  2.     The CA uses the symmetric key to encrypt the private key of the encryption
                         certificate and generates the ciphertext of the encrypted private key.
                  3.     The CA sends the encryption certificate, signature certificate, ciphertext of the
                         symmetric key, and ciphertext of the encrypted private key to the device.
                  4.     The device uses its signature private key to decrypt the ciphertext of the
                         symmetric key to obtain the symmetric key.
                  5.     The device uses the symmetric key to decrypt the ciphertext of the encrypted
                         private key to obtain the plain text.

