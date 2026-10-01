---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-81
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [9880, 10022]
sha256: 680c2d6bccec38920bf26ddb43d96cef3c7647efc46e95b1b2d1ab7e3dcc665a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                                d3:f5
                             Exponent: 65537 (0x10001)
                         X509v3 extensions:
                           X509v3 Authority Key Identifier:
                             keyid:33:06:DB:08:3C:3F:61:9B:C4:04:A5:36:8C:FD:34:D3:4C:73:B9:92

                           X509v3 Subject Key Identifier:
                             78:4C:65:05:E7:F6:B5:C8:48:C3:9D:E5:CE:3F:83:3A:62:84:EA:F9
                           X509v3 CRL Distribution Points:

                              Full Name:
                               DirName:CN = pre_crl, OU = CRL, O = HUAWEI BRAS, C = AT

                              Full Name:
                               URI:http://192.168.1.1:8080/allcrl/pre_crl.crl

                              Full Name:
                               URI:ldap://www.jitldap.com:5389/CN=pre_crl,OU=CRL,O=HUAWEI BRAS,C=AT?
                  certificateRevocationList?base?objectclass=cRLDistributionPoint

                    Signature Algorithm: sha256WithRSAEncryption
                       80:34:0d:ea:a0:7f:b8:a8:cb:8b:ae:a9:b3:85:b3:af:b2:1c:
                       15:fc:7e:75:70:be:ff:37:75:6e:67:f8:37:33:ed:5c:5e:5b:
                       3b:13:dc:44:7e:12:b6:85:b3:5c:b9:49:90:6c:96:33:57:a8:
                       f3:c7:c4:04:2d:36:2a:54:fe:52:9a:16:64:66:a0:2e:a6:f1:
                       0e:e0:29:f0:ac:69:d6:8a:f6:0d:43:41:ff:df:fd:06:03:39:
                       75:8c:36:50:99:c3:89:c7:59:8c:65:7c:0c:6b:86:66:f3:a1:
                       b1:6a:b7:43:0b:6d:3f:7d:82:27:45:b0:75:da:95:07:1d:d2:
                       59:78:88:12:67:26:0f:65:fd:4f:05:4c:7c:74:16:4b:7d:ac:
                       f8:a9:d1:2f:d6:57:4a:ad:aa:a3:ac:7c:30:de:6f:cf:3f:b4:
                       d6:c5:84:e1:55:88:a2:40:52:12:5f:08:d8:50:54:ea:e7:c3:
                       43:e2:6e:98:2a:5d:a4:e9:38:06:36:d6:40:25:a2:2e:0f:e1:
                       95:cc:e8:f9:25:37:75:dd:67:0e:b9:0f:a9:5a:83:9c:6b:6c:
                       f6:e1:bc:9d:fc:c1:a7:76:3f:33:81:e9:6d:25:a2:9f:1b:4e:
                       61:f8:a9:12:de:33:02:2a:98:9d:04:a1:87:98:94:c4:11:cc:
                       af:07:1b:68


                  Pki realm name: abc
                  Certificate file name: abc_local.cer
                  Certificate peer name: -


Configuration Scripts
                  #
                  sysname DeviceA
                  #
                  pki entity user01
                   country cn
                   state jiangsu
                   organization huawei
                   organization-unit info
                   common-name hello
                   fqdn test.abc.com
                   ip-address 10.2.0.2
                   email user@test.abc.com
                  #
                  pki realm abc
                   entity user01
                   rsa local-key-pair rsakey
                  #
                  interface 10GE1/0/1
                   ip address 10.2.0.2 255.255.255.0
                  #
                  interface 10GE1/0/2
                   ip address 10.1.0.2 255.255.255.0
                  #
                  return




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     184
Security Configuration
Security Configuration                                                               10 PKI Configuration




10.7 Applying for and Updating Certificates in Online
Mode Using CMPv2

10.7.1 Understanding Online Certificate Application and
Update Using CMPv2
                  If a device can access a Certificate Authority (CA) that supports the Certificate
                  Management Protocol version 2 (CMPv2), the device can apply for and update its
                  local certificate using CMPv2.

Certificate Application
                  Certificate application, also known as certificate enrollment, is a process in which
                  a PKI entity introduces itself to a CA, which then issues it a certificate. You can
                  apply for a local certificate in online mode using CMPv2. As such, the CA creates a
                  certificate for the PKI entity based on the certificate enrollment request. After the
                  local certificate is created, the CA automatically saves it to the flash:/pki/public
                  directory of the device. You can then manually save the local certificate to the
                  device memory.
                  To obtain a local certificate online, you need to configure PKI entity information,
                  configure an RSA key pair, install the CA certificate, apply for the local certificate,
                  and install the local certificate. Figure 10-13 shows the process of applying for a
                  certificate in online mode using CMPv2.

                  Figure 10-13 Process of applying for a certificate in online mode using CMPv2




                  1.     Create a public/private key pair on DeviceA. The public key information is
                         required during certificate application.
                  2.     Create entity information. When applying for a certificate, DeviceA must
                         provide the CA with information that can prove its identity. The entity
                         information represents identity information, including the common name,
                         fully qualified domain name (FQDN), IP address, and email address. The
                         common name is mandatory, while others are optional.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              185
Security Configuration
Security Configuration                                                                   10 PKI Configuration


                  3.     Install a CA certificate. From the installed CA certificate, the PKI entity can
                         obtain the CA's public key required to be carried in a certificate enrollment
                         request.
                  4.     The PKI entity sends a certificate enrollment request to the CA.
                         Applying for a local certificate using CMPv2 covers two scenarios: initial and
                         non-initial local certificate application. In initial local certificate application,
                         an initialization request (IR) is used. After creating a local certificate, the CA
                         returns the CA certificate together with the local certificate. In signature-
                         based non-initial local certificate application, the PKI entity is authenticated
                         based on the signature and the CA returns only the local certificate, not the
                         CA certificate.
                         –   Initial local certificate application using an IR
                             A CMPv2 server can use either of the following methods to authenticate
                             a PKI entity when a local certificate is requested for the first time:

                             ▪    Message authentication code: The device and CMPv2 server share a
                                  reference value and secret value of the message authentication code.
                                  When applying for the local certificate for the first time, the device
                                  adds these two values to a certificate enrollment request and sends
                                  the request to the CMPv2 server. Then the CMPv2 server validates
                                  the two values to authenticate the device.

