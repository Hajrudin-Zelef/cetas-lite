---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-87
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10659, 10805]
sha256: 498631075c8651332ffe65e62cff2b74efc5f9c50ae99bf233d15ed458ebe83b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           196
Security Configuration
Security Configuration                                                                              10 PKI Configuration


                  ●      After a CA certificate is obtained and imported to memory, run the display
                         pki certificate ca command to view the content of the certificate.
                         [DeviceA] display pki certificate ca filename cmp_ca1.cer
                          The x509 object type is certificate:
                         Certificate:
                            Data:
                               Version: 3 (0x2)
                               Serial Number: 2 (0x2)
                               Signature Algorithm: sha1WithRSAEncryption
                               Issuer: C=cn, ST=beijing, L=BB, O=BB, OU=BB, CN=BB
                               Validity
                                  Not Before: Aug 15 02:38:27 2011 GMT
                                  Not After : Aug 13 02:38:27 2016 GMT
                               Subject: C=cn, ST=jiangsu, O=huawei, OU=info, CN=hello
                               Subject Public Key Info:
                                  Public Key Algorithm: rsaEncryption
                                      RSA Public-Key: (1024 bit)
                                      Modulus:
                                         00:b7:3e:65:7f:3b:3c:18:b8:87:34:39:76:3c:87:
                                         39:f7:a9:b3:35:9b:e0:e0:5b:c7:4f:3c:bb:fa:dd:
                                         da:93:0b:55:6e:eb:ba:52:c8:86:d1:cf:14:1e:1c:
                                         35:c6:53:68:f3:51:e7:2c:d4:b8:fa:0f:b3:04:ef:
                                         3f:a0:b3:4d:78:c1:26:88:26:15:41:3d:14:7f:67:
                                         3e:2f:35:32:ce:c7:73:73:43:5c:12:d3:0f:a0:ec:
                                         96:ae:55:61:27:32:39:a4:f8:32:a1:68:50:e6:3d:
                                         2b:39:6d:42:e8:09:5d:4f:98:46:6e:fc:80:87:0e:
                                         36:ca:09:7a:ca:2f:dd:ad:d3
                                      Exponent: 65537 (0x10001)
                               X509v3 extensions:
                                  X509v3 Basic Constraints: critical
                                      CA:TRUE
                                  X509v3 Subject Key Identifier:
                                      4F:67:F4:CB:F4:C3:F7:61:2C:BD:FF:1D:D1:29:FD:39:28:9F:3B:8B
                                  X509v3 Key Usage:
                                      Certificate Sign, CRL Sign
                                  Netscape Cert Type:
                                      SSL CA, S/MIME CA, Object Signing CA
                                  Netscape Comment:
                                      xca certificate
                            Signature Algorithm: sha1WithRSAEncryption
                               75:43:24:eb:db:ee:7d:05:30:88:b8:1b:d5:32:ca:51:49:74:
                               04:94:fe:d0:31:29:6f:72:c7:4a:86:ac:2a:4c:45:24:9d:3c:
                               b4:30:b5:d1:43:88:29:f7:b4:88:b8:37:dc:dd:f4:fa:42:34:
                               1c:e6:a5:bc:bb:0b:37:ef:db:8c:b2:b0:bd:97:7f:15:ae:6c:
                               71:1b:ff:f1:90:13:74:a4:1f:7c:f7:4e:80:5b:42:aa:6b:22:
                               2a:cf:04:48:29:20:c0:b2:95:38:11:06:be:76:f0:cb:8d:4a:
                               c6:1a:50:af:31:81:58:ac:14:fe:89:f2:e0:bb:95:3c:94:d0:
                               54:96

                         Pki realm name: -
                         Certificate file name: cmp_ca1.cer
                         Certificate peer name: -


Configuration Scripts
                  DeviceA configuration file
                  #
                  sysname DeviceA
                  #
                  pki entity user01
                   country cn
                   state jiangsu
                   organization huawei
                   organization-unit info
                   common-name hello
                   fqdn user@test.abc.com
                   ip-address 10.2.0.2


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        197
Security Configuration
Security Configuration                                                                              10 PKI Configuration

                   email user@user@test.abc.com
                  #
                  interface 10GE1/0/1
                   ip address 10.2.0.2 255.255.255.0
                  #
                  interface 10GE1/0/2
                   ip address 10.1.0.2 255.255.255.0
                  #
                  pki import-certificate ca filename cmp_cal.cer
                  pki import-certificate local filename cmp_ir.cer
                  pki cmp session cmp
                   cmp-request ca-name "C=cn,ST=beijing,L=SD,O=BB,OU=BB,CN=BB"
                   cmp-request authentication-cert cmp_ir.cer
                   cmp-request entity user01
                   cmp-request server url http://10.3.0.1:8080
                   cmp-request rsa local-key-pair rsa_cmp regenerate
                   cmp-request message-authentication-code 1234 %@%##!!!!!!!!!"!!!!'!!!!*!!!!#~Yt'T`/_H5O<-:ydTz$hk./
                  U,Huq3[u0w8!!!!!!!!!!!!!!!~!!!!h#a6(1U`jWv[fB3ZRI\7~b5jYCD+l0/R)RMFWV,:%@%#
                   certificate auto-update enable
                   certificate update expire-time 80
                  #
                  return



10.8 Configuring a Self-signed Certificate
Context
                  If a device fails to request a local certificate from the CA, it can generate a self-
                  signed certificate. The generated certificate is saved as a file in storage,
                  implementing simple certificate issuing. You can export the certificate and transfer
                  it to another device. A self-signed certificate is issued by a device to itself and is
                  signed by the initial CA on the device. That is, the certificate issuer is the same as
                  the certificate subject. This type of certificate contains signature information, and
                  it does not require signature application.

                          NOTE

                         The device does not support lifecycle management (such as certificate update and
                         revocation) of its self-signed certificate. To ensure security of the device and certificate, you
                         are advised to replace the self-signed certificate with a local certificate.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create a self-signed or unsigned certificate.
                  pki create-certificate [ self-signed ] filename file-name

                  During the configuration, you will be prompted to enter the certificate
                  information, such as PKI entity attributes, the certificate file name, the certificate
                  validity period, and the RSA key length.

                  If the self-signed parameter is specified, a self-signed certificate is created. If this
                  parameter is not specified, an unsigned certificate is created. An unsigned
                  certificate, as its name implies, is not signed. It is issued by a device to itself. A
                  signature needs to be obtained from the CA, and the certificate issuer is the CA.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            198
Security Configuration
Security Configuration                                                                10 PKI Configuration


                  The file format of the created self-signed or unsigned certificate is PEM.

