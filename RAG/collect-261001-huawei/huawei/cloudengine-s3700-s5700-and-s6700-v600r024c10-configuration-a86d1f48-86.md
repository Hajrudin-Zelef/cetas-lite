---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-86
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [10532, 10658]
sha256: b9abea5a0ba5aa142f6d9b9b00dd61f225fa9e080d4e7f50c5ef3c0d8f5500e1
---

                  # Specify the PKI entity name referenced by the CMP session.
                  [DeviceA-pki-cmp-session-cmp] cmp-request entity user01

                  # Configure a CA name, for example, C=cn,ST=beijing,L=SD,O=BB,OU=BB,CN=BB.

                          NOTE

                         The field order in the CA name must be the same as that in the CA certificate; otherwise,
                         the server considers the CA name invalid.
                  [DeviceA-pki-cmp-session-cmp] cmp-request ca-name "C=cn,ST=beijing,L=SD,O=BB,OU=BB,CN=BB"

                  # Configure the URL for certificate application.
                  [DeviceA-pki-cmp-session-cmp] cmp-request server url http://10.3.0.1:8080

                  # Specify the RSA key pair used for certificate application and configure the device
                  to update the RSA key pair together with the certificate.
                  [DeviceA-pki-cmp-session-cmp] cmp-request rsa local-key-pair rsa_cmp regenerate

                  # Use the MAC for initial certificate application. Set the MAC reference value to
                  1234 and MAC secret value to Huawei@RSA1234.
                  [DeviceA-pki-cmp-session-cmp] cmp-request message-authentication-code 1234 Huawei@RSA1234
                  [DeviceA-pki-cmp-session-cmp] quit
                  [DeviceA] pki cmp initial-request session cmp

                  The CA and local certificates obtained are named cmp_ca1.cer and cmp_ir.cer
                  respectively, and are stored in the device's storage medium.
         Step 5 Install certificates.
                  After the certificates are imported, the cmp_ca1.cer and cmp_ir.cer files are
                  deleted from the storage medium by default. To keep them, select N as prompted.
                  # Import the CA certificate to memory.
                  [DeviceA] pki import-certificate ca filename cmp_ca1.cer
                   The CA's Subject is /C=cn/ST=beijing/L=BB/O=BB/OU=BB/CN=BB
                   The CA's fingerprint is:
                     SHA1 fingerprint:2C:2B:C0:31:66:A6:95:A0:7A:AC:EF:3D:37:1C:9A:4D:01:BA:09:4D
                     SHA256
                  fingerprint:CA:FC:6B:94:53:E9:E3:D7:D3:E1:F4:75:3F:DB:C4:0F:0A:B9:F1:AD:03:0B:A8:0D:EE:73:4A:83:54:EF:1F:81
                   Is the fingerprint correct?(Y/N):y
                   Info: Succeeded in importing the certificate.
                   Warning: The file in the flash will be deleted. Please select 'N' if you want to keep it. Please select [Y/N]:y
                   Info: Delete Success.

                  # Import the local certificate to memory.
                  [DeviceA] pki import-certificate local filename cmp_ir.cer
                   Info: Succeeded in importing the certificate.
                   Warning: The file in the flash will be deleted. Please select 'N' if you want to keep it. Please select [Y/N]:y
                   Info: Delete Success.

         Step 6 Configure automatic certificate update.
                  # In the CMP session view, enable automatic certificate update using CMPv2.
                  [DeviceA] pki cmp session cmp
                  [DeviceA-pki-cmp-session-cmp] cmp-request authentication-cert cmp_ir.cer


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                    195
Security Configuration
Security Configuration                                                                                 10 PKI Configuration

                  [DeviceA-pki-cmp-session-cmp] certificate auto-update enable
                  [DeviceA-pki-cmp-session-cmp] quit

                  # In the CMP session view, set the automatic certificate update time to 80% of the
                  current certificate validity period.
                  [DeviceA] pki cmp session cmp
                  [DeviceA-pki-cmp-session-cmp] certificate update expire-time 80
                  [DeviceA-pki-cmp-session-cmp] quit

                  ----End

Verifying the Configuration
                  ●      After a local certificate is obtained and imported to memory, run the display
                         pki certificate local command to view the content of the certificate.
                         [DeviceA] display pki certificate local filename cmp_ir.cer
                          The x509_obj type is Cert:
                         Certificate:
                            Data:
                               Version: 3 (0x2)
                               Serial Number: 1144733510 (0x443b3f46)
                               Signature Algorithm: sha1WithRSAEncryption
                               Issuer: C=cn, ST=beijing, L=BB, O=BB, OU=BB, CN=BB
                               Validity
                                  Not Before: Jun 12 09:33:10 2012 GMT
                                  Not After : Aug 13 02:38:27 2016 GMT
                               Subject: C=cn, ST=jiangsu, O=huawei, OU=info, CN=hello
                               Subject Public Key Info:
                                  Public Key Algorithm: rsaEncryption
                                      RSA Public-Key: (3072 bit)
                                      Modulus:
                                         00:d3:12:fe:57:48:c6:a5:10:12:e9:2f:f9:2a:ff:
                                         7b:2a:d8:45:69:11:c4:85:30:c4:9a:4d:0f:ad:58:
                                         e7:56:cd:5c:f0:18:e1:c3:6d:44:c2:c3:5e:64:22:
                                         d1:28:c9:c3:37:3c:34:ed:28:04:7f:62:9e:8b:94:
                                         af:bc:72:de:f6:72:7f:e4:d8:45:31:fd:f9:ac:ce:
                                         5a:b9:c7:1b:23:53:00:28:a6:3b:f5:61:69:5d:ab:
                                         67:cb:bb:e8:96:2f:ce:ab:2c:6b:91:5b:26:91:86:
                                         8f:80:a9:b0:66:c1:16:3d:31:55:a2:d4:b5:5a:af:
                                         85:88:6e:99:f8:f8:53:58:77:26:91:ed:0e:94:ad:
                                         c5:8d:53:67:67:55:08:8d:90:38:e0:5e:96:37:b9:
                                         64:0e:36:e7:cf:9a:d2:77:e4:b0:24:05:a6:eb:03:
                                         6e:ff:f7:ab:be:93:9e:8c:66:7d:31:66:be:6d:c8:
                                         f3:17:9d:86:19:88:21:2d:d9:69:86:5f:b2:55:a4:
                                         db:bc:d7:d0:6b:ac:66:ac:e4:63:9c:66:79:9c:42:
                                         5c:83:b8:9e:4b:6e:67:85:a2:47:19:f1:5c:c0:3c:
                                         c9:a3:47:02:a8:53:69:59:9e:d9:c7:5e:90:83:8d:
                                         ac:cd:21:3c:d5:31:39:49:84:e6:f8:f4:e0:44:dd:
                                         5d:7b
                                      Exponent: 65537 (0x10001)
                               X509v3 extensions:
                                  X509v3 Subject Alternative Name:
                                      IP Address:10.2.0.2, DNS:test.abc.com, email:user@test.abc.com
                            Signature Algorithm: sha1WithRSAEncryption
                               53:d5:79:31:7b:40:52:aa:ec:a9:35:ed:07:62:32:c4:ce:22:
                               d3:37:0e:83:0c:4c:fa:61:dd:8c:db:a8:d3:fd:6a:ca:0e:3c:
                               91:2c:91:ab:92:31:34:b5:87:1e:30:a4:ff:94:9c:d2:71:3c:
                               6b:1f:4f:be:a7:20:f2:e1:c2:ad:71:8b:c2:79:0f:50:1f:3c:
                               f9:87:df:1d:ee:3d:38:8c:f3:30:b7:3b:00:9b:72:38:b0:68:
                               e1:c0:08:f4:02:91:81:a8:fa:51:9e:53:0d:03:b3:6b:0e:e2:
                               62:80:ef:2a:a0:cb:9b:9b:91:21:7c:df:fe:6a:38:cc:03:36:
                               9c:fc

                         Pki realm name: -abc
                         Certificate file name: cmp_ir.cer
                         Certificate peer name: -




