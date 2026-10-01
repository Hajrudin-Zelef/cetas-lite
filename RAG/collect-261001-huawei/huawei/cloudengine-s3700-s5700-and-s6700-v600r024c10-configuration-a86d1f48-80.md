---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-80
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [9759, 9879]
sha256: 9f1576170763483ea97cc3c180a61ab626229527acf6fb33415e029f907586ab
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  After the configurations are complete, run the display pki cert-req command to
                  view content of the certificate request file.
                  [DeviceA] display pki cert-req filename cer_req
                  Certificate Request:
                    Data:
                        Version: 1 (0x0)
                        Subject: C=cn, ST=jiangsu, O=huawei, OU=info, CN=hello
                        Subject Public Key Info:
                           Public Key Algorithm: rsaEncryption
                              RSA Public-Key: (3072 bits)
                              Modulus:
                                 00:a2:db:e3:30:17:8e:f6:2d:2e:64:15:46:51:ad:
                                 70:86:dd:32:c4:bb:6b:58:3a:8c:5f:a0:06:a1:e1:
                                 56:2e:a4:eb:7e:12:06:05:04:28:b2:6d:64:7a:9c:
                                 4f:85:24:c1:aa:b8:99:dc:e9:bb:c4:1e:e2:9d:a0:
                                 18:51:1f:ad:b5:2f:60:18:06:8b:c1:cc:6f:32:58:
                                 f2:21:2c:16:e8:29:c2:a8:c5:aa:9d:6c:1e:ca:14:
                                 fc:7a:e9:bc:07:91:ce:ed:a0:c0:52:d9:0c:e9:ba:
                                 9b:64:43:e0:9a:3f:c5:d1:2c:86:36:96:6b:4b:4f:
                                 d4:df:05:d0:4b:41:2c:ec:0a:d7:0e:45:83:ed:cd:
                                 07:78:40:ed:d5:3d:7f:fe:0f:08:90:04:2e:ac:e5:
                                 42:b9:81:ea:ec:77:e2:cc:04:6e:e4:63:9f:69:ed:
                                 60:06:5e:c7:e8:bf:30:57:6a:5d:e0:46:68:d3:ee:
                                 b0:da:47:24:e3:b6:a5:f3:20:d8:5a:75:92:70:c2:
                                 a9:a6:97:07:07:0d:1c:94:9a:03:6f:f7:8c:db:6f:
                                 b7:06:de:51:50:9e:71:fd:86:f3:b5:c9:99:05:bf:
                                 f1:10:20:28:d3:a6:29:3d:e0:f4:a7:ba:1e:27:85:
                                 a9:66:fc:a9:90:49:f0:35:f7:d9:6d:06:a2:43:3f:
                                 18:87
                              Exponent: 65537 (0x10001)
                        Attributes:
                        Requested Extensions:
                           X509v3 Key Usage:
                              Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment
                           X509v3 Subject Alternative Name:
                              IP Address:10.2.0.2, DNS:test.abc.com, email:user@test.abc.com
                    Signature Algorithm: sha256WithRSAEncryption
                         0e:0a:a5:b7:d5:54:11:10:c4:ea:ff:77:da:f9:24:4b:a9:98:
                         a1:75:36:08:10:59:60:fa:1a:30:70:2c:b7:f6:5f:5e:31:b7:
                         55:a5:7a:26:e5:af:4a:cd:83:c5:f3:90:f3:b9:d5:f9:0a:6d:
                         6e:8f:25:b4:ed:95:9c:75:a5:d7:b6:25:fc:8d:39:89:fb:af:
                         37:fc:01:7b:09:07:9c:96:7c:fa:28:6d:e2:11:49:a7:95:94:
                         ed:26:5b:ca:f8:98:b0:e7:64:7e:dd:2d:75:ff:89:03:b7:0a:
                         92:53:25:d4:a1:23:b9:5c:eb:5b:29:1d:8a:92:8f:36:68:7b:
                         77:32:bc:48:92:48:84:fa:87:5a:d7:2e:3e:be:d5:6b:e4:df:
                         b1:f2:02:35:91:6a:eb:cd:fc:5a:ea:37:85:6c:12:74:5f:a5:
                         5c:c0:05:09:cd:34:59:0d:c6:c8:75:ca:1c:18:d6:48:e5:4b:


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        182
Security Configuration
Security Configuration                                                                                   10 PKI Configuration

                         e7:8e:e3:ff:25:99:0f:2e:a8:b4:c5:8e:4d:8f:dd:64:c5:1f:
                         61:3c:58:21:4f:d5:35:ba:c8:8e:5f:76:41:9f:27:41:0a:94:
                         59:2c:59:25:2d:de:60:5c:92:07:ac:8a:a5:7a:ba:75:af:2c:
                         82:5f:bb:55:a8:48:49:54:0f:99:54:af:8d:12:4d:4b:7d:8b:
                         95:28:ce:dc

         Step 5 Transfer the certificate request file to the CA server in out-of-band mode, for
                example, disk or email, to apply for a local certificate.
                  When the local certificate is successfully registered, download the local certificate
                  abc_local.cer also in out-of-band mode. After the download is complete, you can
                  import the file to the flash:/pki/public directory of the device using a file transfer
                  protocol.
         Step 6 Install the local certificate.
                  After the certificate is imported, the abc_local.cer file in the flash:/pki/public
                  directory is deleted by default. To keep it, select N as prompted.
                  [DeviceA] pki import-certificate local realm abc pem filename abc_local.cer
                   Info: Succeeded in importing the certificate.
                   Warning: The file in the flash will be deleted. Please select 'N' if you want to keep it. Please select [Y/N]:y
                   Info: Delete Success.

                  ----End

Verifying the Configuration
                  After the local certificate is installed, the devices at both ends can use it to protect
                  communication data.
                  [DeviceA] display pki certificate local filename abc_local.cer
                  Info: It will take a few seconds or more to collect data for displaying. Please wait a moment.
                  Total Number: 1

                  Certificate:
                    Data:
                        Version: 3 (0x2)
                        Serial Number: 8372560407419635446 (0x74314f54b0bf46f6)
                        Signature Algorithm: sha256WithRSAEncryption
                        Issuer: CN=HUAWEI BRAS CA, O=HUAWEI BRAS, C=AT
                        Validity
                           Not Before: Sep 12 22:18:27 2022 GMT
                           Not After : Sep 7 22:18:27 2042 GMT
                        Subject: C=cn, ST=jiangsu, O=huawei, OU=info, CN=hello
                        Subject Public Key Info:
                           Public Key Algorithm: rsaEncryption
                               RSA Public-Key: (3072 bits)
                               Modulus:
                                 00:c8:4f:09:9d:6a:53:95:6d:98:fa:22:f4:7c:5e:
                                 f7:4b:08:3b:d2:19:3b:2d:4c:6c:0d:5f:b7:a2:91:
                                 e8:99:de:91:12:df:3d:f5:c4:89:00:30:e7:7c:a6:
                                 7a:03:18:1e:31:6a:65:34:05:cb:8a:29:f8:65:49:
                                 7c:bd:81:cd:93:8d:be:63:e5:87:99:5d:28:6f:b6:
                                 5c:c6:5c:4e:85:dc:26:26:db:a9:81:1a:19:b4:c4:
                                 72:b7:8f:01:8d:55:8c:a0:58:cd:ef:d2:bd:d2:04:
                                 5c:62:ab:3a:c5:71:d8:46:68:db:30:11:9b:48:46:
                                 f7:5a:f7:70:a9:bf:ce:df:67:50:31:6c:c5:b3:f7:
                                 0c:73:74:33:94:69:18:5b:57:74:5b:6b:49:bf:15:
                                 05:17:01:9f:d0:13:71:c0:fe:45:13:07:2d:95:42:
                                 55:e8:9e:77:e8:4e:f8:80:42:97:4f:26:78:a9:81:
                                 61:8e:d3:ac:e8:5e:e0:61:37:84:f4:82:fa:8a:f9:
                                 08:df:c3:70:50:9a:8e:3b:78:a1:f2:5d:3d:0b:fb:
                                 fa:f4:67:ec:31:35:ff:4a:70:29:86:8c:a8:e2:46:
                                 97:39:f7:58:0e:9e:ff:26:f1:7f:10:6b:68:33:f3:
                                 7e:fd:ce:f3:a2:b1:b5:a4:81:88:52:2f:82:e0:28:


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                    183
Security Configuration
Security Configuration                                                                           10 PKI Configuration

