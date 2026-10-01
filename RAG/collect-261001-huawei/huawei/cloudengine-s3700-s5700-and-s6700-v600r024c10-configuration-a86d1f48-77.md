---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-77
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [9364, 9504]
sha256: 085a41c91690e2112409573253d1da4c184cca076f5f3f3ec3bccc7d2789ebf3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         When applying for an SM digital envelope offline, you need to apply for a PEM-encoded
                         digital envelope that complies with GM/T 0009-2012 to ensure successful import.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create a PKI realm and enter its view, or enter the view of an existing PKI realm.
                  pki realm realm-name

                  By default, the system has a PKI realm named default. This realm can be modified
                  but cannot be deleted.
                  A PKI realm is locally available. It is unavailable to CAs or other devices. Each PKI
                  realm has its own parameter settings.
         Step 3 Specify the PKI entity that applies for a certificate.
                  entity entity-name

                  The PKI entity specified by entity-name must have been created using the pki
                  entity command.
         Step 4 Configure the key pair used to apply for a certificate in offline mode as required.
                  ●      This command is added to configure the RSA key pair used in offline
                         certificate application.
                         rsa local-key-pair key-name

                  ●      This command is added to configure the SM2 key pair used in offline
                         certificate application.
                         sm2 local-key-pair key-name

                  ●      This command is added to configure the ECC key pair used in offline
                         certificate application.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  175
Security Configuration
Security Configuration                                                                          10 PKI Configuration

                         ecc local-key-pair key-name

         Step 5 Configure the digest algorithm used to sign certificate enrollment requests.
                  enrollment-request signature message-digest-method { md5 | sha1| sha-256 | sha-384 | sha-512 | sm3 }

                  By default, the digest algorithm used to sign certificate enrollment requests is
                  SHA-256.
                  The digest algorithm used on a PKI entity must be the same as that used on the
                  CA server. Note that the MD5 and SHA1 algorithms are insecure, so you are
                  advised to use the more secure SHA2 algorithms (SHA-256, SHA-384, and
                  SHA-512).
                  In a PKI realm, if the SM2 key pair is used to apply for a certificate in offline
                  mode, the digest algorithm used to sign certificate enrollment requests must be
                  configured as SM3. If the RSA or ECC key pair is used to apply for a certificate in
                  offline mode, the digest algorithm used to sign certificate enrollment requests
                  cannot be configured as SM3. Otherwise, offline certificate application fails.

                          NOTE

                         For security purpose,you are not advised to use the weak security algorithm or weak
                         security protocols provided by this feature. If you need to use the weak security algorithm
                         or protocols, run the install feature-software WEAKEA command to install the weak
                         security algorithm or protocol feature package WEAKEA. By default, the device provides the
                         weak security algorithm or protocol feature package WEAKEA. For details about how to
                         install or uninstall the feature package, see "Upgrade Maintenance Configuration" in CLI
                         Configuration Guide > System Management Configuration.

         Step 6 Optional: Configure the certificate public key usage attribute.
                  key-usage { signature | cipher }
                  quit

         Step 7 Configure the file format in which the device stores the certificate and certificate
                enrollment request in the system view.
                  pki file-format { der | pem }

                  By default, the device stores the certificate and certificate enrollment request into
                  a PEM file.
         Step 8 Set parameters to save certificate enrollment information into a file in PKCS#10
                format.
                  pki enroll-certificate realm realm-name pkcs10 [ filename filename ] [ password password ]

                  The challenge password used on a PKI entity must be the same as that configured
                  on the CA server. If the CA server does not require a challenge password, this
                  challenge password does not need to be configured.
         Step 9 Enable the device to send the certificate enrollment request file to the CA in out-
                of-band mode (for example, by disk or email) to apply for a local certificate.

                  ----End

Verifying the Configuration
                  ●      Run the display pki realm [ realm-name ] command to check the PKI realm
                         information.
                  ●      Run the display pki cert-req filename file-name command to check the
                         content of the certificate enrollment request file.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       176
Security Configuration
Security Configuration                                                                          10 PKI Configuration


10.6.3 Downloading a Local Certificate

Context
                  A device downloads its local certificate using one of following methods, depending
                  on the service type provided by the CA server:
                  ●      LDAP mode: The device's local certificate is downloaded from the server
                         where the certificate is stored using LDAP. After the download is completed,
                         the device automatically saves the local certificate to the flash:/pki/public
                         directory on the device.
                  ●      Out-of-band mode: The device's local certificate is downloaded in out-of-
                         band mode (for example, by disk or email) and then uploaded to the device
                         storage.


Procedure
                  ●      Download the local certificate through LDAP.
                         system-view
                         pki ldap-server-template template-name attribute attr-value save-name dn dn-value

                  ●      Download the local certificate in out-of-band mode (for example, by disk or
                         email).

                         After you obtain the local certificate in out-of-band mode, manually upload it
                         to the device storage. You can also download the local certificate through the
                         management PC and then upload it to the device storage through FTP SFTP.

                  ----End


Verifying the Configuration
                  ●      Run the display pki credential-storage-path command to check the default
                         directory where a certificate is stored.
                  ●      Run the dir command in the user view to check the local certificate file in the
                         device storage.

10.6.4 Installing a Local Certificate

Context
                  After obtaining a local certificate in out-of-band mode (for example, by disk or
                  email), you need to upload the local certificate to the specified storage directory
                  on the device. Before installing a local certificate, you need to upload the
                  certificate to flash:/pki/public.

