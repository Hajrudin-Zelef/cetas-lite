---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-74
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [8967, 9076]
sha256: 557426175694903c6e779818575dc04c2488a48030135763cb9d8493863b5b96
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                                 pki sm2 local-key-pair create    The created SM2 key pair can be
                                 key-name [ exportable ]          exported only when the
                                                                  exportable parameter is specified
                                                                  in the command.

                                 pki ecc curve-name               The created ECC key pair can be
                                 { prime256v1 | ec192wapi |       exported only when the
                                 secp384r1 | secp521r1 }          exportable parameter is specified
                                 local-key-pair create key-       in the command.
                                 name [ exportable ]
                   Import an     pki import rsa-key-pair          If the exclude-cert parameter is
                   RSA/SM2/      keyname [ exclude-cert ]         specified in the command, the
                   ECC key       { pem | pkcs12 } filename        certificate in the file will not be
                   pair.         [ exportable ] [ password        imported to the device.
                                 password ]
                                 pki import sm2-key-pair          The imported SM2 key pair can be
                                 keyname pem filename             exported only when the
                                 [ exportable ] signkey           exportable parameter is specified
                                 signkey-name [ certificate       in the command.
                                 certificate-name ]
                                 pki import ecc-key-pair          If the exclude-cert parameter is
                                 keyname [ exclude-cert ]         specified in the command, the
                                 { pem | pkcs12 } filename        certificate in the file will not be
                                 [ exportable ] [ password        imported to the device.
                                 password ]


                  ----End

Verifying the Configuration
                  ●      Run the display pki rsa local-key-pair { pem | pkcs12 } file-name
                         [ password password ] command to check the RSA key pair information.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               168
Security Configuration
Security Configuration                                                             10 PKI Configuration


                  ●      Run the display pki rsa local-key-pair [ name key-name ] public command
                         to check the RSA public key information.
                  ●      Run the display pki sm2 local-key-pair [ name key-name ] public command
                         to check information about SM2 key pairs and public keys.
                  ●      Run the display pki ecc local-key-pair [ name key-name ] public command
                         to check information about ECC key pairs and public keys.

Follow-up Procedure

                  Table 10-6
                   Operatio      Command                          Description
                   n

                   Export        pki export rsa-key-pair key-     To back up RSA key pairs or use
                   the           name [ and-certificate           them on other devices, run this
                   specified     certificate-name ] { pem         command in the system view to
                   RSA/SM2/      filename [ aes ] | pkcs12        export the specified RSA key pair
                   ECC key       filename } password password     to the device memory. In addition
                   pair                                           to the RSA key pair, its associated
                                                                  certificate and certificate chain will
                                                                  also be exported. You can then
                                                                  obtain the RSA key pair using FTP
                                                                  or SFTP.

                                 pki export sm2-key-pair          To back up SM2 key pairs or use
                                 keyname pem filename             them on other devices, run this
                                 [ password password ]            command in the system view to
                                                                  export the specified SM2 key pair
                                                                  to the device memory. You can
                                                                  then obtain the SM2 key pair
                                                                  using FTP or SFTP.

                                 pki export ecc-key-pair          To back up ECC key pairs or use
                                 keyname [ and-certificate        them on other devices, run this
                                 certificate-name ] { pem         command in the system view to
                                 filename [ aes ] | pkcs12        export the specified ECC key pair
                                 filename } password password     to the device memory. In addition
                                                                  to the ECC key pair, its associated
                                                                  certificate and certificate chain will
                                                                  also be exported. You can then
                                                                  obtain the ECC key pair using FTP
                                                                  or SFTP.

                   Destroy a     pki rsa local-key-pair destroy   When RSA key pairs are leaked,
                   specified     key-name                         damaged, lost, or unused, run this
                   RSA/SM2/                                       command in the system view to
                   ECC key                                        destroy a specified RSA key pair.
                   pair                                           After this command is executed,
                                                                  the system deletes the specified
                                                                  RSA key pair.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            169
Security Configuration
Security Configuration                                                              10 PKI Configuration


                   Operatio       Command                          Description
                   n

                                  pki sm2 local-key-pair           When SM2 key pairs are leaked,
                                  destroy key-name                 damaged, lost, or unused, run this
                                                                   command in the system view to
                                                                   destroy a specified RSA key pair.
                                                                   After this command is executed,
                                                                   the system deletes the specified
                                                                   SM2 key pair.

