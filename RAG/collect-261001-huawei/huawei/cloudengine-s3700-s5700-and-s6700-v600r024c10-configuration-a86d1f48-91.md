---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-91
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [11228, 11394]
sha256: d34d55d60982212c118700f21e3926ab2f4732aa2ef55f8ac92aa4820f5ce821
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Procedure
         Step 1 Configure the action in the default certificate attribute-based access control policy
                to deny, indicating that certificates matching specified certificate attributes are not
                allowed to pass authentication.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] pki certificate access-control-policy default deny

         Step 2 Create a certificate attribute group named group.
                  [DeviceA] pki certificate attribute-group group

         Step 3 Specify the issuer name networkb_ca and subject name cert_ca.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       206
Security Configuration
Security Configuration                                                                            10 PKI Configuration

                  [DeviceA-pki-attribute-group] attribute 1 issuer-name dn equ networkb_ca
                  [DeviceA-pki-attribute-group] attribute 2 subject-name dn equ cert_ca
                  [DeviceA-pki-attribute-group] quit

         Step 4 Create a certificate attribute-based access control policy named policy.
                  [DeviceA] pki certificate access-control-policy name policy

         Step 5 Configure a certificate attribute-based control rule, allowing certificates matching
                the specified attributes in the certificate attribute group to pass the
                authentication.
                  [DeviceA-pki-access-policy] rule 1 permit group
                  [*DeviceA-pki-access-policy] quit

                  ----End

Verifying the Configuration
                  After the configurations are complete, only devices whose certificate issuer name
                  is networkb_ca and subject name is cert_ca can set up an tunnel with DeviceA.

Configuration Scripts
                  #
                  sysname DeviceA
                  #
                  pki certificate access-control-policy default deny
                  #
                  pki certificate attribute-group group
                   attribute 1 issuer-name dn equ networkb_ca
                   attribute 2 subject-name dn equ cert_ca
                  #
                  pki certificate access-control-policy name policy
                   rule 1 permit group
                  #
                  return



10.10 Importing and Exporting a Certificate

10.10.1 Importing the RSA, ECC, or SM2 Key Pair and
Certificate of Another Device
Context
                  On the live network, if a device does not apply for a certificate, you can import the
                  RSA, ECC, or SM2 key pair and certificate of another device.

Procedure
         Step 1 On DeviceA, export its RSA, ECC, or SM2 key pair and certificate.
                  pki export rsa-key-pair keyname [ and-certificate certificate-name ] { pem filename [ aes ] | pkcs12
                  filename } password password
                  pki export sm2-key-pair keyname pem filename [ password password ]
                  pki export ecc-key-pair keyname [ and-certificate certificate-name ] { pem filename [ aes ] | pkcs12
                  filename } password password

         Step 2 Save the key file in the storage of DeviceA to a PC using a file transfer protocol
                such as SFTP.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              207
Security Configuration
Security Configuration                                                                            10 PKI Configuration


         Step 3 Save the key file on the PC to the storage of DeviceB using a file transfer protocol
                such as SFTP.
         Step 4 On DeviceB, import the RSA, ECC, or SM2 key pair and certificate of DeviceA.
                  pki import rsa-key-pair keyname [ exclude-cert ] { pem | pkcs12 } filename [ exportable ] [ password
                  password ]
                  pki import sm2-key-pair keyname pem filename [ exportable ] signkey signkey-name [ certificate
                  certificate-name ]
                  pki import ecc-key-pair keyname [ exclude-cert ] { pem | pkcs12 } filename [ exportable ] [ password
                  password ]

                  ----End

10.10.2 Importing a Peer Certificate
Context
                  Importing a peer certificate applies to large-scale networks.
                  If the imported peer certificate is no longer needed, release the certificate.

Procedure
                  ●      Import a peer certificate to the device memory.
                         system-view
                         pki import-certificate peer peer-name { der | pem | pkcs12 } filename filename [ cert-name cert-
                         name ] [ no-check-same-name ]
                  ●      Release a peer certificate.
                         system-view
                         pki release-certificate peer { name peer-name | all }

                  ----End

Verifying the Configuration
                  Run the display pki peer-certificate { name peer-name | all } command to check
                  the imported peer certificates.

10.10.3 Exporting a Certificate
Context
                  The CA certificate, local certificate, and OCSP server certificate of a device can be
                  exported for use on another device. You can run the commands in "Procedure" to
                  export certificates to the device storage, and then transfer the certificates to
                  another device through FTP or SFTP.

Procedure
                  ●      In the system view, export the CA certificate.
                         pki export-certificate ca realm realm-name { pem | pkcs12 }

                  ●      In the system view, export the default CA certificate.
                         pki export-certificate default ca filename file-name

                  ●      In the system view, export the local certificate.
                         pki export-certificate local realm realm-name { pem | pkcs12 }


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           208
Security Configuration
Security Configuration                                                                      10 PKI Configuration


                  ●      In the system view, export the OCSP server certificate.
                         pki export-certificate ocsp realm realm-name { pem | pkcs12 }

                  ----End

10.10.4 Example for Manually Importing the RSA Key Pair File
and Certificates of Another Device

Networking Requirements
                  As shown in Figure 10-17, DeviceA is deployed at the border of an enterprise
                  network as the egress gateway. DeviceA has applied for a local certificate from the
                  CA server on the public network.

                  The enterprise wants to replace outdated DeviceA with DeviceB. However,
                  DeviceA's RSA key pair file and certificates can only be manually imported to
                  DeviceB because of network issues.


                  Figure 10-17 Manually importing the RSA key pair and certificates of another
                  device




                          NOTE

                         If DeviceA is a non-Huawei device, see its configuration manual for related commands.


Configuration Roadmap
                  The configuration roadmap is as follows:

