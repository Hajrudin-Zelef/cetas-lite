---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-73
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [8837, 8966]
sha256: 5ed17baed320b166c02fcb28bc6486f956735a4134745a8f84efc291a1743720
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                             verifies the PKI entity's identity information, accepting the application
                             and issuing a local certificate to the PKI entity. The CA uses the public key
                             of the PKI entity's external identity certificate to encrypt the local
                             certificate, uses its own private key to digitally sign the local certificate,
                             and sends the local certificate to the PKI entity. At the same time, the CA
                             also sends the local certificate to the certificate/CRL database.
                         –   Message authentication code: The CA uses its own private key to decrypt
                             the certificate enrollment request, and verifies the reference value and
                             secret value of the message authentication code. When the reference
                             value and secret value are the same as those of the CA, the CA verifies
                             the PKI entity's identity information, accepting the application and issuing
                             a local certificate to the PKI entity. The CA then uses the PKI entity's
                             public key to encrypt the local certificate, and issues the local certificate
                             to the PKI entity. At the same time, the CA also sends the local certificate
                             to the certificate/CRL database.
                  3.     After receiving the certificate information from the CA, the PKI entity installs
                         the local certificate to the device memory.
                         –   Signature: The PKI entity uses the private key corresponding to its
                             external identity certificate to decrypt the local certificate, uses the CA
                             certificate's public key to decrypt the digital signature, and verifies the
                             digital fingerprint. If the digital fingerprint is the same as its local one,
                             the PKI entity accepts and installs the local certificate to the device
                             memory.
                         –   Message authentication code: The PKI entity uses its own private key to
                             decrypt the certificate, and verifies the reference value and secret value of
                             the message authentication code. If the reference value and secret value
                             are the same as its local ones, the PKI entity accepts and installs the local
                             certificate to the device memory.
                  4.     Optional:
                         When PKI entities communicate with one another, they must obtain each
                         other's local certificate and CA certificate.
                  5.     The PKI entity uses CRL or OCSP to check whether the peer's local certificate
                         is valid.
                  6.     The PKI entity uses the public key in the peer's local certificate for encrypted
                         communication only after confirming that the peer's local certificate is valid.

                  If an RA is available in a PKI system, the PKI entities also need to download the
                  RA's certificate. The RA verifies local certificate enrollment requests from PKI
                  entities, and forwards the requests to the CA after verifications are passed.


10.3 Configuration Precautions for PKI

10.4 Default Settings for PKI
                  Table 10-4 describes the default settings for PKI.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   166
Security Configuration
Security Configuration                                                               10 PKI Configuration


                  Table 10-4 Default settings for PKI
                   Parameter                                   Default Setting

                   PKI realm                                   The device has a PKI realm named
                                                               default by default. This realm can be
                                                               modified but cannot be deleted.

                   RSA key pair                                By default, the device has an RSA key
                                                               pair file named default.

                   Format of saved certificate request,        PEM
                   certificate, and CRL

                   Certificate check method                    CRL

                   Number of days in advance users are         90 days
                   notified that the local or CA certificate
                   is about to expire

                   Expiration check period of the local        24 hours
                   certificate, CA certificate, or CRL

                   CRL expiration check                        Enabled

                   Prewarning percentage of remaining          5%
                   CRL validity period




10.5 Preconfiguration for Certificate Application
10.5.1 Configuring an RSA/SM2/ECC Key Pair
Context
                  Local certificates are signed and issued by the CA. A local certificate is a bundle of
                  a public key and a PKI entity. As such, before applying for a local certificate, you
                  must configure the RSA/SM2/ECC key pair to generate public and private keys. The
                  public key is sent by the PKI entity to the CA, and the peer uses this key to encrypt
                  clear text. In contrast, the private key is retained by the PKI entity, which uses it to
                  digitally sign and decrypt the peer's ciphertext.
                  You can configure an RSA/SM2/ECC key pair using either of the following
                  methods:
                  ●      Create an RSA/SM2/ECC key pair.
                         You can directly create an RSA/SM2/ECC key pair on the device, without the
                         need to import it to the device memory. During RSA/SM2/ECC key pair
                         creation, the system prompts you to enter the number of bits for the public
                         key, which ranges from 2048 to 4096. A longer public key indicates higher
                         security but slower calculation.
                  ●      Import an RSA/SM2/ECC key pair.
                         To use the RSA/SM2/ECC key pair generated by another PKI entity, upload the
                         RSA/SM2/ECC key pair to the device through FTP or SFTP and then import it

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             167
Security Configuration
Security Configuration                                                             10 PKI Configuration


                         to the device memory. Otherwise, the RSA/SM2/ECC key pair does not take
                         effect on the device.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure an RSA/SM2/ECC key pair using either of the following methods.

                  Table 10-5
                   Operatio      Command                          Description
                   n

                   Create an     pki rsa local-key-pair create    The created RSA key pair can be
                   RSA/SM2/      key-name [ modulus               exported only when the
                   ECC key       modulus-size ] [ exportable ]    exportable parameter is specified
                   pair.                                          in the command.

