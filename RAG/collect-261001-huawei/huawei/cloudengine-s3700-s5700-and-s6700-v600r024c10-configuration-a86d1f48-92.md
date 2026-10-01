---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-92
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [11395, 11510]
sha256: 69b81e55cadb953b4cb7d4ba348eafae054186fb61c7de2512448b876f269fca
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  1.     Export the RSA key pair file and certificates of DeviceA to the storage card.
                  2.     Save the RSA key pair file and certificates in DeviceA's storage card to the PC
                         using SFTP.
                  3.     Save DeviceA's RSA key pair file and certificates on the PC to DeviceB's
                         storage card using SFTP.
                  4.     Import the RSA key pair file and certificates in DeviceB's storage card to its
                         memory.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      209
Security Configuration
Security Configuration                                                                                   10 PKI Configuration


Procedure
         Step 1 Export DeviceA's RSA key pair file and certificates.
                  # Export the RSA key pair rsa_key and corresponding certificate cer_test.cer to
                  the test02.pem file in PEM format and set the encryption mode to AES.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] pki export rsa-key-pair rsa_key and-certificate cer_test.cer pem test02.pem aes password
                  YsHsjx_202206
                   Warning: Exporting the key pair impose security risks, are you sure you want to
                   export it? [y/n]:y
                   Info: Succeeded in exporting the RSA key pair in PEM format.

                          NOTE

                         When the pki rsa local-key-pair create command is executed on DeviceA to create an RSA
                         key pair, the RSA key pair cannot be exported if the exportable parameter is not
                         configured.

                  # Check whether the test02.pem file exists in the storage card.
                  [DeviceA] quit
                  <DeviceA> dir flash:/pki/public/
                  Directory of flash:/pki/public/

                   Idx Attr    Size(Byte) Date    Time      FileName
                     0 -rw-       3,016 Jun 15 2017 18:48:26 test02.pem

                  1,179,616 KB total (434,592 KB free)

         Step 2 Save the test02.pem file in DeviceA's storage card to the PC using SFTP.
         Step 3 Save the test02.pem file on the PC to flash:/pki/public on the storage card of
                DeviceB using SFTP. The directory must be the same as that on DeviceA.
                Otherwise, the import fails.
         Step 4 Import the RSA key pair file and certificates of DeviceA to DeviceB.
                  Import the RSA key pair file test02.pem in PEM format. In the system, the RSA
                  key pair is named rsakey, has password YsHsjx_202206, and is marked
                  exportable.
                  After test02.pem is imported, test02.pem in the storage card is deleted by
                  default. If test02.pem does not need to be deleted, select N as prompted to keep
                  it.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceB
                  [DeviceB] pki import rsa-key-pair rsakey pem test02.pem exportable password YsHsjx_202206
                   Info: Succeeded in importing the RSA key pair in PEM format.
                   Warning: The file in the flash will be deleted. Please select 'N' if you want to keep it. Please select [Y/N]:y
                   Info: Delete Success.

                  After the test02.pem file is imported to DeviceB, the RSA key pair rsakey, local
                  certificate rsakey_local.cer, and CA certificate rsakey_ca.cer are generated in the
                  memory of DeviceB.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                    210
Security Configuration
Security Configuration                                                                          10 PKI Configuration


                          NOTE

                         If the test02.pem file exported from DeviceA does not contain a CA certificate, no CA
                         certificate is generated in the memory of DeviceB after this file is imported to DeviceB. If a
                         CA certificate needs to be imported, run the pki export-certificate ca and pki import-
                         certificate ca commands by following the preceding steps.

                  ----End

Verifying the Configuration
                  1.     Run the display pki rsa local-key-pair command to check information about
                         the RSA key pair file imported to the memory.
                         [DeviceB] display pki rsa local-key-pair name rsakey public
                         Info: It will take a few seconds or more. Please wait a moment.
                         Total Number: 1
                         =====================================================
                         Time of Key pair created: 10:40:22 2021/3/13
                         Key Name: rsakey
                         Key Modulus: 3072 bits
                         Key Exportable: Yes
                         =====================================================
                         RSA Public-Key: (3072 bits)
                         Modulus:
                            00:9d:e2:3b:3b:d9:19:48:3a:62:59:11:c4:af:08:
                            03:dd:9c:4a:61:e8:ed:a3:4b:a2:44:7f:a6:ea:10:
                            12:04:8f:93:f2:ab:dc:09:f9:bc:e5:6b:4c:d3:29:
                            f6:22:9e:da:83:bf:17:b2:8e:6b:65:6c:17:7e:83:
                            dc:8e:33:1f:33:2d:96:4f:3d:ed:03:6d:91:45:47:
                            49:79:8b:89:8a:7b:e5:f8:12:c0:41:45:77:ff:30:
                            4c:a1:d4:f2:d0:9f:02:84:82:6d:02:10:bd:f1:5a:
                            64:d0:8d:21:aa:a5:e6:61:ee:bb:55:a1:99:3f:ad:
                            fb:6c:13:c9:dd:23:c6:ab:02:24:07:e4:76:4b:ef:
                            3e:fa:56:31:80:b2:75:a2:b5:cc:12:0b:33:0a:e7:
                            19:ed:6b:36:93:9f:78:e1:37:13:e2:b5:47:6f:d1:
                            f1:7c:d8:01:49:f6:82:d9:3a:d6:1a:fd:bb:c4:71:
                            05:fd:a4:ea:73:5b:db:b5:1a:2b:a5:e3:e2:78:b4:
                            ec:9b:92:36:72:35:4f:7b:cc:05:91:db:14:1f:da:
                            c5:22:89:f0:64:4a:76:b3:27:69:cf:b6:a6:1d:bd:
                            ec:4c:24:0d:9e:ff:27:46:94:2e:b0:68:61:c6:ce:
                            bd:e3:b0:4b:26:66:ee:f1:8a:3f:8c:30:7f:6f:bd:
                            77:d1
                         Exponent: 65537 (0x10001)

